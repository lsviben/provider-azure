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

type PublicIPPrefixInitParameters struct {

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created. Default is IPv4.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to 28(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The SKU of the Public IP Prefix. Accepted values are Standard. Defaults to Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Public IP Prefix should be located. Changing this forces a new Public IP Prefix to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type PublicIPPrefixObservation struct {

	// The Public IP Prefix ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address prefix value that was allocated.
	IPPrefix *string `json:"ipPrefix,omitempty" tf:"ip_prefix,omitempty"`

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created. Default is IPv4.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to 28(16 addresses). Changing this forces a new resource to be created.
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The name of the resource group in which to create the Public IP Prefix. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of the Public IP Prefix. Accepted values are Standard. Defaults to Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Public IP Prefix should be located. Changing this forces a new Public IP Prefix to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type PublicIPPrefixParameters struct {

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created. Default is IPv4.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the number of bits of the prefix. The value can be set between 0 (4,294,967,296 addresses) and 31 (2 addresses). Defaults to 28(16 addresses). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The name of the resource group in which to create the Public IP Prefix. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Public IP Prefix. Accepted values are Standard. Defaults to Standard. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a list of Availability Zones in which this Public IP Prefix should be located. Changing this forces a new Public IP Prefix to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// PublicIPPrefixSpec defines the desired state of PublicIPPrefix
type PublicIPPrefixSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIPPrefixParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider PublicIPPrefixInitParameters `json:"initProvider,omitempty"`
}

// PublicIPPrefixStatus defines the observed state of PublicIPPrefix.
type PublicIPPrefixStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIPPrefixObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPPrefix is the Schema for the PublicIPPrefixs API. Manages a Public IP Prefix.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PublicIPPrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   PublicIPPrefixSpec   `json:"spec"`
	Status PublicIPPrefixStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPPrefixList contains a list of PublicIPPrefixs
type PublicIPPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIPPrefix `json:"items"`
}

// Repository type metadata.
var (
	PublicIPPrefix_Kind             = "PublicIPPrefix"
	PublicIPPrefix_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicIPPrefix_Kind}.String()
	PublicIPPrefix_KindAPIVersion   = PublicIPPrefix_Kind + "." + CRDGroupVersion.String()
	PublicIPPrefix_GroupVersionKind = CRDGroupVersion.WithKind(PublicIPPrefix_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicIPPrefix{}, &PublicIPPrefixList{})
}
