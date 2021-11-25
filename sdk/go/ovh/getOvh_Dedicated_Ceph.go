// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Dedicated_Ceph(ctx *pulumi.Context, args *GetOvh_Dedicated_CephArgs, opts ...pulumi.InvokeOption) (*GetOvh_Dedicated_CephResult, error) {
	var rv GetOvh_Dedicated_CephResult
	err := ctx.Invoke("ovh:index/getOvh_Dedicated_Ceph:getOvh_Dedicated_Ceph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Dedicated_Ceph.
type GetOvh_Dedicated_CephArgs struct {
	CephVersion *string `pulumi:"cephVersion"`
	ServiceName string  `pulumi:"serviceName"`
	Status      *string `pulumi:"status"`
}

// A collection of values returned by getOvh_Dedicated_Ceph.
type GetOvh_Dedicated_CephResult struct {
	CephMons      []string `pulumi:"cephMons"`
	CephVersion   string   `pulumi:"cephVersion"`
	CrushTunables string   `pulumi:"crushTunables"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Label       string  `pulumi:"label"`
	Region      string  `pulumi:"region"`
	ServiceName string  `pulumi:"serviceName"`
	Size        float64 `pulumi:"size"`
	State       string  `pulumi:"state"`
	Status      string  `pulumi:"status"`
}

func GetOvh_Dedicated_CephOutput(ctx *pulumi.Context, args GetOvh_Dedicated_CephOutputArgs, opts ...pulumi.InvokeOption) GetOvh_Dedicated_CephResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_Dedicated_CephResult, error) {
			args := v.(GetOvh_Dedicated_CephArgs)
			r, err := GetOvh_Dedicated_Ceph(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_Dedicated_CephResultOutput)
}

// A collection of arguments for invoking getOvh_Dedicated_Ceph.
type GetOvh_Dedicated_CephOutputArgs struct {
	CephVersion pulumi.StringPtrInput `pulumi:"cephVersion"`
	ServiceName pulumi.StringInput    `pulumi:"serviceName"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (GetOvh_Dedicated_CephOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Dedicated_CephArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Dedicated_Ceph.
type GetOvh_Dedicated_CephResultOutput struct{ *pulumi.OutputState }

func (GetOvh_Dedicated_CephResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_Dedicated_CephResult)(nil)).Elem()
}

func (o GetOvh_Dedicated_CephResultOutput) ToGetOvh_Dedicated_CephResultOutput() GetOvh_Dedicated_CephResultOutput {
	return o
}

func (o GetOvh_Dedicated_CephResultOutput) ToGetOvh_Dedicated_CephResultOutputWithContext(ctx context.Context) GetOvh_Dedicated_CephResultOutput {
	return o
}

func (o GetOvh_Dedicated_CephResultOutput) CephMons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) []string { return v.CephMons }).(pulumi.StringArrayOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) CephVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.CephVersion }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) CrushTunables() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.CrushTunables }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_Dedicated_CephResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) float64 { return v.Size }).(pulumi.Float64Output)
}

func (o GetOvh_Dedicated_CephResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetOvh_Dedicated_CephResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_Dedicated_CephResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_Dedicated_CephResultOutput{})
}
