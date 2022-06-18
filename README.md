# Form3 challenge

## Authors
- Emiliano Abarca

## Description
The goal of this exercise is to provide a **Go SDK Client (Account)** to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository.

## Requirements
* Docker

## Run with Docker
* Up docker compose
```bash
docker-compose up -d --build
```
* Check logs and test result
```bash
docker logs form3-challenge-accountapi-sdk --follow
```

## How to use
* Fetch Account
```go
result, err := account.Fetch("401b45e0-ead2-41cd-b3e1-9122f3ee10b1")

if err != nil {
    log.Fatalf("error on Fetch: %v", err)
}

fmt.Printf("result from Fetch: %v\n", result)
```

* Create Account
```go
country := "GB"
accountClassification := "Personal"
version := int64(0)
attributes := account.Attributes{
AccountClassification: &accountClassification,
    BankID:                "TEST",
    BankIDCode:            "TEST",
    BaseCurrency:          "GBP",
    Bic:                   "NWBKGB42",
    Country:               &country,
    Name:                  []string{"Emiliano Abarca"},
}
data := account.Account{
    Attributes:     &attributes,
    ID:             uuid.New().String(),
    OrganisationID: uuid.New().String(),
    Type:           "accounts",
    Version:        &version,
}

result, err := account.Create(data)

if err != nil {
    log.Fatalf("error on Create: %v", err)
}

fmt.Printf("result from Create: %v\n", result)
```

* Delete Account
```go
result, err := account.Delete("401b45e0-ead2-41cd-b3e1-9122f3ee10b1", 0)

if err != nil {
    log.Fatalf("error on Delete: %v", err)
}

fmt.Printf("result from Delete: %v\n", result)
```

### Notes
- I'm Golang newbie, be nice to me

### TODO
* Migrate library to Go package
* Add [Viper](https://github.com/spf13/viper) or any other library to manage .env files
* Add linter
* Increase test coverage

## References
* [Challenge Instructions](https://github.com/form3tech-oss/interview-accountapi)
* [Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts)
* [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)