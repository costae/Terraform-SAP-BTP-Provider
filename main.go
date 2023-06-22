package main

import (
	provider "github.com/costae/Terraform-SAP-BTP-Provider/provider"

	plugin "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
