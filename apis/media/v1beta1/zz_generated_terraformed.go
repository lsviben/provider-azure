/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Asset
func (mg *Asset) GetTerraformResourceType() string {
	return "azurerm_media_asset"
}

// GetConnectionDetailsMapping for this Asset
func (tr *Asset) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Asset
func (tr *Asset) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Asset
func (tr *Asset) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Asset
func (tr *Asset) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Asset
func (tr *Asset) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Asset
func (tr *Asset) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Asset using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Asset) LateInitialize(attrs []byte) (bool, error) {
	params := &AssetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Asset) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this AssetFilter
func (mg *AssetFilter) GetTerraformResourceType() string {
	return "azurerm_media_asset_filter"
}

// GetConnectionDetailsMapping for this AssetFilter
func (tr *AssetFilter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AssetFilter
func (tr *AssetFilter) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AssetFilter
func (tr *AssetFilter) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AssetFilter
func (tr *AssetFilter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AssetFilter
func (tr *AssetFilter) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AssetFilter
func (tr *AssetFilter) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AssetFilter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AssetFilter) LateInitialize(attrs []byte) (bool, error) {
	params := &AssetFilterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AssetFilter) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this ContentKeyPolicy
func (mg *ContentKeyPolicy) GetTerraformResourceType() string {
	return "azurerm_media_content_key_policy"
}

// GetConnectionDetailsMapping for this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"policy_option[*].fairplay_configuration[*].ask": "spec.forProvider.policyOption[*].fairplayConfiguration[*].askSecretRef", "policy_option[*].fairplay_configuration[*].pfx": "spec.forProvider.policyOption[*].fairplayConfiguration[*].pfxSecretRef", "policy_option[*].fairplay_configuration[*].pfx_password": "spec.forProvider.policyOption[*].fairplayConfiguration[*].pfxPasswordSecretRef", "policy_option[*].playready_configuration_license[*].grace_period": "spec.forProvider.policyOption[*].playreadyConfigurationLicense[*].gracePeriodSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].rsa_token_key_exponent": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].rsaTokenKeyExponentSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].rsa_token_key_modulus": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].rsaTokenKeyModulusSecretRef", "policy_option[*].token_restriction[*].alternate_key[*].symmetric_token_key": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].symmetricTokenKeySecretRef", "policy_option[*].token_restriction[*].alternate_key[*].x509_token_key_raw": "spec.forProvider.policyOption[*].tokenRestriction[*].alternateKey[*].x509TokenKeyRawSecretRef", "policy_option[*].token_restriction[*].primary_rsa_token_key_exponent": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryRsaTokenKeyExponentSecretRef", "policy_option[*].token_restriction[*].primary_rsa_token_key_modulus": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryRsaTokenKeyModulusSecretRef", "policy_option[*].token_restriction[*].primary_symmetric_token_key": "spec.forProvider.policyOption[*].tokenRestriction[*].primarySymmetricTokenKeySecretRef", "policy_option[*].token_restriction[*].primary_x509_token_key_raw": "spec.forProvider.policyOption[*].tokenRestriction[*].primaryX509TokenKeyRawSecretRef"}
}

// GetObservation of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ContentKeyPolicy
func (tr *ContentKeyPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ContentKeyPolicy
func (tr *ContentKeyPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ContentKeyPolicy
func (tr *ContentKeyPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ContentKeyPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ContentKeyPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &ContentKeyPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ContentKeyPolicy) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this Job
func (mg *Job) GetTerraformResourceType() string {
	return "azurerm_media_job"
}

// GetConnectionDetailsMapping for this Job
func (tr *Job) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Job
func (tr *Job) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Job
func (tr *Job) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Job
func (tr *Job) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Job
func (tr *Job) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Job
func (tr *Job) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Job using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Job) LateInitialize(attrs []byte) (bool, error) {
	params := &JobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Job) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this LiveEvent
func (mg *LiveEvent) GetTerraformResourceType() string {
	return "azurerm_media_live_event"
}

// GetConnectionDetailsMapping for this LiveEvent
func (tr *LiveEvent) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LiveEvent
func (tr *LiveEvent) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LiveEvent
func (tr *LiveEvent) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LiveEvent
func (tr *LiveEvent) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LiveEvent
func (tr *LiveEvent) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LiveEvent
func (tr *LiveEvent) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LiveEvent using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LiveEvent) LateInitialize(attrs []byte) (bool, error) {
	params := &LiveEventParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LiveEvent) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this LiveEventOutput
func (mg *LiveEventOutput) GetTerraformResourceType() string {
	return "azurerm_media_live_event_output"
}

// GetConnectionDetailsMapping for this LiveEventOutput
func (tr *LiveEventOutput) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LiveEventOutput
func (tr *LiveEventOutput) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LiveEventOutput
func (tr *LiveEventOutput) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LiveEventOutput
func (tr *LiveEventOutput) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LiveEventOutput
func (tr *LiveEventOutput) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LiveEventOutput
func (tr *LiveEventOutput) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LiveEventOutput using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LiveEventOutput) LateInitialize(attrs []byte) (bool, error) {
	params := &LiveEventOutputParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LiveEventOutput) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this ServicesAccount
func (mg *ServicesAccount) GetTerraformResourceType() string {
	return "azurerm_media_services_account"
}

// GetConnectionDetailsMapping for this ServicesAccount
func (tr *ServicesAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServicesAccount
func (tr *ServicesAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServicesAccount
func (tr *ServicesAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServicesAccount
func (tr *ServicesAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServicesAccount
func (tr *ServicesAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServicesAccount
func (tr *ServicesAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServicesAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServicesAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &ServicesAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServicesAccount) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this ServicesAccountFilter
func (mg *ServicesAccountFilter) GetTerraformResourceType() string {
	return "azurerm_media_services_account_filter"
}

// GetConnectionDetailsMapping for this ServicesAccountFilter
func (tr *ServicesAccountFilter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServicesAccountFilter
func (tr *ServicesAccountFilter) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServicesAccountFilter
func (tr *ServicesAccountFilter) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServicesAccountFilter
func (tr *ServicesAccountFilter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServicesAccountFilter
func (tr *ServicesAccountFilter) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServicesAccountFilter
func (tr *ServicesAccountFilter) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServicesAccountFilter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServicesAccountFilter) LateInitialize(attrs []byte) (bool, error) {
	params := &ServicesAccountFilterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServicesAccountFilter) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this StreamingEndpoint
func (mg *StreamingEndpoint) GetTerraformResourceType() string {
	return "azurerm_media_streaming_endpoint"
}

// GetConnectionDetailsMapping for this StreamingEndpoint
func (tr *StreamingEndpoint) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this StreamingEndpoint
func (tr *StreamingEndpoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamingEndpoint
func (tr *StreamingEndpoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamingEndpoint
func (tr *StreamingEndpoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamingEndpoint
func (tr *StreamingEndpoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamingEndpoint
func (tr *StreamingEndpoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamingEndpoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamingEndpoint) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamingEndpointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamingEndpoint) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this StreamingLocator
func (mg *StreamingLocator) GetTerraformResourceType() string {
	return "azurerm_media_streaming_locator"
}

// GetConnectionDetailsMapping for this StreamingLocator
func (tr *StreamingLocator) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this StreamingLocator
func (tr *StreamingLocator) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamingLocator
func (tr *StreamingLocator) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamingLocator
func (tr *StreamingLocator) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamingLocator
func (tr *StreamingLocator) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamingLocator
func (tr *StreamingLocator) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamingLocator using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamingLocator) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamingLocatorParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamingLocator) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this StreamingPolicy
func (mg *StreamingPolicy) GetTerraformResourceType() string {
	return "azurerm_media_streaming_policy"
}

// GetConnectionDetailsMapping for this StreamingPolicy
func (tr *StreamingPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this StreamingPolicy
func (tr *StreamingPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamingPolicy
func (tr *StreamingPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamingPolicy
func (tr *StreamingPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamingPolicy
func (tr *StreamingPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamingPolicy
func (tr *StreamingPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamingPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamingPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamingPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamingPolicy) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this Transform
func (mg *Transform) GetTerraformResourceType() string {
	return "azurerm_media_transform"
}

// GetConnectionDetailsMapping for this Transform
func (tr *Transform) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Transform
func (tr *Transform) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Transform
func (tr *Transform) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Transform
func (tr *Transform) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Transform
func (tr *Transform) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Transform
func (tr *Transform) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Transform using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Transform) LateInitialize(attrs []byte) (bool, error) {
	params := &TransformParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Transform) GetTerraformSchemaVersion() int {
	return 1
}
