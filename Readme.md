# keycloak-admin-go

[![](https://godoc.org/github.com/azuka/keycloak-admin-go?status.svg)](http://godoc.org/github.com/azuka/keycloak-admin-go)


Keycloak admin client in go.

This is still highly unstable as more of the admin api endpoints and parameters are added.

## Usage
```shell
go get github.com/azuka/keycloak-admin-go
```

## Local Development
```shell
make init
make test
make integration
```

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
- [ ] Role Mapper
- [ ] Roles
- [ ] Roles (by ID)
- [ ] Scope Mappings
- [ ] User Storage Provider
- [ ] Users
  - [x] Get user
  - [x] Search users
  - [ ] Create user
- [ ] Root

## Thanks to
- https://github.com/go-resty/resty: quick and dirty REST client
- https://github.com/mailru/easyjson: faster JSON serialization
- https://godoc.org/golang.org/x/oauth2: for the shamelessly copied authentication
- https://github.com/fatih/gomodifytags: because I'm too lazy to type json struct tags