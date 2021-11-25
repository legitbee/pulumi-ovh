// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIploadbalancingVrackNetworks(ctx *pulumi.Context, args *GetIploadbalancingVrackNetworksArgs, opts ...pulumi.InvokeOption) (*GetIploadbalancingVrackNetworksResult, error) {
	var rv GetIploadbalancingVrackNetworksResult
	err := ctx.Invoke("ovh:index/getIploadbalancingVrackNetworks:getIploadbalancingVrackNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIploadbalancingVrackNetworks.
type GetIploadbalancingVrackNetworksArgs struct {
	ServiceName string  `pulumi:"serviceName"`
	Subnet      *string `pulumi:"subnet"`
	VlanId      *int    `pulumi:"vlanId"`
}

// A collection of values returned by getIploadbalancingVrackNetworks.
type GetIploadbalancingVrackNetworksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Results     []int   `pulumi:"results"`
	ServiceName string  `pulumi:"serviceName"`
	Subnet      *string `pulumi:"subnet"`
	VlanId      *int    `pulumi:"vlanId"`
}

func GetIploadbalancingVrackNetworksOutput(ctx *pulumi.Context, args GetIploadbalancingVrackNetworksOutputArgs, opts ...pulumi.InvokeOption) GetIploadbalancingVrackNetworksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIploadbalancingVrackNetworksResult, error) {
			args := v.(GetIploadbalancingVrackNetworksArgs)
			r, err := GetIploadbalancingVrackNetworks(ctx, &args, opts...)
			return *r, err
		}).(GetIploadbalancingVrackNetworksResultOutput)
}

// A collection of arguments for invoking getIploadbalancingVrackNetworks.
type GetIploadbalancingVrackNetworksOutputArgs struct {
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	Subnet      pulumi.StringPtrInput `pulumi:"subnet"`
	VlanId      pulumi.IntPtrInput    `pulumi:"vlanId"`
}

func (GetIploadbalancingVrackNetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIploadbalancingVrackNetworksArgs)(nil)).Elem()
}

// A collection of values returned by getIploadbalancingVrackNetworks.
type GetIploadbalancingVrackNetworksResultOutput struct{ *pulumi.OutputState }

func (GetIploadbalancingVrackNetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIploadbalancingVrackNetworksResult)(nil)).Elem()
}

func (o GetIploadbalancingVrackNetworksResultOutput) ToGetIploadbalancingVrackNetworksResultOutput() GetIploadbalancingVrackNetworksResultOutput {
	return o
}

func (o GetIploadbalancingVrackNetworksResultOutput) ToGetIploadbalancingVrackNetworksResultOutputWithContext(ctx context.Context) GetIploadbalancingVrackNetworksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIploadbalancingVrackNetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIploadbalancingVrackNetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIploadbalancingVrackNetworksResultOutput) Results() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetIploadbalancingVrackNetworksResult) []int { return v.Results }).(pulumi.IntArrayOutput)
}

func (o GetIploadbalancingVrackNetworksResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIploadbalancingVrackNetworksResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetIploadbalancingVrackNetworksResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIploadbalancingVrackNetworksResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o GetIploadbalancingVrackNetworksResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetIploadbalancingVrackNetworksResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIploadbalancingVrackNetworksResultOutput{})
}
