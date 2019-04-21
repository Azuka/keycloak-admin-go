# keycloak-admin-go

[![](https://godoc.org/github.com/Azuka/keycloak-admin-go/pkg/keycloak?status.svg)](http://godoc.org/github.com/Azuka/keycloak-admin-go/pkg/keycloak)
[![pipeline status](https://gitlab.com/Azuka/keycloak-admin-go/badges/master/pipeline.svg)](https://gitlab.com/Azuka/keycloak-admin-go/commits/master)
[![coverage report](https://gitlab.com/Azuka/keycloak-admin-go/badges/master/coverage.svg)](https://gitlab.com/Azuka/keycloak-admin-go/commits/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azuka/keycloak-admin-go)](https://goreportcard.com/report/github.com/Azuka/keycloak-admin-go)
[![CircleCI](https://circleci.com/gh/Azuka/keycloak-admin-go.svg?style=svg)](https://circleci.com/gh/Azuka/keycloak-admin-go)

Keycloak admin client in go.

This is still highly unstable as more of the admin api endpoints and parameters are added.

## Usage
```shell
go get -u github.com/Azuka/keycloak-admin-go/...
```

## Local Development
```shell
make init
make test
make integration
#optionally
make integration-clean
```
### Local CI
- Install CircleCI locally: https://circleci.com/docs/2.0/local-cli

## Wish List
- [ ] Attack Detection
- [ ] Authentication Management
- [ ] Client Attribute Certificate
- [ ] Client Initial Access
- [ ] Client Registration Policy
- [ ] Client Role Mappings
- [ ] Client Scopes
- [ ] Clients
    - [x] Get clients
    - [x] Create clients
    - [ ] Update clients
    - [ ] Delete clients
- [ ] Component
- [ ] Groups (needs integration/unit tests)
    - [x] Create
    - [x] Delete
    - [x] Get
    - [x] AddRole
    - [x] DeleteRole
- [ ] Identity Providers
- [ ] Key
- [ ] Protocol Mappers
- [ ] Realms Admin
    - [x] Get realm
    - [x] Create realm
    - [ ] Import realm
    - [ ] Update realm
    - [x] Delete realm
    - [ ] Get admin events
    - [ ] Delete admin events
- [ ] Role Mapper
- [ ] Roles (needs integration/unit tests)
    - [x] Create
    - [x] Get
    - [x] GetMappings
- [ ] Roles (by ID)
- [ ] Scope Mappings
- [ ] User Storage Provider
- [ ] Users
    - [x] Get user
    - [x] Search users
    - [x] Create user
    - [x] Update user
        - [x] Profile information
        - [x] Groups
        - [x] Sessions, Consents
- [ ] Root

## Thanks to
- https://gopkg.in/resty.v1: quick and dirty REST client
- https://godoc.org/golang.org/x/oauth2: for the shamelessly copied authentication
- https://github.com/fatih/gomodifytags: because I'm too lazy to type json struct tags
