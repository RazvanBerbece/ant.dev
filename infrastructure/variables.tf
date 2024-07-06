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