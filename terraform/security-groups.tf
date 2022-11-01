resource "aws_security_group" "node_group_private" {
  name_prefix = "node_group_private"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"

    cidr_blocks = [
      var.vpc_cidr,
    ]
  }
}

resource "aws_security_group" "node_group_public" {
  name_prefix = "node_group_public"
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"

    cidr_blocks = [
      "0.0.0.0/0",
    ]
  }
}
