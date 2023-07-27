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

type MaintenanceAssignmentVirtualMachineInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type MaintenanceAssignmentVirtualMachineObservation struct {

	// The ID of the Maintenance Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId,omitempty" tf:"maintenance_configuration_id,omitempty"`

	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type MaintenanceAssignmentVirtualMachineParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the ID of the Maintenance Configuration Resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/maintenance/v1beta1.MaintenanceConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId,omitempty" tf:"maintenance_configuration_id,omitempty"`

	// Reference to a MaintenanceConfiguration in maintenance to populate maintenanceConfigurationId.
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationIDRef *v1.Reference `json:"maintenanceConfigurationIdRef,omitempty" tf:"-"`

	// Selector for a MaintenanceConfiguration in maintenance to populate maintenanceConfigurationId.
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationIDSelector *v1.Selector `json:"maintenanceConfigurationIdSelector,omitempty" tf:"-"`

	// Specifies the Virtual Machine ID to which the Maintenance Configuration will be assigned. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachine
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`

	// Reference to a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDRef *v1.Reference `json:"virtualMachineIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachine in compute to populate virtualMachineId.
	// +kubebuilder:validation:Optional
	VirtualMachineIDSelector *v1.Selector `json:"virtualMachineIdSelector,omitempty" tf:"-"`
}

// MaintenanceAssignmentVirtualMachineSpec defines the desired state of MaintenanceAssignmentVirtualMachine
type MaintenanceAssignmentVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceAssignmentVirtualMachineParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider MaintenanceAssignmentVirtualMachineInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceAssignmentVirtualMachineStatus defines the observed state of MaintenanceAssignmentVirtualMachine.
type MaintenanceAssignmentVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceAssignmentVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachine is the Schema for the MaintenanceAssignmentVirtualMachines API. Manages a Maintenance Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MaintenanceAssignmentVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   MaintenanceAssignmentVirtualMachineSpec   `json:"spec"`
	Status MaintenanceAssignmentVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachineList contains a list of MaintenanceAssignmentVirtualMachines
type MaintenanceAssignmentVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceAssignmentVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceAssignmentVirtualMachine_Kind             = "MaintenanceAssignmentVirtualMachine"
	MaintenanceAssignmentVirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceAssignmentVirtualMachine_Kind}.String()
	MaintenanceAssignmentVirtualMachine_KindAPIVersion   = MaintenanceAssignmentVirtualMachine_Kind + "." + CRDGroupVersion.String()
	MaintenanceAssignmentVirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceAssignmentVirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceAssignmentVirtualMachine{}, &MaintenanceAssignmentVirtualMachineList{})
}
