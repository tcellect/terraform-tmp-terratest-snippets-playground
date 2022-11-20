terraform {
  required_version = "~> 1.3.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.71.0"
    }
  }
}

provider "aws" {
  region = var.test-region
  profile = "deniss_orlov"
}