package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "os"
)

func setupUser(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return fmt.Errorf("must specify a bucket name with --bucket")
    }
    if args.User == "" {
        return fmt.Errorf("must specify a user with --user")
    }
    if args.DataPolicy == "" {
        return fmt.Errorf("must specify a data policy with --data_policy")
    }

    // ensure user does not exist
    user,_ := getSpectraUser(client, args)
    if user != nil {
        return fmt.Errorf("user %s already exists", args.User)
    }

    // ensure data policy exists
    policy, err := getDataPolicy(client, args)
    if err != nil {
        return fmt.Errorf("get data policy failed: %v", err)
    }
    if policy == nil {
        return fmt.Errorf("data policy %s does not exists: %v", args.DataPolicy, err)
    }

    // ensure bucket does not exist
    err = headBucket(client, args)
    if err == nil {
        return fmt.Errorf("bucket %s already exists", args.Bucket)
    }

    // verification passed
    // create BP user
    user, err = createSpectraUser(client, args)
    if err != nil {
        return fmt.Errorf("create user failed: %v", err)
    }

    // create bucket
    err = putSpectraBucket(client, args)
    if err != nil {
        return fmt.Errorf("create bucket failed: %v", err)
    }

    creds := fmt.Sprintf("Name: %s\r\n", *user.Name) +
    fmt.Sprintf("Bucket: %s\r\n", args.Bucket) +
    fmt.Sprintf("Data Policy: %s\r\n", args.DataPolicy) +
    fmt.Sprintf("AccessKey: %s, SecretKey: %s\r\n", *user.AuthId, *user.SecretKey)

    // write to console
    fmt.Printf(creds)

    // write to local file
    destPath := *user.Name + "_credentials.txt"
    destination, err := os.Create(destPath)
    if err != nil {
        return fmt.Errorf("could not create new file %s, %v", destPath, err)
    }
    defer destination.Close()

    _, err = destination.WriteString(creds + "\n")
    if err != nil {
        return fmt.Errorf("could not write credentials file %s, %v", destPath, err)
    }

    return nil
}

