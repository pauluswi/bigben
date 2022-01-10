# BigBen

Is a microservice which provides simple wallet service for application users. 

[![Build](https://github.com/pauluswi/bigben/actions/workflows/build.yml/badge.svg)](https://github.com/pauluswi/bigben/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/pauluswi/bigben/branch/master/graph/badge.svg)](https://codecov.io/gh/pauluswi/bigben)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)


## Description

A Rest API to access monetary accounts with the current balance of a user. 
The balance can be modified by registering transactions on the account, either debit transactions (removing funds) or credit transactions (adding funds).

A debit transaction will only succeed if there are sufficient funds on the account (balance - debit amount >= 0).

Users are be able to view their wallet balance, transaction history, make point to point transfer, deposit (top up) and withdrawal (cash out).

For Authentication, api key is used at header as a simple way to secure access.

 
## Directory Structure  
```
├── cmd                        // app contains main execution file 
├── configs                    // app configuration files 
├── controller                 // is a directory consist of routing process
├── database                   // is a directory consist of db/data processing
│   └── migration              // is a directory consist of migration files
│   └── seeder                 // is a directory consist of seeding data process
├── entity                     // is a directory consist of struct models used in codebase
├── exception                  // is a directory consist of exception mechanism
├── middleware                 // is a directory consist of app validation layer before hit the main functionalities 
├── model                      // is a directory consist of struct used roe request and response
├── repository                 // is a domain's repository acting to store the data
├── service                    // is a directory consist of functional interface
├── validation                 // is a directory consist of functional validation
```
# Endpoints

## Healthcheck

| Endpoint | Method | Description |
| -------- | ------ | ----------- |
| /ping    | GET    | Check for the service application healthiness |

## Wallet

| Endpoint                                      | Method | Description |
| --------------------------------------------- | ------ | ----------- |
| /v1/ewallet/balance/:account_id               | GET    | Retrieve User's Wallet Balance  |
| /v1/ewallet/transaction/history/:account_id   | GET    | Retrieve User's Transaction History |
| /v1/ewallet/transaction/transfer              | POST   | Make a Point to Point Transfer |
| /v1/ewallet/transaction/deposit               | POST   | Make a User's Wallet Deposit |
| /v1/ewallet/transaction/withdrawal            | POST   | Make a User's Wallet Withdrawal |

# Curl

  ```sh
  curl -X GET \
  http://localhost:3000/ping \
  -H 'x-api-key: 12345' 
  ```
  ```sh
  curl -X GET \
  http://localhost:3000/v1/ewallet/balance/10001 \
  -H 'x-api-key: 12345' 
  ```
  ```sh
  curl -X GET \
  http://localhost:3000/v1/ewallet/transaction/history/10001 \
  -H 'x-api-key: 12345' 
  ```
  ```sh
  curl -X POST \
  http://localhost:3000/v1/ewallet/transaction/transfer \
  -H 'x-api-key: 12345' \
  -H 'Content-Type: application/json' \
  -d '{"from_account_number":10001,"to_account_number":10002,"amount":8}'
  ```
  ```sh
  curl -X POST \
  http://localhost:3000/v1/ewallet/transaction/deposit \
  -H 'x-api-key: 12345' \
  -H 'Content-Type: application/json' \
  -d '{"to_account_number":10001,"amount":8}'
  ```
 ```sh
  curl -X POST \
  http://localhost:3000/v1/ewallet/transaction/withdrawal \
  -H 'x-api-key: 12345' \
  -H 'Content-Type: application/json' \
  -d '{"from_account_number":10001,"amount":8}'
  ```


# Getting started

## Run service dependencies

```sh
make infra-up
```

## Run migration
Migrates the database to the most recent version available.
```
make migrate-up
```

Undo 1 step database migration.
```
make migrate-down
```

## Run service application

```sh
make serve
```

# Technology Used

- Golang (with Fiber Web Framework)
- MySQL

# Unit Testing

For testability purpose, unit testings are provided.
We can use golang test package.

```sh
$ go test -v controller/*.go -race -coverprofile=coverage.out -covermode=atomic
=== RUN   TestEWalletController_GetBalance
--- PASS: TestEWalletController_GetBalance (0.02s)
=== RUN   TestEWalletController_GetTransactionHistory
--- PASS: TestEWalletController_GetTransactionHistory (0.01s)
=== RUN   TestEWalletController_Transfer
--- PASS: TestEWalletController_Transfer (0.02s)
=== RUN   TestEWalletController_Deposit
--- PASS: TestEWalletController_Deposit (0.01s)
=== RUN   TestEWalletController_Withdrawal
--- PASS: TestEWalletController_Withdrawal (0.02s)
=== RUN   TestEWalletController_HealthCheck
--- PASS: TestEWalletController_HealthCheck (0.00s)
PASS
coverage: 81.1% of statements
ok      command-line-arguments  1.833s  coverage: 81.1% of statements

```
