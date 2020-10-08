package commands

import (
    "errors"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "strconv"
)

func headObject(client *ds3.Client, args *Arguments) error {
    // Validate the arguments.
    if args.Bucket == "" && args.Tape == "" {
        return errors.New("Must specify bucket or tape.")
    }
    bucket := args.Bucket
    key := args.Key
    tape := args.Tape

    if tape != "" {
        // do whole tape
        more := true
        totalObjects := 0
        for marker := ""; more; {
            request := models.NewGetBlobsOnTapeSpectraS3Request(args.Tape)
            if len(marker) > 0 {
                request.PageStartMarker = &marker
            }

            response, err := client.GetBlobsOnTapeSpectraS3(request)
            if err != nil {
                return fmt.Errorf("Could not get blobs for tape %s\n%v", args.Tape, err)
            }

            objects := response.BulkObjectList.Objects

            for _, object := range objects {
                err = processHeadObject(client, *object.Bucket, *object.Name)
                if err != nil {
                    return fmt.Errorf("Failed head_object for Bucket %s, key %s\n%v",
                        *object.Bucket, *object.Name, err)
                }
            }

            totalObjects += len(objects)
            // GoSDK pagination decorators not implemented in GetBlobsOnTapeSpectraS3()
            // get next page id count == 1000
            more = len(objects) >= 1000
            marker = *objects[len(objects) - 1].Id
            fmt.Printf("Count: %d, marker: %s, more: %t\n", totalObjects, marker, more)
        }
        fmt.Printf("Accessed %d objects\n", totalObjects)
        return nil
    } else if key == "" {
        // do whole bucket
        request:= models.NewGetBucketRequestWithPrefix(bucket, args.KeyPrefix)
        response, err := client.GetBucket(request)
        if err != nil {
            return fmt.Errorf("Failed to get bucket %s\n%v", bucket, err)
        }

        for _, object := range response.ListBucketResult.Objects {
            err = processHeadObject(client, bucket, *object.Key)
            if err != nil {
                return fmt.Errorf("Failed head_object for Bucket %s, key %s\n%v",
                                    bucket, *object.Key, err)
            }
        }
        return nil
    } else {
        // single object
        return processHeadObject(client, bucket, key)
    }
}

func processHeadObject(client *ds3.Client, bucket string, key string) error {
    // Build the request.
    request := models.NewHeadObjectRequest(bucket, key)

    // Perform the request.
    response, requestErr := client.HeadObject(request)
    if requestErr != nil {
        return requestErr
    }

    checksums := response.BlobChecksums
    checksumType := response.BlobChecksumType

    fmt.Printf("Bucket: %s\n", bucket)
    fmt.Printf("Key: %s\n", key)
    fmt.Printf("Type: %s\n", checksumType)
    for i, c := range checksums {
        fmt.Printf("%s: %s\n",strconv.Itoa(int(i)), c)
    }
    fmt.Printf("-----------------------------\n")

    return nil
}

