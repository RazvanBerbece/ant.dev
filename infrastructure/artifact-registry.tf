resource "google_artifact_registry_repository" "antdev-docker-repo" {
  location      = var.PROJECT_REGION
  repository_id = "antdev-docker-ar"
  description   = "Repository of Docker images for the ant.dev project"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository_iam_binding" "binding" {
  project    = google_artifact_registry_repository.antdev-docker-repo.project
  location   = google_artifact_registry_repository.antdev-docker-repo.location
  repository = google_artifact_registry_repository.antdev-docker-repo.name
  role       = "roles/artifactregistry.admin"
  members = [
    "serviceAccount:${google_service_account.github-webapp-manager-service-account.email}",
  ]
}