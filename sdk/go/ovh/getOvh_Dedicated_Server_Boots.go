// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Dedicated_Server_Boots(ctx *pulumi.Context, args *GetOvh_Dedicated_Server_BootsArgs, opts ...pulumi.InvokeOption) (*GetOvh_Dedicated_Server_BootsResult, error) {
	var rv GetOvh_Dedicated_Server_BootsResult
	err := ctx.Invoke("ovh:index/getOvh_Dedicated_Server_Boots:getOvh_Dedicated_Server_Boots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Dedicated_Server_Boots.
type GetOvh_Dedicated_Server_BootsArgs struct {
	BootType    *string `pulumi:"bootType"`
	Kernel      *string `pulumi:"kernel"`
	ServiceName string  `pulumi:"serviceName"`
}

// A collection of values returned by getOvh_Dedicated_Server_Boots.
type GetOvh_Dedicated_Server_BootsResult struct {
	BootType *string `pulumi:"bootType"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Kernel      *string `pulumi:"kernel"`
	Results     []int   `pulumi:"results"`
	ServiceName string  `pulumi:"serviceName"`
}

func GetOvh_Dedicated_Server_BootsOutput(ctx *pulumi.Context, args GetOvh_Dedicated_Server_BootsOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Dedicated_Server_BootsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Dedicated_Server_BootsResult, error) {
			args := v.(GetOvh_Dedicated_Server_BootsArgs)
			r, err := GetOvh_Dedicated_Server_Boots(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Dedicated_Server_BootsResultOutput)
}

// A collection of arguments for invoking getOvh_Dedicated_Server_Boots.
type GetOvh_Dedicated_Server_BootsOutputArgs struct {
	BootType    pulumi.StringPtrInput `pulumi:"bootType"`
	Kernel      pulumi.StringPtrInput `pulumi:"kernel"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
}

func (GetOvh_Dedicated_Server_BootsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Dedicated_Server_BootsArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Dedicated_Server_Boots.
type GetOvh_Dedicated_Server_BootsResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Dedicated_Server_BootsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Dedicated_Server_BootsResult)(nil)).Elem()
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) ToGetOvh_Dedicated_Server_BootsResultOutput() GetOvh_Dedicated_Server_BootsResultOutput {
	return o
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) ToGetOvh_Dedicated_Server_BootsResultOutputWithContext(ctx context.Context) GetOvh_Dedicated_Server_BootsResultOutput {
	return o
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) BootType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_Server_BootsResult) *string { return v.BootType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Dedicated_Server_BootsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_Server_BootsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) Kernel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_Server_BootsResult) *string { return v.Kernel }).(pulumi.StringPtrOutput)
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) Results() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_Server_BootsResult) []int { return v.Results }).(pulumi.IntArrayOutput)
}

func (o GetOvh_Dedicated_Server_BootsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_Server_BootsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Dedicated_Server_BootsResultOutput{})
}