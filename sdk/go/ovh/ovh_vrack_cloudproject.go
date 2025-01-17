// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_vrack_cloudproject struct {
	pulumi.CustomResourceState

	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewOvh_vrack_cloudproject registers a new resource with the given unique name, arguments, and options.
func NewOvh_vrack_cloudproject(ctx *pulumi.Context,
	name string, args *Ovh_vrack_cloudprojectArgs, opts ...pulumi.ResourceOption) (*Ovh_vrack_cloudproject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_vrack_cloudproject
	err := ctx.RegisterResource("ovh:index/ovh_vrack_cloudproject:ovh_vrack_cloudproject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_vrack_cloudproject gets an existing Ovh_vrack_cloudproject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_vrack_cloudproject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_vrack_cloudprojectState, opts ...pulumi.ResourceOption) (*Ovh_vrack_cloudproject, error) {
	var resource Ovh_vrack_cloudproject
	err := ctx.ReadResource("ovh:index/ovh_vrack_cloudproject:ovh_vrack_cloudproject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_vrack_cloudproject resources.
type ovh_vrack_cloudprojectState struct {
	ProjectId *string `pulumi:"projectId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
}

type Ovh_vrack_cloudprojectState struct {
	ProjectId pulumi.StringPtrInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
}

func (Ovh_vrack_cloudprojectState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_vrack_cloudprojectState)(nil)).Elem()
}

type ovh_vrack_cloudprojectArgs struct {
	ProjectId string `pulumi:"projectId"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_vrack_cloudproject resource.
type Ovh_vrack_cloudprojectArgs struct {
	ProjectId pulumi.StringInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringInput
}

func (Ovh_vrack_cloudprojectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_vrack_cloudprojectArgs)(nil)).Elem()
}

type Ovh_vrack_cloudprojectInput interface {
	pulumi.Input

	ToOvh_vrack_cloudprojectOutput() Ovh_vrack_cloudprojectOutput
	ToOvh_vrack_cloudprojectOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectOutput
}

func (*Ovh_vrack_cloudproject) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_vrack_cloudproject)(nil))
}

func (i *Ovh_vrack_cloudproject) ToOvh_vrack_cloudprojectOutput() Ovh_vrack_cloudprojectOutput {
	return i.ToOvh_vrack_cloudprojectOutputWithContext(context.Background())
}

func (i *Ovh_vrack_cloudproject) ToOvh_vrack_cloudprojectOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_cloudprojectOutput)
}

func (i *Ovh_vrack_cloudproject) ToOvh_vrack_cloudprojectPtrOutput() Ovh_vrack_cloudprojectPtrOutput {
	return i.ToOvh_vrack_cloudprojectPtrOutputWithContext(context.Background())
}

func (i *Ovh_vrack_cloudproject) ToOvh_vrack_cloudprojectPtrOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_cloudprojectPtrOutput)
}

type Ovh_vrack_cloudprojectPtrInput interface {
	pulumi.Input

	ToOvh_vrack_cloudprojectPtrOutput() Ovh_vrack_cloudprojectPtrOutput
	ToOvh_vrack_cloudprojectPtrOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectPtrOutput
}

type ovh_vrack_cloudprojectPtrType Ovh_vrack_cloudprojectArgs

func (*ovh_vrack_cloudprojectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_vrack_cloudproject)(nil))
}

func (i *ovh_vrack_cloudprojectPtrType) ToOvh_vrack_cloudprojectPtrOutput() Ovh_vrack_cloudprojectPtrOutput {
	return i.ToOvh_vrack_cloudprojectPtrOutputWithContext(context.Background())
}

func (i *ovh_vrack_cloudprojectPtrType) ToOvh_vrack_cloudprojectPtrOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_cloudprojectPtrOutput)
}

// Ovh_vrack_cloudprojectArrayInput is an input type that accepts Ovh_vrack_cloudprojectArray and Ovh_vrack_cloudprojectArrayOutput values.
// You can construct a concrete instance of `Ovh_vrack_cloudprojectArrayInput` via:
//
//          Ovh_vrack_cloudprojectArray{ Ovh_vrack_cloudprojectArgs{...} }
type Ovh_vrack_cloudprojectArrayInput interface {
	pulumi.Input

	ToOvh_vrack_cloudprojectArrayOutput() Ovh_vrack_cloudprojectArrayOutput
	ToOvh_vrack_cloudprojectArrayOutputWithContext(context.Context) Ovh_vrack_cloudprojectArrayOutput
}

type Ovh_vrack_cloudprojectArray []Ovh_vrack_cloudprojectInput

func (Ovh_vrack_cloudprojectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_vrack_cloudproject)(nil)).Elem()
}

func (i Ovh_vrack_cloudprojectArray) ToOvh_vrack_cloudprojectArrayOutput() Ovh_vrack_cloudprojectArrayOutput {
	return i.ToOvh_vrack_cloudprojectArrayOutputWithContext(context.Background())
}

func (i Ovh_vrack_cloudprojectArray) ToOvh_vrack_cloudprojectArrayOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_cloudprojectArrayOutput)
}

// Ovh_vrack_cloudprojectMapInput is an input type that accepts Ovh_vrack_cloudprojectMap and Ovh_vrack_cloudprojectMapOutput values.
// You can construct a concrete instance of `Ovh_vrack_cloudprojectMapInput` via:
//
//          Ovh_vrack_cloudprojectMap{ "key": Ovh_vrack_cloudprojectArgs{...} }
type Ovh_vrack_cloudprojectMapInput interface {
	pulumi.Input

	ToOvh_vrack_cloudprojectMapOutput() Ovh_vrack_cloudprojectMapOutput
	ToOvh_vrack_cloudprojectMapOutputWithContext(context.Context) Ovh_vrack_cloudprojectMapOutput
}

type Ovh_vrack_cloudprojectMap map[string]Ovh_vrack_cloudprojectInput

func (Ovh_vrack_cloudprojectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_vrack_cloudproject)(nil)).Elem()
}

func (i Ovh_vrack_cloudprojectMap) ToOvh_vrack_cloudprojectMapOutput() Ovh_vrack_cloudprojectMapOutput {
	return i.ToOvh_vrack_cloudprojectMapOutputWithContext(context.Background())
}

func (i Ovh_vrack_cloudprojectMap) ToOvh_vrack_cloudprojectMapOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_vrack_cloudprojectMapOutput)
}

type Ovh_vrack_cloudprojectOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_cloudprojectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_vrack_cloudproject)(nil))
}

func (o Ovh_vrack_cloudprojectOutput) ToOvh_vrack_cloudprojectOutput() Ovh_vrack_cloudprojectOutput {
	return o
}

func (o Ovh_vrack_cloudprojectOutput) ToOvh_vrack_cloudprojectOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectOutput {
	return o
}

func (o Ovh_vrack_cloudprojectOutput) ToOvh_vrack_cloudprojectPtrOutput() Ovh_vrack_cloudprojectPtrOutput {
	return o.ToOvh_vrack_cloudprojectPtrOutputWithContext(context.Background())
}

func (o Ovh_vrack_cloudprojectOutput) ToOvh_vrack_cloudprojectPtrOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_vrack_cloudproject) *Ovh_vrack_cloudproject {
		return &v
	}).(Ovh_vrack_cloudprojectPtrOutput)
}

type Ovh_vrack_cloudprojectPtrOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_cloudprojectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_vrack_cloudproject)(nil))
}

func (o Ovh_vrack_cloudprojectPtrOutput) ToOvh_vrack_cloudprojectPtrOutput() Ovh_vrack_cloudprojectPtrOutput {
	return o
}

func (o Ovh_vrack_cloudprojectPtrOutput) ToOvh_vrack_cloudprojectPtrOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectPtrOutput {
	return o
}

func (o Ovh_vrack_cloudprojectPtrOutput) Elem() Ovh_vrack_cloudprojectOutput {
	return o.ApplyT(func(v *Ovh_vrack_cloudproject) Ovh_vrack_cloudproject {
		if v != nil {
			return *v
		}
		var ret Ovh_vrack_cloudproject
		return ret
	}).(Ovh_vrack_cloudprojectOutput)
}

type Ovh_vrack_cloudprojectArrayOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_cloudprojectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_vrack_cloudproject)(nil))
}

func (o Ovh_vrack_cloudprojectArrayOutput) ToOvh_vrack_cloudprojectArrayOutput() Ovh_vrack_cloudprojectArrayOutput {
	return o
}

func (o Ovh_vrack_cloudprojectArrayOutput) ToOvh_vrack_cloudprojectArrayOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectArrayOutput {
	return o
}

func (o Ovh_vrack_cloudprojectArrayOutput) Index(i pulumi.IntInput) Ovh_vrack_cloudprojectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_vrack_cloudproject {
		return vs[0].([]Ovh_vrack_cloudproject)[vs[1].(int)]
	}).(Ovh_vrack_cloudprojectOutput)
}

type Ovh_vrack_cloudprojectMapOutput struct{ *pulumi.OutputState }

func (Ovh_vrack_cloudprojectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_vrack_cloudproject)(nil))
}

func (o Ovh_vrack_cloudprojectMapOutput) ToOvh_vrack_cloudprojectMapOutput() Ovh_vrack_cloudprojectMapOutput {
	return o
}

func (o Ovh_vrack_cloudprojectMapOutput) ToOvh_vrack_cloudprojectMapOutputWithContext(ctx context.Context) Ovh_vrack_cloudprojectMapOutput {
	return o
}

func (o Ovh_vrack_cloudprojectMapOutput) MapIndex(k pulumi.StringInput) Ovh_vrack_cloudprojectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_vrack_cloudproject {
		return vs[0].(map[string]Ovh_vrack_cloudproject)[vs[1].(string)]
	}).(Ovh_vrack_cloudprojectOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_vrack_cloudprojectOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_cloudprojectPtrOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_cloudprojectArrayOutput{})
	pulumi.RegisterOutputType(Ovh_vrack_cloudprojectMapOutput{})
}
