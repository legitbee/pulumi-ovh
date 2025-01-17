// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Iploadbalancing struct {
	pulumi.CustomResourceState

	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing pulumi.StringOutput `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringOutput `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer pulumi.StringOutput `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones IploadbalancingOrderableZoneArrayOutput `pulumi:"orderableZones"`
	// Details about an Order
	Orders IploadbalancingOrderArrayOutput `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean pulumi.StringOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IploadbalancingPlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions IploadbalancingPlanOptionArrayOutput `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and
	// Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android
	// 2.3, Java 7. Intermediate if null.
	SslConfiguration pulumi.StringOutput `pulumi:"sslConfiguration"`
	// Current state of your IP
	State pulumi.StringOutput `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility pulumi.BoolOutput `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringOutput `pulumi:"vrackName"`
	// Location where your service is
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewIploadbalancing registers a new resource with the given unique name, arguments, and options.
func NewIploadbalancing(ctx *pulumi.Context,
	name string, args *IploadbalancingArgs, opts ...pulumi.ResourceOption) (*Iploadbalancing, error) {
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
	var resource Iploadbalancing
	err := ctx.RegisterResource("ovh:index/iploadbalancing:Iploadbalancing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIploadbalancing gets an existing Iploadbalancing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIploadbalancing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IploadbalancingState, opts ...pulumi.ResourceOption) (*Iploadbalancing, error) {
	var resource Iploadbalancing
	err := ctx.ReadResource("ovh:index/iploadbalancing:Iploadbalancing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Iploadbalancing resources.
type iploadbalancingState struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing *string `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 *string `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 *string `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken *string `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer *string `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones []IploadbalancingOrderableZone `pulumi:"orderableZones"`
	// Details about an Order
	Orders []IploadbalancingOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *IploadbalancingPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IploadbalancingPlanOption `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and
	// Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android
	// 2.3, Java 7. Intermediate if null.
	SslConfiguration *string `pulumi:"sslConfiguration"`
	// Current state of your IP
	State *string `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility *bool `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName *string `pulumi:"vrackName"`
	// Location where your service is
	Zones []string `pulumi:"zones"`
}

type IploadbalancingState struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Your IP load balancing
	IpLoadbalancing pulumi.StringPtrInput
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringPtrInput
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringPtrInput
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringPtrInput
	// The offer of your IP load balancing
	Offer pulumi.StringPtrInput
	// Available additional zone for your Load Balancer
	OrderableZones IploadbalancingOrderableZoneArrayInput
	// Details about an Order
	Orders IploadbalancingOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan IploadbalancingPlanPtrInput
	// Product Plan to order
	PlanOptions IploadbalancingPlanOptionArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and
	// Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android
	// 2.3, Java 7. Intermediate if null.
	SslConfiguration pulumi.StringPtrInput
	// Current state of your IP
	State pulumi.StringPtrInput
	// Vrack eligibility
	VrackEligibility pulumi.BoolPtrInput
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringPtrInput
	// Location where your service is
	Zones pulumi.StringArrayInput
}

func (IploadbalancingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingState)(nil)).Elem()
}

type iploadbalancingArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	PaymentMean string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IploadbalancingPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IploadbalancingPlanOption `pulumi:"planOptions"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and
	// Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android
	// 2.3, Java 7. Intermediate if null.
	SslConfiguration *string `pulumi:"sslConfiguration"`
}

// The set of arguments for constructing a Iploadbalancing resource.
type IploadbalancingArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode
	PaymentMean pulumi.StringInput
	// Product Plan to order
	Plan IploadbalancingPlanInput
	// Product Plan to order
	PlanOptions IploadbalancingPlanOptionArrayInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and
	// Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android
	// 2.3, Java 7. Intermediate if null.
	SslConfiguration pulumi.StringPtrInput
}

func (IploadbalancingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iploadbalancingArgs)(nil)).Elem()
}

type IploadbalancingInput interface {
	pulumi.Input

	ToIploadbalancingOutput() IploadbalancingOutput
	ToIploadbalancingOutputWithContext(ctx context.Context) IploadbalancingOutput
}

func (*Iploadbalancing) ElementType() reflect.Type {
	return reflect.TypeOf((*Iploadbalancing)(nil))
}

func (i *Iploadbalancing) ToIploadbalancingOutput() IploadbalancingOutput {
	return i.ToIploadbalancingOutputWithContext(context.Background())
}

func (i *Iploadbalancing) ToIploadbalancingOutputWithContext(ctx context.Context) IploadbalancingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingOutput)
}

func (i *Iploadbalancing) ToIploadbalancingPtrOutput() IploadbalancingPtrOutput {
	return i.ToIploadbalancingPtrOutputWithContext(context.Background())
}

func (i *Iploadbalancing) ToIploadbalancingPtrOutputWithContext(ctx context.Context) IploadbalancingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingPtrOutput)
}

type IploadbalancingPtrInput interface {
	pulumi.Input

	ToIploadbalancingPtrOutput() IploadbalancingPtrOutput
	ToIploadbalancingPtrOutputWithContext(ctx context.Context) IploadbalancingPtrOutput
}

type iploadbalancingPtrType IploadbalancingArgs

func (*iploadbalancingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Iploadbalancing)(nil))
}

func (i *iploadbalancingPtrType) ToIploadbalancingPtrOutput() IploadbalancingPtrOutput {
	return i.ToIploadbalancingPtrOutputWithContext(context.Background())
}

func (i *iploadbalancingPtrType) ToIploadbalancingPtrOutputWithContext(ctx context.Context) IploadbalancingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingPtrOutput)
}

// IploadbalancingArrayInput is an input type that accepts IploadbalancingArray and IploadbalancingArrayOutput values.
// You can construct a concrete instance of `IploadbalancingArrayInput` via:
//
//          IploadbalancingArray{ IploadbalancingArgs{...} }
type IploadbalancingArrayInput interface {
	pulumi.Input

	ToIploadbalancingArrayOutput() IploadbalancingArrayOutput
	ToIploadbalancingArrayOutputWithContext(context.Context) IploadbalancingArrayOutput
}

type IploadbalancingArray []IploadbalancingInput

func (IploadbalancingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Iploadbalancing)(nil)).Elem()
}

func (i IploadbalancingArray) ToIploadbalancingArrayOutput() IploadbalancingArrayOutput {
	return i.ToIploadbalancingArrayOutputWithContext(context.Background())
}

func (i IploadbalancingArray) ToIploadbalancingArrayOutputWithContext(ctx context.Context) IploadbalancingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingArrayOutput)
}

// IploadbalancingMapInput is an input type that accepts IploadbalancingMap and IploadbalancingMapOutput values.
// You can construct a concrete instance of `IploadbalancingMapInput` via:
//
//          IploadbalancingMap{ "key": IploadbalancingArgs{...} }
type IploadbalancingMapInput interface {
	pulumi.Input

	ToIploadbalancingMapOutput() IploadbalancingMapOutput
	ToIploadbalancingMapOutputWithContext(context.Context) IploadbalancingMapOutput
}

type IploadbalancingMap map[string]IploadbalancingInput

func (IploadbalancingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Iploadbalancing)(nil)).Elem()
}

func (i IploadbalancingMap) ToIploadbalancingMapOutput() IploadbalancingMapOutput {
	return i.ToIploadbalancingMapOutputWithContext(context.Background())
}

func (i IploadbalancingMap) ToIploadbalancingMapOutputWithContext(ctx context.Context) IploadbalancingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IploadbalancingMapOutput)
}

type IploadbalancingOutput struct{ *pulumi.OutputState }

func (IploadbalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Iploadbalancing)(nil))
}

func (o IploadbalancingOutput) ToIploadbalancingOutput() IploadbalancingOutput {
	return o
}

func (o IploadbalancingOutput) ToIploadbalancingOutputWithContext(ctx context.Context) IploadbalancingOutput {
	return o
}

func (o IploadbalancingOutput) ToIploadbalancingPtrOutput() IploadbalancingPtrOutput {
	return o.ToIploadbalancingPtrOutputWithContext(context.Background())
}

func (o IploadbalancingOutput) ToIploadbalancingPtrOutputWithContext(ctx context.Context) IploadbalancingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Iploadbalancing) *Iploadbalancing {
		return &v
	}).(IploadbalancingPtrOutput)
}

type IploadbalancingPtrOutput struct{ *pulumi.OutputState }

func (IploadbalancingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Iploadbalancing)(nil))
}

func (o IploadbalancingPtrOutput) ToIploadbalancingPtrOutput() IploadbalancingPtrOutput {
	return o
}

func (o IploadbalancingPtrOutput) ToIploadbalancingPtrOutputWithContext(ctx context.Context) IploadbalancingPtrOutput {
	return o
}

func (o IploadbalancingPtrOutput) Elem() IploadbalancingOutput {
	return o.ApplyT(func(v *Iploadbalancing) Iploadbalancing {
		if v != nil {
			return *v
		}
		var ret Iploadbalancing
		return ret
	}).(IploadbalancingOutput)
}

type IploadbalancingArrayOutput struct{ *pulumi.OutputState }

func (IploadbalancingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Iploadbalancing)(nil))
}

func (o IploadbalancingArrayOutput) ToIploadbalancingArrayOutput() IploadbalancingArrayOutput {
	return o
}

func (o IploadbalancingArrayOutput) ToIploadbalancingArrayOutputWithContext(ctx context.Context) IploadbalancingArrayOutput {
	return o
}

func (o IploadbalancingArrayOutput) Index(i pulumi.IntInput) IploadbalancingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Iploadbalancing {
		return vs[0].([]Iploadbalancing)[vs[1].(int)]
	}).(IploadbalancingOutput)
}

type IploadbalancingMapOutput struct{ *pulumi.OutputState }

func (IploadbalancingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Iploadbalancing)(nil))
}

func (o IploadbalancingMapOutput) ToIploadbalancingMapOutput() IploadbalancingMapOutput {
	return o
}

func (o IploadbalancingMapOutput) ToIploadbalancingMapOutputWithContext(ctx context.Context) IploadbalancingMapOutput {
	return o
}

func (o IploadbalancingMapOutput) MapIndex(k pulumi.StringInput) IploadbalancingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Iploadbalancing {
		return vs[0].(map[string]Iploadbalancing)[vs[1].(string)]
	}).(IploadbalancingOutput)
}

func init() {
	pulumi.RegisterOutputType(IploadbalancingOutput{})
	pulumi.RegisterOutputType(IploadbalancingPtrOutput{})
	pulumi.RegisterOutputType(IploadbalancingArrayOutput{})
	pulumi.RegisterOutputType(IploadbalancingMapOutput{})
}
