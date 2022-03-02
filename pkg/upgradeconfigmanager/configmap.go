package upgradeconfigmanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func EnsureConfigmapExist(namespaceName, configmapName string) bool {
	cfg, _ := config.GetConfig()
	c, _ := client.New(cfg, client.Options{})
	cm := &corev1.ConfigMap{}

	_ = c.Get(context.Background(), client.ObjectKey{
		Namespace: namespaceName,
		Name:      configmapName,
	}, cm)

	if cm.GetName() == configmapName {
		return true
	} else {
		return false
	}
}
