terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.37.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.4.0"
    }
  }

  required_version = "~> 1.3.0"
}