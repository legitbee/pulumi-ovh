// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainZoneRedirection struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Keywords    pulumi.StringPtrOutput `pulumi:"keywords"`
	Subdomain   pulumi.StringPtrOutput `pulumi:"subdomain"`
	Target      pulumi.StringOutput    `pulumi:"target"`
	Title       pulumi.StringPtrOutput `pulumi:"title"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Zone        pulumi.StringOutput    `pulumi:"zone"`
}

// NewDomainZoneRedirection registers a new resource with the given unique name, arguments, and options.
func NewDomainZoneRedirection(ctx *pulumi.Context,
	name string, args *DomainZoneRedirectionArgs, opts ...pulumi.ResourceOption) (*DomainZoneRedirection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource DomainZoneRedirection
	err := ctx.RegisterResource("ovh:index/domainZoneRedirection:DomainZoneRedirection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainZoneRedirection gets an existing DomainZoneRedirection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainZoneRedirection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainZoneRedirectionState, opts ...pulumi.ResourceOption) (*DomainZoneRedirection, error) {
	var resource DomainZoneRedirection
	err := ctx.ReadResource("ovh:index/domainZoneRedirection:DomainZoneRedirection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainZoneRedirection resources.
type domainZoneRedirectionState struct {
	Description *string `pulumi:"description"`
	Keywords    *string `pulumi:"keywords"`
	Subdomain   *string `pulumi:"subdomain"`
	Target      *string `pulumi:"target"`
	Title       *string `pulumi:"title"`
	Type        *string `pulumi:"type"`
	Zone        *string `pulumi:"zone"`
}

type DomainZoneRedirectionState struct {
	Description pulumi.StringPtrInput
	Keywords    pulumi.StringPtrInput
	Subdomain   pulumi.StringPtrInput
	Target      pulumi.StringPtrInput
	Title       pulumi.StringPtrInput
	Type        pulumi.StringPtrInput
	Zone        pulumi.StringPtrInput
}

func (DomainZoneRedirectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneRedirectionState)(nil)).Elem()
}

type domainZoneRedirectionArgs struct {
	Description *string `pulumi:"description"`
	Keywords    *string `pulumi:"keywords"`
	Subdomain   *string `pulumi:"subdomain"`
	Target      string  `pulumi:"target"`
	Title       *string `pulumi:"title"`
	Type        string  `pulumi:"type"`
	Zone        string  `pulumi:"zone"`
}

// The set of arguments for constructing a DomainZoneRedirection resource.
type DomainZoneRedirectionArgs struct {
	Description pulumi.StringPtrInput
	Keywords    pulumi.StringPtrInput
	Subdomain   pulumi.StringPtrInput
	Target      pulumi.StringInput
	Title       pulumi.StringPtrInput
	Type        pulumi.StringInput
	Zone        pulumi.StringInput
}

func (DomainZoneRedirectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneRedirectionArgs)(nil)).Elem()
}

type DomainZoneRedirectionInput interface {
	pulumi.Input

	ToDomainZoneRedirectionOutput() DomainZoneRedirectionOutput
	ToDomainZoneRedirectionOutputWithContext(ctx context.Context) DomainZoneRedirectionOutput
}

func (*DomainZoneRedirection) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainZoneRedirection)(nil))
}

func (i *DomainZoneRedirection) ToDomainZoneRedirectionOutput() DomainZoneRedirectionOutput {
	return i.ToDomainZoneRedirectionOutputWithContext(context.Background())
}

func (i *DomainZoneRedirection) ToDomainZoneRedirectionOutputWithContext(ctx context.Context) DomainZoneRedirectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRedirectionOutput)
}

func (i *DomainZoneRedirection) ToDomainZoneRedirectionPtrOutput() DomainZoneRedirectionPtrOutput {
	return i.ToDomainZoneRedirectionPtrOutputWithContext(context.Background())
}

func (i *DomainZoneRedirection) ToDomainZoneRedirectionPtrOutputWithContext(ctx context.Context) DomainZoneRedirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRedirectionPtrOutput)
}

type DomainZoneRedirectionPtrInput interface {
	pulumi.Input

	ToDomainZoneRedirectionPtrOutput() DomainZoneRedirectionPtrOutput
	ToDomainZoneRedirectionPtrOutputWithContext(ctx context.Context) DomainZoneRedirectionPtrOutput
}

type domainZoneRedirectionPtrType DomainZoneRedirectionArgs

func (*domainZoneRedirectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneRedirection)(nil))
}

func (i *domainZoneRedirectionPtrType) ToDomainZoneRedirectionPtrOutput() DomainZoneRedirectionPtrOutput {
	return i.ToDomainZoneRedirectionPtrOutputWithContext(context.Background())
}

func (i *domainZoneRedirectionPtrType) ToDomainZoneRedirectionPtrOutputWithContext(ctx context.Context) DomainZoneRedirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRedirectionPtrOutput)
}

// DomainZoneRedirectionArrayInput is an input type that accepts DomainZoneRedirectionArray and DomainZoneRedirectionArrayOutput values.
// You can construct a concrete instance of `DomainZoneRedirectionArrayInput` via:
//
//          DomainZoneRedirectionArray{ DomainZoneRedirectionArgs{...} }
type DomainZoneRedirectionArrayInput interface {
	pulumi.Input

	ToDomainZoneRedirectionArrayOutput() DomainZoneRedirectionArrayOutput
	ToDomainZoneRedirectionArrayOutputWithContext(context.Context) DomainZoneRedirectionArrayOutput
}

type DomainZoneRedirectionArray []DomainZoneRedirectionInput

func (DomainZoneRedirectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZoneRedirection)(nil)).Elem()
}

func (i DomainZoneRedirectionArray) ToDomainZoneRedirectionArrayOutput() DomainZoneRedirectionArrayOutput {
	return i.ToDomainZoneRedirectionArrayOutputWithContext(context.Background())
}

func (i DomainZoneRedirectionArray) ToDomainZoneRedirectionArrayOutputWithContext(ctx context.Context) DomainZoneRedirectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRedirectionArrayOutput)
}

// DomainZoneRedirectionMapInput is an input type that accepts DomainZoneRedirectionMap and DomainZoneRedirectionMapOutput values.
// You can construct a concrete instance of `DomainZoneRedirectionMapInput` via:
//
//          DomainZoneRedirectionMap{ "key": DomainZoneRedirectionArgs{...} }
type DomainZoneRedirectionMapInput interface {
	pulumi.Input

	ToDomainZoneRedirectionMapOutput() DomainZoneRedirectionMapOutput
	ToDomainZoneRedirectionMapOutputWithContext(context.Context) DomainZoneRedirectionMapOutput
}

type DomainZoneRedirectionMap map[string]DomainZoneRedirectionInput

func (DomainZoneRedirectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZoneRedirection)(nil)).Elem()
}

func (i DomainZoneRedirectionMap) ToDomainZoneRedirectionMapOutput() DomainZoneRedirectionMapOutput {
	return i.ToDomainZoneRedirectionMapOutputWithContext(context.Background())
}

func (i DomainZoneRedirectionMap) ToDomainZoneRedirectionMapOutputWithContext(ctx context.Context) DomainZoneRedirectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRedirectionMapOutput)
}

type DomainZoneRedirectionOutput struct{ *pulumi.OutputState }

func (DomainZoneRedirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainZoneRedirection)(nil))
}

func (o DomainZoneRedirectionOutput) ToDomainZoneRedirectionOutput() DomainZoneRedirectionOutput {
	return o
}

func (o DomainZoneRedirectionOutput) ToDomainZoneRedirectionOutputWithContext(ctx context.Context) DomainZoneRedirectionOutput {
	return o
}

func (o DomainZoneRedirectionOutput) ToDomainZoneRedirectionPtrOutput() DomainZoneRedirectionPtrOutput {
	return o.ToDomainZoneRedirectionPtrOutputWithContext(context.Background())
}

func (o DomainZoneRedirectionOutput) ToDomainZoneRedirectionPtrOutputWithContext(ctx context.Context) DomainZoneRedirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainZoneRedirection) *DomainZoneRedirection {
		return &v
	}).(DomainZoneRedirectionPtrOutput)
}

type DomainZoneRedirectionPtrOutput struct{ *pulumi.OutputState }

func (DomainZoneRedirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneRedirection)(nil))
}

func (o DomainZoneRedirectionPtrOutput) ToDomainZoneRedirectionPtrOutput() DomainZoneRedirectionPtrOutput {
	return o
}

func (o DomainZoneRedirectionPtrOutput) ToDomainZoneRedirectionPtrOutputWithContext(ctx context.Context) DomainZoneRedirectionPtrOutput {
	return o
}

func (o DomainZoneRedirectionPtrOutput) Elem() DomainZoneRedirectionOutput {
	return o.ApplyT(func(v *DomainZoneRedirection) DomainZoneRedirection {
		if v != nil {
			return *v
		}
		var ret DomainZoneRedirection
		return ret
	}).(DomainZoneRedirectionOutput)
}

type DomainZoneRedirectionArrayOutput struct{ *pulumi.OutputState }

func (DomainZoneRedirectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainZoneRedirection)(nil))
}

func (o DomainZoneRedirectionArrayOutput) ToDomainZoneRedirectionArrayOutput() DomainZoneRedirectionArrayOutput {
	return o
}

func (o DomainZoneRedirectionArrayOutput) ToDomainZoneRedirectionArrayOutputWithContext(ctx context.Context) DomainZoneRedirectionArrayOutput {
	return o
}

func (o DomainZoneRedirectionArrayOutput) Index(i pulumi.IntInput) DomainZoneRedirectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainZoneRedirection {
		return vs[0].([]DomainZoneRedirection)[vs[1].(int)]
	}).(DomainZoneRedirectionOutput)
}

type DomainZoneRedirectionMapOutput struct{ *pulumi.OutputState }

func (DomainZoneRedirectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainZoneRedirection)(nil))
}

func (o DomainZoneRedirectionMapOutput) ToDomainZoneRedirectionMapOutput() DomainZoneRedirectionMapOutput {
	return o
}

func (o DomainZoneRedirectionMapOutput) ToDomainZoneRedirectionMapOutputWithContext(ctx context.Context) DomainZoneRedirectionMapOutput {
	return o
}

func (o DomainZoneRedirectionMapOutput) MapIndex(k pulumi.StringInput) DomainZoneRedirectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainZoneRedirection {
		return vs[0].(map[string]DomainZoneRedirection)[vs[1].(string)]
	}).(DomainZoneRedirectionOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainZoneRedirectionOutput{})
	pulumi.RegisterOutputType(DomainZoneRedirectionPtrOutput{})
	pulumi.RegisterOutputType(DomainZoneRedirectionArrayOutput{})
	pulumi.RegisterOutputType(DomainZoneRedirectionMapOutput{})
}
