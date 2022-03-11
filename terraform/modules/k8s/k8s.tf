
provider "kubernetes" {
  host                   = var.host
  client_certificate     = var.client_certificate
  client_key             = var.client_key
  cluster_ca_certificate = var.cluster_ca_certificate
}

locals {
  products_service_port = 8080
  products_service_name = "products-servic"
}

resource "kubernetes_ingress" "products_api_gateway" {

  metadata {
    name = "products-api-gateway"

    labels = {
      name = "products-api-gateway"
    }

    annotations = {
      "kubernetes.io/ingress.class"                = "nginx"
      "nginx.ingress.kubernetes.io/ssl-redirect"   = "false"
      #"cert-manager.io/cluster-issuer"             = "letsencrypt-staging"
      "nginx.ingress.kubernetes.io/rewrite-target" = "/$2"
      "nginx.ingress.kubernetes.io/use-regex"      = "true"
    }
  }

  spec {
    rule {
      host = "aks-demo-341b2aa2.hcp.westeurope.azmk8s.io"

      http {
        path {
          path = "/api(/|$)(.*)"
          backend {
            service_name = "products-cluster-ip-service"
            service_port = local.products_service_port
          }
        }
      }
    }
  }

  wait_for_load_balancer = true
}

resource "kubernetes_deployment" "products_deployment" {
  metadata {
    name = "products-deployment"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = local.products_service_name
      }
    }

    template {
      metadata {
        labels = {
          app = local.products_service_name
        }
      }

      spec {
        container {
          name  = local.products_service_name
          image = "kolan1992/products.service:simple"

          env {
            name  = "environment"
            value = "kube"
          }

          image_pull_policy = "Always"
        }
      }
    }
  }
}

resource "kubernetes_service" "products_cluster_ip_service" {
  metadata {
    name = "products-cluster-ip-service"
  }

  spec {
    port {
      name        = "http"
      protocol    = "TCP"
      port        = local.products_service_port
      target_port = local.products_service_port
    }

    selector = {
      app = local.products_service_name
    }

    type = "ClusterIP"
  }
}

resource "kubernetes_service" "products_node_port_service" {
  metadata {
    name = "products-node-port-service"
  }

  spec {
    port {
      name        = local.products_service_name
      protocol    = "TCP"
      port        = local.products_service_port
      target_port = local.products_service_port
      node_port   = 31001
    }

    selector = {
      app = local.products_service_name
    }

    type = "NodePort"
  }
}


