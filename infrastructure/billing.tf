resource "google_billing_budget" "budget" {

  billing_account = var.BILLING_ACCOUNT_ID
  display_name    = "Billing Budget"

  amount {
    specified_amount {
      currency_code = "GBP"
      units         = "15"
    }
  }

  threshold_rules {
    threshold_percent = 0.5
  }

  threshold_rules {
    threshold_percent = 0.25
    spend_basis       = "FORECASTED_SPEND"
  }

  all_updates_rule {
    monitoring_notification_channels = [
      google_monitoring_notification_channel.budget_notification_channel.id,
    ]
    disable_default_iam_recipients = false
    enable_project_level_recipients = true
  }
}

resource "google_monitoring_notification_channel" "budget_notification_channel" {
  display_name = "Main Budget Notification Channel"
  type         = "email"

  labels = {
    email_address = var.BUDGET_ALERT_CHANNEL_EMAIL
  }
}