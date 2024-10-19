package controller

import (
	"context"

	"github.com/lvlcn-t/loggerhead/logger"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = (*workloadReconciler)(nil)

type workloadReconciler struct {
	client.Client
}

func NewReconciler(c client.Client) reconcile.Reconciler {
	return &workloadReconciler{Client: c}
}

func (r *workloadReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log := logger.FromContext(ctx)
	log.InfoContext(ctx, "Reconciling", "request", req)
	return reconcile.Result{}, nil
}
