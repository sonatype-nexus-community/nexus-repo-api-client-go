module github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395

go 1.23

require gopkg.in/validator.v2 v2.0.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v3.94.0 // shipped with extensive generator-output breaking changes relative to v3.93.2 (renamed operationIds/methods, renamed model types, pointer-vs-value field changes) that were never intended for release; predates the /v395 import path
