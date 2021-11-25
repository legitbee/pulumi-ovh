// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VrackDedicatedServerInterface struct {
	pulumi.CustomResourceState

	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewVrackDedicatedServerInterface registers a new resource with the given unique name, arguments, and options.
func NewVrackDedicatedServerInterface(ctx *pulumi.Context,
	name string, args *VrackDedicatedServerInterfaceArgs, opts ...pulumi.ResourceOption) (*VrackDedicatedServerInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource VrackDedicatedServerInterface
	err := ctx.RegisterResource("ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackDedicatedServerInterface gets an existing VrackDedicatedServerInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackDedicatedServerInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackDedicatedServerInterfaceState, opts ...pulumi.ResourceOption) (*VrackDedicatedServerInterface, error) {
	var resource VrackDedicatedServerInterface
	err := ctx.ReadResource("ovh:index/vrackDedicatedServerInterface:VrackDedicatedServerInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackDedicatedServerInterface resources.
type vrackDedicatedServerInterfaceState struct {
	InterfaceId *string `pulumi:"interfaceId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
}

type VrackDedicatedServerInterfaceState struct {
	InterfaceId pulumi.StringPtrInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
}

func (VrackDedicatedServerInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerInterfaceState)(nil)).Elem()
}

type vrackDedicatedServerInterfaceArgs struct {
	InterfaceId string `pulumi:"interfaceId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VrackDedicatedServerInterface resource.
type VrackDedicatedServerInterfaceArgs struct {
	InterfaceId pulumi.StringInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringInput
}

func (VrackDedicatedServerInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackDedicatedServerInterfaceArgs)(nil)).Elem()
}

type VrackDedicatedServerInterfaceInput interface {
	pulumi.Input

	ToVrackDedicatedServerInterfaceOutput() VrackDedicatedServerInterfaceOutput
	ToVrackDedicatedServerInterfaceOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceOutput
}

func (*VrackDedicatedServerInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*VrackDedicatedServerInterface)(nil))
}

func (i *VrackDedicatedServerInterface) ToVrackDedicatedServerInterfaceOutput() VrackDedicatedServerInterfaceOutput {
	return i.ToVrackDedicatedServerInterfaceOutputWithContext(context.Background())
}

func (i *VrackDedicatedServerInterface) ToVrackDedicatedServerInterfaceOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerInterfaceOutput)
}

func (i *VrackDedicatedServerInterface) ToVrackDedicatedServerInterfacePtrOutput() VrackDedicatedServerInterfacePtrOutput {
	return i.ToVrackDedicatedServerInterfacePtrOutputWithContext(context.Background())
}

func (i *VrackDedicatedServerInterface) ToVrackDedicatedServerInterfacePtrOutputWithContext(ctx context.Context) VrackDedicatedServerInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerInterfacePtrOutput)
}

type VrackDedicatedServerInterfacePtrInput interface {
	pulumi.Input

	ToVrackDedicatedServerInterfacePtrOutput() VrackDedicatedServerInterfacePtrOutput
	ToVrackDedicatedServerInterfacePtrOutputWithContext(ctx context.Context) VrackDedicatedServerInterfacePtrOutput
}

type vrackDedicatedServerInterfacePtrType VrackDedicatedServerInterfaceArgs

func (*vrackDedicatedServerInterfacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedServerInterface)(nil))
}

func (i *vrackDedicatedServerInterfacePtrType) ToVrackDedicatedServerInterfacePtrOutput() VrackDedicatedServerInterfacePtrOutput {
	return i.ToVrackDedicatedServerInterfacePtrOutputWithContext(context.Background())
}

func (i *vrackDedicatedServerInterfacePtrType) ToVrackDedicatedServerInterfacePtrOutputWithContext(ctx context.Context) VrackDedicatedServerInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerInterfacePtrOutput)
}

// VrackDedicatedServerInterfaceArrayInput is an input type that accepts VrackDedicatedServerInterfaceArray and VrackDedicatedServerInterfaceArrayOutput values.
// You can construct a concrete instance of `VrackDedicatedServerInterfaceArrayInput` via:
//
//          VrackDedicatedServerInterfaceArray{ VrackDedicatedServerInterfaceArgs{...} }
type VrackDedicatedServerInterfaceArrayInput interface {
	pulumi.Input

	ToVrackDedicatedServerInterfaceArrayOutput() VrackDedicatedServerInterfaceArrayOutput
	ToVrackDedicatedServerInterfaceArrayOutputWithContext(context.Context) VrackDedicatedServerInterfaceArrayOutput
}

type VrackDedicatedServerInterfaceArray []VrackDedicatedServerInterfaceInput

func (VrackDedicatedServerInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackDedicatedServerInterface)(nil)).Elem()
}

func (i VrackDedicatedServerInterfaceArray) ToVrackDedicatedServerInterfaceArrayOutput() VrackDedicatedServerInterfaceArrayOutput {
	return i.ToVrackDedicatedServerInterfaceArrayOutputWithContext(context.Background())
}

func (i VrackDedicatedServerInterfaceArray) ToVrackDedicatedServerInterfaceArrayOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerInterfaceArrayOutput)
}

// VrackDedicatedServerInterfaceMapInput is an input type that accepts VrackDedicatedServerInterfaceMap and VrackDedicatedServerInterfaceMapOutput values.
// You can construct a concrete instance of `VrackDedicatedServerInterfaceMapInput` via:
//
//          VrackDedicatedServerInterfaceMap{ "key": VrackDedicatedServerInterfaceArgs{...} }
type VrackDedicatedServerInterfaceMapInput interface {
	pulumi.Input

	ToVrackDedicatedServerInterfaceMapOutput() VrackDedicatedServerInterfaceMapOutput
	ToVrackDedicatedServerInterfaceMapOutputWithContext(context.Context) VrackDedicatedServerInterfaceMapOutput
}

type VrackDedicatedServerInterfaceMap map[string]VrackDedicatedServerInterfaceInput

func (VrackDedicatedServerInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackDedicatedServerInterface)(nil)).Elem()
}

func (i VrackDedicatedServerInterfaceMap) ToVrackDedicatedServerInterfaceMapOutput() VrackDedicatedServerInterfaceMapOutput {
	return i.ToVrackDedicatedServerInterfaceMapOutputWithContext(context.Background())
}

func (i VrackDedicatedServerInterfaceMap) ToVrackDedicatedServerInterfaceMapOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackDedicatedServerInterfaceMapOutput)
}

type VrackDedicatedServerInterfaceOutput struct{ *pulumi.OutputState }

func (VrackDedicatedServerInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VrackDedicatedServerInterface)(nil))
}

func (o VrackDedicatedServerInterfaceOutput) ToVrackDedicatedServerInterfaceOutput() VrackDedicatedServerInterfaceOutput {
	return o
}

func (o VrackDedicatedServerInterfaceOutput) ToVrackDedicatedServerInterfaceOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceOutput {
	return o
}

func (o VrackDedicatedServerInterfaceOutput) ToVrackDedicatedServerInterfacePtrOutput() VrackDedicatedServerInterfacePtrOutput {
	return o.ToVrackDedicatedServerInterfacePtrOutputWithContext(context.Background())
}

func (o VrackDedicatedServerInterfaceOutput) ToVrackDedicatedServerInterfacePtrOutputWithContext(ctx context.Context) VrackDedicatedServerInterfacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VrackDedicatedServerInterface) *VrackDedicatedServerInterface {
		return &v
	}).(VrackDedicatedServerInterfacePtrOutput)
}

type VrackDedicatedServerInterfacePtrOutput struct{ *pulumi.OutputState }

func (VrackDedicatedServerInterfacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackDedicatedServerInterface)(nil))
}

func (o VrackDedicatedServerInterfacePtrOutput) ToVrackDedicatedServerInterfacePtrOutput() VrackDedicatedServerInterfacePtrOutput {
	return o
}

func (o VrackDedicatedServerInterfacePtrOutput) ToVrackDedicatedServerInterfacePtrOutputWithContext(ctx context.Context) VrackDedicatedServerInterfacePtrOutput {
	return o
}

func (o VrackDedicatedServerInterfacePtrOutput) Elem() VrackDedicatedServerInterfaceOutput {
	return o.ApplyT(func(v *VrackDedicatedServerInterface) VrackDedicatedServerInterface {
		if v != nil {
			return *v
		}
		var ret VrackDedicatedServerInterface
		return ret
	}).(VrackDedicatedServerInterfaceOutput)
}

type VrackDedicatedServerInterfaceArrayOutput struct{ *pulumi.OutputState }

func (VrackDedicatedServerInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VrackDedicatedServerInterface)(nil))
}

func (o VrackDedicatedServerInterfaceArrayOutput) ToVrackDedicatedServerInterfaceArrayOutput() VrackDedicatedServerInterfaceArrayOutput {
	return o
}

func (o VrackDedicatedServerInterfaceArrayOutput) ToVrackDedicatedServerInterfaceArrayOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceArrayOutput {
	return o
}

func (o VrackDedicatedServerInterfaceArrayOutput) Index(i pulumi.IntInput) VrackDedicatedServerInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VrackDedicatedServerInterface {
		return vs[0].([]VrackDedicatedServerInterface)[vs[1].(int)]
	}).(VrackDedicatedServerInterfaceOutput)
}

type VrackDedicatedServerInterfaceMapOutput struct{ *pulumi.OutputState }

func (VrackDedicatedServerInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VrackDedicatedServerInterface)(nil))
}

func (o VrackDedicatedServerInterfaceMapOutput) ToVrackDedicatedServerInterfaceMapOutput() VrackDedicatedServerInterfaceMapOutput {
	return o
}

func (o VrackDedicatedServerInterfaceMapOutput) ToVrackDedicatedServerInterfaceMapOutputWithContext(ctx context.Context) VrackDedicatedServerInterfaceMapOutput {
	return o
}

func (o VrackDedicatedServerInterfaceMapOutput) MapIndex(k pulumi.StringInput) VrackDedicatedServerInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VrackDedicatedServerInterface {
		return vs[0].(map[string]VrackDedicatedServerInterface)[vs[1].(string)]
	}).(VrackDedicatedServerInterfaceOutput)
}

func init() {
	pulumi.RegisterOutputType(VrackDedicatedServerInterfaceOutput{})
	pulumi.RegisterOutputType(VrackDedicatedServerInterfacePtrOutput{})
	pulumi.RegisterOutputType(VrackDedicatedServerInterfaceArrayOutput{})
	pulumi.RegisterOutputType(VrackDedicatedServerInterfaceMapOutput{})
}
