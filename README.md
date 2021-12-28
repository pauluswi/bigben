# BigBen

BigBen is a microservice which provides simple wallet service for application users. 

## Description

A Rest API to access monetary accounts with the current balance of a user. 
The balance can be modified by registering transactions on the account, either debit transactions (removing funds) or credit transactions (adding funds).

A debit transaction will only succeed if there are sufficient funds on the account (balance - debit amount >= 0).

Users are be able to view their wallet balance, transaction history, make point to point transfer, deposit (top up) and withdrawal (cash out).

 
## Directory Structure  

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

