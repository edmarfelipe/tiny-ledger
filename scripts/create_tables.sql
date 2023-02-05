\connect ledger

create table if not exists public.persons (
   id        uuid not null primary key,
   name      varchar(100) not null,
   doc       char(11) not null unique,
   birth     date not null
);

create table if not exists public.accounts (
    id        uuid not null primary key,
    person_id uuid not null references persons,
    balance   decimal(12,2) not null,
    date      timestamp not null,
    enable    boolean  not null
);

create table if not exists public.transactions (
    id         uuid not null primary key,
    account_id uuid not null references accounts,
    amount     decimal(12,2) not null,
    date       timestamp not null
);
