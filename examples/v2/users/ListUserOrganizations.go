// Get a user organization returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.ListUserOrganizations(ctx, "00000000-0000-9999-0000-000000000000")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserOrganizations`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUserOrganizations`:\n%s\n", responseContent)
}
