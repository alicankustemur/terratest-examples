variable "region" {
  type = string
}

variable "alias_name" {
  description = "The name of the key alias"
  type        = string
}

variable "deletion_window_in_days" {
  description = "The duration in days after which the key is deleted after destruction of the resource"
  type        = string
  default     = 7
}

variable "description" {
  description = "The description of this KMS key"
  type        = string
  default     = "description"
}

variable "environment" {
  description = "The environment this KMS key belongs to"
  type        = string
}

variable "product_domain" {
  description = "The name of the product domain"
  type        = string
  default     = "test"
}

variable "additional_tags" {
  type        = map(string)
  description = "Additional tags to be added to kms-cmk"
  default     = {}
}
