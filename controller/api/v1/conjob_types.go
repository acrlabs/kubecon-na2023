package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConJobSpec defines the desired state of ConJob
type ConJobSpec struct {
	// Foo is an example field of ConJob. Edit conjob_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ConJobStatus defines the observed state of ConJob
type ConJobStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConJob is the Schema for the conjobs API
type ConJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConJobSpec   `json:"spec,omitempty"`
	Status ConJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConJobList contains a list of ConJob
type ConJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConJob{}, &ConJobList{})
}