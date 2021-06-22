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

// Code generated by ack-generate. DO NOT EDIT.

package template

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ses"

	svcapitypes "github.com/crossplane/provider-aws/apis/ses/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetTemplateInput returns input for read
// operation.
func GenerateGetTemplateInput(cr *svcapitypes.Template) *svcsdk.GetTemplateInput {
	res := &svcsdk.GetTemplateInput{}

	return res
}

// GenerateTemplate returns the current state in the form of *svcapitypes.Template.
func GenerateTemplate(resp *svcsdk.GetTemplateOutput) *svcapitypes.Template {
	cr := &svcapitypes.Template{}

	return cr
}

// GenerateCreateTemplateInput returns a create input.
func GenerateCreateTemplateInput(cr *svcapitypes.Template) *svcsdk.CreateTemplateInput {
	res := &svcsdk.CreateTemplateInput{}

	if cr.Spec.ForProvider.Template != nil {
		f0 := &svcsdk.Template{}
		if cr.Spec.ForProvider.Template.HTMLPart != nil {
			f0.SetHtmlPart(*cr.Spec.ForProvider.Template.HTMLPart)
		}
		if cr.Spec.ForProvider.Template.SubjectPart != nil {
			f0.SetSubjectPart(*cr.Spec.ForProvider.Template.SubjectPart)
		}
		if cr.Spec.ForProvider.Template.TemplateName != nil {
			f0.SetTemplateName(*cr.Spec.ForProvider.Template.TemplateName)
		}
		if cr.Spec.ForProvider.Template.TextPart != nil {
			f0.SetTextPart(*cr.Spec.ForProvider.Template.TextPart)
		}
		res.SetTemplate(f0)
	}

	return res
}

// GenerateUpdateTemplateInput returns an update input.
func GenerateUpdateTemplateInput(cr *svcapitypes.Template) *svcsdk.UpdateTemplateInput {
	res := &svcsdk.UpdateTemplateInput{}

	if cr.Spec.ForProvider.Template != nil {
		f0 := &svcsdk.Template{}
		if cr.Spec.ForProvider.Template.HTMLPart != nil {
			f0.SetHtmlPart(*cr.Spec.ForProvider.Template.HTMLPart)
		}
		if cr.Spec.ForProvider.Template.SubjectPart != nil {
			f0.SetSubjectPart(*cr.Spec.ForProvider.Template.SubjectPart)
		}
		if cr.Spec.ForProvider.Template.TemplateName != nil {
			f0.SetTemplateName(*cr.Spec.ForProvider.Template.TemplateName)
		}
		if cr.Spec.ForProvider.Template.TextPart != nil {
			f0.SetTextPart(*cr.Spec.ForProvider.Template.TextPart)
		}
		res.SetTemplate(f0)
	}

	return res
}

// GenerateDeleteTemplateInput returns a deletion input.
func GenerateDeleteTemplateInput(cr *svcapitypes.Template) *svcsdk.DeleteTemplateInput {
	res := &svcsdk.DeleteTemplateInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
