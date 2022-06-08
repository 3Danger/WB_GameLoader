https://habr.com/ru/post/254425/


CREATE TABLE account
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    login VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE customer (
      id SERIAL PRIMARY KEY NOT NULL UNIQUE,
      account_id INTEGER REFERENCES account (Id),
      money FLOAT NOT NULL
);

CREATE TABLE loader (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    account_id INTEGER REFERENCES account (Id) UNIQUE,
    customer_id INTEGER,
    money FLOAT NOT NULL,
    salary FLOAT NOT NULL,
    max_weight_trans FLOAT NOT NULL ,
    fatigue FLOAT NOT NULL,
    drunk BOOLEAN NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    account_id INTEGER REFERENCES account (Id),
    name VARCHAR NOT NULL,
    weight FLOAT NOT NULL
);


SELECT * FROM account INNER JOIN customer ON account.id = customer.users_id;

SELECT * FROM customer
    INNER JOIN account ON account.id = customer.users_id
    INNER JOIN loader ON customer.id = loader.customer_id;

INSERT INTO loader (account_id, customer_id, money, salary, maxweighttrans, fatigue, drunk)
VALUES (2, 1, 0, 20000, 20, 0.2, false);




SELECT * FROM account
WHERE data @> '{"money":4}';

... WHERE data ->> 'money' > '31';


CREATE TABLE tasks
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name varchar,
    weight float,
    user_id int
)


