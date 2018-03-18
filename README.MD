# goval
TODO

## Tags & Constraints
| Tag | Constraints | Additional Information |
| ------ | ------ | ------ |
| string | min,max,required | |
| uuid | required | |
| email | min,max,required,domain | |
| number | min,max,required | integer |
| enum | values,required | |
| enum | values,required | |
| country_code | exclude,excludeEu,required | ISO-Alpha-2 |
| currency | exclude,required | ISO 4217 |

## Usage
### Install
```bash
go get https://github.com/remychantenay/goval
```
### Import
```golang
import (
  goval https://github.com/remychantenay/goval
)
```

## License
TODO
