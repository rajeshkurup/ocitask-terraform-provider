package ocitaskprovider

import (
	"context"
	"ocitaskclient"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/**
 * @brief Define OCI Task Service operations that can be performed by Terraform provider
 */
type OciTaskOperation struct {
	// Empty
}

/**
 * @brief Constructor for OciTaskOperation
 * @return Instance of OciTaskOperation
 */
func MakeOciTaskOperation() *OciTaskOperation {
	return &OciTaskOperation{}
}

/**
 * @brief Create new Task in OCI Task System
 * @param ctx Context to Terraform Provider
 * @param rd Contains Task instance defined in Terraform scripts
 * @param m Contains OCI Task Service Client pluged into Terraform Provider
 * @return Collection of diag.Diagnostics instances if failed, otherwise empty
 */
func (ociTaskOperation *OciTaskOperation) OciTaskCreate(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	items := rd.Get("items").([]interface{})
	if len(items) > 0 {
		ociRequest, err := ocitaskclient.MakeOciTaskServRequest(&items[0])
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to make OCI Task Service Request",
				Detail:   err.Error(),
			})
		} else {
			ociClient := m.(ocitaskclient.OciTaskServClientInterface)
			ociResponse, err := ociClient.CreateTask(ociRequest)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to create task",
					Detail:   err.Error(),
				})
			} else {
				if ociResponse.Err != nil {
					ociErr, _ := ociResponse.Err.Serialize()
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to create task",
						Detail:   ociErr,
					})
				} else {
					rd.SetId(strconv.FormatInt(*ociResponse.TaskId, 10))
					ociTaskOperation.OciTaskRead(ctx, rd, m)
				}
			}
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid Argument - No tasks found",
			Detail:   "No Items found in incoming Resource data",
		})
	}

	return diags
}

/**
 * @brief Update existing Task in OCI Task System
 * @param ctx Context to Terraform Provider
 * @param rd Contains Task instance defined in Terraform scripts
 * @param m Contains OCI Task Service Client pluged into Terraform Provider
 * @return Collection of diag.Diagnostics instances if failed, otherwise empty
 */
func (ociTaskOperation *OciTaskOperation) OciTaskUpdate(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	taskId, err := strconv.ParseInt(rd.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get Id from resource data",
			Detail:   err.Error(),
		})
	} else {
		items := rd.Get("items").([]interface{})
		if len(items) > 0 {
			ociRequest, err := ocitaskclient.MakeOciTaskServRequest(&items[0])
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to make OCI Task Service Request",
					Detail:   err.Error(),
				})
			} else {
				ociClient := m.(ocitaskclient.OciTaskServClientInterface)
				ociResponse, err := ociClient.UpdateTask(&taskId, ociRequest)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to update task",
						Detail:   err.Error(),
					})
				} else {
					if ociResponse.Err != nil {
						ociErr, _ := ociResponse.Err.Serialize()
						diags = append(diags, diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Failed to update task",
							Detail:   ociErr,
						})
					} else {
						rd.SetId(strconv.FormatInt(*ociResponse.TaskId, 10))
						ociTaskOperation.OciTaskRead(ctx, rd, m)
					}
				}
			}
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Invalid Argument - No tasks found",
				Detail:   "No Items found in incoming Resource data",
			})
		}
	}

	return diags
}

/**
 * @brief Read Task in OCI Task System
 * @param ctx Context to Terraform Provider
 * @param rd Contains Task Identifier defined in Terraform scripts
 * @param m Contains OCI Task Service Client pluged into Terraform Provider
 * @return Collection of diag.Diagnostics instances if failed, otherwise empty
 */
func (ociTaskOperation *OciTaskOperation) OciTaskRead(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	taskId, err := strconv.ParseInt(rd.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get Id from resource data",
			Detail:   err.Error(),
		})
	} else {
		ociClient := m.(ocitaskclient.OciTaskServClientInterface)
		ociResponse, err := ociClient.GetTask(&taskId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to read task",
				Detail:   err.Error(),
			})
		} else {
			if ociResponse.Err != nil {
				ociErr, _ := ociResponse.Err.Serialize()
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to read task",
					Detail:   ociErr,
				})
			} else {
				ociTasks, flatDiag := ocitaskclient.FlattenOciTask(ociResponse.Task)
				if len(flatDiag) == 0 {
					err := rd.Set("items", ociTasks)
					if err != nil {
						diags = append(diags, diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Failed to set task into resource data",
							Detail:   err.Error(),
						})
					}
				} else {
					diags = append(diags, flatDiag...)
				}
			}
		}
	}

	return diags
}

/**
 * @brief Delete Task in OCI Task System
 * @param ctx Context to Terraform Provider
 * @param rd Contains Task Identifier defined in Terraform scripts
 * @param m Contains OCI Task Service Client pluged into Terraform Provider
 * @return Collection of diag.Diagnostics instances if failed, otherwise empty
 */
func (ociTaskOperation *OciTaskOperation) OciTaskDelete(ctx context.Context, rd *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	taskId, err := strconv.ParseInt(rd.Id(), 10, 64)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get Id from resource data",
			Detail:   err.Error(),
		})
	} else {
		ociClient := m.(ocitaskclient.OciTaskServClientInterface)
		ociResponse, err := ociClient.DeleteTask(&taskId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to delete task",
				Detail:   err.Error(),
			})
		} else {
			if ociResponse.Err != nil {
				ociErr, _ := ociResponse.Err.Serialize()
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to delete task",
					Detail:   ociErr,
				})
			} else {
				rd.SetId("")
			}
		}
	}

	return diags
}
