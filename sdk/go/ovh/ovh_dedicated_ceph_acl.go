// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_dedicated_ceph_acl struct {
	pulumi.CustomResourceState

	Family      pulumi.StringOutput `pulumi:"family"`
	Netmask     pulumi.StringOutput `pulumi:"netmask"`
	Network     pulumi.StringOutput `pulumi:"network"`
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewOvh_dedicated_ceph_acl registers a new resource with the given unique name, arguments, and options.
func NewOvh_dedicated_ceph_acl(ctx *pulumi.Context,
	name string, args *Ovh_dedicated_ceph_aclArgs, opts ...pulumi.ResourceOption) (*Ovh_dedicated_ceph_acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Netmask == nil {
		return nil, errors.New("invalid value for required argument 'Netmask'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_dedicated_ceph_acl
	err := ctx.RegisterResource("ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_dedicated_ceph_acl gets an existing Ovh_dedicated_ceph_acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_dedicated_ceph_acl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_dedicated_ceph_aclState, opts ...pulumi.ResourceOption) (*Ovh_dedicated_ceph_acl, error) {
	var resource Ovh_dedicated_ceph_acl
	err := ctx.ReadResource("ovh:index/ovh_dedicated_ceph_acl:ovh_dedicated_ceph_acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_dedicated_ceph_acl resources.
type ovh_dedicated_ceph_aclState struct {
	Family      *string `pulumi:"family"`
	Netmask     *string `pulumi:"netmask"`
	Network     *string `pulumi:"network"`
	ServiceName *string `pulumi:"serviceName"`
}

type Ovh_dedicated_ceph_aclState struct {
	Family      pulumi.StringPtrInput
	Netmask     pulumi.StringPtrInput
	Network     pulumi.StringPtrInput
	ServiceName pulumi.StringPtrInput
}

func (Ovh_dedicated_ceph_aclState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_dedicated_ceph_aclState)(nil)).Elem()
}

type ovh_dedicated_ceph_aclArgs struct {
	Netmask     string `pulumi:"netmask"`
	Network     string `pulumi:"network"`
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_dedicated_ceph_acl resource.
type Ovh_dedicated_ceph_aclArgs struct {
	Netmask     pulumi.StringInput
	Network     pulumi.StringInput
	ServiceName pulumi.StringInput
}

func (Ovh_dedicated_ceph_aclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_dedicated_ceph_aclArgs)(nil)).Elem()
}

type Ovh_dedicated_ceph_aclInput interface {
	pulumi.Input

	ToOvh_dedicated_ceph_aclOutput() Ovh_dedicated_ceph_aclOutput
	ToOvh_dedicated_ceph_aclOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclOutput
}

func (*Ovh_dedicated_ceph_acl) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_dedicated_ceph_acl)(nil))
}

func (i *Ovh_dedicated_ceph_acl) ToOvh_dedicated_ceph_aclOutput() Ovh_dedicated_ceph_aclOutput {
	return i.ToOvh_dedicated_ceph_aclOutputWithContext(context.Background())
}

func (i *Ovh_dedicated_ceph_acl) ToOvh_dedicated_ceph_aclOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_ceph_aclOutput)
}

func (i *Ovh_dedicated_ceph_acl) ToOvh_dedicated_ceph_aclPtrOutput() Ovh_dedicated_ceph_aclPtrOutput {
	return i.ToOvh_dedicated_ceph_aclPtrOutputWithContext(context.Background())
}

func (i *Ovh_dedicated_ceph_acl) ToOvh_dedicated_ceph_aclPtrOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_ceph_aclPtrOutput)
}

type Ovh_dedicated_ceph_aclPtrInput interface {
	pulumi.Input

	ToOvh_dedicated_ceph_aclPtrOutput() Ovh_dedicated_ceph_aclPtrOutput
	ToOvh_dedicated_ceph_aclPtrOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclPtrOutput
}

type ovh_dedicated_ceph_aclPtrType Ovh_dedicated_ceph_aclArgs

func (*ovh_dedicated_ceph_aclPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_dedicated_ceph_acl)(nil))
}

func (i *ovh_dedicated_ceph_aclPtrType) ToOvh_dedicated_ceph_aclPtrOutput() Ovh_dedicated_ceph_aclPtrOutput {
	return i.ToOvh_dedicated_ceph_aclPtrOutputWithContext(context.Background())
}

func (i *ovh_dedicated_ceph_aclPtrType) ToOvh_dedicated_ceph_aclPtrOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_ceph_aclPtrOutput)
}

// Ovh_dedicated_ceph_aclArrayInput is an input type that accepts Ovh_dedicated_ceph_aclArray and Ovh_dedicated_ceph_aclArrayOutput values.
// You can construct a concrete instance of `Ovh_dedicated_ceph_aclArrayInput` via:
//
//          Ovh_dedicated_ceph_aclArray{ Ovh_dedicated_ceph_aclArgs{...} }
type Ovh_dedicated_ceph_aclArrayInput interface {
	pulumi.Input

	ToOvh_dedicated_ceph_aclArrayOutput() Ovh_dedicated_ceph_aclArrayOutput
	ToOvh_dedicated_ceph_aclArrayOutputWithContext(context.Context) Ovh_dedicated_ceph_aclArrayOutput
}

type Ovh_dedicated_ceph_aclArray []Ovh_dedicated_ceph_aclInput

func (Ovh_dedicated_ceph_aclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_dedicated_ceph_acl)(nil)).Elem()
}

func (i Ovh_dedicated_ceph_aclArray) ToOvh_dedicated_ceph_aclArrayOutput() Ovh_dedicated_ceph_aclArrayOutput {
	return i.ToOvh_dedicated_ceph_aclArrayOutputWithContext(context.Background())
}

func (i Ovh_dedicated_ceph_aclArray) ToOvh_dedicated_ceph_aclArrayOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_ceph_aclArrayOutput)
}

// Ovh_dedicated_ceph_aclMapInput is an input type that accepts Ovh_dedicated_ceph_aclMap and Ovh_dedicated_ceph_aclMapOutput values.
// You can construct a concrete instance of `Ovh_dedicated_ceph_aclMapInput` via:
//
//          Ovh_dedicated_ceph_aclMap{ "key": Ovh_dedicated_ceph_aclArgs{...} }
type Ovh_dedicated_ceph_aclMapInput interface {
	pulumi.Input

	ToOvh_dedicated_ceph_aclMapOutput() Ovh_dedicated_ceph_aclMapOutput
	ToOvh_dedicated_ceph_aclMapOutputWithContext(context.Context) Ovh_dedicated_ceph_aclMapOutput
}

type Ovh_dedicated_ceph_aclMap map[string]Ovh_dedicated_ceph_aclInput

func (Ovh_dedicated_ceph_aclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_dedicated_ceph_acl)(nil)).Elem()
}

func (i Ovh_dedicated_ceph_aclMap) ToOvh_dedicated_ceph_aclMapOutput() Ovh_dedicated_ceph_aclMapOutput {
	return i.ToOvh_dedicated_ceph_aclMapOutputWithContext(context.Background())
}

func (i Ovh_dedicated_ceph_aclMap) ToOvh_dedicated_ceph_aclMapOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_ceph_aclMapOutput)
}

type Ovh_dedicated_ceph_aclOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_ceph_aclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_dedicated_ceph_acl)(nil))
}

func (o Ovh_dedicated_ceph_aclOutput) ToOvh_dedicated_ceph_aclOutput() Ovh_dedicated_ceph_aclOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclOutput) ToOvh_dedicated_ceph_aclOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclOutput) ToOvh_dedicated_ceph_aclPtrOutput() Ovh_dedicated_ceph_aclPtrOutput {
	return o.ToOvh_dedicated_ceph_aclPtrOutputWithContext(context.Background())
}

func (o Ovh_dedicated_ceph_aclOutput) ToOvh_dedicated_ceph_aclPtrOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_dedicated_ceph_acl) *Ovh_dedicated_ceph_acl {
		return &v
	}).(Ovh_dedicated_ceph_aclPtrOutput)
}

type Ovh_dedicated_ceph_aclPtrOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_ceph_aclPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_dedicated_ceph_acl)(nil))
}

func (o Ovh_dedicated_ceph_aclPtrOutput) ToOvh_dedicated_ceph_aclPtrOutput() Ovh_dedicated_ceph_aclPtrOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclPtrOutput) ToOvh_dedicated_ceph_aclPtrOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclPtrOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclPtrOutput) Elem() Ovh_dedicated_ceph_aclOutput {
	return o.ApplyT(func(v *Ovh_dedicated_ceph_acl) Ovh_dedicated_ceph_acl {
		if v != nil {
			return *v
		}
		var ret Ovh_dedicated_ceph_acl
		return ret
	}).(Ovh_dedicated_ceph_aclOutput)
}

type Ovh_dedicated_ceph_aclArrayOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_ceph_aclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_dedicated_ceph_acl)(nil))
}

func (o Ovh_dedicated_ceph_aclArrayOutput) ToOvh_dedicated_ceph_aclArrayOutput() Ovh_dedicated_ceph_aclArrayOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclArrayOutput) ToOvh_dedicated_ceph_aclArrayOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclArrayOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclArrayOutput) Index(i pulumi.IntInput) Ovh_dedicated_ceph_aclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_dedicated_ceph_acl {
		return vs[0].([]Ovh_dedicated_ceph_acl)[vs[1].(int)]
	}).(Ovh_dedicated_ceph_aclOutput)
}

type Ovh_dedicated_ceph_aclMapOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_ceph_aclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_dedicated_ceph_acl)(nil))
}

func (o Ovh_dedicated_ceph_aclMapOutput) ToOvh_dedicated_ceph_aclMapOutput() Ovh_dedicated_ceph_aclMapOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclMapOutput) ToOvh_dedicated_ceph_aclMapOutputWithContext(ctx context.Context) Ovh_dedicated_ceph_aclMapOutput {
	return o
}

func (o Ovh_dedicated_ceph_aclMapOutput) MapIndex(k pulumi.StringInput) Ovh_dedicated_ceph_aclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_dedicated_ceph_acl {
		return vs[0].(map[string]Ovh_dedicated_ceph_acl)[vs[1].(string)]
	}).(Ovh_dedicated_ceph_aclOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_dedicated_ceph_aclOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_ceph_aclPtrOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_ceph_aclArrayOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_ceph_aclMapOutput{})
}
