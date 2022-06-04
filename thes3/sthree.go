package thes3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func HeadObject() {
	// ctx := context.Background()
	sess, err := session.NewSession(
		&aws.Config{
			Region:           aws.String("us-west-2"),
			Credentials:      credentials.NewStaticCredentials("1", "1", ""),
			Endpoint:         aws.String("http://localhost:4566"),
			S3ForcePathStyle: aws.Bool(true),
		},
	)
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	out, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String("examplebucket"),
		Key:    aws.String("happyface.jpg"),
	})
	if err != nil {
		// return nil, errors.Wrap(err, errStatFileMsg)
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println("aerr.Code", aerr.Code(), "message", aerr.Message())
			switch aerr.Code() {
			case "NotFound":
				fmt.Println("besides that NOT FOUND")
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	fmt.Println("out", out)
}
