terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.36.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "5.36.0"
    }
  }
}

provider "google" {
  project = var.PROJECT_ID
  region  = var.PROJECT_REGION
  zone    = var.PROJECT_ZONE

  credentials = var.GCP_CREDENTIALS
}

provider "google-beta" {
  project = var.PROJECT_ID
}