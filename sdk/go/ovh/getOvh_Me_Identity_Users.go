// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Me_Identity_Users(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOvh_Me_Identity_UsersResult, error) {
	var rv GetOvh_Me_Identity_UsersResult
	err := ctx.Invoke("ovh:index/getOvh_Me_Identity_Users:getOvh_Me_Identity_Users", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOvh_Me_Identity_Users.
type GetOvh_Me_Identity_UsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id    string   `pulumi:"id"`
	Users []string `pulumi:"users"`
}
