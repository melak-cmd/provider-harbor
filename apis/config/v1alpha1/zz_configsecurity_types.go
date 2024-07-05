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

type ConfigSecurityInitParameters struct {

	// 123", "CVE-145"] or ["CVE-123"]
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// (Number) The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt *float64 `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`
}

type ConfigSecurityObservation struct {

	// (String) Time of creation of the list.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// 123", "CVE-145"] or ["CVE-123"]
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// (Number) The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt *float64 `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Time of update of the list.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ConfigSecurityParameters struct {

	// 123", "CVE-145"] or ["CVE-123"]
	// +kubebuilder:validation:Optional
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// (Number) The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	// +kubebuilder:validation:Optional
	ExpiresAt *float64 `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`
}

// ConfigSecuritySpec defines the desired state of ConfigSecurity
type ConfigSecuritySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigSecurityParameters `json:"forProvider"`
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
	InitProvider ConfigSecurityInitParameters `json:"initProvider,omitempty"`
}

// ConfigSecurityStatus defines the observed state of ConfigSecurity.
type ConfigSecurityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigSecurityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigSecurity is the Schema for the ConfigSecuritys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type ConfigSecurity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cveAllowlist) || (has(self.initProvider) && has(self.initProvider.cveAllowlist))",message="spec.forProvider.cveAllowlist is a required parameter"
	Spec   ConfigSecuritySpec   `json:"spec"`
	Status ConfigSecurityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigSecurityList contains a list of ConfigSecuritys
type ConfigSecurityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigSecurity `json:"items"`
}

// Repository type metadata.
var (
	ConfigSecurity_Kind             = "ConfigSecurity"
	ConfigSecurity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigSecurity_Kind}.String()
	ConfigSecurity_KindAPIVersion   = ConfigSecurity_Kind + "." + CRDGroupVersion.String()
	ConfigSecurity_GroupVersionKind = CRDGroupVersion.WithKind(ConfigSecurity_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigSecurity{}, &ConfigSecurityList{})
}
