// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDedicatedServers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDedicatedServersResult, error) {
	var rv GetDedicatedServersResult
	err := ctx.Invoke("ovh:index/getDedicatedServers:getDedicatedServers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDedicatedServers.
type GetDedicatedServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	Results []string `pulumi:"results"`
}