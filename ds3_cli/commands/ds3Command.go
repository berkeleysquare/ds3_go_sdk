package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
)

type command func(*ds3.Client, *Arguments) error

var availableCommands = map[string]command {
    "get_service": getService,
    "get_bucket": getBucket,
    "get_object": getObject,
    "head_object": headObject,
    "get_tape": getTape,

    "put_bucket": putSpectraBucket,
    "put_object": putObject,

    "delete_bucket": deleteBucket,
    "delete_object": deleteObject,

    "bulk_get": bulkGet,
    "bulk_put": bulkPut,

	"get_users": getUsers,
    "get_user": getUser,
    "create_user": createUser,
    "setup_user": setupUser,

    "get_data_policies": getDataPolicies,

    "performance_test":  performanceTest,
}

func RunCommand(client *ds3.Client, args *Arguments) error {
    cmd, ok := availableCommands[args.Command]
    if ok {
        return cmd(client, args)
    } else {
        return fmt.Errorf("Unsupported command: '%s'", args.Command)
    }
}

func ListCommands(args *Arguments) error {
    fmt.Printf("Usage: ds3_cli --command <command>\n",)
    for key, _ := range availableCommands {
        fmt.Printf("%s\n", key)
    }
    return nil
}

