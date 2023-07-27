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

type OutputSynapseInitParameters struct {

	// The name of the Azure SQL database. Changing this forces a new resource to be created.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The name of the SQL server containing the Azure SQL database. Changing this forces a new resource to be created.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The name of the table in the Azure SQL database. Changing this forces a new resource to be created.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type OutputSynapseObservation struct {

	// The name of the Azure SQL database. Changing this forces a new resource to be created.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The ID of the Stream Analytics Output to an Azure Synapse Analytics Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the SQL server containing the Azure SQL database. Changing this forces a new resource to be created.
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// The name of the table in the Azure SQL database. Changing this forces a new resource to be created.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The user name that will be used to connect to the Azure SQL database. Changing this forces a new resource to be created.
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type OutputSynapseParameters struct {

	// The name of the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The password that will be used to connect to the Azure SQL database.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SQL server containing the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Server *string `json:"server,omitempty" tf:"server,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// The name of the table in the Azure SQL database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// The user name that will be used to connect to the Azure SQL database. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("sql_administrator_login",false)
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a Workspace in synapse to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// OutputSynapseSpec defines the desired state of OutputSynapse
type OutputSynapseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputSynapseParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider OutputSynapseInitParameters `json:"initProvider,omitempty"`
}

// OutputSynapseStatus defines the observed state of OutputSynapse.
type OutputSynapseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputSynapseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapse is the Schema for the OutputSynapses API. Manages a Stream Analytics Output to an Azure Synapse Analytics Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputSynapse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.database) || has(self.initProvider.database)",message="database is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.server) || has(self.initProvider.server)",message="server is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.table) || has(self.initProvider.table)",message="table is a required parameter"
	Spec   OutputSynapseSpec   `json:"spec"`
	Status OutputSynapseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapseList contains a list of OutputSynapses
type OutputSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputSynapse `json:"items"`
}

// Repository type metadata.
var (
	OutputSynapse_Kind             = "OutputSynapse"
	OutputSynapse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputSynapse_Kind}.String()
	OutputSynapse_KindAPIVersion   = OutputSynapse_Kind + "." + CRDGroupVersion.String()
	OutputSynapse_GroupVersionKind = CRDGroupVersion.WithKind(OutputSynapse_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputSynapse{}, &OutputSynapseList{})
}
