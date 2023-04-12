# Go parrallel

```bash
pnpm install -g @openapitools/openapi-generator-cli
```

```bash
openapi-generator-cli generate -g go \
 -i api-schemas/lobstrio-openapi-3.json -o api/lobstrio \
 --additional-properties packageName=lobstrio \
 --additional-properties packageVersion=latest \
 --git-repo-id go-parrallel-train/api/lobstrio --git-user-id esgi-lyon

rm -rf api/lobstrio/go.mod api/lobstrio/go.sum
```

## Run

```bash
export LOBSTRIOS_API_TOKEN=your_token
go run main.go
```
