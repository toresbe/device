terraform {
  backend "gcs" {
    bucket = "nais-device-tfstate"
    prefix = "bootstrap-api"
  }
}

provider "google" {
  project = "nais-device"
  region  = "europe-north1"
  version = "3.14"
}

provider "google-beta" {
  project = "nais-device"
  region  = "europe-north1"
  version = "3.14"
}
