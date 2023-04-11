# Go parrallel

```
pnpm install -g apib2swagger
apib2swagger -i data/lobstrio.apib --open-api-3 -o data/openapi.json
pnpm install -g @openapitools/openapi-generator-cli
openapi-generator-cli generate -g go -i data/openapi.json -o lobstrio-go-client --skip-validate-spec
```
