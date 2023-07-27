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

type AceInitParameters struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AceObservation struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AceParameters struct {

	// Specifies the Object ID of the Azure Active Directory User or Group that the entry relates to. Only valid for user or group entries.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the permissions for the entry in rwx form. For example, rwx gives full permissions but r-- only gives read permissions.
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// Specifies whether the ACE represents an access entry or a default entry. Default value is access.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Specifies the type of entry. Can be user, group, mask or other.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataLakeGen2FileSystemInitParameters struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	Ace []AceInitParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group of the root path (i.e. /). Possible values also include $superuser.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user of the root path (i.e. /). Possible values also include $superuser.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataLakeGen2FileSystemObservation struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	Ace []AceObservation `json:"ace,omitempty" tf:"ace,omitempty"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group of the root path (i.e. /). Possible values also include $superuser.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// The ID of the Data Lake Gen2 File System.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user of the root path (i.e. /). Possible values also include $superuser.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type DataLakeGen2FileSystemParameters struct {

	// One or more ace blocks as defined below to specify the entries for the ACL for the path.
	// +kubebuilder:validation:Optional
	Ace []AceParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// Specifies the Object ID of the Azure Active Directory Group to make the owning group of the root path (i.e. /). Possible values also include $superuser.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the Object ID of the Azure Active Directory User to make the owning user of the root path (i.e. /). Possible values also include $superuser.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// DataLakeGen2FileSystemSpec defines the desired state of DataLakeGen2FileSystem
type DataLakeGen2FileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataLakeGen2FileSystemParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DataLakeGen2FileSystemInitParameters `json:"initProvider,omitempty"`
}

// DataLakeGen2FileSystemStatus defines the observed state of DataLakeGen2FileSystem.
type DataLakeGen2FileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataLakeGen2FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2FileSystem is the Schema for the DataLakeGen2FileSystems API. Manages a Data Lake Gen2 File System within an Azure Storage Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataLakeGen2FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeGen2FileSystemSpec   `json:"spec"`
	Status            DataLakeGen2FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2FileSystemList contains a list of DataLakeGen2FileSystems
type DataLakeGen2FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLakeGen2FileSystem `json:"items"`
}

// Repository type metadata.
var (
	DataLakeGen2FileSystem_Kind             = "DataLakeGen2FileSystem"
	DataLakeGen2FileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataLakeGen2FileSystem_Kind}.String()
	DataLakeGen2FileSystem_KindAPIVersion   = DataLakeGen2FileSystem_Kind + "." + CRDGroupVersion.String()
	DataLakeGen2FileSystem_GroupVersionKind = CRDGroupVersion.WithKind(DataLakeGen2FileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&DataLakeGen2FileSystem{}, &DataLakeGen2FileSystemList{})
}
