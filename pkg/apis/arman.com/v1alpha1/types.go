package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BlackMan struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec BlackManSpec `json:"spec,omitempty"`
	Status ArmanStatus `json:"status,omitempty"`
}

type ArmanStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

type BlackManSpec struct {
	Name     string `json:"name"`
	Region   string `json:"region"`
	Version  string `json:"version"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BlackManList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items []BlackMan `json:"items,omitempty"`
}
