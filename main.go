package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)




func setup(ctx *pulumi.Context){

	vpc := vpcSetup{}
	vpc.Init(ctx).
		SetRegion("us-east-1").
		SetName("custom").
		SetGroupName("konebone").
		SetEnvironmentName("dev").
		SetPulumiStackName("dev").
		CREATE.
			VPC("10.0.0.0/16").
			PublicSubnet("public1", "10.0.0.0/24", "a").
			PublicSubnet("public2", "10.0.1.0/24", "b").
			PrivateSubnet("private1", "10.0.2.0/24", "a").
			PrivateSubnet("private2", "10.0.3.0/24", "b").

			InternetGateway("igw").
			RouteTable("internet").
			InternetRoute("internet", "internet", "0.0.0.0/0", "igw").
			SubnetToRouteTableAssociation("public1", "internet").
			SubnetToRouteTableAssociation("public2", "internet").

			Eip("nat_gw1").
			Eip("nat_gw2").
			NatGateway("nat_gw1", "nat_gw1", "public1").
			NatGateway("nat_gw2", "nat_gw2", "public2").
			RouteTable("private1").
			RouteTable("private2").
			RouteTable("private3").
			NatRoute("internet_over_nat1", "private1", "0.0.0.0/0", "nat_gw1").
			NatRoute("internet_over_nat2", "private2", "0.0.0.0/0", "nat_gw2").
			SubnetToRouteTableAssociation("private1", "private1").
			SubnetToRouteTableAssociation("private2", "private2").

			SecurityGroup("bastionSg", "Bastion security group").
				IngressCIDR("91.222.250.80/28", 22, "tcp", "Allow access from Kharkiv HQ").
				IngressCIDR("0.0.0.0/0", 1194, "udp", "Allow OpenVPN").
				IngressCIDR("0.0.0.0/0", 443, "tcp", "Allow OpenVPN HTTPS port").
			Done().
			SecurityGroup("albSg", "ALB security group").
				IngressCIDR("0.0.0.0/0", 80, "tcp", "HTTP").
				IngressCIDR("0.0.0.0/0", 443, "tcp", "HTTPS").
			Done().
			SecurityGroup("webSg", "Web service security group").
				IngressSG("albSg", 80, "tcp", "Allow access from ALB").
				IngressSG("albSg", 443, "tcp", "Allow access from ALB").
				IngressSG("albSg", 8080, "tcp", "Allow access from ALB").
				IngressSG("albSg", 22, "tcp", "Allow access from ALB").
				IngressSG("bastionSg", 22, "tcp", "Allow access from Bastion").
				IngressSG("bastionSg", 8080, "tcp", "Allow access from Bastion").
				IngressCIDR("10.0.0.0/16", 8080, "tcp", "Allow access from VPC").
			Done().
			SecurityGroup("rdsSg", "RDS security group").
				IngressSG("webSg", 5432, "tcp", "Allow access from webSg").
				IngressSG("bastionSg", 5432, "tcp", "Allow access from Bastion").
			Done().
		DONE().
		Export()
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		setup(ctx)

		return nil
	})
}

func deployVpc(){
}
