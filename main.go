package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jetstack/terraform-provider-awstag/awstag"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: awstag.Provider})
}
