package v1

import (
	"context"
	"errors"

	"github.com/lvlcn-t/go-kit/env"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// OperatorConfig is the configuration for the operator.
type OperatorConfig struct {
	// MaintainanceWindow is the maintenance window the operator allows changes to workloads.
	MaintainanceWindow MaintainanceWindow `json:"maintainanceWindow" yaml:"maintainanceWindow" mapstructure:"maintainanceWindow" validate:"required"`
}

// MaintainanceWindow represents the maintenance window.
type MaintainanceWindow struct {
	// Start is the start time of the maintenance window.
	// Allowed formats are [time.TimeOnly] and "15:04:05Z07:00".
	Start Time `json:"start" yaml:"start" mapstructure:"start"`
	// End is the end time of the maintenance window.
	// Allowed formats are [time.TimeOnly] and "15:04:05Z07:00".
	End Time `json:"end" yaml:"end" mapstructure:"end"`
}

// Validate checks if the maintenance window is valid.
// It implements the [config.Validator] interface.
func (m *MaintainanceWindow) Validate() error {
	var err error
	if m.Start.IsZero() {
		err = errors.New("start time cannot be zero")
	}
	if m.End.IsZero() {
		err = errors.Join(err, errors.New("end time cannot be zero"))
	}
	if m.Start.After(m.End.Time) {
		err = errors.Join(err, errors.New("start time cannot be after end time"))
	}
	return err
}

type WorkloadUpdateConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Spec              OperatorConfig `json:"spec" yaml:"spec"`
}

// DeepCopyObject implements the [runtime.Object] interface.
func (o *WorkloadUpdateConfig) DeepCopyObject() runtime.Object {
	out := &WorkloadUpdateConfig{}
	o.DeepCopyInto(&out.ObjectMeta)
	return out
}

// LoadConfig loads the operator configuration from the Kubernetes API.
func LoadConfig(ctx context.Context, c client.Client) (*OperatorConfig, error) {
	namespace := env.GetWithFallback("WORKLOAD_UPDATE_OPERATOR_NAMESPACE", metav1.NamespaceDefault)
	cfg := &WorkloadUpdateConfig{}
	if err := c.Get(ctx, client.ObjectKey{Name: "workload-update-operator", Namespace: namespace}, cfg); err != nil {
		return nil, err
	}
	return &cfg.Spec, nil
}
