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

type NATGatewayInitParameters struct {

	// The idle timeout which should be used in minutes. Defaults to 4.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The SKU which should be used. At this time the only supported value is Standard. Defaults to Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Availability Zones in which this NAT Gateway should be located. Changing this forces a new NAT Gateway to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type NATGatewayObservation struct {

	// The ID of the NAT Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The idle timeout which should be used in minutes. Defaults to 4.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource GUID property of the NAT Gateway.
	ResourceGUID *string `json:"resourceGuid,omitempty" tf:"resource_guid,omitempty"`

	// Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU which should be used. At this time the only supported value is Standard. Defaults to Standard.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Availability Zones in which this NAT Gateway should be located. Changing this forces a new NAT Gateway to be created.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type NATGatewayParameters struct {

	// The idle timeout which should be used in minutes. Defaults to 4.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the NAT Gateway should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name of the Resource Group in which the NAT Gateway should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU which should be used. At this time the only supported value is Standard. Defaults to Standard.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of Availability Zones in which this NAT Gateway should be located. Changing this forces a new NAT Gateway to be created.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// NATGatewaySpec defines the desired state of NATGateway
type NATGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATGatewayParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NATGatewayInitParameters `json:"initProvider,omitempty"`
}

// NATGatewayStatus defines the observed state of NATGateway.
type NATGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATGateway is the Schema for the NATGateways API. Manages a Azure NAT Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NATGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   NATGatewaySpec   `json:"spec"`
	Status NATGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATGatewayList contains a list of NATGateways
type NATGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATGateway `json:"items"`
}

// Repository type metadata.
var (
	NATGateway_Kind             = "NATGateway"
	NATGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATGateway_Kind}.String()
	NATGateway_KindAPIVersion   = NATGateway_Kind + "." + CRDGroupVersion.String()
	NATGateway_GroupVersionKind = CRDGroupVersion.WithKind(NATGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&NATGateway{}, &NATGatewayList{})
}
