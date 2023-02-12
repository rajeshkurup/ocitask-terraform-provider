package ocitaskprovider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type OciTaskDataSource struct {
	ociTaskOperation *OciTaskOperation
}

func MakeOciTaskDataSource() *OciTaskDataSource {
	return &OciTaskDataSource{
		ociTaskOperation: &OciTaskOperation{},
	}
}

func (ociTaskDataSource *OciTaskDataSource) DataSourceOciTasks() *schema.Resource {
	return &schema.Resource{
		ReadContext: ociTaskDataSource.ociTaskOperation.OciTaskRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
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
	}
}
