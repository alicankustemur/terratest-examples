resource "aws_kms_key" "key" {
  description             = var.description
  key_usage               = "ENCRYPT_DECRYPT"
  policy                  = data.aws_iam_policy_document.kms_key_policy.json
  deletion_window_in_days = var.deletion_window_in_days
  is_enabled              = true
  enable_key_rotation     = true
  tags = merge(
    {
      Description   = var.description
      Environment   = var.environment
      Name          = var.alias_name
      ProductDomain = var.product_domain
      ManagedBy     = "terraform"
    },
    var.additional_tags,
  )
}

resource "aws_kms_alias" "key_alias" {
  name          = "alias/${var.alias_name}"
  target_key_id = aws_kms_key.key.id
}
