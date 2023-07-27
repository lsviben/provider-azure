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

type DNSZoneInitParameters struct {

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	SoaRecord []SoaRecordInitParameters `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DNSZoneObservation struct {

	// The DNS Zone ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of Records in the zone. Defaults to 1000.
	MaxNumberOfRecordSets *float64 `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets,omitempty"`

	// A list of values that make up the NS record for the zone.
	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	// The number of records already in the zone.
	NumberOfRecordSets *float64 `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	SoaRecord []SoaRecordObservation `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DNSZoneParameters struct {

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SoaRecord []SoaRecordParameters `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordInitParameters struct {

	// The email contact for the SOA record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The domain name of the authoritative name server for the SOA record.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 300.
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The serial number for the SOA record. Defaults to 1.
	SerialNumber *float64 `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordObservation struct {

	// The email contact for the SOA record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The domain name of the authoritative name server for the SOA record.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 300.
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The serial number for the SOA record. Defaults to 1.
	SerialNumber *float64 `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordParameters struct {

	// The email contact for the SOA record.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	// +kubebuilder:validation:Optional
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The domain name of the authoritative name server for the SOA record.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 300.
	// +kubebuilder:validation:Optional
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	// +kubebuilder:validation:Optional
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	// +kubebuilder:validation:Optional
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The serial number for the SOA record. Defaults to 1.
	// +kubebuilder:validation:Optional
	SerialNumber *float64 `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DNSZoneSpec defines the desired state of DNSZone
type DNSZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSZoneParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DNSZoneInitParameters `json:"initProvider,omitempty"`
}

// DNSZoneStatus defines the observed state of DNSZone.
type DNSZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSZone is the Schema for the DNSZones API. Manages a DNS Zone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DNSZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSZoneSpec   `json:"spec"`
	Status            DNSZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSZoneList contains a list of DNSZones
type DNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSZone `json:"items"`
}

// Repository type metadata.
var (
	DNSZone_Kind             = "DNSZone"
	DNSZone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSZone_Kind}.String()
	DNSZone_KindAPIVersion   = DNSZone_Kind + "." + CRDGroupVersion.String()
	DNSZone_GroupVersionKind = CRDGroupVersion.WithKind(DNSZone_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSZone{}, &DNSZoneList{})
}
