### sample swagger run

* [reference](https://rotational.io/blog/documenting-grpc-with-openapi/)

```
make build-sample-proto
docker run -p 80:8080 \
    -e SWAGGER_JSON=/protos/sample/sample.swagger.json \
    -v $PWD/protos/sample:/protos/sample \
    swaggerapi/swagger-ui
```