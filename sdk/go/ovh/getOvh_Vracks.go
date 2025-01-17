// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Vracks(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOvh_VracksResult, error) {
	var rv GetOvh_VracksResult
	err := ctx.Invoke("ovh:index/getOvh_Vracks:getOvh_Vracks", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOvh_Vracks.
type GetOvh_VracksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	Results []string `pulumi:"results"`
}
