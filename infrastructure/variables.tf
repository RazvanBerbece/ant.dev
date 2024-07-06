variable "PROJECT_ID" {
  type        = string
  description = "The ID as shown on Google Cloud for the target project"
  sensitive   = true
}

variable "PROJECT_REGION" {
  type        = string
  description = "The location of the Google Cloud project"
  default     = "europe-west2"
}

variable "PROJECT_ZONE" {
  type        = string
  description = "The zone of the Google Cloud project"
  default     = "europe-west2-c"
}

variable "GCP_CREDENTIALS" {
  type        = string
  sensitive   = true
  description = "Google Cloud service account credentials"
}

variable "BUDGET_ALERT_CHANNEL_EMAIL" {
  type        = string
  sensitive   = true
  description = "Email to which budget usage alerts are sent to"
}

variable "BILLING_ACCOUNT_ID" {
  type        = string
  sensitive   = true
  description = "The ID of the billing account associated with the Google Cloud project"
}