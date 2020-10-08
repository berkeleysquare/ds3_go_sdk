package commands

import (
    "errors"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func putBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing put_bucket.")
    }

    // Run request.
    _, err := client.PutBucket(models.NewPutBucketRequest(args.Bucket))
    return err
}

func putSpectraBucket(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("must specify a bucket name")
    }
    if args.User == "" || args.DataPolicy == "" {
        // s3 bucket
        return putBucket(client, args)
    }

    // Spectra S3 Bucket
    request := models.NewPutBucketSpectraS3Request(args.Bucket).
        WithDataPolicyId(args.DataPolicy).
        WithUserId(args.User)

    _, err := client.PutBucketSpectraS3(request)
    return err
}

