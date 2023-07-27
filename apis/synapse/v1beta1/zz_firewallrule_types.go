/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallRuleInitParameters struct {

	// The ending IP address to allow through the firewall for this rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The starting IP address to allow through the firewall for this rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type FirewallRuleObservation struct {

	// The ending IP address to allow through the firewall for this rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The Synapse Firewall Rule ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The starting IP address to allow through the firewall for this rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`

	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`
}

type FirewallRuleParameters struct {

	// The ending IP address to allow through the firewall for this rule.
	// +kubebuilder:validation:Optional
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The starting IP address to allow through the firewall for this rule.
	// +kubebuilder:validation:Optional
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`

	// The ID of the Synapse Workspace on which to create the Firewall Rule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`
}

// FirewallRuleSpec defines the desired state of FirewallRule
type FirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// FirewallRuleStatus defines the observed state of FirewallRule.
type FirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRule is the Schema for the FirewallRules API. Manages a Synapse Firewall Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endIpAddress) || has(self.initProvider.endIpAddress)",message="endIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startIpAddress) || has(self.initProvider.startIpAddress)",message="startIpAddress is a required parameter"
	Spec   FirewallRuleSpec   `json:"spec"`
	Status FirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleList contains a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallRule_Kind             = "FirewallRule"
	FirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRule_Kind}.String()
	FirewallRule_KindAPIVersion   = FirewallRule_Kind + "." + CRDGroupVersion.String()
	FirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRule{}, &FirewallRuleList{})
}
