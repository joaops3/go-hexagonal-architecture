```bash


docker exec -it appproduct bash

mockgen -destination=application/mocks/application.go -source=application/product.go application

cobra-cli init github.com/joaops3/go-hexagonal

cobra-cli add cli

go main.go cli -a=create -n="produto cli" -p=25.0

```
