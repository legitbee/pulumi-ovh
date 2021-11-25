// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDbaasLogsInputEngine(ctx *pulumi.Context, args *GetDbaasLogsInputEngineArgs, opts ...pulumi.InvokeOption) (*GetDbaasLogsInputEngineResult, error) {
	var rv GetDbaasLogsInputEngineResult
	err := ctx.Invoke("ovh:index/getDbaasLogsInputEngine:getDbaasLogsInputEngine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDbaasLogsInputEngine.
type GetDbaasLogsInputEngineArgs struct {
	IsDeprecated *bool   `pulumi:"isDeprecated"`
	Name         *string `pulumi:"name"`
	Version      *string `pulumi:"version"`
}

// A collection of values returned by getDbaasLogsInputEngine.
type GetDbaasLogsInputEngineResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	IsDeprecated bool   `pulumi:"isDeprecated"`
	Name         string `pulumi:"name"`
	Version      string `pulumi:"version"`
}

func GetDbaasLogsInputEngineOutput(ctx *pulumi.Context, args GetDbaasLogsInputEngineOutputArgs, opts ...pulumi.InvokeOption) GetDbaasLogsInputEngineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDbaasLogsInputEngineResult, error) {
			args := v.(GetDbaasLogsInputEngineArgs)
			r, err := GetDbaasLogsInputEngine(ctx, &args, opts...)
			return *r, err
		}).(GetDbaasLogsInputEngineResultOutput)
}

// A collection of arguments for invoking getDbaasLogsInputEngine.
type GetDbaasLogsInputEngineOutputArgs struct {
	IsDeprecated pulumi.BoolPtrInput   `pulumi:"isDeprecated"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (GetDbaasLogsInputEngineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbaasLogsInputEngineArgs)(nil)).Elem()
}

// A collection of values returned by getDbaasLogsInputEngine.
type GetDbaasLogsInputEngineResultOutput struct{ *pulumi.OutputState }

func (GetDbaasLogsInputEngineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDbaasLogsInputEngineResult)(nil)).Elem()
}

func (o GetDbaasLogsInputEngineResultOutput) ToGetDbaasLogsInputEngineResultOutput() GetDbaasLogsInputEngineResultOutput {
	return o
}

func (o GetDbaasLogsInputEngineResultOutput) ToGetDbaasLogsInputEngineResultOutputWithContext(ctx context.Context) GetDbaasLogsInputEngineResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDbaasLogsInputEngineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbaasLogsInputEngineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDbaasLogsInputEngineResultOutput) IsDeprecated() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDbaasLogsInputEngineResult) bool { return v.IsDeprecated }).(pulumi.BoolOutput)
}

func (o GetDbaasLogsInputEngineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbaasLogsInputEngineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDbaasLogsInputEngineResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetDbaasLogsInputEngineResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDbaasLogsInputEngineResultOutput{})
}
