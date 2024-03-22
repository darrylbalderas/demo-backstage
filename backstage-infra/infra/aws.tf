resource "aws_s3_bucket" "backstage" {
  bucket = "demo-backstage-us-west-2-03192024"

  tags = {
    Application = "demo-backstage"
    Environment = "Development"
  }
}
