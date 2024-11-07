package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) (awscdk.Stack, error) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Build go
	buildCommand := exec.Command("go", "build", "../lambda/main.go")
	if err := buildCommand.Run(); err != nil {
		return nil, err
	}

	return stack, nil
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	_, err := NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_REGION")),
	}
}
