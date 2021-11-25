// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_cloud_project_kube_nodepool struct {
	pulumi.CustomResourceState

	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolPtrOutput `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolPtrOutput `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntOutput `pulumi:"availableNodes"`
	// Creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntOutput `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.IntOutput `pulumi:"desiredNodes"`
	// Flavor name
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Flavor name
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// Kube ID
	KubeId pulumi.StringOutput `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes pulumi.IntOutput `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes pulumi.IntOutput `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolPtrOutput `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Project id
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringOutput `pulumi:"sizeStatus"`
	// Current status
	Status pulumi.StringOutput `pulumi:"status"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.IntOutput `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewOvh_cloud_project_kube_nodepool registers a new resource with the given unique name, arguments, and options.
func NewOvh_cloud_project_kube_nodepool(ctx *pulumi.Context,
	name string, args *Ovh_cloud_project_kube_nodepoolArgs, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_kube_nodepool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorName == nil {
		return nil, errors.New("invalid value for required argument 'FlavorName'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_cloud_project_kube_nodepool
	err := ctx.RegisterResource("ovh:index/ovh_cloud_project_kube_nodepool:ovh_cloud_project_kube_nodepool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_cloud_project_kube_nodepool gets an existing Ovh_cloud_project_kube_nodepool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_cloud_project_kube_nodepool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_cloud_project_kube_nodepoolState, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_kube_nodepool, error) {
	var resource Ovh_cloud_project_kube_nodepool
	err := ctx.ReadResource("ovh:index/ovh_cloud_project_kube_nodepool:ovh_cloud_project_kube_nodepool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_cloud_project_kube_nodepool resources.
type ovh_cloud_project_kube_nodepoolState struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale *bool `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes *int `pulumi:"availableNodes"`
	// Creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes *int `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes *int `pulumi:"desiredNodes"`
	// Flavor name
	Flavor *string `pulumi:"flavor"`
	// Flavor name
	FlavorName *string `pulumi:"flavorName"`
	// Kube ID
	KubeId *string `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes *int `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes *int `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name *string `pulumi:"name"`
	// Project id
	ProjectId *string `pulumi:"projectId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus *string `pulumi:"sizeStatus"`
	// Current status
	Status *string `pulumi:"status"`
	// Number of nodes with latest version installed in the pool
	UpToDateNodes *int `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt *string `pulumi:"updatedAt"`
}

type Ovh_cloud_project_kube_nodepoolState struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolPtrInput
	// Number of nodes which are actually ready in the pool
	AvailableNodes pulumi.IntPtrInput
	// Creation date
	CreatedAt pulumi.StringPtrInput
	// Number of nodes present in the pool
	CurrentNodes pulumi.IntPtrInput
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.IntPtrInput
	// Flavor name
	Flavor pulumi.StringPtrInput
	// Flavor name
	FlavorName pulumi.StringPtrInput
	// Kube ID
	KubeId pulumi.StringPtrInput
	// Number of nodes you desire in the pool
	MaxNodes pulumi.IntPtrInput
	// Number of nodes you desire in the pool
	MinNodes pulumi.IntPtrInput
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolPtrInput
	// NodePool resource name
	Name pulumi.StringPtrInput
	// Project id
	ProjectId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus pulumi.StringPtrInput
	// Current status
	Status pulumi.StringPtrInput
	// Number of nodes with latest version installed in the pool
	UpToDateNodes pulumi.IntPtrInput
	// Last update date
	UpdatedAt pulumi.StringPtrInput
}

func (Ovh_cloud_project_kube_nodepoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_kube_nodepoolState)(nil)).Elem()
}

type ovh_cloud_project_kube_nodepoolArgs struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity *bool `pulumi:"antiAffinity"`
	// Enable auto-scaling for the pool
	Autoscale *bool `pulumi:"autoscale"`
	// Number of nodes you desire in the pool
	DesiredNodes *int `pulumi:"desiredNodes"`
	// Flavor name
	FlavorName string `pulumi:"flavorName"`
	// Kube ID
	KubeId string `pulumi:"kubeId"`
	// Number of nodes you desire in the pool
	MaxNodes *int `pulumi:"maxNodes"`
	// Number of nodes you desire in the pool
	MinNodes *int `pulumi:"minNodes"`
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled *bool `pulumi:"monthlyBilled"`
	// NodePool resource name
	Name *string `pulumi:"name"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_cloud_project_kube_nodepool resource.
type Ovh_cloud_project_kube_nodepoolArgs struct {
	// Enable anti affinity groups for nodes in the pool
	AntiAffinity pulumi.BoolPtrInput
	// Enable auto-scaling for the pool
	Autoscale pulumi.BoolPtrInput
	// Number of nodes you desire in the pool
	DesiredNodes pulumi.IntPtrInput
	// Flavor name
	FlavorName pulumi.StringInput
	// Kube ID
	KubeId pulumi.StringInput
	// Number of nodes you desire in the pool
	MaxNodes pulumi.IntPtrInput
	// Number of nodes you desire in the pool
	MinNodes pulumi.IntPtrInput
	// Enable monthly billing on all nodes in the pool
	MonthlyBilled pulumi.BoolPtrInput
	// NodePool resource name
	Name pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringInput
}

func (Ovh_cloud_project_kube_nodepoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_kube_nodepoolArgs)(nil)).Elem()
}

type Ovh_cloud_project_kube_nodepoolInput interface {
	pulumi.Input

	ToOvh_cloud_project_kube_nodepoolOutput() Ovh_cloud_project_kube_nodepoolOutput
	ToOvh_cloud_project_kube_nodepoolOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolOutput
}

func (*Ovh_cloud_project_kube_nodepool) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_kube_nodepool)(nil))
}

func (i *Ovh_cloud_project_kube_nodepool) ToOvh_cloud_project_kube_nodepoolOutput() Ovh_cloud_project_kube_nodepoolOutput {
	return i.ToOvh_cloud_project_kube_nodepoolOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_kube_nodepool) ToOvh_cloud_project_kube_nodepoolOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kube_nodepoolOutput)
}

func (i *Ovh_cloud_project_kube_nodepool) ToOvh_cloud_project_kube_nodepoolPtrOutput() Ovh_cloud_project_kube_nodepoolPtrOutput {
	return i.ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_kube_nodepool) ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kube_nodepoolPtrOutput)
}

type Ovh_cloud_project_kube_nodepoolPtrInput interface {
	pulumi.Input

	ToOvh_cloud_project_kube_nodepoolPtrOutput() Ovh_cloud_project_kube_nodepoolPtrOutput
	ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolPtrOutput
}

type ovh_cloud_project_kube_nodepoolPtrType Ovh_cloud_project_kube_nodepoolArgs

func (*ovh_cloud_project_kube_nodepoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_kube_nodepool)(nil))
}

func (i *ovh_cloud_project_kube_nodepoolPtrType) ToOvh_cloud_project_kube_nodepoolPtrOutput() Ovh_cloud_project_kube_nodepoolPtrOutput {
	return i.ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(context.Background())
}

func (i *ovh_cloud_project_kube_nodepoolPtrType) ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kube_nodepoolPtrOutput)
}

// Ovh_cloud_project_kube_nodepoolArrayInput is an input type that accepts Ovh_cloud_project_kube_nodepoolArray and Ovh_cloud_project_kube_nodepoolArrayOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_kube_nodepoolArrayInput` via:
//
//          Ovh_cloud_project_kube_nodepoolArray{ Ovh_cloud_project_kube_nodepoolArgs{...} }
type Ovh_cloud_project_kube_nodepoolArrayInput interface {
	pulumi.Input

	ToOvh_cloud_project_kube_nodepoolArrayOutput() Ovh_cloud_project_kube_nodepoolArrayOutput
	ToOvh_cloud_project_kube_nodepoolArrayOutputWithContext(context.Context) Ovh_cloud_project_kube_nodepoolArrayOutput
}

type Ovh_cloud_project_kube_nodepoolArray []Ovh_cloud_project_kube_nodepoolInput

func (Ovh_cloud_project_kube_nodepoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_cloud_project_kube_nodepool)(nil)).Elem()
}

func (i Ovh_cloud_project_kube_nodepoolArray) ToOvh_cloud_project_kube_nodepoolArrayOutput() Ovh_cloud_project_kube_nodepoolArrayOutput {
	return i.ToOvh_cloud_project_kube_nodepoolArrayOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_kube_nodepoolArray) ToOvh_cloud_project_kube_nodepoolArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kube_nodepoolArrayOutput)
}

// Ovh_cloud_project_kube_nodepoolMapInput is an input type that accepts Ovh_cloud_project_kube_nodepoolMap and Ovh_cloud_project_kube_nodepoolMapOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_kube_nodepoolMapInput` via:
//
//          Ovh_cloud_project_kube_nodepoolMap{ "key": Ovh_cloud_project_kube_nodepoolArgs{...} }
type Ovh_cloud_project_kube_nodepoolMapInput interface {
	pulumi.Input

	ToOvh_cloud_project_kube_nodepoolMapOutput() Ovh_cloud_project_kube_nodepoolMapOutput
	ToOvh_cloud_project_kube_nodepoolMapOutputWithContext(context.Context) Ovh_cloud_project_kube_nodepoolMapOutput
}

type Ovh_cloud_project_kube_nodepoolMap map[string]Ovh_cloud_project_kube_nodepoolInput

func (Ovh_cloud_project_kube_nodepoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_cloud_project_kube_nodepool)(nil)).Elem()
}

func (i Ovh_cloud_project_kube_nodepoolMap) ToOvh_cloud_project_kube_nodepoolMapOutput() Ovh_cloud_project_kube_nodepoolMapOutput {
	return i.ToOvh_cloud_project_kube_nodepoolMapOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_kube_nodepoolMap) ToOvh_cloud_project_kube_nodepoolMapOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kube_nodepoolMapOutput)
}

type Ovh_cloud_project_kube_nodepoolOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kube_nodepoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_kube_nodepool)(nil))
}

func (o Ovh_cloud_project_kube_nodepoolOutput) ToOvh_cloud_project_kube_nodepoolOutput() Ovh_cloud_project_kube_nodepoolOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolOutput) ToOvh_cloud_project_kube_nodepoolOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolOutput) ToOvh_cloud_project_kube_nodepoolPtrOutput() Ovh_cloud_project_kube_nodepoolPtrOutput {
	return o.ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(context.Background())
}

func (o Ovh_cloud_project_kube_nodepoolOutput) ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_cloud_project_kube_nodepool) *Ovh_cloud_project_kube_nodepool {
		return &v
	}).(Ovh_cloud_project_kube_nodepoolPtrOutput)
}

type Ovh_cloud_project_kube_nodepoolPtrOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kube_nodepoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_kube_nodepool)(nil))
}

func (o Ovh_cloud_project_kube_nodepoolPtrOutput) ToOvh_cloud_project_kube_nodepoolPtrOutput() Ovh_cloud_project_kube_nodepoolPtrOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolPtrOutput) ToOvh_cloud_project_kube_nodepoolPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolPtrOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolPtrOutput) Elem() Ovh_cloud_project_kube_nodepoolOutput {
	return o.ApplyT(func(v *Ovh_cloud_project_kube_nodepool) Ovh_cloud_project_kube_nodepool {
		if v != nil {
			return *v
		}
		var ret Ovh_cloud_project_kube_nodepool
		return ret
	}).(Ovh_cloud_project_kube_nodepoolOutput)
}

type Ovh_cloud_project_kube_nodepoolArrayOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kube_nodepoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_cloud_project_kube_nodepool)(nil))
}

func (o Ovh_cloud_project_kube_nodepoolArrayOutput) ToOvh_cloud_project_kube_nodepoolArrayOutput() Ovh_cloud_project_kube_nodepoolArrayOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolArrayOutput) ToOvh_cloud_project_kube_nodepoolArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolArrayOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolArrayOutput) Index(i pulumi.IntInput) Ovh_cloud_project_kube_nodepoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_cloud_project_kube_nodepool {
		return vs[0].([]Ovh_cloud_project_kube_nodepool)[vs[1].(int)]
	}).(Ovh_cloud_project_kube_nodepoolOutput)
}

type Ovh_cloud_project_kube_nodepoolMapOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kube_nodepoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_cloud_project_kube_nodepool)(nil))
}

func (o Ovh_cloud_project_kube_nodepoolMapOutput) ToOvh_cloud_project_kube_nodepoolMapOutput() Ovh_cloud_project_kube_nodepoolMapOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolMapOutput) ToOvh_cloud_project_kube_nodepoolMapOutputWithContext(ctx context.Context) Ovh_cloud_project_kube_nodepoolMapOutput {
	return o
}

func (o Ovh_cloud_project_kube_nodepoolMapOutput) MapIndex(k pulumi.StringInput) Ovh_cloud_project_kube_nodepoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_cloud_project_kube_nodepool {
		return vs[0].(map[string]Ovh_cloud_project_kube_nodepool)[vs[1].(string)]
	}).(Ovh_cloud_project_kube_nodepoolOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_cloud_project_kube_nodepoolOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kube_nodepoolPtrOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kube_nodepoolArrayOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kube_nodepoolMapOutput{})
}