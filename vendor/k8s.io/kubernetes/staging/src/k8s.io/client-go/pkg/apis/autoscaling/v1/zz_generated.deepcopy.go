// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_CrossVersionObjectReference, InType: reflect.TypeOf(&CrossVersionObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_HorizontalPodAutoscaler, InType: reflect.TypeOf(&HorizontalPodAutoscaler{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_HorizontalPodAutoscalerList, InType: reflect.TypeOf(&HorizontalPodAutoscalerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_HorizontalPodAutoscalerSpec, InType: reflect.TypeOf(&HorizontalPodAutoscalerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_HorizontalPodAutoscalerStatus, InType: reflect.TypeOf(&HorizontalPodAutoscalerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Scale, InType: reflect.TypeOf(&Scale{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ScaleSpec, InType: reflect.TypeOf(&ScaleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ScaleStatus, InType: reflect.TypeOf(&ScaleStatus{})},
	)
}

func DeepCopy_v1_CrossVersionObjectReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CrossVersionObjectReference)
		out := out.(*CrossVersionObjectReference)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_HorizontalPodAutoscaler(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscaler)
		out := out.(*HorizontalPodAutoscaler)
		*out = *in
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_HorizontalPodAutoscalerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerList)
		out := out.(*HorizontalPodAutoscalerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]HorizontalPodAutoscaler, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_HorizontalPodAutoscalerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerSpec)
		out := out.(*HorizontalPodAutoscalerSpec)
		*out = *in
		if in.MinReplicas != nil {
			in, out := &in.MinReplicas, &out.MinReplicas
			*out = new(int32)
			**out = **in
		}
		if in.TargetCPUUtilizationPercentage != nil {
			in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_HorizontalPodAutoscalerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerStatus)
		out := out.(*HorizontalPodAutoscalerStatus)
		*out = *in
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		}
		if in.LastScaleTime != nil {
			in, out := &in.LastScaleTime, &out.LastScaleTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.CurrentCPUUtilizationPercentage != nil {
			in, out := &in.CurrentCPUUtilizationPercentage, &out.CurrentCPUUtilizationPercentage
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_Scale(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Scale)
		out := out.(*Scale)
		*out = *in
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_ScaleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleSpec)
		out := out.(*ScaleSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_ScaleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleStatus)
		out := out.(*ScaleStatus)
		*out = *in
		return nil
	}
}