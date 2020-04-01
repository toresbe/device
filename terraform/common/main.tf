terraform {
  backend "gcs" {
    bucket = "nais-device-tfstate"
    prefix = "common"
  }
}

provider "google" {
  project = "nais-device"
  region  = "europe-north1"
  version = "3.14"
}