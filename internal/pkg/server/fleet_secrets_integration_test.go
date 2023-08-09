// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofrs/uuid"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elastic/fleet-server/v7/internal/pkg/api"
	"github.com/elastic/fleet-server/v7/internal/pkg/apikey"
	"github.com/elastic/fleet-server/v7/internal/pkg/bulk"
	"github.com/elastic/fleet-server/v7/internal/pkg/dl"
	"github.com/elastic/fleet-server/v7/internal/pkg/model"
)

type SecretResponse struct {
	ID string
}

func createSecret(t *testing.T, ctx context.Context, bulker bulk.Bulk) string {
	t.Log("Setup secret for fleet integration test")
	esClient := bulker.Client()

	req, err := http.NewRequestWithContext(ctx, "POST", "/_fleet/secret/", bytes.NewBuffer([]byte("{\"value\":\"secret_value\"}")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// kibana_system:changeme base64 encoded
	req.Header.Set("Authorization", "Basic a2liYW5hX3N5c3RlbTpjaGFuZ2VtZQ==")
	if err != nil {
		t.Fatal(err)
	}

	res, err := esClient.Perform(req)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()
	var resp SecretResponse

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("secret id created: %s", resp.ID)

	return resp.ID
}

func createAgentPolicyWithSecrets(t *testing.T, ctx context.Context, bulker bulk.Bulk, secretID string, secretRef string) string {
	policyID := uuid.Must(uuid.NewV4()).String()
	var policyDataWSecrets = map[string]interface{}{
		"name": "TestPolicy " + policyID,
		"outputs": map[string]interface{}{
			"default": map[string]string{
				"type": "elasticsearch",
			}},
		"output_permissions": map[string]interface{}{
			"default": map[string]string{},
		},
		"inputs": []map[string]string{{
			"type":               "fleet-server",
			"package_var_secret": secretRef,
		}},
		"secret_references": []map[string]string{{
			"id": secretID,
		}},
	}
	policyDataJSON, _ := json.Marshal(policyDataWSecrets)

	_, err := dl.CreatePolicy(ctx, bulker, model.Policy{
		PolicyID:           policyID,
		RevisionIdx:        1,
		DefaultFleetServer: true,
		Data:               policyDataJSON,
	})
	if err != nil {
		t.Fatal(err)
	}

	// In order to create a functional enrollement token we need to use the ES endpoint to create a new api key
	// then add the key (id/value) to the enrollment index
	esCfg := elasticsearch.Config{
		Username: "elastic",
		Password: "changeme",
	}
	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		t.Fatal(err)
	}
	key, err := apikey.Create(ctx, es, "default", "", "true", []byte(`{
	    "fleet-apikey-enroll": {
		"cluster": [],
		"index": [],
		"applications": [{
		    "application": "fleet",
		    "privileges": ["no-privileges"],
		    "resources": ["*"]
		}]
	    }
	}`), map[string]interface{}{
		"managed_by": "fleet",
		"managed":    true,
		"type":       "enroll",
		"policy_id":  policyID,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = dl.CreateEnrollmentAPIKey(ctx, bulker, model.EnrollmentAPIKey{
		Name:     "Default",
		APIKey:   key.Key,
		APIKeyID: key.ID,
		PolicyID: policyID,
		Active:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	return key.Token()
}

func Test_Agent_Policy_Secrets(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start test server
	srv, err := startTestServer(t, ctx)
	require.NoError(t, err)

	// create secret with kibana_system user
	secretID := createSecret(t, ctx, srv.bulker)
	secretRef := fmt.Sprintf("$co.elastic.secret{%s}", secretID)

	// create agent policy with secret reference
	enrollKey := createAgentPolicyWithSecrets(t, ctx, srv.bulker, secretID, secretRef)
	cli := cleanhttp.DefaultClient()
	// enroll an agent
	t.Log("Enroll an agent")
	req, err := http.NewRequestWithContext(ctx, "POST", srv.baseURL()+"/api/fleet/agents/enroll", strings.NewReader(enrollBody))
	require.NoError(t, err)
	req.Header.Set("Authorization", "ApiKey "+enrollKey)
	req.Header.Set("User-Agent", "elastic agent "+serverVersion)
	req.Header.Set("Content-Type", "application/json")
	res, err := cli.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, res.StatusCode)
	t.Log("Agent enrollment successful")
	p, _ := io.ReadAll(res.Body)
	res.Body.Close()
	var obj map[string]interface{}
	err = json.Unmarshal(p, &obj)
	require.NoError(t, err)

	item := obj["item"]
	mm, ok := item.(map[string]interface{})
	require.True(t, ok, "expected attribute item to be an object")
	id := mm["id"]
	str, ok := id.(string)
	require.True(t, ok, "expected attribute id to be a string")

	apiKey := mm["access_api_key"]
	key, ok := apiKey.(string)
	require.True(t, ok, "expected attribute apiKey to be a string")

	// checkin
	t.Logf("Fake a checkin for agent %s", str)
	req, err = http.NewRequestWithContext(ctx, "POST", srv.baseURL()+"/api/fleet/agents/"+str+"/checkin", strings.NewReader(checkinBody))
	require.NoError(t, err)
	req.Header.Set("Authorization", "ApiKey "+key)
	req.Header.Set("User-Agent", "elastic agent "+serverVersion)
	req.Header.Set("Content-Type", "application/json")
	res, err = cli.Do(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, res.StatusCode)
	t.Log("Checkin successful, verify body")
	var checkinResponse api.CheckinResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&checkinResponse)
	res.Body.Close()

	require.NoError(t, err)

	// expect 1 POLICY_CHANGE action
	assert.Equal(t, 1, len(*checkinResponse.Actions))
	var actionDataRaw interface{}
	for _, action := range *checkinResponse.Actions {
		actionDataRaw = action.Data
		assert.Equal(t, "POLICY_CHANGE", action.Type)
	}

	actionData, ok := actionDataRaw.(map[string]interface{})
	require.True(t, ok, "expected attribute action.Data to be an object")

	policy, ok := actionData["policy"].(map[string]interface{})
	require.True(t, ok, "expected attribute policy to be an object")
	inputs, ok := policy["inputs"].([]interface{})
	require.True(t, ok, "expected attribute inputs to be an array")

	input, ok := inputs[0].(map[string]interface{})
	require.True(t, ok, "expected first input to be an object")

	// expect secret reference replaced with secret value
	assert.Equal(t, map[string]interface{}{
		"package_var_secret": "secret_value",
		"type":               "fleet-server",
	}, input)
}
