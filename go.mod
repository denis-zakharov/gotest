module github.com/denis-zakharov/gotest

go 1.22.4

replace github.com/denis-zakharov/gotest/fooer => ./

require (
	github.com/denis-zakharov/gotest/fooer v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
