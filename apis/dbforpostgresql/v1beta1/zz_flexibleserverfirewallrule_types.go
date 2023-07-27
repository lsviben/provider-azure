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

type FlexibleServerFirewallRuleInitParameters struct {

	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type FlexibleServerFirewallRuleObservation struct {

	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The ID of the PostgreSQL Flexible Server Firewall Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

type FlexibleServerFirewallRuleParameters struct {

	// The End IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	// +kubebuilder:validation:Optional
	EndIPAddress *string `json:"endIpAddress,omitempty" tf:"end_ip_address,omitempty"`

	// The ID of the PostgreSQL Flexible Server from which to create this PostgreSQL Flexible Server Firewall Rule. Changing this forces a new PostgreSQL Flexible Server Firewall Rule to be created.
	// +crossplane:generate:reference:type=FlexibleServer
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a FlexibleServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a FlexibleServer to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The Start IP Address associated with this PostgreSQL Flexible Server Firewall Rule.
	// +kubebuilder:validation:Optional
	StartIPAddress *string `json:"startIpAddress,omitempty" tf:"start_ip_address,omitempty"`
}

// FlexibleServerFirewallRuleSpec defines the desired state of FlexibleServerFirewallRule
type FlexibleServerFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerFirewallRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FlexibleServerFirewallRuleInitParameters `json:"initProvider,omitempty"`
}

// FlexibleServerFirewallRuleStatus defines the observed state of FlexibleServerFirewallRule.
type FlexibleServerFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRule is the Schema for the FlexibleServerFirewallRules API. Manages a PostgreSQL Flexible Server Firewall Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServerFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endIpAddress) || has(self.initProvider.endIpAddress)",message="endIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.startIpAddress) || has(self.initProvider.startIpAddress)",message="startIpAddress is a required parameter"
	Spec   FlexibleServerFirewallRuleSpec   `json:"spec"`
	Status FlexibleServerFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerFirewallRuleList contains a list of FlexibleServerFirewallRules
type FlexibleServerFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServerFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServerFirewallRule_Kind             = "FlexibleServerFirewallRule"
	FlexibleServerFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServerFirewallRule_Kind}.String()
	FlexibleServerFirewallRule_KindAPIVersion   = FlexibleServerFirewallRule_Kind + "." + CRDGroupVersion.String()
	FlexibleServerFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServerFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServerFirewallRule{}, &FlexibleServerFirewallRuleList{})
}
