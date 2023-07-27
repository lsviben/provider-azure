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

type RepositoryInitParameters struct {

	// Specifies the SSH public key of git repository.
	HostKey *string `json:"hostKey,omitempty" tf:"host_key,omitempty"`

	// Specifies the SSH key algorithm of git repository.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Specifies the label of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specifies the name which should be used for this repository.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the collection of patterns of the repository.
	Patterns []*string `json:"patterns,omitempty" tf:"patterns,omitempty"`

	// Specifies a list of searching path of the repository
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// Specifies whether enable the strict host key checking.
	StrictHostKeyChecking *bool `json:"strictHostKeyChecking,omitempty" tf:"strict_host_key_checking,omitempty"`

	// Specifies the URI of the repository.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the username of git repository basic auth.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RepositoryObservation struct {

	// Specifies the SSH public key of git repository.
	HostKey *string `json:"hostKey,omitempty" tf:"host_key,omitempty"`

	// Specifies the SSH key algorithm of git repository.
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Specifies the label of the repository.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specifies the name which should be used for this repository.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the collection of patterns of the repository.
	Patterns []*string `json:"patterns,omitempty" tf:"patterns,omitempty"`

	// Specifies a list of searching path of the repository
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// Specifies whether enable the strict host key checking.
	StrictHostKeyChecking *bool `json:"strictHostKeyChecking,omitempty" tf:"strict_host_key_checking,omitempty"`

	// Specifies the URI of the repository.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the username of git repository basic auth.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RepositoryParameters struct {

	// Specifies the SSH public key of git repository.
	// +kubebuilder:validation:Optional
	HostKey *string `json:"hostKey,omitempty" tf:"host_key,omitempty"`

	// Specifies the SSH key algorithm of git repository.
	// +kubebuilder:validation:Optional
	HostKeyAlgorithm *string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`

	// Specifies the label of the repository.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Specifies the name which should be used for this repository.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the password of git repository basic auth.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies the collection of patterns of the repository.
	// +kubebuilder:validation:Optional
	Patterns []*string `json:"patterns,omitempty" tf:"patterns,omitempty"`

	// Specifies the SSH private key of git repository.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef *v1.SecretKeySelector `json:"privateKeySecretRef,omitempty" tf:"-"`

	// Specifies a list of searching path of the repository
	// +kubebuilder:validation:Optional
	SearchPaths []*string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`

	// Specifies whether enable the strict host key checking.
	// +kubebuilder:validation:Optional
	StrictHostKeyChecking *bool `json:"strictHostKeyChecking,omitempty" tf:"strict_host_key_checking,omitempty"`

	// Specifies the URI of the repository.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Specifies the username of git repository basic auth.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SpringCloudConfigurationServiceInitParameters struct {

	// One or more repository blocks as defined below.
	Repository []RepositoryInitParameters `json:"repository,omitempty" tf:"repository,omitempty"`
}

type SpringCloudConfigurationServiceObservation struct {

	// The ID of the Spring Cloud Configuration Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more repository blocks as defined below.
	Repository []RepositoryObservation `json:"repository,omitempty" tf:"repository,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Configuration Service to be created.
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`
}

type SpringCloudConfigurationServiceParameters struct {

	// One or more repository blocks as defined below.
	// +kubebuilder:validation:Optional
	Repository []RepositoryParameters `json:"repository,omitempty" tf:"repository,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Configuration Service to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`
}

// SpringCloudConfigurationServiceSpec defines the desired state of SpringCloudConfigurationService
type SpringCloudConfigurationServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudConfigurationServiceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SpringCloudConfigurationServiceInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudConfigurationServiceStatus defines the observed state of SpringCloudConfigurationService.
type SpringCloudConfigurationServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudConfigurationServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudConfigurationService is the Schema for the SpringCloudConfigurationServices API. Manages a Spring Cloud Configuration Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudConfigurationService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudConfigurationServiceSpec   `json:"spec"`
	Status            SpringCloudConfigurationServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudConfigurationServiceList contains a list of SpringCloudConfigurationServices
type SpringCloudConfigurationServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudConfigurationService `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudConfigurationService_Kind             = "SpringCloudConfigurationService"
	SpringCloudConfigurationService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudConfigurationService_Kind}.String()
	SpringCloudConfigurationService_KindAPIVersion   = SpringCloudConfigurationService_Kind + "." + CRDGroupVersion.String()
	SpringCloudConfigurationService_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudConfigurationService_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudConfigurationService{}, &SpringCloudConfigurationServiceList{})
}
