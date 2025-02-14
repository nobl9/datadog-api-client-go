// Create role returns "OK" response

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
	body := datadogV2.RoleCreateRequest{
		Data: datadogV2.RoleCreateData{
			Type: datadogV2.ROLESTYPE_ROLES.Ptr(),
			Attributes: datadogV2.RoleCreateAttributes{
				Name: "Example-Create_role_returns_OK_response",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRolesApi(apiClient)
	resp, r, err := api.CreateRole(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateRole`:\n%s\n", responseContent)
}
