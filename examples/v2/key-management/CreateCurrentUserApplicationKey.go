// Create an application key for current user returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ApplicationKeyCreateRequest{
		Data: datadogV2.ApplicationKeyCreateData{
			Attributes: datadogV2.ApplicationKeyCreateAttributes{
				Name: "Application Key for managing dashboards",
				Scopes: []string{
					"dashboards_read",
					"dashboards_write",
					"dashboards_public_share",
				},
			},
			Type: datadogV2.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewKeyManagementApi(apiClient)
	resp, r, err := api.CreateCurrentUserApplicationKey(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateCurrentUserApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `KeyManagementApi.CreateCurrentUserApplicationKey`:\n%s\n", responseContent)
}
