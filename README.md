# Overview #

Merchant API service

A simple merchant API that can manage Merchant and related members


### Prerequisites

* Go 1.19

# Tech Stacks #
- [Go](https://go.dev/)
- [core-go: log, config, middleware, sql](https://github.com/core-go)
- hexagonal-architecture

# API #
## Heath check:
GET /heath

## Merchants:
```
GET /merchants?pageSize=1&pageIdx=2
GET /merchants/{code}
POST /merchants
PUT /merchants/{code}
DELETE /merchants/{code}
```

* `GET /merchants?pageSize=1&pageIdx=2` - Get merchants
    * Params:
        * `pageSize`: page size     
        * `pageIdx`: page index
    * Response
        * `data`: array of merchants
        * `pagination`: paging information

    * Example
    ```sh
    > curl -X GET <host>/merchants?pageSize=1&pageIdx=2
    ```

    ```json
    // Response
    HTTP 200 OK
    {
        "data": [
            {
                "code": "00000000-0000-0000-0000-00000_code01",
                "contactName": "contact Name",
                "province": "province",
                "district": "district",
                "street": "street",
                "contactEmail": "contact email",
                "contactPhoneNo": "contact phone",
                "ownerId": "00000000-0000-0000-0000-00000_owner01",
                "taxId": "123456780",
                "status": "Active",
                "createdAt": "2022-10-01T03:39:28+07:00"
            }
        ],
        "pagination": {
            "total": 1,
            "pageIndex": 2,
            "pageSize": 1
        }
    }
    ```

* `GET /merchants/{code}` - Get merchant by code
    * Params:
        * `code`: merchant code
    * Response
        * `object`: merchant detail

    * Example
    ```sh
    > curl -X GET <host>/merchants/00000000-0000-0000-0000-00000_code01
    ```

    ```json
     // Response
    HTTP 200 OK
    {
        "code": "00000000-0000-0000-0000-00000_code01",
        "contactName": "contact Name",
        "province": "province",
        "district": "district",
        "street": "street",
        "contactEmail": "contact email",
        "contactPhoneNo": "contact phone",
        "ownerId": "00000000-0000-0000-0000-00000_owner01",
        "taxId": "123456780",
        "status": "Active",
        "createdAt": "2022-10-01T03:39:28+07:00"
    }
    ```

* `POST /merchants` - Create merchant

    * Request body:
        * `object`: merchant detail
    * Response
        * `number`: affected rows

    * Example
     ```sh
    > curl -X POST <host>/merchants -d '{"contactName": "contact Name" ...}'
    ```

    ```json
    // Request body
    {
        "contactName": "contact Name",
        "province": "province",
        "district": "district",
        "street": "street",
        "contactEmail": "contact email",
        "contactPhoneNo": "contact phone",
        "ownerId": "00000000-0000-0000-0000-00000_owner01",
        "taxId": "123456780",
        "status": "Active"
    }

    // Response
    HTTP 200 OK
    1
    ```

* `PUT /merchants/{code}` - Update merchant
    * Params:
        * `code`: merchant code
    * Request body:
        * `object`: merchant detail
    * Response
        * `number`: affected rows

    * Example
    ```sh
    > curl -X PUT <host>/merchants/00000000-0000-0000-0000-00000_code01 -d '{"contactName": "contact Name" ...}'
    ```

    ```json
    // Request body
    {
        "contactName": "contact Name",
        "province": "province",
        "district": "district",
        "street": "street",
        "contactEmail": "contact email",
        "contactPhoneNo": "contact phone",
        "ownerId": "00000000-0000-0000-0000-00000_owner01",
        "taxId": "123456780",
        "status": "Active",
        "createdAt": "2022-10-01T03:39:28+07:00"
    }

    // Response
    HTTP 200 OK
    1
    ```

* `DELETE /merchants/{code}` - DELETE merchant
    * Params:
        * `code`: merchant code
    * Response
        * `number`: affected rows

    * Example
    ```sh
    > curl -X DELETE <host>/merchants/00000000-0000-0000-0000-00000_code01
    ```

    ```json
    // Response
    HTTP 200 OK
    1
    ```

* Detail: [postman sample](data/testMerchants_CollectionAPI.postman_collection.json) 

## Team Members:
```
GET /team_members?pageSize=1&pageIdx=2
GET /team_members/{code}
GET /team_members/merchants/{merchant_code}
POST /team_members
PUT /team_members/{code}
DELETE /team_members/{code}
```

* `GET /team_members?pageSize=1&pageIdx=2` - Get team members
    * Params:
        * `pageSize`: page size     
        * `pageIdx`: page index
    * Response
        * `data`: array of members
        * `pagination`: paging information

    * Example
    ```sh
    > curl -X GET <host>/team_members?pageSize=1&pageIdx=2
    ```

    ```json
    // Response
    HTTP 200 OK
    {
        "data": [
            {
                "id": "00000000-0000-0000-0000-00000_member01",
                "username": "member100",
                "password": "",
                "firstName": "Fname1",
                "lastName": "Lname1",
                "birthDate": "1991-05-06T00:00:00+07:00",
                "nationality": "nationality",
                "contactEmail": "",
                "contactPhoneNo": "",
                "province": "province",
                "district": "district",
                "street": "street",
                "merchantCode": "00000000-0000-0000-0000-00000_code01",
                "role": "Staff",
                "createdAt": "2022-10-01T03:47:35+07:00"
            }
        ],
        "pagination": {
            "total": 1,
            "pageIndex": 2,
            "pageSize": 1
        }
    }
    ```

* `GET /team_members/{code}` - Get member by code
    * Params:
        * `code`: member code
    * Response
        * `object`: merchant detail

    * Example
    ```sh
    > curl -X GET <host>/team_members/00000000-0000-0000-0000-00000_member01
    ```

    ```json
     // Response
    HTTP 200 OK
    {
        "id": "00000000-0000-0000-0000-00000_member01",
        "username": "member100",
        "password": "",
        "firstName": "Fname1",
        "lastName": "Lname1",
        "birthDate": "1991-05-06T00:00:00+07:00",
        "nationality": "nationality",
        "contactEmail": "",
        "contactPhoneNo": "",
        "province": "province",
        "district": "district",
        "street": "street",
        "merchantCode": "00000000-0000-0000-0000-00000_code01",
        "role": "Staff",
        "createdAt": "2022-10-01T03:47:35+07:00"
    }
    ```

* `GET /team_members/merchants/{merchant_code}?pageSize=1&pageIdx=2` - Get team members by merchant
    * Params:
        * `merchant_code`: merchant code
        * `pageSize`: page size     
        * `pageIdx`: page index
    * Response
        * `data`: array of members
        * `pagination`: paging information

    * Example
    ```sh
    > curl -X GET <host>/team_members/merchants/00000000-0000-0000-0000-00000_code01?pageSize=1&pageIdx=2
    ```

    ```json
    // Response
    HTTP 200 OK
    {
        "data": [
            {
                "id": "00000000-0000-0000-0000-00000_member01",
                "username": "member100",
                "password": "",
                "firstName": "Fname1",
                "lastName": "Lname1",
                "birthDate": "1991-05-06T00:00:00+07:00",
                "nationality": "nationality",
                "contactEmail": "",
                "contactPhoneNo": "",
                "province": "province",
                "district": "district",
                "street": "street",
                "merchantCode": "00000000-0000-0000-0000-00000_code01",
                "role": "Staff",
                "createdAt": "2022-10-01T03:47:35+07:00"
            }
        ],
        "pagination": {
            "total": 1,
            "pageIndex": 2,
            "pageSize": 1
        }
    }
    ```

* `POST /team_members` - Create member

    * Request body:
        * `object`: merchant detail
    * Response
        * `number`: affected rows

    * Example
    ```sh
    > curl -X POST <host>/team_members -d '{"username": "member100"...}'
    ```


    ```json
    // Request body
    {
        "username": "member100",
        "password": "password@6",
        "firstName": "first name",
        "lastName": "last name",
        "birthDate": "1991-05-06T00:00:00+07:00",
        "nationality": "nationality",
        "email": "test@gmail.com",
        "phone": "12345678",
        "province": "province",
        "district": "district",
        "street": "street",
        "merchantCode": "00000000-0000-0000-0000-00000_code01",
        "role": "Staff"
    }

    // Response
    HTTP 200 OK
    1
    ```

* `PUT /team_members/{code}` - Update member
    * Params:
        * `code`: member code
    * Request body:
        * `object`: member detail
    * Response
        * `number`: affected rows

    * Example
    ```sh
    > curl -X PUT <host>/team_members/00000000-0000-0000-0000-00000_member01 -d '{"username": "member100"...}'
    ```

    ```json
    // Request body
    {
        "username": "member100",
        "password": "password@6",
        "firstName": "first name",
        "lastName": "last name",
        "birthDate": "1991-05-06T00:00:00+07:00",
        "nationality": "nationality",
        "email": "test@gmail.com",
        "phone": "12345678",
        "province": "province",
        "district": "district",
        "street": "street",
        "merchantCode": "00000000-0000-0000-0000-00000_code01",
        "role": "Staff"
    }

    // Response
    HTTP 200 OK
    1
    ```

* `DELETE /team_members/{code}` - DELETE member
    * Params:
        * `code`: member code
    * Response
        * `number`: affected rows

    * Example
    ```sh
    > curl -X DELETE <host>/team_members/00000000-0000-0000-0000-00000_member01
    ```

    ```json
    // Response
    HTTP 200 OK
    1
    ```


* Detail: [postman sample](data/testTeamMembers_CollectionAPI.postman_collection.json)