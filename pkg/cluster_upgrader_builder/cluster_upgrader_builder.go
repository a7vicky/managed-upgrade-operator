package cluster_upgrader_builder

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/openshift/managed-upgrade-operator/pkg/eventmanager"
	"github.com/openshift/managed-upgrade-operator/pkg/upgraders"
	"sigs.k8s.io/controller-runtime/pkg/client"

	upgradev1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	"github.com/openshift/managed-upgrade-operator/pkg/configmanager"
	"github.com/openshift/managed-upgrade-operator/pkg/metrics"
)

// Interface describing the functions of a cluster upgrader.
//go:generate mockgen -destination=mocks/cluster_upgrader.go -package=mocks github.com/openshift/managed-upgrade-operator/pkg/cluster_upgrader_builder ClusterUpgrader
type ClusterUpgrader interface {
	UpgradeCluster(ctx context.Context, upgradeConfig *upgradev1alpha1.UpgradeConfig, logger logr.Logger) (upgradev1alpha1.UpgradePhase, *upgradev1alpha1.UpgradeCondition, error)
}

//go:generate mockgen -destination=mocks/cluster_upgrader_builder.go -package=mocks github.com/openshift/managed-upgrade-operator/pkg/cluster_upgrader_builder ClusterUpgraderBuilder
type ClusterUpgraderBuilder interface {
	NewClient(client.Client, configmanager.ConfigManager, metrics.Metrics, eventmanager.EventManager, upgradev1alpha1.UpgradeType) (ClusterUpgrader, error)
}

func NewBuilder() ClusterUpgraderBuilder {
	return &clusterUpgraderBuilder{}
}

type clusterUpgraderBuilder struct{}

func (cub *clusterUpgraderBuilder) NewClient(c client.Client, cfm configmanager.ConfigManager, mc metrics.Metrics, nc eventmanager.EventManager, upgradeType upgradev1alpha1.UpgradeType) (ClusterUpgrader, error) {
	switch upgradeType {
	case upgradev1alpha1.OSD:
		cu, err := upgraders.NewOSDUpgrader(c, cfm, mc, nc)
		if err != nil {
			return nil, err
		}
		return cu, nil
	case upgradev1alpha1.ARO:
		cu, err := upgraders.NewAROUpgrader(c, cfm, mc, nc)
		if err != nil {
			return nil, err
		}
		return cu, nil
	default:
		cu, err := upgraders.NewOSDUpgrader(c, cfm, mc, nc)
		if err != nil {
			return nil, err
		}
		return cu, nil
	}
}
