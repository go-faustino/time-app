module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.18.1"

  name = "go-faustino-test-vpc"

  cidr = var.vpc_cidr
  azs  = var.availability_zone_names

  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets

  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true

  # public_subnet_tags = {
  #   "kubernetes.io/cluster/${local.cluster_name}" = "shared"
  #   "kubernetes.io/role/elb"                      = 1
  # }

  # private_subnet_tags = {
  #   "kubernetes.io/cluster/${local.cluster_name}" = "shared"
  #   "kubernetes.io/role/internal-elb"             = 1
  # }
}
