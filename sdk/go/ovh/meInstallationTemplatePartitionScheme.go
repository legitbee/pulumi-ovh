// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MeInstallationTemplatePartitionScheme struct {
	pulumi.CustomResourceState

	// name of this partitioning scheme
	Name pulumi.StringOutput `pulumi:"name"`
	// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
	// among all the compatible partitioning schemes (given the underlying hardware specifications)
	Priority pulumi.IntOutput `pulumi:"priority"`
	// This template name
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewMeInstallationTemplatePartitionScheme registers a new resource with the given unique name, arguments, and options.
func NewMeInstallationTemplatePartitionScheme(ctx *pulumi.Context,
	name string, args *MeInstallationTemplatePartitionSchemeArgs, opts ...pulumi.ResourceOption) (*MeInstallationTemplatePartitionScheme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	var resource MeInstallationTemplatePartitionScheme
	err := ctx.RegisterResource("ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeInstallationTemplatePartitionScheme gets an existing MeInstallationTemplatePartitionScheme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeInstallationTemplatePartitionScheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeInstallationTemplatePartitionSchemeState, opts ...pulumi.ResourceOption) (*MeInstallationTemplatePartitionScheme, error) {
	var resource MeInstallationTemplatePartitionScheme
	err := ctx.ReadResource("ovh:index/meInstallationTemplatePartitionScheme:MeInstallationTemplatePartitionScheme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeInstallationTemplatePartitionScheme resources.
type meInstallationTemplatePartitionSchemeState struct {
	// name of this partitioning scheme
	Name *string `pulumi:"name"`
	// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
	// among all the compatible partitioning schemes (given the underlying hardware specifications)
	Priority *int `pulumi:"priority"`
	// This template name
	TemplateName *string `pulumi:"templateName"`
}

type MeInstallationTemplatePartitionSchemeState struct {
	// name of this partitioning scheme
	Name pulumi.StringPtrInput
	// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
	// among all the compatible partitioning schemes (given the underlying hardware specifications)
	Priority pulumi.IntPtrInput
	// This template name
	TemplateName pulumi.StringPtrInput
}

func (MeInstallationTemplatePartitionSchemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*meInstallationTemplatePartitionSchemeState)(nil)).Elem()
}

type meInstallationTemplatePartitionSchemeArgs struct {
	// name of this partitioning scheme
	Name *string `pulumi:"name"`
	// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
	// among all the compatible partitioning schemes (given the underlying hardware specifications)
	Priority int `pulumi:"priority"`
	// This template name
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a MeInstallationTemplatePartitionScheme resource.
type MeInstallationTemplatePartitionSchemeArgs struct {
	// name of this partitioning scheme
	Name pulumi.StringPtrInput
	// on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default,
	// among all the compatible partitioning schemes (given the underlying hardware specifications)
	Priority pulumi.IntInput
	// This template name
	TemplateName pulumi.StringInput
}

func (MeInstallationTemplatePartitionSchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meInstallationTemplatePartitionSchemeArgs)(nil)).Elem()
}

type MeInstallationTemplatePartitionSchemeInput interface {
	pulumi.Input

	ToMeInstallationTemplatePartitionSchemeOutput() MeInstallationTemplatePartitionSchemeOutput
	ToMeInstallationTemplatePartitionSchemeOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeOutput
}

func (*MeInstallationTemplatePartitionScheme) ElementType() reflect.Type {
	return reflect.TypeOf((*MeInstallationTemplatePartitionScheme)(nil))
}

func (i *MeInstallationTemplatePartitionScheme) ToMeInstallationTemplatePartitionSchemeOutput() MeInstallationTemplatePartitionSchemeOutput {
	return i.ToMeInstallationTemplatePartitionSchemeOutputWithContext(context.Background())
}

func (i *MeInstallationTemplatePartitionScheme) ToMeInstallationTemplatePartitionSchemeOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeInstallationTemplatePartitionSchemeOutput)
}

func (i *MeInstallationTemplatePartitionScheme) ToMeInstallationTemplatePartitionSchemePtrOutput() MeInstallationTemplatePartitionSchemePtrOutput {
	return i.ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(context.Background())
}

func (i *MeInstallationTemplatePartitionScheme) ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeInstallationTemplatePartitionSchemePtrOutput)
}

type MeInstallationTemplatePartitionSchemePtrInput interface {
	pulumi.Input

	ToMeInstallationTemplatePartitionSchemePtrOutput() MeInstallationTemplatePartitionSchemePtrOutput
	ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemePtrOutput
}

type meInstallationTemplatePartitionSchemePtrType MeInstallationTemplatePartitionSchemeArgs

func (*meInstallationTemplatePartitionSchemePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MeInstallationTemplatePartitionScheme)(nil))
}

func (i *meInstallationTemplatePartitionSchemePtrType) ToMeInstallationTemplatePartitionSchemePtrOutput() MeInstallationTemplatePartitionSchemePtrOutput {
	return i.ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(context.Background())
}

func (i *meInstallationTemplatePartitionSchemePtrType) ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeInstallationTemplatePartitionSchemePtrOutput)
}

// MeInstallationTemplatePartitionSchemeArrayInput is an input type that accepts MeInstallationTemplatePartitionSchemeArray and MeInstallationTemplatePartitionSchemeArrayOutput values.
// You can construct a concrete instance of `MeInstallationTemplatePartitionSchemeArrayInput` via:
//
//          MeInstallationTemplatePartitionSchemeArray{ MeInstallationTemplatePartitionSchemeArgs{...} }
type MeInstallationTemplatePartitionSchemeArrayInput interface {
	pulumi.Input

	ToMeInstallationTemplatePartitionSchemeArrayOutput() MeInstallationTemplatePartitionSchemeArrayOutput
	ToMeInstallationTemplatePartitionSchemeArrayOutputWithContext(context.Context) MeInstallationTemplatePartitionSchemeArrayOutput
}

type MeInstallationTemplatePartitionSchemeArray []MeInstallationTemplatePartitionSchemeInput

func (MeInstallationTemplatePartitionSchemeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MeInstallationTemplatePartitionScheme)(nil)).Elem()
}

func (i MeInstallationTemplatePartitionSchemeArray) ToMeInstallationTemplatePartitionSchemeArrayOutput() MeInstallationTemplatePartitionSchemeArrayOutput {
	return i.ToMeInstallationTemplatePartitionSchemeArrayOutputWithContext(context.Background())
}

func (i MeInstallationTemplatePartitionSchemeArray) ToMeInstallationTemplatePartitionSchemeArrayOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeInstallationTemplatePartitionSchemeArrayOutput)
}

// MeInstallationTemplatePartitionSchemeMapInput is an input type that accepts MeInstallationTemplatePartitionSchemeMap and MeInstallationTemplatePartitionSchemeMapOutput values.
// You can construct a concrete instance of `MeInstallationTemplatePartitionSchemeMapInput` via:
//
//          MeInstallationTemplatePartitionSchemeMap{ "key": MeInstallationTemplatePartitionSchemeArgs{...} }
type MeInstallationTemplatePartitionSchemeMapInput interface {
	pulumi.Input

	ToMeInstallationTemplatePartitionSchemeMapOutput() MeInstallationTemplatePartitionSchemeMapOutput
	ToMeInstallationTemplatePartitionSchemeMapOutputWithContext(context.Context) MeInstallationTemplatePartitionSchemeMapOutput
}

type MeInstallationTemplatePartitionSchemeMap map[string]MeInstallationTemplatePartitionSchemeInput

func (MeInstallationTemplatePartitionSchemeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MeInstallationTemplatePartitionScheme)(nil)).Elem()
}

func (i MeInstallationTemplatePartitionSchemeMap) ToMeInstallationTemplatePartitionSchemeMapOutput() MeInstallationTemplatePartitionSchemeMapOutput {
	return i.ToMeInstallationTemplatePartitionSchemeMapOutputWithContext(context.Background())
}

func (i MeInstallationTemplatePartitionSchemeMap) ToMeInstallationTemplatePartitionSchemeMapOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeInstallationTemplatePartitionSchemeMapOutput)
}

type MeInstallationTemplatePartitionSchemeOutput struct{ *pulumi.OutputState }

func (MeInstallationTemplatePartitionSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeInstallationTemplatePartitionScheme)(nil))
}

func (o MeInstallationTemplatePartitionSchemeOutput) ToMeInstallationTemplatePartitionSchemeOutput() MeInstallationTemplatePartitionSchemeOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeOutput) ToMeInstallationTemplatePartitionSchemeOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeOutput) ToMeInstallationTemplatePartitionSchemePtrOutput() MeInstallationTemplatePartitionSchemePtrOutput {
	return o.ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(context.Background())
}

func (o MeInstallationTemplatePartitionSchemeOutput) ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MeInstallationTemplatePartitionScheme) *MeInstallationTemplatePartitionScheme {
		return &v
	}).(MeInstallationTemplatePartitionSchemePtrOutput)
}

type MeInstallationTemplatePartitionSchemePtrOutput struct{ *pulumi.OutputState }

func (MeInstallationTemplatePartitionSchemePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeInstallationTemplatePartitionScheme)(nil))
}

func (o MeInstallationTemplatePartitionSchemePtrOutput) ToMeInstallationTemplatePartitionSchemePtrOutput() MeInstallationTemplatePartitionSchemePtrOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemePtrOutput) ToMeInstallationTemplatePartitionSchemePtrOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemePtrOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemePtrOutput) Elem() MeInstallationTemplatePartitionSchemeOutput {
	return o.ApplyT(func(v *MeInstallationTemplatePartitionScheme) MeInstallationTemplatePartitionScheme {
		if v != nil {
			return *v
		}
		var ret MeInstallationTemplatePartitionScheme
		return ret
	}).(MeInstallationTemplatePartitionSchemeOutput)
}

type MeInstallationTemplatePartitionSchemeArrayOutput struct{ *pulumi.OutputState }

func (MeInstallationTemplatePartitionSchemeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MeInstallationTemplatePartitionScheme)(nil))
}

func (o MeInstallationTemplatePartitionSchemeArrayOutput) ToMeInstallationTemplatePartitionSchemeArrayOutput() MeInstallationTemplatePartitionSchemeArrayOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeArrayOutput) ToMeInstallationTemplatePartitionSchemeArrayOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeArrayOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeArrayOutput) Index(i pulumi.IntInput) MeInstallationTemplatePartitionSchemeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MeInstallationTemplatePartitionScheme {
		return vs[0].([]MeInstallationTemplatePartitionScheme)[vs[1].(int)]
	}).(MeInstallationTemplatePartitionSchemeOutput)
}

type MeInstallationTemplatePartitionSchemeMapOutput struct{ *pulumi.OutputState }

func (MeInstallationTemplatePartitionSchemeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MeInstallationTemplatePartitionScheme)(nil))
}

func (o MeInstallationTemplatePartitionSchemeMapOutput) ToMeInstallationTemplatePartitionSchemeMapOutput() MeInstallationTemplatePartitionSchemeMapOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeMapOutput) ToMeInstallationTemplatePartitionSchemeMapOutputWithContext(ctx context.Context) MeInstallationTemplatePartitionSchemeMapOutput {
	return o
}

func (o MeInstallationTemplatePartitionSchemeMapOutput) MapIndex(k pulumi.StringInput) MeInstallationTemplatePartitionSchemeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MeInstallationTemplatePartitionScheme {
		return vs[0].(map[string]MeInstallationTemplatePartitionScheme)[vs[1].(string)]
	}).(MeInstallationTemplatePartitionSchemeOutput)
}

func init() {
	pulumi.RegisterOutputType(MeInstallationTemplatePartitionSchemeOutput{})
	pulumi.RegisterOutputType(MeInstallationTemplatePartitionSchemePtrOutput{})
	pulumi.RegisterOutputType(MeInstallationTemplatePartitionSchemeArrayOutput{})
	pulumi.RegisterOutputType(MeInstallationTemplatePartitionSchemeMapOutput{})
}
