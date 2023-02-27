"# astrid-be" 
# Run elasticsearch
make docker-up
# Run go app
go run main.go

# Insert example

endpoint: http://localhost:3000/api/v1/kinds
payload:
{
  "apiVersion": "backstage.io/v1alpha1",
  "kind": "Component",
  "metadata": {
    "name": "name",
    "annotations": "annotation",
    "description": "The place to be, for great artists",
    "links": "https://admin.example-org.com"
  },
  "spec": {
    "lifecycle": "production",
    "owner": "artist-relations-team",
    "type": "website"
  }
}

# Read example
endpoint: http://localhost:3000/api/v1/kinds/{id}