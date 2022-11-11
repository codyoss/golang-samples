// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secretmanager

// [START secretmanager_enable_secret_version_with_etag]
import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// enableSecretVersionWithEtag enables the given secret version, enabling it to be
// accessed after previously being disabled. Other secrets versions are
// unaffected.
func enableSecretVersionWithEtag(name, etag string) error {
	// name := "projects/my-project/secrets/my-secret/versions/5"
	// etag := `"123"`

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.EnableSecretVersionRequest{
		Name: name,
		Etag: etag,
	}

	// Call the API.
	if _, err := client.EnableSecretVersion(ctx, req); err != nil {
		return fmt.Errorf("failed to enable secret version: %v", err)
	}
	return nil
}

// [END secretmanager_enable_secret_version_with_etag]
