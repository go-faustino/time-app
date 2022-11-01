variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-2"
}

variable "vpc_cidr" {
  description = "VPC CIDR"
  type    = string
  default = "10.0.0.0/16"
}

variable "availability_zone_names" {
  description = "AWS availability zones"
  type    = list(string)
  default = ["us-east-2a", "us-east-2b"]
}

variable "private_subnets" {
  description = "AWS private subnets"
  type    = list(string)
  default = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "public_subnets" {
  description = "AWS public subnets"
  type    = list(string)
  default = ["10.0.3.0/24", "10.0.4.0/24"]
}
