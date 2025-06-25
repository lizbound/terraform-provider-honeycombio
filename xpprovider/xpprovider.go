package xpprovider

import (
	"context"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/honeycombio/terraform-provider-honeycombio/honeycombio"
	"github.com/honeycombio/terraform-provider-honeycombio/internal/provider"
)

func GetProvider(_ context.Context) (fwprovider.Provider, *schema.Provider) {
	fwProvider := provider.New("0.36.0")
	sdkProvider := honeycombio.Provider("0.36.0")
	return fwProvider, sdkProvider
}
