resource "google_service_account" "github-tf-manager-service-account" {
  project      = var.PROJECT_ID
  account_id   = "gh-tf-manager"
  display_name = "TF Manager (GitHub Access)"
}

resource "google_service_account" "github-webapp-manager-service-account" {
  project      = var.PROJECT_ID
  account_id   = "gh-webapp-manager"
  display_name = "Webapp Manager (GitHub Access)"
}

resource "google_service_account_iam_binding" "github-tf-manager-service-account-owner-iam" {
  service_account_id = google_service_account.github-tf-manager-service-account.name
  role               = "roles/owner"
  members = [
    "serviceAccount:${google_service_account.github-tf-manager-service-account.email}",
  ]
}

resource "google_service_account_iam_binding" "github-webapp-manager-service-account-owner-iam" {
  service_account_id = google_service_account.github-webapp-manager-service-account.name
  role               = "roles/owner"
  members = [
    "serviceAccount:${google_service_account.github-webapp-manager-service-account.email}",
  ]
}

resource "google_service_account_iam_binding" "github-webapp-manager-service-account-owner-iam" {
  service_account_id = google_service_account.github-webapp-manager-service-account.name
  role               = "roles/roles/run.admin"
  members = [
    "serviceAccount:${google_service_account.github-webapp-manager-service-account.email}",
  ]
}

module "gh_oidc" {
  source            = "terraform-google-modules/github-actions-runners/google//modules/gh-oidc"
  version           = "v3.1.2"
  project_id        = var.PROJECT_ID
  pool_id           = "cd-identity-pool"
  pool_display_name = "CD Identity Pool"
  provider_id       = "cd-provider-antdev"
  sa_mapping = {
    (google_service_account.github-webapp-manager-service-account.account_id) = {
      sa_name   = google_service_account.github-webapp-manager-service-account.name
      attribute = "*"
    }
  }
}