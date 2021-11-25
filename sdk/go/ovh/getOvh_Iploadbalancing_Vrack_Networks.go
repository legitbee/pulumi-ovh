// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Iploadbalancing_Vrack_Networks(ctx *pulumi.Context, args *GetOvh_Iploadbalancing_Vrack_NetworksArgs, opts ...pulumi.InvokeOption) (*GetOvh_Iploadbalancing_Vrack_NetworksResult, error) {
	var rv GetOvh_Iploadbalancing_Vrack_NetworksResult
	err := ctx.Invoke("ovh:index/getOvh_Iploadbalancing_Vrack_Networks:getOvh_Iploadbalancing_Vrack_Networks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Iploadbalancing_Vrack_Networks.
type GetOvh_Iploadbalancing_Vrack_NetworksArgs struct {
	ServiceName string  `pulumi:"serviceName"`
	Subnet      *string `pulumi:"subnet"`
	VlanId      *int    `pulumi:"vlanId"`
}

// A collection of values returned by getOvh_Iploadbalancing_Vrack_Networks.
type GetOvh_Iploadbalancing_Vrack_NetworksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Results     []int   `pulumi:"results"`
	ServiceName string  `pulumi:"serviceName"`
	Subnet      *string `pulumi:"subnet"`
	VlanId      *int    `pulumi:"vlanId"`
}

func GetOvh_Iploadbalancing_Vrack_NetworksOutput(ctx *pulumi.Context, args GetOvh_Iploadbalancing_Vrack_NetworksOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Iploadbalancing_Vrack_NetworksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Iploadbalancing_Vrack_NetworksResult, error) {
			args := v.(GetOvh_Iploadbalancing_Vrack_NetworksArgs)
			r, err := GetOvh_Iploadbalancing_Vrack_Networks(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Iploadbalancing_Vrack_NetworksResultOutput)
}

// A collection of arguments for invoking getOvh_Iploadbalancing_Vrack_Networks.
type GetOvh_Iploadbalancing_Vrack_NetworksOutputArgs struct {
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	Subnet      pulumi.StringPtrInput `pulumi:"subnet"`
	VlanId      pulumi.IntPtrInput    `pulumi:"vlanId"`
}

func (GetOvh_Iploadbalancing_Vrack_NetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Iploadbalancing_Vrack_NetworksArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Iploadbalancing_Vrack_Networks.
type GetOvh_Iploadbalancing_Vrack_NetworksResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Iploadbalancing_Vrack_NetworksResult)(nil)).Elem()
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) ToGetOvh_Iploadbalancing_Vrack_NetworksResultOutput() GetOvh_Iploadbalancing_Vrack_NetworksResultOutput {
	return o
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) ToGetOvh_Iploadbalancing_Vrack_NetworksResultOutputWithContext(ctx context.Context) GetOvh_Iploadbalancing_Vrack_NetworksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Iploadbalancing_Vrack_NetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) Results() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOvh_Iploadbalancing_Vrack_NetworksResult) []int { return v.Results }).(pulumi.IntArrayOutput)
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Iploadbalancing_Vrack_NetworksResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Iploadbalancing_Vrack_NetworksResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o GetOvh_Iploadbalancing_Vrack_NetworksResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetOvh_Iploadbalancing_Vrack_NetworksResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Iploadbalancing_Vrack_NetworksResultOutput{})
}