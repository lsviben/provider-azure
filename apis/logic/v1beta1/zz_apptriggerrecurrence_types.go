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

type AppTriggerRecurrenceInitParameters struct {

	// Specifies the Frequency at which this Trigger should be run. Possible values include Month, Week, Day, Hour, Minute and Second.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies interval used for the Frequency, for example a value of 4 for interval and hour for frequency would run the Trigger every 4 hours.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// A schedule block as specified below.
	Schedule []AppTriggerRecurrenceScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Specifies the start date and time for this trigger in RFC3339 format: 2000-01-02T03:04:05Z.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for this trigger. Supported time zone options are listed here
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AppTriggerRecurrenceObservation struct {

	// Specifies the Frequency at which this Trigger should be run. Possible values include Month, Week, Day, Hour, Minute and Second.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// The ID of the Recurrence Trigger within the Logic App Workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies interval used for the Frequency, for example a value of 4 for interval and hour for frequency would run the Trigger every 4 hours.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// A schedule block as specified below.
	Schedule []AppTriggerRecurrenceScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Specifies the start date and time for this trigger in RFC3339 format: 2000-01-02T03:04:05Z.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for this trigger. Supported time zone options are listed here
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AppTriggerRecurrenceParameters struct {

	// Specifies the Frequency at which this Trigger should be run. Possible values include Month, Week, Day, Hour, Minute and Second.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies interval used for the Frequency, for example a value of 4 for interval and hour for frequency would run the Trigger every 4 hours.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppWorkflow
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Reference to a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDRef *v1.Reference `json:"logicAppIdRef,omitempty" tf:"-"`

	// Selector for a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDSelector *v1.Selector `json:"logicAppIdSelector,omitempty" tf:"-"`

	// A schedule block as specified below.
	// +kubebuilder:validation:Optional
	Schedule []AppTriggerRecurrenceScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Specifies the start date and time for this trigger in RFC3339 format: 2000-01-02T03:04:05Z.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone for this trigger. Supported time zone options are listed here
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AppTriggerRecurrenceScheduleInitParameters struct {

	// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
	AtTheseHours []*float64 `json:"atTheseHours,omitempty" tf:"at_these_hours,omitempty"`

	// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
	AtTheseMinutes []*float64 `json:"atTheseMinutes,omitempty" tf:"at_these_minutes,omitempty"`

	// Specifies a list of days when the trigger should run. Valid values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	OnTheseDays []*string `json:"onTheseDays,omitempty" tf:"on_these_days,omitempty"`
}

type AppTriggerRecurrenceScheduleObservation struct {

	// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
	AtTheseHours []*float64 `json:"atTheseHours,omitempty" tf:"at_these_hours,omitempty"`

	// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
	AtTheseMinutes []*float64 `json:"atTheseMinutes,omitempty" tf:"at_these_minutes,omitempty"`

	// Specifies a list of days when the trigger should run. Valid values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	OnTheseDays []*string `json:"onTheseDays,omitempty" tf:"on_these_days,omitempty"`
}

type AppTriggerRecurrenceScheduleParameters struct {

	// Specifies a list of hours when the trigger should run. Valid values are between 0 and 23.
	// +kubebuilder:validation:Optional
	AtTheseHours []*float64 `json:"atTheseHours,omitempty" tf:"at_these_hours,omitempty"`

	// Specifies a list of minutes when the trigger should run. Valid values are between 0 and 59.
	// +kubebuilder:validation:Optional
	AtTheseMinutes []*float64 `json:"atTheseMinutes,omitempty" tf:"at_these_minutes,omitempty"`

	// Specifies a list of days when the trigger should run. Valid values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	// +kubebuilder:validation:Optional
	OnTheseDays []*string `json:"onTheseDays,omitempty" tf:"on_these_days,omitempty"`
}

// AppTriggerRecurrenceSpec defines the desired state of AppTriggerRecurrence
type AppTriggerRecurrenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppTriggerRecurrenceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AppTriggerRecurrenceInitParameters `json:"initProvider,omitempty"`
}

// AppTriggerRecurrenceStatus defines the observed state of AppTriggerRecurrence.
type AppTriggerRecurrenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppTriggerRecurrenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerRecurrence is the Schema for the AppTriggerRecurrences API. Manages a Recurrence Trigger within a Logic App Workflow
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppTriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.frequency) || has(self.initProvider.frequency)",message="frequency is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interval) || has(self.initProvider.interval)",message="interval is a required parameter"
	Spec   AppTriggerRecurrenceSpec   `json:"spec"`
	Status AppTriggerRecurrenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerRecurrenceList contains a list of AppTriggerRecurrences
type AppTriggerRecurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppTriggerRecurrence `json:"items"`
}

// Repository type metadata.
var (
	AppTriggerRecurrence_Kind             = "AppTriggerRecurrence"
	AppTriggerRecurrence_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppTriggerRecurrence_Kind}.String()
	AppTriggerRecurrence_KindAPIVersion   = AppTriggerRecurrence_Kind + "." + CRDGroupVersion.String()
	AppTriggerRecurrence_GroupVersionKind = CRDGroupVersion.WithKind(AppTriggerRecurrence_Kind)
)

func init() {
	SchemeBuilder.Register(&AppTriggerRecurrence{}, &AppTriggerRecurrenceList{})
}
