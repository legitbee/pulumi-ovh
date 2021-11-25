// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_dedicated_server_reboot_task struct {
	pulumi.CustomResourceState

	// Details of this task
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Completion date
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name
	Function pulumi.StringOutput `pulumi:"function"`
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayOutput `pulumi:"keepers"`
	// Last update
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// The internal name of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task Creation date
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewOvh_dedicated_server_reboot_task registers a new resource with the given unique name, arguments, and options.
func NewOvh_dedicated_server_reboot_task(ctx *pulumi.Context,
	name string, args *Ovh_dedicated_server_reboot_taskArgs, opts ...pulumi.ResourceOption) (*Ovh_dedicated_server_reboot_task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keepers == nil {
		return nil, errors.New("invalid value for required argument 'Keepers'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_dedicated_server_reboot_task
	err := ctx.RegisterResource("ovh:index/ovh_dedicated_server_reboot_task:ovh_dedicated_server_reboot_task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_dedicated_server_reboot_task gets an existing Ovh_dedicated_server_reboot_task resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_dedicated_server_reboot_task(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_dedicated_server_reboot_taskState, opts ...pulumi.ResourceOption) (*Ovh_dedicated_server_reboot_task, error) {
	var resource Ovh_dedicated_server_reboot_task
	err := ctx.ReadResource("ovh:index/ovh_dedicated_server_reboot_task:ovh_dedicated_server_reboot_task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_dedicated_server_reboot_task resources.
type ovh_dedicated_server_reboot_taskState struct {
	// Details of this task
	Comment *string `pulumi:"comment"`
	// Completion date
	DoneDate *string `pulumi:"doneDate"`
	// Function name
	Function *string `pulumi:"function"`
	// Change this value to recreate a reboot task.
	Keepers []string `pulumi:"keepers"`
	// Last update
	LastUpdate *string `pulumi:"lastUpdate"`
	// The internal name of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task Creation date
	StartDate *string `pulumi:"startDate"`
	// Task status
	Status *string `pulumi:"status"`
}

type Ovh_dedicated_server_reboot_taskState struct {
	// Details of this task
	Comment pulumi.StringPtrInput
	// Completion date
	DoneDate pulumi.StringPtrInput
	// Function name
	Function pulumi.StringPtrInput
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayInput
	// Last update
	LastUpdate pulumi.StringPtrInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task Creation date
	StartDate pulumi.StringPtrInput
	// Task status
	Status pulumi.StringPtrInput
}

func (Ovh_dedicated_server_reboot_taskState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_dedicated_server_reboot_taskState)(nil)).Elem()
}

type ovh_dedicated_server_reboot_taskArgs struct {
	// Change this value to recreate a reboot task.
	Keepers []string `pulumi:"keepers"`
	// The internal name of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_dedicated_server_reboot_task resource.
type Ovh_dedicated_server_reboot_taskArgs struct {
	// Change this value to recreate a reboot task.
	Keepers pulumi.StringArrayInput
	// The internal name of your dedicated server.
	ServiceName pulumi.StringInput
}

func (Ovh_dedicated_server_reboot_taskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_dedicated_server_reboot_taskArgs)(nil)).Elem()
}

type Ovh_dedicated_server_reboot_taskInput interface {
	pulumi.Input

	ToOvh_dedicated_server_reboot_taskOutput() Ovh_dedicated_server_reboot_taskOutput
	ToOvh_dedicated_server_reboot_taskOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskOutput
}

func (*Ovh_dedicated_server_reboot_task) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_dedicated_server_reboot_task)(nil))
}

func (i *Ovh_dedicated_server_reboot_task) ToOvh_dedicated_server_reboot_taskOutput() Ovh_dedicated_server_reboot_taskOutput {
	return i.ToOvh_dedicated_server_reboot_taskOutputWithContext(context.Background())
}

func (i *Ovh_dedicated_server_reboot_task) ToOvh_dedicated_server_reboot_taskOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_server_reboot_taskOutput)
}

func (i *Ovh_dedicated_server_reboot_task) ToOvh_dedicated_server_reboot_taskPtrOutput() Ovh_dedicated_server_reboot_taskPtrOutput {
	return i.ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(context.Background())
}

func (i *Ovh_dedicated_server_reboot_task) ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_server_reboot_taskPtrOutput)
}

type Ovh_dedicated_server_reboot_taskPtrInput interface {
	pulumi.Input

	ToOvh_dedicated_server_reboot_taskPtrOutput() Ovh_dedicated_server_reboot_taskPtrOutput
	ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskPtrOutput
}

type ovh_dedicated_server_reboot_taskPtrType Ovh_dedicated_server_reboot_taskArgs

func (*ovh_dedicated_server_reboot_taskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_dedicated_server_reboot_task)(nil))
}

func (i *ovh_dedicated_server_reboot_taskPtrType) ToOvh_dedicated_server_reboot_taskPtrOutput() Ovh_dedicated_server_reboot_taskPtrOutput {
	return i.ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(context.Background())
}

func (i *ovh_dedicated_server_reboot_taskPtrType) ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_server_reboot_taskPtrOutput)
}

// Ovh_dedicated_server_reboot_taskArrayInput is an input type that accepts Ovh_dedicated_server_reboot_taskArray and Ovh_dedicated_server_reboot_taskArrayOutput values.
// You can construct a concrete instance of `Ovh_dedicated_server_reboot_taskArrayInput` via:
//
//          Ovh_dedicated_server_reboot_taskArray{ Ovh_dedicated_server_reboot_taskArgs{...} }
type Ovh_dedicated_server_reboot_taskArrayInput interface {
	pulumi.Input

	ToOvh_dedicated_server_reboot_taskArrayOutput() Ovh_dedicated_server_reboot_taskArrayOutput
	ToOvh_dedicated_server_reboot_taskArrayOutputWithContext(context.Context) Ovh_dedicated_server_reboot_taskArrayOutput
}

type Ovh_dedicated_server_reboot_taskArray []Ovh_dedicated_server_reboot_taskInput

func (Ovh_dedicated_server_reboot_taskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_dedicated_server_reboot_task)(nil)).Elem()
}

func (i Ovh_dedicated_server_reboot_taskArray) ToOvh_dedicated_server_reboot_taskArrayOutput() Ovh_dedicated_server_reboot_taskArrayOutput {
	return i.ToOvh_dedicated_server_reboot_taskArrayOutputWithContext(context.Background())
}

func (i Ovh_dedicated_server_reboot_taskArray) ToOvh_dedicated_server_reboot_taskArrayOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_server_reboot_taskArrayOutput)
}

// Ovh_dedicated_server_reboot_taskMapInput is an input type that accepts Ovh_dedicated_server_reboot_taskMap and Ovh_dedicated_server_reboot_taskMapOutput values.
// You can construct a concrete instance of `Ovh_dedicated_server_reboot_taskMapInput` via:
//
//          Ovh_dedicated_server_reboot_taskMap{ "key": Ovh_dedicated_server_reboot_taskArgs{...} }
type Ovh_dedicated_server_reboot_taskMapInput interface {
	pulumi.Input

	ToOvh_dedicated_server_reboot_taskMapOutput() Ovh_dedicated_server_reboot_taskMapOutput
	ToOvh_dedicated_server_reboot_taskMapOutputWithContext(context.Context) Ovh_dedicated_server_reboot_taskMapOutput
}

type Ovh_dedicated_server_reboot_taskMap map[string]Ovh_dedicated_server_reboot_taskInput

func (Ovh_dedicated_server_reboot_taskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_dedicated_server_reboot_task)(nil)).Elem()
}

func (i Ovh_dedicated_server_reboot_taskMap) ToOvh_dedicated_server_reboot_taskMapOutput() Ovh_dedicated_server_reboot_taskMapOutput {
	return i.ToOvh_dedicated_server_reboot_taskMapOutputWithContext(context.Background())
}

func (i Ovh_dedicated_server_reboot_taskMap) ToOvh_dedicated_server_reboot_taskMapOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_dedicated_server_reboot_taskMapOutput)
}

type Ovh_dedicated_server_reboot_taskOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_server_reboot_taskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_dedicated_server_reboot_task)(nil))
}

func (o Ovh_dedicated_server_reboot_taskOutput) ToOvh_dedicated_server_reboot_taskOutput() Ovh_dedicated_server_reboot_taskOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskOutput) ToOvh_dedicated_server_reboot_taskOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskOutput) ToOvh_dedicated_server_reboot_taskPtrOutput() Ovh_dedicated_server_reboot_taskPtrOutput {
	return o.ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(context.Background())
}

func (o Ovh_dedicated_server_reboot_taskOutput) ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_dedicated_server_reboot_task) *Ovh_dedicated_server_reboot_task {
		return &v
	}).(Ovh_dedicated_server_reboot_taskPtrOutput)
}

type Ovh_dedicated_server_reboot_taskPtrOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_server_reboot_taskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_dedicated_server_reboot_task)(nil))
}

func (o Ovh_dedicated_server_reboot_taskPtrOutput) ToOvh_dedicated_server_reboot_taskPtrOutput() Ovh_dedicated_server_reboot_taskPtrOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskPtrOutput) ToOvh_dedicated_server_reboot_taskPtrOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskPtrOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskPtrOutput) Elem() Ovh_dedicated_server_reboot_taskOutput {
	return o.ApplyT(func(v *Ovh_dedicated_server_reboot_task) Ovh_dedicated_server_reboot_task {
		if v != nil {
			return *v
		}
		var ret Ovh_dedicated_server_reboot_task
		return ret
	}).(Ovh_dedicated_server_reboot_taskOutput)
}

type Ovh_dedicated_server_reboot_taskArrayOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_server_reboot_taskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_dedicated_server_reboot_task)(nil))
}

func (o Ovh_dedicated_server_reboot_taskArrayOutput) ToOvh_dedicated_server_reboot_taskArrayOutput() Ovh_dedicated_server_reboot_taskArrayOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskArrayOutput) ToOvh_dedicated_server_reboot_taskArrayOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskArrayOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskArrayOutput) Index(i pulumi.IntInput) Ovh_dedicated_server_reboot_taskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_dedicated_server_reboot_task {
		return vs[0].([]Ovh_dedicated_server_reboot_task)[vs[1].(int)]
	}).(Ovh_dedicated_server_reboot_taskOutput)
}

type Ovh_dedicated_server_reboot_taskMapOutput struct{ *pulumi.OutputState }

func (Ovh_dedicated_server_reboot_taskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_dedicated_server_reboot_task)(nil))
}

func (o Ovh_dedicated_server_reboot_taskMapOutput) ToOvh_dedicated_server_reboot_taskMapOutput() Ovh_dedicated_server_reboot_taskMapOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskMapOutput) ToOvh_dedicated_server_reboot_taskMapOutputWithContext(ctx context.Context) Ovh_dedicated_server_reboot_taskMapOutput {
	return o
}

func (o Ovh_dedicated_server_reboot_taskMapOutput) MapIndex(k pulumi.StringInput) Ovh_dedicated_server_reboot_taskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_dedicated_server_reboot_task {
		return vs[0].(map[string]Ovh_dedicated_server_reboot_task)[vs[1].(string)]
	}).(Ovh_dedicated_server_reboot_taskOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_dedicated_server_reboot_taskOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_server_reboot_taskPtrOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_server_reboot_taskArrayOutput{})
	pulumi.RegisterOutputType(Ovh_dedicated_server_reboot_taskMapOutput{})
}