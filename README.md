# xmarvel

[![codecov](https://codecov.io/gh/ivantedja/xmarvel/branch/master/graph/badge.svg)](https://codecov.io/gh/ivantedja/xmarvel)

## Usage

Run via console:
```
$ PORT=1234 MARVEL_HOST=https://gateway.marvel.com MARVEL_PUBLIC_KEY=<your_public_key> MARVEL_PRIVATE_KEY=<your_private_key> go run cmd/api/main.go
$ curl localhost:1234/characters
```

## To do

- [ ] Add characters domain
  - [ ] Add redis repository
  - [ ] Add usecase
    - [ ] Bulk call Marvels' API
    - [ ] Set cache key for list of ids
    - [ ] Set cache key for single id
- [ ] Warm up cache
- [ ] Add datadog metrics
  - [ ] Counter Number of API call to Marvel
  - [ ] Middleware this service's endpoint
