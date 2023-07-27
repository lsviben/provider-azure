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

type AssignToUserInitParameters struct {

	// User’s AAD Object Id.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// User’s AAD Tenant Id.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AssignToUserObservation struct {

	// User’s AAD Object Id.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// User’s AAD Tenant Id.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AssignToUserParameters struct {

	// User’s AAD Object Id.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// User’s AAD Tenant Id.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ComputeInstanceIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Instance. Changing this forces a new resource to be created.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Instance. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeInstanceIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Instance. Changing this forces a new resource to be created.
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Instance.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Machine Learning Compute Instance.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Instance. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeInstanceIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Machine Learning Compute Instance. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Machine Learning Compute Instance. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ComputeInstanceInitParameters struct {

	// A assign_to_user block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
	AssignToUser []AssignToUserInitParameters `json:"assignToUser,omitempty" tf:"assign_to_user,omitempty"`

	// The Compute Instance Authorization type. Possible values include: personal. Changing this forces a new Machine Learning Compute Instance to be created.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
	Identity []ComputeInstanceIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Instance to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A ssh block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
	SSH []ComputeInstanceSSHInitParameters `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
	VirtualMachineSize *string `json:"virtualMachineSize,omitempty" tf:"virtual_machine_size,omitempty"`
}

type ComputeInstanceObservation struct {

	// A assign_to_user block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
	AssignToUser []AssignToUserObservation `json:"assignToUser,omitempty" tf:"assign_to_user,omitempty"`

	// The Compute Instance Authorization type. Possible values include: personal. Changing this forces a new Machine Learning Compute Instance to be created.
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Machine Learning Compute Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
	Identity []ComputeInstanceIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Instance to be created.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// A ssh block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
	SSH []ComputeInstanceSSHObservation `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
	VirtualMachineSize *string `json:"virtualMachineSize,omitempty" tf:"virtual_machine_size,omitempty"`
}

type ComputeInstanceParameters struct {

	// A assign_to_user block as defined below. A user explicitly assigned to a personal compute instance. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	AssignToUser []AssignToUserParameters `json:"assignToUser,omitempty" tf:"assign_to_user,omitempty"`

	// The Compute Instance Authorization type. Possible values include: personal. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// The description of the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An identity block as defined below. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	Identity []ComputeInstanceIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Whether local authentication methods is enabled. Defaults to true. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// The Azure Region where the Machine Learning Compute Instance should exist. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Machine Learning Workspace. Changing this forces a new Machine Learning Compute Instance to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/machinelearningservices/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceId,omitempty" tf:"machine_learning_workspace_id,omitempty"`

	// Reference to a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDRef *v1.Reference `json:"machineLearningWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in machinelearningservices to populate machineLearningWorkspaceId.
	// +kubebuilder:validation:Optional
	MachineLearningWorkspaceIDSelector *v1.Selector `json:"machineLearningWorkspaceIdSelector,omitempty" tf:"-"`

	// A ssh block as defined below. Specifies policy and settings for SSH access. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	SSH []ComputeInstanceSSHParameters `json:"ssh,omitempty" tf:"ssh,omitempty"`

	// Virtual network subnet resource ID the compute nodes belong to. Changing this forces a new Machine Learning Compute Instance to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id,omitempty"`

	// Reference to a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDRef *v1.Reference `json:"subnetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetResourceId.
	// +kubebuilder:validation:Optional
	SubnetResourceIDSelector *v1.Selector `json:"subnetResourceIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Machine Learning Compute Instance. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Virtual Machine Size. Changing this forces a new Machine Learning Compute Instance to be created.
	// +kubebuilder:validation:Optional
	VirtualMachineSize *string `json:"virtualMachineSize,omitempty" tf:"virtual_machine_size,omitempty"`
}

type ComputeInstanceSSHInitParameters struct {

	// Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH key pairs.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type ComputeInstanceSSHObservation struct {

	// Describes the port for connecting through SSH.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH key pairs.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The admin username of this Machine Learning Compute Instance.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ComputeInstanceSSHParameters struct {

	// Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH key pairs.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

// ComputeInstanceSpec defines the desired state of ComputeInstance
type ComputeInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComputeInstanceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ComputeInstanceInitParameters `json:"initProvider,omitempty"`
}

// ComputeInstanceStatus defines the observed state of ComputeInstance.
type ComputeInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComputeInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeInstance is the Schema for the ComputeInstances API. Manages a Machine Learning Compute Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.virtualMachineSize) || has(self.initProvider.virtualMachineSize)",message="virtualMachineSize is a required parameter"
	Spec   ComputeInstanceSpec   `json:"spec"`
	Status ComputeInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeInstanceList contains a list of ComputeInstances
type ComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstance `json:"items"`
}

// Repository type metadata.
var (
	ComputeInstance_Kind             = "ComputeInstance"
	ComputeInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ComputeInstance_Kind}.String()
	ComputeInstance_KindAPIVersion   = ComputeInstance_Kind + "." + CRDGroupVersion.String()
	ComputeInstance_GroupVersionKind = CRDGroupVersion.WithKind(ComputeInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ComputeInstance{}, &ComputeInstanceList{})
}
