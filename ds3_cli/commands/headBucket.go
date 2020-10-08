package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func headBucket(client *ds3.Client, args *Arguments) error {
    if args.Bucket == "" {
        return fmt.Errorf("must specify a bucket")
    }

    request := models.NewHeadBucketRequest(args.Bucket)
    _, err := client.HeadBucket(request)
    return err
}
