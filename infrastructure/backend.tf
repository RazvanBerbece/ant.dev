terraform {
  cloud {
    organization = "antonio-dev-org"

    workspaces {
      name = "ant-dev-webapp"
    }
  }
}