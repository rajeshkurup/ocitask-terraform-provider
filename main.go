package main

import (
	"log"
	"ocitaskprovider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

/**
 * @brief Entry point for OCI Task Service Terraform Provider
 */
func main() {
	log.Println("OCI Task Management Service Terraform Provider Start")

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			prov := ocitaskprovider.MakeOciTaskServProvider()
			return prov.Provider()
		},
	})

	log.Println("All Done")
}
