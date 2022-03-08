## How to run app

addr = flag.String("addr", "postgresql://roach:Q7gc8rEdS@localhost:26257/productsdb", "the address of the database")

## Create DB:

CREATE USER roach WITH PASSWORD 'Q7gc8rEdS';

## JSON Api examples

```
GET    /customer
    curl http://localhost:4100/customer

GET    /customer/:id
    curl http://localhost:4100/customer/1

POST   /customer
    curl -X POST -d '{"id": 1, "name": "bob"}' http://localhost:4100/customer

PUT    /customer/:id
    curl -X PUT -d '{"id": 2, "name": "robert"}' http://localhost:4100/customer/1

DELETE /customer
    curl -X DELETE http://localhost:4100/customer/1

GET    /product
    curl http://localhost:4100/product

GET    /product/:id
    curl http://localhost:4100/product/1

POST   /product
    curl -X POST -d '{"id": 1, "name": "apple", "price": 0.30}' http://localhost:4100/product

PUT    /product
DELETE /product

GET    /order
    curl http://localhost:4100/order

GET    /order/:id
    curl http://localhost:4100/order/1

POST   /order
    curl -X POST -d '{"id": 1, "subtotal": 18.2, "customer": {"id": 1}}' http://localhost:4100/order

PUT    /order
DELETE /order
```
