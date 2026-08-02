package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	frpv1alpha1 "github.com/zufardhiyaulhaq/frp-operator/api/v1alpha1"
	"github.com/zufardhiyaulhaq/frp-operator/pkg/client/status"
)

// TestClientReconciler_StatusWiring_ImageUnresolved exercises the status-writing
// methods the image-unresolved fail-loud branch in Reconcile relies on:
// setCondition + updateClientStatus together must flip a Client to phase=Failed
// and persist a Ready=False / ReasonImageUnresolved condition through a real
// status subresource update.
//
// This covers the new ReasonImageUnresolved wiring without standing up a full
// envtest harness; end-to-end Reconcile coverage remains a follow-up.
func TestClientReconciler_StatusWiring_ImageUnresolved(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := frpv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("add to scheme: %v", err)
	}

	clientObj := &frpv1alpha1.Client{
		ObjectMeta: metav1.ObjectMeta{Name: "test-client", Namespace: "default"},
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(&frpv1alpha1.Client{}).
		WithObjects(clientObj).
		Build()

	r := &ClientReconciler{Client: fakeClient, Scheme: scheme}

	const msg = "no frpc image configured: set spec.podTemplate.image or the operator --frpc-default-image flag"

	// Mirror the fail-loud branch in Reconcile (controllers/client_controller.go).
	r.setCondition(clientObj, status.ConditionTypeReady, metav1.ConditionFalse, status.ReasonImageUnresolved, msg)
	if err := r.updateClientStatus(context.Background(), clientObj, status.ClientPhaseFailed, msg, 0, 0); err != nil {
		t.Fatalf("updateClientStatus: %v", err)
	}

	// Re-fetch from the fake client to confirm the status update persisted.
	got := &frpv1alpha1.Client{}
	if err := fakeClient.Get(context.Background(), types.NamespacedName{Name: "test-client", Namespace: "default"}, got); err != nil {
		t.Fatalf("get client: %v", err)
	}

	if got.Status.Phase != status.ClientPhaseFailed {
		t.Errorf("phase: got %q, want %q", got.Status.Phase, status.ClientPhaseFailed)
	}
	if got.Status.Message != msg {
		t.Errorf("message: got %q, want %q", got.Status.Message, msg)
	}

	var cond *metav1.Condition
	for i := range got.Status.Conditions {
		if got.Status.Conditions[i].Type == status.ConditionTypeReady {
			cond = &got.Status.Conditions[i]
			break
		}
	}
	if cond == nil {
		t.Fatalf("no Ready condition set on status; conditions=%#v", got.Status.Conditions)
	}
	if cond.Status != metav1.ConditionFalse {
		t.Errorf("Ready condition status: got %s, want False", cond.Status)
	}
	if cond.Reason != status.ReasonImageUnresolved {
		t.Errorf("Ready condition reason: got %q, want %q", cond.Reason, status.ReasonImageUnresolved)
	}
}
