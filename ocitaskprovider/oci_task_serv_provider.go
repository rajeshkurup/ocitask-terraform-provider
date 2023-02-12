package ocitaskprovider

import (
	"context"
	"ocitaskclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type OciTaskServProvider struct {
	resource   *OciTaskResource
	dataSource *OciTaskDataSource
}

func MakeOciTaskServProvider() *OciTaskServProvider {
	return &OciTaskServProvider{
		resource:   MakeOciTaskResource(),
		dataSource: MakeOciTaskDataSource(),
	}
}

func (ociTaskServProvider *OciTaskServProvider) Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ocitask_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OCITASK_HOST", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ocitask_task": ociTaskServProvider.resource.ResourceOciTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ocitask_tasks": ociTaskServProvider.dataSource.DataSourceOciTasks(),
		},
		ConfigureContextFunc: ociTaskServProvider.providerConfigure,
	}
}

func (ociTaskServProvider *OciTaskServProvider) providerConfigure(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var ociTaskHost *string

	hVal, ok := rd.GetOk("ocitask_host")
	if ok {
		tempHost := hVal.(string)
		ociTaskHost = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ociTaskClient := ocitaskclient.MakeOciTaskServClient(ociTaskHost)
	return ociTaskClient, diags
}
