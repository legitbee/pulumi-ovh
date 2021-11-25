// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Inputs
{

    public sealed class DedicatedServerInstallTaskDetailsGetArgs : Pulumi.ResourceArgs
    {
        [Input("changeLog")]
        public Input<string>? ChangeLog { get; set; }

        [Input("customHostname")]
        public Input<string>? CustomHostname { get; set; }

        [Input("diskGroupId")]
        public Input<int>? DiskGroupId { get; set; }

        [Input("installRtm")]
        public Input<bool>? InstallRtm { get; set; }

        [Input("installSqlServer")]
        public Input<bool>? InstallSqlServer { get; set; }

        [Input("language")]
        public Input<string>? Language { get; set; }

        [Input("noRaid")]
        public Input<bool>? NoRaid { get; set; }

        [Input("postInstallationScriptLink")]
        public Input<string>? PostInstallationScriptLink { get; set; }

        [Input("postInstallationScriptReturn")]
        public Input<string>? PostInstallationScriptReturn { get; set; }

        [Input("resetHwRaid")]
        public Input<bool>? ResetHwRaid { get; set; }

        [Input("softRaidDevices")]
        public Input<int>? SoftRaidDevices { get; set; }

        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        [Input("useDistribKernel")]
        public Input<bool>? UseDistribKernel { get; set; }

        [Input("useSpla")]
        public Input<bool>? UseSpla { get; set; }

        public DedicatedServerInstallTaskDetailsGetArgs()
        {
        }
    }
}