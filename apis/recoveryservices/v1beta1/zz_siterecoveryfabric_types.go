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

type SiteRecoveryFabricInitParameters struct {

	// In what region should the fabric be located. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type SiteRecoveryFabricObservation struct {

	// The ID of the Site Recovery Fabric.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// In what region should the fabric be located. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type SiteRecoveryFabricParameters struct {

	// In what region should the fabric be located. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the vault that should be updated. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault that should be updated is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// SiteRecoveryFabricSpec defines the desired state of SiteRecoveryFabric
type SiteRecoveryFabricSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryFabricParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SiteRecoveryFabricInitParameters `json:"initProvider,omitempty"`
}

// SiteRecoveryFabricStatus defines the observed state of SiteRecoveryFabric.
type SiteRecoveryFabricStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryFabricObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryFabric is the Schema for the SiteRecoveryFabrics API. Manages a Site Recovery Replication Fabric on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryFabric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   SiteRecoveryFabricSpec   `json:"spec"`
	Status SiteRecoveryFabricStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryFabricList contains a list of SiteRecoveryFabrics
type SiteRecoveryFabricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryFabric `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryFabric_Kind             = "SiteRecoveryFabric"
	SiteRecoveryFabric_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SiteRecoveryFabric_Kind}.String()
	SiteRecoveryFabric_KindAPIVersion   = SiteRecoveryFabric_Kind + "." + CRDGroupVersion.String()
	SiteRecoveryFabric_GroupVersionKind = CRDGroupVersion.WithKind(SiteRecoveryFabric_Kind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryFabric{}, &SiteRecoveryFabricList{})
}
