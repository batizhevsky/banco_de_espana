Create Table if not exists clients(
    id integer primary key asc,
    name text,
    email text,
    phone integer
);

Create Table if not exists accounts(
    id integer primary key asc,
    client_id integer,
    balance integer
);

Create Table if not exists transactions(
    id integer primary key asc,
    account_id integer,
    amount integer,
    subject text,
    timestamp integer 
);