package octopusdeploy

import (
	"context"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceScriptModules() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about existing script modules.",
		ReadContext: dataSourceScriptModulesRead,
		Schema:      getScriptModuleDataSchema(),
	}
}

func dataSourceScriptModulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	query := octopusdeploy.LibraryVariablesQuery{
		ContentType: "ScriptModule",
		IDs:         expandArray(d.Get("ids").([]interface{})),
		PartialName: d.Get("partial_name").(string),
		Skip:        d.Get("skip").(int),
		Take:        d.Get("take").(int),
	}

	client := m.(*octopusdeploy.Client)
	libraryVariableSets, err := client.ScriptModules.Get(query)
	if err != nil {
		return diag.FromErr(err)
	}

	flattenedScriptModules := []interface{}{}
	for _, libraryVariableSet := range libraryVariableSets.Items {
		flattenedScriptModules = append(flattenedScriptModules, flattenScriptModule(libraryVariableSet))
	}

	d.Set("script_modules", flattenedScriptModules)
	d.SetId("Script Modules " + time.Now().UTC().String())

	return nil
}
