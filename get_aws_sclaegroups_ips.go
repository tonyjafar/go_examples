package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type arrayFlags []string

var myFlags arrayFlags

func (i *arrayFlags) String() string {
	return "Scal Groups Names"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func getInstancesIDs(s string) [][]byte {
	svc := autoscaling.New(session.New())
	input := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(s),
		},
	}
	result, err := svc.DescribeAutoScalingGroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeInvalidNextToken:
				fmt.Println(autoscaling.ErrCodeInvalidNextToken, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	match := regexp.MustCompile(`InstanceId: \"(.*)\"`).FindAll([]byte(result.String()), -1)
	return match
}

func getInstances(s string) [][]byte {
	svc := ec2.New(session.New())
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(s),
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	match := regexp.MustCompile(`PublicIpAddress: \"(.*)\"`).FindAll([]byte(result.String()), -1)

	return match
}

func main() {
	os.Setenv("AWS_PROFILE", "ec2-start")
	os.Setenv("AWS_REGION", "eu-central-1")
	flag.Var(&myFlags, "scaleNames", "Scale Groups names sepated by a space")
	flag.Parse()
	servers := make(map[string][]string)
	if len(myFlags) == 0 {
		os.Exit(2)
	}
	for _, myFlag := range myFlags {
		IDs := getInstancesIDs(string(myFlag))
		for _, id := range IDs {
			myID := strings.Split(string(id), " ")[1][1 : len(strings.Split(string(id), " ")[1])-1]
			IPAddresses := getInstances(myID)
			for _, IP := range IPAddresses {
				ip := strings.Split(string(IP), " ")[1][1 : len(strings.Split(string(IP), " ")[1])-1]
				servers[myFlag] = append(servers[myFlag], string(ip))
			}
		}
	}
	for server, ip := range servers {
		fmt.Println(server, ip)
	}

}
