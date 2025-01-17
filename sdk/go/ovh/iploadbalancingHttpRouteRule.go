// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IploadbalancingHttpRouteRule struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	Field       pulumi.StringOutput    `pulumi:"field"`
	Match       pulumi.StringOutput    `pulumi:"match"`
	Negate      pulumi.BoolOutput      `pulumi:"negate"`
	Pattern     pulumi.StringPtrOutput `pulumi:"pattern"`
	RouteId     pulumi.StringOutput    `pulumi:"routeId"`
	ServiceName pulumi.StringOutput    `pulumi:"serviceName"`
	SubField    pulumi.StringPtrOutput `pulumi:"subField"`
}

// NewIploadbalancingHttpRouteRule registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingHttpRouteRule(ctx *pulumi.Context,
	name string, args *IploadbalancingHttpRouteRuleArgs, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRouteRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource IploadbalancingHttpRouteRule
	err := ctx.RegisterResource("ovh:index/iploadbalancingHttpRouteRule:IploadbalancingHttpRouteRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingHttpRouteRule gets an existing IploadbalancingHttpRouteRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingHttpRouteRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingHttpRouteRuleState, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRouteRule, error) {
	var resource IploadbalancingHttpRouteRule
	err := ctx.ReadResource("ovh:index/iploadbalancingHttpRouteRule:IploadbalancingHttpRouteRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingHttpRouteRule resources.
type iploadbalancingHttpRouteRuleState struct {
	DisplayName *string `pulumi:"displayName"`
	Field       *string `pulumi:"field"`
	Match       *string `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     *string `pulumi:"routeId"`
	ServiceName *string `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

type IploadbalancingHttpRouteRuleState struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringPtrInput
	Match       pulumi.StringPtrInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingHttpRouteRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteRuleState)(nil)).Elem()
}

type iploadbalancingHttpRouteRuleArgs struct {
	DisplayName *string `pulumi:"displayName"`
	Field       string  `pulumi:"field"`
	Match       string  `pulumi:"match"`
	Negate      *bool   `pulumi:"negate"`
	Pattern     *string `pulumi:"pattern"`
	RouteId     string  `pulumi:"routeId"`
	ServiceName string  `pulumi:"serviceName"`
	SubField    *string `pulumi:"subField"`
}

// The set of arguments for constructing a IploadbalancingHttpRouteRule resource.
type IploadbalancingHttpRouteRuleArgs struct {
	DisplayName pulumi.StringPtrInput
	Field       pulumi.StringInput
	Match       pulumi.StringInput
	Negate      pulumi.BoolPtrInput
	Pattern     pulumi.StringPtrInput
	RouteId     pulumi.StringInput
	ServiceName pulumi.StringInput
	SubField    pulumi.StringPtrInput
}

func (IploadbalancingHttpRouteRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteRuleArgs)(nil)).Elem()
}

type IploadbalancingHttpRouteRuleInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput
	ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput
}

func (*IploadbalancingHttpRouteRule) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingHttpRouteRule)(nil))
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput {
	return i.ToIploadbalancingHttpRouteRuleOutputWithContext(context.Background())
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRuleOutput)
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRulePtrOutput() IploadbalancingHttpRouteRulePtrOutput {
	return i.ToIploadbalancingHttpRouteRulePtrOutputWithContext(context.Background())
}

func (i *IploadbalancingHttpRouteRule) ToIploadbalancingHttpRouteRulePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRulePtrOutput)
}

type IploadbalancingHttpRouteRulePtrInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteRulePtrOutput() IploadbalancingHttpRouteRulePtrOutput
	ToIploadbalancingHttpRouteRulePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRulePtrOutput
}

type iploadbalancingHttpRouteRulePtrType IploadbalancingHttpRouteRuleArgs

func (*iploadbalancingHttpRouteRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRouteRule)(nil))
}

func (i *iploadbalancingHttpRouteRulePtrType) ToIploadbalancingHttpRouteRulePtrOutput() IploadbalancingHttpRouteRulePtrOutput {
	return i.ToIploadbalancingHttpRouteRulePtrOutputWithContext(context.Background())
}

func (i *iploadbalancingHttpRouteRulePtrType) ToIploadbalancingHttpRouteRulePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRulePtrOutput)
}

// IploadbalancingHttpRouteRuleArrayInput is an input type that accepts IploadbalancingHttpRouteRuleArray and IploadbalancingHttpRouteRuleArrayOutput values.
// You can construct a concrete instance of `IploadbalancingHttpRouteRuleArrayInput` via:
//
//          IploadbalancingHttpRouteRuleArray{ IploadbalancingHttpRouteRuleArgs{...} }
type IploadbalancingHttpRouteRuleArrayInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteRuleArrayOutput() IploadbalancingHttpRouteRuleArrayOutput
	ToIploadbalancingHttpRouteRuleArrayOutputWithContext(context.Context) IploadbalancingHttpRouteRuleArrayOutput
}

type IploadbalancingHttpRouteRuleArray []IploadbalancingHttpRouteRuleInput

func (IploadbalancingHttpRouteRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IploadbalancingHttpRouteRule)(nil)).Elem()
}

func (i IploadbalancingHttpRouteRuleArray) ToIploadbalancingHttpRouteRuleArrayOutput() IploadbalancingHttpRouteRuleArrayOutput {
	return i.ToIploadbalancingHttpRouteRuleArrayOutputWithContext(context.Background())
}

func (i IploadbalancingHttpRouteRuleArray) ToIploadbalancingHttpRouteRuleArrayOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRuleArrayOutput)
}

// IploadbalancingHttpRouteRuleMapInput is an input type that accepts IploadbalancingHttpRouteRuleMap and IploadbalancingHttpRouteRuleMapOutput values.
// You can construct a concrete instance of `IploadbalancingHttpRouteRuleMapInput` via:
//
//          IploadbalancingHttpRouteRuleMap{ "key": IploadbalancingHttpRouteRuleArgs{...} }
type IploadbalancingHttpRouteRuleMapInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteRuleMapOutput() IploadbalancingHttpRouteRuleMapOutput
	ToIploadbalancingHttpRouteRuleMapOutputWithContext(context.Context) IploadbalancingHttpRouteRuleMapOutput
}

type IploadbalancingHttpRouteRuleMap map[string]IploadbalancingHttpRouteRuleInput

func (IploadbalancingHttpRouteRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IploadbalancingHttpRouteRule)(nil)).Elem()
}

func (i IploadbalancingHttpRouteRuleMap) ToIploadbalancingHttpRouteRuleMapOutput() IploadbalancingHttpRouteRuleMapOutput {
	return i.ToIploadbalancingHttpRouteRuleMapOutputWithContext(context.Background())
}

func (i IploadbalancingHttpRouteRuleMap) ToIploadbalancingHttpRouteRuleMapOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteRuleMapOutput)
}

type IploadbalancingHttpRouteRuleOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingHttpRouteRule)(nil))
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRuleOutput() IploadbalancingHttpRouteRuleOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRuleOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRulePtrOutput() IploadbalancingHttpRouteRulePtrOutput {
	return o.ToIploadbalancingHttpRouteRulePtrOutputWithContext(context.Background())
}

func (o IploadbalancingHttpRouteRuleOutput) ToIploadbalancingHttpRouteRulePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IploadbalancingHttpRouteRule) *IploadbalancingHttpRouteRule {
		return &v
	}).(IploadbalancingHttpRouteRulePtrOutput)
}

type IploadbalancingHttpRouteRulePtrOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRouteRule)(nil))
}

func (o IploadbalancingHttpRouteRulePtrOutput) ToIploadbalancingHttpRouteRulePtrOutput() IploadbalancingHttpRouteRulePtrOutput {
	return o
}

func (o IploadbalancingHttpRouteRulePtrOutput) ToIploadbalancingHttpRouteRulePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRulePtrOutput {
	return o
}

func (o IploadbalancingHttpRouteRulePtrOutput) Elem() IploadbalancingHttpRouteRuleOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRouteRule) IploadbalancingHttpRouteRule {
		if v != nil {
			return *v
		}
		var ret IploadbalancingHttpRouteRule
		return ret
	}).(IploadbalancingHttpRouteRuleOutput)
}

type IploadbalancingHttpRouteRuleArrayOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IploadbalancingHttpRouteRule)(nil))
}

func (o IploadbalancingHttpRouteRuleArrayOutput) ToIploadbalancingHttpRouteRuleArrayOutput() IploadbalancingHttpRouteRuleArrayOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleArrayOutput) ToIploadbalancingHttpRouteRuleArrayOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleArrayOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleArrayOutput) Index(i pulumi.IntInput) IploadbalancingHttpRouteRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IploadbalancingHttpRouteRule {
		return vs[0].([]IploadbalancingHttpRouteRule)[vs[1].(int)]
	}).(IploadbalancingHttpRouteRuleOutput)
}

type IploadbalancingHttpRouteRuleMapOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IploadbalancingHttpRouteRule)(nil))
}

func (o IploadbalancingHttpRouteRuleMapOutput) ToIploadbalancingHttpRouteRuleMapOutput() IploadbalancingHttpRouteRuleMapOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleMapOutput) ToIploadbalancingHttpRouteRuleMapOutputWithContext(ctx context.Context) IploadbalancingHttpRouteRuleMapOutput {
	return o
}

func (o IploadbalancingHttpRouteRuleMapOutput) MapIndex(k pulumi.StringInput) IploadbalancingHttpRouteRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IploadbalancingHttpRouteRule {
		return vs[0].(map[string]IploadbalancingHttpRouteRule)[vs[1].(string)]
	}).(IploadbalancingHttpRouteRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(IploadbalancingHttpRouteRuleOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteRulePtrOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteRuleArrayOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteRuleMapOutput{})
}
