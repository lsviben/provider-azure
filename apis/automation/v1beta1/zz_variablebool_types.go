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

type VariableBoolInitParameters struct {

	// The description of the Automation Variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The value of the Automation Variable as a boolean.
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

type VariableBoolObservation struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// The description of the Automation Variable.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ID of the Automation Variable.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The value of the Automation Variable as a boolean.
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

type VariableBoolParameters struct {

	// The name of the automation account in which the Variable is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// The description of the Automation Variable.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies if the Automation Variable is encrypted. Defaults to false.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The name of the resource group in which to create the Automation Variable. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The value of the Automation Variable as a boolean.
	// +kubebuilder:validation:Optional
	Value *bool `json:"value,omitempty" tf:"value,omitempty"`
}

// VariableBoolSpec defines the desired state of VariableBool
type VariableBoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableBoolParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider VariableBoolInitParameters `json:"initProvider,omitempty"`
}

// VariableBoolStatus defines the observed state of VariableBool.
type VariableBoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableBoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VariableBool is the Schema for the VariableBools API. Manages a boolean variable in Azure Automation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VariableBool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VariableBoolSpec   `json:"spec"`
	Status            VariableBoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableBoolList contains a list of VariableBools
type VariableBoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VariableBool `json:"items"`
}

// Repository type metadata.
var (
	VariableBool_Kind             = "VariableBool"
	VariableBool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VariableBool_Kind}.String()
	VariableBool_KindAPIVersion   = VariableBool_Kind + "." + CRDGroupVersion.String()
	VariableBool_GroupVersionKind = CRDGroupVersion.WithKind(VariableBool_Kind)
)

func init() {
	SchemeBuilder.Register(&VariableBool{}, &VariableBoolList{})
}
