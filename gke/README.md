# pulumi-sdk-castai
Example of Pulumi converted sdk from CASTAI Terraform provider

## Requirements

- pulumi cli
- go

## Trying it out

Clone the repo

Create a local pulumi stack
```bash
pulumi login --local

pulumi stack init
```

Set the CASTAI API TOKEN
```bash
pulumi config set castai:apiToken <CAST_API_TOKEN> --secret
```

Populate the subnet, projectid, clustername, etc., in the main.go then save.

```go
		projectId := "<YOUR_PROJECT>"
		clusterName := "<EXISTING_GKE_CLUSTER_NAME"
		gcpLocation := "<GKE_LOCATION>"
		subnets := []string{
			"<SUBNET>",
		}
		enableAutoScaler := true
		enableUnschedulablePods := true
		aggressiveModeEnabled := false
		enableEvictor := false
		enableDownScaler := false
```


To preview/plan
```bash
pulumi preview
```

To create/apply the preview
```bash
pulumi up
```

To destroy
```bash
pulumi destroy
```


