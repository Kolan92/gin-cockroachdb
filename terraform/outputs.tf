# cluster

# output "kube_config" {
#   value     = module.cluster.kube_config
#   sensitive = true
# }

# output "cluster_ca_certificate" {
#   value = module.cluster.cluster_ca_certificate
# }

# output "client_certificate" {
#   value = module.cluster.client_certificate
# }

# output "client_key" {
#   value = module.cluster.client_key
# }

output "host" {
  value = module.cluster.host
}

# k8s
output "load_balancer_hostname" {
  value = module.k8s.load_balancer_hostname
}

output "load_balancer_ip" {
  value = module.k8s.load_balancer_ip
}

