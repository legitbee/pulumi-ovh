// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpService struct {
	pulumi.CustomResourceState

	CanBeTerminated pulumi.BoolOutput   `pulumi:"canBeTerminated"`
	Country         pulumi.StringOutput `pulumi:"country"`
	// Custom description on your ip
	Description pulumi.StringOutput `pulumi:"description"`
	Ip          pulumi.StringOutput `pulumi:"ip"`
	// Details about an Order
	Orders         IpServiceOrderArrayOutput `pulumi:"orders"`
	OrganisationId pulumi.StringOutput       `pulumi:"organisationId"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean pulumi.StringOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IpServicePlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayOutput `pulumi:"planOptions"`
	// Routage information
	RoutedTos   IpServiceRoutedToArrayOutput `pulumi:"routedTos"`
	ServiceName pulumi.StringOutput          `pulumi:"serviceName"`
	// Possible values for ip type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpService registers a new resource with the given unique name, arguments, and options.
func NewIpService(ctx *pulumi.Context,
	name string, args *IpServiceArgs, opts ...pulumi.ResourceOption) (*IpService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OvhSubsidiary == nil {
		return nil, errors.New("invalid value for required argument 'OvhSubsidiary'")
	}
	if args.PaymentMean == nil {
		return nil, errors.New("invalid value for required argument 'PaymentMean'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	var resource IpService
	err := ctx.RegisterResource("ovh:index/ipService:IpService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpService gets an existing IpService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpServiceState, opts ...pulumi.ResourceOption) (*IpService, error) {
	var resource IpService
	err := ctx.ReadResource("ovh:index/ipService:IpService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpService resources.
type ipServiceState struct {
	CanBeTerminated *bool   `pulumi:"canBeTerminated"`
	Country         *string `pulumi:"country"`
	// Custom description on your ip
	Description *string `pulumi:"description"`
	Ip          *string `pulumi:"ip"`
	// Details about an Order
	Orders         []IpServiceOrder `pulumi:"orders"`
	OrganisationId *string          `pulumi:"organisationId"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *IpServicePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpServicePlanOption `pulumi:"planOptions"`
	// Routage information
	RoutedTos   []IpServiceRoutedTo `pulumi:"routedTos"`
	ServiceName *string             `pulumi:"serviceName"`
	// Possible values for ip type
	Type *string `pulumi:"type"`
}

type IpServiceState struct {
	CanBeTerminated pulumi.BoolPtrInput
	Country         pulumi.StringPtrInput
	// Custom description on your ip
	Description pulumi.StringPtrInput
	Ip          pulumi.StringPtrInput
	// Details about an Order
	Orders         IpServiceOrderArrayInput
	OrganisationId pulumi.StringPtrInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan IpServicePlanPtrInput
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayInput
	// Routage information
	RoutedTos   IpServiceRoutedToArrayInput
	ServiceName pulumi.StringPtrInput
	// Possible values for ip type
	Type pulumi.StringPtrInput
}

func (IpServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipServiceState)(nil)).Elem()
}

type ipServiceArgs struct {
	// Custom description on your ip
	Description *string `pulumi:"description"`
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IpServicePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpServicePlanOption `pulumi:"planOptions"`
}

// The set of arguments for constructing a IpService resource.
type IpServiceArgs struct {
	// Custom description on your ip
	Description pulumi.StringPtrInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode
	PaymentMean pulumi.StringInput
	// Product Plan to order
	Plan IpServicePlanInput
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayInput
}

func (IpServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipServiceArgs)(nil)).Elem()
}

type IpServiceInput interface {
	pulumi.Input

	ToIpServiceOutput() IpServiceOutput
	ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput
}

func (*IpService) ElementType() reflect.Type {
	return reflect.TypeOf((*IpService)(nil))
}

func (i *IpService) ToIpServiceOutput() IpServiceOutput {
	return i.ToIpServiceOutputWithContext(context.Background())
}

func (i *IpService) ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceOutput)
}

func (i *IpService) ToIpServicePtrOutput() IpServicePtrOutput {
	return i.ToIpServicePtrOutputWithContext(context.Background())
}

func (i *IpService) ToIpServicePtrOutputWithContext(ctx context.Context) IpServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServicePtrOutput)
}

type IpServicePtrInput interface {
	pulumi.Input

	ToIpServicePtrOutput() IpServicePtrOutput
	ToIpServicePtrOutputWithContext(ctx context.Context) IpServicePtrOutput
}

type ipServicePtrType IpServiceArgs

func (*ipServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpService)(nil))
}

func (i *ipServicePtrType) ToIpServicePtrOutput() IpServicePtrOutput {
	return i.ToIpServicePtrOutputWithContext(context.Background())
}

func (i *ipServicePtrType) ToIpServicePtrOutputWithContext(ctx context.Context) IpServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServicePtrOutput)
}

// IpServiceArrayInput is an input type that accepts IpServiceArray and IpServiceArrayOutput values.
// You can construct a concrete instance of `IpServiceArrayInput` via:
//
//          IpServiceArray{ IpServiceArgs{...} }
type IpServiceArrayInput interface {
	pulumi.Input

	ToIpServiceArrayOutput() IpServiceArrayOutput
	ToIpServiceArrayOutputWithContext(context.Context) IpServiceArrayOutput
}

type IpServiceArray []IpServiceInput

func (IpServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpService)(nil)).Elem()
}

func (i IpServiceArray) ToIpServiceArrayOutput() IpServiceArrayOutput {
	return i.ToIpServiceArrayOutputWithContext(context.Background())
}

func (i IpServiceArray) ToIpServiceArrayOutputWithContext(ctx context.Context) IpServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceArrayOutput)
}

// IpServiceMapInput is an input type that accepts IpServiceMap and IpServiceMapOutput values.
// You can construct a concrete instance of `IpServiceMapInput` via:
//
//          IpServiceMap{ "key": IpServiceArgs{...} }
type IpServiceMapInput interface {
	pulumi.Input

	ToIpServiceMapOutput() IpServiceMapOutput
	ToIpServiceMapOutputWithContext(context.Context) IpServiceMapOutput
}

type IpServiceMap map[string]IpServiceInput

func (IpServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpService)(nil)).Elem()
}

func (i IpServiceMap) ToIpServiceMapOutput() IpServiceMapOutput {
	return i.ToIpServiceMapOutputWithContext(context.Background())
}

func (i IpServiceMap) ToIpServiceMapOutputWithContext(ctx context.Context) IpServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceMapOutput)
}

type IpServiceOutput struct{ *pulumi.OutputState }

func (IpServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpService)(nil))
}

func (o IpServiceOutput) ToIpServiceOutput() IpServiceOutput {
	return o
}

func (o IpServiceOutput) ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput {
	return o
}

func (o IpServiceOutput) ToIpServicePtrOutput() IpServicePtrOutput {
	return o.ToIpServicePtrOutputWithContext(context.Background())
}

func (o IpServiceOutput) ToIpServicePtrOutputWithContext(ctx context.Context) IpServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpService) *IpService {
		return &v
	}).(IpServicePtrOutput)
}

type IpServicePtrOutput struct{ *pulumi.OutputState }

func (IpServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpService)(nil))
}

func (o IpServicePtrOutput) ToIpServicePtrOutput() IpServicePtrOutput {
	return o
}

func (o IpServicePtrOutput) ToIpServicePtrOutputWithContext(ctx context.Context) IpServicePtrOutput {
	return o
}

func (o IpServicePtrOutput) Elem() IpServiceOutput {
	return o.ApplyT(func(v *IpService) IpService {
		if v != nil {
			return *v
		}
		var ret IpService
		return ret
	}).(IpServiceOutput)
}

type IpServiceArrayOutput struct{ *pulumi.OutputState }

func (IpServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpService)(nil))
}

func (o IpServiceArrayOutput) ToIpServiceArrayOutput() IpServiceArrayOutput {
	return o
}

func (o IpServiceArrayOutput) ToIpServiceArrayOutputWithContext(ctx context.Context) IpServiceArrayOutput {
	return o
}

func (o IpServiceArrayOutput) Index(i pulumi.IntInput) IpServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpService {
		return vs[0].([]IpService)[vs[1].(int)]
	}).(IpServiceOutput)
}

type IpServiceMapOutput struct{ *pulumi.OutputState }

func (IpServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IpService)(nil))
}

func (o IpServiceMapOutput) ToIpServiceMapOutput() IpServiceMapOutput {
	return o
}

func (o IpServiceMapOutput) ToIpServiceMapOutputWithContext(ctx context.Context) IpServiceMapOutput {
	return o
}

func (o IpServiceMapOutput) MapIndex(k pulumi.StringInput) IpServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IpService {
		return vs[0].(map[string]IpService)[vs[1].(string)]
	}).(IpServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(IpServiceOutput{})
	pulumi.RegisterOutputType(IpServicePtrOutput{})
	pulumi.RegisterOutputType(IpServiceArrayOutput{})
	pulumi.RegisterOutputType(IpServiceMapOutput{})
}
