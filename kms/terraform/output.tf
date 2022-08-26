output "description" {
  value = lookup(aws_kms_key.key.tags_all, "Description")
}