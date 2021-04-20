# go-pokemorm

### How to run

#### Requirements
Glide installed
[Glide installed on your machine](https://glide.sh/)
[A pokemon DB](https://github.com/marty-crane/pokemORM)

#### Run
`glide install`

`cd go-pokemorm/cmd/main && go build && go run main.go`

#### Postman/CURL
URL:http://localhost:8080
##### Routes
###### /trainer `get`
Get all trainers
###### /trainer `post`
Add new trainer
###### /trainer/1 `get`
Get specific trainer
###### basic auth
user: cheeseman
password: cheese-cheese-cheese

#### todo
- GORM relationships
- Extract application config - DB credentials, auth keys etc
- Updates
- Deletes
- Routes for other tables/models
- Migrations so we don't rely on Laravel Pokemorm
