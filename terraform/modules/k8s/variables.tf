variable "host" {
}

variable "client_certificate" {
  sensitive = true
}

variable "client_key" {
  sensitive = true
}

variable "cluster_ca_certificate" {
  sensitive = true
}
