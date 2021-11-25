// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudProjectKubeNodepool struct {
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

// NewCloudProjectKubeNodepool registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectKubeNodepool(ctx *pulumi.Context,
	name string, args *CloudProjectKubeNodepoolArgs, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodepool, error) {
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
	var resource CloudProjectKubeNodepool
	err := ctx.RegisterResource("ovh:index/cloudProjectKubeNodepool:CloudProjectKubeNodepool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectKubeNodepool gets an existing CloudProjectKubeNodepool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectKubeNodepool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectKubeNodepoolState, opts ...pulumi.ResourceOption) (*CloudProjectKubeNodepool, error) {
	var resource CloudProjectKubeNodepool
	err := ctx.ReadResource("ovh:index/cloudProjectKubeNodepool:CloudProjectKubeNodepool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectKubeNodepool resources.
type cloudProjectKubeNodepoolState struct {
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

type CloudProjectKubeNodepoolState struct {
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

func (CloudProjectKubeNodepoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodepoolState)(nil)).Elem()
}

type cloudProjectKubeNodepoolArgs struct {
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

// The set of arguments for constructing a CloudProjectKubeNodepool resource.
type CloudProjectKubeNodepoolArgs struct {
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

func (CloudProjectKubeNodepoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectKubeNodepoolArgs)(nil)).Elem()
}

type CloudProjectKubeNodepoolInput interface {
	pulumi.Input

	ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput
	ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput
}

func (*CloudProjectKubeNodepool) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectKubeNodepool)(nil))
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput {
	return i.ToCloudProjectKubeNodepoolOutputWithContext(context.Background())
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolOutput)
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolPtrOutput() CloudProjectKubeNodepoolPtrOutput {
	return i.ToCloudProjectKubeNodepoolPtrOutputWithContext(context.Background())
}

func (i *CloudProjectKubeNodepool) ToCloudProjectKubeNodepoolPtrOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolPtrOutput)
}

type CloudProjectKubeNodepoolPtrInput interface {
	pulumi.Input

	ToCloudProjectKubeNodepoolPtrOutput() CloudProjectKubeNodepoolPtrOutput
	ToCloudProjectKubeNodepoolPtrOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolPtrOutput
}

type cloudProjectKubeNodepoolPtrType CloudProjectKubeNodepoolArgs

func (*cloudProjectKubeNodepoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodepool)(nil))
}

func (i *cloudProjectKubeNodepoolPtrType) ToCloudProjectKubeNodepoolPtrOutput() CloudProjectKubeNodepoolPtrOutput {
	return i.ToCloudProjectKubeNodepoolPtrOutputWithContext(context.Background())
}

func (i *cloudProjectKubeNodepoolPtrType) ToCloudProjectKubeNodepoolPtrOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolPtrOutput)
}

// CloudProjectKubeNodepoolArrayInput is an input type that accepts CloudProjectKubeNodepoolArray and CloudProjectKubeNodepoolArrayOutput values.
// You can construct a concrete instance of `CloudProjectKubeNodepoolArrayInput` via:
//
//          CloudProjectKubeNodepoolArray{ CloudProjectKubeNodepoolArgs{...} }
type CloudProjectKubeNodepoolArrayInput interface {
	pulumi.Input

	ToCloudProjectKubeNodepoolArrayOutput() CloudProjectKubeNodepoolArrayOutput
	ToCloudProjectKubeNodepoolArrayOutputWithContext(context.Context) CloudProjectKubeNodepoolArrayOutput
}

type CloudProjectKubeNodepoolArray []CloudProjectKubeNodepoolInput

func (CloudProjectKubeNodepoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectKubeNodepool)(nil)).Elem()
}

func (i CloudProjectKubeNodepoolArray) ToCloudProjectKubeNodepoolArrayOutput() CloudProjectKubeNodepoolArrayOutput {
	return i.ToCloudProjectKubeNodepoolArrayOutputWithContext(context.Background())
}

func (i CloudProjectKubeNodepoolArray) ToCloudProjectKubeNodepoolArrayOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolArrayOutput)
}

// CloudProjectKubeNodepoolMapInput is an input type that accepts CloudProjectKubeNodepoolMap and CloudProjectKubeNodepoolMapOutput values.
// You can construct a concrete instance of `CloudProjectKubeNodepoolMapInput` via:
//
//          CloudProjectKubeNodepoolMap{ "key": CloudProjectKubeNodepoolArgs{...} }
type CloudProjectKubeNodepoolMapInput interface {
	pulumi.Input

	ToCloudProjectKubeNodepoolMapOutput() CloudProjectKubeNodepoolMapOutput
	ToCloudProjectKubeNodepoolMapOutputWithContext(context.Context) CloudProjectKubeNodepoolMapOutput
}

type CloudProjectKubeNodepoolMap map[string]CloudProjectKubeNodepoolInput

func (CloudProjectKubeNodepoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectKubeNodepool)(nil)).Elem()
}

func (i CloudProjectKubeNodepoolMap) ToCloudProjectKubeNodepoolMapOutput() CloudProjectKubeNodepoolMapOutput {
	return i.ToCloudProjectKubeNodepoolMapOutputWithContext(context.Background())
}

func (i CloudProjectKubeNodepoolMap) ToCloudProjectKubeNodepoolMapOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectKubeNodepoolMapOutput)
}

type CloudProjectKubeNodepoolOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodepoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectKubeNodepool)(nil))
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolOutput() CloudProjectKubeNodepoolOutput {
	return o
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolOutput {
	return o
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolPtrOutput() CloudProjectKubeNodepoolPtrOutput {
	return o.ToCloudProjectKubeNodepoolPtrOutputWithContext(context.Background())
}

func (o CloudProjectKubeNodepoolOutput) ToCloudProjectKubeNodepoolPtrOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProjectKubeNodepool) *CloudProjectKubeNodepool {
		return &v
	}).(CloudProjectKubeNodepoolPtrOutput)
}

type CloudProjectKubeNodepoolPtrOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodepoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectKubeNodepool)(nil))
}

func (o CloudProjectKubeNodepoolPtrOutput) ToCloudProjectKubeNodepoolPtrOutput() CloudProjectKubeNodepoolPtrOutput {
	return o
}

func (o CloudProjectKubeNodepoolPtrOutput) ToCloudProjectKubeNodepoolPtrOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolPtrOutput {
	return o
}

func (o CloudProjectKubeNodepoolPtrOutput) Elem() CloudProjectKubeNodepoolOutput {
	return o.ApplyT(func(v *CloudProjectKubeNodepool) CloudProjectKubeNodepool {
		if v != nil {
			return *v
		}
		var ret CloudProjectKubeNodepool
		return ret
	}).(CloudProjectKubeNodepoolOutput)
}

type CloudProjectKubeNodepoolArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodepoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudProjectKubeNodepool)(nil))
}

func (o CloudProjectKubeNodepoolArrayOutput) ToCloudProjectKubeNodepoolArrayOutput() CloudProjectKubeNodepoolArrayOutput {
	return o
}

func (o CloudProjectKubeNodepoolArrayOutput) ToCloudProjectKubeNodepoolArrayOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolArrayOutput {
	return o
}

func (o CloudProjectKubeNodepoolArrayOutput) Index(i pulumi.IntInput) CloudProjectKubeNodepoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudProjectKubeNodepool {
		return vs[0].([]CloudProjectKubeNodepool)[vs[1].(int)]
	}).(CloudProjectKubeNodepoolOutput)
}

type CloudProjectKubeNodepoolMapOutput struct{ *pulumi.OutputState }

func (CloudProjectKubeNodepoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudProjectKubeNodepool)(nil))
}

func (o CloudProjectKubeNodepoolMapOutput) ToCloudProjectKubeNodepoolMapOutput() CloudProjectKubeNodepoolMapOutput {
	return o
}

func (o CloudProjectKubeNodepoolMapOutput) ToCloudProjectKubeNodepoolMapOutputWithContext(ctx context.Context) CloudProjectKubeNodepoolMapOutput {
	return o
}

func (o CloudProjectKubeNodepoolMapOutput) MapIndex(k pulumi.StringInput) CloudProjectKubeNodepoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudProjectKubeNodepool {
		return vs[0].(map[string]CloudProjectKubeNodepool)[vs[1].(string)]
	}).(CloudProjectKubeNodepoolOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudProjectKubeNodepoolOutput{})
	pulumi.RegisterOutputType(CloudProjectKubeNodepoolPtrOutput{})
	pulumi.RegisterOutputType(CloudProjectKubeNodepoolArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectKubeNodepoolMapOutput{})
}
