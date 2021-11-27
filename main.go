package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create CloudTrail client
	svc := cloudtrail.New(sess)

	resp, err := svc.DescribeTrails(&cloudtrail.DescribeTrailsInput{TrailNameList: nil})
	if err != nil {
		fmt.Println("Got error calling CreateTrail:")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Found", len(resp.TrailList), "trail(s)")
	fmt.Println("")

	var trailFound *cloudtrail.Trail = nil

	for _, trail := range resp.TrailList {
		fmt.Println("Trail name:  " + *trail.Name)
		fmt.Println("Bucket name: " + *trail.S3BucketName)
		fmt.Println("")

		if *trail.Name == "bugbash-instance-trail" {
			fmt.Println("found what we were looking for!")
			trailFound = trail
			break
		}
	}

	params := &cloudtrail.LookupEventsInput{EndTime: aws.Time(time.Now())}

	lookupEventsOutput, err := svc.LookupEvents(params)

}
