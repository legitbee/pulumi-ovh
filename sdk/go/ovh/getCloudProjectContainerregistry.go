// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudProjectContainerregistry(ctx *pulumi.Context, args *LookupCloudProjectContainerregistryArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectContainerregistryResult, error) {
	var rv LookupCloudProjectContainerregistryResult
	err := ctx.Invoke("ovh:index/getCloudProjectContainerregistry:getCloudProjectContainerregistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectContainerregistry.
type LookupCloudProjectContainerregistryArgs struct {
	RegistryId  string `pulumi:"registryId"`
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectContainerregistry.
type LookupCloudProjectContainerregistryResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	ProjectId   string `pulumi:"projectId"`
	Region      string `pulumi:"region"`
	RegistryId  string `pulumi:"registryId"`
	ServiceName string `pulumi:"serviceName"`
	Size        int    `pulumi:"size"`
	Status      string `pulumi:"status"`
	UpdatedAt   string `pulumi:"updatedAt"`
	Url         string `pulumi:"url"`
	Version     string `pulumi:"version"`
}

func LookupCloudProjectContainerregistryOutput(ctx *pulumi.Context, args LookupCloudProjectContainerregistryOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectContainerregistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectContainerregistryResult, error) {
			args := v.(LookupCloudProjectContainerregistryArgs)
			r, err := LookupCloudProjectContainerregistry(ctx, &args, opts...)
			return *r, err
		}).(LookupCloudProjectContainerregistryResultOutput)
}

// A collection of arguments for invoking getCloudProjectContainerregistry.
type LookupCloudProjectContainerregistryOutputArgs struct {
	RegistryId  pulumi.StringInput `pulumi:"registryId"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCloudProjectContainerregistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectContainerregistryArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectContainerregistry.
type LookupCloudProjectContainerregistryResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectContainerregistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectContainerregistryResult)(nil)).Elem()
}

func (o LookupCloudProjectContainerregistryResultOutput) ToLookupCloudProjectContainerregistryResultOutput() LookupCloudProjectContainerregistryResultOutput {
	return o
}

func (o LookupCloudProjectContainerregistryResultOutput) ToLookupCloudProjectContainerregistryResultOutputWithContext(ctx context.Context) LookupCloudProjectContainerregistryResultOutput {
	return o
}

func (o LookupCloudProjectContainerregistryResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProjectContainerregistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Url }).(pulumi.StringOutput)
}

func (o LookupCloudProjectContainerregistryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectContainerregistryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectContainerregistryResultOutput{})
}