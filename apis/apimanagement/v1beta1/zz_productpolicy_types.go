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

type ProductPolicyInitParameters struct {

	// The XML Content for this Policy.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ProductPolicyObservation struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The ID of the API Management Product Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The XML Content for this Policy.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ProductPolicyParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Product
	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Reference to a Product to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDRef *v1.Reference `json:"productIdRef,omitempty" tf:"-"`

	// Selector for a Product to populate productId.
	// +kubebuilder:validation:Optional
	ProductIDSelector *v1.Selector `json:"productIdSelector,omitempty" tf:"-"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The XML Content for this Policy.
	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// ProductPolicySpec defines the desired state of ProductPolicy
type ProductPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProductPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ProductPolicyInitParameters `json:"initProvider,omitempty"`
}

// ProductPolicyStatus defines the observed state of ProductPolicy.
type ProductPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProductPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProductPolicy is the Schema for the ProductPolicys API. Manages an API Management Product Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ProductPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProductPolicySpec   `json:"spec"`
	Status            ProductPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductPolicyList contains a list of ProductPolicys
type ProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductPolicy `json:"items"`
}

// Repository type metadata.
var (
	ProductPolicy_Kind             = "ProductPolicy"
	ProductPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProductPolicy_Kind}.String()
	ProductPolicy_KindAPIVersion   = ProductPolicy_Kind + "." + CRDGroupVersion.String()
	ProductPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ProductPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ProductPolicy{}, &ProductPolicyList{})
}
