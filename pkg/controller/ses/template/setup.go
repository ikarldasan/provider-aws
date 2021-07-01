/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package template

import (
	"context"
	"fmt"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	svcsdk "github.com/aws/aws-sdk-go/service/ses"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	aws "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/google/go-cmp/cmp"

	svcapitypes "github.com/crossplane/provider-aws/apis/ses/v1alpha1"
)

// SetupTemplate adds a controller that reconciles Template.
func SetupTemplate(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(svcapitypes.TemplateGroupKind)
	opts := []option{
		func(e *external) {
			// e.preCreate = preCreate
			e.postCreate = postCreate
			e.preObserve = preObserve
			e.postObserve = postObserve
			// e.preUpdate = preUpdate
			e.preDelete = preDelete
			// e.postDelete = postDelete
			e.isUpToDate = isUpToDate

		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.Template{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.TemplateGroupVersionKind),
			managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient())),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func preObserve(_ context.Context, cr *svcapitypes.Template, obj *svcsdk.GetTemplateInput) error {
	fmt.Println("---------PRE OBSERVE CALLED---------")

	fmt.Println("cr: ", cr)
	fmt.Println("cr: ", cr.Status)
	fmt.Println("obj: ", obj)
	fmt.Println("obj.String: ", obj.String())
	fmt.Println("obj.GoString: ", obj.GoString())
	fmt.Println("aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))

	obj.TemplateName = aws.String(meta.GetExternalName(cr))
	fmt.Println("obj.TemplateName: ", *obj.TemplateName)

	// meta.SetExternalName(cr, obj.TemplateName)
	return nil
}

func postObserve(_ context.Context, cr *svcapitypes.Template, obj *svcsdk.GetTemplateOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	fmt.Println("---------POST OBSERVE CALLED---------")
	fmt.Println("cr: ", cr)
	fmt.Println("obj: ", obj)
	fmt.Println("obj.String: ", obj.String())
	fmt.Println("obj.GoString: ", obj.GoString())
	fmt.Println("aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	if obj.Template != nil {
		fmt.Println("------------------------------")
	}

	cr.SetConditions(xpv1.Available())
	return obs, err
}

// func preCreate(_ context.Context, cr *svcapitypes.Template, obj *svcsdk.CreateTemplateInput) error {
// 	fmt.Println("---------PRE CREATE CALLED---------")
// 	fmt.Println("aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))
// 	// obj.Template.TemplateName = aws.String(meta.GetExternalName(cr))
// 	fmt.Println("obj: ", obj)
// 	return nil
// }

// func preUpdate(_ context.Context, cr *svcapitypes.Template, obj *svcsdk.UpdateTemplateInput) error {
// 	fmt.Println("---------PRE UPDATE CALLED---------")
// 	// obj.TemplateName = aws.String(meta.GetExternalName(cr))
// 	obj.Template.TemplateName = aws.String(meta.GetExternalName(cr))
// 	return nil
// }

func preDelete(context context.Context, cr *svcapitypes.Template, obj *svcsdk.DeleteTemplateInput) (bool, error) {
	fmt.Println("---------PRE DELETE CALLED---------")
	// obj.TemplateName = aws.String(meta.GetExternalName(cr))
	obj.TemplateName = aws.String(meta.GetExternalName(cr))
	return false, nil
}

func postCreate(context context.Context, cr *svcapitypes.Template, obj *svcsdk.CreateTemplateOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	fmt.Println("---------POST CREATE CALLED---------")
	fmt.Println("cr: ", cr)
	fmt.Println("obj: ", obj)
	fmt.Println("obj.String: ", obj.String())
	fmt.Println("obj.GoString: ", obj.GoString())
	fmt.Println("Before: aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))

	if err != nil {
		return managed.ExternalCreation{}, err
	}

	meta.SetExternalName(cr, aws.StringValue(cr.Spec.ForProvider.Template.TemplateName))
	// meta.SetExternalName(cr, aws.StringValue())
	fmt.Println("After: meta.GetExternalName(cr): ", meta.GetExternalName(cr))
	fmt.Println("After: aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))
	cre.ExternalNameAssigned = true
	return cre, nil
}

func postDelete(_ context.Context, cr *svcapitypes.Template, obj *svcsdk.DeleteTemplateOutput, err error) error {
	fmt.Println("---------POST DELETE CALLED---------")

	fmt.Println("cr: ", cr)
	fmt.Println("cr.Status: ", cr.Status)
	fmt.Println("obj: ", obj)
	fmt.Println("obj.String: ", obj.String())
	fmt.Println("obj.GoString: ", obj.GoString())
	fmt.Println("Before: aws.StringValue(cr.Spec.ForProvider.Template.TemplateName): ", aws.StringValue(cr.Spec.ForProvider.Template.TemplateName))
	fmt.Println("Before: aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))

	// if err != nil {
	// 	return err
	// }

	// cr = nil

	// meta.SetExternalName(cr, "")
	// fmt.Println("After: meta.GetExternalName(cr): ", meta.GetExternalName(cr))
	// fmt.Println("After: aws.String(meta.GetExternalName(cr)): ", aws.String(meta.GetExternalName(cr)))
	return err
}

func isUpToDate(cr *svcapitypes.Template, obj *svcsdk.GetTemplateOutput) (bool, error) {
	if !cmp.Equal(cr.Spec.ForProvider.Template.TemplateName, obj.Template.TemplateName) {
		return false, nil
	}
	if !cmp.Equal(cr.Spec.ForProvider.Template.HTMLPart, obj.Template.HtmlPart) {
		return false, nil
	}
	if !cmp.Equal(cr.Spec.ForProvider.Template.SubjectPart, obj.Template.SubjectPart) {
		return false, nil
	}
	if !cmp.Equal(cr.Spec.ForProvider.Template.TextPart, obj.Template.TextPart) {
		return false, nil
	}
	return true, nil
}
