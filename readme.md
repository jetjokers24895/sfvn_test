#HOW TO RUN

```
docker-compose up -d

go run ./cmd/main.go
```
# About Coingecko
when user request come
1. Read from redis, if had then return, otherwise next step
2. Read from postgres, if had then return, otherwise next step
3. Read from API. Then cache to redis and save to postgres for later