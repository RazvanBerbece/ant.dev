resource "google_service_account" "github-tf-manager-service-account" {
  project      = var.PROJECT_ID
  account_id   = "gh-tf-manager"
  display_name = "TF Manager (GitHub Access)"
}

resource "google_service_account_iam_binding" "github-tf-manager-service-account-owner-iam" {
  service_account_id = google_service_account.github-tf-manager-service-account.name
  role               = "roles/owner"
  members = [
    "serviceAccount:${google_service_account.github-tf-manager-service-account.email}",
  ]
}

resource "google_service_account_iam_binding" "github-tf-manager-service-account-billing-iam" {
  service_account_id = google_service_account.github-tf-manager-service-account.name
  role               = "roles/billing.admin"
  members = [
    "serviceAccount:${google_service_account.github-tf-manager-service-account.email}",
  ]
}