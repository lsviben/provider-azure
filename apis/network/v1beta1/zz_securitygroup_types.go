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

type SecurityGroupInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// List of objects representing security rules, as defined below.
	SecurityRule []SecurityRuleInitParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityGroupObservation struct {

	// The ID of the Network Security Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// List of objects representing security rules, as defined below.
	SecurityRule []SecurityRuleObservation `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityGroupParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// List of objects representing security rules, as defined below.
	// +kubebuilder:validation:Optional
	SecurityRule []SecurityRuleParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecurityRuleInitParameters struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	Access *string `json:"access,omitempty" tf:"access"`

	// A description for this rule. Restricted to 140 characters.
	Description *string `json:"description,omitempty" tf:"description"`

	// CIDR or destination IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if destination_address_prefixes is not specified.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes"`

	// A List of destination Application Security Group IDs
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	Direction *string `json:"direction,omitempty" tf:"direction"`

	// The name of the security rule.
	Name *string `json:"name,omitempty" tf:"name"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// CIDR or source IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if source_address_prefixes is not specified.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes"`

	// A List of source Application Security Group IDs
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges"`
}

type SecurityRuleObservation struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	Access *string `json:"access,omitempty" tf:"access,omitempty"`

	// A description for this rule. Restricted to 140 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// CIDR or destination IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if destination_address_prefixes is not specified.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`

	// A List of destination Application Security Group IDs
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids,omitempty"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// The name of the security rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// CIDR or source IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if source_address_prefixes is not specified.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`

	// A List of source Application Security Group IDs
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids,omitempty"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

type SecurityRuleParameters struct {

	// Specifies whether network traffic is allowed or denied. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	Access *string `json:"access,omitempty" tf:"access"`

	// A description for this rule. Restricted to 140 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// CIDR or destination IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if destination_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix"`

	// List of destination address prefixes. Tags may not be used. This is required if destination_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	DestinationAddressPrefixes []*string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes"`

	// A List of destination Application Security Group IDs
	// +kubebuilder:validation:Optional
	DestinationApplicationSecurityGroupIds []*string `json:"destinationApplicationSecurityGroupIds,omitempty" tf:"destination_application_security_group_ids"`

	// Destination Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if destination_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`

	// List of destination ports or port ranges. This is required if destination_port_range is not specified.
	// +kubebuilder:validation:Optional
	DestinationPortRanges []*string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges"`

	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are Inbound and Outbound.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction"`

	// The name of the security rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	// Network protocol this rule applies to. Possible values include Tcp, Udp, Icmp, Esp, Ah or * (which matches all).
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	// CIDR or source IP range or * to match any IP. Tags such as VirtualNetwork, AzureLoadBalancer and Internet can also be used. This is required if source_address_prefixes is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix"`

	// List of source address prefixes. Tags may not be used. This is required if source_address_prefix is not specified.
	// +kubebuilder:validation:Optional
	SourceAddressPrefixes []*string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes"`

	// A List of source Application Security Group IDs
	// +kubebuilder:validation:Optional
	SourceApplicationSecurityGroupIds []*string `json:"sourceApplicationSecurityGroupIds,omitempty" tf:"source_application_security_group_ids"`

	// Source Port or Range. Integer or range between 0 and 65535 or * to match any. This is required if source_port_ranges is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`

	// List of source ports or port ranges. This is required if source_port_range is not specified.
	// +kubebuilder:validation:Optional
	SourcePortRanges []*string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. Manages a network security group that contains a list of network security rules. Network security groups enable inbound or outbound traffic to be enabled or denied.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   SecurityGroupSpec   `json:"spec"`
	Status SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
