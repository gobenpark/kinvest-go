<p align="center">
<h1 align="center">kinvest-go</h1>
<p align="center">korea investment client library for Go (inspired by <a href="https://github.com/koreainvestment/open-trading-api">open-trading-api)</p>
</p>


## Installation

```bash
go get github.com/gobenpark/kinvest-go
```


## Usage

```go
cli := NewKinvest(&Config{
    AppKey:    AppKey,
    SecretKey: SecretKey,
    Imitation: false,
    Customer:  Person,
    Token:     Token,
    Account:   "",
})

// Do something
client... 
```
