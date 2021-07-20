package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type vpcSetupCreateHelper struct{

	ctx *pulumi.Context

	vpc *ec2.Vpc
	subnets map[string]*ec2.Subnet
	internetGateways map[string] *ec2.InternetGateway
	routeTables map[string] *ec2.RouteTable
	routes map[string]*ec2.Route
	routeTableAssociations map[string]map[string] *ec2.RouteTableAssociation
	eips map[string]*ec2.Eip
	natGateways map[string]*ec2.NatGateway
	securityGroups map[string]*ec2.SecurityGroup

	conf *VpcConf

	vpcSetup *vpcSetup
}

func (v *vpcSetupCreateHelper) DONE() *vpcSetup{
	return v.vpcSetup
}

func (v *vpcSetupCreateHelper) GetVpc() *ec2.Vpc{
	return v.vpc
}
func (v *vpcSetupCreateHelper) GetSubnets() map[string]*ec2.Subnet{
	return v.subnets
}
func (v *vpcSetupCreateHelper) GetEips() map[string]*ec2.Eip{
	return v.eips
}
func (v *vpcSetupCreateHelper) GetSecurityGroups() map[string]*ec2.SecurityGroup{
	return v.securityGroups
}

func (v *vpcSetupCreateHelper) VPC(cidr string) *vpcSetupCreateHelper{
	v.vpc, _ = ec2.NewVpc(v.ctx, v.conf.Name.R(), &ec2.VpcArgs{
		CidrBlock:          pulumi.String(cidr),
		EnableDnsHostnames: TRUE,
		EnableDnsSupport:   TRUE,
		InstanceTenancy:    pulumi.StringPtr("default"),
		Tags:               v.conf.GetTags(),
	})

	return v
}
func (v *vpcSetupCreateHelper) VPCWithArgs(args *ec2.VpcArgs) *vpcSetupCreateHelper{
	v.vpc, _ = ec2.NewVpc(v.ctx, v.conf.Name.R(), args)

	return v
}

func (v *vpcSetupCreateHelper) PublicSubnet(name, cidr, zone string) *vpcSetupCreateHelper{
	availabilityZoneName := fmt.Sprintf("%s%s", v.conf.Region, zone)
	v.subnets[name], _ = ec2.NewSubnet(v.ctx, name, &ec2.SubnetArgs{
		AvailabilityZone:    pulumi.StringPtr(availabilityZoneName),
		CidrBlock:           pulumi.String(cidr),
		MapPublicIpOnLaunch: TRUE,
		Tags:                v.conf.GetTags(),
		VpcId:               v.vpc.ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) PrivateSubnet(name, cidr, zone string) *vpcSetupCreateHelper{
	availabilityZoneName := fmt.Sprintf("%s%s", v.conf.Region, zone)
	v.subnets[name], _ = ec2.NewSubnet(v.ctx, name, &ec2.SubnetArgs{
		AvailabilityZone:    pulumi.StringPtr(availabilityZoneName),
		CidrBlock:           pulumi.String(cidr),
		MapPublicIpOnLaunch: FALSE,
		Tags:                v.conf.GetTags(),
		VpcId:               v.vpc.ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) InternetGateway(name string) *vpcSetupCreateHelper{
	v.internetGateways[name], _ = ec2.NewInternetGateway(v.ctx, name, &ec2.InternetGatewayArgs{
		Tags:    v.conf.GetTags(),
		VpcId:   v.vpc.ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) RouteTable(name string) *vpcSetupCreateHelper{
	v.routeTables[name], _ = ec2.NewRouteTable(v.ctx, name, &ec2.RouteTableArgs{
		Tags:            v.conf.GetTags(),
		VpcId:           v.vpc.ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) InternetRoute(name, routeTable, destinationCidr, internetGateway string) *vpcSetupCreateHelper{
	v.routes[name], _ = ec2.NewRoute(v.ctx, name, &ec2.RouteArgs{
		RouteTableId: v.routeTables[routeTable].ID(),
		DestinationCidrBlock: pulumi.String(destinationCidr),
		GatewayId: v.internetGateways[internetGateway].ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) NatRoute(name, routeTable, destinationCidr, natGateway string) *vpcSetupCreateHelper{
	v.routes[name], _ = ec2.NewRoute(v.ctx, name, &ec2.RouteArgs{
		RouteTableId: v.routeTables[routeTable].ID(),
		DestinationCidrBlock: pulumi.String(destinationCidr),
		NatGatewayId: v.natGateways[natGateway].ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) SubnetToRouteTableAssociation(subnet, routeTable string) *vpcSetupCreateHelper{

	var err error

	if v.routeTableAssociations[subnet] == nil {
		v.routeTableAssociations[subnet] = map[string]*ec2.RouteTableAssociation{}
	}
	v.routeTableAssociations[subnet][routeTable], err = ec2.NewRouteTableAssociation(v.ctx, fmt.Sprintf("%s-%s", subnet, routeTable), &ec2.RouteTableAssociationArgs{
		SubnetId: v.subnets[subnet].ID(),
		RouteTableId: v.routeTables[routeTable].ID(),
	})

	if err != nil {
		panic(err)
	}

	return v
}

func (v *vpcSetupCreateHelper) Eip(name string) *vpcSetupCreateHelper{
	v.eips[name], _ = ec2.NewEip(v.ctx, name, &ec2.EipArgs{
		Tags:                   v.conf.GetTags(),
	})

	return v
}

func (v *vpcSetupCreateHelper) NatGateway(name, eip, subnet string) *vpcSetupCreateHelper{
	v.natGateways[name], _ = ec2.NewNatGateway(v.ctx, name, &ec2.NatGatewayArgs{
		AllocationId: v.eips[eip].ID(),
		SubnetId: v.subnets[subnet].ID(),
	})

	return v
}

func (v *vpcSetupCreateHelper) SecurityGroup(name, description string) *securityGroup{

	var s securityGroup
	return s.
		SetName(name).
		SetDescription(description).
		SetVpcSetupCreateHelperObject(v)
}