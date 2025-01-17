// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Ovh_me_identity_user struct {
	pulumi.CustomResourceState

	// Creation date of this user
	Creation pulumi.StringOutput `pulumi:"creation"`
	// User description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User's email
	Email pulumi.StringOutput `pulumi:"email"`
	// User's group
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// Last update of this user
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// User's login suffix
	Login pulumi.StringOutput `pulumi:"login"`
	// User's password
	Password pulumi.StringOutput `pulumi:"password"`
	// When the user changed his password for the last time
	PasswordLastUpdate pulumi.StringOutput `pulumi:"passwordLastUpdate"`
	// Current user's status
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewOvh_me_identity_user registers a new resource with the given unique name, arguments, and options.
func NewOvh_me_identity_user(ctx *pulumi.Context,
	name string, args *Ovh_me_identity_userArgs, opts ...pulumi.ResourceOption) (*Ovh_me_identity_user, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	var resource Ovh_me_identity_user
	err := ctx.RegisterResource("ovh:index/ovh_me_identity_user:ovh_me_identity_user", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOvh_me_identity_user gets an existing Ovh_me_identity_user resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOvh_me_identity_user(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ovh_me_identity_userState, opts ...pulumi.ResourceOption) (*Ovh_me_identity_user, error) {
	var resource Ovh_me_identity_user
	err := ctx.ReadResource("ovh:index/ovh_me_identity_user:ovh_me_identity_user", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ovh_me_identity_user resources.
type ovh_me_identity_userState struct {
	// Creation date of this user
	Creation *string `pulumi:"creation"`
	// User description
	Description *string `pulumi:"description"`
	// User's email
	Email *string `pulumi:"email"`
	// User's group
	Group *string `pulumi:"group"`
	// Last update of this user
	LastUpdate *string `pulumi:"lastUpdate"`
	// User's login suffix
	Login *string `pulumi:"login"`
	// User's password
	Password *string `pulumi:"password"`
	// When the user changed his password for the last time
	PasswordLastUpdate *string `pulumi:"passwordLastUpdate"`
	// Current user's status
	Status *string `pulumi:"status"`
}

type Ovh_me_identity_userState struct {
	// Creation date of this user
	Creation pulumi.StringPtrInput
	// User description
	Description pulumi.StringPtrInput
	// User's email
	Email pulumi.StringPtrInput
	// User's group
	Group pulumi.StringPtrInput
	// Last update of this user
	LastUpdate pulumi.StringPtrInput
	// User's login suffix
	Login pulumi.StringPtrInput
	// User's password
	Password pulumi.StringPtrInput
	// When the user changed his password for the last time
	PasswordLastUpdate pulumi.StringPtrInput
	// Current user's status
	Status pulumi.StringPtrInput
}

func (Ovh_me_identity_userState) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_me_identity_userState)(nil)).Elem()
}

type ovh_me_identity_userArgs struct {
	// User description
	Description *string `pulumi:"description"`
	// User's email
	Email string `pulumi:"email"`
	// User's group
	Group *string `pulumi:"group"`
	// User's login suffix
	Login string `pulumi:"login"`
	// User's password
	Password string `pulumi:"password"`
}

// The set of arguments for constructing a Ovh_me_identity_user resource.
type Ovh_me_identity_userArgs struct {
	// User description
	Description pulumi.StringPtrInput
	// User's email
	Email pulumi.StringInput
	// User's group
	Group pulumi.StringPtrInput
	// User's login suffix
	Login pulumi.StringInput
	// User's password
	Password pulumi.StringInput
}

func (Ovh_me_identity_userArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ovh_me_identity_userArgs)(nil)).Elem()
}

type Ovh_me_identity_userInput interface {
	pulumi.Input

	ToOvh_me_identity_userOutput() Ovh_me_identity_userOutput
	ToOvh_me_identity_userOutputWithContext(ctx context.Context) Ovh_me_identity_userOutput
}

func (*Ovh_me_identity_user) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_me_identity_user)(nil))
}

func (i *Ovh_me_identity_user) ToOvh_me_identity_userOutput() Ovh_me_identity_userOutput {
	return i.ToOvh_me_identity_userOutputWithContext(context.Background())
}

func (i *Ovh_me_identity_user) ToOvh_me_identity_userOutputWithContext(ctx context.Context) Ovh_me_identity_userOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_identity_userOutput)
}

func (i *Ovh_me_identity_user) ToOvh_me_identity_userPtrOutput() Ovh_me_identity_userPtrOutput {
	return i.ToOvh_me_identity_userPtrOutputWithContext(context.Background())
}

func (i *Ovh_me_identity_user) ToOvh_me_identity_userPtrOutputWithContext(ctx context.Context) Ovh_me_identity_userPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_identity_userPtrOutput)
}

type Ovh_me_identity_userPtrInput interface {
	pulumi.Input

	ToOvh_me_identity_userPtrOutput() Ovh_me_identity_userPtrOutput
	ToOvh_me_identity_userPtrOutputWithContext(ctx context.Context) Ovh_me_identity_userPtrOutput
}

type ovh_me_identity_userPtrType Ovh_me_identity_userArgs

func (*ovh_me_identity_userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_me_identity_user)(nil))
}

func (i *ovh_me_identity_userPtrType) ToOvh_me_identity_userPtrOutput() Ovh_me_identity_userPtrOutput {
	return i.ToOvh_me_identity_userPtrOutputWithContext(context.Background())
}

func (i *ovh_me_identity_userPtrType) ToOvh_me_identity_userPtrOutputWithContext(ctx context.Context) Ovh_me_identity_userPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_identity_userPtrOutput)
}

// Ovh_me_identity_userArrayInput is an input type that accepts Ovh_me_identity_userArray and Ovh_me_identity_userArrayOutput values.
// You can construct a concrete instance of `Ovh_me_identity_userArrayInput` via:
//
//          Ovh_me_identity_userArray{ Ovh_me_identity_userArgs{...} }
type Ovh_me_identity_userArrayInput interface {
	pulumi.Input

	ToOvh_me_identity_userArrayOutput() Ovh_me_identity_userArrayOutput
	ToOvh_me_identity_userArrayOutputWithContext(context.Context) Ovh_me_identity_userArrayOutput
}

type Ovh_me_identity_userArray []Ovh_me_identity_userInput

func (Ovh_me_identity_userArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ovh_me_identity_user)(nil)).Elem()
}

func (i Ovh_me_identity_userArray) ToOvh_me_identity_userArrayOutput() Ovh_me_identity_userArrayOutput {
	return i.ToOvh_me_identity_userArrayOutputWithContext(context.Background())
}

func (i Ovh_me_identity_userArray) ToOvh_me_identity_userArrayOutputWithContext(ctx context.Context) Ovh_me_identity_userArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_identity_userArrayOutput)
}

// Ovh_me_identity_userMapInput is an input type that accepts Ovh_me_identity_userMap and Ovh_me_identity_userMapOutput values.
// You can construct a concrete instance of `Ovh_me_identity_userMapInput` via:
//
//          Ovh_me_identity_userMap{ "key": Ovh_me_identity_userArgs{...} }
type Ovh_me_identity_userMapInput interface {
	pulumi.Input

	ToOvh_me_identity_userMapOutput() Ovh_me_identity_userMapOutput
	ToOvh_me_identity_userMapOutputWithContext(context.Context) Ovh_me_identity_userMapOutput
}

type Ovh_me_identity_userMap map[string]Ovh_me_identity_userInput

func (Ovh_me_identity_userMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ovh_me_identity_user)(nil)).Elem()
}

func (i Ovh_me_identity_userMap) ToOvh_me_identity_userMapOutput() Ovh_me_identity_userMapOutput {
	return i.ToOvh_me_identity_userMapOutputWithContext(context.Background())
}

func (i Ovh_me_identity_userMap) ToOvh_me_identity_userMapOutputWithContext(ctx context.Context) Ovh_me_identity_userMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ovh_me_identity_userMapOutput)
}

type Ovh_me_identity_userOutput struct{ *pulumi.OutputState }

func (Ovh_me_identity_userOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ovh_me_identity_user)(nil))
}

func (o Ovh_me_identity_userOutput) ToOvh_me_identity_userOutput() Ovh_me_identity_userOutput {
	return o
}

func (o Ovh_me_identity_userOutput) ToOvh_me_identity_userOutputWithContext(ctx context.Context) Ovh_me_identity_userOutput {
	return o
}

func (o Ovh_me_identity_userOutput) ToOvh_me_identity_userPtrOutput() Ovh_me_identity_userPtrOutput {
	return o.ToOvh_me_identity_userPtrOutputWithContext(context.Background())
}

func (o Ovh_me_identity_userOutput) ToOvh_me_identity_userPtrOutputWithContext(ctx context.Context) Ovh_me_identity_userPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ovh_me_identity_user) *Ovh_me_identity_user {
		return &v
	}).(Ovh_me_identity_userPtrOutput)
}

type Ovh_me_identity_userPtrOutput struct{ *pulumi.OutputState }

func (Ovh_me_identity_userPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ovh_me_identity_user)(nil))
}

func (o Ovh_me_identity_userPtrOutput) ToOvh_me_identity_userPtrOutput() Ovh_me_identity_userPtrOutput {
	return o
}

func (o Ovh_me_identity_userPtrOutput) ToOvh_me_identity_userPtrOutputWithContext(ctx context.Context) Ovh_me_identity_userPtrOutput {
	return o
}

func (o Ovh_me_identity_userPtrOutput) Elem() Ovh_me_identity_userOutput {
	return o.ApplyT(func(v *Ovh_me_identity_user) Ovh_me_identity_user {
		if v != nil {
			return *v
		}
		var ret Ovh_me_identity_user
		return ret
	}).(Ovh_me_identity_userOutput)
}

type Ovh_me_identity_userArrayOutput struct{ *pulumi.OutputState }

func (Ovh_me_identity_userArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ovh_me_identity_user)(nil))
}

func (o Ovh_me_identity_userArrayOutput) ToOvh_me_identity_userArrayOutput() Ovh_me_identity_userArrayOutput {
	return o
}

func (o Ovh_me_identity_userArrayOutput) ToOvh_me_identity_userArrayOutputWithContext(ctx context.Context) Ovh_me_identity_userArrayOutput {
	return o
}

func (o Ovh_me_identity_userArrayOutput) Index(i pulumi.IntInput) Ovh_me_identity_userOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ovh_me_identity_user {
		return vs[0].([]Ovh_me_identity_user)[vs[1].(int)]
	}).(Ovh_me_identity_userOutput)
}

type Ovh_me_identity_userMapOutput struct{ *pulumi.OutputState }

func (Ovh_me_identity_userMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ovh_me_identity_user)(nil))
}

func (o Ovh_me_identity_userMapOutput) ToOvh_me_identity_userMapOutput() Ovh_me_identity_userMapOutput {
	return o
}

func (o Ovh_me_identity_userMapOutput) ToOvh_me_identity_userMapOutputWithContext(ctx context.Context) Ovh_me_identity_userMapOutput {
	return o
}

func (o Ovh_me_identity_userMapOutput) MapIndex(k pulumi.StringInput) Ovh_me_identity_userOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ovh_me_identity_user {
		return vs[0].(map[string]Ovh_me_identity_user)[vs[1].(string)]
	}).(Ovh_me_identity_userOutput)
}

func init() {
	pulumi.RegisterOutputType(Ovh_me_identity_userOutput{})
	pulumi.RegisterOutputType(Ovh_me_identity_userPtrOutput{})
	pulumi.RegisterOutputType(Ovh_me_identity_userArrayOutput{})
	pulumi.RegisterOutputType(Ovh_me_identity_userMapOutput{})
}
