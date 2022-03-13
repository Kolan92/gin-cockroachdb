#! /bin/bash

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.5.1/install/crds.yaml --kubeconfig kubeconfig
echo 'Finished crds.yaml'

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.5.1/install/operator.yaml --kubeconfig kubeconfig
echo 'Finished operator.yaml'

kubectl apply -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/v2.5.1/examples/example.yaml --kubeconfig kubeconfig
echo 'Finished example.yaml'

kubectl create -f https://raw.githubusercontent.com/cockroachdb/cockroach-operator/master/examples/client-secure-operator.yaml --kubeconfig kubeconfig
echo 'Finished client-secure-operator.yaml'

#TODO store passowrds in secure way...
kubectl exec -it cockroachdb-client-secure \
--kubeconfig kubeconfig \
-- ./cockroach sql \
--certs-dir=/cockroach/cockroach-certs \
--host=cockroachdb-public \
--execute="CREATE DATABASE IF NOT EXISTS productsdb; CREATE USER IF NOT EXISTS roach WITH PASSWORD 'Q7gc8rEdS'; GRANT ALL TO roach;" 

echo 'Finished CREATE DATABASE and CREATE USER'

kubectl apply -f ../k8s/cockroachdb.yaml --kubeconfig kubeconfig
echo 'Finished cockroachdb.yaml'
