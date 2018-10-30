package helpers

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "sync"
)

type putTransceiver struct {
    BucketName string
    WriteObjects *[]helperModels.PutObject
    Strategy *WriteTransferStrategy
    Client *ds3.Client
}

func newPutTransceiver(bucketName string, writeObjects *[]helperModels.PutObject, strategy *WriteTransferStrategy, client *ds3.Client) *putTransceiver {
    return &putTransceiver{
        BucketName:bucketName,
        WriteObjects:writeObjects,
        Strategy:strategy,
        Client:client,
    }
}

// Creates the bulk put request from the list of write objects and put bulk job options
func newBulkPutRequest(bucketName string, writeObjects *[]helperModels.PutObject, options WriteBulkJobOptions) *ds3Models.PutBulkJobSpectraS3Request {
    var putObjects []ds3Models.Ds3PutObject
    for _, obj := range *writeObjects {
        putObjects = append(putObjects, obj.PutObject)
    }

    bulkPut := ds3Models.NewPutBulkJobSpectraS3Request(bucketName, putObjects)
    if options.Aggregating != nil {
        bulkPut = bulkPut.WithAggregating(*options.Aggregating)
    }
    if options.ImplicitJobIdResolution != nil {
        bulkPut = bulkPut.WithImplicitJobIdResolution(*options.ImplicitJobIdResolution)
    }
    if options.MaxUploadSize != nil {
        bulkPut = bulkPut.WithMaxUploadSize(*options.MaxUploadSize)
    }
    if options.MinimizeSpanningAcrossMedia != nil {
        bulkPut = bulkPut.WithMinimizeSpanningAcrossMedia(*options.MinimizeSpanningAcrossMedia)
    }
    if options.Priority != nil {
        bulkPut = bulkPut.WithPriority(*options.Priority)
    }
    if options.VerifyAfterWrite != nil {
        bulkPut = bulkPut.WithVerifyAfterWrite(*options.VerifyAfterWrite)
    }
    if options.Name != nil {
        bulkPut = bulkPut.WithName(*options.Name)
    }
    if options.Force != nil && *options.Force {
        bulkPut = bulkPut.WithForce()
    }
    if options.IgnoreNamingConflicts != nil && *options.IgnoreNamingConflicts {
        bulkPut = bulkPut.WithIgnoreNamingConflicts()
    }
    return bulkPut
}

func (transceiver *putTransceiver) transfer() error {
    // create bulk put job
    bulkPut := newBulkPutRequest(transceiver.BucketName, transceiver.WriteObjects, transceiver.Strategy.Options)

    bulkPutResponse, err := transceiver.Client.PutBulkJobSpectraS3(bulkPut)
    if err != nil {
        return err
    }

    // init queue, producer and consumer
    var waitGroup sync.WaitGroup

    queue := newOperationQueue(transceiver.Strategy.BlobStrategy.maxWaitingTransfers())
    producer := newPutProducer(&bulkPutResponse.MasterObjectList, transceiver.WriteObjects, &queue, transceiver.Strategy, transceiver.Client, &waitGroup)
    consumer := newConsumer(&queue, &waitGroup, transceiver.Strategy.BlobStrategy.maxConcurrentTransfers())

    // Wait for completion of producer-consumer goroutines
    waitGroup.Add(2)  // adding producer and consumer goroutines to wait group

    var aggErr ds3Models.AggregateError
    go producer.run(&aggErr) // producer will add to waitGroup for every blob added to queue, and each transfer performed will decrement from waitGroup
    go consumer.run()
    waitGroup.Wait()

    return aggErr.GetErrors()
}

/*
func (transfernator *putTransceiver) cancel() {
    panic("Not yet implemented")
}
*/