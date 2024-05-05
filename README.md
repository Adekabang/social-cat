## Migrate Up
```sh
migrate -database "postgres://postgres:password@127.0.0.1:5432/social_cat?sslmode=disable" -path db/migrations up
```
## Migrate Drop
```sh
migrate -database "postgres://postgres:password@127.0.0.1:5432/social_cat?sslmode=disable" -path db/migrations drop -f
```

# Requirement:
- [x]  Authentication & Authorization
    - [x]  POST /v1/user/register
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
        - [x]  Error Handling
    - [x]  POST /v1/user/login
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
        - [x]  Error Handling
- [x]  Manage Cats
    - [x]  POST /v1/cat
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  GET /v1/cat
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  GET /v1/cat/{id}
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  DELETE /v1/cat/{id}
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
- [x]  Match Cat
    - [x]  POST /v1/cat/match
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  GET /v1/cat/match
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  POST /v1/cat/match/approve
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  POST /v1/cat/match/reject
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
    - [x]  DELETE /v1/cat/match/{id}
        - [x]  Request Format
        - [x]  Process
        - [x]  Response Format
