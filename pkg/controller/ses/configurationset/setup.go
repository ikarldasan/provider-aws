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

package configurationset

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ses"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/ses/v1alpha1"
)

// SetupConfigurationSet adds a controller that reconciles ConfigurationSet.
func SetupConfigurationSet(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(svcapitypes.ConfigurationSetGroupKind)
	opts := []option{
		func(e *external) {
			e.preObserve = preObserve
			// e.postObserve = postObserve
			// e.preCreate = preCreate
			// e.preUpdate = preUpdate
			// e.preDelete = preDelete
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&svcapitypes.ConfigurationSet{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.ConfigurationSetGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func preObserve(context context.Context, cr *svcapitypes.ConfigurationSet, obj *svcsdk.DescribeConfigurationSetInput) error {
	obj.ConfigurationSetName = aws.String(meta.GetExternalName(cr))
	return nil
}

// func postObserve(_ context.Context, cr *svcapitypes.ConfigurationSet, obj *svcsdk.GetTemplateOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
// 	fmt.Println("---------POST OBSERVE CALLED---------")
// 	obj.ConfigurationSetName = aws.String(meta.GetExternalName(cr))
// 	return managed.ExternalObservation{}, nil
// }

// func preCreate(_ context.Context, cr *svcapitypes.ConfigurationSet, obj *svcsdk.CreateTemplateInput) error {
// 	fmt.Println("---------PRE CREATE CALLED---------")
// 	return nil
// }

// func preUpdate(_ context.Context, cr *svcapitypes.ConfigurationSet, obj *svcsdk.UpdateTemplateInput) error {
// 	return nil
// }
