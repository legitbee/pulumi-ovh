// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Vps(ctx *pulumi.Context, args *GetOvh_VpsArgs, opts ...pulumi.InvokeOption) (*GetOvh_VpsResult, error) {
	var rv GetOvh_VpsResult
	err := ctx.Invoke("ovh:index/getOvh_Vps:getOvh_Vps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Vps.
type GetOvh_VpsArgs struct {
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getOvh_Vps.
type GetOvh_VpsResult struct {
	Cluster     string            `pulumi:"cluster"`
	Datacenter  map[string]string `pulumi:"datacenter"`
	Displayname string            `pulumi:"displayname"`
	// The provider-assigned unique ID for this managed resource.
	Id            string            `pulumi:"id"`
	Ips           []string          `pulumi:"ips"`
	Keymap        string            `pulumi:"keymap"`
	Memory        int               `pulumi:"memory"`
	Model         map[string]string `pulumi:"model"`
	Name          string            `pulumi:"name"`
	Netbootmode   string            `pulumi:"netbootmode"`
	Offertype     string            `pulumi:"offertype"`
	ServiceName   string            `pulumi:"serviceName"`
	Slamonitoring bool              `pulumi:"slamonitoring"`
	State         string            `pulumi:"state"`
	Type          string            `pulumi:"type"`
	Vcore         int               `pulumi:"vcore"`
	Zone          string            `pulumi:"zone"`
}

func GetOvh_VpsOutput(ctx *pulumi.Context, args GetOvh_VpsOutputArgs, opts ...pulumi.InvokeOption) GetOvh_VpsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_VpsResult, error) {
			args := v.(GetOvh_VpsArgs)
			r, err := GetOvh_Vps(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_VpsResultOutput)
}

// A collection of arguments for invoking getOvh_Vps.
type GetOvh_VpsOutputArgs struct {
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetOvh_VpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_VpsArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Vps.
type GetOvh_VpsResultOutput struct{ *pulumi.OutputState }

func (GetOvh_VpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_VpsResult)(nil)).Elem()
}

func (o GetOvh_VpsResultOutput) ToGetOvh_VpsResultOutput() GetOvh_VpsResultOutput {
	return o
}

func (o GetOvh_VpsResultOutput) ToGetOvh_VpsResultOutputWithContext(ctx context.Context) GetOvh_VpsResultOutput {
	return o
}

func (o GetOvh_VpsResultOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Cluster }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Datacenter() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) map[string]string { return v.Datacenter }).(pulumi.StringMapOutput)
}

func (o GetOvh_VpsResultOutput) Displayname() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Displayname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_VpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

func (o GetOvh_VpsResultOutput) Keymap() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Keymap }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) int { return v.Memory }).(pulumi.IntOutput)
}

func (o GetOvh_VpsResultOutput) Model() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) map[string]string { return v.Model }).(pulumi.StringMapOutput)
}

func (o GetOvh_VpsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Netbootmode() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Netbootmode }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Offertype() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Offertype }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Slamonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) bool { return v.Slamonitoring }).(pulumi.BoolOutput)
}

func (o GetOvh_VpsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetOvh_VpsResultOutput) Vcore() pulumi.IntOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) int { return v.Vcore }).(pulumi.IntOutput)
}

func (o GetOvh_VpsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_VpsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_VpsResultOutput{})
}
