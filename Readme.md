# keycloak-admin-go

[![](https://godoc.org/github.com/Azuka/keycloak-admin-go/keycloak?status.svg)](http://godoc.org/github.com/Azuka/keycloak-admin-go/keycloak)
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
```
### Local CI
- Install CircleCI locally: https://circleci.com/docs/2.0/local-cli

## Wish List
- [x] Add authentication integration tests
- [ ] Attack Detection
- [ ] Authentication Management
- [ ] Client Attribute Certificate
- [ ] Client Initial Access
- [ ] Client Registration Policy
- [ ] Client Role Mappings
- [ ] Client Scopes
- [ ] Clients
- [ ] Component
- [ ] Groups
- [ ] Identity Providers
- [ ] Key
- [ ] Protocol Mappers
- [ ] Realms Admin
  - [x] Get realm
  - [ ] Import realm
  - [ ] Update realm
  - [x] Delete realm
  - [ ] Get admin events
  - [ ] Delete admin events
- [ ] Role Mapper
- [ ] Roles
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
- https://github.com/mailru/easyjson: faster JSON serialization
- https://godoc.org/golang.org/x/oauth2: for the shamelessly copied authentication
- https://github.com/fatih/gomodifytags: because I'm too lazy to type json struct tags
