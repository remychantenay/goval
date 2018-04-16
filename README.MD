[![Build Status](https://travis-ci.org/remychantenay/goval.svg?branch=master)](https://travis-ci.org/remychantenay/goval)
[![GoDoc](https://godoc.org/github.com/remychantenay/goval?status.svg)](https://godoc.org/github.com/remychantenay/goval)

# goval
Simple validation library for Go.

## Tags & Constraints
| Tag | Constraints | Additional Information |
| ------ | ------ | ------ |
| string | min,max,required | |
| uuid | required | |
| email | min,max,required,domain | |
| number | min,max,required | integer |
| enum | values,required | |
| country_code | exclude,excludeEu,required | [ISO-Alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) |
| currency | exclude,required | [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) |

## Installation
```bash
go get -u github.com/remychantenay/goval
```

## Usage
```go
import (
  goval github.com/remychantenay/goval
)

type SomeStruct struct {
	Id      	string		`goval:"uuid,required=true"`
	Country 	string		`goval:"country_code,required=true,exclude=US,excludeEu=true"`
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

## Documentation
https://godoc.org/github.com/remychantenay/goval

## License
Copyright (C) 2018 Chantenay Remy

Distributed under the MIT License.
