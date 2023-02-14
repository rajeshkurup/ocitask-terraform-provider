package ocitaskprovider

import (
	"context"
	"ocitaskclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/**
 * @brief Terraform Provider for OCI Task Service
 */
type OciTaskServProvider struct {
	resource   *OciTaskResource
	dataSource *OciTaskDataSource
}

/**
 * @brief Constructor for OciTaskServProvider
 * @return Instance of OciTaskServProvider
 */
func MakeOciTaskServProvider() *OciTaskServProvider {
	return &OciTaskServProvider{
		resource:   MakeOciTaskResource(),
		dataSource: MakeOciTaskDataSource(),
	}
}

/**
 * @brief Build schema.Provider with provider schema, resource and data source map
 * @return Instance of schema.Provider
 */
func (ociTaskServProvider *OciTaskServProvider) Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ocitask_host": {
				Type:     schema.TypeString,
				Required: true,
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

/**
 * @brief Configure Terraform Provider for OCI Task Service and build Client to OCI Task Service
 * @param ctx Terraform Provider context
 * @param rd Instance of schema.ResourceData contains provider configuration from Terraform scripts
 * @return Generic interface instance equivalent to Client to OCI Task Service if succeeded
 * @return Instance of diag.Diagnostics collection with error details if failed
 */
func (ociTaskServProvider *OciTaskServProvider) providerConfigure(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var ociTaskHost *string

	hVal, ok := rd.GetOk("ocitask_host")
	if ok {
		tempHost := hVal.(string)
		ociTaskHost = &tempHost
	} else {
		return nil, diags
	}

	ociTaskClient := ocitaskclient.MakeOciTaskServClient(ociTaskHost)
	return ociTaskClient, diags
}
