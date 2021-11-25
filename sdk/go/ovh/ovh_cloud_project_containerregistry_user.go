// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_cloud_project_containerregistry_user struct {
	pulumi.CustomResourceState

	// User email.
	Email pulumi.StringOutput `pulumi:"email"`
	// Registry name
	Login pulumi.StringOutput `pulumi:"login"`
	// User password
	Password pulumi.StringOutput `pulumi:"password"`
	// RegistryID
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// User name
	User pulumi.StringOutput `pulumi:"user"`
}

// NewOvh_cloud_project_containerregistry_user registers a new resource with the given unique name, arguments, and options.
func NewOvh_cloud_project_containerregistry_user(ctx *pulumi.Context,
	name string, args *Ovh_cloud_project_containerregistry_userArgs, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_containerregistry_user, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource Ovh_cloud_project_containerregistry_user
	err := ctx.RegisterResource("ovh:index/ovh_cloud_project_containerregistry_user:ovh_cloud_project_containerregistry_user", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_cloud_project_containerregistry_user gets an existing Ovh_cloud_project_containerregistry_user resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_cloud_project_containerregistry_user(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_cloud_project_containerregistry_userState, opts ...pulumi.ResourceOption) (*Ovh_cloud_project_containerregistry_user, error) {
	var resource Ovh_cloud_project_containerregistry_user
	err := ctx.ReadResource("ovh:index/ovh_cloud_project_containerregistry_user:ovh_cloud_project_containerregistry_user", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_cloud_project_containerregistry_user resources.
type ovh_cloud_project_containerregistry_userState struct {
	// User email.
	Email *string `pulumi:"email"`
	// Registry name
	Login *string `pulumi:"login"`
	// User password
	Password *string `pulumi:"password"`
	// RegistryID
	RegistryId *string `pulumi:"registryId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// User name
	User *string `pulumi:"user"`
}

type Ovh_cloud_project_containerregistry_userState struct {
	// User email.
	Email pulumi.StringPtrInput
	// Registry name
	Login pulumi.StringPtrInput
	// User password
	Password pulumi.StringPtrInput
	// RegistryID
	RegistryId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// User name
	User pulumi.StringPtrInput
}

func (Ovh_cloud_project_containerregistry_userState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_containerregistry_userState)(nil)).Elem()
}

type ovh_cloud_project_containerregistry_userArgs struct {
	// User email.
	Email string `pulumi:"email"`
	// Registry name
	Login string `pulumi:"login"`
	// RegistryID
	RegistryId string `pulumi:"registryId"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Ovh_cloud_project_containerregistry_user resource.
type Ovh_cloud_project_containerregistry_userArgs struct {
	// User email.
	Email pulumi.StringInput
	// Registry name
	Login pulumi.StringInput
	// RegistryID
	RegistryId pulumi.StringInput
	// Service name
	ServiceName pulumi.StringInput
}

func (Ovh_cloud_project_containerregistry_userArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_cloud_project_containerregistry_userArgs)(nil)).Elem()
}

type Ovh_cloud_project_containerregistry_userInput interface {
	pulumi.Input

	ToOvh_cloud_project_containerregistry_userOutput() Ovh_cloud_project_containerregistry_userOutput
	ToOvh_cloud_project_containerregistry_userOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userOutput
}

func (*Ovh_cloud_project_containerregistry_user) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_containerregistry_user)(nil))
}

func (i *Ovh_cloud_project_containerregistry_user) ToOvh_cloud_project_containerregistry_userOutput() Ovh_cloud_project_containerregistry_userOutput {
	return i.ToOvh_cloud_project_containerregistry_userOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_containerregistry_user) ToOvh_cloud_project_containerregistry_userOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_containerregistry_userOutput)
}

func (i *Ovh_cloud_project_containerregistry_user) ToOvh_cloud_project_containerregistry_userPtrOutput() Ovh_cloud_project_containerregistry_userPtrOutput {
	return i.ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(context.Background())
}

func (i *Ovh_cloud_project_containerregistry_user) ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_containerregistry_userPtrOutput)
}

type Ovh_cloud_project_containerregistry_userPtrInput interface {
	pulumi.Input

	ToOvh_cloud_project_containerregistry_userPtrOutput() Ovh_cloud_project_containerregistry_userPtrOutput
	ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userPtrOutput
}

type ovh_cloud_project_containerregistry_userPtrType Ovh_cloud_project_containerregistry_userArgs

func (*ovh_cloud_project_containerregistry_userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_containerregistry_user)(nil))
}

func (i *ovh_cloud_project_containerregistry_userPtrType) ToOvh_cloud_project_containerregistry_userPtrOutput() Ovh_cloud_project_containerregistry_userPtrOutput {
	return i.ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(context.Background())
}

func (i *ovh_cloud_project_containerregistry_userPtrType) ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_containerregistry_userPtrOutput)
}

// Ovh_cloud_project_containerregistry_userArrayInput is an input type that accepts Ovh_cloud_project_containerregistry_userArray and Ovh_cloud_project_containerregistry_userArrayOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_containerregistry_userArrayInput` via:
//
//          Ovh_cloud_project_containerregistry_userArray{ Ovh_cloud_project_containerregistry_userArgs{...} }
type Ovh_cloud_project_containerregistry_userArrayInput interface {
	pulumi.Input

	ToOvh_cloud_project_containerregistry_userArrayOutput() Ovh_cloud_project_containerregistry_userArrayOutput
	ToOvh_cloud_project_containerregistry_userArrayOutputWithContext(context.Context) Ovh_cloud_project_containerregistry_userArrayOutput
}

type Ovh_cloud_project_containerregistry_userArray []Ovh_cloud_project_containerregistry_userInput

func (Ovh_cloud_project_containerregistry_userArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_cloud_project_containerregistry_user)(nil)).Elem()
}

func (i Ovh_cloud_project_containerregistry_userArray) ToOvh_cloud_project_containerregistry_userArrayOutput() Ovh_cloud_project_containerregistry_userArrayOutput {
	return i.ToOvh_cloud_project_containerregistry_userArrayOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_containerregistry_userArray) ToOvh_cloud_project_containerregistry_userArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_containerregistry_userArrayOutput)
}

// Ovh_cloud_project_containerregistry_userMapInput is an input type that accepts Ovh_cloud_project_containerregistry_userMap and Ovh_cloud_project_containerregistry_userMapOutput values.
// You can construct a concrete instance of `Ovh_cloud_project_containerregistry_userMapInput` via:
//
//          Ovh_cloud_project_containerregistry_userMap{ "key": Ovh_cloud_project_containerregistry_userArgs{...} }
type Ovh_cloud_project_containerregistry_userMapInput interface {
	pulumi.Input

	ToOvh_cloud_project_containerregistry_userMapOutput() Ovh_cloud_project_containerregistry_userMapOutput
	ToOvh_cloud_project_containerregistry_userMapOutputWithContext(context.Context) Ovh_cloud_project_containerregistry_userMapOutput
}

type Ovh_cloud_project_containerregistry_userMap map[string]Ovh_cloud_project_containerregistry_userInput

func (Ovh_cloud_project_containerregistry_userMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_cloud_project_containerregistry_user)(nil)).Elem()
}

func (i Ovh_cloud_project_containerregistry_userMap) ToOvh_cloud_project_containerregistry_userMapOutput() Ovh_cloud_project_containerregistry_userMapOutput {
	return i.ToOvh_cloud_project_containerregistry_userMapOutputWithContext(context.Background())
}

func (i Ovh_cloud_project_containerregistry_userMap) ToOvh_cloud_project_containerregistry_userMapOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_cloud_project_containerregistry_userMapOutput)
}

type Ovh_cloud_project_containerregistry_userOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_containerregistry_userOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_cloud_project_containerregistry_user)(nil))
}

func (o Ovh_cloud_project_containerregistry_userOutput) ToOvh_cloud_project_containerregistry_userOutput() Ovh_cloud_project_containerregistry_userOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userOutput) ToOvh_cloud_project_containerregistry_userOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userOutput) ToOvh_cloud_project_containerregistry_userPtrOutput() Ovh_cloud_project_containerregistry_userPtrOutput {
	return o.ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(context.Background())
}

func (o Ovh_cloud_project_containerregistry_userOutput) ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_cloud_project_containerregistry_user) *Ovh_cloud_project_containerregistry_user {
		return &v
	}).(Ovh_cloud_project_containerregistry_userPtrOutput)
}

type Ovh_cloud_project_containerregistry_userPtrOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_containerregistry_userPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_cloud_project_containerregistry_user)(nil))
}

func (o Ovh_cloud_project_containerregistry_userPtrOutput) ToOvh_cloud_project_containerregistry_userPtrOutput() Ovh_cloud_project_containerregistry_userPtrOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userPtrOutput) ToOvh_cloud_project_containerregistry_userPtrOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userPtrOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userPtrOutput) Elem() Ovh_cloud_project_containerregistry_userOutput {
	return o.ApplyT(func(v *Ovh_cloud_project_containerregistry_user) Ovh_cloud_project_containerregistry_user {
		if v != nil {
			return *v
		}
		var ret Ovh_cloud_project_containerregistry_user
		return ret
	}).(Ovh_cloud_project_containerregistry_userOutput)
}

type Ovh_cloud_project_containerregistry_userArrayOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_containerregistry_userArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_cloud_project_containerregistry_user)(nil))
}

func (o Ovh_cloud_project_containerregistry_userArrayOutput) ToOvh_cloud_project_containerregistry_userArrayOutput() Ovh_cloud_project_containerregistry_userArrayOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userArrayOutput) ToOvh_cloud_project_containerregistry_userArrayOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userArrayOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userArrayOutput) Index(i pulumi.IntInput) Ovh_cloud_project_containerregistry_userOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_cloud_project_containerregistry_user {
		return vs[0].([]Ovh_cloud_project_containerregistry_user)[vs[1].(int)]
	}).(Ovh_cloud_project_containerregistry_userOutput)
}

type Ovh_cloud_project_containerregistry_userMapOutput struct{ *pulumi.OutputState }

func (Ovh_cloud_project_containerregistry_userMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_cloud_project_containerregistry_user)(nil))
}

func (o Ovh_cloud_project_containerregistry_userMapOutput) ToOvh_cloud_project_containerregistry_userMapOutput() Ovh_cloud_project_containerregistry_userMapOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userMapOutput) ToOvh_cloud_project_containerregistry_userMapOutputWithContext(ctx context.Context) Ovh_cloud_project_containerregistry_userMapOutput {
	return o
}

func (o Ovh_cloud_project_containerregistry_userMapOutput) MapIndex(k pulumi.StringInput) Ovh_cloud_project_containerregistry_userOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_cloud_project_containerregistry_user {
		return vs[0].(map[string]Ovh_cloud_project_containerregistry_user)[vs[1].(string)]
	}).(Ovh_cloud_project_containerregistry_userOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_cloud_project_containerregistry_userOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_containerregistry_userPtrOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_containerregistry_userArrayOutput{})
	pulumi.RegisterOutputType(Ovh_cloud_project_containerregistry_userMapOutput{})
}
