resource "google_cloud_run_v2_service" "ant-dev-webapp-container-instance" {
  name     = "ant-dev-webapp-service"
  location = var.PROJECT_REGION
  ingress  = "INGRESS_TRAFFIC_INTERNAL_ONLY"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}