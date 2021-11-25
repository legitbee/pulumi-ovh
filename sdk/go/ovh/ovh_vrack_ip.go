// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_vrack_ip struct {
	pulumi.CustomResourceState

	// Your IP block.
	Block pulumi.StringOutput `pulumi:"block"`
	// Your gateway
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Your IP block
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewOvh_vrack_ip registers a new resource with the given unique name, arguments, and options.
func NewOvh_vrack_ip(ctx *pulumi.Context,
	name string, args *Ovh_vrack_ipArgs, opts ...pulumi.ResourceOption) (*Ovh_vrack_ip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Block == nil {
		return nil, errors.New("invalid value for required argument 'Block'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_vrack_ip
	err := ctx.RegisterResource("ovh:index/ovh_vrack_ip:ovh_vrack_ip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_vrack_ip gets an existing Ovh_vrack_ip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_vrack_ip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_vrack_ipState, opts ...pulumi.ResourceOption) (*Ovh_vrack_ip, error) {
	var resource Ovh_vrack_ip
	err := ctx.ReadResource("ovh:index/ovh_vrack_ip:ovh_vrack_ip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_vrack_ip resources.
type ovh_vrack_ipState struct {
	// Your IP block.
	Block *string `pulumi:"block"`
	// Your gateway
	Gateway *string `pulumi:"gateway"`
	// Your IP block
	Ip *string `pulumi:"ip"`
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone *string `pulumi:"zone"`
}

type Ovh_vrack_ipState struct {
	// Your IP block.
	Block pulumi.StringPtrInput
	// Your gateway
	Gateway pulumi.StringPtrInput
	// Your IP block
	Ip pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
	// Where you want your block announced on the network
	Zone pulumi.StringPtrInput
}

func (Ovh_vrack_ipState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_vrack_ipState)(nil)).Elem()
}

type ovh_vrack_ipArgs struct {
	// Your IP block.
	Block string `pulumi:"block"`
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_vrack_ip resource.
type Ovh_vrack_ipArgs struct {
	// Your IP block.
	Block pulumi.StringInput
	// The internal name of your vrack
	ServiceName pulumi.StringInput
}

func (Ovh_vrack_ipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_vrack_ipArgs)(nil)).Elem()
}

type Ovh_vrack_ipInput interface {
	pulumi.Input

	ToOvh_vrack_ipOutput() Ovh_vrack_ipOutput
	ToOvh_vrack_ipOutputWithContext(ctx context.Context) Ovh_vrack_ipOutput
}

func (*Ovh_vrack_ip) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_vrack_ip)(nil))
}

func (i *Ovh_vrack_ip) ToOvh_vrack_ipOutput() Ovh_vrack_ipOutput {
	return i.ToOvh_vrack_ipOutputWithContext(context.Background())
}

func (i *Ovh_vrack_ip) ToOvh_vrack_ipOutputWithContext(ctx context.Context) Ovh_vrack_ipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_ipOutput)
}

func (i *Ovh_vrack_ip) ToOvh_vrack_ipPtrOutput() Ovh_vrack_ipPtrOutput {
	return i.ToOvh_vrack_ipPtrOutputWithContext(context.Background())
}

func (i *Ovh_vrack_ip) ToOvh_vrack_ipPtrOutputWithContext(ctx context.Context) Ovh_vrack_ipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_ipPtrOutput)
}

type Ovh_vrack_ipPtrInput interface {
	pulumi.Input

	ToOvh_vrack_ipPtrOutput() Ovh_vrack_ipPtrOutput
	ToOvh_vrack_ipPtrOutputWithContext(ctx context.Context) Ovh_vrack_ipPtrOutput
}

type ovh_vrack_ipPtrType Ovh_vrack_ipArgs

func (*ovh_vrack_ipPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_vrack_ip)(nil))
}

func (i *ovh_vrack_ipPtrType) ToOvh_vrack_ipPtrOutput() Ovh_vrack_ipPtrOutput {
	return i.ToOvh_vrack_ipPtrOutputWithContext(context.Background())
}

func (i *ovh_vrack_ipPtrType) ToOvh_vrack_ipPtrOutputWithContext(ctx context.Context) Ovh_vrack_ipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_ipPtrOutput)
}

// Ovh_vrack_ipArrayInput is an input type that accepts Ovh_vrack_ipArray and Ovh_vrack_ipArrayOutput values.
// You can construct a concrete instance of `Ovh_vrack_ipArrayInput` via:
//
//          Ovh_vrack_ipArray{ Ovh_vrack_ipArgs{...} }
type Ovh_vrack_ipArrayInput interface {
	pulumi.Input

	ToOvh_vrack_ipArrayOutput() Ovh_vrack_ipArrayOutput
	ToOvh_vrack_ipArrayOutputWithContext(context.Context) Ovh_vrack_ipArrayOutput
}

type Ovh_vrack_ipArray []Ovh_vrack_ipInput

func (Ovh_vrack_ipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_vrack_ip)(nil)).Elem()
}

func (i Ovh_vrack_ipArray) ToOvh_vrack_ipArrayOutput() Ovh_vrack_ipArrayOutput {
	return i.ToOvh_vrack_ipArrayOutputWithContext(context.Background())
}

func (i Ovh_vrack_ipArray) ToOvh_vrack_ipArrayOutputWithContext(ctx context.Context) Ovh_vrack_ipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_ipArrayOutput)
}

// Ovh_vrack_ipMapInput is an input type that accepts Ovh_vrack_ipMap and Ovh_vrack_ipMapOutput values.
// You can construct a concrete instance of `Ovh_vrack_ipMapInput` via:
//
//          Ovh_vrack_ipMap{ "key": Ovh_vrack_ipArgs{...} }
type Ovh_vrack_ipMapInput interface {
	pulumi.Input

	ToOvh_vrack_ipMapOutput() Ovh_vrack_ipMapOutput
	ToOvh_vrack_ipMapOutputWithContext(context.Context) Ovh_vrack_ipMapOutput
}

type Ovh_vrack_ipMap map[string]Ovh_vrack_ipInput

func (Ovh_vrack_ipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_vrack_ip)(nil)).Elem()
}

func (i Ovh_vrack_ipMap) ToOvh_vrack_ipMapOutput() Ovh_vrack_ipMapOutput {
	return i.ToOvh_vrack_ipMapOutputWithContext(context.Background())
}

func (i Ovh_vrack_ipMap) ToOvh_vrack_ipMapOutputWithContext(ctx context.Context) Ovh_vrack_ipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_ipMapOutput)
}

type Ovh_vrack_ipOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_ipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_vrack_ip)(nil))
}

func (o Ovh_vrack_ipOutput) ToOvh_vrack_ipOutput() Ovh_vrack_ipOutput {
	return o
}

func (o Ovh_vrack_ipOutput) ToOvh_vrack_ipOutputWithContext(ctx context.Context) Ovh_vrack_ipOutput {
	return o
}

func (o Ovh_vrack_ipOutput) ToOvh_vrack_ipPtrOutput() Ovh_vrack_ipPtrOutput {
	return o.ToOvh_vrack_ipPtrOutputWithContext(context.Background())
}

func (o Ovh_vrack_ipOutput) ToOvh_vrack_ipPtrOutputWithContext(ctx context.Context) Ovh_vrack_ipPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_vrack_ip) *Ovh_vrack_ip {
		return &v
	}).(Ovh_vrack_ipPtrOutput)
}

type Ovh_vrack_ipPtrOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_ipPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_vrack_ip)(nil))
}

func (o Ovh_vrack_ipPtrOutput) ToOvh_vrack_ipPtrOutput() Ovh_vrack_ipPtrOutput {
	return o
}

func (o Ovh_vrack_ipPtrOutput) ToOvh_vrack_ipPtrOutputWithContext(ctx context.Context) Ovh_vrack_ipPtrOutput {
	return o
}

func (o Ovh_vrack_ipPtrOutput) Elem() Ovh_vrack_ipOutput {
	return o.ApplyT(func(v *Ovh_vrack_ip) Ovh_vrack_ip {
		if v != nil {
			return *v
		}
		var ret Ovh_vrack_ip
		return ret
	}).(Ovh_vrack_ipOutput)
}

type Ovh_vrack_ipArrayOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_ipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_vrack_ip)(nil))
}

func (o Ovh_vrack_ipArrayOutput) ToOvh_vrack_ipArrayOutput() Ovh_vrack_ipArrayOutput {
	return o
}

func (o Ovh_vrack_ipArrayOutput) ToOvh_vrack_ipArrayOutputWithContext(ctx context.Context) Ovh_vrack_ipArrayOutput {
	return o
}

func (o Ovh_vrack_ipArrayOutput) Index(i pulumi.IntInput) Ovh_vrack_ipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_vrack_ip {
		return vs[0].([]Ovh_vrack_ip)[vs[1].(int)]
	}).(Ovh_vrack_ipOutput)
}

type Ovh_vrack_ipMapOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_ipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_vrack_ip)(nil))
}

func (o Ovh_vrack_ipMapOutput) ToOvh_vrack_ipMapOutput() Ovh_vrack_ipMapOutput {
	return o
}

func (o Ovh_vrack_ipMapOutput) ToOvh_vrack_ipMapOutputWithContext(ctx context.Context) Ovh_vrack_ipMapOutput {
	return o
}

func (o Ovh_vrack_ipMapOutput) MapIndex(k pulumi.StringInput) Ovh_vrack_ipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_vrack_ip {
		return vs[0].(map[string]Ovh_vrack_ip)[vs[1].(string)]
	}).(Ovh_vrack_ipOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_vrack_ipOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_ipPtrOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_ipArrayOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_ipMapOutput{})
}
