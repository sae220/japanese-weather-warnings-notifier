package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	LAMBDA_DIR        = "../lambda"
	LAMBDA_BUILD_FILE = "build/main"
	LAMBDA_FILE       = "main.go"
	LAMBDA_NAME       = "JapaneseWeatherWarningsNotifyFunction"
	SCHEDULE_NAME     = "JapaneseWeatherWarningsNotifySchedule"
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
	buildCommand := exec.Command("go", "build", "-C", LAMBDA_DIR, "-o", LAMBDA_BUILD_FILE, LAMBDA_FILE)
	if err := buildCommand.Run(); err != nil {
		return nil, err
	}

	// Define Lambda Function
	lambdaFunction := awslambda.NewFunction(stack, jsii.String(LAMBDA_NAME), &awslambda.FunctionProps{
		FunctionName: jsii.String(LAMBDA_NAME),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Code:         awslambda.Code_FromAsset(jsii.String("../lambda/build"), nil),
		Handler:      jsii.String("main"),
		Environment: &map[string]*string{
			"LINE_CHANNEL_TOKEN": jsii.String(os.Getenv("LINE_CHANNEL_TOKEN")),
		},
	})

	// Define EventBridge Scheduler Rule
	rule := awsevents.NewRule(stack, jsii.String(SCHEDULE_NAME), &awsevents.RuleProps{
		RuleName: jsii.String(SCHEDULE_NAME),
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			WeekDay: jsii.String(os.Getenv("SCHEDULE_WEEKDAY")),
			Hour:    jsii.String(os.Getenv("SCHEDULE_TIME")[0:2]),
			Minute:  jsii.String(os.Getenv("SCHEDULE_TIME")[2:4]),
		}),
	})

	// Add Lambda Function to EventBridge Scheduler Rule
	rule.AddTarget(awseventstargets.NewLambdaFunction(lambdaFunction, &awseventstargets.LambdaFunctionProps{
		RetryAttempts: jsii.Number(2),
	}))

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
