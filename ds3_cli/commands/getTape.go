package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func getTape(client *ds3.Client, args *Arguments) error {
    request := models.NewGetBlobsOnTapeSpectraS3Request(args.Tape)
    response, err := client.GetBlobsOnTapeSpectraS3(request)
    if err != nil {
        return fmt.Errorf("Could not get blobs for tape %s\n%v", args.Tape, err)
    }

    for _, object := range response.BulkObjectList.Objects {
        err = processHeadObject(client, *object.Bucket, *object.Name)
        if err != nil {
            return fmt.Errorf("Failed head_object for Bucket %s, key %s\n%v",
                *object.Bucket, *object.Name, err)
        }
    }
    return nil
}


