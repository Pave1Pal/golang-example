create table product
(
    id    uuid not null
        primary key,
    price bigint,
    name  varchar
);

create table cart
(
    id    uuid   primary key,
    price bigint not null
);

create table purchase
(
    id      uuid                     not null
        primary key,
    person  varchar(255)             not null,
    address varchar(255)             not null,
    date    timestamp with time zone not null,
    cart_id uuid                     not null
        constraint fk_purchase_cart
            references cart (id)
);

create table cart_product_merge
(
    cart_id    uuid not null
        constraint fk_cart_product_cart references cart (id),

    product_id uuid not null
        constraint fk_cart_product_product references product (id)
);

insert into product (id, price, name)
values (
        gen_random_uuid(),
        40000,
        'Мягкий диван'
       );

insert into product (id, price, name)
values (
        gen_random_uuid(),
        50000,
        'Мягкое кресло'
       );

insert into product (id, price, name)
values (
        gen_random_uuid(),
        35000,
        'Мягкий стол'
       );