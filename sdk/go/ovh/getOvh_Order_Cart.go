// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Order_Cart(ctx *pulumi.Context, args *GetOvh_Order_CartArgs, opts ...pulumi.InvokeOption) (*GetOvh_Order_CartResult, error) {
	var rv GetOvh_Order_CartResult
	err := ctx.Invoke("ovh:index/getOvh_Order_Cart:getOvh_Order_Cart", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Order_Cart.
type GetOvh_Order_CartArgs struct {
	Description   *string `pulumi:"description"`
	Expire        *string `pulumi:"expire"`
	OvhSubsidiary string  `pulumi:"ovhSubsidiary"`
}

// A collection of values returned by getOvh_Order_Cart.
type GetOvh_Order_CartResult struct {
	CartId      string  `pulumi:"cartId"`
	Description *string `pulumi:"description"`
	Expire      string  `pulumi:"expire"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	Items         []int  `pulumi:"items"`
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	ReadOnly      bool   `pulumi:"readOnly"`
}

func GetOvh_Order_CartOutput(ctx *pulumi.Context, args GetOvh_Order_CartOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Order_CartResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Order_CartResult, error) {
			args := v.(GetOvh_Order_CartArgs)
			r, err := GetOvh_Order_Cart(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Order_CartResultOutput)
}

// A collection of arguments for invoking getOvh_Order_Cart.
type GetOvh_Order_CartOutputArgs struct {
	Description   pulumi.StringPtrInput `pulumi:"description"`
	Expire        pulumi.StringPtrInput `pulumi:"expire"`
	OvhSubsidiary pulumi.StringInput    `pulumi:"ovhSubsidiary"`
}

func (GetOvh_Order_CartOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Order_CartArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Order_Cart.
type GetOvh_Order_CartResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Order_CartResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Order_CartResult)(nil)).Elem()
}

func (o GetOvh_Order_CartResultOutput) ToGetOvh_Order_CartResultOutput() GetOvh_Order_CartResultOutput {
	return o
}

func (o GetOvh_Order_CartResultOutput) ToGetOvh_Order_CartResultOutputWithContext(ctx context.Context) GetOvh_Order_CartResultOutput {
	return o
}

func (o GetOvh_Order_CartResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetOvh_Order_CartResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetOvh_Order_CartResultOutput) Expire() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) string { return v.Expire }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Order_CartResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Order_CartResultOutput) Items() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) []int { return v.Items }).(pulumi.IntArrayOutput)
}

func (o GetOvh_Order_CartResultOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) string { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

func (o GetOvh_Order_CartResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOvh_Order_CartResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Order_CartResultOutput{})
}
