create database if not exists simplemerchantdb;

use simplemerchantdb;

create table if not exists merchants (
    code varchar(40) not null,
    name varchar(120) not null,
    province varchar(120) not null,
    country varchar(120) not null,
    address varchar(120) not null,
    email varchar(120) not null,
    phone varchar(120) not null,
    status varchar(40) not null,
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
    email varchar(120) not null unique,
    phone varchar(120) not null unique,
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

insert into merchants (code, name, status, province, district, street, email, phone, owner_id, created_at) values('734cb079-1a2e-11ed-b4c1-7c10c91fb7f4', 'First Merchant', 'Active', "province1", "district1", "street1", "email1", "0123456781", "748a5586-1bbb-11ed-861d-0242ac120002", "2006-01-02 03:04:05");
insert into merchants (code, name, status, province, district, street, email, phone, owner_id, created_at) values('9bf4271c-fe0b-419d-95fd-5592b7a527dd', 'Second Merchant', 'Active', "province1", "district1", "street1", "email2", "0123456782", "748a5586-1bbb-11ed-861d-0242ac120002","2006-01-02 03:04:06");
insert into merchants (code, name, status, province, district, street, email, phone, owner_id, created_at) values('07e7a76c-1bbb-11ed-861d-0242ac120002', 'Third Merchant', 'Active', "province1", "district1", "street1", "email3", "0123456783", "748a5586-1bbb-11ed-861d-0242ac120002", "2006-01-02 03:04:07");

insert into team_members (id, username, password, first_name, last_name, birth_date, nationality, email, phone, province, district, street, merchant_code, role, created_at) values ("6a077d3c-1bbb-11ed-861d-0242ac120002", "member1", "password@1", "Fname1","Lname1","1991-05-06", "nationality1", "email1@gmail.com", "0987654321", "province1", "district1", "street1", '734cb079-1a2e-11ed-b4c1-7c10c91fb7f4', "Staff","2006-01-02 03:04:05");
insert into team_members (id, username, password, first_name, last_name, birth_date, nationality, email, phone, province, district, street, merchant_code, role, created_at) values("6dc0bb0a-1bbb-11ed-861d-0242ac120002",  "member2", "password@2", "Fname2","Lname2","1991-05-06", "nationality2", "email2@gmail.com", "0987654322", "province1", "district1", "street1", '9bf4271c-fe0b-419d-95fd-5592b7a527dd', "Staff", "2006-01-02 03:04:05");
insert into team_members (id, username, password, first_name, last_name, birth_date, nationality, email, phone, province, district, street, merchant_code, role, created_at) values("70c55c98-1bbb-11ed-861d-0242ac120002",  "member3", "password@3", "Fname3","Lname3","1991-05-06", "nationality3", "email3@gmail.com", "0987654323", "province1", "district1", "street1", '9bf4271c-fe0b-419d-95fd-5592b7a527dd', "Manager", "2006-01-02 03:04:05");
insert into team_members (id, username, password, first_name, last_name, birth_date, nationality, email, phone, province, district, street, merchant_code, role, created_at) values("748a5586-1bbb-11ed-861d-0242ac120002",  "member4", "password@4", "Fname4","Lname4","1991-05-06", "nationality4", "email4@gmail.com", "0987654324", "province1", "district1", "street1", '734cb079-1a2e-11ed-b4c1-7c10c91fb7f4', "Owner", "2006-01-02 03:04:05");
insert into team_members (id, username, password, first_name, last_name, birth_date, nationality, email, phone, province, district, street, merchant_code, role, created_at) values("772effbc-1bbb-11ed-861d-0242ac120002",  "member5", "password@5", "Fname5","Lname5","1991-05-06", "nationality5", "email5@gmail.com", "0987654325", "province1", "district1", "street1", '07e7a76c-1bbb-11ed-861d-0242ac120002', "Manager", "2006-01-02 03:04:05");
