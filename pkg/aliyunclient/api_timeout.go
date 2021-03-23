package aliyunclient

import (
	"encoding/json"
	"strings"
	"time"
)

var apiTimeouts = `{
  "ecs": {
      "ActivateRouterInterface": 10,
      "AddTags": 61,
      "AllocateDedicatedHosts": 10,
      "AllocateEipAddress": 17,
      "AllocatePublicIpAddress": 36,
      "ApplyAutoSnapshotPolicy": 10,
      "AssignIpv6Addresses": 10,
      "AssignPrivateIpAddresses": 10,
      "AssociateEipAddress": 17,
      "AttachClassicLinkVpc": 14,
      "AttachDisk": 36,
      "AttachInstanceRamRole": 11,
      "AttachKeyPair": 16,
      "AttachNetworkInterface": 16,
      "AuthorizeSecurityGroupEgress": 16,
      "AuthorizeSecurityGroup": 16,
      "CancelAutoSnapshotPolicy": 10,
      "CancelCopyImage": 10,
      "CancelPhysicalConnection": 10,
      "CancelSimulatedSystemEvents": 10,
      "CancelTask": 10,
      "ConnectRouterInterface": 10,
      "ConvertNatPublicIpToEip": 12,
      "CopyImage": 10,
      "CreateAutoSnapshotPolicy": 10,
      "CreateCommand": 16,
      "CreateDeploymentSet": 16,
      "CreateDisk": 36,
      "CreateHpcCluster": 10,
      "CreateImage": 36,
      "CreateInstance": 86,
      "CreateKeyPair": 10,
      "CreateLaunchTemplate": 10,
      "CreateLaunchTemplateVersion": 10,
      "CreateNatGateway": 36,
      "CreateNetworkInterfacePermission": 13,
      "CreateNetworkInterface": 16,
      "CreatePhysicalConnection": 10,
      "CreateRouteEntry": 17,
      "CreateRouterInterface": 10,
      "CreateSecurityGroup": 86,
      "CreateSimulatedSystemEvents": 10,
      "CreateSnapshot": 86,
      "CreateVirtualBorderRouter": 10,
      "CreateVpc": 16,
      "CreateVSwitch": 17,
      "DeactivateRouterInterface": 10,
      "DeleteAutoSnapshotPolicy": 10,
      "DeleteBandwidthPackage": 10,
      "DeleteCommand": 16,
      "DeleteDeploymentSet": 12,
      "DeleteDisk": 16,
      "DeleteHpcCluster": 10,
      "DeleteImage": 36,
      "DeleteInstance": 66,
      "DeleteKeyPairs": 10,
      "DeleteLaunchTemplate": 10,
      "DeleteLaunchTemplateVersion": 10,
      "DeleteNatGateway": 10,
      "DeleteNetworkInterfacePermission": 10,
      "DeleteNetworkInterface": 16,
      "DeletePhysicalConnection": 10,
      "DeleteRouteEntry": 16,
      "DeleteRouterInterface": 10,
      "DeleteSecurityGroup": 87,
      "DeleteSnapshot": 17,
      "DeleteVirtualBorderRouter": 10,
      "DeleteVpc": 17,
      "DeleteVSwitch": 17,
      "DescribeAccessPoints": 10,
      "DescribeAccountAttributes": 10,
      "DescribeAutoSnapshotPolicyEx": 16,
      "DescribeAvailableResource": 10,
      "DescribeBandwidthLimitation": 16,
      "DescribeBandwidthPackages": 10,
      "DescribeClassicLinkInstances": 15,
      "DescribeCloudAssistantStatus": 16,
      "DescribeClusters": 10,
      "DescribeCommands": 16,
      "DescribeDedicatedHosts": 10,
      "DescribeDedicatedHostTypes": 10,
      "DescribeDeploymentSets": 26,
      "DescribeDiskMonitorData": 16,
      "DescribeDisksFullStatus": 14,
      "DescribeDisks": 19,
      "DescribeEipAddresses": 16,
      "DescribeEipMonitorData": 16,
      "DescribeEniMonitorData": 10,
      "DescribeHaVips": 10,
      "DescribeHpcClusters": 16,
      "DescribeImageSharePermission": 10,
      "DescribeImages": 38,
      "DescribeImageSupportInstanceTypes": 16,
      "DescribeInstanceAttribute": 36,
      "DescribeInstanceAutoRenewAttribute": 17,
      "DescribeInstanceHistoryEvents": 19,
      "DescribeInstanceMonitorData": 19,
      "DescribeInstancePhysicalAttribute": 10,
      "DescribeInstanceRamRole": 11,
      "DescribeInstancesFullStatus": 14,
      "DescribeInstances": 10,
      "DescribeInstanceStatus": 26,
      "DescribeInstanceTopology": 12,
      "DescribeInstanceTypeFamilies": 17,
      "DescribeInstanceTypes": 17,
      "DescribeInstanceVncPasswd": 10,
      "DescribeInstanceVncUrl": 36,
      "DescribeInvocationResults": 16,
      "DescribeInvocations": 16,
      "DescribeKeyPairs": 12,
      "DescribeLaunchTemplates": 16,
      "DescribeLaunchTemplateVersions": 16,
      "DescribeLimitation": 36,
      "DescribeNatGateways": 10,
      "DescribeNetworkInterfacePermissions": 13,
      "DescribeNetworkInterfaces": 16,
      "DescribeNewProjectEipMonitorData": 16,
      "DescribePhysicalConnections": 10,
      "DescribePrice": 16,
      "DescribeRecommendInstanceType": 10,
      "DescribeRegions": 19,
      "DescribeRenewalPrice": 16,
      "DescribeResourceByTags": 10,
      "DescribeResourcesModification": 17,
      "DescribeRouterInterfaces": 10,
      "DescribeRouteTables": 17,
      "DescribeSecurityGroupAttribute": 133,
      "DescribeSecurityGroupReferences": 16,
      "DescribeSecurityGroups": 25,
      "DescribeSnapshotLinks": 17,
      "DescribeSnapshotMonitorData": 12,
      "DescribeSnapshotPackage": 10,
      "DescribeSnapshots": 26,
      "DescribeSnapshotsUsage": 26,
      "DescribeSpotPriceHistory": 22,
      "DescribeTags": 17,
      "DescribeTaskAttribute": 10,
      "DescribeTasks": 11,
      "DescribeUserBusinessBehavior": 13,
      "DescribeUserData": 10,
      "DescribeVirtualBorderRoutersForPhysicalConnection": 10,
      "DescribeVirtualBorderRouters": 10,
      "DescribeVpcs": 41,
      "DescribeVRouters": 17,
      "DescribeVSwitches": 17,
      "DescribeZones": 103,
      "DetachClassicLinkVpc": 14,
      "DetachDisk": 17,
      "DetachInstanceRamRole": 10,
      "DetachKeyPair": 10,
      "DetachNetworkInterface": 16,
      "EipFillParams": 19,
      "EipFillProduct": 13,
      "EipNotifyPaid": 10,
      "EnablePhysicalConnection": 10,
      "ExportImage": 10,
      "GetInstanceConsoleOutput": 14,
      "GetInstanceScreenshot": 14,
      "ImportImage": 29,
      "ImportKeyPair": 10,
      "InstallCloudAssistant": 10,
      "InvokeCommand": 16,
      "JoinResourceGroup": 10,
      "JoinSecurityGroup": 66,
      "LeaveSecurityGroup": 66,
      "ModifyAutoSnapshotPolicyEx": 10,
      "ModifyBandwidthPackageSpec": 11,
      "ModifyCommand": 10,
      "ModifyDeploymentSetAttribute": 10,
      "ModifyDiskAttribute": 16,
      "ModifyDiskChargeType": 13,
      "ModifyEipAddressAttribute": 14,
      "ModifyImageAttribute": 10,
      "ModifyImageSharePermission": 16,
      "ModifyInstanceAttribute": 22,
      "ModifyInstanceAutoReleaseTime": 15,
      "ModifyInstanceAutoRenewAttribute": 16,
      "ModifyInstanceChargeType": 22,
      "ModifyInstanceDeployment": 10,
      "ModifyInstanceNetworkSpec": 36,
      "ModifyInstanceSpec": 62,
      "ModifyInstanceVncPasswd": 35,
      "ModifyInstanceVpcAttribute": 15,
      "ModifyLaunchTemplateDefaultVersion": 10,
      "ModifyNetworkInterfaceAttribute": 10,
      "ModifyPhysicalConnectionAttribute": 10,
      "ModifyPrepayInstanceSpec": 13,
      "ModifyRouterInterfaceAttribute": 10,
      "ModifySecurityGroupAttribute": 10,
      "ModifySecurityGroupEgressRule": 10,
      "ModifySecurityGroupPolicy": 10,
      "ModifySecurityGroupRule": 16,
      "ModifySnapshotAttribute": 10,
      "ModifyUserBusinessBehavior": 10,
      "ModifyVirtualBorderRouterAttribute": 10,
      "ModifyVpcAttribute": 10,
      "ModifyVRouterAttribute": 10,
      "ModifyVSwitchAttribute": 10,
      "ReActivateInstances": 10,
      "RebootInstance": 27,
      "RedeployInstance": 14,
      "ReInitDisk": 16,
      "ReleaseDedicatedHost": 10,
      "ReleaseEipAddress": 16,
      "ReleasePublicIpAddress": 10,
      "RemoveTags": 10,
      "RenewInstance": 19,
      "ReplaceSystemDisk": 36,
      "ResetDisk": 36,
      "ResizeDisk": 11,
      "RevokeSecurityGroupEgress": 13,
      "RevokeSecurityGroup": 16,
      "RunInstances": 86,
      "StartInstance": 46,
      "StopInstance": 27,
      "StopInvocation": 10,
      "TerminatePhysicalConnection": 10,
      "TerminateVirtualBorderRouter": 10,
      "UnassignIpv6Addresses": 10,
      "UnassignPrivateIpAddresses": 10,
      "UnassociateEipAddress": 16 
  }
}
`

func getAPIMaxTimeout(product, actionName string) (time.Duration, bool) {
	timeout := make(map[string]map[string]int)
	err := json.Unmarshal([]byte(apiTimeouts), &timeout)
	if err != nil {
		return 0 * time.Millisecond, false
	}

	obj := timeout[strings.ToLower(product)]
	if obj != nil && obj[actionName] != 0 {
		return time.Duration(obj[actionName]) * time.Second, true
	}

	return 0 * time.Millisecond, false
}
