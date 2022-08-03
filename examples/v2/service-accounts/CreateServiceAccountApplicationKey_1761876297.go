// Create an app key for this service account returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	// there is a valid "service_account_user" in the system
	ServiceAccountUserDataID := os.Getenv("SERVICE_ACCOUNT_USER_DATA_ID")

	body := datadog.ApplicationKeyCreateRequest{
		Data: datadog.ApplicationKeyCreateData{
			Attributes: datadog.ApplicationKeyCreateAttributes{
				Name: "Example-Create_an_app_key_for_this_service_account_returns_Created_response",
			},
			Type: datadog.APPLICATIONKEYSTYPE_APPLICATION_KEYS,
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewServiceAccountsApi(apiClient)
	resp, r, err := api.CreateServiceAccountApplicationKey(ctx, ServiceAccountUserDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountApplicationKey`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountApplicationKey`:\n%s\n", responseContent)
}
