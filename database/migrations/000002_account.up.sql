create table account
(
    account_number  int           not null primary key,
    customer_number int           not null,
    balance         int default 0 not null
);
alter table account add constraint account_account_number_uindex unique (account_number);
alter table account add foreign key (customer_number) references customer (customer_number);