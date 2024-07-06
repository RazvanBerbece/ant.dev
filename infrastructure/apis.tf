resource "google_project_service" "apis" {
  for_each = toset([
    "cloudbilling.googleapis.com",
    "iam.googleapis.com",
    "billingbudgets.googleapis.com"
  ])

  service            = each.value
  disable_on_destroy = true
}
