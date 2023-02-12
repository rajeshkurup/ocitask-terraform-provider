package ocitaskprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type OciTaskResource struct {
	ociTaskOperation *OciTaskOperation
}

func MakeOciTaskResource() *OciTaskResource {
	return &OciTaskResource{
		ociTaskOperation: &OciTaskOperation{},
	}
}

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
