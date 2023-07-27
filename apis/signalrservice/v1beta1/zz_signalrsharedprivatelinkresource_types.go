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

type SignalrSharedPrivateLinkResourceInitParameters struct {

	// The name of the Signalr Shared Private Link Resource. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// The sub resource name which the Signalr Private Endpoint can connect to. Possible values are sites, vault. Changing this forces a new resource to be created.
	SubResourceName *string `json:"subResourceName,omitempty" tf:"sub_resource_name,omitempty"`
}

type SignalrSharedPrivateLinkResourceObservation struct {

	// The ID of the Signalr Shared Private Link resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Signalr Shared Private Link Resource. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// The id of the Signalr Service. Changing this forces a new resource to be created.
	SignalrServiceID *string `json:"signalrServiceId,omitempty" tf:"signalr_service_id,omitempty"`

	// The status of a private endpoint connection. Possible values are Pending, Approved, Rejected or Disconnected.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The sub resource name which the Signalr Private Endpoint can connect to. Possible values are sites, vault. Changing this forces a new resource to be created.
	SubResourceName *string `json:"subResourceName,omitempty" tf:"sub_resource_name,omitempty"`

	// The ID of the Shared Private Link Enabled Remote Resource which this Signalr Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type SignalrSharedPrivateLinkResourceParameters struct {

	// The name of the Signalr Shared Private Link Resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	// +kubebuilder:validation:Optional
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// The id of the Signalr Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/signalrservice/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SignalrServiceID *string `json:"signalrServiceId,omitempty" tf:"signalr_service_id,omitempty"`

	// Reference to a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDRef *v1.Reference `json:"signalrServiceIdRef,omitempty" tf:"-"`

	// Selector for a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDSelector *v1.Selector `json:"signalrServiceIdSelector,omitempty" tf:"-"`

	// The sub resource name which the Signalr Private Endpoint can connect to. Possible values are sites, vault. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubResourceName *string `json:"subResourceName,omitempty" tf:"sub_resource_name,omitempty"`

	// The ID of the Shared Private Link Enabled Remote Resource which this Signalr Private Endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Vault in keyvault to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Vault in keyvault to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// SignalrSharedPrivateLinkResourceSpec defines the desired state of SignalrSharedPrivateLinkResource
type SignalrSharedPrivateLinkResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SignalrSharedPrivateLinkResourceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SignalrSharedPrivateLinkResourceInitParameters `json:"initProvider,omitempty"`
}

// SignalrSharedPrivateLinkResourceStatus defines the observed state of SignalrSharedPrivateLinkResource.
type SignalrSharedPrivateLinkResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SignalrSharedPrivateLinkResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrSharedPrivateLinkResource is the Schema for the SignalrSharedPrivateLinkResources API. Manages the Shared Private Link Resource for a Signalr service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SignalrSharedPrivateLinkResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subResourceName) || has(self.initProvider.subResourceName)",message="subResourceName is a required parameter"
	Spec   SignalrSharedPrivateLinkResourceSpec   `json:"spec"`
	Status SignalrSharedPrivateLinkResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignalrSharedPrivateLinkResourceList contains a list of SignalrSharedPrivateLinkResources
type SignalrSharedPrivateLinkResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignalrSharedPrivateLinkResource `json:"items"`
}

// Repository type metadata.
var (
	SignalrSharedPrivateLinkResource_Kind             = "SignalrSharedPrivateLinkResource"
	SignalrSharedPrivateLinkResource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SignalrSharedPrivateLinkResource_Kind}.String()
	SignalrSharedPrivateLinkResource_KindAPIVersion   = SignalrSharedPrivateLinkResource_Kind + "." + CRDGroupVersion.String()
	SignalrSharedPrivateLinkResource_GroupVersionKind = CRDGroupVersion.WithKind(SignalrSharedPrivateLinkResource_Kind)
)

func init() {
	SchemeBuilder.Register(&SignalrSharedPrivateLinkResource{}, &SignalrSharedPrivateLinkResourceList{})
}
