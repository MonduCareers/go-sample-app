# go-sample-app

go-sample-app is a simple bank API that implements the following endpoints:
1. Create account
2. Get account balance
3. Deposit money to the account

## Run locally
### Start postgres container
```bash
docker-compose up -V
```
The above command will start `postgres` container and create DB schema from `./db/schema/*.sql` directory

To verify that schema is correctly created, run:
```bash
docker exec -it postgres psql -d mondu_dev
\d
```

### Start application
```bash
go run main.go
```

### Testing
Note: postgres and the application must be running.
```bash
go test -v ./...
```

## API

### Create account
Request:
```bash
curl -XPOST http://localhost:8080/account/create
```
Sample response:
```
# 201 Created
{"id":"f301561b-02a7-4d96-8812-34ca4cbb2a91"}
```

### Get account balance
Request:
```bash
curl http://localhost:8080/account/balance?id=<Account ID>
```
Sample response:
```
# 200 OK
{"balance":0}
```

### Deposit
Request:
```bash
curl -XPOST http://localhost:8080/account/deposit -d '{"id":"<Account ID>","amount":5}'
```

Sample response:
```
# 204 No Content
```
