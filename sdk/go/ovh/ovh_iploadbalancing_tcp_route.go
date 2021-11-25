// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_iploadbalancing_tcp_route struct {
	pulumi.CustomResourceState

	// Action triggered when all rules match
	Action Ovh_iploadbalancing_tcp_routeActionOutput `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId pulumi.IntOutput `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules Ovh_iploadbalancing_tcp_routeRuleArrayOutput `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringOutput `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewOvh_iploadbalancing_tcp_route registers a new resource with the given unique name, arguments, and options.
func NewOvh_iploadbalancing_tcp_route(ctx *pulumi.Context,
	name string, args *Ovh_iploadbalancing_tcp_routeArgs, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_iploadbalancing_tcp_route
	err := ctx.RegisterResource("ovh:index/ovh_iploadbalancing_tcp_route:ovh_iploadbalancing_tcp_route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_iploadbalancing_tcp_route gets an existing Ovh_iploadbalancing_tcp_route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_iploadbalancing_tcp_route(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_iploadbalancing_tcp_routeState, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_route, error) {
	var resource Ovh_iploadbalancing_tcp_route
	err := ctx.ReadResource("ovh:index/ovh_iploadbalancing_tcp_route:ovh_iploadbalancing_tcp_route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_iploadbalancing_tcp_route resources.
type ovh_iploadbalancing_tcp_routeState struct {
	// Action triggered when all rules match
	Action *Ovh_iploadbalancing_tcp_routeAction `pulumi:"action"`
	// Human readable name for your route, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Route traffic for this frontend
	FrontendId *int `pulumi:"frontendId"`
	// List of rules to match to trigger action
	Rules []Ovh_iploadbalancing_tcp_routeRule `pulumi:"rules"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Route status. Routes in "ok" state are ready to operate
	Status *string `pulumi:"status"`
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight *int `pulumi:"weight"`
}

type Ovh_iploadbalancing_tcp_routeState struct {
	// Action triggered when all rules match
	Action Ovh_iploadbalancing_tcp_routeActionPtrInput
	// Human readable name for your route, this field is for you
	DisplayName pulumi.StringPtrInput
	// Route traffic for this frontend
	FrontendId pulumi.IntPtrInput
	// List of rules to match to trigger action
	Rules Ovh_iploadbalancing_tcp_routeRuleArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Route status. Routes in "ok" state are ready to operate
	Status pulumi.StringPtrInput
	// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated last. Only the first matching route will
	// trigger an action
	Weight pulumi.IntPtrInput
}

func (Ovh_iploadbalancing_tcp_routeState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_routeState)(nil)).Elem()
}

type ovh_iploadbalancing_tcp_routeArgs struct {
	// Action triggered when all rules match
	Action Ovh_iploadbalancing_tcp_routeAction `pulumi:"action"`
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

// The set of arguments for constructing a Ovh_iploadbalancing_tcp_route resource.
type Ovh_iploadbalancing_tcp_routeArgs struct {
	// Action triggered when all rules match
	Action Ovh_iploadbalancing_tcp_routeActionInput
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

func (Ovh_iploadbalancing_tcp_routeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_routeArgs)(nil)).Elem()
}

type Ovh_iploadbalancing_tcp_routeInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_routeOutput() Ovh_iploadbalancing_tcp_routeOutput
	ToOvh_iploadbalancing_tcp_routeOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeOutput
}

func (*Ovh_iploadbalancing_tcp_route) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_route)(nil))
}

func (i *Ovh_iploadbalancing_tcp_route) ToOvh_iploadbalancing_tcp_routeOutput() Ovh_iploadbalancing_tcp_routeOutput {
	return i.ToOvh_iploadbalancing_tcp_routeOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_route) ToOvh_iploadbalancing_tcp_routeOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_routeOutput)
}

func (i *Ovh_iploadbalancing_tcp_route) ToOvh_iploadbalancing_tcp_routePtrOutput() Ovh_iploadbalancing_tcp_routePtrOutput {
	return i.ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_route) ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_routePtrOutput)
}

type Ovh_iploadbalancing_tcp_routePtrInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_routePtrOutput() Ovh_iploadbalancing_tcp_routePtrOutput
	ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routePtrOutput
}

type ovh_iploadbalancing_tcp_routePtrType Ovh_iploadbalancing_tcp_routeArgs

func (*ovh_iploadbalancing_tcp_routePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_route)(nil))
}

func (i *ovh_iploadbalancing_tcp_routePtrType) ToOvh_iploadbalancing_tcp_routePtrOutput() Ovh_iploadbalancing_tcp_routePtrOutput {
	return i.ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(context.Background())
}

func (i *ovh_iploadbalancing_tcp_routePtrType) ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_routePtrOutput)
}

// Ovh_iploadbalancing_tcp_routeArrayInput is an input type that accepts Ovh_iploadbalancing_tcp_routeArray and Ovh_iploadbalancing_tcp_routeArrayOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_routeArrayInput` via:
//
//          Ovh_iploadbalancing_tcp_routeArray{ Ovh_iploadbalancing_tcp_routeArgs{...} }
type Ovh_iploadbalancing_tcp_routeArrayInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_routeArrayOutput() Ovh_iploadbalancing_tcp_routeArrayOutput
	ToOvh_iploadbalancing_tcp_routeArrayOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_routeArrayOutput
}

type Ovh_iploadbalancing_tcp_routeArray []Ovh_iploadbalancing_tcp_routeInput

func (Ovh_iploadbalancing_tcp_routeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_iploadbalancing_tcp_route)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_routeArray) ToOvh_iploadbalancing_tcp_routeArrayOutput() Ovh_iploadbalancing_tcp_routeArrayOutput {
	return i.ToOvh_iploadbalancing_tcp_routeArrayOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_routeArray) ToOvh_iploadbalancing_tcp_routeArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_routeArrayOutput)
}

// Ovh_iploadbalancing_tcp_routeMapInput is an input type that accepts Ovh_iploadbalancing_tcp_routeMap and Ovh_iploadbalancing_tcp_routeMapOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_routeMapInput` via:
//
//          Ovh_iploadbalancing_tcp_routeMap{ "key": Ovh_iploadbalancing_tcp_routeArgs{...} }
type Ovh_iploadbalancing_tcp_routeMapInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_routeMapOutput() Ovh_iploadbalancing_tcp_routeMapOutput
	ToOvh_iploadbalancing_tcp_routeMapOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_routeMapOutput
}

type Ovh_iploadbalancing_tcp_routeMap map[string]Ovh_iploadbalancing_tcp_routeInput

func (Ovh_iploadbalancing_tcp_routeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_iploadbalancing_tcp_route)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_routeMap) ToOvh_iploadbalancing_tcp_routeMapOutput() Ovh_iploadbalancing_tcp_routeMapOutput {
	return i.ToOvh_iploadbalancing_tcp_routeMapOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_routeMap) ToOvh_iploadbalancing_tcp_routeMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_routeMapOutput)
}

type Ovh_iploadbalancing_tcp_routeOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_routeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_route)(nil))
}

func (o Ovh_iploadbalancing_tcp_routeOutput) ToOvh_iploadbalancing_tcp_routeOutput() Ovh_iploadbalancing_tcp_routeOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeOutput) ToOvh_iploadbalancing_tcp_routeOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeOutput) ToOvh_iploadbalancing_tcp_routePtrOutput() Ovh_iploadbalancing_tcp_routePtrOutput {
	return o.ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(context.Background())
}

func (o Ovh_iploadbalancing_tcp_routeOutput) ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_iploadbalancing_tcp_route) *Ovh_iploadbalancing_tcp_route {
		return &v
	}).(Ovh_iploadbalancing_tcp_routePtrOutput)
}

type Ovh_iploadbalancing_tcp_routePtrOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_routePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_route)(nil))
}

func (o Ovh_iploadbalancing_tcp_routePtrOutput) ToOvh_iploadbalancing_tcp_routePtrOutput() Ovh_iploadbalancing_tcp_routePtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routePtrOutput) ToOvh_iploadbalancing_tcp_routePtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routePtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routePtrOutput) Elem() Ovh_iploadbalancing_tcp_routeOutput {
	return o.ApplyT(func(v *Ovh_iploadbalancing_tcp_route) Ovh_iploadbalancing_tcp_route {
		if v != nil {
			return *v
		}
		var ret Ovh_iploadbalancing_tcp_route
		return ret
	}).(Ovh_iploadbalancing_tcp_routeOutput)
}

type Ovh_iploadbalancing_tcp_routeArrayOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_routeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_iploadbalancing_tcp_route)(nil))
}

func (o Ovh_iploadbalancing_tcp_routeArrayOutput) ToOvh_iploadbalancing_tcp_routeArrayOutput() Ovh_iploadbalancing_tcp_routeArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeArrayOutput) ToOvh_iploadbalancing_tcp_routeArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeArrayOutput) Index(i pulumi.IntInput) Ovh_iploadbalancing_tcp_routeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_route {
		return vs[0].([]Ovh_iploadbalancing_tcp_route)[vs[1].(int)]
	}).(Ovh_iploadbalancing_tcp_routeOutput)
}

type Ovh_iploadbalancing_tcp_routeMapOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_routeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_iploadbalancing_tcp_route)(nil))
}

func (o Ovh_iploadbalancing_tcp_routeMapOutput) ToOvh_iploadbalancing_tcp_routeMapOutput() Ovh_iploadbalancing_tcp_routeMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeMapOutput) ToOvh_iploadbalancing_tcp_routeMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_routeMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_routeMapOutput) MapIndex(k pulumi.StringInput) Ovh_iploadbalancing_tcp_routeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_route {
		return vs[0].(map[string]Ovh_iploadbalancing_tcp_route)[vs[1].(string)]
	}).(Ovh_iploadbalancing_tcp_routeOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_routeOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_routePtrOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_routeArrayOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_routeMapOutput{})
}
