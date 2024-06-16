docker exec -it appproduct bash

mockgen -destination=application/mocks/application.go -source=application/product.go application
