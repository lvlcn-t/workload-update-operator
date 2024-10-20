package controller

import (
	"context"

	"github.com/lvlcn-t/loggerhead/logger"
	apiv1 "github.com/lvlcn-t/workload-update-operator/controller/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = (*workloadReconciler)(nil)

type workloadReconciler struct {
	client.Client
	config apiv1.OperatorConfig
}

func NewReconciler(ctx context.Context, c client.Client) (reconcile.Reconciler, error) {
	cfg, err := apiv1.LoadConfig(ctx, c)
	if err != nil {
		return nil, err
	}
	return &workloadReconciler{Client: c, config: *cfg}, nil
}

func (r *workloadReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log := logger.FromContext(ctx)
	log.InfoContext(ctx, "Reconciling", "request", req)
	return reconcile.Result{}, nil
}
