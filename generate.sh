#!/bin/sh
set -eu

image='openapitools/openapi-generator-cli@sha256:5d76b35fe7d16b88771dd5871cc63d2102b40d9a854f2b375f8f622243585583'
output="$PWD/.generated"
trap 'rm -rf -- "$output"' EXIT

mkdir -p "$output"

docker run --rm \
  -v "$PWD:/work" \
  "$image" generate \
  -i /work/api/openapi.yaml \
  -g go \
  -o /work/.generated \
  --global-property 'apis,models,supportingFiles,apiTests=false,apiDocs=false,modelTests=false,modelDocs=false' \
  --additional-properties 'disallowAdditionalPropertiesIfNotPresent=false,enumClassPrefix=true,packageName=figma,packageVersion=0.42.0,useOneOfDiscriminatorLookup=true'

cp "$output"/*.go .
gofmt -w ./*.go
go mod tidy
