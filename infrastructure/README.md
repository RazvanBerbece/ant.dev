# Cloud Infrastructure for `ant.dev`

The containerised webapp is deployed to Google Cloud Run, mainly because it offers ~100K free requests and plenty of free CPU time a month. 

Most of the Google Cloud infrastructure (service accounts, artifact registry, role assignments and billing stuff) is managed through Terraform and deployed automatically via GitHub Actions.

The container app is managed and deployed by the `gcloud run deploy` command which manages both the creation of the container resource on-Cloud and the update of the revision with the new code. This workflow is automated via GitHub Actions.