// Code generated by schema-generate. DO NOT EDIT.

package model

import (
	"encoding/json"
)

// Root
type Root interface{}

type ESInitializer interface {
	ESInitialize(id string, seqno, version int64)
}

type ESDocument struct {
	Id      string `json:"-"`
	Version int64  `json:"-"`
	SeqNo   int64  `json:"-"`
}

func (d *ESDocument) ESInitialize(id string, seqno, version int64) {
	d.Id = id
	d.SeqNo = seqno
	d.Version = version
}

// Action An Elastic Agent action
type Action struct {
	ESDocument

	// The unique identifier for the Elastic Agent action. There could be multiple documents with the same action_id if the action is split into two separate documents.
	ActionID string `json:"action_id,omitempty"`

	// The Agent IDs the action is intended for. No support for json.RawMessage with the current generator. Could be useful to lazy parse the agent ids
	Agents []string `json:"agents,omitempty"`

	// The opaque payload.
	Data json.RawMessage `json:"data,omitempty"`

	// The action expiration date/time
	Expiration string `json:"expiration,omitempty"`

	// The input type the actions should be routed to.
	InputType string `json:"input_type,omitempty"`

	// The minimum time (in seconds) provided for an action execution when scheduled by fleet-server.
	MinimumExecutionDuration int64 `json:"minimum_execution_duration,omitempty"`

	// The action start date/time
	StartTime string `json:"start_time,omitempty"`

	// The optional action timeout in seconds
	Timeout int64 `json:"timeout,omitempty"`

	// Date/time the action was created
	Timestamp string `json:"@timestamp,omitempty"`

	// The action type. INPUT_ACTION is the value for the actions that suppose to be routed to the endpoints/beats.
	Type string `json:"type,omitempty"`

	// The ID of the user who created the action.
	UserID string `json:"user_id,omitempty"`
}

// ActionData The opaque payload.
type ActionData struct {
}

// ActionResponse The custom action response payload.
type ActionResponse struct {
}

// ActionResult An Elastic Agent action results
type ActionResult struct {
	ESDocument

	// The opaque payload.
	ActionData json.RawMessage `json:"action_data,omitempty"`

	// The action id.
	ActionID string `json:"action_id,omitempty"`

	// The input type of the original action.
	ActionInputType string `json:"action_input_type,omitempty"`

	// The custom action response payload.
	ActionResponse json.RawMessage `json:"action_response,omitempty"`

	// The agent id.
	AgentID string `json:"agent_id,omitempty"`

	// Date/time the action was completed
	CompletedAt string `json:"completed_at,omitempty"`

	// The opaque payload.
	Data json.RawMessage `json:"data,omitempty"`

	// The action error message.
	Error string `json:"error,omitempty"`

	// Date/time the action was started
	StartedAt string `json:"started_at,omitempty"`

	// Date/time the action was created
	Timestamp string `json:"@timestamp,omitempty"`
}

// Agent An Elastic Agent that has enrolled into Fleet
type Agent struct {
	ESDocument

	// ID of the API key the Elastic Agent must used to contact Fleet Server
	AccessAPIKeyID string `json:"access_api_key_id,omitempty"`

	// The last acknowledged action sequence number for the Elastic Agent
	ActionSeqNo []int64 `json:"action_seq_no,omitempty"`

	// Active flag
	Active bool           `json:"active"`
	Agent  *AgentMetadata `json:"agent,omitempty"`

	// API key the Elastic Agent uses to authenticate with elasticsearch
	DefaultAPIKey string `json:"default_api_key,omitempty"`

	// Default API Key History
	DefaultAPIKeyHistory []DefaultAPIKeyHistoryItems `json:"default_api_key_history,omitempty"`

	// ID of the API key the Elastic Agent uses to authenticate with elasticsearch
	DefaultAPIKeyID string `json:"default_api_key_id,omitempty"`

	// Date/time the Elastic Agent enrolled
	EnrolledAt string `json:"enrolled_at"`

	// Date/time the Elastic Agent checked in last time
	LastCheckin string `json:"last_checkin,omitempty"`

	// Last checkin status
	LastCheckinStatus string `json:"last_checkin_status,omitempty"`

	// Date/time the Elastic Agent was last updated
	LastUpdated string `json:"last_updated,omitempty"`

	// Local metadata information for the Elastic Agent
	LocalMetadata json.RawMessage `json:"local_metadata,omitempty"`

	// Packages array
	Packages []string `json:"packages,omitempty"`

	// The current policy coordinator for the Elastic Agent
	PolicyCoordinatorIdx int64 `json:"policy_coordinator_idx,omitempty"`

	// The policy ID for the Elastic Agent
	PolicyID string `json:"policy_id,omitempty"`

	// The policy output permissions hash
	PolicyOutputPermissionsHash string `json:"policy_output_permissions_hash,omitempty"`

	// The current policy revision_idx for the Elastic Agent
	PolicyRevisionIdx int64 `json:"policy_revision_idx,omitempty"`

	// Shared ID
	SharedID string `json:"shared_id,omitempty"`

	// User provided tags for the Elastic Agent
	Tags []string `json:"tags,omitempty"`

	// Type
	Type string `json:"type"`

	// Date/time the Elastic Agent unenrolled
	UnenrolledAt string `json:"unenrolled_at,omitempty"`

	// Reason the Elastic Agent was unenrolled
	UnenrolledReason string `json:"unenrolled_reason,omitempty"`

	// Date/time the Elastic Agent unenrolled started
	UnenrollmentStartedAt string `json:"unenrollment_started_at,omitempty"`

	// Date/time the Elastic Agent was last updated
	UpdatedAt string `json:"updated_at,omitempty"`

	// Date/time the Elastic Agent started the current upgrade
	UpgradeStartedAt string `json:"upgrade_started_at,omitempty"`

	// Upgrade status
	UpgradeStatus string `json:"upgrade_status,omitempty"`

	// Date/time the Elastic Agent was last upgraded
	UpgradedAt string `json:"upgraded_at,omitempty"`

	// User provided metadata information for the Elastic Agent
	UserProvidedMetadata json.RawMessage `json:"user_provided_metadata,omitempty"`
}

// AgentMetadata An Elastic Agent metadata
type AgentMetadata struct {

	// The unique identifier for the Elastic Agent
	ID string `json:"id"`

	// The version of the Elastic Agent
	Version string `json:"version"`
}

// Artifact An artifact served by Fleet
type Artifact struct {
	ESDocument

	// Encoded artifact data
	Body json.RawMessage `json:"body"`

	// Name of compression algorithm applied to artifact
	CompressionAlgorithm string `json:"compression_algorithm,omitempty"`

	// Timestamp artifact was created
	Created string `json:"created"`

	// SHA256 of artifact before encoding has been applied
	DecodedSha256 string `json:"decoded_sha256,omitempty"`

	// Size of artifact before encoding has been applied
	DecodedSize int64 `json:"decoded_size,omitempty"`

	// SHA256 of artifact after encoding has been applied
	EncodedSha256 string `json:"encoded_sha256,omitempty"`

	// Size of artifact after encoding has been applied
	EncodedSize int64 `json:"encoded_size,omitempty"`

	// Name of encryption algorithm applied to artifact
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`

	// Human readable artifact identifier
	Identifier string `json:"identifier"`

	// Name of the package that owns this artifact
	PackageName string `json:"package_name,omitempty"`
}

// Body Encoded artifact data
type Body struct {
}

// Data The opaque payload.
type Data struct {
}

// DefaultAPIKeyHistoryItems
type DefaultAPIKeyHistoryItems struct {

	// API Key identifier
	ID string `json:"id,omitempty"`

	// Date/time the API key was retired
	RetiredAt string `json:"retired_at,omitempty"`
}

// EnrollmentAPIKey An Elastic Agent enrollment API key
type EnrollmentAPIKey struct {
	ESDocument

	// Api key
	APIKey string `json:"api_key"`

	// The unique identifier for the enrollment key, currently xid
	APIKeyID string `json:"api_key_id"`

	// True when the key is active
	Active    bool   `json:"active,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ExpireAt  string `json:"expire_at,omitempty"`

	// Enrollment key name
	Name      string `json:"name,omitempty"`
	PolicyID  string `json:"policy_id,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// HostMetadata The host metadata for the Elastic Agent
type HostMetadata struct {

	// The architecture for the Elastic Agent
	Architecture string `json:"architecture"`

	// The ID of the host
	ID string `json:"id"`

	// The IP addresses of the Elastic Agent
	Ip []string `json:"ip,omitempty"`

	// The hostname of the Elastic Agent
	Name string `json:"name"`
}

// LocalMetadata Local metadata information for the Elastic Agent
type LocalMetadata struct {
}

// Policy A policy that an Elastic Agent is attached to
type Policy struct {
	ESDocument

	// The coordinator index of the policy
	CoordinatorIdx int64 `json:"coordinator_idx"`

	// The opaque payload.
	Data json.RawMessage `json:"data"`

	// True when this policy is the default policy to start Fleet Server
	DefaultFleetServer bool `json:"default_fleet_server"`

	// The ID of the policy
	PolicyID string `json:"policy_id"`

	// The revision index of the policy
	RevisionIdx int64 `json:"revision_idx"`

	// Date/time the policy revision was created
	Timestamp string `json:"@timestamp,omitempty"`

	// Timeout (seconds) that an Elastic Agent should be un-enrolled.
	UnenrollTimeout int64 `json:"unenroll_timeout,omitempty"`
}

// PolicyLeader The current leader Fleet Server for a policy
type PolicyLeader struct {
	ESDocument
	Server *ServerMetadata `json:"server"`

	// Date/time the leader was taken or held
	Timestamp string `json:"@timestamp,omitempty"`
}

// Server A Fleet Server
type Server struct {
	ESDocument
	Agent  *AgentMetadata  `json:"agent"`
	Host   *HostMetadata   `json:"host"`
	Server *ServerMetadata `json:"server"`

	// Date/time the server was updated
	Timestamp string `json:"@timestamp,omitempty"`
}

// ServerMetadata A Fleet Server metadata
type ServerMetadata struct {

	// The unique identifier for the Fleet Server
	ID string `json:"id"`

	// The version of the Fleet Server
	Version string `json:"version"`
}

// UserProvidedMetadata User provided metadata information for the Elastic Agent
type UserProvidedMetadata struct {
}
