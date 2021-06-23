package upgraders

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	upgradev1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	"github.com/openshift/managed-upgrade-operator/pkg/scaler"
)

// EnsureExtraUpgradeWorkers will scale up new workers to ensure customer capacity while upgrading.
func (c *clusterUpgrader) EnsureExtraUpgradeWorkers(ctx context.Context, logger logr.Logger) (bool, error) {
	// Skip the step scale up worker node if capacity reservation is set to false
	if !c.upgradeConfig.Spec.CapacityReservation {
		logger.Info("Do not need to scale up extra node(s) since the capacity reservation is disabled")
		return true, nil
	}

	upgradeCommenced, err := c.cvClient.HasUpgradeCommenced(c.upgradeConfig)
	if err != nil {
		return false, err
	}
	desired := c.upgradeConfig.Spec.Desired
	if upgradeCommenced {
		logger.Info(fmt.Sprintf("ClusterVersion is already set to Channel %s Version %s, skipping %s", desired.Channel, desired.Version, upgradev1alpha1.UpgradeScaleUpExtraNodes))
		return true, nil
	}

	isScaled, err := c.scaler.EnsureScaleUpNodes(c.client, c.config.GetScaleDuration(), logger)
	if err != nil {
		if scaler.IsScaleTimeOutError(err) {
			c.metrics.UpdateMetricScalingFailed(c.upgradeConfig.Name)
		}
		return false, err
	}

	if isScaled {
		c.metrics.UpdateMetricScalingSucceeded(c.upgradeConfig.Name)
	}

	return isScaled, nil
}


// RemoveExtraScaledNodes will scale down the extra workers added pre upgrade.
func (c *clusterUpgrader) RemoveExtraScaledNodes(ctx context.Context, logger logr.Logger) (bool, error) {
	// Skip the step scale down worker node if capacity reservation is set to false
	if !c.upgradeConfig.Spec.CapacityReservation {
		logger.Info("Do not need to remove nodes since the capacity reservation is disabled")
		return true, nil
	}

	nds, err := c.drainstrategyBuilder.NewNodeDrainStrategy(c.client, c.upgradeConfig, &c.config.NodeDrain)
	if err != nil {
		return false, err
	}
	isScaledDown, err := c.scaler.EnsureScaleDownNodes(c.client, nds, logger)
	if err != nil {
		dtErr, ok := scaler.IsDrainTimeOutError(err)
		if ok {
			c.metrics.UpdateMetricNodeDrainFailed(dtErr.GetNodeName())
		}
		logger.Error(err, "Extra upgrade node failed to drain in time")
		return false, err
	}

	if isScaledDown {
		c.metrics.ResetAllMetricNodeDrainFailed()
	}

	return isScaledDown, nil
}

