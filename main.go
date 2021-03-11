package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	settings, err := auth.GetSettingsFromFile()
	if err != nil {
		log.Fatalf("unable to authorize with azure: %s\n", err)
	}

	fmt.Printf("export AZURE_SUBSCRIPTION_ID=%s\n", settings.Values[auth.SubscriptionID])
	fmt.Printf("export AZURE_TENANT_ID=%s\n", settings.Values[auth.TenantID])
	fmt.Printf("export AZURE_CLIENT_ID=%s\n", settings.Values[auth.ClientID])
	fmt.Printf("export AZURE_CLIENT_SECRET=%s\n", settings.Values[auth.ClientSecret])

	fmt.Printf("export AZURE_SUBSCRIPTION_ID_B64=%s\n", base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.SubscriptionID])))
	fmt.Printf("export AZURE_TENANT_ID_B64=%s\n", base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.TenantID])))
	fmt.Printf("export AZURE_CLIENT_ID_B64=%s\n", base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.ClientID])))
	fmt.Printf("export AZURE_CLIENT_SECRET_B64=%s\n", base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.ClientSecret])))
}
