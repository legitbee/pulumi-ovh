// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOvh_Iploadbalancing(ctx *pulumi.Context, args *GetOvh_IploadbalancingArgs, opts ...pulumi.InvokeOption) (*GetOvh_IploadbalancingResult, error) {
	var rv GetOvh_IploadbalancingResult
	err := ctx.Invoke("ovh:index/getOvh_Iploadbalancing:getOvh_Iploadbalancing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvh_Iploadbalancing.
type GetOvh_IploadbalancingArgs struct {
	DisplayName      *string  `pulumi:"displayName"`
	IpLoadbalancing  *string  `pulumi:"ipLoadbalancing"`
	Ipv4             *string  `pulumi:"ipv4"`
	Ipv6             *string  `pulumi:"ipv6"`
	Offer            *string  `pulumi:"offer"`
	ServiceName      *string  `pulumi:"serviceName"`
	SslConfiguration *string  `pulumi:"sslConfiguration"`
	State            *string  `pulumi:"state"`
	VrackEligibility *bool    `pulumi:"vrackEligibility"`
	VrackName        *string  `pulumi:"vrackName"`
	Zones            []string `pulumi:"zones"`
}

// A collection of values returned by getOvh_Iploadbalancing.
type GetOvh_IploadbalancingResult struct {
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                `pulumi:"id"`
	IpLoadbalancing  string                                `pulumi:"ipLoadbalancing"`
	Ipv4             string                                `pulumi:"ipv4"`
	Ipv6             string                                `pulumi:"ipv6"`
	MetricsToken     string                                `pulumi:"metricsToken"`
	Offer            string                                `pulumi:"offer"`
	OrderableZones   []GetOvh_IploadbalancingOrderableZone `pulumi:"orderableZones"`
	ServiceName      string                                `pulumi:"serviceName"`
	SslConfiguration string                                `pulumi:"sslConfiguration"`
	State            string                                `pulumi:"state"`
	VrackEligibility bool                                  `pulumi:"vrackEligibility"`
	VrackName        string                                `pulumi:"vrackName"`
	Zones            []string                              `pulumi:"zones"`
}

func GetOvh_IploadbalancingOutput(ctx *pulumi.Context, args GetOvh_IploadbalancingOutputArgs, opts ...pulumi.InvokeOption) GetOvh_IploadbalancingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOvh_IploadbalancingResult, error) {
			args := v.(GetOvh_IploadbalancingArgs)
			r, err := GetOvh_Iploadbalancing(ctx, &args, opts...)
			return *r, err
		}).(GetOvh_IploadbalancingResultOutput)
}

// A collection of arguments for invoking getOvh_Iploadbalancing.
type GetOvh_IploadbalancingOutputArgs struct {
	DisplayName      pulumi.StringPtrInput   `pulumi:"displayName"`
	IpLoadbalancing  pulumi.StringPtrInput   `pulumi:"ipLoadbalancing"`
	Ipv4             pulumi.StringPtrInput   `pulumi:"ipv4"`
	Ipv6             pulumi.StringPtrInput   `pulumi:"ipv6"`
	Offer            pulumi.StringPtrInput   `pulumi:"offer"`
	ServiceName      pulumi.StringPtrInput   `pulumi:"serviceName"`
	SslConfiguration pulumi.StringPtrInput   `pulumi:"sslConfiguration"`
	State            pulumi.StringPtrInput   `pulumi:"state"`
	VrackEligibility pulumi.BoolPtrInput     `pulumi:"vrackEligibility"`
	VrackName        pulumi.StringPtrInput   `pulumi:"vrackName"`
	Zones            pulumi.StringArrayInput `pulumi:"zones"`
}

func (GetOvh_IploadbalancingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_IploadbalancingArgs)(nil)).Elem()
}

// A collection of values returned by getOvh_Iploadbalancing.
type GetOvh_IploadbalancingResultOutput struct{ *pulumi.OutputState }

func (GetOvh_IploadbalancingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOvh_IploadbalancingResult)(nil)).Elem()
}

func (o GetOvh_IploadbalancingResultOutput) ToGetOvh_IploadbalancingResultOutput() GetOvh_IploadbalancingResultOutput {
	return o
}

func (o GetOvh_IploadbalancingResultOutput) ToGetOvh_IploadbalancingResultOutputWithContext(ctx context.Context) GetOvh_IploadbalancingResultOutput {
	return o
}

func (o GetOvh_IploadbalancingResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOvh_IploadbalancingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) IpLoadbalancing() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.IpLoadbalancing }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.Ipv4 }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.Ipv6 }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) MetricsToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.MetricsToken }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.Offer }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) OrderableZones() GetOvh_IploadbalancingOrderableZoneArrayOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) []GetOvh_IploadbalancingOrderableZone { return v.OrderableZones }).(GetOvh_IploadbalancingOrderableZoneArrayOutput)
}

func (o GetOvh_IploadbalancingResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) SslConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.SslConfiguration }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) VrackEligibility() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) bool { return v.VrackEligibility }).(pulumi.BoolOutput)
}

func (o GetOvh_IploadbalancingResultOutput) VrackName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) string { return v.VrackName }).(pulumi.StringOutput)
}

func (o GetOvh_IploadbalancingResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOvh_IploadbalancingResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOvh_IploadbalancingResultOutput{})
}
