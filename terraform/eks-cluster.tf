module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "18.30.2"

  cluster_name    = local.cluster_name
  cluster_version = "1.23"

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"

    attach_cluster_primary_security_group = false

    # Disabling and using externally provided security groups
    create_security_group = false
  }

  eks_managed_node_groups = {
    node_group_private = {
      name = "node-group-private"

      instance_types = ["t2.micro"]
      capacity_type  = "SPOT"

      min_size     = 1
      max_size     = 5
      desired_size = 3

      vpc_security_group_ids = [
        aws_security_group.node_group_private.id
      ]

      labels = {
        Access = "private"
      }

      tags = {
        Terraform   = "true"
      }
    }

    node_group_public = {
      name = "node-group-public"

      instance_types = ["t2.micro"]
      capacity_type  = "SPOT"

      min_size     = 1
      max_size     = 5
      desired_size = 3

      vpc_security_group_ids = [
        aws_security_group.node_group_public.id
      ]

      labels = {
        Access = "public"
      }

      tags = {
        Terraform   = "true"
      }

      taints = {
        dedicated = {
          key    = "Access"
          value  = "public"
          effect = "NO_SCHEDULE"
        }
      }

      # subnet_ids = module.vpc.private_subnets
    }
  }
}
