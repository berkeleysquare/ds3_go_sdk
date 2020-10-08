package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func getSpectraUserList(client *ds3.Client, args *Arguments) (*models.SpectraUserList, error) {
    requestAll := models.NewGetUsersSpectraS3Request()
    responseAll, err := client.GetUsersSpectraS3(requestAll)
    if err != nil {
        return nil, fmt.Errorf("Could not get User List\n%v", err)
    }
    return &responseAll.SpectraUserList, nil
}

func getUsers(client *ds3.Client, args *Arguments) error {
    users, err := getSpectraUserList(client, args)
    if err != nil {
        return fmt.Errorf("Could not get Users\n%v", err)
    }

    for _, object := range users.SpectraUsers {
        fmt.Printf("Name: %s, Id: %s\n", *object.Name, object.Id)
    }
    return nil
}

func getSpectraUser(client *ds3.Client, args *Arguments) (*models.SpectraUser, error) {
    users, err := getSpectraUserList(client, args)
    if err != nil {
        return nil, fmt.Errorf("Could not get Users\n%v", err)
    }

    for _, object := range users.SpectraUsers {
        if *object.Name == args.User {
            return &object, nil
        }
    }
    return nil, fmt.Errorf("user %s not found", args.User )
}

func getUser(client *ds3.Client, args *Arguments) error {
    user, err := getSpectraUser(client, args)
    if err != nil {
        return fmt.Errorf("Could not get User\n%v", err)
    }
    fmt.Printf("Name: %s, Id: %s\n", *user.Name, user.Id)
    fmt.Printf("AccessKey: %s, SectretKey: %s\n", *user.AuthId, *user.SecretKey)
    return nil
}

func createSpectraUser(client *ds3.Client, args *Arguments) (*models.SpectraUser, error) {
    request := models.NewDelegateCreateUserSpectraS3Request(args.User)
    _, err := client.DelegateCreateUserSpectraS3(request)
    if err != nil {
        return nil, fmt.Errorf("Could not create Spectra user %s\n%v", args.User, err)
    }
    return getSpectraUser(client, args)
}

func createUser(client *ds3.Client, args *Arguments) error {
    user, err := createSpectraUser(client, args)
    if err != nil {
        return fmt.Errorf("Could not create user %s\n%v", args.User, err)
    }
    fmt.Printf("Name: %s, Id: %s\n", *user.Name, user.Id)
    fmt.Printf("AccessKey: %s, SectretKey: %s\n", *user.AuthId, *user.SecretKey)
    return nil
}

