# gotapi
API SDK for crypto exchanges by Gotbit on Golang

## Installation

```bash
go get https://github.com/StrRoma/gotapi
```

## Usage

```golang
import (
    https://github.com/StrRoma/gotapi
)

const exchange = "exchange" // Name of exchnage (check exchnage list)
const apiKey = "apiKey" // Your Api Open key (depends on exchange, check exchnage list) 
const apiSecret = "apiSecret" // Your Api Secret key (depends on exchange, check exchnage list)
const accountID = "accountID" // Your Api additional information (depends on exchange, check exchnage list)

var apiClient = gotapi.Init(exchange, apiKey, apiSecret, accountID)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. You also can add new exchanges and after checking we will add it.

## License
[MIT](https://choosealicense.com/licenses/mit/)
