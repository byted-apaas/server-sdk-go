module github.com/byted-apaas/server-sdk-go

go 1.16

require (
	github.com/byted-apaas/server-common-go v0.0.39
	github.com/google/uuid v1.3.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/muesli/cache2go v0.0.0-20221011235721-518229cd8021
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/gjson v1.14.2
	go.mongodb.org/mongo-driver v1.10.1 // indirect
	golang.org/x/sys v0.0.0-20220817070843-5a390386f1f2 // indirect
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

// Deprecated
retract v0.0.26

// Deprecated
retract v0.0.27

// Deprecated
retract v0.0.29

// Deprecated
retract v0.0.30

// Deprecated
retract v0.0.31

// Deprecated
retract v0.0.32
