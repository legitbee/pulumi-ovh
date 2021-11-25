// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_cloud_project_network_private struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated: use the regions_attributes field instead
	Regions           pulumi.StringArrayOutput                                     `pulumi:"regions"`
	RegionsAttributes Ovh_cloud_project_network_privateRegionsAttributeArrayOutput `pulumi:"regionsAttributes"`
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses Ovh_cloud_project_network_privateRegionsStatusArrayOutput `pulumi:"regionsStatuses"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	Status      pulumi.StringOutput `pulumi:"status"`
	Type        pulumi.StringOutput `pulumi:"type"`
	VlanId      pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewOvh_cloud_project_network_private registers a new resource with the given unique name, arguments, and options.
func NewOvh_cloud_project_network_private(ctx *pulumi.Context,
	name string, args *Ovh_cloud_project_network_privateArgs, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_network_private, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_cloud_project_network_private
	err := ctx.RegisterResource("ovh:index/ovh_cloud_project_network_private:ovh_cloud_project_network_private", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_cloud_project_network_private gets an existing Ovh_cloud_project_network_private resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_cloud_project_network_private(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_cloud_project_network_privateState, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_network_private, error) {
	var resource Ovh_cloud_project_network_private
	err := ctx.ReadResource("ovh:index/ovh_cloud_project_network_private:ovh_cloud_project_network_private", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_cloud_project_network_private resources.
type ovh_cloud_project_network_privateState struct {
	Name *string `pulumi:"name"`
	// Deprecated: use the regions_attributes field instead
	Regions           []string                                            `pulumi:"regions"`
	RegionsAttributes []Ovh_cloud_project_network_privateRegionsAttribute `pulumi:"regionsAttributes"`
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses []Ovh_cloud_project_network_privateRegionsStatus `pulumi:"regionsStatuses"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
	Status      *string `pulumi:"status"`
	Type        *string `pulumi:"type"`
	VlanId      *int    `pulumi:"vlanId"`
}

type Ovh_cloud_project_network_privateState struct {
	Name pulumi.StringPtrInput
	// Deprecated: use the regions_attributes field instead
	Regions           pulumi.StringArrayInput
	RegionsAttributes Ovh_cloud_project_network_privateRegionsAttributeArrayInput
	// Deprecated: use the regions_attributes field instead
	RegionsStatuses Ovh_cloud_project_network_privateRegionsStatusArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	Type        pulumi.StringPtrInput
	VlanId      pulumi.IntPtrInput
}

func (Ovh_cloud_project_network_privateState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_network_privateState)(nil)).Elem()
}

type ovh_cloud_project_network_privateArgs struct {
	Name *string `pulumi:"name"`
	// Deprecated: use the regions_attributes field instead
	Regions []string `pulumi:"regions"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName string `pulumi:"serviceName"`
	VlanId      *int   `pulumi:"vlanId"`
}

// The set of arguments for constructing a Ovh_cloud_project_network_private resource.
type Ovh_cloud_project_network_privateArgs struct {
	Name pulumi.StringPtrInput
	// Deprecated: use the regions_attributes field instead
	Regions pulumi.StringArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringInput
	VlanId      pulumi.IntPtrInput
}

func (Ovh_cloud_project_network_privateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_network_privateArgs)(nil)).Elem()
}

type Ovh_cloud_project_network_privateInput interface {
	pulumi.Input

	ToOvh_cloud_project_network_privateOutput() Ovh_cloud_project_network_privateOutput
	ToOvh_cloud_project_network_privateOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateOutput
}

func (*Ovh_cloud_project_network_private) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_network_private)(nil))
}

func (i *Ovh_cloud_project_network_private) ToOvh_cloud_project_network_privateOutput() Ovh_cloud_project_network_privateOutput {
	return i.ToOvh_cloud_project_network_privateOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_network_private) ToOvh_cloud_project_network_privateOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_network_privateOutput)
}

func (i *Ovh_cloud_project_network_private) ToOvh_cloud_project_network_privatePtrOutput() Ovh_cloud_project_network_privatePtrOutput {
	return i.ToOvh_cloud_project_network_privatePtrOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_network_private) ToOvh_cloud_project_network_privatePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_network_privatePtrOutput)
}

type Ovh_cloud_project_network_privatePtrInput interface {
	pulumi.Input

	ToOvh_cloud_project_network_privatePtrOutput() Ovh_cloud_project_network_privatePtrOutput
	ToOvh_cloud_project_network_privatePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privatePtrOutput
}

type ovh_cloud_project_network_privatePtrType Ovh_cloud_project_network_privateArgs

func (*ovh_cloud_project_network_privatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_network_private)(nil))
}

func (i *ovh_cloud_project_network_privatePtrType) ToOvh_cloud_project_network_privatePtrOutput() Ovh_cloud_project_network_privatePtrOutput {
	return i.ToOvh_cloud_project_network_privatePtrOutputWithContext(context.Background())
}

func (i *ovh_cloud_project_network_privatePtrType) ToOvh_cloud_project_network_privatePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_network_privatePtrOutput)
}

// Ovh_cloud_project_network_privateArrayInput is an input type that accepts Ovh_cloud_project_network_privateArray and Ovh_cloud_project_network_privateArrayOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_network_privateArrayInput` via:
//
//          Ovh_cloud_project_network_privateArray{ Ovh_cloud_project_network_privateArgs{...} }
type Ovh_cloud_project_network_privateArrayInput interface {
	pulumi.Input

	ToOvh_cloud_project_network_privateArrayOutput() Ovh_cloud_project_network_privateArrayOutput
	ToOvh_cloud_project_network_privateArrayOutputWithContext(context.Context) Ovh_cloud_project_network_privateArrayOutput
}

type Ovh_cloud_project_network_privateArray []Ovh_cloud_project_network_privateInput

func (Ovh_cloud_project_network_privateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_cloud_project_network_private)(nil)).Elem()
}

func (i Ovh_cloud_project_network_privateArray) ToOvh_cloud_project_network_privateArrayOutput() Ovh_cloud_project_network_privateArrayOutput {
	return i.ToOvh_cloud_project_network_privateArrayOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_network_privateArray) ToOvh_cloud_project_network_privateArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_network_privateArrayOutput)
}

// Ovh_cloud_project_network_privateMapInput is an input type that accepts Ovh_cloud_project_network_privateMap and Ovh_cloud_project_network_privateMapOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_network_privateMapInput` via:
//
//          Ovh_cloud_project_network_privateMap{ "key": Ovh_cloud_project_network_privateArgs{...} }
type Ovh_cloud_project_network_privateMapInput interface {
	pulumi.Input

	ToOvh_cloud_project_network_privateMapOutput() Ovh_cloud_project_network_privateMapOutput
	ToOvh_cloud_project_network_privateMapOutputWithContext(context.Context) Ovh_cloud_project_network_privateMapOutput
}

type Ovh_cloud_project_network_privateMap map[string]Ovh_cloud_project_network_privateInput

func (Ovh_cloud_project_network_privateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_cloud_project_network_private)(nil)).Elem()
}

func (i Ovh_cloud_project_network_privateMap) ToOvh_cloud_project_network_privateMapOutput() Ovh_cloud_project_network_privateMapOutput {
	return i.ToOvh_cloud_project_network_privateMapOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_network_privateMap) ToOvh_cloud_project_network_privateMapOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_network_privateMapOutput)
}

type Ovh_cloud_project_network_privateOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_network_privateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_network_private)(nil))
}

func (o Ovh_cloud_project_network_privateOutput) ToOvh_cloud_project_network_privateOutput() Ovh_cloud_project_network_privateOutput {
	return o
}

func (o Ovh_cloud_project_network_privateOutput) ToOvh_cloud_project_network_privateOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateOutput {
	return o
}

func (o Ovh_cloud_project_network_privateOutput) ToOvh_cloud_project_network_privatePtrOutput() Ovh_cloud_project_network_privatePtrOutput {
	return o.ToOvh_cloud_project_network_privatePtrOutputWithContext(context.Background())
}

func (o Ovh_cloud_project_network_privateOutput) ToOvh_cloud_project_network_privatePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_cloud_project_network_private) *Ovh_cloud_project_network_private {
		return &v
	}).(Ovh_cloud_project_network_privatePtrOutput)
}

type Ovh_cloud_project_network_privatePtrOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_network_privatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_network_private)(nil))
}

func (o Ovh_cloud_project_network_privatePtrOutput) ToOvh_cloud_project_network_privatePtrOutput() Ovh_cloud_project_network_privatePtrOutput {
	return o
}

func (o Ovh_cloud_project_network_privatePtrOutput) ToOvh_cloud_project_network_privatePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privatePtrOutput {
	return o
}

func (o Ovh_cloud_project_network_privatePtrOutput) Elem() Ovh_cloud_project_network_privateOutput {
	return o.ApplyT(func(v *Ovh_cloud_project_network_private) Ovh_cloud_project_network_private {
		if v != nil {
			return *v
		}
		var ret Ovh_cloud_project_network_private
		return ret
	}).(Ovh_cloud_project_network_privateOutput)
}

type Ovh_cloud_project_network_privateArrayOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_network_privateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_cloud_project_network_private)(nil))
}

func (o Ovh_cloud_project_network_privateArrayOutput) ToOvh_cloud_project_network_privateArrayOutput() Ovh_cloud_project_network_privateArrayOutput {
	return o
}

func (o Ovh_cloud_project_network_privateArrayOutput) ToOvh_cloud_project_network_privateArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateArrayOutput {
	return o
}

func (o Ovh_cloud_project_network_privateArrayOutput) Index(i pulumi.IntInput) Ovh_cloud_project_network_privateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_cloud_project_network_private {
		return vs[0].([]Ovh_cloud_project_network_private)[vs[1].(int)]
	}).(Ovh_cloud_project_network_privateOutput)
}

type Ovh_cloud_project_network_privateMapOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_network_privateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_cloud_project_network_private)(nil))
}

func (o Ovh_cloud_project_network_privateMapOutput) ToOvh_cloud_project_network_privateMapOutput() Ovh_cloud_project_network_privateMapOutput {
	return o
}

func (o Ovh_cloud_project_network_privateMapOutput) ToOvh_cloud_project_network_privateMapOutputWithContext(ctx context.Context) Ovh_cloud_project_network_privateMapOutput {
	return o
}

func (o Ovh_cloud_project_network_privateMapOutput) MapIndex(k pulumi.StringInput) Ovh_cloud_project_network_privateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_cloud_project_network_private {
		return vs[0].(map[string]Ovh_cloud_project_network_private)[vs[1].(string)]
	}).(Ovh_cloud_project_network_privateOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_cloud_project_network_privateOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_network_privatePtrOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_network_privateArrayOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_network_privateMapOutput{})
}
