/*
Copyright 2023 Redhat.

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

package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	backendv1alpha1 "github.com/k8s-analytics/anomaly-operator/api/v1alpha1"
)

// AnomalyDataReconciler reconciles a AnomalyData object
type AnomalyDataReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=backend.anomaly.io,resources=anomalydata,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=backend.anomaly.io,resources=anomalydata/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=backend.anomaly.io,resources=anomalydata/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AnomalyData object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *AnomalyDataReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// // Define the path to your YAML manifest file
	// manifestPath := "deploy/anomaly_storage_crd_config.yaml"

	// // Create a Manifestival object with the path to your manifest file
	// manifest, err := mf.NewManifest(manifestPath)
	// if err != nil {
	// 	fmt.Printf("Error creating Manifestival: %v\n", err)
	// 	os.Exit(1)
	// }

	// // Read the manifest file and apply its resources to the cluster
	// manifest.Apply(nil)

	return ctrl.Result{RequeueAfter: time.Duration(30 * time.Second)}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AnomalyDataReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&backendv1alpha1.AnomalyData{}).
		Complete(r)
}
