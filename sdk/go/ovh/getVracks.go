// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVracks(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetVracksResult, error) {
	var rv GetVracksResult
	err := ctx.Invoke("ovh:index/getVracks:getVracks", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getVracks.
type GetVracksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	Results []string `pulumi:"results"`
}