// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package azurefetcher

import (
	"context"

	"github.com/elastic/elastic-agent-libs/logp"

	"github.com/elastic/cloudbeat/internal/inventory"
	azurelib "github.com/elastic/cloudbeat/internal/resources/providers/azurelib/inventory"
)

type resourceGraphFetcher struct {
	logger   *logp.Logger
	provider resourceGraphProvider
}

type (
	resourceGraphProvider interface {
		ListAllAssetTypesByName(ctx context.Context, assetGroup string, assets []string) ([]azurelib.AzureAsset, error)
	}
)

func newResourceGraphFetcher(logger *logp.Logger, provider resourceGraphProvider) inventory.AssetFetcher {
	return &resourceGraphFetcher{
		logger:   logger,
		provider: provider,
	}
}

func (f *resourceGraphFetcher) Fetch(ctx context.Context, assetChan chan<- inventory.AssetEvent) {
	resourcesToFetch := []struct {
		name           string
		azureGroup     string
		azureType      string
		classification inventory.AssetClassification
	}{
		{"App Services", azurelib.AssetGroupResources, azurelib.WebsitesAssetType, inventory.AssetClassificationAzureAppService},
		{"Container Registries", azurelib.AssetGroupResources, azurelib.ContainerRegistryAssetType, inventory.AssetClassificationAzureContainerRegistry},
		{"Cosmos DB Accounts", azurelib.AssetGroupResources, azurelib.DocumentDBDatabaseAccountAssetType, inventory.AssetClassificationAzureCosmosDBAccount},
		{"Cosmos DB SQL Databases", azurelib.AssetGroupResources, azurelib.CosmosDBForSQLDatabaseAssetType, inventory.AssetClassificationAzureCosmosDBSQLDatabase},
		{"Disks", azurelib.AssetGroupResources, azurelib.DiskAssetType, inventory.AssetClassificationAzureDisk},
		{"Elastic Pools", azurelib.AssetGroupResources, azurelib.ElasticPoolAssetType, inventory.AssetClassificationAzureElasticPool},
		{"MySQL Flexible Servers", azurelib.AssetGroupResources, azurelib.FlexibleMySQLDBAssetType, inventory.AssetClassificationAzureSQLServer},
		{"Resource Groups", azurelib.AssetGroupResourceContainers, azurelib.ResouceGroupAssetType, inventory.AssetClassificationAzureResourceGroup},
		{"SQL Databases", azurelib.AssetGroupResources, azurelib.MySQLDatabaseAssetType, inventory.AssetClassificationAzureSQLDatabase},
		{"Snapshots", azurelib.AssetGroupResources, azurelib.SnapshotAssetType, inventory.AssetClassificationAzureSnapshot},
		{"Storage Accounts", azurelib.AssetGroupResources, azurelib.StorageAccountAssetType, inventory.AssetClassificationAzureStorageAccount},
		{"Virtual Machines", azurelib.AssetGroupResources, azurelib.VirtualMachineAssetType, inventory.AssetClassificationAzureVirtualMachine},
	}
	for _, r := range resourcesToFetch {
		f.fetch(ctx, r.name, r.azureGroup, r.azureType, r.classification, assetChan)
	}
}

func (f *resourceGraphFetcher) fetch(ctx context.Context, resourceName, resourceGroup, resourceType string, classification inventory.AssetClassification, assetChan chan<- inventory.AssetEvent) {
	f.logger.Infof("Fetching %s", resourceName)
	defer f.logger.Infof("Fetching %s - Finished", resourceName)

	azureAssets, err := f.provider.ListAllAssetTypesByName(ctx, resourceGroup, []string{resourceType})
	if err != nil {
		f.logger.Errorf("Could not fetch %s: %v", resourceName, err)
		return
	}

	for _, item := range azureAssets {
		name := item.Name
		if name == "" {
			name = item.DisplayName
		}
		assetChan <- inventory.NewAssetEvent(
			classification,
			[]string{item.Id},
			name,
			inventory.WithRawAsset(item),
			inventory.WithCloud(inventory.AssetCloud{
				Provider: inventory.AzureCloudProvider,
				Account: inventory.AssetCloudAccount{
					Id: item.TenantId,
				},
				Service: &inventory.AssetCloudService{
					Name: "Azure",
				},
			}),
		)
	}
}