// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_iploadbalancing_tcp_farm_server struct {
	pulumi.CustomResourceState

	Address              pulumi.StringOutput    `pulumi:"address"`
	Backup               pulumi.BoolPtrOutput   `pulumi:"backup"`
	Chain                pulumi.StringPtrOutput `pulumi:"chain"`
	DisplayName          pulumi.StringPtrOutput `pulumi:"displayName"`
	FarmId               pulumi.IntOutput       `pulumi:"farmId"`
	Port                 pulumi.IntPtrOutput    `pulumi:"port"`
	Probe                pulumi.BoolPtrOutput   `pulumi:"probe"`
	ProxyProtocolVersion pulumi.StringPtrOutput `pulumi:"proxyProtocolVersion"`
	ServiceName          pulumi.StringOutput    `pulumi:"serviceName"`
	Ssl                  pulumi.BoolPtrOutput   `pulumi:"ssl"`
	Status               pulumi.StringOutput    `pulumi:"status"`
	Weight               pulumi.IntPtrOutput    `pulumi:"weight"`
}

// NewOvh_iploadbalancing_tcp_farm_server registers a new resource with the given unique name, arguments, and options.
func NewOvh_iploadbalancing_tcp_farm_server(ctx *pulumi.Context,
	name string, args *Ovh_iploadbalancing_tcp_farm_serverArgs, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_farm_server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.FarmId == nil {
		return nil, errors.New("invalid value for required argument 'FarmId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource Ovh_iploadbalancing_tcp_farm_server
	err := ctx.RegisterResource("ovh:index/ovh_iploadbalancing_tcp_farm_server:ovh_iploadbalancing_tcp_farm_server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_iploadbalancing_tcp_farm_server gets an existing Ovh_iploadbalancing_tcp_farm_server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_iploadbalancing_tcp_farm_server(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_iploadbalancing_tcp_farm_serverState, opts ...pulumi.ResourceOption) (*Ovh_iploadbalancing_tcp_farm_server, error) {
	var resource Ovh_iploadbalancing_tcp_farm_server
	err := ctx.ReadResource("ovh:index/ovh_iploadbalancing_tcp_farm_server:ovh_iploadbalancing_tcp_farm_server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_iploadbalancing_tcp_farm_server resources.
type ovh_iploadbalancing_tcp_farm_serverState struct {
	Address              *string `pulumi:"address"`
	Backup               *bool   `pulumi:"backup"`
	Chain                *string `pulumi:"chain"`
	DisplayName          *string `pulumi:"displayName"`
	FarmId               *int    `pulumi:"farmId"`
	Port                 *int    `pulumi:"port"`
	Probe                *bool   `pulumi:"probe"`
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	ServiceName          *string `pulumi:"serviceName"`
	Ssl                  *bool   `pulumi:"ssl"`
	Status               *string `pulumi:"status"`
	Weight               *int    `pulumi:"weight"`
}

type Ovh_iploadbalancing_tcp_farm_serverState struct {
	Address              pulumi.StringPtrInput
	Backup               pulumi.BoolPtrInput
	Chain                pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	FarmId               pulumi.IntPtrInput
	Port                 pulumi.IntPtrInput
	Probe                pulumi.BoolPtrInput
	ProxyProtocolVersion pulumi.StringPtrInput
	ServiceName          pulumi.StringPtrInput
	Ssl                  pulumi.BoolPtrInput
	Status               pulumi.StringPtrInput
	Weight               pulumi.IntPtrInput
}

func (Ovh_iploadbalancing_tcp_farm_serverState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_farm_serverState)(nil)).Elem()
}

type ovh_iploadbalancing_tcp_farm_serverArgs struct {
	Address              string  `pulumi:"address"`
	Backup               *bool   `pulumi:"backup"`
	Chain                *string `pulumi:"chain"`
	DisplayName          *string `pulumi:"displayName"`
	FarmId               int     `pulumi:"farmId"`
	Port                 *int    `pulumi:"port"`
	Probe                *bool   `pulumi:"probe"`
	ProxyProtocolVersion *string `pulumi:"proxyProtocolVersion"`
	ServiceName          string  `pulumi:"serviceName"`
	Ssl                  *bool   `pulumi:"ssl"`
	Status               string  `pulumi:"status"`
	Weight               *int    `pulumi:"weight"`
}

// The set of arguments for constructing a Ovh_iploadbalancing_tcp_farm_server resource.
type Ovh_iploadbalancing_tcp_farm_serverArgs struct {
	Address              pulumi.StringInput
	Backup               pulumi.BoolPtrInput
	Chain                pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	FarmId               pulumi.IntInput
	Port                 pulumi.IntPtrInput
	Probe                pulumi.BoolPtrInput
	ProxyProtocolVersion pulumi.StringPtrInput
	ServiceName          pulumi.StringInput
	Ssl                  pulumi.BoolPtrInput
	Status               pulumi.StringInput
	Weight               pulumi.IntPtrInput
}

func (Ovh_iploadbalancing_tcp_farm_serverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_iploadbalancing_tcp_farm_serverArgs)(nil)).Elem()
}

type Ovh_iploadbalancing_tcp_farm_serverInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farm_serverOutput() Ovh_iploadbalancing_tcp_farm_serverOutput
	ToOvh_iploadbalancing_tcp_farm_serverOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverOutput
}

func (*Ovh_iploadbalancing_tcp_farm_server) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (i *Ovh_iploadbalancing_tcp_farm_server) ToOvh_iploadbalancing_tcp_farm_serverOutput() Ovh_iploadbalancing_tcp_farm_serverOutput {
	return i.ToOvh_iploadbalancing_tcp_farm_serverOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_farm_server) ToOvh_iploadbalancing_tcp_farm_serverOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farm_serverOutput)
}

func (i *Ovh_iploadbalancing_tcp_farm_server) ToOvh_iploadbalancing_tcp_farm_serverPtrOutput() Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(context.Background())
}

func (i *Ovh_iploadbalancing_tcp_farm_server) ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farm_serverPtrOutput)
}

type Ovh_iploadbalancing_tcp_farm_serverPtrInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farm_serverPtrOutput() Ovh_iploadbalancing_tcp_farm_serverPtrOutput
	ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverPtrOutput
}

type ovh_iploadbalancing_tcp_farm_serverPtrType Ovh_iploadbalancing_tcp_farm_serverArgs

func (*ovh_iploadbalancing_tcp_farm_serverPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (i *ovh_iploadbalancing_tcp_farm_serverPtrType) ToOvh_iploadbalancing_tcp_farm_serverPtrOutput() Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return i.ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(context.Background())
}

func (i *ovh_iploadbalancing_tcp_farm_serverPtrType) ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farm_serverPtrOutput)
}

// Ovh_iploadbalancing_tcp_farm_serverArrayInput is an input type that accepts Ovh_iploadbalancing_tcp_farm_serverArray and Ovh_iploadbalancing_tcp_farm_serverArrayOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_farm_serverArrayInput` via:
//
//          Ovh_iploadbalancing_tcp_farm_serverArray{ Ovh_iploadbalancing_tcp_farm_serverArgs{...} }
type Ovh_iploadbalancing_tcp_farm_serverArrayInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farm_serverArrayOutput() Ovh_iploadbalancing_tcp_farm_serverArrayOutput
	ToOvh_iploadbalancing_tcp_farm_serverArrayOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_farm_serverArrayOutput
}

type Ovh_iploadbalancing_tcp_farm_serverArray []Ovh_iploadbalancing_tcp_farm_serverInput

func (Ovh_iploadbalancing_tcp_farm_serverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_iploadbalancing_tcp_farm_server)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_farm_serverArray) ToOvh_iploadbalancing_tcp_farm_serverArrayOutput() Ovh_iploadbalancing_tcp_farm_serverArrayOutput {
	return i.ToOvh_iploadbalancing_tcp_farm_serverArrayOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_farm_serverArray) ToOvh_iploadbalancing_tcp_farm_serverArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farm_serverArrayOutput)
}

// Ovh_iploadbalancing_tcp_farm_serverMapInput is an input type that accepts Ovh_iploadbalancing_tcp_farm_serverMap and Ovh_iploadbalancing_tcp_farm_serverMapOutput values.
// You can construct a concrete instance of `Ovh_iploadbalancing_tcp_farm_serverMapInput` via:
//
//          Ovh_iploadbalancing_tcp_farm_serverMap{ "key": Ovh_iploadbalancing_tcp_farm_serverArgs{...} }
type Ovh_iploadbalancing_tcp_farm_serverMapInput interface {
	pulumi.Input

	ToOvh_iploadbalancing_tcp_farm_serverMapOutput() Ovh_iploadbalancing_tcp_farm_serverMapOutput
	ToOvh_iploadbalancing_tcp_farm_serverMapOutputWithContext(context.Context) Ovh_iploadbalancing_tcp_farm_serverMapOutput
}

type Ovh_iploadbalancing_tcp_farm_serverMap map[string]Ovh_iploadbalancing_tcp_farm_serverInput

func (Ovh_iploadbalancing_tcp_farm_serverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_iploadbalancing_tcp_farm_server)(nil)).Elem()
}

func (i Ovh_iploadbalancing_tcp_farm_serverMap) ToOvh_iploadbalancing_tcp_farm_serverMapOutput() Ovh_iploadbalancing_tcp_farm_serverMapOutput {
	return i.ToOvh_iploadbalancing_tcp_farm_serverMapOutputWithContext(context.Background())
}

func (i Ovh_iploadbalancing_tcp_farm_serverMap) ToOvh_iploadbalancing_tcp_farm_serverMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_iploadbalancing_tcp_farm_serverMapOutput)
}

type Ovh_iploadbalancing_tcp_farm_serverOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farm_serverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (o Ovh_iploadbalancing_tcp_farm_serverOutput) ToOvh_iploadbalancing_tcp_farm_serverOutput() Ovh_iploadbalancing_tcp_farm_serverOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverOutput) ToOvh_iploadbalancing_tcp_farm_serverOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverOutput) ToOvh_iploadbalancing_tcp_farm_serverPtrOutput() Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return o.ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(context.Background())
}

func (o Ovh_iploadbalancing_tcp_farm_serverOutput) ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_iploadbalancing_tcp_farm_server) *Ovh_iploadbalancing_tcp_farm_server {
		return &v
	}).(Ovh_iploadbalancing_tcp_farm_serverPtrOutput)
}

type Ovh_iploadbalancing_tcp_farm_serverPtrOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farm_serverPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (o Ovh_iploadbalancing_tcp_farm_serverPtrOutput) ToOvh_iploadbalancing_tcp_farm_serverPtrOutput() Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverPtrOutput) ToOvh_iploadbalancing_tcp_farm_serverPtrOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverPtrOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverPtrOutput) Elem() Ovh_iploadbalancing_tcp_farm_serverOutput {
	return o.ApplyT(func(v *Ovh_iploadbalancing_tcp_farm_server) Ovh_iploadbalancing_tcp_farm_server {
		if v != nil {
			return *v
		}
		var ret Ovh_iploadbalancing_tcp_farm_server
		return ret
	}).(Ovh_iploadbalancing_tcp_farm_serverOutput)
}

type Ovh_iploadbalancing_tcp_farm_serverArrayOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farm_serverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (o Ovh_iploadbalancing_tcp_farm_serverArrayOutput) ToOvh_iploadbalancing_tcp_farm_serverArrayOutput() Ovh_iploadbalancing_tcp_farm_serverArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverArrayOutput) ToOvh_iploadbalancing_tcp_farm_serverArrayOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverArrayOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverArrayOutput) Index(i pulumi.IntInput) Ovh_iploadbalancing_tcp_farm_serverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_farm_server {
		return vs[0].([]Ovh_iploadbalancing_tcp_farm_server)[vs[1].(int)]
	}).(Ovh_iploadbalancing_tcp_farm_serverOutput)
}

type Ovh_iploadbalancing_tcp_farm_serverMapOutput struct{ *pulumi.OutputState }

func (Ovh_iploadbalancing_tcp_farm_serverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_iploadbalancing_tcp_farm_server)(nil))
}

func (o Ovh_iploadbalancing_tcp_farm_serverMapOutput) ToOvh_iploadbalancing_tcp_farm_serverMapOutput() Ovh_iploadbalancing_tcp_farm_serverMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverMapOutput) ToOvh_iploadbalancing_tcp_farm_serverMapOutputWithContext(ctx context.Context) Ovh_iploadbalancing_tcp_farm_serverMapOutput {
	return o
}

func (o Ovh_iploadbalancing_tcp_farm_serverMapOutput) MapIndex(k pulumi.StringInput) Ovh_iploadbalancing_tcp_farm_serverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_iploadbalancing_tcp_farm_server {
		return vs[0].(map[string]Ovh_iploadbalancing_tcp_farm_server)[vs[1].(string)]
	}).(Ovh_iploadbalancing_tcp_farm_serverOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farm_serverOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farm_serverPtrOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farm_serverArrayOutput{})
	pulumi.RegisterOutputType(Ovh_iploadbalancing_tcp_farm_serverMapOutput{})
}
