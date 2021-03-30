package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type Data struct {
	ClientID       string `json:"client-id"`
	ClientSecret   string `json:"client-secret"`
	TenantID       string `json:"tenant-id"`
	SubscriptionID string `json:"subscription-id"`
}

type Secret struct {
	Data Data `json:"data"`
}

func main() {
	var format string
	flag.StringVar(&format, "output", "json", "Output format (json, env)")
	flag.Parse()

	settings, err := auth.GetSettingsFromFile()
	if err != nil {
		log.Fatalf("unable to authorize with azure: %s\n", err)
	}

	s := Secret{
		Data: Data{
			ClientID:       base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.ClientID])),
			ClientSecret:   base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.ClientSecret])),
			TenantID:       base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.TenantID])),
			SubscriptionID: base64.StdEncoding.EncodeToString([]byte(settings.Values[auth.SubscriptionID])),
		},
	}

	out, err := outputSecret(&s, format)
	if err != nil {
		log.Fatalf("unable to output secret: %s", err)
	}
	fmt.Print(out)
}

func outputSecret(s *Secret, format string) (string, error) {
	switch format {
	case "json":
		b, err := json.Marshal(&s)
		return string(b), err
	case "env":
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("export AZURE_SUBSCRIPTION_ID_B64=%s\n", s.Data.SubscriptionID))
		b.WriteString(fmt.Sprintf("export AZURE_TENANT_ID_B64=%s\n", s.Data.TenantID))
		b.WriteString(fmt.Sprintf("export AZURE_CLIENT_ID_B64=%s\n", s.Data.ClientID))
		b.WriteString(fmt.Sprintf("export AZURE_CLIENT_SECRET_B64=%s\n", s.Data.ClientSecret))
		return b.String(), nil
	default:
		return "", fmt.Errorf("unknown format %s", format)
	}
}
