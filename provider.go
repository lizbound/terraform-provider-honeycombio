package main

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/honeycombio/terraform-provider-honeycombio/internal/provider"
)

func Version() string {
	return providerVersion
}

func Provider() fwprovider.Provider {
	return provider.New(Version())
}
