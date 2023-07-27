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

type SpringCloudCustomDomainInitParameters struct {

	// Specifies the name of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when thumbprint is specified
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Specifies the name of the Spring Cloud Custom Domain. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when certificate_name is specified. Changing this forces a new resource to be created.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudCustomDomainObservation struct {

	// Specifies the name of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when thumbprint is specified
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// The ID of the Spring Cloud Custom Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Spring Cloud Custom Domain. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the resource ID of the Spring Cloud Application. Changing this forces a new resource to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when certificate_name is specified. Changing this forces a new resource to be created.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudCustomDomainParameters struct {

	// Specifies the name of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when thumbprint is specified
	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Specifies the name of the Spring Cloud Custom Domain. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the resource ID of the Spring Cloud Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud Custom Domain. Required when certificate_name is specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

// SpringCloudCustomDomainSpec defines the desired state of SpringCloudCustomDomain
type SpringCloudCustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudCustomDomainParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SpringCloudCustomDomainInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudCustomDomainStatus defines the observed state of SpringCloudCustomDomain.
type SpringCloudCustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudCustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudCustomDomain is the Schema for the SpringCloudCustomDomains API. Manages an Azure Spring Cloud Custom Domain.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SpringCloudCustomDomainSpec   `json:"spec"`
	Status SpringCloudCustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudCustomDomainList contains a list of SpringCloudCustomDomains
type SpringCloudCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudCustomDomain `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudCustomDomain_Kind             = "SpringCloudCustomDomain"
	SpringCloudCustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudCustomDomain_Kind}.String()
	SpringCloudCustomDomain_KindAPIVersion   = SpringCloudCustomDomain_Kind + "." + CRDGroupVersion.String()
	SpringCloudCustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudCustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudCustomDomain{}, &SpringCloudCustomDomainList{})
}
