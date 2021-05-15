# xmarvel

[![codecov](https://codecov.io/gh/ivantedja/xmarvel/branch/master/graph/badge.svg)](https://codecov.io/gh/ivantedja/xmarvel)

## Usage

1. Register to https://developer.marvel.com/ to get API credentials

2. Prepare environment variables (**adjust the credentials accordingly**):
```
$ cp env.sample .env
```

3. Run via console:
```
$ go run cmd/api/main.go
```

4. Test the endpoint:
```
$ curl localhost:8080/characters
```

## To do

- [ ] Warm up cache
- [ ] Add datadog metrics
  - [ ] Counter Number of API call to Marvel
  - [ ] Middleware this service's endpoint
