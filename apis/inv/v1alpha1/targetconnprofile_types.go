/*
Copyright 2024 Nokia.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Encoding string

const (
	Encoding_Unknown   Encoding = "unknown"
	Encoding_JSON      Encoding = "JSON"
	Encoding_JSON_IETF Encoding = "JSON_IETF"
	Encoding_Bytes     Encoding = "bytes"
	Encoding_Protobuf  Encoding = "protobuf"
	Encoding_Ascii     Encoding = "ASCII"
	Encoding_Config    Encoding = "config"
)

type Protocol string

const (
	Protocol_Unknown Protocol = "unknown"
	Protocol_GNMI    Protocol = "gnmi"
	Protocol_NETCONF Protocol = "netconf"
	Protocol_NOOP    Protocol = "noop"
)

type CommitCandidate string

const (
	CommitCandidate_Candidate CommitCandidate = "candidate"
	CommitCandidate_Running   CommitCandidate = "running"
)

// TargetConnectionProfileSpec defines the desired state of TargetConnectionProfile
type TargetConnectionProfileSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="connectRetry is immutable"
	// +kubebuilder:default:="10s"
	ConnectRetry metav1.Duration `json:"connectRetry,omitempty" yaml:"connectRetry,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="timeout is immutable"
	// +kubebuilder:default:="10s"
	Timeout metav1.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="protocol is immutable"
	// +kubebuilder:validation:Enum=unknown;gnmi;netconf;noop;
	// +kubebuilder:default:="gnmi"
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="port is immutable"
	// +kubebuilder:default:=57400
	// Port defines the port on which the scan runs
	Port uint `json:"port" yaml:"port"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="encoding is immutable"
	// +kubebuilder:validation:Enum=unknown;JSON;JSON_IETF;bytes;protobuf;ASCII;config;
	// +kubebuilder:default:="ASCII"
	Encoding Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="preferredNetconfVersion is immutable"
	// +kubebuilder:validation:Enum="1.0";"1.1";
	// +kubebuilder:default:="1.0"
	PreferredNetconfVersion string `json:"preferredNetconfVersion,omitempty" yaml:"preferredNetconfVersion,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="insecure is immutable"
	// +kubebuilder:default:=false
	Insecure bool `json:"insecure,omitempty" yaml:"insecure,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="skipVerify is immutable"
	// +kubebuilder:default:=true
	SkipVerify bool `json:"skipVerify,omitempty" yaml:"skipVerify,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="includeNS is immutable"
	// +kubebuilder:default:=false
	IncludeNS bool `json:"includeNS,omitempty" yaml:"include-ns,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="operationWithNS is immutable"
	// +kubebuilder:default:=false
	OperationWithNS bool `json:"operationWithNS,omitempty" yaml:"operationWithNS,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UseOperationRemove is immutable"
	// +kubebuilder:default:=false
	UseOperationRemove bool `json:"useOperationRemove,omitempty" yaml:"useOperationRemove,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UseOperationRemove is immutable"
	// +kubebuilder:validation:Enum=candidate;running;
	// +kubebuilder:default:="candidate"
	CommitCandidate CommitCandidate `json:"commitCandidate,omitempty" yaml:"commitCandidate,omitempty"`
}

// +kubebuilder:object:root=true
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={sdc,inv}
// TargetConnectionProfile is the Schema for the TargetConnectionProfile API
// +k8s:openapi-gen=true
type TargetConnectionProfile struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec TargetConnectionProfileSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// TargetConnectionProfileList contains a list of TargetConnectionProfile
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TargetConnectionProfileList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []TargetConnectionProfile `json:"items" yaml:"items"`
}

func init() {
	SchemeBuilder.Register(&TargetConnectionProfile{}, &TargetConnectionProfileList{})
}

var (
	TargetConnectionProfileKind              = reflect.TypeOf(TargetConnectionProfile{}).Name()
	TargetConnectionProfileGroupKind         = schema.GroupKind{Group: SchemeGroupVersion.Group, Kind: TargetConnectionProfileKind}.String()
	TargetConnectionProfileKindAPIVersion    = TargetKind + "." + SchemeGroupVersion.String()
	TargetConnectionProfileGroupVersionKind = SchemeGroupVersion.WithKind(TargetConnectionProfileKind)
)
