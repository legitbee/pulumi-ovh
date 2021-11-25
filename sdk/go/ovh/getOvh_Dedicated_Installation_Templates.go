// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Dedicated_Installation_Templates(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOvh_Dedicated_Installation_TemplatesResult, error) {
	var rv GetOvh_Dedicated_Installation_TemplatesResult
	err := ctx.Invoke("ovh:index/getOvh_Dedicated_Installation_Templates:getOvh_Dedicated_Installation_Templates", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOvh_Dedicated_Installation_Templates.
type GetOvh_Dedicated_Installation_TemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string   `pulumi:"id"`
	Results []string `pulumi:"results"`
}