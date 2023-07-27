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

type LinkedServiceWebInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The URL of the web service endpoint (e.g. https://www.microsoft.com).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username for Basic authentication. Required if authentication_type sets to Basic.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceWebObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The URL of the web service endpoint (e.g. https://www.microsoft.com).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username for Basic authentication. Required if authentication_type sets to Basic.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LinkedServiceWebParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The type of authentication used to connect to the web table source. Valid options are Anonymous, Basic and ClientCertificate.
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The password for Basic authentication. Required if authentication_type sets to Basic.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The URL of the web service endpoint (e.g. https://www.microsoft.com).
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The username for Basic authentication. Required if authentication_type sets to Basic.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// LinkedServiceWebSpec defines the desired state of LinkedServiceWeb
type LinkedServiceWebSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceWebParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider LinkedServiceWebInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceWebStatus defines the observed state of LinkedServiceWeb.
type LinkedServiceWebStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceWebObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceWeb is the Schema for the LinkedServiceWebs API. Manages a Linked Service (connection) between a Web Server and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationType) || has(self.initProvider.authenticationType)",message="authenticationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || has(self.initProvider.url)",message="url is a required parameter"
	Spec   LinkedServiceWebSpec   `json:"spec"`
	Status LinkedServiceWebStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceWebList contains a list of LinkedServiceWebs
type LinkedServiceWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceWeb `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceWeb_Kind             = "LinkedServiceWeb"
	LinkedServiceWeb_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceWeb_Kind}.String()
	LinkedServiceWeb_KindAPIVersion   = LinkedServiceWeb_Kind + "." + CRDGroupVersion.String()
	LinkedServiceWeb_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceWeb_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceWeb{}, &LinkedServiceWebList{})
}
