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

type ProjectInitParameters struct {

	// 123", "CVE-145"] or ["CVE-123"]
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// empty)
	DeploymentSecurity *string `json:"deploymentSecurity,omitempty" tf:"deployment_security,omitempty"`

	// (Boolean) Enables Content Trust for project. When enabled it queries the embedded docker notary server. (Default: false).
	EnableContentTrust *bool `json:"enableContentTrust,omitempty" tf:"enable_content_trust,omitempty"`

	// (Boolean) Enables Content Trust Cosign for project. When enabled it queries Cosign. (Default: false)
	EnableContentTrustCosign *bool `json:"enableContentTrustCosign,omitempty" tf:"enable_content_trust_cosign,omitempty"`

	// (Boolean) A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String) The name of the project that will be created in harbor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) The project will be public accessibility.(Default: false)
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// (Number) The storage quota of the project in GB's.
	StorageQuota *float64 `json:"storageQuota,omitempty" tf:"storage_quota,omitempty"`

	// (Boolean) Images will be scanned for vulnerabilities when push to harbor. (Default: true)
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

type ProjectObservation struct {

	// 123", "CVE-145"] or ["CVE-123"]
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// empty)
	DeploymentSecurity *string `json:"deploymentSecurity,omitempty" tf:"deployment_security,omitempty"`

	// (Boolean) Enables Content Trust for project. When enabled it queries the embedded docker notary server. (Default: false).
	EnableContentTrust *bool `json:"enableContentTrust,omitempty" tf:"enable_content_trust,omitempty"`

	// (Boolean) Enables Content Trust Cosign for project. When enabled it queries Cosign. (Default: false)
	EnableContentTrustCosign *bool `json:"enableContentTrustCosign,omitempty" tf:"enable_content_trust_cosign,omitempty"`

	// (Boolean) A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the project that will be created in harbor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The project id of this resource.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Boolean) The project will be public accessibility.(Default: false)
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// (Number) To enable project as Proxy Cache.
	RegistryID *float64 `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// (Number) The storage quota of the project in GB's.
	StorageQuota *float64 `json:"storageQuota,omitempty" tf:"storage_quota,omitempty"`

	// (Boolean) Images will be scanned for vulnerabilities when push to harbor. (Default: true)
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

type ProjectParameters struct {

	// 123", "CVE-145"] or ["CVE-123"]
	// +kubebuilder:validation:Optional
	CveAllowlist []*string `json:"cveAllowlist,omitempty" tf:"cve_allowlist,omitempty"`

	// empty)
	// +kubebuilder:validation:Optional
	DeploymentSecurity *string `json:"deploymentSecurity,omitempty" tf:"deployment_security,omitempty"`

	// (Boolean) Enables Content Trust for project. When enabled it queries the embedded docker notary server. (Default: false).
	// +kubebuilder:validation:Optional
	EnableContentTrust *bool `json:"enableContentTrust,omitempty" tf:"enable_content_trust,omitempty"`

	// (Boolean) Enables Content Trust Cosign for project. When enabled it queries Cosign. (Default: false)
	// +kubebuilder:validation:Optional
	EnableContentTrustCosign *bool `json:"enableContentTrustCosign,omitempty" tf:"enable_content_trust_cosign,omitempty"`

	// (Boolean) A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String) The name of the project that will be created in harbor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) The project will be public accessibility.(Default: false)
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// (Number) To enable project as Proxy Cache.
	// +crossplane:generate:reference:type=github.com/globallogicuki/provider-harbor/apis/registry/v1alpha1.Registry
	// +crossplane:generate:reference:extractor=github.com/globallogicuki/provider-harbor/config/common.ExtractRegistryId()
	// +kubebuilder:validation:Optional
	RegistryID *float64 `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a Registry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a Registry in registry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	// (Number) The storage quota of the project in GB's.
	// +kubebuilder:validation:Optional
	StorageQuota *float64 `json:"storageQuota,omitempty" tf:"storage_quota,omitempty"`

	// (Boolean) Images will be scanned for vulnerabilities when push to harbor. (Default: true)
	// +kubebuilder:validation:Optional
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
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
	InitProvider ProjectInitParameters `json:"initProvider,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
