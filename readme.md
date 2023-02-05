# Tiny Ledger

A Tiny Ledger service

### API

| Description | Verb | Path |
| -- | -- |--|
| Create Person | POST  | /persons  |
| Create Account | POST | /accounts |
| Disable Account | PATCH | /accounts/{accountId} |
| Create Deposit | POST | /accounts/{accountId}/deposit |
| Get Balance | GET | /accounts/{accountId}/balance |
| Get Transactions | GET | /accounts/{accountId}/transactions |
| Get Transactions by Date | GET | /accounts/{accountId}/transactions?begin=2022-11-17&end=2022-12-31

### Prerequisites

Requirements for the software and other tools to build
- [Go](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/)

### Installing

To run on your local machine, just run

```shell
docker compose up -d
```

