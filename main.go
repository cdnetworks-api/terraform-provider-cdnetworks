package main

import (
	"flag"
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./example/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
// Before executing this command, please set debuggable to false
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name cdnetworks

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cdnetworks.Provider,
		ProviderAddr: "registry.terraform.io/cdnetworks-api/cdnetworks",
		Debug:        debugMode,
	})
}
