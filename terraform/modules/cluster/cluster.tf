resource "azurerm_resource_group" "aks-demo" {
  name     = "aks-demo"
  location = var.location
}

resource "azurerm_kubernetes_cluster" "aks-demo" {
  name                = "aks-demo"
  location            = azurerm_resource_group.aks-demo.location
  resource_group_name = azurerm_resource_group.aks-demo.name
  dns_prefix          = "aks-demo"
  kubernetes_version  = var.kubernetes_version

  default_node_pool {
    name       = "default"
    node_count = 3
    vm_size    = "standard_d2_v2"
    type       = "VirtualMachineScaleSets"
  }

  service_principal {
    client_id     = var.serviceprinciple_id
    client_secret = var.serviceprinciple_key
  }

  linux_profile {
    admin_username = "azureuser"
    ssh_key {
      key_data = var.ssh_key
    }
  }

  network_profile {
    network_plugin    = "kubenet"
    load_balancer_sku = "Standard"
  }

  addon_profile {
    http_application_routing {
      enabled = true
    }
  }
}

/*
resource "azurerm_kubernetes_cluster_node_pool" "monitoring" {
  name                  = "monitoring"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.aks-demo.id
  vm_size               = "Standard_DS2_v2"
  node_count            = 1
  os_disk_size_gb       = 250
  os_type               = "Linux"
}

*/
