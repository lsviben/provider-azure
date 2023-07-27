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

type EventSourceEventHubInitParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name,omitempty"`
}

type EventSourceEventHubObservation struct {

	// Specifies the name of the EventHub Consumer Group that holds the partitions from which events will be read.
	ConsumerGroupName *string `json:"consumerGroupName,omitempty" tf:"consumer_group_name,omitempty"`

	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Specifies the name of the EventHub which will be associated with this resource.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Specifies the resource id where events will be coming from.
	EventSourceResourceID *string `json:"eventSourceResourceId,omitempty" tf:"event_source_resource_id,omitempty"`

	// The ID of the IoT Time Series Insights EventHub Event Source.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the EventHub Namespace name.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Specifies the name of the Shared Access key that grants the Event Source access to the EventHub.
	SharedAccessKeyName *string `json:"sharedAccessKeyName,omitempty" tf:"shared_access_key_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name,omitempty"`
}

type EventSourceEventHubParameters struct {

	// Specifies the name of the EventHub Consumer Group that holds the partitions from which events will be read.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.ConsumerGroup
	// +kubebuilder:validation:Optional
	ConsumerGroupName *string `json:"consumerGroupName,omitempty" tf:"consumer_group_name,omitempty"`

	// Reference to a ConsumerGroup in eventhub to populate consumerGroupName.
	// +kubebuilder:validation:Optional
	ConsumerGroupNameRef *v1.Reference `json:"consumerGroupNameRef,omitempty" tf:"-"`

	// Selector for a ConsumerGroup in eventhub to populate consumerGroupName.
	// +kubebuilder:validation:Optional
	ConsumerGroupNameSelector *v1.Selector `json:"consumerGroupNameSelector,omitempty" tf:"-"`

	// Specifies the id of the IoT Time Series Insights Environment that the Event Source should be associated with. Changing this forces a new resource to created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/timeseriesinsights/v1beta1.Gen2Environment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Reference to a Gen2Environment in timeseriesinsights to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDRef *v1.Reference `json:"environmentIdRef,omitempty" tf:"-"`

	// Selector for a Gen2Environment in timeseriesinsights to populate environmentId.
	// +kubebuilder:validation:Optional
	EnvironmentIDSelector *v1.Selector `json:"environmentIdSelector,omitempty" tf:"-"`

	// Specifies the name of the EventHub which will be associated with this resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// Specifies the resource id where events will be coming from.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EventSourceResourceID *string `json:"eventSourceResourceId,omitempty" tf:"event_source_resource_id,omitempty"`

	// Reference to a EventHub in eventhub to populate eventSourceResourceId.
	// +kubebuilder:validation:Optional
	EventSourceResourceIDRef *v1.Reference `json:"eventSourceResourceIdRef,omitempty" tf:"-"`

	// Selector for a EventHub in eventhub to populate eventSourceResourceId.
	// +kubebuilder:validation:Optional
	EventSourceResourceIDSelector *v1.Selector `json:"eventSourceResourceIdSelector,omitempty" tf:"-"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the EventHub Namespace name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a EventHubNamespace in eventhub to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace in eventhub to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Shared Access key that grants the Event Source access to the EventHub.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/eventhub/v1beta1.AuthorizationRule
	// +kubebuilder:validation:Optional
	SharedAccessKeyName *string `json:"sharedAccessKeyName,omitempty" tf:"shared_access_key_name,omitempty"`

	// Reference to a AuthorizationRule in eventhub to populate sharedAccessKeyName.
	// +kubebuilder:validation:Optional
	SharedAccessKeyNameRef *v1.Reference `json:"sharedAccessKeyNameRef,omitempty" tf:"-"`

	// Selector for a AuthorizationRule in eventhub to populate sharedAccessKeyName.
	// +kubebuilder:validation:Optional
	SharedAccessKeyNameSelector *v1.Selector `json:"sharedAccessKeyNameSelector,omitempty" tf:"-"`

	// Specifies the value of the Shared Access Policy key that grants the Time Series Insights service read access to the EventHub.
	// +kubebuilder:validation:Optional
	SharedAccessKeySecretRef v1.SecretKeySelector `json:"sharedAccessKeySecretRef" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the value that will be used as the event source's timestamp. This value defaults to the event creation time.
	// +kubebuilder:validation:Optional
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name,omitempty"`
}

// EventSourceEventHubSpec defines the desired state of EventSourceEventHub
type EventSourceEventHubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSourceEventHubParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EventSourceEventHubInitParameters `json:"initProvider,omitempty"`
}

// EventSourceEventHubStatus defines the observed state of EventSourceEventHub.
type EventSourceEventHubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSourceEventHubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceEventHub is the Schema for the EventSourceEventHubs API. Manages an Azure IoT Time Series Insights EventHub Event Source.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventSourceEventHub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sharedAccessKeySecretRef)",message="sharedAccessKeySecretRef is a required parameter"
	Spec   EventSourceEventHubSpec   `json:"spec"`
	Status EventSourceEventHubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceEventHubList contains a list of EventSourceEventHubs
type EventSourceEventHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSourceEventHub `json:"items"`
}

// Repository type metadata.
var (
	EventSourceEventHub_Kind             = "EventSourceEventHub"
	EventSourceEventHub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSourceEventHub_Kind}.String()
	EventSourceEventHub_KindAPIVersion   = EventSourceEventHub_Kind + "." + CRDGroupVersion.String()
	EventSourceEventHub_GroupVersionKind = CRDGroupVersion.WithKind(EventSourceEventHub_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSourceEventHub{}, &EventSourceEventHubList{})
}
