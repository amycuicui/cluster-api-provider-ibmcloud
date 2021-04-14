/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha3

import clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"

const (
	// MachineAPIServerPodHealthyCondition reports a machine's kube-apiserver's operational status.
	MachineAPIServerPodHealthyCondition clusterv1.ConditionType = "APIServerPodHealthy"

	// MachineControllerManagerPodHealthyCondition reports a machine's kube-controller-manager's health status.
	MachineControllerManagerPodHealthyCondition clusterv1.ConditionType = "ControllerManagerPodHealthy"

	// MachineSchedulerPodHealthyCondition reports a machine's kube-scheduler's operational status.
	MachineSchedulerPodHealthyCondition clusterv1.ConditionType = "SchedulerPodHealthy"

	// MachineEtcdPodHealthyCondition reports a machine's etcd pod's operational status.
	// NOTE: This conditions exists only if a stacked etcd cluster is used.
	MachineEtcdPodHealthyCondition clusterv1.ConditionType = "EtcdPodHealthy"

	// MachineEtcdMemberHealthyCondition report the machine's etcd member's health status.
	// NOTE: This conditions exists only if a stacked etcd cluster is used.
	MachineEtcdMemberHealthyCondition clusterv1.ConditionType = "EtcdMemberHealthy"
)
