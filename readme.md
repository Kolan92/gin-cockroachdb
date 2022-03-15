# Products api

This demo app shows how to run simple rest api written in gin backed by cockroachdb.
Kubernetes and cockroachdb setup is based on this [article](https://github.com/cockroachdb/cockroach-operator).

## Local development

For local development cockroachdb is still need. Next section describes how to run kubernetes with cockrachdb.
To make database accesible for application running outside of cluster run:
`kubectl port-forward service/cockroachdb-public 26257`

To run gin app just run `go run .`in products-service directory
To run unit tests just run `go test`in products-service directory

## Kubernetes

1. Run cockroach setup [script](k8s/setup_cockroachDb.sh)
2. Run `kubectl apply -f ` for all files in k8s directory

## Azure + terraform

Kubernetes can be deployed to Azure with the terraform.
Terraform requires several variables, which can be conveniently stored as environment variables with `TF_VAR_` prefix:

```bash
export TF_VAR_subscription_id={Your subscription id}
export TF_VAR_tenant_id={Your tenant id}
export TF_VAR_serviceprinciple_id={Your service principal id, used to mange aks}
export TF_VAR_serviceprinciple_key={Your service principal password}
export TF_VAR_ssh_key={Public key path or key contents to install on node VMs for SSH access.}
```

Then you can deploy it using:

```bash
terraform init
terraform plan -out k8s.tfplan
terraform apply "k8s.tfplan"
```

## JSON Api examples

```bash
GET    /customer
    curl http://localhost:8080/customer

GET    /customer/:id
    curl http://localhost:8080/customer/1

POST   /customer
    curl -X POST -d '{"id": 1, "name": "bob"}' http://localhost:8080/customer

PUT    /customer/:id
    curl -X PUT -d '{"id": 2, "name": "robert"}' http://localhost:8080/customer/1

DELETE /customer
    curl -X DELETE http://localhost:8080/customer/1

GET    /product
    curl http://localhost:8080/product

GET    /product/:id
    curl http://localhost:8080/product/1

POST   /product
    curl -X POST -d '{"id": 1, "name": "apple", "price": 0.30}' http://localhost:8080/product

PUT    /product
DELETE /product

GET    /order
    curl http://localhost:8080/order

GET    /order/:id
    curl http://localhost:8080/order/1

POST   /order
    curl -X POST -d '{"id": 1, "subtotal": 18.2, "customer": {"id": 1}}' http://localhost:8080/order

PUT    /order
DELETE /order
```
