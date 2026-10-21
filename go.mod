module github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397

go 1.23

require gopkg.in/validator.v2 v2.0.1

require (
	github.com/stretchr/testify v1.12.1 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)

retract v3.94.0 // shipped with extensive generator-output breaking changes relative to v3.93.2 (renamed operationIds/methods, renamed model types, pointer-vs-value field changes) that were never intended for release; predates the /v395 import path
retract v395.97.0 // generated from NXRM 3.97.0 and published on the /v395 import path, but NXRM 3.97.0 introduces generator-output breaking changes for Go; superseded by the /v397 import path (v397.97.1 onwards)
