// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories is used to instantiate a provider during acceptance testing.
// The factory function is called for each Terraform CLI command to create a provider
// server that the CLI can connect to and interact with.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"n8ncloud": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_ACC"); v == "" {
		t.Fatal("TF_ACC must be set for acceptance tests")
	}
	if v := os.Getenv("N8N_API_KEY"); v == "" {
		t.Fatal("N8N_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("N8N_INSTANCE_URL"); v == "" {
		t.Fatal("N8N_INSTANCE_URL must be set for acceptance tests")
	}
}
