// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudProjectUser struct {
	pulumi.CustomResourceState

	CreationDate pulumi.StringOutput             `pulumi:"creationDate"`
	Description  pulumi.StringPtrOutput          `pulumi:"description"`
	OpenstackRc  pulumi.MapOutput                `pulumi:"openstackRc"`
	Password     pulumi.StringOutput             `pulumi:"password"`
	RoleName     pulumi.StringPtrOutput          `pulumi:"roleName"`
	RoleNames    pulumi.StringArrayOutput        `pulumi:"roleNames"`
	Roles        CloudProjectUserRoleArrayOutput `pulumi:"roles"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	Status      pulumi.StringOutput `pulumi:"status"`
	Username    pulumi.StringOutput `pulumi:"username"`
}

// NewCloudProjectUser registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectUser(ctx *pulumi.Context,
	name string, args *CloudProjectUserArgs, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	var resource CloudProjectUser
	err := ctx.RegisterResource("ovh:index/cloudProjectUser:CloudProjectUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectUser gets an existing CloudProjectUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectUserState, opts ...pulumi.ResourceOption) (*CloudProjectUser, error) {
	var resource CloudProjectUser
	err := ctx.ReadResource("ovh:index/cloudProjectUser:CloudProjectUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectUser resources.
type cloudProjectUserState struct {
	CreationDate *string                `pulumi:"creationDate"`
	Description  *string                `pulumi:"description"`
	OpenstackRc  map[string]interface{} `pulumi:"openstackRc"`
	Password     *string                `pulumi:"password"`
	RoleName     *string                `pulumi:"roleName"`
	RoleNames    []string               `pulumi:"roleNames"`
	Roles        []CloudProjectUserRole `pulumi:"roles"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName *string `pulumi:"serviceName"`
	Status      *string `pulumi:"status"`
	Username    *string `pulumi:"username"`
}

type CloudProjectUserState struct {
	CreationDate pulumi.StringPtrInput
	Description  pulumi.StringPtrInput
	OpenstackRc  pulumi.MapInput
	Password     pulumi.StringPtrInput
	RoleName     pulumi.StringPtrInput
	RoleNames    pulumi.StringArrayInput
	Roles        CloudProjectUserRoleArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringPtrInput
	Status      pulumi.StringPtrInput
	Username    pulumi.StringPtrInput
}

func (CloudProjectUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserState)(nil)).Elem()
}

type cloudProjectUserArgs struct {
	Description *string                `pulumi:"description"`
	OpenstackRc map[string]interface{} `pulumi:"openstackRc"`
	RoleName    *string                `pulumi:"roleName"`
	RoleNames   []string               `pulumi:"roleNames"`
	// Service name of the resource representing the id of the cloud project.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectUser resource.
type CloudProjectUserArgs struct {
	Description pulumi.StringPtrInput
	OpenstackRc pulumi.MapInput
	RoleName    pulumi.StringPtrInput
	RoleNames   pulumi.StringArrayInput
	// Service name of the resource representing the id of the cloud project.
	ServiceName pulumi.StringInput
}

func (CloudProjectUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectUserArgs)(nil)).Elem()
}

type CloudProjectUserInput interface {
	pulumi.Input

	ToCloudProjectUserOutput() CloudProjectUserOutput
	ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput
}

func (*CloudProjectUser) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectUser)(nil))
}

func (i *CloudProjectUser) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return i.ToCloudProjectUserOutputWithContext(context.Background())
}

func (i *CloudProjectUser) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserOutput)
}

func (i *CloudProjectUser) ToCloudProjectUserPtrOutput() CloudProjectUserPtrOutput {
	return i.ToCloudProjectUserPtrOutputWithContext(context.Background())
}

func (i *CloudProjectUser) ToCloudProjectUserPtrOutputWithContext(ctx context.Context) CloudProjectUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserPtrOutput)
}

type CloudProjectUserPtrInput interface {
	pulumi.Input

	ToCloudProjectUserPtrOutput() CloudProjectUserPtrOutput
	ToCloudProjectUserPtrOutputWithContext(ctx context.Context) CloudProjectUserPtrOutput
}

type cloudProjectUserPtrType CloudProjectUserArgs

func (*cloudProjectUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil))
}

func (i *cloudProjectUserPtrType) ToCloudProjectUserPtrOutput() CloudProjectUserPtrOutput {
	return i.ToCloudProjectUserPtrOutputWithContext(context.Background())
}

func (i *cloudProjectUserPtrType) ToCloudProjectUserPtrOutputWithContext(ctx context.Context) CloudProjectUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserPtrOutput)
}

// CloudProjectUserArrayInput is an input type that accepts CloudProjectUserArray and CloudProjectUserArrayOutput values.
// You can construct a concrete instance of `CloudProjectUserArrayInput` via:
//
//          CloudProjectUserArray{ CloudProjectUserArgs{...} }
type CloudProjectUserArrayInput interface {
	pulumi.Input

	ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput
	ToCloudProjectUserArrayOutputWithContext(context.Context) CloudProjectUserArrayOutput
}

type CloudProjectUserArray []CloudProjectUserInput

func (CloudProjectUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectUser)(nil)).Elem()
}

func (i CloudProjectUserArray) ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput {
	return i.ToCloudProjectUserArrayOutputWithContext(context.Background())
}

func (i CloudProjectUserArray) ToCloudProjectUserArrayOutputWithContext(ctx context.Context) CloudProjectUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserArrayOutput)
}

// CloudProjectUserMapInput is an input type that accepts CloudProjectUserMap and CloudProjectUserMapOutput values.
// You can construct a concrete instance of `CloudProjectUserMapInput` via:
//
//          CloudProjectUserMap{ "key": CloudProjectUserArgs{...} }
type CloudProjectUserMapInput interface {
	pulumi.Input

	ToCloudProjectUserMapOutput() CloudProjectUserMapOutput
	ToCloudProjectUserMapOutputWithContext(context.Context) CloudProjectUserMapOutput
}

type CloudProjectUserMap map[string]CloudProjectUserInput

func (CloudProjectUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectUser)(nil)).Elem()
}

func (i CloudProjectUserMap) ToCloudProjectUserMapOutput() CloudProjectUserMapOutput {
	return i.ToCloudProjectUserMapOutputWithContext(context.Background())
}

func (i CloudProjectUserMap) ToCloudProjectUserMapOutputWithContext(ctx context.Context) CloudProjectUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectUserMapOutput)
}

type CloudProjectUserOutput struct{ *pulumi.OutputState }

func (CloudProjectUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudProjectUser)(nil))
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutput() CloudProjectUserOutput {
	return o
}

func (o CloudProjectUserOutput) ToCloudProjectUserOutputWithContext(ctx context.Context) CloudProjectUserOutput {
	return o
}

func (o CloudProjectUserOutput) ToCloudProjectUserPtrOutput() CloudProjectUserPtrOutput {
	return o.ToCloudProjectUserPtrOutputWithContext(context.Background())
}

func (o CloudProjectUserOutput) ToCloudProjectUserPtrOutputWithContext(ctx context.Context) CloudProjectUserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudProjectUser) *CloudProjectUser {
		return &v
	}).(CloudProjectUserPtrOutput)
}

type CloudProjectUserPtrOutput struct{ *pulumi.OutputState }

func (CloudProjectUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectUser)(nil))
}

func (o CloudProjectUserPtrOutput) ToCloudProjectUserPtrOutput() CloudProjectUserPtrOutput {
	return o
}

func (o CloudProjectUserPtrOutput) ToCloudProjectUserPtrOutputWithContext(ctx context.Context) CloudProjectUserPtrOutput {
	return o
}

func (o CloudProjectUserPtrOutput) Elem() CloudProjectUserOutput {
	return o.ApplyT(func(v *CloudProjectUser) CloudProjectUser {
		if v != nil {
			return *v
		}
		var ret CloudProjectUser
		return ret
	}).(CloudProjectUserOutput)
}

type CloudProjectUserArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CloudProjectUser)(nil))
}

func (o CloudProjectUserArrayOutput) ToCloudProjectUserArrayOutput() CloudProjectUserArrayOutput {
	return o
}

func (o CloudProjectUserArrayOutput) ToCloudProjectUserArrayOutputWithContext(ctx context.Context) CloudProjectUserArrayOutput {
	return o
}

func (o CloudProjectUserArrayOutput) Index(i pulumi.IntInput) CloudProjectUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CloudProjectUser {
		return vs[0].([]CloudProjectUser)[vs[1].(int)]
	}).(CloudProjectUserOutput)
}

type CloudProjectUserMapOutput struct{ *pulumi.OutputState }

func (CloudProjectUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CloudProjectUser)(nil))
}

func (o CloudProjectUserMapOutput) ToCloudProjectUserMapOutput() CloudProjectUserMapOutput {
	return o
}

func (o CloudProjectUserMapOutput) ToCloudProjectUserMapOutputWithContext(ctx context.Context) CloudProjectUserMapOutput {
	return o
}

func (o CloudProjectUserMapOutput) MapIndex(k pulumi.StringInput) CloudProjectUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CloudProjectUser {
		return vs[0].(map[string]CloudProjectUser)[vs[1].(string)]
	}).(CloudProjectUserOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudProjectUserOutput{})
	pulumi.RegisterOutputType(CloudProjectUserPtrOutput{})
	pulumi.RegisterOutputType(CloudProjectUserArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectUserMapOutput{})
}
