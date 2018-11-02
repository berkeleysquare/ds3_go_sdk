package helpers_integration

import (
    "bytes"
    "testing"
    "log"
    "os"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "github.com/SpectraLogic/ds3_go_sdk/helpers"
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_integration/utils"
    "io/ioutil"
    "github.com/SpectraLogic/ds3_go_sdk/samples/utils"
    "github.com/SpectraLogic/ds3_go_sdk/helpers/channels"
    "io"
)

var client *ds3.Client
var testBucket = "GoHelperTestBucket"
var envTestNameSpace = "GoHelperTest"
var defaultUser = "Administrator"

func TestMain(m *testing.M) {
    var err error
    var exitVal int
    var ids *testutils.TestIds
    client, ids, err = testutils.SetupTestEnv(testBucket, defaultUser, envTestNameSpace)
    if err != nil {
        log.Printf("Unable to setup test environment '%s'.", err.Error())
        exitVal = 1
    } else {
        exitVal = m.Run()
    }
    testutils.TeardownTestEnv(client, ids, testBucket)
    os.Exit(exitVal)
}

func TestPutBulk(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy()

    writeObjects, err := getTestBooksAsWriteObjects()
    ds3Testing.AssertNilError(t, err)

    err = helper.PutObjects(testBucket, *writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    // verify all books are on BP
    getBucket, getBucketErr := client.GetBucket(ds3Models.NewGetBucketRequest(testBucket))
    ds3Testing.AssertNilError(t, getBucketErr)
    if len(getBucket.ListBucketResult.Objects) != len(*writeObjects) {
        t.Fatalf("Expected '%d' objects in bucket '%s', but found '%d'.", len(*writeObjects), testBucket, len(getBucket.ListBucketResult.Objects))
    }
    for i, obj := range getBucket.ListBucketResult.Objects {
        ds3Testing.AssertNonNilStringPtr(t, "BookName", testutils.BookTitles[i], obj.Key)
    }

    testutils.VerifyFilesOnBP(t, testBucket, testutils.BookTitles, testutils.BookPath, client)
}

func TestPutBulkBlobSpanningChunksRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy()

    writeObj, err := getTestWriteObjectRandomAccess(LargeBookTitle, LargeBookPath + LargeBookTitle)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    err = helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)


    testutils.VerifyFilesOnBP(t, testBucket, []string {LargeBookTitle}, LargeBookPath, client)
}

func TestPutBulkBlobSpanningChunksStreamAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    helper := helpers.NewHelpers(client)

    strategy := newTestTransferStrategy()

    writeObj, err := getTestWriteObjectStreamAccess(LargeBookTitle, LargeBookPath + LargeBookTitle)

    var writeObjects []helperModels.PutObject
    writeObjects = append(writeObjects, *writeObj)

    ds3Testing.AssertNilError(t, err)

    err = helper.PutObjects(testBucket, writeObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    testutils.VerifyFilesOnBP(t, testBucket, []string {LargeBookTitle}, LargeBookPath, client)
}


func TestGetBulk(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)
    err := testutils.PutTestBooks(client, testBucket)
    ds3Testing.AssertNilError(t, err)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
    }

    file0, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file0.Close()
    defer os.Remove(file0.Name())

    file1, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file1.Close()
    defer os.Remove(file1.Name())

    file2, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file2.Close()
    defer os.Remove(file2.Name())

    file3, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file3.Close()
    defer os.Remove(file3.Name())

    readObjects := []helperModels.GetObject {
        {Name: testutils.BookTitles[0], ChannelBuilder: channels.NewWriteChannelBuilder(file0.Name())},
        {Name: testutils.BookTitles[1], ChannelBuilder: channels.NewWriteChannelBuilder(file1.Name())},
        {Name: testutils.BookTitles[2], ChannelBuilder: channels.NewWriteChannelBuilder(file2.Name())},
        {Name: testutils.BookTitles[3], ChannelBuilder: channels.NewWriteChannelBuilder(file3.Name())},
    }

    err = helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    utils.VerifyBookContent(testutils.BookTitles[0], file0)
    utils.VerifyBookContent(testutils.BookTitles[1], file1)
    utils.VerifyBookContent(testutils.BookTitles[2], file2)
    utils.VerifyBookContent(testutils.BookTitles[3], file3)
}

func TestGetBulkBlobSpanningChunksRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: channels.NewWriteChannelBuilder(file.Name())},
    }

    err = helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    err = VerifyLargeBookContent(file)
    ds3Testing.AssertNilError(t, err)
}

func TestGetBulkPartialObjectRandomAccess(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    LoadLargeFile(testBucket, client)

    helper := helpers.NewHelpers(client)

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{}, // use default job options
        BlobStrategy: newTestBlobStrategy(),
    }

    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    ranges := []ds3Models.Range {
        {0, 100},
        {200, 300},
        {301, 400},
        {500, 600},
    }

    readObjects := []helperModels.GetObject{
        {Name: LargeBookTitle, ChannelBuilder: channels.NewPartialObjectChannelBuilder(file.Name(), ranges), Ranges: ranges},
    }

    err = helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    file.Seek(0, io.SeekStart)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 101, 0, file)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 201, 200, file)
    testutils.VerifyPartialFile(t, LargeBookPath + LargeBookTitle, 101, 500, file)
}

func TestGettingBlobbedFile(t *testing.T) {
    defer testutils.DeleteBucketContents(client, testBucket)

    filePath, err := writeAFileThatWillGetBlobbed()
    ds3Testing.AssertNilError(t, err)

    defer os.Remove(filePath)

    putObject, err := getTestWriteObjectStreamAccess(filePath, filePath)
    ds3Testing.AssertNilError(t, err)

    helper := helpers.NewHelpers(client)

    err = helper.PutObjects(testBucket, []helperModels.PutObject{*putObject}, newTestTransferStrategy())
    ds3Testing.AssertNilError(t, err)

    renamedOriginalFile := filePath + ".orig"
    err = os.Rename(filePath, renamedOriginalFile)
    ds3Testing.AssertNilError(t, err)

    defer os.Remove(renamedOriginalFile)

    readObjects := []helperModels.GetObject {
        {Name: filePath, ChannelBuilder: channels.NewWriteChannelBuilder(filePath)},
    }

    strategy := helpers.ReadTransferStrategy{
        Options: helpers.ReadBulkJobOptions{},
        BlobStrategy: newTestBlobStrategy(),
    }

    err = helper.GetObjects(testBucket, readObjects, strategy)
    ds3Testing.AssertNilError(t, err)

    ds3Testing.AssertBool(t, "files have different content", true, filesAreTheSame(filePath, renamedOriginalFile))
}

func writeAFileThatWillGetBlobbed() (filePath string, err error) {
    filePath = "aBlobbedFile"

    file, err := os.Create(filePath)
    if err != nil {
        return
    }

    defer file.Close()

    biggerThanAChunkSize := helpers.MinUploadSize * 2 + 1024

    numIntsInBiggerThanAChunkSize := biggerThanAChunkSize / 4

    dataArray := make([]byte, 4)

    for i := int64(0); i < numIntsInBiggerThanAChunkSize; i++ {
        dataArray[0] = byte(i)
        dataArray[1] = byte(i >> 8)
        dataArray[2] = byte(i >> 16)
        dataArray[3] = byte(i >> 24)

        _, err = file.Write(dataArray)
        if err != nil {
            return
        }
    }

    return
}

const chunkSize = 64000

func filesAreTheSame(file1, file2 string) bool {
    f1, err := os.Open(file1)
    if err != nil {
        log.Fatal(err)
    }

    f2, err := os.Open(file2)
    if err != nil {
        log.Fatal(err)
    }

    for {
        b1 := make([]byte, chunkSize)
        _, err1 := f1.Read(b1)

        b2 := make([]byte, chunkSize)
        _, err2 := f2.Read(b2)

        if err1 != nil || err2 != nil {
            if err1 == io.EOF && err2 == io.EOF {
                return true
            } else if err1 == io.EOF || err2 == io.EOF {
                return false
            } else {
                log.Fatal(err1, err2)
            }
        }

        if !bytes.Equal(b1, b2) {
            return false
        }
    }
}
