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

type CassandraClusterIdentityInitParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Cassandra Cluster. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CassandraClusterIdentityObservation struct {

	// The ID of the Cassandra Cluster.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Cassandra Cluster.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Cassandra Cluster. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CassandraClusterIdentityParameters struct {

	// Specifies the type of Managed Service Identity that should be configured on this Cassandra Cluster. The only possible value is SystemAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CassandraClusterInitParameters struct {

	// The authentication method that is used to authenticate clients. Possible values are None and Cassandra. Defaults to Cassandra.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// A list of TLS certificates that is used to authorize client connecting to the Cassandra Cluster.
	ClientCertificatePems []*string `json:"clientCertificatePems,omitempty" tf:"client_certificate_pems,omitempty"`

	// A list of TLS certificates that is used to authorize gossip from unmanaged Cassandra Data Center.
	ExternalGossipCertificatePems []*string `json:"externalGossipCertificatePems,omitempty" tf:"external_gossip_certificate_pems,omitempty"`

	// A list of IP Addresses of the seed nodes in unmanaged the Cassandra Data Center which will be added to the seed node lists of all managed nodes.
	ExternalSeedNodeIPAddresses []*string `json:"externalSeedNodeIpAddresses,omitempty" tf:"external_seed_node_ip_addresses,omitempty"`

	// The number of hours to wait between taking a backup of the Cassandra Cluster. Defaults to 24.
	HoursBetweenBackups *float64 `json:"hoursBetweenBackups,omitempty" tf:"hours_between_backups,omitempty"`

	// An identity block as defined below.
	Identity []CassandraClusterIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Cassandra Cluster should exist. Changing this forces a new Cassandra Cluster to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is the automatic repair enabled on the Cassandra Cluster? Defaults to true.
	RepairEnabled *bool `json:"repairEnabled,omitempty" tf:"repair_enabled,omitempty"`

	// A mapping of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of Cassandra what the Cluster converges to run. Possible values are 3.11 and 4.0. Defaults to 3.11. Changing this forces a new Cassandra Cluster to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CassandraClusterObservation struct {

	// The authentication method that is used to authenticate clients. Possible values are None and Cassandra. Defaults to Cassandra.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// A list of TLS certificates that is used to authorize client connecting to the Cassandra Cluster.
	ClientCertificatePems []*string `json:"clientCertificatePems,omitempty" tf:"client_certificate_pems,omitempty"`

	// The ID of the delegated management subnet for this Cassandra Cluster. Changing this forces a new Cassandra Cluster to be created.
	DelegatedManagementSubnetID *string `json:"delegatedManagementSubnetId,omitempty" tf:"delegated_management_subnet_id,omitempty"`

	// A list of TLS certificates that is used to authorize gossip from unmanaged Cassandra Data Center.
	ExternalGossipCertificatePems []*string `json:"externalGossipCertificatePems,omitempty" tf:"external_gossip_certificate_pems,omitempty"`

	// A list of IP Addresses of the seed nodes in unmanaged the Cassandra Data Center which will be added to the seed node lists of all managed nodes.
	ExternalSeedNodeIPAddresses []*string `json:"externalSeedNodeIpAddresses,omitempty" tf:"external_seed_node_ip_addresses,omitempty"`

	// The number of hours to wait between taking a backup of the Cassandra Cluster. Defaults to 24.
	HoursBetweenBackups *float64 `json:"hoursBetweenBackups,omitempty" tf:"hours_between_backups,omitempty"`

	// The ID of the Cassandra Cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []CassandraClusterIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Cassandra Cluster should exist. Changing this forces a new Cassandra Cluster to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is the automatic repair enabled on the Cassandra Cluster? Defaults to true.
	RepairEnabled *bool `json:"repairEnabled,omitempty" tf:"repair_enabled,omitempty"`

	// The name of the Resource Group where the Cassandra Cluster should exist. Changing this forces a new Cassandra Cluster to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of Cassandra what the Cluster converges to run. Possible values are 3.11 and 4.0. Defaults to 3.11. Changing this forces a new Cassandra Cluster to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CassandraClusterParameters struct {

	// The authentication method that is used to authenticate clients. Possible values are None and Cassandra. Defaults to Cassandra.
	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// A list of TLS certificates that is used to authorize client connecting to the Cassandra Cluster.
	// +kubebuilder:validation:Optional
	ClientCertificatePems []*string `json:"clientCertificatePems,omitempty" tf:"client_certificate_pems,omitempty"`

	// The initial admin password for this Cassandra Cluster. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DefaultAdminPasswordSecretRef v1.SecretKeySelector `json:"defaultAdminPasswordSecretRef" tf:"-"`

	// The ID of the delegated management subnet for this Cassandra Cluster. Changing this forces a new Cassandra Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetID *string `json:"delegatedManagementSubnetId,omitempty" tf:"delegated_management_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate delegatedManagementSubnetId.
	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetIDRef *v1.Reference `json:"delegatedManagementSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate delegatedManagementSubnetId.
	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetIDSelector *v1.Selector `json:"delegatedManagementSubnetIdSelector,omitempty" tf:"-"`

	// A list of TLS certificates that is used to authorize gossip from unmanaged Cassandra Data Center.
	// +kubebuilder:validation:Optional
	ExternalGossipCertificatePems []*string `json:"externalGossipCertificatePems,omitempty" tf:"external_gossip_certificate_pems,omitempty"`

	// A list of IP Addresses of the seed nodes in unmanaged the Cassandra Data Center which will be added to the seed node lists of all managed nodes.
	// +kubebuilder:validation:Optional
	ExternalSeedNodeIPAddresses []*string `json:"externalSeedNodeIpAddresses,omitempty" tf:"external_seed_node_ip_addresses,omitempty"`

	// The number of hours to wait between taking a backup of the Cassandra Cluster. Defaults to 24.
	// +kubebuilder:validation:Optional
	HoursBetweenBackups *float64 `json:"hoursBetweenBackups,omitempty" tf:"hours_between_backups,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []CassandraClusterIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Cassandra Cluster should exist. Changing this forces a new Cassandra Cluster to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Is the automatic repair enabled on the Cassandra Cluster? Defaults to true.
	// +kubebuilder:validation:Optional
	RepairEnabled *bool `json:"repairEnabled,omitempty" tf:"repair_enabled,omitempty"`

	// The name of the Resource Group where the Cassandra Cluster should exist. Changing this forces a new Cassandra Cluster to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version of Cassandra what the Cluster converges to run. Possible values are 3.11 and 4.0. Defaults to 3.11. Changing this forces a new Cassandra Cluster to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// CassandraClusterSpec defines the desired state of CassandraCluster
type CassandraClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraClusterParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CassandraClusterInitParameters `json:"initProvider,omitempty"`
}

// CassandraClusterStatus defines the observed state of CassandraCluster.
type CassandraClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraCluster is the Schema for the CassandraClusters API. Manages a Cassandra Cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CassandraCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAdminPasswordSecretRef)",message="defaultAdminPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   CassandraClusterSpec   `json:"spec"`
	Status CassandraClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraClusterList contains a list of CassandraClusters
type CassandraClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraCluster `json:"items"`
}

// Repository type metadata.
var (
	CassandraCluster_Kind             = "CassandraCluster"
	CassandraCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraCluster_Kind}.String()
	CassandraCluster_KindAPIVersion   = CassandraCluster_Kind + "." + CRDGroupVersion.String()
	CassandraCluster_GroupVersionKind = CRDGroupVersion.WithKind(CassandraCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraCluster{}, &CassandraClusterList{})
}
