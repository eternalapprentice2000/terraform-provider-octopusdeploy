package octopusdeploy

import (
	"context"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandDeploymentProcess(d *schema.ResourceData, client *octopusdeploy.Client) *octopusdeploy.DeploymentProcess {
	projectID := d.Get("project_id").(string)
	deploymentProcess := octopusdeploy.NewDeploymentProcess(projectID)
	deploymentProcess.ID = d.Id()

	if v, ok := d.GetOk("branch"); ok {
		deploymentProcess.Branch = v.(string)
	} else {
		project, err := client.Projects.GetByID(projectID)
		if err != nil {
			return nil
		}

		if project.PersistenceSettings != nil && project.PersistenceSettings.GetType() == "VersionControlled" {
			deploymentProcess.Branch = project.PersistenceSettings.(*octopusdeploy.GitPersistenceSettings).DefaultBranch
		}
	}

	if v, ok := d.GetOk("last_snapshot_id"); ok {
		deploymentProcess.LastSnapshotID = v.(string)
	}

	if v, ok := d.GetOk("space_id"); ok {
		deploymentProcess.SpaceID = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		deploymentProcess.Version = int32(v.(int))
	}

	if v, ok := d.GetOk("step"); ok {
		steps := v.([]interface{})
		for _, step := range steps {
			deploymentStep := expandDeploymentStep(step.(map[string]interface{}))
			deploymentProcess.Steps = append(deploymentProcess.Steps, *deploymentStep)
		}
	}

	return deploymentProcess
}

func setDeploymentProcess(ctx context.Context, d *schema.ResourceData, deploymentProcess *octopusdeploy.DeploymentProcess) error {
	d.Set("branch", deploymentProcess.Branch)
	d.Set("last_snapshot_id", deploymentProcess.LastSnapshotID)
	d.Set("project_id", deploymentProcess.ProjectID)
	d.Set("space_id", deploymentProcess.SpaceID)
	d.Set("version", deploymentProcess.Version)

	if err := d.Set("step", flattenDeploymentSteps(deploymentProcess.Steps)); err != nil {
		return fmt.Errorf("error setting step: %s", err)
	}

	return nil
}
