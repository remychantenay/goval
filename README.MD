[![Build Status](https://travis-ci.org/remychantenay/goval.svg?branch=master)](https://travis-ci.org/remychantenay/goval)
[![GoDoc](https://godoc.org/github.com/remychantenay/goval?status.svg)](https://godoc.org/github.com/remychantenay/goval)
[![Go Report Card](https://goreportcard.com/badge/github.com/remychantenay/goval)](https://goreportcard.com/report/github.com/remychantenay/goval)
[![codebeat badge](https://codebeat.co/badges/1cb648d7-39ea-4772-8330-03a2ed351e4b)](https://codebeat.co/projects/github-com-remychantenay-goval-master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# goval
Simple Struct validation library for Go. Uses reflection.

## Tags & Constraints
| Tag | Constraints | Additional Information |
| ------ | ------ | ------ |
| string | min,max,required | |
| uuid | required | |
| email | min,max,required,domain | |
| number | min,max,required | int64 |
| enum | values,required | |
| country_code | exclude,excludeEU,required | [ISO-3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) |
| currency | exclude,required | [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) |

## Usage
### Import
```bash
go get github.com/remychantenay/goval
```

### Example
```go
import "github.com/remychantenay/goval"

type SomeStruct struct {
	ID      	string		`goval:"uuid,required=true"`
	Country 	string		`goval:"country_code,required=true,exclude=US,excludeEU=true"`
	Currency 	string		`goval:"currency,required=true,exclude=EUR|GBP"`
	Email    	string		`goval:"email,required=true,domain=google.com"`
	Enum     	string		`goval:"enum,required=true,values=SOMETHING|SOMETHING_ELSE"`
	Number   	int		    `goval:"number,min=10,max=15,required=true"`
	String   	string		`goval:"string,min=6,max=255,required=true"`
}

myStruct := SomeStruct{...}

result := goval.ValidateStruct(myStruct)
for i, err := range result {
	fmt.Printf("\t%d. %s\n", i+1, err.Error())
}
```

## Transitive Dependencies
```
github.com/satori/go.uuid
```

## Documentation
https://godoc.org/github.com/remychantenay/goval

## License
Apache License Version 2.0