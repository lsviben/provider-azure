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

type SentinelWatchlistInitParameters struct {

	// The default duration in ISO8601 duration form of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	DefaultDuration *string `json:"defaultDuration,omitempty" tf:"default_duration,omitempty"`

	// The description of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The key used to optimize query performance when using Watchlist for joins with other data. Changing this forces a new Sentinel Watchlist to be created.
	ItemSearchKey *string `json:"itemSearchKey,omitempty" tf:"item_search_key,omitempty"`

	// Specifies a list of labels related to this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type SentinelWatchlistObservation struct {

	// The default duration in ISO8601 duration form of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	DefaultDuration *string `json:"defaultDuration,omitempty" tf:"default_duration,omitempty"`

	// The description of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Sentinel Watchlist.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key used to optimize query performance when using Watchlist for joins with other data. Changing this forces a new Sentinel Watchlist to be created.
	ItemSearchKey *string `json:"itemSearchKey,omitempty" tf:"item_search_key,omitempty"`

	// Specifies a list of labels related to this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the Log Analytics Workspace where this Sentinel Watchlist resides in. Changing this forces a new Sentinel Watchlist to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`
}

type SentinelWatchlistParameters struct {

	// The default duration in ISO8601 duration form of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	// +kubebuilder:validation:Optional
	DefaultDuration *string `json:"defaultDuration,omitempty" tf:"default_duration,omitempty"`

	// The description of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The key used to optimize query performance when using Watchlist for joins with other data. Changing this forces a new Sentinel Watchlist to be created.
	// +kubebuilder:validation:Optional
	ItemSearchKey *string `json:"itemSearchKey,omitempty" tf:"item_search_key,omitempty"`

	// Specifies a list of labels related to this Sentinel Watchlist. Changing this forces a new Sentinel Watchlist to be created.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the Log Analytics Workspace where this Sentinel Watchlist resides in. Changing this forces a new Sentinel Watchlist to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/securityinsights/v1beta1.SentinelLogAnalyticsWorkspaceOnboarding
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("workspace_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`
}

// SentinelWatchlistSpec defines the desired state of SentinelWatchlist
type SentinelWatchlistSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelWatchlistParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SentinelWatchlistInitParameters `json:"initProvider,omitempty"`
}

// SentinelWatchlistStatus defines the observed state of SentinelWatchlist.
type SentinelWatchlistStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelWatchlistObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelWatchlist is the Schema for the SentinelWatchlists API. Manages a Sentinel Watchlist.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelWatchlist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.itemSearchKey) || has(self.initProvider.itemSearchKey)",message="itemSearchKey is a required parameter"
	Spec   SentinelWatchlistSpec   `json:"spec"`
	Status SentinelWatchlistStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelWatchlistList contains a list of SentinelWatchlists
type SentinelWatchlistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelWatchlist `json:"items"`
}

// Repository type metadata.
var (
	SentinelWatchlist_Kind             = "SentinelWatchlist"
	SentinelWatchlist_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelWatchlist_Kind}.String()
	SentinelWatchlist_KindAPIVersion   = SentinelWatchlist_Kind + "." + CRDGroupVersion.String()
	SentinelWatchlist_GroupVersionKind = CRDGroupVersion.WithKind(SentinelWatchlist_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelWatchlist{}, &SentinelWatchlistList{})
}
