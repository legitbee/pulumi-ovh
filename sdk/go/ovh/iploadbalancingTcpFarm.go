// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IploadbalancingTcpFarm struct {
	pulumi.CustomResourceState

	Balance        pulumi.StringPtrOutput               `pulumi:"balance"`
	DisplayName    pulumi.StringPtrOutput               `pulumi:"displayName"`
	Port           pulumi.IntPtrOutput                  `pulumi:"port"`
	Probe          IploadbalancingTcpFarmProbePtrOutput `pulumi:"probe"`
	ServiceName    pulumi.StringOutput                  `pulumi:"serviceName"`
	Stickiness     pulumi.StringPtrOutput               `pulumi:"stickiness"`
	VrackNetworkId pulumi.IntPtrOutput                  `pulumi:"vrackNetworkId"`
	Zone           pulumi.StringOutput                  `pulumi:"zone"`
}

// NewIploadbalancingTcpFarm registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancingTcpFarm(ctx *pulumi.Context,
	name string, args *IploadbalancingTcpFarmArgs, opts ...pulumi.ResourceOption) (*IploadbalancingTcpFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource IploadbalancingTcpFarm
	err := ctx.RegisterResource("ovh:index/iploadbalancingTcpFarm:IploadbalancingTcpFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancingTcpFarm gets an existing IploadbalancingTcpFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancingTcpFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingTcpFarmState, opts ...pulumi.ResourceOption) (*IploadbalancingTcpFarm, error) {
	var resource IploadbalancingTcpFarm
	err := ctx.ReadResource("ovh:index/iploadbalancingTcpFarm:IploadbalancingTcpFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IploadbalancingTcpFarm resources.
type iploadbalancingTcpFarmState struct {
	Balance        *string                      `pulumi:"balance"`
	DisplayName    *string                      `pulumi:"displayName"`
	Port           *int                         `pulumi:"port"`
	Probe          *IploadbalancingTcpFarmProbe `pulumi:"probe"`
	ServiceName    *string                      `pulumi:"serviceName"`
	Stickiness     *string                      `pulumi:"stickiness"`
	VrackNetworkId *int                         `pulumi:"vrackNetworkId"`
	Zone           *string                      `pulumi:"zone"`
}

type IploadbalancingTcpFarmState struct {
	Balance        pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.IntPtrInput
	Probe          IploadbalancingTcpFarmProbePtrInput
	ServiceName    pulumi.StringPtrInput
	Stickiness     pulumi.StringPtrInput
	VrackNetworkId pulumi.IntPtrInput
	Zone           pulumi.StringPtrInput
}

func (IploadbalancingTcpFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpFarmState)(nil)).Elem()
}

type iploadbalancingTcpFarmArgs struct {
	Balance        *string                      `pulumi:"balance"`
	DisplayName    *string                      `pulumi:"displayName"`
	Port           *int                         `pulumi:"port"`
	Probe          *IploadbalancingTcpFarmProbe `pulumi:"probe"`
	ServiceName    string                       `pulumi:"serviceName"`
	Stickiness     *string                      `pulumi:"stickiness"`
	VrackNetworkId *int                         `pulumi:"vrackNetworkId"`
	Zone           string                       `pulumi:"zone"`
}

// The set of arguments for constructing a IploadbalancingTcpFarm resource.
type IploadbalancingTcpFarmArgs struct {
	Balance        pulumi.StringPtrInput
	DisplayName    pulumi.StringPtrInput
	Port           pulumi.IntPtrInput
	Probe          IploadbalancingTcpFarmProbePtrInput
	ServiceName    pulumi.StringInput
	Stickiness     pulumi.StringPtrInput
	VrackNetworkId pulumi.IntPtrInput
	Zone           pulumi.StringInput
}

func (IploadbalancingTcpFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingTcpFarmArgs)(nil)).Elem()
}

type IploadbalancingTcpFarmInput interface {
	pulumi.Input

	ToIploadbalancingTcpFarmOutput() IploadbalancingTcpFarmOutput
	ToIploadbalancingTcpFarmOutputWithContext(ctx context.Context) IploadbalancingTcpFarmOutput
}

func (*IploadbalancingTcpFarm) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingTcpFarm)(nil))
}

func (i *IploadbalancingTcpFarm) ToIploadbalancingTcpFarmOutput() IploadbalancingTcpFarmOutput {
	return i.ToIploadbalancingTcpFarmOutputWithContext(context.Background())
}

func (i *IploadbalancingTcpFarm) ToIploadbalancingTcpFarmOutputWithContext(ctx context.Context) IploadbalancingTcpFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFarmOutput)
}

func (i *IploadbalancingTcpFarm) ToIploadbalancingTcpFarmPtrOutput() IploadbalancingTcpFarmPtrOutput {
	return i.ToIploadbalancingTcpFarmPtrOutputWithContext(context.Background())
}

func (i *IploadbalancingTcpFarm) ToIploadbalancingTcpFarmPtrOutputWithContext(ctx context.Context) IploadbalancingTcpFarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFarmPtrOutput)
}

type IploadbalancingTcpFarmPtrInput interface {
	pulumi.Input

	ToIploadbalancingTcpFarmPtrOutput() IploadbalancingTcpFarmPtrOutput
	ToIploadbalancingTcpFarmPtrOutputWithContext(ctx context.Context) IploadbalancingTcpFarmPtrOutput
}

type iploadbalancingTcpFarmPtrType IploadbalancingTcpFarmArgs

func (*iploadbalancingTcpFarmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpFarm)(nil))
}

func (i *iploadbalancingTcpFarmPtrType) ToIploadbalancingTcpFarmPtrOutput() IploadbalancingTcpFarmPtrOutput {
	return i.ToIploadbalancingTcpFarmPtrOutputWithContext(context.Background())
}

func (i *iploadbalancingTcpFarmPtrType) ToIploadbalancingTcpFarmPtrOutputWithContext(ctx context.Context) IploadbalancingTcpFarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFarmPtrOutput)
}

// IploadbalancingTcpFarmArrayInput is an input type that accepts IploadbalancingTcpFarmArray and IploadbalancingTcpFarmArrayOutput values.
// You can construct a concrete instance of `IploadbalancingTcpFarmArrayInput` via:
//
//          IploadbalancingTcpFarmArray{ IploadbalancingTcpFarmArgs{...} }
type IploadbalancingTcpFarmArrayInput interface {
	pulumi.Input

	ToIploadbalancingTcpFarmArrayOutput() IploadbalancingTcpFarmArrayOutput
	ToIploadbalancingTcpFarmArrayOutputWithContext(context.Context) IploadbalancingTcpFarmArrayOutput
}

type IploadbalancingTcpFarmArray []IploadbalancingTcpFarmInput

func (IploadbalancingTcpFarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IploadbalancingTcpFarm)(nil)).Elem()
}

func (i IploadbalancingTcpFarmArray) ToIploadbalancingTcpFarmArrayOutput() IploadbalancingTcpFarmArrayOutput {
	return i.ToIploadbalancingTcpFarmArrayOutputWithContext(context.Background())
}

func (i IploadbalancingTcpFarmArray) ToIploadbalancingTcpFarmArrayOutputWithContext(ctx context.Context) IploadbalancingTcpFarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFarmArrayOutput)
}

// IploadbalancingTcpFarmMapInput is an input type that accepts IploadbalancingTcpFarmMap and IploadbalancingTcpFarmMapOutput values.
// You can construct a concrete instance of `IploadbalancingTcpFarmMapInput` via:
//
//          IploadbalancingTcpFarmMap{ "key": IploadbalancingTcpFarmArgs{...} }
type IploadbalancingTcpFarmMapInput interface {
	pulumi.Input

	ToIploadbalancingTcpFarmMapOutput() IploadbalancingTcpFarmMapOutput
	ToIploadbalancingTcpFarmMapOutputWithContext(context.Context) IploadbalancingTcpFarmMapOutput
}

type IploadbalancingTcpFarmMap map[string]IploadbalancingTcpFarmInput

func (IploadbalancingTcpFarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IploadbalancingTcpFarm)(nil)).Elem()
}

func (i IploadbalancingTcpFarmMap) ToIploadbalancingTcpFarmMapOutput() IploadbalancingTcpFarmMapOutput {
	return i.ToIploadbalancingTcpFarmMapOutputWithContext(context.Background())
}

func (i IploadbalancingTcpFarmMap) ToIploadbalancingTcpFarmMapOutputWithContext(ctx context.Context) IploadbalancingTcpFarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingTcpFarmMapOutput)
}

type IploadbalancingTcpFarmOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IploadbalancingTcpFarm)(nil))
}

func (o IploadbalancingTcpFarmOutput) ToIploadbalancingTcpFarmOutput() IploadbalancingTcpFarmOutput {
	return o
}

func (o IploadbalancingTcpFarmOutput) ToIploadbalancingTcpFarmOutputWithContext(ctx context.Context) IploadbalancingTcpFarmOutput {
	return o
}

func (o IploadbalancingTcpFarmOutput) ToIploadbalancingTcpFarmPtrOutput() IploadbalancingTcpFarmPtrOutput {
	return o.ToIploadbalancingTcpFarmPtrOutputWithContext(context.Background())
}

func (o IploadbalancingTcpFarmOutput) ToIploadbalancingTcpFarmPtrOutputWithContext(ctx context.Context) IploadbalancingTcpFarmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IploadbalancingTcpFarm) *IploadbalancingTcpFarm {
		return &v
	}).(IploadbalancingTcpFarmPtrOutput)
}

type IploadbalancingTcpFarmPtrOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpFarmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IploadbalancingTcpFarm)(nil))
}

func (o IploadbalancingTcpFarmPtrOutput) ToIploadbalancingTcpFarmPtrOutput() IploadbalancingTcpFarmPtrOutput {
	return o
}

func (o IploadbalancingTcpFarmPtrOutput) ToIploadbalancingTcpFarmPtrOutputWithContext(ctx context.Context) IploadbalancingTcpFarmPtrOutput {
	return o
}

func (o IploadbalancingTcpFarmPtrOutput) Elem() IploadbalancingTcpFarmOutput {
	return o.ApplyT(func(v *IploadbalancingTcpFarm) IploadbalancingTcpFarm {
		if v != nil {
			return *v
		}
		var ret IploadbalancingTcpFarm
		return ret
	}).(IploadbalancingTcpFarmOutput)
}

type IploadbalancingTcpFarmArrayOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpFarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IploadbalancingTcpFarm)(nil))
}

func (o IploadbalancingTcpFarmArrayOutput) ToIploadbalancingTcpFarmArrayOutput() IploadbalancingTcpFarmArrayOutput {
	return o
}

func (o IploadbalancingTcpFarmArrayOutput) ToIploadbalancingTcpFarmArrayOutputWithContext(ctx context.Context) IploadbalancingTcpFarmArrayOutput {
	return o
}

func (o IploadbalancingTcpFarmArrayOutput) Index(i pulumi.IntInput) IploadbalancingTcpFarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IploadbalancingTcpFarm {
		return vs[0].([]IploadbalancingTcpFarm)[vs[1].(int)]
	}).(IploadbalancingTcpFarmOutput)
}

type IploadbalancingTcpFarmMapOutput struct{ *pulumi.OutputState }

func (IploadbalancingTcpFarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IploadbalancingTcpFarm)(nil))
}

func (o IploadbalancingTcpFarmMapOutput) ToIploadbalancingTcpFarmMapOutput() IploadbalancingTcpFarmMapOutput {
	return o
}

func (o IploadbalancingTcpFarmMapOutput) ToIploadbalancingTcpFarmMapOutputWithContext(ctx context.Context) IploadbalancingTcpFarmMapOutput {
	return o
}

func (o IploadbalancingTcpFarmMapOutput) MapIndex(k pulumi.StringInput) IploadbalancingTcpFarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IploadbalancingTcpFarm {
		return vs[0].(map[string]IploadbalancingTcpFarm)[vs[1].(string)]
	}).(IploadbalancingTcpFarmOutput)
}

func init() {
	pulumi.RegisterOutputType(IploadbalancingTcpFarmOutput{})
	pulumi.RegisterOutputType(IploadbalancingTcpFarmPtrOutput{})
	pulumi.RegisterOutputType(IploadbalancingTcpFarmArrayOutput{})
	pulumi.RegisterOutputType(IploadbalancingTcpFarmMapOutput{})
}
