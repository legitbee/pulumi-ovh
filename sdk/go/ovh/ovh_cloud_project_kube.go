// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_cloud_project_kube struct {
	pulumi.CustomResourceState

	ControlPlaneIsUpToDate pulumi.BoolOutput        `pulumi:"controlPlaneIsUpToDate"`
	IsUpToDate             pulumi.BoolOutput        `pulumi:"isUpToDate"`
	Kubeconfig             pulumi.StringOutput      `pulumi:"kubeconfig"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	NextUpgradeVersions    pulumi.StringArrayOutput `pulumi:"nextUpgradeVersions"`
	NodesUrl               pulumi.StringOutput      `pulumi:"nodesUrl"`
	PrivateNetworkId       pulumi.StringPtrOutput   `pulumi:"privateNetworkId"`
	Region                 pulumi.StringOutput      `pulumi:"region"`
	ServiceName            pulumi.StringOutput      `pulumi:"serviceName"`
	Status                 pulumi.StringOutput      `pulumi:"status"`
	UpdatePolicy           pulumi.StringOutput      `pulumi:"updatePolicy"`
	Url                    pulumi.StringOutput      `pulumi:"url"`
	Version                pulumi.StringOutput      `pulumi:"version"`
}

// NewOvh_cloud_project_kube registers a new resource with the given unique name, arguments, and options.
func NewOvh_cloud_project_kube(ctx *pulumi.Context,
	name string, args *Ovh_cloud_project_kubeArgs, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_kube, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_cloud_project_kube
	err := ctx.RegisterResource("ovh:index/ovh_cloud_project_kube:ovh_cloud_project_kube", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_cloud_project_kube gets an existing Ovh_cloud_project_kube resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_cloud_project_kube(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_cloud_project_kubeState, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_kube, error) {
	var resource Ovh_cloud_project_kube
	err := ctx.ReadResource("ovh:index/ovh_cloud_project_kube:ovh_cloud_project_kube", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_cloud_project_kube resources.
type ovh_cloud_project_kubeState struct {
	ControlPlaneIsUpToDate *bool    `pulumi:"controlPlaneIsUpToDate"`
	IsUpToDate             *bool    `pulumi:"isUpToDate"`
	Kubeconfig             *string  `pulumi:"kubeconfig"`
	Name                   *string  `pulumi:"name"`
	NextUpgradeVersions    []string `pulumi:"nextUpgradeVersions"`
	NodesUrl               *string  `pulumi:"nodesUrl"`
	PrivateNetworkId       *string  `pulumi:"privateNetworkId"`
	Region                 *string  `pulumi:"region"`
	ServiceName            *string  `pulumi:"serviceName"`
	Status                 *string  `pulumi:"status"`
	UpdatePolicy           *string  `pulumi:"updatePolicy"`
	Url                    *string  `pulumi:"url"`
	Version                *string  `pulumi:"version"`
}

type Ovh_cloud_project_kubeState struct {
	ControlPlaneIsUpToDate pulumi.BoolPtrInput
	IsUpToDate             pulumi.BoolPtrInput
	Kubeconfig             pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	NextUpgradeVersions    pulumi.StringArrayInput
	NodesUrl               pulumi.StringPtrInput
	PrivateNetworkId       pulumi.StringPtrInput
	Region                 pulumi.StringPtrInput
	ServiceName            pulumi.StringPtrInput
	Status                 pulumi.StringPtrInput
	UpdatePolicy           pulumi.StringPtrInput
	Url                    pulumi.StringPtrInput
	Version                pulumi.StringPtrInput
}

func (Ovh_cloud_project_kubeState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_kubeState)(nil)).Elem()
}

type ovh_cloud_project_kubeArgs struct {
	Name             *string `pulumi:"name"`
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	Region           string  `pulumi:"region"`
	ServiceName      string  `pulumi:"serviceName"`
	Version          *string `pulumi:"version"`
}

// The set of arguments for constructing a Ovh_cloud_project_kube resource.
type Ovh_cloud_project_kubeArgs struct {
	Name             pulumi.StringPtrInput
	PrivateNetworkId pulumi.StringPtrInput
	Region           pulumi.StringInput
	ServiceName      pulumi.StringInput
	Version          pulumi.StringPtrInput
}

func (Ovh_cloud_project_kubeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_kubeArgs)(nil)).Elem()
}

type Ovh_cloud_project_kubeInput interface {
	pulumi.Input

	ToOvh_cloud_project_kubeOutput() Ovh_cloud_project_kubeOutput
	ToOvh_cloud_project_kubeOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeOutput
}

func (*Ovh_cloud_project_kube) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_kube)(nil))
}

func (i *Ovh_cloud_project_kube) ToOvh_cloud_project_kubeOutput() Ovh_cloud_project_kubeOutput {
	return i.ToOvh_cloud_project_kubeOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_kube) ToOvh_cloud_project_kubeOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kubeOutput)
}

func (i *Ovh_cloud_project_kube) ToOvh_cloud_project_kubePtrOutput() Ovh_cloud_project_kubePtrOutput {
	return i.ToOvh_cloud_project_kubePtrOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_kube) ToOvh_cloud_project_kubePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kubePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kubePtrOutput)
}

type Ovh_cloud_project_kubePtrInput interface {
	pulumi.Input

	ToOvh_cloud_project_kubePtrOutput() Ovh_cloud_project_kubePtrOutput
	ToOvh_cloud_project_kubePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kubePtrOutput
}

type ovh_cloud_project_kubePtrType Ovh_cloud_project_kubeArgs

func (*ovh_cloud_project_kubePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_kube)(nil))
}

func (i *ovh_cloud_project_kubePtrType) ToOvh_cloud_project_kubePtrOutput() Ovh_cloud_project_kubePtrOutput {
	return i.ToOvh_cloud_project_kubePtrOutputWithContext(context.Background())
}

func (i *ovh_cloud_project_kubePtrType) ToOvh_cloud_project_kubePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kubePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kubePtrOutput)
}

// Ovh_cloud_project_kubeArrayInput is an input type that accepts Ovh_cloud_project_kubeArray and Ovh_cloud_project_kubeArrayOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_kubeArrayInput` via:
//
//          Ovh_cloud_project_kubeArray{ Ovh_cloud_project_kubeArgs{...} }
type Ovh_cloud_project_kubeArrayInput interface {
	pulumi.Input

	ToOvh_cloud_project_kubeArrayOutput() Ovh_cloud_project_kubeArrayOutput
	ToOvh_cloud_project_kubeArrayOutputWithContext(context.Context) Ovh_cloud_project_kubeArrayOutput
}

type Ovh_cloud_project_kubeArray []Ovh_cloud_project_kubeInput

func (Ovh_cloud_project_kubeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_cloud_project_kube)(nil)).Elem()
}

func (i Ovh_cloud_project_kubeArray) ToOvh_cloud_project_kubeArrayOutput() Ovh_cloud_project_kubeArrayOutput {
	return i.ToOvh_cloud_project_kubeArrayOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_kubeArray) ToOvh_cloud_project_kubeArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kubeArrayOutput)
}

// Ovh_cloud_project_kubeMapInput is an input type that accepts Ovh_cloud_project_kubeMap and Ovh_cloud_project_kubeMapOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_kubeMapInput` via:
//
//          Ovh_cloud_project_kubeMap{ "key": Ovh_cloud_project_kubeArgs{...} }
type Ovh_cloud_project_kubeMapInput interface {
	pulumi.Input

	ToOvh_cloud_project_kubeMapOutput() Ovh_cloud_project_kubeMapOutput
	ToOvh_cloud_project_kubeMapOutputWithContext(context.Context) Ovh_cloud_project_kubeMapOutput
}

type Ovh_cloud_project_kubeMap map[string]Ovh_cloud_project_kubeInput

func (Ovh_cloud_project_kubeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_cloud_project_kube)(nil)).Elem()
}

func (i Ovh_cloud_project_kubeMap) ToOvh_cloud_project_kubeMapOutput() Ovh_cloud_project_kubeMapOutput {
	return i.ToOvh_cloud_project_kubeMapOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_kubeMap) ToOvh_cloud_project_kubeMapOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_kubeMapOutput)
}

type Ovh_cloud_project_kubeOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kubeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_kube)(nil))
}

func (o Ovh_cloud_project_kubeOutput) ToOvh_cloud_project_kubeOutput() Ovh_cloud_project_kubeOutput {
	return o
}

func (o Ovh_cloud_project_kubeOutput) ToOvh_cloud_project_kubeOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeOutput {
	return o
}

func (o Ovh_cloud_project_kubeOutput) ToOvh_cloud_project_kubePtrOutput() Ovh_cloud_project_kubePtrOutput {
	return o.ToOvh_cloud_project_kubePtrOutputWithContext(context.Background())
}

func (o Ovh_cloud_project_kubeOutput) ToOvh_cloud_project_kubePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kubePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_cloud_project_kube) *Ovh_cloud_project_kube {
		return &v
	}).(Ovh_cloud_project_kubePtrOutput)
}

type Ovh_cloud_project_kubePtrOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kubePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_kube)(nil))
}

func (o Ovh_cloud_project_kubePtrOutput) ToOvh_cloud_project_kubePtrOutput() Ovh_cloud_project_kubePtrOutput {
	return o
}

func (o Ovh_cloud_project_kubePtrOutput) ToOvh_cloud_project_kubePtrOutputWithContext(ctx context.Context) Ovh_cloud_project_kubePtrOutput {
	return o
}

func (o Ovh_cloud_project_kubePtrOutput) Elem() Ovh_cloud_project_kubeOutput {
	return o.ApplyT(func(v *Ovh_cloud_project_kube) Ovh_cloud_project_kube {
		if v != nil {
			return *v
		}
		var ret Ovh_cloud_project_kube
		return ret
	}).(Ovh_cloud_project_kubeOutput)
}

type Ovh_cloud_project_kubeArrayOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kubeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_cloud_project_kube)(nil))
}

func (o Ovh_cloud_project_kubeArrayOutput) ToOvh_cloud_project_kubeArrayOutput() Ovh_cloud_project_kubeArrayOutput {
	return o
}

func (o Ovh_cloud_project_kubeArrayOutput) ToOvh_cloud_project_kubeArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeArrayOutput {
	return o
}

func (o Ovh_cloud_project_kubeArrayOutput) Index(i pulumi.IntInput) Ovh_cloud_project_kubeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_cloud_project_kube {
		return vs[0].([]Ovh_cloud_project_kube)[vs[1].(int)]
	}).(Ovh_cloud_project_kubeOutput)
}

type Ovh_cloud_project_kubeMapOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_kubeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_cloud_project_kube)(nil))
}

func (o Ovh_cloud_project_kubeMapOutput) ToOvh_cloud_project_kubeMapOutput() Ovh_cloud_project_kubeMapOutput {
	return o
}

func (o Ovh_cloud_project_kubeMapOutput) ToOvh_cloud_project_kubeMapOutputWithContext(ctx context.Context) Ovh_cloud_project_kubeMapOutput {
	return o
}

func (o Ovh_cloud_project_kubeMapOutput) MapIndex(k pulumi.StringInput) Ovh_cloud_project_kubeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_cloud_project_kube {
		return vs[0].(map[string]Ovh_cloud_project_kube)[vs[1].(string)]
	}).(Ovh_cloud_project_kubeOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_cloud_project_kubeOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kubePtrOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kubeArrayOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_kubeMapOutput{})
}
