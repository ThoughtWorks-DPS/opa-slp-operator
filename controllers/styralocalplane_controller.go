/*
Copyright 2022.

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
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	opav1alpha1 "github.com/ThoughtWorks-DPS/opa-slp-operator/api/v1alpha1"
	"opa-slp-operator/controllers/styra/system"
)

// StyraLocalPlaneReconciler reconciles a StyraLocalPlane object
type StyraLocalPlaneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=opa.twdps,resources=styralocalplanes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=opa.twdps,resources=styralocalplanes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=opa.twdps,resources=styralocalplanes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StyraLocalPlane object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *StyraLocalPlaneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrllog.FromContext(ctx)

	// Fetch the OpaSidecar instance
	styraLocalPlane := &opav1alpha1.StyraLocalPlane{}
	if err := r.Get(ctx, req.NamespacedName, styraLocalPlane); err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("StyraLocalPlane resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get StyraLocalPlane")
		return ctrl.Result{}, err
	}

	// Log reconcilation event
	log.Info(fmt.Sprintf("Reconciling StyraLocalPlane %s/%s", req.NamespacedName.Namespace, req.NamespacedName.Name))
	
	// Get Styra System opa token

	// Set status conditions
	//styraLocalPlane.Status.Ready =
	//styraLocalPlane.Status.Status =

	// req.NamespacedName.Namespace

	err := r.Status().Update(ctx, styraLocalPlane)
	if err != nil {
		log.Error(err, "Failed to update styraLocalPlane status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StyraLocalPlaneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&opav1alpha1.StyraLocalPlane{}).
		Complete(r)
}
