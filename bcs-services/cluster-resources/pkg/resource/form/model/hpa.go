/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

// HPA 表单化建模
type HPA struct {
	Metadata Metadata `structs:"metadata"`
	Spec     HPASpec  `structs:"spec"`
}

// HPASpec ...
type HPASpec struct {
	Ref          HPATargetRef       `structs:"ref"`
	Resource     ResourceMetric     `structs:"resource"`
	ContainerRes ContainerResMetric `structs:"containerRes"`
	External     ExternalMetric     `structs:"external"`
	Object       ObjectMetric       `structs:"object"`
	Pod          PodMetric          `structs:"pod"`
}

// HPATargetRef ...
type HPATargetRef struct {
	Kind        string `structs:"kind"`
	APIVersion  string `structs:"apiVersion"`
	ResName     string `structs:"resName"`
	MinReplicas int64  `structs:"minReplicas"`
	MaxReplicas int64  `structs:"maxReplicas"`
}

// ResourceMetric ...
type ResourceMetric struct {
	Items []ResourceMetricItem `structs:"items"`
}

// ResourceMetricItem ...
type ResourceMetricItem struct {
	Name  string `structs:"name"`
	Type  string `structs:"type"`
	Value string `structs:"value"`
}

// ContainerResMetric NOTE ContainerResource 指标目前未启用
type ContainerResMetric struct {
	Items []ContainerResMetricItem `structs:"items"`
}

// ContainerResMetricItem ...
type ContainerResMetricItem struct {
	Name          string `structs:"name"`
	ContainerName string `structs:"containerName"`
	Type          string `structs:"type"`
	Value         string `structs:"value"`
}

// ExternalMetric ...
type ExternalMetric struct {
	Items []ExternalMetricItem `structs:"items"`
}

// ExternalMetricItem ...
type ExternalMetricItem struct {
	Name     string         `structs:"name"`
	Type     string         `structs:"type"`
	Value    string         `structs:"value"`
	Selector MetricSelector `structs:"selector"`
}

// MetricSelector ...
type MetricSelector struct {
	Expressions []ExpSelector   `structs:"expressions"`
	Labels      []LabelSelector `structs:"labels"`
}

// ObjectMetric ...
type ObjectMetric struct {
	Items []ObjectMetricItem `structs:"items"`
}

// ObjectMetricItem ...
type ObjectMetricItem struct {
	Name       string         `structs:"name"`
	APIVersion string         `structs:"apiVersion"`
	Kind       string         `structs:"kind"`
	ResName    string         `structs:"resName"`
	Type       string         `structs:"type"`
	Value      string         `structs:"value"`
	Selector   MetricSelector `structs:"selector"`
}

// PodMetric ...
type PodMetric struct {
	Items []PodMetricItem `structs:"items"`
}

// PodMetricItem ...
type PodMetricItem struct {
	Name     string         `structs:"name"`
	Type     string         `structs:"type"`
	Value    string         `structs:"value"`
	Selector MetricSelector `structs:"selector"`
}
