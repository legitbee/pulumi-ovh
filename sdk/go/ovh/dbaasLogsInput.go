// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DbaasLogsInput struct {
	pulumi.CustomResourceState

	// IP blocks
	AllowedNetworks pulumi.StringArrayOutput `pulumi:"allowedNetworks"`
	// Input configuration
	Configuration DbaasLogsInputConfigurationOutput `pulumi:"configuration"`
	// Input creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Input description
	Description pulumi.StringOutput `pulumi:"description"`
	// Input engine ID
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// Port
	ExposedPort pulumi.StringOutput `pulumi:"exposedPort"`
	// Hostname
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Input ID
	InputId pulumi.StringOutput `pulumi:"inputId"`
	// Indicate if input need to be restarted
	IsRestartRequired pulumi.BoolOutput `pulumi:"isRestartRequired"`
	// Number of instance running
	NbInstance pulumi.IntOutput `pulumi:"nbInstance"`
	// Input IP address
	PublicAddress pulumi.StringOutput `pulumi:"publicAddress"`
	ServiceName   pulumi.StringOutput `pulumi:"serviceName"`
	// Input SSL certificate
	SslCertificate pulumi.StringOutput `pulumi:"sslCertificate"`
	// init: configuration required, pending: ready to start, running: available
	Status pulumi.StringOutput `pulumi:"status"`
	// Associated Graylog stream
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// Input title
	Title pulumi.StringOutput `pulumi:"title"`
	// Input last update
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDbaasLogsInput registers a new resource with the given unique name, arguments, and options.
func NewDbaasLogsInput(ctx *pulumi.Context,
	name string, args *DbaasLogsInputArgs, opts ...pulumi.ResourceOption) (*DbaasLogsInput, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.EngineId == nil {
		return nil, errors.New("invalid value for required argument 'EngineId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.StreamId == nil {
		return nil, errors.New("invalid value for required argument 'StreamId'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource DbaasLogsInput
	err := ctx.RegisterResource("ovh:index/dbaasLogsInput:DbaasLogsInput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbaasLogsInput gets an existing DbaasLogsInput resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbaasLogsInput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbaasLogsInputState, opts ...pulumi.ResourceOption) (*DbaasLogsInput, error) {
	var resource DbaasLogsInput
	err := ctx.ReadResource("ovh:index/dbaasLogsInput:DbaasLogsInput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbaasLogsInput resources.
type dbaasLogsInputState struct {
	// IP blocks
	AllowedNetworks []string `pulumi:"allowedNetworks"`
	// Input configuration
	Configuration *DbaasLogsInputConfiguration `pulumi:"configuration"`
	// Input creation
	CreatedAt *string `pulumi:"createdAt"`
	// Input description
	Description *string `pulumi:"description"`
	// Input engine ID
	EngineId *string `pulumi:"engineId"`
	// Port
	ExposedPort *string `pulumi:"exposedPort"`
	// Hostname
	Hostname *string `pulumi:"hostname"`
	// Input ID
	InputId *string `pulumi:"inputId"`
	// Indicate if input need to be restarted
	IsRestartRequired *bool `pulumi:"isRestartRequired"`
	// Number of instance running
	NbInstance *int `pulumi:"nbInstance"`
	// Input IP address
	PublicAddress *string `pulumi:"publicAddress"`
	ServiceName   *string `pulumi:"serviceName"`
	// Input SSL certificate
	SslCertificate *string `pulumi:"sslCertificate"`
	// init: configuration required, pending: ready to start, running: available
	Status *string `pulumi:"status"`
	// Associated Graylog stream
	StreamId *string `pulumi:"streamId"`
	// Input title
	Title *string `pulumi:"title"`
	// Input last update
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DbaasLogsInputState struct {
	// IP blocks
	AllowedNetworks pulumi.StringArrayInput
	// Input configuration
	Configuration DbaasLogsInputConfigurationPtrInput
	// Input creation
	CreatedAt pulumi.StringPtrInput
	// Input description
	Description pulumi.StringPtrInput
	// Input engine ID
	EngineId pulumi.StringPtrInput
	// Port
	ExposedPort pulumi.StringPtrInput
	// Hostname
	Hostname pulumi.StringPtrInput
	// Input ID
	InputId pulumi.StringPtrInput
	// Indicate if input need to be restarted
	IsRestartRequired pulumi.BoolPtrInput
	// Number of instance running
	NbInstance pulumi.IntPtrInput
	// Input IP address
	PublicAddress pulumi.StringPtrInput
	ServiceName   pulumi.StringPtrInput
	// Input SSL certificate
	SslCertificate pulumi.StringPtrInput
	// init: configuration required, pending: ready to start, running: available
	Status pulumi.StringPtrInput
	// Associated Graylog stream
	StreamId pulumi.StringPtrInput
	// Input title
	Title pulumi.StringPtrInput
	// Input last update
	UpdatedAt pulumi.StringPtrInput
}

func (DbaasLogsInputState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbaasLogsInputState)(nil)).Elem()
}

type dbaasLogsInputArgs struct {
	// IP blocks
	AllowedNetworks []string `pulumi:"allowedNetworks"`
	// Input configuration
	Configuration DbaasLogsInputConfiguration `pulumi:"configuration"`
	// Input description
	Description string `pulumi:"description"`
	// Input engine ID
	EngineId string `pulumi:"engineId"`
	// Port
	ExposedPort *string `pulumi:"exposedPort"`
	// Number of instance running
	NbInstance  *int   `pulumi:"nbInstance"`
	ServiceName string `pulumi:"serviceName"`
	// Associated Graylog stream
	StreamId string `pulumi:"streamId"`
	// Input title
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a DbaasLogsInput resource.
type DbaasLogsInputArgs struct {
	// IP blocks
	AllowedNetworks pulumi.StringArrayInput
	// Input configuration
	Configuration DbaasLogsInputConfigurationInput
	// Input description
	Description pulumi.StringInput
	// Input engine ID
	EngineId pulumi.StringInput
	// Port
	ExposedPort pulumi.StringPtrInput
	// Number of instance running
	NbInstance  pulumi.IntPtrInput
	ServiceName pulumi.StringInput
	// Associated Graylog stream
	StreamId pulumi.StringInput
	// Input title
	Title pulumi.StringInput
}

func (DbaasLogsInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbaasLogsInputArgs)(nil)).Elem()
}

type DbaasLogsInputInput interface {
	pulumi.Input

	ToDbaasLogsInputOutput() DbaasLogsInputOutput
	ToDbaasLogsInputOutputWithContext(ctx context.Context) DbaasLogsInputOutput
}

func (*DbaasLogsInput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbaasLogsInput)(nil))
}

func (i *DbaasLogsInput) ToDbaasLogsInputOutput() DbaasLogsInputOutput {
	return i.ToDbaasLogsInputOutputWithContext(context.Background())
}

func (i *DbaasLogsInput) ToDbaasLogsInputOutputWithContext(ctx context.Context) DbaasLogsInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbaasLogsInputOutput)
}

func (i *DbaasLogsInput) ToDbaasLogsInputPtrOutput() DbaasLogsInputPtrOutput {
	return i.ToDbaasLogsInputPtrOutputWithContext(context.Background())
}

func (i *DbaasLogsInput) ToDbaasLogsInputPtrOutputWithContext(ctx context.Context) DbaasLogsInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbaasLogsInputPtrOutput)
}

type DbaasLogsInputPtrInput interface {
	pulumi.Input

	ToDbaasLogsInputPtrOutput() DbaasLogsInputPtrOutput
	ToDbaasLogsInputPtrOutputWithContext(ctx context.Context) DbaasLogsInputPtrOutput
}

type dbaasLogsInputPtrType DbaasLogsInputArgs

func (*dbaasLogsInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DbaasLogsInput)(nil))
}

func (i *dbaasLogsInputPtrType) ToDbaasLogsInputPtrOutput() DbaasLogsInputPtrOutput {
	return i.ToDbaasLogsInputPtrOutputWithContext(context.Background())
}

func (i *dbaasLogsInputPtrType) ToDbaasLogsInputPtrOutputWithContext(ctx context.Context) DbaasLogsInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbaasLogsInputPtrOutput)
}

// DbaasLogsInputArrayInput is an input type that accepts DbaasLogsInputArray and DbaasLogsInputArrayOutput values.
// You can construct a concrete instance of `DbaasLogsInputArrayInput` via:
//
//          DbaasLogsInputArray{ DbaasLogsInputArgs{...} }
type DbaasLogsInputArrayInput interface {
	pulumi.Input

	ToDbaasLogsInputArrayOutput() DbaasLogsInputArrayOutput
	ToDbaasLogsInputArrayOutputWithContext(context.Context) DbaasLogsInputArrayOutput
}

type DbaasLogsInputArray []DbaasLogsInputInput

func (DbaasLogsInputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbaasLogsInput)(nil)).Elem()
}

func (i DbaasLogsInputArray) ToDbaasLogsInputArrayOutput() DbaasLogsInputArrayOutput {
	return i.ToDbaasLogsInputArrayOutputWithContext(context.Background())
}

func (i DbaasLogsInputArray) ToDbaasLogsInputArrayOutputWithContext(ctx context.Context) DbaasLogsInputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbaasLogsInputArrayOutput)
}

// DbaasLogsInputMapInput is an input type that accepts DbaasLogsInputMap and DbaasLogsInputMapOutput values.
// You can construct a concrete instance of `DbaasLogsInputMapInput` via:
//
//          DbaasLogsInputMap{ "key": DbaasLogsInputArgs{...} }
type DbaasLogsInputMapInput interface {
	pulumi.Input

	ToDbaasLogsInputMapOutput() DbaasLogsInputMapOutput
	ToDbaasLogsInputMapOutputWithContext(context.Context) DbaasLogsInputMapOutput
}

type DbaasLogsInputMap map[string]DbaasLogsInputInput

func (DbaasLogsInputMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbaasLogsInput)(nil)).Elem()
}

func (i DbaasLogsInputMap) ToDbaasLogsInputMapOutput() DbaasLogsInputMapOutput {
	return i.ToDbaasLogsInputMapOutputWithContext(context.Background())
}

func (i DbaasLogsInputMap) ToDbaasLogsInputMapOutputWithContext(ctx context.Context) DbaasLogsInputMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbaasLogsInputMapOutput)
}

type DbaasLogsInputOutput struct{ *pulumi.OutputState }

func (DbaasLogsInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbaasLogsInput)(nil))
}

func (o DbaasLogsInputOutput) ToDbaasLogsInputOutput() DbaasLogsInputOutput {
	return o
}

func (o DbaasLogsInputOutput) ToDbaasLogsInputOutputWithContext(ctx context.Context) DbaasLogsInputOutput {
	return o
}

func (o DbaasLogsInputOutput) ToDbaasLogsInputPtrOutput() DbaasLogsInputPtrOutput {
	return o.ToDbaasLogsInputPtrOutputWithContext(context.Background())
}

func (o DbaasLogsInputOutput) ToDbaasLogsInputPtrOutputWithContext(ctx context.Context) DbaasLogsInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DbaasLogsInput) *DbaasLogsInput {
		return &v
	}).(DbaasLogsInputPtrOutput)
}

type DbaasLogsInputPtrOutput struct{ *pulumi.OutputState }

func (DbaasLogsInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbaasLogsInput)(nil))
}

func (o DbaasLogsInputPtrOutput) ToDbaasLogsInputPtrOutput() DbaasLogsInputPtrOutput {
	return o
}

func (o DbaasLogsInputPtrOutput) ToDbaasLogsInputPtrOutputWithContext(ctx context.Context) DbaasLogsInputPtrOutput {
	return o
}

func (o DbaasLogsInputPtrOutput) Elem() DbaasLogsInputOutput {
	return o.ApplyT(func(v *DbaasLogsInput) DbaasLogsInput {
		if v != nil {
			return *v
		}
		var ret DbaasLogsInput
		return ret
	}).(DbaasLogsInputOutput)
}

type DbaasLogsInputArrayOutput struct{ *pulumi.OutputState }

func (DbaasLogsInputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbaasLogsInput)(nil))
}

func (o DbaasLogsInputArrayOutput) ToDbaasLogsInputArrayOutput() DbaasLogsInputArrayOutput {
	return o
}

func (o DbaasLogsInputArrayOutput) ToDbaasLogsInputArrayOutputWithContext(ctx context.Context) DbaasLogsInputArrayOutput {
	return o
}

func (o DbaasLogsInputArrayOutput) Index(i pulumi.IntInput) DbaasLogsInputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbaasLogsInput {
		return vs[0].([]DbaasLogsInput)[vs[1].(int)]
	}).(DbaasLogsInputOutput)
}

type DbaasLogsInputMapOutput struct{ *pulumi.OutputState }

func (DbaasLogsInputMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DbaasLogsInput)(nil))
}

func (o DbaasLogsInputMapOutput) ToDbaasLogsInputMapOutput() DbaasLogsInputMapOutput {
	return o
}

func (o DbaasLogsInputMapOutput) ToDbaasLogsInputMapOutputWithContext(ctx context.Context) DbaasLogsInputMapOutput {
	return o
}

func (o DbaasLogsInputMapOutput) MapIndex(k pulumi.StringInput) DbaasLogsInputOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DbaasLogsInput {
		return vs[0].(map[string]DbaasLogsInput)[vs[1].(string)]
	}).(DbaasLogsInputOutput)
}

func init() {
	pulumi.RegisterOutputType(DbaasLogsInputOutput{})
	pulumi.RegisterOutputType(DbaasLogsInputPtrOutput{})
	pulumi.RegisterOutputType(DbaasLogsInputArrayOutput{})
	pulumi.RegisterOutputType(DbaasLogsInputMapOutput{})
}
