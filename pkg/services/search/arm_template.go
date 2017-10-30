package search

// nolint: lll
var armTemplateBytes = []byte(`
{
  "$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {
    "location": {
      "type": "string"
    },
    "searchServiceName": {
      "type": "string",
      "minLength": 2,
      "maxLength": 60,
      "metadata": {
        "description": "Service name must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and is limited between 2 and 60 characters in length."
      }
    },
    "searchServiceSku": {
      "type": "string",
      "defaultValue": "standard",
      "allowedValues": [
        "free",
        "basic",
        "standard",
        "standard2",
        "standard3"
      ],
      "metadata": {
        "description": "The SKU of the search service you want to create. E.g. free or standard"
      }
    },
    "replicaCount": {
      "type": "int",
      "minValue": 1,
      "maxValue": 12,
      "defaultValue": 1,
      "metadata": {
        "description": "Replicas distribute search workloads across the service. You need 2 or more to support high availability (applies to Basic and Standard only)."
      }
    },
    "partitionCount": {
      "type": "int",
      "allowedValues": [
        1,
        2,
        3,
        4,
        6,
        12
      ],
      "defaultValue": 1,
      "metadata": {
        "description": "Partitions allow for scaling of document count as well as faster indexing by sharding your index over multiple Azure Search units."
      }
    },
    "hostingMode": {
      "type": "string",
      "defaultValue": "default",
      "allowedValues": [
        "default",
        "highDensity"
      ],
      "metadata": {
        "description": "Applicable only for SKU set to standard3. You can set this property to enable a single, high density partition that allows up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU."
      }
    },
		"tags": {
			"type": "object"
		}
  },
  "resources": [
    {
      "apiVersion": "2015-08-19",
      "name": "[parameters('searchServiceName')]",
      "type": "Microsoft.Search/searchServices",
      "location": "[parameters('location')]",
      "tags": "[parameters('tags')]",
      "sku": {
        "name": "[toLower(parameters('searchServiceSku'))]"
      },
      "properties": {
        "replicaCount": "[parameters('replicaCount')]",
        "partitionCount": "[parameters('partitionCount')]",
        "hostingMode": "[parameters('hostingMode')]"
      }
    }
  ],
  "outputs": {
    "apiKey": {
        "type": "string",
        "value": "[listAdminKeys(resourceId('Microsoft.Search/searchServices', parameters('searchServiceName')), '2015-08-19').primaryKey]"
    },
    "searchServiceName": {
      "type": "string",
      "value": "[parameters('searchServiceName')]"
    }
  }   
}
`)
