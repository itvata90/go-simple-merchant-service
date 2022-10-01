create database if not exists simplemerchantdb;

use simplemerchantdb;

create table if not exists merchants (
    code varchar(50) not null,
    contact_name varchar(100) not null,
    province varchar(100) not null,
    district varchar(100) not null,
    street varchar(100) not null,
    contact_email varchar(100) not null unique,
    contact_phone_no varchar(100) not null unique,
    owner_id varchar(50) not null,
    tax_id  varchar(50) not null,
    status varchar(50) not null,
    created_at datetime not null,
    created_by varchar(40),
    updated_at datetime,
    updated_by varchar(40),
    primary key (code)
);

create table if not exists team_members (
    id varchar(40) not null,
    username varchar(120) not null,
    password varchar(120) not null,
    first_name varchar(120) not null,
    last_name varchar(120) not null,
    birth_date date not null,
    nationality varchar(120) not null,
    contact_email varchar(100) not null unique,
    contact_phone_no varchar(100) not null unique,
    province varchar(120) not null,
    district varchar(120) not null,
    street varchar(120) not null,
    merchant_code varchar(40) not null,
    role varchar(40) not null,
    created_at datetime not null,
    created_by varchar(40),
    updated_at datetime,
    updated_by varchar(40),
    primary key(id),
    foreign key (merchant_code) references merchants(code)
);