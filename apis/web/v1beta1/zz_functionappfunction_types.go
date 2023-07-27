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

type FileInitParameters struct {

	// The content of the file. Changing this forces a new resource to be created.
	// The content of the file.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The filename of the file to be uploaded. Changing this forces a new resource to be created.
	// The filename of the file to be uploaded.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FileObservation struct {

	// The content of the file. Changing this forces a new resource to be created.
	// The content of the file.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The filename of the file to be uploaded. Changing this forces a new resource to be created.
	// The filename of the file to be uploaded.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FileParameters struct {

	// The content of the file. Changing this forces a new resource to be created.
	// The content of the file.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The filename of the file to be uploaded. Changing this forces a new resource to be created.
	// The filename of the file to be uploaded.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FunctionAppFunctionInitParameters struct {

	// The config for this Function in JSON format.
	// The config for this Function in JSON format.
	ConfigJSON *string `json:"configJson,omitempty" tf:"config_json,omitempty"`

	// Should this function be enabled. Defaults to true.
	// Should this function be enabled. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A file block as detailed below. Changing this forces a new resource to be created.
	File []FileInitParameters `json:"file,omitempty" tf:"file,omitempty"`

	// The language the Function is written in. Possible values are CSharp, Custom, Java, Javascript, Python, PowerShell, and TypeScript.
	// The language the Function is written in.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The name of the function. Changing this forces a new resource to be created.
	// The name of the function.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The test data for the function.
	// The test data for the function.
	TestData *string `json:"testData,omitempty" tf:"test_data,omitempty"`
}

type FunctionAppFunctionObservation struct {

	// The config for this Function in JSON format.
	// The config for this Function in JSON format.
	ConfigJSON *string `json:"configJson,omitempty" tf:"config_json,omitempty"`

	// The URL of the configuration JSON.
	// The URL of the configuration JSON.
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// Should this function be enabled. Defaults to true.
	// Should this function be enabled. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A file block as detailed below. Changing this forces a new resource to be created.
	File []FileObservation `json:"file,omitempty" tf:"file,omitempty"`

	// The ID of the Function App in which this function should reside. Changing this forces a new resource to be created.
	// The ID of the Function App in which this function should reside.
	FunctionAppID *string `json:"functionAppId,omitempty" tf:"function_app_id,omitempty"`

	// The ID of the Function App Function
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The invocation URL.
	// The invocation URL.
	InvocationURL *string `json:"invocationUrl,omitempty" tf:"invocation_url,omitempty"`

	// The language the Function is written in. Possible values are CSharp, Custom, Java, Javascript, Python, PowerShell, and TypeScript.
	// The language the Function is written in.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The name of the function. Changing this forces a new resource to be created.
	// The name of the function.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Script root path URL.
	// The Script root path URL.
	ScriptRootPathURL *string `json:"scriptRootPathUrl,omitempty" tf:"script_root_path_url,omitempty"`

	// The script URL.
	// The script URL.
	ScriptURL *string `json:"scriptUrl,omitempty" tf:"script_url,omitempty"`

	// The URL for the Secrets File.
	// The URL for the Secrets File.
	SecretsFileURL *string `json:"secretsFileUrl,omitempty" tf:"secrets_file_url,omitempty"`

	// The test data for the function.
	// The test data for the function.
	TestData *string `json:"testData,omitempty" tf:"test_data,omitempty"`

	// The Test data URL.
	// The Test data URL.
	TestDataURL *string `json:"testDataUrl,omitempty" tf:"test_data_url,omitempty"`

	// The function URL.
	// The function URL.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type FunctionAppFunctionParameters struct {

	// The config for this Function in JSON format.
	// The config for this Function in JSON format.
	// +kubebuilder:validation:Optional
	ConfigJSON *string `json:"configJson,omitempty" tf:"config_json,omitempty"`

	// Should this function be enabled. Defaults to true.
	// Should this function be enabled. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A file block as detailed below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	File []FileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// The ID of the Function App in which this function should reside. Changing this forces a new resource to be created.
	// The ID of the Function App in which this function should reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.LinuxFunctionApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FunctionAppID *string `json:"functionAppId,omitempty" tf:"function_app_id,omitempty"`

	// Reference to a LinuxFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDRef *v1.Reference `json:"functionAppIdRef,omitempty" tf:"-"`

	// Selector for a LinuxFunctionApp in web to populate functionAppId.
	// +kubebuilder:validation:Optional
	FunctionAppIDSelector *v1.Selector `json:"functionAppIdSelector,omitempty" tf:"-"`

	// The language the Function is written in. Possible values are CSharp, Custom, Java, Javascript, Python, PowerShell, and TypeScript.
	// The language the Function is written in.
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// The name of the function. Changing this forces a new resource to be created.
	// The name of the function.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The test data for the function.
	// The test data for the function.
	// +kubebuilder:validation:Optional
	TestData *string `json:"testData,omitempty" tf:"test_data,omitempty"`
}

// FunctionAppFunctionSpec defines the desired state of FunctionAppFunction
type FunctionAppFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppFunctionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FunctionAppFunctionInitParameters `json:"initProvider,omitempty"`
}

// FunctionAppFunctionStatus defines the observed state of FunctionAppFunction.
type FunctionAppFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppFunction is the Schema for the FunctionAppFunctions API. Manages a Function App Function.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionAppFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configJson) || has(self.initProvider.configJson)",message="configJson is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   FunctionAppFunctionSpec   `json:"spec"`
	Status FunctionAppFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppFunctionList contains a list of FunctionAppFunctions
type FunctionAppFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppFunction `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppFunction_Kind             = "FunctionAppFunction"
	FunctionAppFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppFunction_Kind}.String()
	FunctionAppFunction_KindAPIVersion   = FunctionAppFunction_Kind + "." + CRDGroupVersion.String()
	FunctionAppFunction_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppFunction{}, &FunctionAppFunctionList{})
}
