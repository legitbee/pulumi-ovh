// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOrderCartProduct(ctx *pulumi.Context, args *GetOrderCartProductArgs, opts ...pulumi.InvokeOption) (*GetOrderCartProductResult, error) {
	var rv GetOrderCartProductResult
	err := ctx.Invoke("ovh:index/getOrderCartProduct:getOrderCartProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderCartProduct.
type GetOrderCartProductArgs struct {
	CartId  string `pulumi:"cartId"`
	Product string `pulumi:"product"`
}

// A collection of values returned by getOrderCartProduct.
type GetOrderCartProductResult struct {
	CartId string `pulumi:"cartId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string                      `pulumi:"id"`
	Product string                      `pulumi:"product"`
	Results []GetOrderCartProductResult `pulumi:"results"`
}

func GetOrderCartProductOutput(ctx *pulumi.Context, args GetOrderCartProductOutputArgs, opts ...pulumi.InvokeOption) GetOrderCartProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrderCartProductResult, error) {
			args := v.(GetOrderCartProductArgs)
			r, err := GetOrderCartProduct(ctx, &args, opts...)
			return *r, err
		}).(GetOrderCartProductResultOutput)
}

// A collection of arguments for invoking getOrderCartProduct.
type GetOrderCartProductOutputArgs struct {
	CartId  pulumi.StringInput `pulumi:"cartId"`
	Product pulumi.StringInput `pulumi:"product"`
}

func (GetOrderCartProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartProductArgs)(nil)).Elem()
}

// A collection of values returned by getOrderCartProduct.
type GetOrderCartProductResultOutput struct{ *pulumi.OutputState }

func (GetOrderCartProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartProductResult)(nil)).Elem()
}

func (o GetOrderCartProductResultOutput) ToGetOrderCartProductResultOutput() GetOrderCartProductResultOutput {
	return o
}

func (o GetOrderCartProductResultOutput) ToGetOrderCartProductResultOutputWithContext(ctx context.Context) GetOrderCartProductResultOutput {
	return o
}

func (o GetOrderCartProductResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductResult) string { return v.CartId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrderCartProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrderCartProductResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductResult) string { return v.Product }).(pulumi.StringOutput)
}

func (o GetOrderCartProductResultOutput) Results() GetOrderCartProductResultArrayOutput {
	return o.ApplyT(func(v GetOrderCartProductResult) []GetOrderCartProductResult { return v.Results }).(GetOrderCartProductResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrderCartProductResultOutput{})
}
