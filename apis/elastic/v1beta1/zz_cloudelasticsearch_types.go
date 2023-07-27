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

type CloudElasticsearchInitParameters struct {

	// Specifies the Email Address which should be associated with this Elasticsearch account. Changing this forces a new Elasticsearch to be created.
	ElasticCloudEmailAddress *string `json:"elasticCloudEmailAddress,omitempty" tf:"elastic_cloud_email_address,omitempty"`

	// The Azure Region where the Elasticsearch resource should exist. Changing this forces a new Elasticsearch to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A logs block as defined below.
	Logs []LogsInitParameters `json:"logs,omitempty" tf:"logs,omitempty"`

	// Specifies if the Elasticsearch should have monitoring configured? Defaults to true. Changing this forces a new Elasticsearch to be created.
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty" tf:"monitoring_enabled,omitempty"`

	// Specifies the name of the SKU for this Elasticsearch. Changing this forces a new Elasticsearch to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Elasticsearch resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudElasticsearchObservation struct {

	// The ID of the Deployment within Elastic Cloud.
	ElasticCloudDeploymentID *string `json:"elasticCloudDeploymentId,omitempty" tf:"elastic_cloud_deployment_id,omitempty"`

	// Specifies the Email Address which should be associated with this Elasticsearch account. Changing this forces a new Elasticsearch to be created.
	ElasticCloudEmailAddress *string `json:"elasticCloudEmailAddress,omitempty" tf:"elastic_cloud_email_address,omitempty"`

	// The Default URL used for Single Sign On (SSO) to Elastic Cloud.
	ElasticCloudSsoDefaultURL *string `json:"elasticCloudSsoDefaultUrl,omitempty" tf:"elastic_cloud_sso_default_url,omitempty"`

	// The ID of the User Account within Elastic Cloud.
	ElasticCloudUserID *string `json:"elasticCloudUserId,omitempty" tf:"elastic_cloud_user_id,omitempty"`

	// The URL to the Elasticsearch Service associated with this Elasticsearch.
	ElasticsearchServiceURL *string `json:"elasticsearchServiceUrl,omitempty" tf:"elasticsearch_service_url,omitempty"`

	// The ID of the Elasticsearch.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL to the Kibana Dashboard associated with this Elasticsearch.
	KibanaServiceURL *string `json:"kibanaServiceUrl,omitempty" tf:"kibana_service_url,omitempty"`

	// The URI used for SSO to the Kibana Dashboard associated with this Elasticsearch.
	KibanaSsoURI *string `json:"kibanaSsoUri,omitempty" tf:"kibana_sso_uri,omitempty"`

	// The Azure Region where the Elasticsearch resource should exist. Changing this forces a new Elasticsearch to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A logs block as defined below.
	Logs []LogsObservation `json:"logs,omitempty" tf:"logs,omitempty"`

	// Specifies if the Elasticsearch should have monitoring configured? Defaults to true. Changing this forces a new Elasticsearch to be created.
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty" tf:"monitoring_enabled,omitempty"`

	// The name of the Resource Group where the Elasticsearch resource should exist. Changing this forces a new Elasticsearch to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the SKU for this Elasticsearch. Changing this forces a new Elasticsearch to be created.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Elasticsearch resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudElasticsearchParameters struct {

	// Specifies the Email Address which should be associated with this Elasticsearch account. Changing this forces a new Elasticsearch to be created.
	// +kubebuilder:validation:Optional
	ElasticCloudEmailAddress *string `json:"elasticCloudEmailAddress,omitempty" tf:"elastic_cloud_email_address,omitempty"`

	// The Azure Region where the Elasticsearch resource should exist. Changing this forces a new Elasticsearch to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A logs block as defined below.
	// +kubebuilder:validation:Optional
	Logs []LogsParameters `json:"logs,omitempty" tf:"logs,omitempty"`

	// Specifies if the Elasticsearch should have monitoring configured? Defaults to true. Changing this forces a new Elasticsearch to be created.
	// +kubebuilder:validation:Optional
	MonitoringEnabled *bool `json:"monitoringEnabled,omitempty" tf:"monitoring_enabled,omitempty"`

	// The name of the Resource Group where the Elasticsearch resource should exist. Changing this forces a new Elasticsearch to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the SKU for this Elasticsearch. Changing this forces a new Elasticsearch to be created.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A mapping of tags which should be assigned to the Elasticsearch resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilteringTagInitParameters struct {

	// Specifies the type of action which should be taken when the Tag matches the name and value. Possible values are Exclude and Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name which should be used for this Elasticsearch resource. Changing this forces a new Elasticsearch to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the Tag which should be filtered.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilteringTagObservation struct {

	// Specifies the type of action which should be taken when the Tag matches the name and value. Possible values are Exclude and Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name which should be used for this Elasticsearch resource. Changing this forces a new Elasticsearch to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the Tag which should be filtered.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilteringTagParameters struct {

	// Specifies the type of action which should be taken when the Tag matches the name and value. Possible values are Exclude and Include.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name which should be used for this Elasticsearch resource. Changing this forces a new Elasticsearch to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the Tag which should be filtered.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LogsInitParameters struct {

	// A list of filtering_tag blocks as defined above.
	FilteringTag []FilteringTagInitParameters `json:"filteringTag,omitempty" tf:"filtering_tag,omitempty"`

	// Specifies if the Azure Activity Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Specifies if the AzureAD Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendAzureadLogs *bool `json:"sendAzureadLogs,omitempty" tf:"send_azuread_logs,omitempty"`

	// Specifies if the Azure Subscription Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`
}

type LogsObservation struct {

	// A list of filtering_tag blocks as defined above.
	FilteringTag []FilteringTagObservation `json:"filteringTag,omitempty" tf:"filtering_tag,omitempty"`

	// Specifies if the Azure Activity Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Specifies if the AzureAD Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendAzureadLogs *bool `json:"sendAzureadLogs,omitempty" tf:"send_azuread_logs,omitempty"`

	// Specifies if the Azure Subscription Logs should be sent to the Elasticsearch cluster. Defaults to false.
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`
}

type LogsParameters struct {

	// A list of filtering_tag blocks as defined above.
	// +kubebuilder:validation:Optional
	FilteringTag []FilteringTagParameters `json:"filteringTag,omitempty" tf:"filtering_tag,omitempty"`

	// Specifies if the Azure Activity Logs should be sent to the Elasticsearch cluster. Defaults to false.
	// +kubebuilder:validation:Optional
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Specifies if the AzureAD Logs should be sent to the Elasticsearch cluster. Defaults to false.
	// +kubebuilder:validation:Optional
	SendAzureadLogs *bool `json:"sendAzureadLogs,omitempty" tf:"send_azuread_logs,omitempty"`

	// Specifies if the Azure Subscription Logs should be sent to the Elasticsearch cluster. Defaults to false.
	// +kubebuilder:validation:Optional
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`
}

// CloudElasticsearchSpec defines the desired state of CloudElasticsearch
type CloudElasticsearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudElasticsearchParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CloudElasticsearchInitParameters `json:"initProvider,omitempty"`
}

// CloudElasticsearchStatus defines the observed state of CloudElasticsearch.
type CloudElasticsearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudElasticsearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudElasticsearch is the Schema for the CloudElasticsearchs API. Manages an Elasticsearch cluster in Elastic Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CloudElasticsearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.elasticCloudEmailAddress) || has(self.initProvider.elasticCloudEmailAddress)",message="elasticCloudEmailAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || has(self.initProvider.skuName)",message="skuName is a required parameter"
	Spec   CloudElasticsearchSpec   `json:"spec"`
	Status CloudElasticsearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudElasticsearchList contains a list of CloudElasticsearchs
type CloudElasticsearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudElasticsearch `json:"items"`
}

// Repository type metadata.
var (
	CloudElasticsearch_Kind             = "CloudElasticsearch"
	CloudElasticsearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudElasticsearch_Kind}.String()
	CloudElasticsearch_KindAPIVersion   = CloudElasticsearch_Kind + "." + CRDGroupVersion.String()
	CloudElasticsearch_GroupVersionKind = CRDGroupVersion.WithKind(CloudElasticsearch_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudElasticsearch{}, &CloudElasticsearchList{})
}
