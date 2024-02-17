package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get defines how to get resources.
type Get struct {
	// Timeout for the operation. Overrides the global timeout set in the Configuration.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Resource type.
	Resource string `json:"resource"`

	// ObjectLabelsSelector determines the selection process of referenced objects.
	ObjectLabelsSelector `json:",inline"`
}
