resource "aws_ecr_repository" "foo" {
  name                 = "time-app-ecr"

  image_scanning_configuration {
    scan_on_push = true
  }
}
