-- Active: 1679021533728@@127.0.0.1@5432@product_sales
SELECT
    *
FROM
    roles;

ALTER TABLE
    roles
add
    constraint roles_pkey PRIMARY KEY (id);

ALTER Table
    users
ADD
    CONSTRAINT fk_roles_id Foreign Key (role_id) REFERENCES roles (id);