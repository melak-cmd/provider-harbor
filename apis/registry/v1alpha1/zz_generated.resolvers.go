/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	common "github.com/globallogicuki/provider-harbor/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Replication.
func (mg *Replication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RegistryID),
		Extract:      common.ExtractRegistryID(),
		Reference:    mg.Spec.ForProvider.RegistryIDRef,
		Selector:     mg.Spec.ForProvider.RegistryIDSelector,
		To: reference.To{
			List:    &RegistryList{},
			Managed: &Registry{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RegistryID")
	}
	mg.Spec.ForProvider.RegistryID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegistryIDRef = rsp.ResolvedReference

	return nil
}
