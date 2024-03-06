package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Azure/azure-sdk-for-go/profiles/latest/kubernetesmanagement/mgmt/containerservice"
    "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
    "github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
    // Authenticate with Azure
    authorizer, err := auth.NewAuthorizerFromEnvironment()
    if err != nil {
        log.Fatalf("Failed to get OAuth config: %v", err)
    }

    // Create a Kubernetes client
    kubeClient := containerservice.NewManagedClustersClient("<subscriptionID>")
    kubeClient.Authorizer = authorizer

    // List Kubernetes clusters
    clusters, err := kubeClient.List(context.Background(), "<resourceGroupName>")
    if err != nil {
        log.Fatalf("Failed to list Kubernetes clusters: %v", err)
    }

    // Check if the cluster exists
    clusterName := "<clusterName>"
    exists := false
    for _, cluster := range clusters.Values() {
        if *cluster.Name == clusterName {
            exists = true
            break
        }
    }

    if exists {
        fmt.Printf("Kubernetes cluster %s exists.\n", clusterName)
    } else {
        fmt.Printf("Kubernetes cluster %s does not exist.\n", clusterName)
    }
}