package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yamamoto-febc/terraform-provider-sakuraiot/builtin/providers/sakuraiot"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sakuraiot.Provider,
	})
}
