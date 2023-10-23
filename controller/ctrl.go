package main

import (
	"context"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	kubeconv1 "github.com/acrlabs/kubecon-na2023/controller/api/v1"
)

var (
	conjobScheme = runtime.NewScheme()
	setupLog     = ctrl.Log.WithName("setup")
)

// ConJobReconciler reconciles a ConJob object
type ConJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=kubecon.appliedcomputing.io,resources=conjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kubecon.appliedcomputing.io,resources=conjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kubecon.appliedcomputing.io,resources=conjobs/finalizers,verbs=update

func (r *ConJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("got new request %v", "request", req)

	return ctrl.Result{}, nil
}

func (r *ConJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubeconv1.ConJob{}).
		Complete(r)
}

func main() {
	opts := zap.Options{
		Development: true,
	}

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: conjobScheme,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&ConJobReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConJob")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func init() {
	utilruntime.Must(scheme.AddToScheme(conjobScheme))
	utilruntime.Must(kubeconv1.AddToScheme(conjobScheme))
}
