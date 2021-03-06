package controllers

import (
	nsmv1alpha1 "github.com/networkservicemesh/nsm-operator/apis/nsm/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *NSMReconciler) serviceForWebhook(nsm *nsmv1alpha1.NSM) *corev1.Service {

	annotations := map[string]string{}

	if r.isPlatformOpenShift() {
		annotations = map[string]string{"service.beta.openshift.io/serving-cert-secret-name": "nsm-admission-webhook-certs"}
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webhookServiceName,
			Namespace: nsmNamespace,
			Labels:    labelsForNSMAdmissionWebhook(nsm.Name),

			// TODO: Solve TLS Certs for OCP - This annotation below is specific to OpenShift and needs to be addressed other way
			Annotations: annotations,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Name: webhookName + "-port", Port: webhookServicePort, TargetPort: intstr.FromInt(webhookServiceTargetPort)},
			},
			Selector: map[string]string{"app": webhookName},
		},
	}
	// Set NSM instance as the owner and controller
	controllerutil.SetControllerReference(nsm, service, r.Scheme)
	return service
}
