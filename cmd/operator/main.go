package main

import (
	"context"
	"time"

	"github.com/lvlcn-t/loggerhead/logger"
	"github.com/lvlcn-t/workload-update-operator/controller"
	appsv1 "k8s.io/api/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// version is the version of the operator.
// This is set at compile time using the -ldflags "-X main.version=$VERSION" flag.
var version string

var (
	// gracePeriod is the duration given to the operator to gracefully shutdown.
	gracePeriod = 30 * time.Second
)

func main() {
	ctx := logger.IntoContext(context.Background(), logger.NewLogger())
	log := logger.FromContext(ctx).With("version", version)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), manager.Options{
		BaseContext:             func() context.Context { return ctx },
		GracefulShutdownTimeout: &gracePeriod,
	})
	if err != nil {
		log.FatalContext(ctx, "Unable to start manager", "error", err)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		For(&appsv1.StatefulSet{}).
		Complete(controller.NewReconciler(mgr.GetClient()))
	if err != nil {
		log.FatalContext(ctx, "Unable to create controller", "error", err)
	}

	if err := mgr.Start(ctx); err != nil {
		log.FatalContext(ctx, "Manager exited non-zero", "error", err)
	}
}
