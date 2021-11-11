package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := support.NewFromConfig(cfg)

	input := &support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("ja"),
	}
	resp, err := svc.DescribeTrustedAdvisorChecks(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, v := range resp.Checks {
		fmt.Println(v)
	}

}
