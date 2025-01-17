// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Order_Cart_Product_Plan(ctx *pulumi.Context, args *GetOvh_Order_Cart_Product_PlanArgs, opts ...pulumi.InvokeOption) (*GetOvh_Order_Cart_Product_PlanResult, error) {
	var rv GetOvh_Order_Cart_Product_PlanResult
	err := ctx.Invoke("ovh:index/getOvh_Order_Cart_Product_Plan:getOvh_Order_Cart_Product_Plan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Order_Cart_Product_Plan.
type GetOvh_Order_Cart_Product_PlanArgs struct {
	CartId        string  `pulumi:"cartId"`
	CatalogName   *string `pulumi:"catalogName"`
	PlanCode      string  `pulumi:"planCode"`
	PriceCapacity string  `pulumi:"priceCapacity"`
	Product       string  `pulumi:"product"`
}

// A collection of values returned by getOvh_Order_Cart_Product_Plan.
type GetOvh_Order_Cart_Product_PlanResult struct {
	CartId      string  `pulumi:"cartId"`
	CatalogName *string `pulumi:"catalogName"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                                        `pulumi:"id"`
	PlanCode       string                                        `pulumi:"planCode"`
	PriceCapacity  string                                        `pulumi:"priceCapacity"`
	Prices         []GetOvh_Order_Cart_Product_PlanPrice         `pulumi:"prices"`
	Product        string                                        `pulumi:"product"`
	ProductName    string                                        `pulumi:"productName"`
	ProductType    string                                        `pulumi:"productType"`
	SelectedPrices []GetOvh_Order_Cart_Product_PlanSelectedPrice `pulumi:"selectedPrices"`
}

func GetOvh_Order_Cart_Product_PlanOutput(ctx *pulumi.Context, args GetOvh_Order_Cart_Product_PlanOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Order_Cart_Product_PlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Order_Cart_Product_PlanResult, error) {
			args := v.(GetOvh_Order_Cart_Product_PlanArgs)
			r, err := GetOvh_Order_Cart_Product_Plan(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Order_Cart_Product_PlanResultOutput)
}

// A collection of arguments for invoking getOvh_Order_Cart_Product_Plan.
type GetOvh_Order_Cart_Product_PlanOutputArgs struct {
	CartId        pulumi.StringInput    `pulumi:"cartId"`
	CatalogName   pulumi.StringPtrInput `pulumi:"catalogName"`
	PlanCode      pulumi.StringInput    `pulumi:"planCode"`
	PriceCapacity pulumi.StringInput    `pulumi:"priceCapacity"`
	Product       pulumi.StringInput    `pulumi:"product"`
}

func (GetOvh_Order_Cart_Product_PlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Order_Cart_Product_PlanArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Order_Cart_Product_Plan.
type GetOvh_Order_Cart_Product_PlanResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Order_Cart_Product_PlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Order_Cart_Product_PlanResult)(nil)).Elem()
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) ToGetOvh_Order_Cart_Product_PlanResultOutput() GetOvh_Order_Cart_Product_PlanResultOutput {
	return o
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) ToGetOvh_Order_Cart_Product_PlanResultOutputWithContext(ctx context.Context) GetOvh_Order_Cart_Product_PlanResultOutput {
	return o
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) CatalogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) *string { return v.CatalogName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Order_Cart_Product_PlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) PriceCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.PriceCapacity }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) Prices() GetOvh_Order_Cart_Product_PlanPriceArrayOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) []GetOvh_Order_Cart_Product_PlanPrice { return v.Prices }).(GetOvh_Order_Cart_Product_PlanPriceArrayOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.Product }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.ProductName }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) string { return v.ProductType }).(pulumi.StringOutput)
}

func (o GetOvh_Order_Cart_Product_PlanResultOutput) SelectedPrices() GetOvh_Order_Cart_Product_PlanSelectedPriceArrayOutput {
	return o.ApplyT(func(v GetOvh_Order_Cart_Product_PlanResult) []GetOvh_Order_Cart_Product_PlanSelectedPrice {
		return v.SelectedPrices
	}).(GetOvh_Order_Cart_Product_PlanSelectedPriceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Order_Cart_Product_PlanResultOutput{})
}
