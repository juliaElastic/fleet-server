// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildlimits/buildlimits.go - DO NOT EDIT.

package config

import (
	"testing"

	testlog "github.com/elastic/fleet-server/v7/internal/pkg/testing/log"

	"github.com/stretchr/testify/require"
	"github.com/rs/zerolog/log"
)

func TestLoadLimits(t *testing.T) {
	testCases := []struct {
		Name                 string
		ConfiguredAgentLimit int
		ExpectedAgentLimit   int
	}{
		{"few agents", 5, 49},
		{"512", 512, 4999},
		{"precise", 7499, 7499},
		{"10k", 10050, 29999},
		{"30k", 30050, 49999},
		{"50k", 50050, 99999},
		{"above max", 100001, int(getMaxInt())},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			log.Logger = testlog.SetLogger(t)
			l := loadLimitsForAgents(tc.ConfiguredAgentLimit)

			require.Equal(t, tc.ExpectedAgentLimit, l.Agents.Max)
		})
	}
}
