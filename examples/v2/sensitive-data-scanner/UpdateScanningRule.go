// Update Scanning Rule returns "OK" response

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
	// the "scanning_group" has a "scanning_rule"
	RuleDataID := os.Getenv("RULE_DATA_ID")

	// there is a valid "scanning_group" in the system
	GroupDataID := os.Getenv("GROUP_DATA_ID")

	body := datadogV2.SensitiveDataScannerRuleUpdateRequest{
		Meta: datadogV2.SensitiveDataScannerMetaVersionOnly{},
		Data: datadogV2.SensitiveDataScannerRuleUpdate{
			Id:   datadog.PtrString(RuleDataID),
			Type: datadogV2.SENSITIVEDATASCANNERRULETYPE_SENSITIVE_DATA_SCANNER_RULE.Ptr(),
			Attributes: &datadogV2.SensitiveDataScannerRuleAttributes{
				Name:    datadog.PtrString("Example-Update_Scanning_Rule_returns_OK_response"),
				Pattern: datadog.PtrString("pattern"),
				TextReplacement: &datadogV2.SensitiveDataScannerTextReplacement{
					Type: datadogV2.SENSITIVEDATASCANNERTEXTREPLACEMENTTYPE_NONE.Ptr(),
				},
				Tags: []string{
					"sensitive_data:true",
				},
				IsEnabled: datadog.PtrBool(true),
			},
			Relationships: &datadogV2.SensitiveDataScannerRuleRelationships{
				Group: &datadogV2.SensitiveDataScannerGroupData{
					Data: &datadogV2.SensitiveDataScannerGroup{
						Type: datadogV2.SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP.Ptr(),
						Id:   datadog.PtrString(GroupDataID),
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSensitiveDataScannerApi(apiClient)
	resp, r, err := api.UpdateScanningRule(ctx, RuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SensitiveDataScannerApi.UpdateScanningRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SensitiveDataScannerApi.UpdateScanningRule`:\n%s\n", responseContent)
}
