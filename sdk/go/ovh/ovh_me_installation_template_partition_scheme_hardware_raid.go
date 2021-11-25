// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_me_installation_template_partition_scheme_hardware_raid struct {
	pulumi.CustomResourceState

	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks pulumi.StringArrayOutput `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Hardware RAID name
	Name pulumi.StringOutput `pulumi:"name"`
	// name of this partitioning scheme
	SchemeName pulumi.StringOutput `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID
	Step pulumi.IntOutput `pulumi:"step"`
	// Template name
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewOvh_me_installation_template_partition_scheme_hardware_raid registers a new resource with the given unique name, arguments, and options.
func NewOvh_me_installation_template_partition_scheme_hardware_raid(ctx *pulumi.Context,
	name string, args *Ovh_me_installation_template_partition_scheme_hardware_raidArgs, opts ...pulumi.ResourceOption) (*Ovh_me_installation_template_partition_scheme_hardware_raid, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disks == nil {
		return nil, errors.New("invalid value for required argument 'Disks'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.SchemeName == nil {
		return nil, errors.New("invalid value for required argument 'SchemeName'")
	}
	if args.Step == nil {
		return nil, errors.New("invalid value for required argument 'Step'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	var resource Ovh_me_installation_template_partition_scheme_hardware_raid
	err := ctx.RegisterResource("ovh:index/ovh_me_installation_template_partition_scheme_hardware_raid:ovh_me_installation_template_partition_scheme_hardware_raid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_me_installation_template_partition_scheme_hardware_raid gets an existing Ovh_me_installation_template_partition_scheme_hardware_raid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_me_installation_template_partition_scheme_hardware_raid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_me_installation_template_partition_scheme_hardware_raidState, opts ...pulumi.ResourceOption) (*Ovh_me_installation_template_partition_scheme_hardware_raid, error) {
	var resource Ovh_me_installation_template_partition_scheme_hardware_raid
	err := ctx.ReadResource("ovh:index/ovh_me_installation_template_partition_scheme_hardware_raid:ovh_me_installation_template_partition_scheme_hardware_raid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_me_installation_template_partition_scheme_hardware_raid resources.
type ovh_me_installation_template_partition_scheme_hardware_raidState struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks []string `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode *string `pulumi:"mode"`
	// Hardware RAID name
	Name *string `pulumi:"name"`
	// name of this partitioning scheme
	SchemeName *string `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID
	Step *int `pulumi:"step"`
	// Template name
	TemplateName *string `pulumi:"templateName"`
}

type Ovh_me_installation_template_partition_scheme_hardware_raidState struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks pulumi.StringArrayInput
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode pulumi.StringPtrInput
	// Hardware RAID name
	Name pulumi.StringPtrInput
	// name of this partitioning scheme
	SchemeName pulumi.StringPtrInput
	// Specifies the creation order of the hardware RAID
	Step pulumi.IntPtrInput
	// Template name
	TemplateName pulumi.StringPtrInput
}

func (Ovh_me_installation_template_partition_scheme_hardware_raidState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_me_installation_template_partition_scheme_hardware_raidState)(nil)).Elem()
}

type ovh_me_installation_template_partition_scheme_hardware_raidArgs struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks []string `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode string `pulumi:"mode"`
	// Hardware RAID name
	Name *string `pulumi:"name"`
	// name of this partitioning scheme
	SchemeName string `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID
	Step int `pulumi:"step"`
	// Template name
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a Ovh_me_installation_template_partition_scheme_hardware_raid resource.
type Ovh_me_installation_template_partition_scheme_hardware_raidArgs struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id
	Disks pulumi.StringArrayInput
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60)
	Mode pulumi.StringInput
	// Hardware RAID name
	Name pulumi.StringPtrInput
	// name of this partitioning scheme
	SchemeName pulumi.StringInput
	// Specifies the creation order of the hardware RAID
	Step pulumi.IntInput
	// Template name
	TemplateName pulumi.StringInput
}

func (Ovh_me_installation_template_partition_scheme_hardware_raidArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_me_installation_template_partition_scheme_hardware_raidArgs)(nil)).Elem()
}

type Ovh_me_installation_template_partition_scheme_hardware_raidInput interface {
	pulumi.Input

	ToOvh_me_installation_template_partition_scheme_hardware_raidOutput() Ovh_me_installation_template_partition_scheme_hardware_raidOutput
	ToOvh_me_installation_template_partition_scheme_hardware_raidOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidOutput
}

func (*Ovh_me_installation_template_partition_scheme_hardware_raid) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (i *Ovh_me_installation_template_partition_scheme_hardware_raid) ToOvh_me_installation_template_partition_scheme_hardware_raidOutput() Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return i.ToOvh_me_installation_template_partition_scheme_hardware_raidOutputWithContext(context.Background())
}

func (i *Ovh_me_installation_template_partition_scheme_hardware_raid) ToOvh_me_installation_template_partition_scheme_hardware_raidOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_installation_template_partition_scheme_hardware_raidOutput)
}

func (i *Ovh_me_installation_template_partition_scheme_hardware_raid) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutput() Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return i.ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(context.Background())
}

func (i *Ovh_me_installation_template_partition_scheme_hardware_raid) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput)
}

type Ovh_me_installation_template_partition_scheme_hardware_raidPtrInput interface {
	pulumi.Input

	ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutput() Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput
	ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput
}

type ovh_me_installation_template_partition_scheme_hardware_raidPtrType Ovh_me_installation_template_partition_scheme_hardware_raidArgs

func (*ovh_me_installation_template_partition_scheme_hardware_raidPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (i *ovh_me_installation_template_partition_scheme_hardware_raidPtrType) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutput() Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return i.ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(context.Background())
}

func (i *ovh_me_installation_template_partition_scheme_hardware_raidPtrType) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput)
}

// Ovh_me_installation_template_partition_scheme_hardware_raidArrayInput is an input type that accepts Ovh_me_installation_template_partition_scheme_hardware_raidArray and Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput values.
// You can construct a concrete instance of `Ovh_me_installation_template_partition_scheme_hardware_raidArrayInput` via:
//
//          Ovh_me_installation_template_partition_scheme_hardware_raidArray{ Ovh_me_installation_template_partition_scheme_hardware_raidArgs{...} }
type Ovh_me_installation_template_partition_scheme_hardware_raidArrayInput interface {
	pulumi.Input

	ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutput() Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput
	ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutputWithContext(context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput
}

type Ovh_me_installation_template_partition_scheme_hardware_raidArray []Ovh_me_installation_template_partition_scheme_hardware_raidInput

func (Ovh_me_installation_template_partition_scheme_hardware_raidArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_me_installation_template_partition_scheme_hardware_raid)(nil)).Elem()
}

func (i Ovh_me_installation_template_partition_scheme_hardware_raidArray) ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutput() Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput {
	return i.ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutputWithContext(context.Background())
}

func (i Ovh_me_installation_template_partition_scheme_hardware_raidArray) ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput)
}

// Ovh_me_installation_template_partition_scheme_hardware_raidMapInput is an input type that accepts Ovh_me_installation_template_partition_scheme_hardware_raidMap and Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput values.
// You can construct a concrete instance of `Ovh_me_installation_template_partition_scheme_hardware_raidMapInput` via:
//
//          Ovh_me_installation_template_partition_scheme_hardware_raidMap{ "key": Ovh_me_installation_template_partition_scheme_hardware_raidArgs{...} }
type Ovh_me_installation_template_partition_scheme_hardware_raidMapInput interface {
	pulumi.Input

	ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutput() Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput
	ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutputWithContext(context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput
}

type Ovh_me_installation_template_partition_scheme_hardware_raidMap map[string]Ovh_me_installation_template_partition_scheme_hardware_raidInput

func (Ovh_me_installation_template_partition_scheme_hardware_raidMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_me_installation_template_partition_scheme_hardware_raid)(nil)).Elem()
}

func (i Ovh_me_installation_template_partition_scheme_hardware_raidMap) ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutput() Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput {
	return i.ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutputWithContext(context.Background())
}

func (i Ovh_me_installation_template_partition_scheme_hardware_raidMap) ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput)
}

type Ovh_me_installation_template_partition_scheme_hardware_raidOutput struct{ *pulumi.OutputState }

func (Ovh_me_installation_template_partition_scheme_hardware_raidOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidOutput() Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutput() Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return o.ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(context.Background())
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_me_installation_template_partition_scheme_hardware_raid) *Ovh_me_installation_template_partition_scheme_hardware_raid {
		return &v
	}).(Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput)
}

type Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput struct{ *pulumi.OutputState }

func (Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutput() Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidPtrOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput) Elem() Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return o.ApplyT(func(v *Ovh_me_installation_template_partition_scheme_hardware_raid) Ovh_me_installation_template_partition_scheme_hardware_raid {
		if v != nil {
			return *v
		}
		var ret Ovh_me_installation_template_partition_scheme_hardware_raid
		return ret
	}).(Ovh_me_installation_template_partition_scheme_hardware_raidOutput)
}

type Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput struct{ *pulumi.OutputState }

func (Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutput() Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidArrayOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput) Index(i pulumi.IntInput) Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_me_installation_template_partition_scheme_hardware_raid {
		return vs[0].([]Ovh_me_installation_template_partition_scheme_hardware_raid)[vs[1].(int)]
	}).(Ovh_me_installation_template_partition_scheme_hardware_raidOutput)
}

type Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput struct{ *pulumi.OutputState }

func (Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_me_installation_template_partition_scheme_hardware_raid)(nil))
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutput() Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput) ToOvh_me_installation_template_partition_scheme_hardware_raidMapOutputWithContext(ctx context.Context) Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput {
	return o
}

func (o Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput) MapIndex(k pulumi.StringInput) Ovh_me_installation_template_partition_scheme_hardware_raidOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_me_installation_template_partition_scheme_hardware_raid {
		return vs[0].(map[string]Ovh_me_installation_template_partition_scheme_hardware_raid)[vs[1].(string)]
	}).(Ovh_me_installation_template_partition_scheme_hardware_raidOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_me_installation_template_partition_scheme_hardware_raidOutput{})
	pulumi.RegisterOutputType(Ovh_me_installation_template_partition_scheme_hardware_raidPtrOutput{})
	pulumi.RegisterOutputType(Ovh_me_installation_template_partition_scheme_hardware_raidArrayOutput{})
	pulumi.RegisterOutputType(Ovh_me_installation_template_partition_scheme_hardware_raidMapOutput{})
}