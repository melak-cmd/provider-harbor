// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FiltersInitParameters struct {

	// (String) Matches or excludes the result. Can be one of the following. matches, excludes
	Decoration *string `json:"decoration,omitempty" tf:"decoration,omitempty"`

	// (List of String) Filter on the resource according to labels.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the replication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Filter on the resource type. Can be one of the following types. chart, artifact
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// (String) Filter on the tag/version of the resource.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FiltersObservation struct {

	// (String) Matches or excludes the result. Can be one of the following. matches, excludes
	Decoration *string `json:"decoration,omitempty" tf:"decoration,omitempty"`

	// (List of String) Filter on the resource according to labels.
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the replication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Filter on the resource type. Can be one of the following types. chart, artifact
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// (String) Filter on the tag/version of the resource.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FiltersParameters struct {

	// (String) Matches or excludes the result. Can be one of the following. matches, excludes
	// +kubebuilder:validation:Optional
	Decoration *string `json:"decoration,omitempty" tf:"decoration,omitempty"`

	// (List of String) Filter on the resource according to labels.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the replication.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Filter on the resource type. Can be one of the following types. chart, artifact
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// (String) Filter on the tag/version of the resource.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type ReplicationInitParameters struct {

	// (String)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (Boolean) Specify whether to delete the remote resources when locally deleted. (Default: false)
	Deletion *bool `json:"deletion,omitempty" tf:"deletion,omitempty"`

	// (String) Description of the replication policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Specify the destination namespace. if empty, the resource will be put under the same namespace as the source.
	DestNamespace *string `json:"destNamespace,omitempty" tf:"dest_namespace,omitempty"`

	// 1 to 3 are valid values in the harbor API. A value of -1 will 'Flatten All Levels', 0 means 'No Flattening', 1 'Flatten 1 Level', 2 'Flatten 2 Levels', 3 'Flatten 3 Levels' (Default: -1, see Replication Rules for more details)
	DestNamespaceReplace *float64 `json:"destNamespaceReplace,omitempty" tf:"dest_namespace_replace,omitempty"`

	// (Boolean) Specify whether the replication is enabled. (Default: true)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Specify whether to execute the replication rule if new or modified. (Default: false)
	ExecuteOnChanged *bool `json:"executeOnChanged,omitempty" tf:"execute_on_changed,omitempty"`

	// (Block Set) (see below for nested schema)
	Filters []FiltersInitParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// (String) The name of the replication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Specify whether to override the resources at the destination if a resources with the same name exist. (Default: true)
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// (String) The scheduled time of when the container register will be push / pull. In cron base format. Hourly "0 0 * * * *", Daily "0 0 0 * * *", Monthly "0 0 0 * * 0". Can be one of the following: event_based, manual, cron format (Default: manual)
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// 1 (unlimited).
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`
}

type ReplicationObservation struct {

	// (String)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (Boolean) Specify whether to delete the remote resources when locally deleted. (Default: false)
	Deletion *bool `json:"deletion,omitempty" tf:"deletion,omitempty"`

	// (String) Description of the replication policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Specify the destination namespace. if empty, the resource will be put under the same namespace as the source.
	DestNamespace *string `json:"destNamespace,omitempty" tf:"dest_namespace,omitempty"`

	// 1 to 3 are valid values in the harbor API. A value of -1 will 'Flatten All Levels', 0 means 'No Flattening', 1 'Flatten 1 Level', 2 'Flatten 2 Levels', 3 'Flatten 3 Levels' (Default: -1, see Replication Rules for more details)
	DestNamespaceReplace *float64 `json:"destNamespaceReplace,omitempty" tf:"dest_namespace_replace,omitempty"`

	// (Boolean) Specify whether the replication is enabled. (Default: true)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Specify whether to execute the replication rule if new or modified. (Default: false)
	ExecuteOnChanged *bool `json:"executeOnChanged,omitempty" tf:"execute_on_changed,omitempty"`

	// (Block Set) (see below for nested schema)
	Filters []FiltersObservation `json:"filters,omitempty" tf:"filters,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the replication.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Specify whether to override the resources at the destination if a resources with the same name exist. (Default: true)
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// (Number) The registry ID of the Registry Endpoint.
	RegistryID *float64 `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// (Number)
	ReplicationPolicyID *float64 `json:"replicationPolicyId,omitempty" tf:"replication_policy_id,omitempty"`

	// (String) The scheduled time of when the container register will be push / pull. In cron base format. Hourly "0 0 * * * *", Daily "0 0 0 * * *", Monthly "0 0 0 * * 0". Can be one of the following: event_based, manual, cron format (Default: manual)
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// 1 (unlimited).
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`
}

type ReplicationParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (Boolean) Specify whether to delete the remote resources when locally deleted. (Default: false)
	// +kubebuilder:validation:Optional
	Deletion *bool `json:"deletion,omitempty" tf:"deletion,omitempty"`

	// (String) Description of the replication policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Specify the destination namespace. if empty, the resource will be put under the same namespace as the source.
	// +kubebuilder:validation:Optional
	DestNamespace *string `json:"destNamespace,omitempty" tf:"dest_namespace,omitempty"`

	// 1 to 3 are valid values in the harbor API. A value of -1 will 'Flatten All Levels', 0 means 'No Flattening', 1 'Flatten 1 Level', 2 'Flatten 2 Levels', 3 'Flatten 3 Levels' (Default: -1, see Replication Rules for more details)
	// +kubebuilder:validation:Optional
	DestNamespaceReplace *float64 `json:"destNamespaceReplace,omitempty" tf:"dest_namespace_replace,omitempty"`

	// (Boolean) Specify whether the replication is enabled. (Default: true)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Boolean) Specify whether to execute the replication rule if new or modified. (Default: false)
	// +kubebuilder:validation:Optional
	ExecuteOnChanged *bool `json:"executeOnChanged,omitempty" tf:"execute_on_changed,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// (String) The name of the replication.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Specify whether to override the resources at the destination if a resources with the same name exist. (Default: true)
	// +kubebuilder:validation:Optional
	Override *bool `json:"override,omitempty" tf:"override,omitempty"`

	// (Number) The registry ID of the Registry Endpoint.
	// +crossplane:generate:reference:type=github.com/globallogicuki/provider-harbor/apis/registry/v1alpha1.Registry
	// +crossplane:generate:reference:extractor=github.com/globallogicuki/provider-harbor/config/common.ExtractRegistryID()
	// +kubebuilder:validation:Optional
	RegistryID *float64 `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a Registry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a Registry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	// (String) The scheduled time of when the container register will be push / pull. In cron base format. Hourly "0 0 * * * *", Daily "0 0 0 * * *", Monthly "0 0 0 * * 0". Can be one of the following: event_based, manual, cron format (Default: manual)
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// 1 (unlimited).
	// +kubebuilder:validation:Optional
	Speed *float64 `json:"speed,omitempty" tf:"speed,omitempty"`
}

// ReplicationSpec defines the desired state of Replication
type ReplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ReplicationInitParameters `json:"initProvider,omitempty"`
}

// ReplicationStatus defines the observed state of Replication.
type ReplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Replication is the Schema for the Replications API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type Replication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ReplicationSpec   `json:"spec"`
	Status ReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationList contains a list of Replications
type ReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Replication `json:"items"`
}

// Repository type metadata.
var (
	Replication_Kind             = "Replication"
	Replication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Replication_Kind}.String()
	Replication_KindAPIVersion   = Replication_Kind + "." + CRDGroupVersion.String()
	Replication_GroupVersionKind = CRDGroupVersion.WithKind(Replication_Kind)
)

func init() {
	SchemeBuilder.Register(&Replication{}, &ReplicationList{})
}
