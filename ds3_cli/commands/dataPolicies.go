package commands

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func getDataPoliciesList(client *ds3.Client, args *Arguments) (*models.DataPolicyList, error) {
    request := models.NewGetDataPoliciesSpectraS3Request()
    response, err := client.GetDataPoliciesSpectraS3(request)
    if err != nil {
        return nil, fmt.Errorf("Could not get Data Policies\n%v", err)
    }
    return &response.DataPolicyList, nil
}

func getDataPolicies(client *ds3.Client, args *Arguments) error {
    policies, err := getDataPoliciesList(client, args)
    if err != nil {
        return fmt.Errorf("Could not get Policies\n%v", err)
    }

    for _, object := range policies.DataPolicies {
        fmt.Printf("Name: %s, Id: %s\n", *object.Name, object.Id)
    }
    return nil
}

func getDataPolicy(client *ds3.Client, args *Arguments) (*models.DataPolicy ,error) {
    policies, err := getDataPoliciesList(client, args)
    if err != nil {
        return nil, fmt.Errorf("Could not get Policies\n%v", err)
    }

    for _, object := range policies.DataPolicies {
        if *object.Name == args.DataPolicy {
            return &object, nil
        }
    }
    return nil, fmt.Errorf("data Policy %s not found", args.DataPolicy)
}
