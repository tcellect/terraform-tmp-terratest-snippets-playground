resource "aws_s3_bucket" "terratest_bucket" {
  bucket = "terratest${var.bucket_name}"
  versioning {
    enabled = true
  }
}