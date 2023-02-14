package ocitaskprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/**
 * @brief Define schema for Task resource in OCI Task System
 */
type OciTaskResource struct {
	ociTaskOperation *OciTaskOperation
}

/**
 * @brief Constructor for OciTaskResource
 * @return Instance of OciTaskResource
 */
func MakeOciTaskResource() *OciTaskResource {
	return &OciTaskResource{
		ociTaskOperation: MakeOciTaskOperation(),
	}
}

/**
 * @brief Build schema for Task resource in OCI Task System
 * @return Instance of schema.Resource contains schema for Task resource in OCI Task System
 */
func (ociTaskResource *OciTaskResource) ResourceOciTask() *schema.Resource {
	return &schema.Resource{
		CreateContext: ociTaskResource.ociTaskOperation.OciTaskCreate,
		ReadContext:   ociTaskResource.ociTaskOperation.OciTaskRead,
		UpdateContext: ociTaskResource.ociTaskOperation.OciTaskUpdate,
		DeleteContext: ociTaskResource.ociTaskOperation.OciTaskDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"items": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"completed": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"start_date": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"due_date": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}
