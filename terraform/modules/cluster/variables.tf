variable "serviceprinciple_id" {
  sensitive = true
}

variable "serviceprinciple_key" {
  sensitive = true
}

variable "location" {
  default = "westeurope"
}

variable "kubernetes_version" {
  default = "1.21.9"
}

variable "ssh_key" {
  sensitive = true
}
