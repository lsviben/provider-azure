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

type WebhookInitParameters struct {

	// Controls if Webhook is enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime *string `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of input parameters passed to runbook.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup *string `json:"runOnWorkerGroup,omitempty" tf:"run_on_worker_group,omitempty"`
}

type WebhookObservation struct {

	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Controls if Webhook is enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	ExpiryTime *string `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	// The Automation Webhook ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of input parameters passed to runbook.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Name of the Automation Runbook to execute by Webhook.
	RunBookName *string `json:"runbookName,omitempty" tf:"runbook_name,omitempty"`

	// Name of the hybrid worker group the Webhook job will run on.
	RunOnWorkerGroup *string `json:"runOnWorkerGroup,omitempty" tf:"run_on_worker_group,omitempty"`
}

type WebhookParameters struct {

	// The name of the automation account in which the Webhook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// Controls if Webhook is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp when the webhook expires. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ExpiryTime *string `json:"expiryTime,omitempty" tf:"expiry_time,omitempty"`

	// Specifies the name of the Webhook. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of input parameters passed to runbook.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The name of the resource group in which the Webhook is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Name of the Automation Runbook to execute by Webhook.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.RunBook
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	RunBookName *string `json:"runbookName,omitempty" tf:"runbook_name,omitempty"`

	// Reference to a RunBook in automation to populate runbookName.
	// +kubebuilder:validation:Optional
	RunBookNameRef *v1.Reference `json:"runbookNameRef,omitempty" tf:"-"`

	// Selector for a RunBook in automation to populate runbookName.
	// +kubebuilder:validation:Optional
	RunBookNameSelector *v1.Selector `json:"runbookNameSelector,omitempty" tf:"-"`

	// Name of the hybrid worker group the Webhook job will run on.
	// +kubebuilder:validation:Optional
	RunOnWorkerGroup *string `json:"runOnWorkerGroup,omitempty" tf:"run_on_worker_group,omitempty"`

	// URI to initiate the webhook. Can be generated using Generate URI API. By default, new URI is generated on each new resource creation. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	URISecretRef *v1.SecretKeySelector `json:"uriSecretRef,omitempty" tf:"-"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Webhook is the Schema for the Webhooks API. Manages an Automation Runbook's Webhook.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expiryTime) || has(self.initProvider.expiryTime)",message="expiryTime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}
