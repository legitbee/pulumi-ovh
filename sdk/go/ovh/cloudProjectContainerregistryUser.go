// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudProjectContainerregistryUser struct {
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

// NewCloudProjectContainerregistryUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectContainerregistryUser(ctx *pulumi.Context,
	name string, args *CloudProjectContainerregistryUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistryUser, error) {
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
	var resource CloudProjectContainerregistryUser
	err := ctx.RegisterResource("ovh:index/cloudProjectContainerregistryUser:CloudProjectContainerregistryUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectContainerregistryUser gets an existing CloudProjectContainerregistryUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectContainerregistryUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectContainerregistryUserState, opts ...pulumi.ResourceOption) (*CloudProjectContainerregistryUser, error) {
	var resource CloudProjectContainerregistryUser
	err := ctx.ReadResource("ovh:index/cloudProjectContainerregistryUser:CloudProjectContainerregistryUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectContainerregistryUser resources.
type cloudProjectContainerregistryUserState struct {
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

type CloudProjectContainerregistryUserState struct {
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

func (CloudProjectContainerregistryUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryUserState)(nil)).Elem()
}

type cloudProjectContainerregistryUserArgs struct {
	// User email.
	Email string `pulumi:"email"`
	// Registry name
	Login string `pulumi:"login"`
	// RegistryID
	RegistryId string `pulumi:"registryId"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectContainerregistryUser resource.
type CloudProjectContainerregistryUserArgs struct {
	// User email.
	Email pulumi.StringInput
	// Registry name
	Login pulumi.StringInput
	// RegistryID
	RegistryId pulumi.StringInput
	// Service name
	ServiceName pulumi.StringInput
}

func (CloudProjectContainerregistryUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerregistryUserArgs)(nil)).Elem()
}

type CloudProjectContainerregistryUserInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryUserOutput() CloudProjectContainerregistryUserOutput
	ToCloudProjectContainerregistryUserOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserOutput
}

func (*CloudProjectContainerregistryUser) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectContainerregistryUser)(nil))
}

func (i *CloudProjectContainerregistryUser) ToCloudProjectContainerregistryUserOutput() CloudProjectContainerregistryUserOutput {
	return i.ToCloudProjectContainerregistryUserOutputWithContext(context.Background())
}

func (i *CloudProjectContainerregistryUser) ToCloudProjectContainerregistryUserOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryUserOutput)
}

func (i *CloudProjectContainerregistryUser) ToCloudProjectContainerregistryUserPtrOutput() CloudProjectContainerregistryUserPtrOutput {
	return i.ToCloudProjectContainerregistryUserPtrOutputWithContext(context.Background())
}

func (i *CloudProjectContainerregistryUser) ToCloudProjectContainerregistryUserPtrOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryUserPtrOutput)
}

type CloudProjectContainerregistryUserPtrInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryUserPtrOutput() CloudProjectContainerregistryUserPtrOutput
	ToCloudProjectContainerregistryUserPtrOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserPtrOutput
}

type cloudProjectContainerregistryUserPtrType CloudProjectContainerregistryUserArgs

func (*cloudProjectContainerregistryUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistryUser)(nil))
}

func (i *cloudProjectContainerregistryUserPtrType) ToCloudProjectContainerregistryUserPtrOutput() CloudProjectContainerregistryUserPtrOutput {
	return i.ToCloudProjectContainerregistryUserPtrOutputWithContext(context.Background())
}

func (i *cloudProjectContainerregistryUserPtrType) ToCloudProjectContainerregistryUserPtrOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryUserPtrOutput)
}

// CloudProjectContainerregistryUserArrayInput is an input type that accepts CloudProjectContainerregistryUserArray and CloudProjectContainerregistryUserArrayOutput values.
// You can construct a concrete instance of `CloudProjectContainerregistryUserArrayInput` via:
//
//          CloudProjectContainerregistryUserArray{ CloudProjectContainerregistryUserArgs{...} }
type CloudProjectContainerregistryUserArrayInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryUserArrayOutput() CloudProjectContainerregistryUserArrayOutput
	ToCloudProjectContainerregistryUserArrayOutputWithContext(context.Context) CloudProjectContainerregistryUserArrayOutput
}

type CloudProjectContainerregistryUserArray []CloudProjectContainerregistryUserInput

func (CloudProjectContainerregistryUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectContainerregistryUser)(nil)).Elem()
}

func (i CloudProjectContainerregistryUserArray) ToCloudProjectContainerregistryUserArrayOutput() CloudProjectContainerregistryUserArrayOutput {
	return i.ToCloudProjectContainerregistryUserArrayOutputWithContext(context.Background())
}

func (i CloudProjectContainerregistryUserArray) ToCloudProjectContainerregistryUserArrayOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryUserArrayOutput)
}

// CloudProjectContainerregistryUserMapInput is an input type that accepts CloudProjectContainerregistryUserMap and CloudProjectContainerregistryUserMapOutput values.
// You can construct a concrete instance of `CloudProjectContainerregistryUserMapInput` via:
//
//          CloudProjectContainerregistryUserMap{ "key": CloudProjectContainerregistryUserArgs{...} }
type CloudProjectContainerregistryUserMapInput interface {
	pulumi.Input

	ToCloudProjectContainerregistryUserMapOutput() CloudProjectContainerregistryUserMapOutput
	ToCloudProjectContainerregistryUserMapOutputWithContext(context.Context) CloudProjectContainerregistryUserMapOutput
}

type CloudProjectContainerregistryUserMap map[string]CloudProjectContainerregistryUserInput

func (CloudProjectContainerregistryUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectContainerregistryUser)(nil)).Elem()
}

func (i CloudProjectContainerregistryUserMap) ToCloudProjectContainerregistryUserMapOutput() CloudProjectContainerregistryUserMapOutput {
	return i.ToCloudProjectContainerregistryUserMapOutputWithContext(context.Background())
}

func (i CloudProjectContainerregistryUserMap) ToCloudProjectContainerregistryUserMapOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerregistryUserMapOutput)
}

type CloudProjectContainerregistryUserOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectContainerregistryUser)(nil))
}

func (o CloudProjectContainerregistryUserOutput) ToCloudProjectContainerregistryUserOutput() CloudProjectContainerregistryUserOutput {
	return o
}

func (o CloudProjectContainerregistryUserOutput) ToCloudProjectContainerregistryUserOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserOutput {
	return o
}

func (o CloudProjectContainerregistryUserOutput) ToCloudProjectContainerregistryUserPtrOutput() CloudProjectContainerregistryUserPtrOutput {
	return o.ToCloudProjectContainerregistryUserPtrOutputWithContext(context.Background())
}

func (o CloudProjectContainerregistryUserOutput) ToCloudProjectContainerregistryUserPtrOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProjectContainerregistryUser) *CloudProjectContainerregistryUser {
		return &v
	}).(CloudProjectContainerregistryUserPtrOutput)
}

type CloudProjectContainerregistryUserPtrOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerregistryUser)(nil))
}

func (o CloudProjectContainerregistryUserPtrOutput) ToCloudProjectContainerregistryUserPtrOutput() CloudProjectContainerregistryUserPtrOutput {
	return o
}

func (o CloudProjectContainerregistryUserPtrOutput) ToCloudProjectContainerregistryUserPtrOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserPtrOutput {
	return o
}

func (o CloudProjectContainerregistryUserPtrOutput) Elem() CloudProjectContainerregistryUserOutput {
	return o.ApplyT(func(v *CloudProjectContainerregistryUser) CloudProjectContainerregistryUser {
		if v != nil {
			return *v
		}
		var ret CloudProjectContainerregistryUser
		return ret
	}).(CloudProjectContainerregistryUserOutput)
}

type CloudProjectContainerregistryUserArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudProjectContainerregistryUser)(nil))
}

func (o CloudProjectContainerregistryUserArrayOutput) ToCloudProjectContainerregistryUserArrayOutput() CloudProjectContainerregistryUserArrayOutput {
	return o
}

func (o CloudProjectContainerregistryUserArrayOutput) ToCloudProjectContainerregistryUserArrayOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserArrayOutput {
	return o
}

func (o CloudProjectContainerregistryUserArrayOutput) Index(i pulumi.IntInput) CloudProjectContainerregistryUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudProjectContainerregistryUser {
		return vs[0].([]CloudProjectContainerregistryUser)[vs[1].(int)]
	}).(CloudProjectContainerregistryUserOutput)
}

type CloudProjectContainerregistryUserMapOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerregistryUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudProjectContainerregistryUser)(nil))
}

func (o CloudProjectContainerregistryUserMapOutput) ToCloudProjectContainerregistryUserMapOutput() CloudProjectContainerregistryUserMapOutput {
	return o
}

func (o CloudProjectContainerregistryUserMapOutput) ToCloudProjectContainerregistryUserMapOutputWithContext(ctx context.Context) CloudProjectContainerregistryUserMapOutput {
	return o
}

func (o CloudProjectContainerregistryUserMapOutput) MapIndex(k pulumi.StringInput) CloudProjectContainerregistryUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudProjectContainerregistryUser {
		return vs[0].(map[string]CloudProjectContainerregistryUser)[vs[1].(string)]
	}).(CloudProjectContainerregistryUserOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudProjectContainerregistryUserOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerregistryUserPtrOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerregistryUserArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerregistryUserMapOutput{})
}
