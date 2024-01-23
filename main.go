package main

import (
	"github.com/deepshore/terraform-provider-netapp-cloudmanager/cloudmanager"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudmanager.Provider,
	})
}
