// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIploadbalancing(ctx *pulumi.Context, args *LookupIploadbalancingArgs, opts ...pulumi.InvokeOption) (*LookupIploadbalancingResult, error) {
	var rv LookupIploadbalancingResult
	err := ctx.Invoke("ovh:index/getIploadbalancing:getIploadbalancing", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIploadbalancing.
type LookupIploadbalancingArgs struct {
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

// A collection of values returned by getIploadbalancing.
type LookupIploadbalancingResult struct {
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                            `pulumi:"id"`
	IpLoadbalancing  string                            `pulumi:"ipLoadbalancing"`
	Ipv4             string                            `pulumi:"ipv4"`
	Ipv6             string                            `pulumi:"ipv6"`
	MetricsToken     string                            `pulumi:"metricsToken"`
	Offer            string                            `pulumi:"offer"`
	OrderableZones   []GetIploadbalancingOrderableZone `pulumi:"orderableZones"`
	ServiceName      string                            `pulumi:"serviceName"`
	SslConfiguration string                            `pulumi:"sslConfiguration"`
	State            string                            `pulumi:"state"`
	VrackEligibility bool                              `pulumi:"vrackEligibility"`
	VrackName        string                            `pulumi:"vrackName"`
	Zones            []string                          `pulumi:"zones"`
}

func LookupIploadbalancingOutput(ctx *pulumi.Context, args LookupIploadbalancingOutputArgs, opts ...pulumi.InvokeOption) LookupIploadbalancingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIploadbalancingResult, error) {
			args := v.(LookupIploadbalancingArgs)
			r, err := LookupIploadbalancing(ctx, &args, opts...)
			return *r, err
		}).(LookupIploadbalancingResultOutput)
}

// A collection of arguments for invoking getIploadbalancing.
type LookupIploadbalancingOutputArgs struct {
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

func (LookupIploadbalancingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIploadbalancingArgs)(nil)).Elem()
}

// A collection of values returned by getIploadbalancing.
type LookupIploadbalancingResultOutput struct{ *pulumi.OutputState }

func (LookupIploadbalancingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIploadbalancingResult)(nil)).Elem()
}

func (o LookupIploadbalancingResultOutput) ToLookupIploadbalancingResultOutput() LookupIploadbalancingResultOutput {
	return o
}

func (o LookupIploadbalancingResultOutput) ToLookupIploadbalancingResultOutputWithContext(ctx context.Context) LookupIploadbalancingResultOutput {
	return o
}

func (o LookupIploadbalancingResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIploadbalancingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) IpLoadbalancing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.IpLoadbalancing }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.Ipv4 }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.Ipv6 }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) MetricsToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.MetricsToken }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.Offer }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) OrderableZones() GetIploadbalancingOrderableZoneArrayOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) []GetIploadbalancingOrderableZone { return v.OrderableZones }).(GetIploadbalancingOrderableZoneArrayOutput)
}

func (o LookupIploadbalancingResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) SslConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.SslConfiguration }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) VrackEligibility() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) bool { return v.VrackEligibility }).(pulumi.BoolOutput)
}

func (o LookupIploadbalancingResultOutput) VrackName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) string { return v.VrackName }).(pulumi.StringOutput)
}

func (o LookupIploadbalancingResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIploadbalancingResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIploadbalancingResultOutput{})
}