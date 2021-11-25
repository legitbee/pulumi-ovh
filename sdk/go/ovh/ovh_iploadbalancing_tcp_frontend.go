// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_iploadbalancing_tcp_frontend struct {
	pulumi.CustomResourceState

	AllowedSources pulumi.StringArrayOutput `pulumi:"allowedSources"`
	DedicatedIpfos pulumi.StringArrayOutput `pulumi:"dedicatedIpfos"`
	DefaultFarmId  pulumi.IntOutput         `pulumi:"defaultFarmId"`
	DefaultSslId   pulumi.IntOutput         `pulumi:"defaultSslId"`
	Disabled       pulumi.BoolPtrOutput     `pulumi:"disabled"`
	DisplayName    pulumi.StringPtrOutput   `pulumi:"displayName"`
	Port           pulumi.StringOutput      `pulumi:"port"`
	ServiceName    pulumi.StringOutput      `pulumi:"serviceName"`
	Ssl            pulumi.BoolPtrOutput     `pulumi:"ssl"`
	Zone           pulumi.StringOutput      `pulumi:"zone"`
}

// NewOvh_iploadbalancing_tcp_frontend registers a new resource with the given unique name, arguments, and options.
func NewOvh_iploadbalancing_tcp_frontend(ctx *pulumi.Context,
	name string, args *Ovh_iploadbalancing_tcp_frontendArgs, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_frontend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Ovh_iploadbalancing_tcp_frontend
	err := ctx.RegisterResource("ovh:index/ovh_iploadbalancing_tcp_frontend:ovh_iploadbalancing_tcp_frontend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_iploadbalancing_tcp_frontend gets an existing Ovh_iploadbalancing_tcp_frontend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_iploadbalancing_tcp_frontend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_iploadbalancing_tcp_frontendState, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_frontend, error) {
	var resource Ovh_iploadbalancing_tcp_frontend
	err := ctx.ReadResource("ovh:index/ovh_iploadbalancing_tcp_frontend:ovh_iploadbalancing_tcp_frontend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_iploadbalancing_tcp_frontend resources.
type ovh_iploadbalancing_tcp_frontendState struct {
	AllowedSources []string `pulumi:"allowedSources"`
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	DefaultFarmId  *int     `pulumi:"defaultFarmId"`
	DefaultSslId   *int     `pulumi:"defaultSslId"`
	Disabled       *bool    `pulumi:"disabled"`
	DisplayName    *string  `pulumi:"displayName"`
	Port           *string  `pulumi:"port"`
	ServiceName    *string  `pulumi:"serviceName"`
	Ssl            *bool    `pulumi:"ssl"`
	Zone           *string  `pulumi:"zone"`
}

type Ovh_iploadbalancing_tcp_frontendState struct {
	AllowedSources pulumi.StringArrayInput
	DedicatedIpfos pulumi.StringArrayInput
	DefaultFarmId  pulumi.IntPtrInput
	DefaultSslId   pulumi.IntPtrInput
	Disabled       pulumi.BoolPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.StringPtrInput
	ServiceName    pulumi.StringPtrInput
	Ssl            pulumi.BoolPtrInput
	Zone           pulumi.StringPtrInput
}

func (Ovh_iploadbalancing_tcp_frontendState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_frontendState)(nil)).Elem()
}

type ovh_iploadbalancing_tcp_frontendArgs struct {
	AllowedSources []string `pulumi:"allowedSources"`
	DedicatedIpfos []string `pulumi:"dedicatedIpfos"`
	DefaultFarmId  *int     `pulumi:"defaultFarmId"`
	DefaultSslId   *int     `pulumi:"defaultSslId"`
	Disabled       *bool    `pulumi:"disabled"`
	DisplayName    *string  `pulumi:"displayName"`
	Port           string   `pulumi:"port"`
	ServiceName    string   `pulumi:"serviceName"`
	Ssl            *bool    `pulumi:"ssl"`
	Zone           string   `pulumi:"zone"`
}

// The set of arguments for constructing a Ovh_iploadbalancing_tcp_frontend resource.
type Ovh_iploadbalancing_tcp_frontendArgs struct {
	AllowedSources pulumi.StringArrayInput
	DedicatedIpfos pulumi.StringArrayInput
	DefaultFarmId  pulumi.IntPtrInput
	DefaultSslId   pulumi.IntPtrInput
	Disabled       pulumi.BoolPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.StringInput
	ServiceName    pulumi.StringInput
	Ssl            pulumi.BoolPtrInput
	Zone           pulumi.StringInput
}

func (Ovh_iploadbalancing_tcp_frontendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_frontendArgs)(nil)).Elem()
}

type Ovh_iploadbalancing_tcp_frontendInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_frontendOutput() Ovh_iploadbalancing_tcp_frontendOutput
	ToOvh_iploadbalancing_tcp_frontendOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendOutput
}

func (*Ovh_iploadbalancing_tcp_frontend) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (i *Ovh_iploadbalancing_tcp_frontend) ToOvh_iploadbalancing_tcp_frontendOutput() Ovh_iploadbalancing_tcp_frontendOutput {
	return i.ToOvh_iploadbalancing_tcp_frontendOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_frontend) ToOvh_iploadbalancing_tcp_frontendOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_frontendOutput)
}

func (i *Ovh_iploadbalancing_tcp_frontend) ToOvh_iploadbalancing_tcp_frontendPtrOutput() Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_frontend) ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_frontendPtrOutput)
}

type Ovh_iploadbalancing_tcp_frontendPtrInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_frontendPtrOutput() Ovh_iploadbalancing_tcp_frontendPtrOutput
	ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendPtrOutput
}

type ovh_iploadbalancing_tcp_frontendPtrType Ovh_iploadbalancing_tcp_frontendArgs

func (*ovh_iploadbalancing_tcp_frontendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (i *ovh_iploadbalancing_tcp_frontendPtrType) ToOvh_iploadbalancing_tcp_frontendPtrOutput() Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(context.Background())
}

func (i *ovh_iploadbalancing_tcp_frontendPtrType) ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_frontendPtrOutput)
}

// Ovh_iploadbalancing_tcp_frontendArrayInput is an input type that accepts Ovh_iploadbalancing_tcp_frontendArray and Ovh_iploadbalancing_tcp_frontendArrayOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_frontendArrayInput` via:
//
//          Ovh_iploadbalancing_tcp_frontendArray{ Ovh_iploadbalancing_tcp_frontendArgs{...} }
type Ovh_iploadbalancing_tcp_frontendArrayInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_frontendArrayOutput() Ovh_iploadbalancing_tcp_frontendArrayOutput
	ToOvh_iploadbalancing_tcp_frontendArrayOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_frontendArrayOutput
}

type Ovh_iploadbalancing_tcp_frontendArray []Ovh_iploadbalancing_tcp_frontendInput

func (Ovh_iploadbalancing_tcp_frontendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_iploadbalancing_tcp_frontend)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_frontendArray) ToOvh_iploadbalancing_tcp_frontendArrayOutput() Ovh_iploadbalancing_tcp_frontendArrayOutput {
	return i.ToOvh_iploadbalancing_tcp_frontendArrayOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_frontendArray) ToOvh_iploadbalancing_tcp_frontendArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_frontendArrayOutput)
}

// Ovh_iploadbalancing_tcp_frontendMapInput is an input type that accepts Ovh_iploadbalancing_tcp_frontendMap and Ovh_iploadbalancing_tcp_frontendMapOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_frontendMapInput` via:
//
//          Ovh_iploadbalancing_tcp_frontendMap{ "key": Ovh_iploadbalancing_tcp_frontendArgs{...} }
type Ovh_iploadbalancing_tcp_frontendMapInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_frontendMapOutput() Ovh_iploadbalancing_tcp_frontendMapOutput
	ToOvh_iploadbalancing_tcp_frontendMapOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_frontendMapOutput
}

type Ovh_iploadbalancing_tcp_frontendMap map[string]Ovh_iploadbalancing_tcp_frontendInput

func (Ovh_iploadbalancing_tcp_frontendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_iploadbalancing_tcp_frontend)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_frontendMap) ToOvh_iploadbalancing_tcp_frontendMapOutput() Ovh_iploadbalancing_tcp_frontendMapOutput {
	return i.ToOvh_iploadbalancing_tcp_frontendMapOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_frontendMap) ToOvh_iploadbalancing_tcp_frontendMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_frontendMapOutput)
}

type Ovh_iploadbalancing_tcp_frontendOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_frontendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (o Ovh_iploadbalancing_tcp_frontendOutput) ToOvh_iploadbalancing_tcp_frontendOutput() Ovh_iploadbalancing_tcp_frontendOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendOutput) ToOvh_iploadbalancing_tcp_frontendOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendOutput) ToOvh_iploadbalancing_tcp_frontendPtrOutput() Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return o.ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(context.Background())
}

func (o Ovh_iploadbalancing_tcp_frontendOutput) ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_iploadbalancing_tcp_frontend) *Ovh_iploadbalancing_tcp_frontend {
		return &v
	}).(Ovh_iploadbalancing_tcp_frontendPtrOutput)
}

type Ovh_iploadbalancing_tcp_frontendPtrOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_frontendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (o Ovh_iploadbalancing_tcp_frontendPtrOutput) ToOvh_iploadbalancing_tcp_frontendPtrOutput() Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendPtrOutput) ToOvh_iploadbalancing_tcp_frontendPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendPtrOutput) Elem() Ovh_iploadbalancing_tcp_frontendOutput {
	return o.ApplyT(func(v *Ovh_iploadbalancing_tcp_frontend) Ovh_iploadbalancing_tcp_frontend {
		if v != nil {
			return *v
		}
		var ret Ovh_iploadbalancing_tcp_frontend
		return ret
	}).(Ovh_iploadbalancing_tcp_frontendOutput)
}

type Ovh_iploadbalancing_tcp_frontendArrayOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_frontendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (o Ovh_iploadbalancing_tcp_frontendArrayOutput) ToOvh_iploadbalancing_tcp_frontendArrayOutput() Ovh_iploadbalancing_tcp_frontendArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendArrayOutput) ToOvh_iploadbalancing_tcp_frontendArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendArrayOutput) Index(i pulumi.IntInput) Ovh_iploadbalancing_tcp_frontendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_frontend {
		return vs[0].([]Ovh_iploadbalancing_tcp_frontend)[vs[1].(int)]
	}).(Ovh_iploadbalancing_tcp_frontendOutput)
}

type Ovh_iploadbalancing_tcp_frontendMapOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_frontendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_iploadbalancing_tcp_frontend)(nil))
}

func (o Ovh_iploadbalancing_tcp_frontendMapOutput) ToOvh_iploadbalancing_tcp_frontendMapOutput() Ovh_iploadbalancing_tcp_frontendMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendMapOutput) ToOvh_iploadbalancing_tcp_frontendMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_frontendMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_frontendMapOutput) MapIndex(k pulumi.StringInput) Ovh_iploadbalancing_tcp_frontendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_frontend {
		return vs[0].(map[string]Ovh_iploadbalancing_tcp_frontend)[vs[1].(string)]
	}).(Ovh_iploadbalancing_tcp_frontendOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_frontendOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_frontendPtrOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_frontendArrayOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_frontendMapOutput{})
}
