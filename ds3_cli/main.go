package main

import (
    "log"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_cli/commands"
)

func main() {
    // Parse the arguments.
    args, argsErr := commands.ParseArgs()
    if argsErr != nil {
        if argsErr.Error() == "Must specify a command." {
            commands.ListCommands(args)
            return
        }
        log.Fatal(argsErr)
    }

    if args.Command == "list_commands" {
        commands.ListCommands(args)
        return
    }

    // Build the client.
    client, clientErr := buildclient.FromArgs(args)
    if clientErr != nil {
        log.Fatal(clientErr)
    }

    // Run the command
    if cmdErr := commands.RunCommand(client, args); cmdErr != nil {
        log.Fatal(cmdErr)
    }
}



