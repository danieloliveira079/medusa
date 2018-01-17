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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Medusa is a specification for a Medusa resource
type Medusa struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MedusaSpec   `json:"spec"`
	Status MedusaStatus `json:"status"`
}

// MedusaSpec is the spec for a Medusa resource
type MedusaSpec struct {
	Name        string
	CurrrentSHA string `json:"current_sha"`
	Pusher      string
	RepoURL     string `json:"repo_url"`
	RepoName    string `json:"repo_name"`
	Destroy     bool
	//TODO - Must be an array of objects due to the nature of different
	// possible templates. Deployments, Services, Ingress, etc.
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// MedusaStatus is the status for a Medusa resource
type MedusaStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MedusaList is a list of Medusa resources
type MedusaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Medusa `json:"items"`
}
