package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"regexp"
)

type securityGroup struct {
	Name         String
	Description  String
	ingressRules []securityGroupRule
	egressRules  []securityGroupRule

	vpcSetupCreateHelper *vpcSetupCreateHelper
}

func (s *securityGroup) SetName(name string) *securityGroup{
	s.Name = String(name)

	return s
}
func (s *securityGroup) SetDescription(description string) *securityGroup{
	s.Description = String(description)

	return s
}
func (s *securityGroup) SetVpcSetupCreateHelperObject(v *vpcSetupCreateHelper) *securityGroup{
	s.vpcSetupCreateHelper = v

	return s
}

func (s *securityGroup) handleIngressRulesForDone() ec2.SecurityGroupIngressArray{
	var pulumiIngressRules ec2.SecurityGroupIngressArray
	if s.ingressRules != nil {
		for _, i := range s.ingressRules {

			var cidrBlocks = pulumi.StringArray{}
			if i.CidrBlock.R() != ""{
				cidrBlocks = append(cidrBlocks, i.CidrBlock.S())
			} else {
				cidrBlocks = nil
			}

			var securityGroups = pulumi.StringArray{}
			sgIdRegex := regexp.MustCompile(`sg-[a-z0-9]*`)
			if i.SecurityGroup.R() != ""{
				// checking if the security group provided corresponds to AWS SecurityGroupID
				if sgIdRegex.MatchString(i.SecurityGroup.R()){
					securityGroups = append(securityGroups, i.SecurityGroup.S())
				// or just our internal naming
				} else {
					securityGroups = append(securityGroups, s.vpcSetupCreateHelper.securityGroups[i.SecurityGroup.R()].ID())
				}
			} else {
				securityGroups = nil
			}

			pulumiIngressRules = append(pulumiIngressRules, ec2.SecurityGroupIngressArgs{
				CidrBlocks:     cidrBlocks,
				Description:    i.Description.SP(),
				FromPort:       pulumi.Int(i.ToPort),
				Protocol:       i.Protocol.S(),
				SecurityGroups: securityGroups,
				ToPort:         pulumi.Int(i.ToPort),
			})
		}
		if len(s.ingressRules) == 0{
			pulumiIngressRules = nil
		}
	}else {
		pulumiIngressRules = nil
	}

	return pulumiIngressRules
}

func (s *securityGroup) handleEgressRulesForDone() ec2.SecurityGroupEgressArray{
	var pulumiEgressRules ec2.SecurityGroupEgressArray
	if s.egressRules != nil {
		for _, i := range s.egressRules {

			var cidrBlocks = pulumi.StringArray{}
			if i.CidrBlock.R() != ""{
				cidrBlocks = append(cidrBlocks, i.CidrBlock.S())
			} else {
				cidrBlocks = nil
			}

			var securityGroups = pulumi.StringArray{}
			sgIdRegex := regexp.MustCompile(`sg-[a-z0-9]*`)
			if i.SecurityGroup.R() != ""{
				// checking if the security group provided corresponds to AWS SecurityGroupID
				if sgIdRegex.MatchString(i.SecurityGroup.R()){
					securityGroups = append(securityGroups, i.SecurityGroup.S())
					// or just our internal naming
				} else {
					securityGroups = append(securityGroups, s.vpcSetupCreateHelper.securityGroups[i.SecurityGroup.R()].ID())
				}
			} else {
				securityGroups = nil
			}

			pulumiEgressRules = append(pulumiEgressRules, ec2.SecurityGroupEgressArgs{
				CidrBlocks:     cidrBlocks,
				Description:    i.Description.SP(),
				FromPort:       pulumi.Int(i.ToPort),
				Protocol:       i.Protocol.S(),
				SecurityGroups: securityGroups,
				ToPort:         pulumi.Int(i.ToPort),
			})
		}
		if len(s.egressRules) == 0{
			pulumiEgressRules = nil
		}
	}else {
		pulumiEgressRules = nil
	}

	return pulumiEgressRules
}

func (s *securityGroup) Done() *vpcSetupCreateHelper{

	var pulumiIngressRules = s.handleIngressRulesForDone()
	var pulumiEgressRules = s.handleEgressRulesForDone()

	s.vpcSetupCreateHelper.securityGroups[s.Name.R()], _ = ec2.NewSecurityGroup(s.vpcSetupCreateHelper.ctx, s.Name.R(), &ec2.SecurityGroupArgs{

		Description:         s.Description.SP(),
		Egress:              pulumiEgressRules,
		Ingress:             pulumiIngressRules,
		Name:                s.Name.SP(),
		RevokeRulesOnDelete: TRUE,
		Tags:                s.vpcSetupCreateHelper.conf.GetTags(),
		VpcId:               s.vpcSetupCreateHelper.vpc.ID(),
	})

	return s.vpcSetupCreateHelper
}

func (s *securityGroup) ingress(cidr, securityGroup String, toPort int, protocol, description String) *securityGroup{
	s.ingressRules = append(s.ingressRules, securityGroupRule{
		CidrBlock:     cidr,
		SecurityGroup: securityGroup,
		Description:   description,
		ToPort:        toPort,
		Protocol:      protocol,
	})

	return s
}

func (s *securityGroup) IngressCIDR(cidr String, toPort int, protocol, description String) *securityGroup{
	s.ingress(cidr, "", toPort, protocol, description)

	return s
}

func (s *securityGroup) IngressSG(securityGroup String, toPort int, protocol, description String) *securityGroup{
	s.ingress("", securityGroup, toPort, protocol, description)

	return s
}

func (s *securityGroup) egress(cidr, securityGroup String, toPort int, protocol, description String) *securityGroup{
	s.egressRules = append(s.egressRules, securityGroupRule{
		CidrBlock:     cidr,
		SecurityGroup: securityGroup,
		Description:   description,
		ToPort:        toPort,
		Protocol:      protocol,
	})

	return s
}

func (s *securityGroup) EgressCIDR(cidr String, toPort int, protocol, description String) *securityGroup{
	s.egress(cidr, "", toPort, protocol, description)

	return s
}

func (s *securityGroup) EgressSG(securityGroup String, toPort int, protocol, description String) *securityGroup{
	s.egress(securityGroup, "", toPort, protocol, description)

	return s
}

type securityGroupRule struct {
	CidrBlock String
	SecurityGroup String
	Description String
	ToPort int
	Protocol String
}
