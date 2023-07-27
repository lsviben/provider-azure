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

type HPCCacheBlobNFSTargetInitParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The client-facing file path of the HPC Cache Blob NFS Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob NFS Target. Changing this forces a new resource to be created.
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// The type of usage of the HPC Cache Blob NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

type HPCCacheBlobNFSTargetObservation struct {

	// The name of the access policy applied to this target. Defaults to default.
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The name of the HPC Cache, which the HPC Cache Blob NFS Target will be added to. Changing this forces a new HPC Cache Blob NFS Target to be created.
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// The ID of the HPC Cache Blob NFS Target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The client-facing file path of the HPC Cache Blob NFS Target.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The name of the Resource Group where the HPC Cache Blob NFS Target should exist. Changing this forces a new HPC Cache Blob NFS Target to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob NFS Target. Changing this forces a new resource to be created.
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// The type of usage of the HPC Cache Blob NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

type HPCCacheBlobNFSTargetParameters struct {

	// The name of the access policy applied to this target. Defaults to default.
	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// The name of the HPC Cache, which the HPC Cache Blob NFS Target will be added to. Changing this forces a new HPC Cache Blob NFS Target to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storagecache/v1beta1.HPCCache
	// +kubebuilder:validation:Optional
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// Reference to a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// Selector for a HPCCache in storagecache to populate cacheName.
	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// The client-facing file path of the HPC Cache Blob NFS Target.
	// +kubebuilder:validation:Optional
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// The name of the Resource Group where the HPC Cache Blob NFS Target should exist. Changing this forces a new HPC Cache Blob NFS Target to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Resource Manager ID of the Storage Container used as the HPC Cache Blob NFS Target. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// The type of usage of the HPC Cache Blob NFS Target. Possible values are: READ_HEAVY_INFREQ, READ_HEAVY_CHECK_180, WRITE_WORKLOAD_15, WRITE_AROUND, WRITE_WORKLOAD_CHECK_30, WRITE_WORKLOAD_CHECK_60 and WRITE_WORKLOAD_CLOUDWS.
	// +kubebuilder:validation:Optional
	UsageModel *string `json:"usageModel,omitempty" tf:"usage_model,omitempty"`
}

// HPCCacheBlobNFSTargetSpec defines the desired state of HPCCacheBlobNFSTarget
type HPCCacheBlobNFSTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheBlobNFSTargetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider HPCCacheBlobNFSTargetInitParameters `json:"initProvider,omitempty"`
}

// HPCCacheBlobNFSTargetStatus defines the observed state of HPCCacheBlobNFSTarget.
type HPCCacheBlobNFSTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheBlobNFSTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobNFSTarget is the Schema for the HPCCacheBlobNFSTargets API. Manages a Blob NFSv3 Target within a HPC Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCacheBlobNFSTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.namespacePath) || has(self.initProvider.namespacePath)",message="namespacePath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageContainerId) || has(self.initProvider.storageContainerId)",message="storageContainerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usageModel) || has(self.initProvider.usageModel)",message="usageModel is a required parameter"
	Spec   HPCCacheBlobNFSTargetSpec   `json:"spec"`
	Status HPCCacheBlobNFSTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobNFSTargetList contains a list of HPCCacheBlobNFSTargets
type HPCCacheBlobNFSTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheBlobNFSTarget `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheBlobNFSTarget_Kind             = "HPCCacheBlobNFSTarget"
	HPCCacheBlobNFSTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheBlobNFSTarget_Kind}.String()
	HPCCacheBlobNFSTarget_KindAPIVersion   = HPCCacheBlobNFSTarget_Kind + "." + CRDGroupVersion.String()
	HPCCacheBlobNFSTarget_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheBlobNFSTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheBlobNFSTarget{}, &HPCCacheBlobNFSTargetList{})
}
