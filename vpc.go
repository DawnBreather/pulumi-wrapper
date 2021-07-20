package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type vpcSetup struct{
	//vpc *ec2.Vpc
	//
	//subnets map[string]*ec2.Subnet
	//
	//internetGateway *ec2.InternetGateway
	//
	//routeTables map[string] *ec2.RouteTable
	//
	//routes map[string]*ec2.Route
	//
	//routeTableAssociations map[string]map[string] *ec2.RouteTableAssociation
	//
	//eips map[string]*ec2.Eip
	//
	//natGateways map[string]*ec2.NatGateway
	//
	//securityGroups map[string]*ec2.SecurityGroup

	CREATE *vpcSetupCreateHelper
}
func (v *vpcSetup) Items() vpcSetupCreateHelper{
	return *v.CREATE
}


func (v *vpcSetup) Init(ctx *pulumi.Context) *vpcSetup{
	v.CREATE = &vpcSetupCreateHelper{
		ctx:                    ctx,
		subnets:                map[string]*ec2.Subnet{},
		internetGateways:       map[string]*ec2.InternetGateway{},
		routeTables:            map[string]*ec2.RouteTable{},
		routes:                 map[string]*ec2.Route{},
		routeTableAssociations: map[string]map[string]*ec2.RouteTableAssociation{},
		eips:                   map[string]*ec2.Eip{},
		natGateways:            map[string]*ec2.NatGateway{},
		securityGroups:         map[string]*ec2.SecurityGroup{},
		conf:                   &VpcConf{},
		vpcSetup:               v,
	}

	return v
}

func (v *vpcSetup) SetName(name string) *vpcSetup{
	v.CREATE.conf.Name = String(name)
	return v
}
func (v *vpcSetup) SetRegion(region string) *vpcSetup{
	v.CREATE.conf.Region = String(region)
	return v
}
func (v *vpcSetup) SetGroupName(groupName string) *vpcSetup{
	v.CREATE.conf.GroupName = String(groupName)
	return v
}

func (v *vpcSetup) SetEnvironmentName(environmentName string) *vpcSetup{
	v.CREATE.conf.EnvironmentName = String(environmentName)
	return v
}

func (v *vpcSetup) SetPulumiStackName(pulumiStackName string) *vpcSetup{
	v.CREATE.conf.PulumiStackName = String(pulumiStackName)
	return v
}

func (v *vpcSetup) Export(){
	v.CREATE.ctx.Export(fmt.Sprintf("vpc-%s-name", v.Items().conf.Name), pulumi.String(v.Items().conf.Name))
}