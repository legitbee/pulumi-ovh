// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_iploadbalancing_tcp_farm struct {
	pulumi.CustomResourceState

	Balance        pulumi.StringPtrOutput                     `pulumi:"balance"`
	DisplayName    pulumi.StringPtrOutput                     `pulumi:"displayName"`
	Port           pulumi.IntPtrOutput                        `pulumi:"port"`
	Probe          Ovh_iploadbalancing_tcp_farmProbePtrOutput `pulumi:"probe"`
	ServiceName    pulumi.StringOutput                        `pulumi:"serviceName"`
	Stickiness     pulumi.StringPtrOutput                     `pulumi:"stickiness"`
	VrackNetworkId pulumi.IntPtrOutput                        `pulumi:"vrackNetworkId"`
	Zone           pulumi.StringOutput                        `pulumi:"zone"`
}

// NewOvh_iploadbalancing_tcp_farm registers a new resource with the given unique name, arguments, and options.
func NewOvh_iploadbalancing_tcp_farm(ctx *pulumi.Context,
	name string, args *Ovh_iploadbalancing_tcp_farmArgs, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_farm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Ovh_iploadbalancing_tcp_farm
	err := ctx.RegisterResource("ovh:index/ovh_iploadbalancing_tcp_farm:ovh_iploadbalancing_tcp_farm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_iploadbalancing_tcp_farm gets an existing Ovh_iploadbalancing_tcp_farm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_iploadbalancing_tcp_farm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_iploadbalancing_tcp_farmState, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_farm, error) {
	var resource Ovh_iploadbalancing_tcp_farm
	err := ctx.ReadResource("ovh:index/ovh_iploadbalancing_tcp_farm:ovh_iploadbalancing_tcp_farm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_iploadbalancing_tcp_farm resources.
type ovh_iploadbalancing_tcp_farmState struct {
	Balance        *string                            `pulumi:"balance"`
	DisplayName    *string                            `pulumi:"displayName"`
	Port           *int                               `pulumi:"port"`
	Probe          *Ovh_iploadbalancing_tcp_farmProbe `pulumi:"probe"`
	ServiceName    *string                            `pulumi:"serviceName"`
	Stickiness     *string                            `pulumi:"stickiness"`
	VrackNetworkId *int                               `pulumi:"vrackNetworkId"`
	Zone           *string                            `pulumi:"zone"`
}

type Ovh_iploadbalancing_tcp_farmState struct {
	Balance        pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.IntPtrInput
	Probe          Ovh_iploadbalancing_tcp_farmProbePtrInput
	ServiceName    pulumi.StringPtrInput
	Stickiness     pulumi.StringPtrInput
	VrackNetworkId pulumi.IntPtrInput
	Zone           pulumi.StringPtrInput
}

func (Ovh_iploadbalancing_tcp_farmState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_farmState)(nil)).Elem()
}

type ovh_iploadbalancing_tcp_farmArgs struct {
	Balance        *string                            `pulumi:"balance"`
	DisplayName    *string                            `pulumi:"displayName"`
	Port           *int                               `pulumi:"port"`
	Probe          *Ovh_iploadbalancing_tcp_farmProbe `pulumi:"probe"`
	ServiceName    string                             `pulumi:"serviceName"`
	Stickiness     *string                            `pulumi:"stickiness"`
	VrackNetworkId *int                               `pulumi:"vrackNetworkId"`
	Zone           string                             `pulumi:"zone"`
}

// The set of arguments for constructing a Ovh_iploadbalancing_tcp_farm resource.
type Ovh_iploadbalancing_tcp_farmArgs struct {
	Balance        pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.IntPtrInput
	Probe          Ovh_iploadbalancing_tcp_farmProbePtrInput
	ServiceName    pulumi.StringInput
	Stickiness     pulumi.StringPtrInput
	VrackNetworkId pulumi.IntPtrInput
	Zone           pulumi.StringInput
}

func (Ovh_iploadbalancing_tcp_farmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_farmArgs)(nil)).Elem()
}

type Ovh_iploadbalancing_tcp_farmInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farmOutput() Ovh_iploadbalancing_tcp_farmOutput
	ToOvh_iploadbalancing_tcp_farmOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmOutput
}

func (*Ovh_iploadbalancing_tcp_farm) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_farm)(nil))
}

func (i *Ovh_iploadbalancing_tcp_farm) ToOvh_iploadbalancing_tcp_farmOutput() Ovh_iploadbalancing_tcp_farmOutput {
	return i.ToOvh_iploadbalancing_tcp_farmOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_farm) ToOvh_iploadbalancing_tcp_farmOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farmOutput)
}

func (i *Ovh_iploadbalancing_tcp_farm) ToOvh_iploadbalancing_tcp_farmPtrOutput() Ovh_iploadbalancing_tcp_farmPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_farm) ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farmPtrOutput)
}

type Ovh_iploadbalancing_tcp_farmPtrInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farmPtrOutput() Ovh_iploadbalancing_tcp_farmPtrOutput
	ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmPtrOutput
}

type ovh_iploadbalancing_tcp_farmPtrType Ovh_iploadbalancing_tcp_farmArgs

func (*ovh_iploadbalancing_tcp_farmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_farm)(nil))
}

func (i *ovh_iploadbalancing_tcp_farmPtrType) ToOvh_iploadbalancing_tcp_farmPtrOutput() Ovh_iploadbalancing_tcp_farmPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(context.Background())
}

func (i *ovh_iploadbalancing_tcp_farmPtrType) ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farmPtrOutput)
}

// Ovh_iploadbalancing_tcp_farmArrayInput is an input type that accepts Ovh_iploadbalancing_tcp_farmArray and Ovh_iploadbalancing_tcp_farmArrayOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_farmArrayInput` via:
//
//          Ovh_iploadbalancing_tcp_farmArray{ Ovh_iploadbalancing_tcp_farmArgs{...} }
type Ovh_iploadbalancing_tcp_farmArrayInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farmArrayOutput() Ovh_iploadbalancing_tcp_farmArrayOutput
	ToOvh_iploadbalancing_tcp_farmArrayOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_farmArrayOutput
}

type Ovh_iploadbalancing_tcp_farmArray []Ovh_iploadbalancing_tcp_farmInput

func (Ovh_iploadbalancing_tcp_farmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_iploadbalancing_tcp_farm)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_farmArray) ToOvh_iploadbalancing_tcp_farmArrayOutput() Ovh_iploadbalancing_tcp_farmArrayOutput {
	return i.ToOvh_iploadbalancing_tcp_farmArrayOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_farmArray) ToOvh_iploadbalancing_tcp_farmArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farmArrayOutput)
}

// Ovh_iploadbalancing_tcp_farmMapInput is an input type that accepts Ovh_iploadbalancing_tcp_farmMap and Ovh_iploadbalancing_tcp_farmMapOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_farmMapInput` via:
//
//          Ovh_iploadbalancing_tcp_farmMap{ "key": Ovh_iploadbalancing_tcp_farmArgs{...} }
type Ovh_iploadbalancing_tcp_farmMapInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farmMapOutput() Ovh_iploadbalancing_tcp_farmMapOutput
	ToOvh_iploadbalancing_tcp_farmMapOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_farmMapOutput
}

type Ovh_iploadbalancing_tcp_farmMap map[string]Ovh_iploadbalancing_tcp_farmInput

func (Ovh_iploadbalancing_tcp_farmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_iploadbalancing_tcp_farm)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_farmMap) ToOvh_iploadbalancing_tcp_farmMapOutput() Ovh_iploadbalancing_tcp_farmMapOutput {
	return i.ToOvh_iploadbalancing_tcp_farmMapOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_farmMap) ToOvh_iploadbalancing_tcp_farmMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farmMapOutput)
}

type Ovh_iploadbalancing_tcp_farmOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_farm)(nil))
}

func (o Ovh_iploadbalancing_tcp_farmOutput) ToOvh_iploadbalancing_tcp_farmOutput() Ovh_iploadbalancing_tcp_farmOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmOutput) ToOvh_iploadbalancing_tcp_farmOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmOutput) ToOvh_iploadbalancing_tcp_farmPtrOutput() Ovh_iploadbalancing_tcp_farmPtrOutput {
	return o.ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(context.Background())
}

func (o Ovh_iploadbalancing_tcp_farmOutput) ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_iploadbalancing_tcp_farm) *Ovh_iploadbalancing_tcp_farm {
		return &v
	}).(Ovh_iploadbalancing_tcp_farmPtrOutput)
}

type Ovh_iploadbalancing_tcp_farmPtrOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_farm)(nil))
}

func (o Ovh_iploadbalancing_tcp_farmPtrOutput) ToOvh_iploadbalancing_tcp_farmPtrOutput() Ovh_iploadbalancing_tcp_farmPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmPtrOutput) ToOvh_iploadbalancing_tcp_farmPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmPtrOutput) Elem() Ovh_iploadbalancing_tcp_farmOutput {
	return o.ApplyT(func(v *Ovh_iploadbalancing_tcp_farm) Ovh_iploadbalancing_tcp_farm {
		if v != nil {
			return *v
		}
		var ret Ovh_iploadbalancing_tcp_farm
		return ret
	}).(Ovh_iploadbalancing_tcp_farmOutput)
}

type Ovh_iploadbalancing_tcp_farmArrayOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_iploadbalancing_tcp_farm)(nil))
}

func (o Ovh_iploadbalancing_tcp_farmArrayOutput) ToOvh_iploadbalancing_tcp_farmArrayOutput() Ovh_iploadbalancing_tcp_farmArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmArrayOutput) ToOvh_iploadbalancing_tcp_farmArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmArrayOutput) Index(i pulumi.IntInput) Ovh_iploadbalancing_tcp_farmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_farm {
		return vs[0].([]Ovh_iploadbalancing_tcp_farm)[vs[1].(int)]
	}).(Ovh_iploadbalancing_tcp_farmOutput)
}

type Ovh_iploadbalancing_tcp_farmMapOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_iploadbalancing_tcp_farm)(nil))
}

func (o Ovh_iploadbalancing_tcp_farmMapOutput) ToOvh_iploadbalancing_tcp_farmMapOutput() Ovh_iploadbalancing_tcp_farmMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmMapOutput) ToOvh_iploadbalancing_tcp_farmMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farmMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farmMapOutput) MapIndex(k pulumi.StringInput) Ovh_iploadbalancing_tcp_farmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_farm {
		return vs[0].(map[string]Ovh_iploadbalancing_tcp_farm)[vs[1].(string)]
	}).(Ovh_iploadbalancing_tcp_farmOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farmOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farmPtrOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farmArrayOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farmMapOutput{})
}
