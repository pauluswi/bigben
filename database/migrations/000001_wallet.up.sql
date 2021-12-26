create table account
(
    account_number  int           not null primary key,
    customer_number int           not null,
    balance         int default 0 not null
);
alter table account add constraint account_account_number_uindex unique (account_number);
alter table account add foreign key (customer_number) references customer (customer_number);

CREATE TABLE `wallet` (
  `accountid` int(11) NOT NULL,
  `balance` int(11) DEFAULT NULL,
  `modifieddate` datetime DEFAULT NULL,
  PRIMARY KEY (`accountid`),
  UNIQUE KEY `idwallet_UNIQUE` (`accountid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `wallet_trx` (
  `trxid` int(11) NOT NULL AUTO_INCREMENT,
  `accountid` int(11) NOT NULL,
  `trxtype` varchar(45) NOT NULL,
  `dc` varchar(45) NOT NULL,
  `trxamount` decimal(10,0) NOT NULL,
  `createdby` varchar(45) NOT NULL,
  `createddate` datetime NOT NULL,
  `modifiedby` varchar(45) NOT NULL,
  `modifieddate` datetime NOT NULL,
  PRIMARY KEY (`trxid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;