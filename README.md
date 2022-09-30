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

Detail: [postman sample](data\testMerchants_CollectionAPI.postman_collection.json) 

## Team Members:
```
GET /team_members?pageSize=1&pageIdx=2
GET /team_members/{code}
GET /team_members/merchants/{merchant_code}
POST /team_members
PUT /team_members/{code}
DELETE /team_members/{code}
```

Detail: [postman sample](data\testTeamMembers_CollectionAPI.postman_collection.json)