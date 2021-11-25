// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IploadbalancingHttpRoute struct {
	pulumi.CustomResourceState

	// Action triggered when all rules match
	Action IploadbalancingHttpRouteActionOutput `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId pulumi.IntOutput `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules IploadbalancingHttpRouteRuleTypeArrayOutput `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringOutput `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewIploadbalancingHttpRoute registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingHttpRoute(ctx *pulumi.Context,
	name string, args *IploadbalancingHttpRouteArgs, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource IploadbalancingHttpRoute
	err := ctx.RegisterResource("ovh:index/iploadbalancingHttpRoute:IploadbalancingHttpRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingHttpRoute gets an existing IploadbalancingHttpRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingHttpRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingHttpRouteState, opts ...pulumi.ResourceOption) (*IploadbalancingHttpRoute, error) {
	var resource IploadbalancingHttpRoute
	err := ctx.ReadResource("ovh:index/iploadbalancingHttpRoute:IploadbalancingHttpRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingHttpRoute resources.
type iploadbalancingHttpRouteState struct {
	// Action triggered when all rules match
	Action *IploadbalancingHttpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules []IploadbalancingHttpRouteRuleType `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status *string `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight *int `pulumi:"weight"`
}

type IploadbalancingHttpRouteState struct {
	// Action triggered when all rules match
	Action IploadbalancingHttpRouteActionPtrInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// List of rules to match to trigger action
	Rules IploadbalancingHttpRouteRuleTypeArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringPtrInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight pulumi.IntPtrInput
}

func (IploadbalancingHttpRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteState)(nil)).Elem()
}

type iploadbalancingHttpRouteArgs struct {
	// Action triggered when all rules match
	Action IploadbalancingHttpRouteAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a IploadbalancingHttpRoute resource.
type IploadbalancingHttpRouteArgs struct {
	// Action triggered when all rules match
	Action IploadbalancingHttpRouteActionInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight pulumi.IntPtrInput
}

func (IploadbalancingHttpRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingHttpRouteArgs)(nil)).Elem()
}

type IploadbalancingHttpRouteInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteOutput() IploadbalancingHttpRouteOutput
	ToIploadbalancingHttpRouteOutputWithContext(ctx context.Context) IploadbalancingHttpRouteOutput
}

func (*IploadbalancingHttpRoute) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingHttpRoute)(nil))
}

func (i *IploadbalancingHttpRoute) ToIploadbalancingHttpRouteOutput() IploadbalancingHttpRouteOutput {
	return i.ToIploadbalancingHttpRouteOutputWithContext(context.Background())
}

func (i *IploadbalancingHttpRoute) ToIploadbalancingHttpRouteOutputWithContext(ctx context.Context) IploadbalancingHttpRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteOutput)
}

func (i *IploadbalancingHttpRoute) ToIploadbalancingHttpRoutePtrOutput() IploadbalancingHttpRoutePtrOutput {
	return i.ToIploadbalancingHttpRoutePtrOutputWithContext(context.Background())
}

func (i *IploadbalancingHttpRoute) ToIploadbalancingHttpRoutePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRoutePtrOutput)
}

type IploadbalancingHttpRoutePtrInput interface {
	pulumi.Input

	ToIploadbalancingHttpRoutePtrOutput() IploadbalancingHttpRoutePtrOutput
	ToIploadbalancingHttpRoutePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRoutePtrOutput
}

type iploadbalancingHttpRoutePtrType IploadbalancingHttpRouteArgs

func (*iploadbalancingHttpRoutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRoute)(nil))
}

func (i *iploadbalancingHttpRoutePtrType) ToIploadbalancingHttpRoutePtrOutput() IploadbalancingHttpRoutePtrOutput {
	return i.ToIploadbalancingHttpRoutePtrOutputWithContext(context.Background())
}

func (i *iploadbalancingHttpRoutePtrType) ToIploadbalancingHttpRoutePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRoutePtrOutput)
}

// IploadbalancingHttpRouteArrayInput is an input type that accepts IploadbalancingHttpRouteArray and IploadbalancingHttpRouteArrayOutput values.
// You can construct a concrete instance of `IploadbalancingHttpRouteArrayInput` via:
//
//          IploadbalancingHttpRouteArray{ IploadbalancingHttpRouteArgs{...} }
type IploadbalancingHttpRouteArrayInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteArrayOutput() IploadbalancingHttpRouteArrayOutput
	ToIploadbalancingHttpRouteArrayOutputWithContext(context.Context) IploadbalancingHttpRouteArrayOutput
}

type IploadbalancingHttpRouteArray []IploadbalancingHttpRouteInput

func (IploadbalancingHttpRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IploadbalancingHttpRoute)(nil)).Elem()
}

func (i IploadbalancingHttpRouteArray) ToIploadbalancingHttpRouteArrayOutput() IploadbalancingHttpRouteArrayOutput {
	return i.ToIploadbalancingHttpRouteArrayOutputWithContext(context.Background())
}

func (i IploadbalancingHttpRouteArray) ToIploadbalancingHttpRouteArrayOutputWithContext(ctx context.Context) IploadbalancingHttpRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteArrayOutput)
}

// IploadbalancingHttpRouteMapInput is an input type that accepts IploadbalancingHttpRouteMap and IploadbalancingHttpRouteMapOutput values.
// You can construct a concrete instance of `IploadbalancingHttpRouteMapInput` via:
//
//          IploadbalancingHttpRouteMap{ "key": IploadbalancingHttpRouteArgs{...} }
type IploadbalancingHttpRouteMapInput interface {
	pulumi.Input

	ToIploadbalancingHttpRouteMapOutput() IploadbalancingHttpRouteMapOutput
	ToIploadbalancingHttpRouteMapOutputWithContext(context.Context) IploadbalancingHttpRouteMapOutput
}

type IploadbalancingHttpRouteMap map[string]IploadbalancingHttpRouteInput

func (IploadbalancingHttpRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IploadbalancingHttpRoute)(nil)).Elem()
}

func (i IploadbalancingHttpRouteMap) ToIploadbalancingHttpRouteMapOutput() IploadbalancingHttpRouteMapOutput {
	return i.ToIploadbalancingHttpRouteMapOutputWithContext(context.Background())
}

func (i IploadbalancingHttpRouteMap) ToIploadbalancingHttpRouteMapOutputWithContext(ctx context.Context) IploadbalancingHttpRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingHttpRouteMapOutput)
}

type IploadbalancingHttpRouteOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingHttpRoute)(nil))
}

func (o IploadbalancingHttpRouteOutput) ToIploadbalancingHttpRouteOutput() IploadbalancingHttpRouteOutput {
	return o
}

func (o IploadbalancingHttpRouteOutput) ToIploadbalancingHttpRouteOutputWithContext(ctx context.Context) IploadbalancingHttpRouteOutput {
	return o
}

func (o IploadbalancingHttpRouteOutput) ToIploadbalancingHttpRoutePtrOutput() IploadbalancingHttpRoutePtrOutput {
	return o.ToIploadbalancingHttpRoutePtrOutputWithContext(context.Background())
}

func (o IploadbalancingHttpRouteOutput) ToIploadbalancingHttpRoutePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRoutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IploadbalancingHttpRoute) *IploadbalancingHttpRoute {
		return &v
	}).(IploadbalancingHttpRoutePtrOutput)
}

type IploadbalancingHttpRoutePtrOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingHttpRoute)(nil))
}

func (o IploadbalancingHttpRoutePtrOutput) ToIploadbalancingHttpRoutePtrOutput() IploadbalancingHttpRoutePtrOutput {
	return o
}

func (o IploadbalancingHttpRoutePtrOutput) ToIploadbalancingHttpRoutePtrOutputWithContext(ctx context.Context) IploadbalancingHttpRoutePtrOutput {
	return o
}

func (o IploadbalancingHttpRoutePtrOutput) Elem() IploadbalancingHttpRouteOutput {
	return o.ApplyT(func(v *IploadbalancingHttpRoute) IploadbalancingHttpRoute {
		if v != nil {
			return *v
		}
		var ret IploadbalancingHttpRoute
		return ret
	}).(IploadbalancingHttpRouteOutput)
}

type IploadbalancingHttpRouteArrayOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IploadbalancingHttpRoute)(nil))
}

func (o IploadbalancingHttpRouteArrayOutput) ToIploadbalancingHttpRouteArrayOutput() IploadbalancingHttpRouteArrayOutput {
	return o
}

func (o IploadbalancingHttpRouteArrayOutput) ToIploadbalancingHttpRouteArrayOutputWithContext(ctx context.Context) IploadbalancingHttpRouteArrayOutput {
	return o
}

func (o IploadbalancingHttpRouteArrayOutput) Index(i pulumi.IntInput) IploadbalancingHttpRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IploadbalancingHttpRoute {
		return vs[0].([]IploadbalancingHttpRoute)[vs[1].(int)]
	}).(IploadbalancingHttpRouteOutput)
}

type IploadbalancingHttpRouteMapOutput struct{ *pulumi.OutputState }

func (IploadbalancingHttpRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IploadbalancingHttpRoute)(nil))
}

func (o IploadbalancingHttpRouteMapOutput) ToIploadbalancingHttpRouteMapOutput() IploadbalancingHttpRouteMapOutput {
	return o
}

func (o IploadbalancingHttpRouteMapOutput) ToIploadbalancingHttpRouteMapOutputWithContext(ctx context.Context) IploadbalancingHttpRouteMapOutput {
	return o
}

func (o IploadbalancingHttpRouteMapOutput) MapIndex(k pulumi.StringInput) IploadbalancingHttpRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IploadbalancingHttpRoute {
		return vs[0].(map[string]IploadbalancingHttpRoute)[vs[1].(string)]
	}).(IploadbalancingHttpRouteOutput)
}

func init() {
	pulumi.RegisterOutputType(IploadbalancingHttpRouteOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRoutePtrOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteArrayOutput{})
	pulumi.RegisterOutputType(IploadbalancingHttpRouteMapOutput{})
}