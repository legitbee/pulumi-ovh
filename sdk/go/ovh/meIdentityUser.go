// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MeIdentityUser struct {
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

// NewMeIdentityUser registers a new resource with the given unique name, arguments, and options.
func NewMeIdentityUser(ctx *pulumi.Context,
	name string, args *MeIdentityUserArgs, opts ...pulumi.ResourceOption) (*MeIdentityUser, error) {
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
	var resource MeIdentityUser
	err := ctx.RegisterResource("ovh:index/meIdentityUser:MeIdentityUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMeIdentityUser gets an existing MeIdentityUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMeIdentityUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeIdentityUserState, opts ...pulumi.ResourceOption) (*MeIdentityUser, error) {
	var resource MeIdentityUser
	err := ctx.ReadResource("ovh:index/meIdentityUser:MeIdentityUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MeIdentityUser resources.
type meIdentityUserState struct {
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

type MeIdentityUserState struct {
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

func (MeIdentityUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*meIdentityUserState)(nil)).Elem()
}

type meIdentityUserArgs struct {
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

// The set of arguments for constructing a MeIdentityUser resource.
type MeIdentityUserArgs struct {
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

func (MeIdentityUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meIdentityUserArgs)(nil)).Elem()
}

type MeIdentityUserInput interface {
	pulumi.Input

	ToMeIdentityUserOutput() MeIdentityUserOutput
	ToMeIdentityUserOutputWithContext(ctx context.Context) MeIdentityUserOutput
}

func (*MeIdentityUser) ElementType() reflect.Type {
	return reflect.TypeOf((*MeIdentityUser)(nil))
}

func (i *MeIdentityUser) ToMeIdentityUserOutput() MeIdentityUserOutput {
	return i.ToMeIdentityUserOutputWithContext(context.Background())
}

func (i *MeIdentityUser) ToMeIdentityUserOutputWithContext(ctx context.Context) MeIdentityUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeIdentityUserOutput)
}

func (i *MeIdentityUser) ToMeIdentityUserPtrOutput() MeIdentityUserPtrOutput {
	return i.ToMeIdentityUserPtrOutputWithContext(context.Background())
}

func (i *MeIdentityUser) ToMeIdentityUserPtrOutputWithContext(ctx context.Context) MeIdentityUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeIdentityUserPtrOutput)
}

type MeIdentityUserPtrInput interface {
	pulumi.Input

	ToMeIdentityUserPtrOutput() MeIdentityUserPtrOutput
	ToMeIdentityUserPtrOutputWithContext(ctx context.Context) MeIdentityUserPtrOutput
}

type meIdentityUserPtrType MeIdentityUserArgs

func (*meIdentityUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MeIdentityUser)(nil))
}

func (i *meIdentityUserPtrType) ToMeIdentityUserPtrOutput() MeIdentityUserPtrOutput {
	return i.ToMeIdentityUserPtrOutputWithContext(context.Background())
}

func (i *meIdentityUserPtrType) ToMeIdentityUserPtrOutputWithContext(ctx context.Context) MeIdentityUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeIdentityUserPtrOutput)
}

// MeIdentityUserArrayInput is an input type that accepts MeIdentityUserArray and MeIdentityUserArrayOutput values.
// You can construct a concrete instance of `MeIdentityUserArrayInput` via:
//
//          MeIdentityUserArray{ MeIdentityUserArgs{...} }
type MeIdentityUserArrayInput interface {
	pulumi.Input

	ToMeIdentityUserArrayOutput() MeIdentityUserArrayOutput
	ToMeIdentityUserArrayOutputWithContext(context.Context) MeIdentityUserArrayOutput
}

type MeIdentityUserArray []MeIdentityUserInput

func (MeIdentityUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MeIdentityUser)(nil)).Elem()
}

func (i MeIdentityUserArray) ToMeIdentityUserArrayOutput() MeIdentityUserArrayOutput {
	return i.ToMeIdentityUserArrayOutputWithContext(context.Background())
}

func (i MeIdentityUserArray) ToMeIdentityUserArrayOutputWithContext(ctx context.Context) MeIdentityUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeIdentityUserArrayOutput)
}

// MeIdentityUserMapInput is an input type that accepts MeIdentityUserMap and MeIdentityUserMapOutput values.
// You can construct a concrete instance of `MeIdentityUserMapInput` via:
//
//          MeIdentityUserMap{ "key": MeIdentityUserArgs{...} }
type MeIdentityUserMapInput interface {
	pulumi.Input

	ToMeIdentityUserMapOutput() MeIdentityUserMapOutput
	ToMeIdentityUserMapOutputWithContext(context.Context) MeIdentityUserMapOutput
}

type MeIdentityUserMap map[string]MeIdentityUserInput

func (MeIdentityUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MeIdentityUser)(nil)).Elem()
}

func (i MeIdentityUserMap) ToMeIdentityUserMapOutput() MeIdentityUserMapOutput {
	return i.ToMeIdentityUserMapOutputWithContext(context.Background())
}

func (i MeIdentityUserMap) ToMeIdentityUserMapOutputWithContext(ctx context.Context) MeIdentityUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeIdentityUserMapOutput)
}

type MeIdentityUserOutput struct{ *pulumi.OutputState }

func (MeIdentityUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeIdentityUser)(nil))
}

func (o MeIdentityUserOutput) ToMeIdentityUserOutput() MeIdentityUserOutput {
	return o
}

func (o MeIdentityUserOutput) ToMeIdentityUserOutputWithContext(ctx context.Context) MeIdentityUserOutput {
	return o
}

func (o MeIdentityUserOutput) ToMeIdentityUserPtrOutput() MeIdentityUserPtrOutput {
	return o.ToMeIdentityUserPtrOutputWithContext(context.Background())
}

func (o MeIdentityUserOutput) ToMeIdentityUserPtrOutputWithContext(ctx context.Context) MeIdentityUserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MeIdentityUser) *MeIdentityUser {
		return &v
	}).(MeIdentityUserPtrOutput)
}

type MeIdentityUserPtrOutput struct{ *pulumi.OutputState }

func (MeIdentityUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MeIdentityUser)(nil))
}

func (o MeIdentityUserPtrOutput) ToMeIdentityUserPtrOutput() MeIdentityUserPtrOutput {
	return o
}

func (o MeIdentityUserPtrOutput) ToMeIdentityUserPtrOutputWithContext(ctx context.Context) MeIdentityUserPtrOutput {
	return o
}

func (o MeIdentityUserPtrOutput) Elem() MeIdentityUserOutput {
	return o.ApplyT(func(v *MeIdentityUser) MeIdentityUser {
		if v != nil {
			return *v
		}
		var ret MeIdentityUser
		return ret
	}).(MeIdentityUserOutput)
}

type MeIdentityUserArrayOutput struct{ *pulumi.OutputState }

func (MeIdentityUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MeIdentityUser)(nil))
}

func (o MeIdentityUserArrayOutput) ToMeIdentityUserArrayOutput() MeIdentityUserArrayOutput {
	return o
}

func (o MeIdentityUserArrayOutput) ToMeIdentityUserArrayOutputWithContext(ctx context.Context) MeIdentityUserArrayOutput {
	return o
}

func (o MeIdentityUserArrayOutput) Index(i pulumi.IntInput) MeIdentityUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MeIdentityUser {
		return vs[0].([]MeIdentityUser)[vs[1].(int)]
	}).(MeIdentityUserOutput)
}

type MeIdentityUserMapOutput struct{ *pulumi.OutputState }

func (MeIdentityUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MeIdentityUser)(nil))
}

func (o MeIdentityUserMapOutput) ToMeIdentityUserMapOutput() MeIdentityUserMapOutput {
	return o
}

func (o MeIdentityUserMapOutput) ToMeIdentityUserMapOutputWithContext(ctx context.Context) MeIdentityUserMapOutput {
	return o
}

func (o MeIdentityUserMapOutput) MapIndex(k pulumi.StringInput) MeIdentityUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MeIdentityUser {
		return vs[0].(map[string]MeIdentityUser)[vs[1].(string)]
	}).(MeIdentityUserOutput)
}

func init() {
	pulumi.RegisterOutputType(MeIdentityUserOutput{})
	pulumi.RegisterOutputType(MeIdentityUserPtrOutput{})
	pulumi.RegisterOutputType(MeIdentityUserArrayOutput{})
	pulumi.RegisterOutputType(MeIdentityUserMapOutput{})
}
