/*
Copyright 2020 Alexei Ledenev.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KillContainerSpec defines the desired state of KillContainer
type KillContainerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Selector is used to select pods that are used to inject chaos action.
	Selector SelectorSpec `json:"selector"`

	// Signal POSIX signal, see more https://man7.org/linux/man-pages/man7/signal.7.html
	// +kubebuilder:validation:Maximum:=1
	// +kubebuilder:validation:Maximum:=31
	// +optional
	Signal int `json:"signal,omitempty"`

	// Scheduler defines some schedule rules to
	// control the running time of the chaos experiment about pods.
	Scheduler *SchedulerSpec `json:"scheduler,omitempty"`
}

// KillContainerStatus defines the observed state of KillContainer
type KillContainerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KillContainer is the Schema for the killcontainers API
type KillContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KillContainerSpec   `json:"spec,omitempty"`
	Status KillContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KillContainerList contains a list of KillContainer
type KillContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KillContainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KillContainer{}, &KillContainerList{})
}
