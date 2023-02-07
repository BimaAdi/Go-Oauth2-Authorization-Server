# OAuth2 Authorization Server in Golang
ON PROGRESS

## TODO
- Database / ORM
    - Gorm Integration [v]
- CRUD User
    - Create table user [v]
    - Get All User (search and pagination) [v]
    - Get All User Test [v]
    - Get Detail User [v]
    - Get Detail User Test [v]
    - Create User [v] (don't forget to json validation)
    - Create User Test [v]
    - Update User [v] (don't forget to json validation)
    - Update User Test [v]
    - Delete User [v] (soft delete and logic if user already exists)
    - Delete User Test [v]
    - Authorization For CRUD User []
- Authentication
    - Simple email password login /token []
    - JWT token []
- Testing
    - testify integration [v]
- Swagger
    - swagger integration []
- CLI using https://github.com/urfave/cli
    - CLI integration []
    - create superuser through cli []
    - run server througj cli []
- Oauth2 Flow
    - create table oauth2_session []
    - client registration api /client-registration (generate client_id and client_secret) []
    - Oauth2 redirect ui
        - login ui []
        - forgot password ui []
    - exchange authorization code api /token (sharing with login user) []
    ```
    POST https://authorization-server.com/token 

    request:
    grant_type=authorization_code
    &client_id=pCQhYDc-nX02fzza_R9rBJ6J
    &client_secret=0eevNqvAB71F_lLVO_GgBLkwxPPWpsC7CDRW6n6DhhToSh7a
    &redirect_uri=https://www.oauth.com/playground/authorization-code.html
    &code=7ovthyS5JuUryY1q6gbG6pfrQsAcFjxNomumPUduc626FLSK

    response:
    {
        "token_type": "Bearer",
        "expires_in": 86400,
        "access_token": "SuisKC3OxrJBmHrQHKmriQIkeibbazCy8QIkf-AlLAIuDzwSyQF6UiQMj_yi0E0KVZpBYYMC",
        "scope": "photo offline_access",
        "refresh_token": "4u_WctybF_VLtSEpCa3igJGf"
    }

    ```
- CRUD Role
    - Create table role []
    - Get All Role (search and pagination) []
    - Get All Role Test []
    - Get Detail Role []
    - Get Detauk Role Test []
    - Create Role []
    - Create Role Test []
    - Update Role []
    - Update Role Test []
    - Delete Role [] (soft delete and logic if role already exists)
    - Delete Role Test []
- CRUD Module
    - Create table module []
    - Get All Module (search and pagination) []
    - Get All Module Test []
    - Get Detail Module []
    - Get Detauk Module Test []
    - Create Module []
    - Create Module Test []
    - Update Module []
    - Update Module Test []
    - Delete Module [] (soft delete and logic if module already exists)
    - Delete Module Test []
- CRUD Action
    - Create table action []
    - Get All Action (search and pagination) []
    - Get All Action Test []
    - Get Detail Action []
    - Get Detauk Action Test []
    - Create Action []
    - Create Action Test []
    - Update Action []
    - Update Action Test []
    - Delete Action [] (soft delete and logic if action already exists)
    - Delete Action Test []
- CRUD Permission
    - Create table permission []
    - Get All Permission (search and pagination) []
    - Get All Permission Test []
    - Get Detail Permission []
    - Get Detauk Permission Test []
    - Create Permission []
    - Create Permission Test []
    - Update Permission []
    - Update Permission Test []
    - Delete Permission [] (soft delete and logic if permission already exists)
    - Delete Permission Test []
- CRUD Many to Many
    - Create Tabel user_role []
    - Create Tabel user_permission []
    - Create Tabel role_permission []

## Testing

- run all testing `go test ./...`
- run all test in folder `go test ./{folder name}/... ./{another folder name}/...`
- run all test in file `go test ./{folder name}/{file name}`
- run specific test function `go test --run '^{function name}$' ./{folder name}/{file name}` (Note: --run input is a regex)
- run test verbosely (show log) `go test ./... -v`
- remove all test cache `go clean -testcache ./...`
