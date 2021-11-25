// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Me_Paymentmean_Bankaccount(ctx *pulumi.Context, args *GetOvh_Me_Paymentmean_BankaccountArgs, opts ...pulumi.InvokeOption) (*GetOvh_Me_Paymentmean_BankaccountResult, error) {
	var rv GetOvh_Me_Paymentmean_BankaccountResult
	err := ctx.Invoke("ovh:index/getOvh_Me_Paymentmean_Bankaccount:getOvh_Me_Paymentmean_Bankaccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Me_Paymentmean_Bankaccount.
type GetOvh_Me_Paymentmean_BankaccountArgs struct {
	DescriptionRegexp *string `pulumi:"descriptionRegexp"`
	State             *string `pulumi:"state"`
	UseDefault        *bool   `pulumi:"useDefault"`
	UseOldest         *bool   `pulumi:"useOldest"`
}

// A collection of values returned by getOvh_Me_Paymentmean_Bankaccount.
type GetOvh_Me_Paymentmean_BankaccountResult struct {
	Default           bool    `pulumi:"default"`
	Description       string  `pulumi:"description"`
	DescriptionRegexp *string `pulumi:"descriptionRegexp"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	State      string `pulumi:"state"`
	UseDefault *bool  `pulumi:"useDefault"`
	UseOldest  *bool  `pulumi:"useOldest"`
}

func GetOvh_Me_Paymentmean_BankaccountOutput(ctx *pulumi.Context, args GetOvh_Me_Paymentmean_BankaccountOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Me_Paymentmean_BankaccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Me_Paymentmean_BankaccountResult, error) {
			args := v.(GetOvh_Me_Paymentmean_BankaccountArgs)
			r, err := GetOvh_Me_Paymentmean_Bankaccount(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Me_Paymentmean_BankaccountResultOutput)
}

// A collection of arguments for invoking getOvh_Me_Paymentmean_Bankaccount.
type GetOvh_Me_Paymentmean_BankaccountOutputArgs struct {
	DescriptionRegexp pulumi.StringPtrInput `pulumi:"descriptionRegexp"`
	State             pulumi.StringPtrInput `pulumi:"state"`
	UseDefault        pulumi.BoolPtrInput   `pulumi:"useDefault"`
	UseOldest         pulumi.BoolPtrInput   `pulumi:"useOldest"`
}

func (GetOvh_Me_Paymentmean_BankaccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Me_Paymentmean_BankaccountArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Me_Paymentmean_Bankaccount.
type GetOvh_Me_Paymentmean_BankaccountResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Me_Paymentmean_BankaccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Me_Paymentmean_BankaccountResult)(nil)).Elem()
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) ToGetOvh_Me_Paymentmean_BankaccountResultOutput() GetOvh_Me_Paymentmean_BankaccountResultOutput {
	return o
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) ToGetOvh_Me_Paymentmean_BankaccountResultOutputWithContext(ctx context.Context) GetOvh_Me_Paymentmean_BankaccountResultOutput {
	return o
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) bool { return v.Default }).(pulumi.BoolOutput)
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) DescriptionRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) *string { return v.DescriptionRegexp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) UseDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) *bool { return v.UseDefault }).(pulumi.BoolPtrOutput)
}

func (o GetOvh_Me_Paymentmean_BankaccountResultOutput) UseOldest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOvh_Me_Paymentmean_BankaccountResult) *bool { return v.UseOldest }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Me_Paymentmean_BankaccountResultOutput{})
}
