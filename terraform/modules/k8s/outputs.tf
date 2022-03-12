output "load_balancer_hostname" {
  value = kubernetes_ingress.products_api_gateway.status[0].load_balancer[0].ingress[*].hostname
}

output "load_balancer_ip" {
  value = kubernetes_ingress.products_api_gateway.status[0].load_balancer[0].ingress[*].ip
}
