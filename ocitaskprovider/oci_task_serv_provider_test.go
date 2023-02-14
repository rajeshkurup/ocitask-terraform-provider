package ocitaskprovider

import (
	"ocitaskclient"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestProvider(test *testing.T) {

	ociTaskServProvider := MakeOciTaskServProvider()

	assert.NotNil(test, ociTaskServProvider, "TestProvider Failed: Valid Provider expected")
	assert.NotNil(test, ociTaskServProvider.dataSource, "TestProvider Failed: Valid Data Source expected")
	assert.NotNil(test, ociTaskServProvider.resource, "TestProvider Failed: Valid Resource expected")

	provider := ociTaskServProvider.Provider()

	assert.NotNil(test, provider, "TestProvider Failed: Valid native Provider expected")

	hostSchema := provider.Schema["ocitask_host"]

	assert.Equal(test, schema.TypeString, hostSchema.Type, "TestProvider Failed: Host Schema Type doesn't match with expected value")
	assert.Equal(test, true, hostSchema.Required, "TestProvider Failed: Host Schema Required flag doesn't match with expected value")
	assert.Equal(test, true, hostSchema.Required, "TestProvider Failed: Host Schema Required flag doesn't match with expected value")

	resource := provider.ResourcesMap["ocitask_task"]

	assert.NotNil(test, resource, "TestProvider Failed: Resource expected")
	assert.NotNil(test, resource.CreateContext, "TestProvider Failed: CreateContext expected")
	assert.NotNil(test, resource.UpdateContext, "TestProvider Failed: UpdateContext expected")
	assert.NotNil(test, resource.ReadContext, "TestProvider Failed: ReadContext expected")
	assert.NotNil(test, resource.DeleteContext, "TestProvider Failed: DeleteContext expected")

	resourceSchema := resource.Schema

	assert.Equal(test, schema.TypeString, resourceSchema["last_updated"].Type, "TestProvider Failed: Resource Schema last_updated Type doesn't match with expected value")
	assert.Equal(test, true, resourceSchema["last_updated"].Optional, "TestProvider Failed: Resource Schema last_updated Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceSchema["last_updated"].Computed, "TestProvider Failed: Resource Schema last_updated Computed flag doesn't match with expected value")
	assert.Equal(test, schema.TypeList, resourceSchema["items"].Type, "TestProvider Failed: Resource Schema items Type doesn't match with expected value")
	assert.Equal(test, true, resourceSchema["items"].Required, "TestProvider Failed: Resource Schema items Required flag doesn't match with expected value")
	assert.Equal(test, 1, resourceSchema["items"].MaxItems, "TestProvider Failed: Resource Schema items MaxItems doesn't match with expected value")

	resourceItem := resourceSchema["items"].Elem.(*schema.Resource)

	assert.Equal(test, schema.TypeInt, resourceItem.Schema["id"].Type, "TestProvider Failed: Resource Item Id Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["id"].Optional, "TestProvider Failed: Resource Item Id Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["id"].Computed, "TestProvider Failed: Resource Item Id Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["title"].Type, "TestProvider Failed: Resource Item Title Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["title"].Required, "TestProvider Failed: Resource Item Title Required flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["description"].Type, "TestProvider Failed: Resource Item Description Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["description"].Optional, "TestProvider Failed: Resource Item Description Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["description"].Computed, "TestProvider Failed: Resource Item Description Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeInt, resourceItem.Schema["priority"].Type, "TestProvider Failed: Resource Item Priority Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["priority"].Optional, "TestProvider Failed: Resource Item Priority Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["priority"].Computed, "TestProvider Failed: Resource Item Priority Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeBool, resourceItem.Schema["completed"].Type, "TestProvider Failed: Resource Item Completed Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["completed"].Optional, "TestProvider Failed: Resource Item Completed Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["completed"].Computed, "TestProvider Failed: Resource Item Completed Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["start_date"].Type, "TestProvider Failed: Resource Item start_date Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["start_date"].Optional, "TestProvider Failed: Resource Item start_date Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["start_date"].Computed, "TestProvider Failed: Resource Item start_date Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["due_date"].Type, "TestProvider Failed: Resource Item due_date Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["due_date"].Optional, "TestProvider Failed: Resource Item due_date Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["due_date"].Computed, "TestProvider Failed: Resource Item due_date Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["time_created"].Type, "TestProvider Failed: Resource Item time_created Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["time_created"].Optional, "TestProvider Failed: Resource Item time_created Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["time_created"].Computed, "TestProvider Failed: Resource Item time_created Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, resourceItem.Schema["time_updated"].Type, "TestProvider Failed: Resource Item time_updated Type doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["time_updated"].Optional, "TestProvider Failed: Resource Item time_updated Optional flag doesn't match with expected value")
	assert.Equal(test, true, resourceItem.Schema["time_updated"].Computed, "TestProvider Failed: Resource Item time_updated Computed flag doesn't match with expected value")

	dataSource := provider.DataSourcesMap["ocitask_tasks"]

	assert.NotNil(test, dataSource, "TestProvider Failed: DataSource expected")
	assert.Nil(test, dataSource.CreateContext, "TestProvider Failed: CreateContext not expected")
	assert.Nil(test, dataSource.UpdateContext, "TestProvider Failed: UpdateContext not expected")
	assert.NotNil(test, dataSource.ReadContext, "TestProvider Failed: ReadContext expected")
	assert.Nil(test, dataSource.DeleteContext, "TestProvider Failed: DeleteContext not expected")

	dataSourceSchema := dataSource.Schema

	assert.Equal(test, schema.TypeInt, dataSourceSchema["id"].Type, "TestProvider Failed: DataSource Schema Id Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceSchema["id"].Required, "TestProvider Failed: DataSource Schema id Required flag doesn't match with expected value")
	assert.Equal(test, schema.TypeList, dataSourceSchema["items"].Type, "TestProvider Failed: DataSource Schema items Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceSchema["items"].Computed, "TestProvider Failed: DataSource Schema items Computed flag doesn't match with expected value")

	dataSourceItem := dataSourceSchema["items"].Elem.(*schema.Resource)

	assert.Equal(test, schema.TypeInt, dataSourceItem.Schema["id"].Type, "TestProvider Failed: DataSource Item Id Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["id"].Optional, "TestProvider Failed: DataSource Item Id Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["id"].Computed, "TestProvider Failed: DataSource Item Id Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["title"].Type, "TestProvider Failed: DataSource Item Title Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["title"].Required, "TestProvider Failed: DataSource Item Title Required flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["description"].Type, "TestProvider Failed: DataSource Item Description Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["description"].Optional, "TestProvider Failed: DataSource Item Description Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["description"].Computed, "TestProvider Failed: DataSource Item Description Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeInt, dataSourceItem.Schema["priority"].Type, "TestProvider Failed: DataSource Item Priority Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["priority"].Optional, "TestProvider Failed: DataSource Item Priority Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["priority"].Computed, "TestProvider Failed: DataSource Item Priority Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeBool, dataSourceItem.Schema["completed"].Type, "TestProvider Failed: DataSource Item Completed Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["completed"].Optional, "TestProvider Failed: DataSource Item Completed Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["completed"].Computed, "TestProvider Failed: DataSource Item Completed Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["start_date"].Type, "TestProvider Failed: DataSource Item start_date Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["start_date"].Optional, "TestProvider Failed: DataSource Item start_date Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["start_date"].Computed, "TestProvider Failed: DataSource Item start_date Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["due_date"].Type, "TestProvider Failed: DataSource Item due_date Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["due_date"].Optional, "TestProvider Failed: DataSource Item due_date Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["due_date"].Computed, "TestProvider Failed: DataSource Item due_date Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["time_created"].Type, "TestProvider Failed: DataSource Item time_created Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["time_created"].Optional, "TestProvider Failed: DataSource Item time_created Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["time_created"].Computed, "TestProvider Failed: DataSource Item time_created Computed flag doesn't match with expected value")

	assert.Equal(test, schema.TypeString, dataSourceItem.Schema["time_updated"].Type, "TestProvider Failed: DataSource Item time_updated Type doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["time_updated"].Optional, "TestProvider Failed: DataSource Item time_updated Optional flag doesn't match with expected value")
	assert.Equal(test, true, dataSourceItem.Schema["time_updated"].Computed, "TestProvider Failed: DataSource Item time_updated Computed flag doesn't match with expected value")

	assert.NotNil(test, provider.ConfigureContextFunc, "TestProvider Failed: Provider ConfigureContextFunc expected")

	var testSchema = map[string]*schema.Schema{
		"ocitask_host": {Type: schema.TypeString},
	}

	hostUrl := make(map[string]interface{})
	hostUrl["ocitask_host"] = "http://localhost"

	rd := schema.TestResourceDataRaw(test, testSchema, hostUrl)

	iOciTaskClient, _ := provider.ConfigureContextFunc(nil, rd)

	assert.NotNil(test, iOciTaskClient, "TestProvider Failed: Generic OciTaskClient expected from ConfigureContextFunc")

	ociTaskClient := iOciTaskClient.(*ocitaskclient.OciTaskServClient)

	assert.NotNil(test, ociTaskClient, "TestProvider Failed: OciTaskClient expected from ConfigureContextFunc")

	assert.Equal(test, "http://localhost", ociTaskClient.GetUrl(), "TestProvider Failed: OciTaskClient Host URL doesn't match with expected value")
}

func TestProviderConfigureSuccess(test *testing.T) {
	ociTaskServProvider := MakeOciTaskServProvider()

	assert.NotNil(test, ociTaskServProvider, "TestProviderConfigureSuccess Failed: Valid Provider expected")
	assert.NotNil(test, ociTaskServProvider.dataSource, "TestProviderConfigureSuccess Failed: Valid Data Source expected")
	assert.NotNil(test, ociTaskServProvider.resource, "TestProviderConfigureSuccess Failed: Valid Resource expected")

	provider := ociTaskServProvider.Provider()

	var testSchema = map[string]*schema.Schema{
		"ocitask_host": {Type: schema.TypeString},
	}

	hostUrl := make(map[string]interface{})
	hostUrl["ocitask_host"] = "http://localhost"

	rd := schema.TestResourceDataRaw(test, testSchema, hostUrl)

	iOciTaskClient, _ := provider.ConfigureContextFunc(nil, rd)

	assert.NotNil(test, iOciTaskClient, "TestProviderConfigureSuccess Failed: Generic OciTaskClient expected from ConfigureContextFunc")

	ociTaskClient := iOciTaskClient.(*ocitaskclient.OciTaskServClient)

	assert.NotNil(test, ociTaskClient, "TestProviderConfigureSuccess Failed: OciTaskClient expected from ConfigureContextFunc")

	assert.Equal(test, "http://localhost", ociTaskClient.GetUrl(), "TestProviderConfigureSuccess Failed: OciTaskClient Host URL doesn't match with expected value")
}

func TestProviderConfigureFailed(test *testing.T) {
	ociTaskServProvider := MakeOciTaskServProvider()

	assert.NotNil(test, ociTaskServProvider, "TestProviderConfigureFailed Failed: Valid Provider expected")
	assert.NotNil(test, ociTaskServProvider.dataSource, "TestProviderConfigureFailed Failed: Valid Data Source expected")
	assert.NotNil(test, ociTaskServProvider.resource, "TestProviderConfigureFailed Failed: Valid Resource expected")

	provider := ociTaskServProvider.Provider()

	var testSchema = map[string]*schema.Schema{
		"ocitask_host": {Type: schema.TypeString},
	}

	hostUrl := make(map[string]interface{})

	rd := schema.TestResourceDataRaw(test, testSchema, hostUrl)

	iOciTaskClient, _ := provider.ConfigureContextFunc(nil, rd)

	assert.Nil(test, iOciTaskClient, "TestProviderConfigureFailed Failed: Generic OciTaskClient not expected from ConfigureContextFunc")
}
