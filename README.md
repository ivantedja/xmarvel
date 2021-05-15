# xmarvel

[![codecov](https://codecov.io/gh/ivantedja/xmarvel/branch/master/graph/badge.svg)](https://codecov.io/gh/ivantedja/xmarvel)

## Prerequisites

1. [Install Docker](https://docs.docker.com/engine/install/) to setup dependencies
2. [Install Golang](https://golang.org/dl/) > 1.15
3. Register to https://developer.marvel.com/ to get API credentials

## API Documentation

See https://github.com/ivantedja/xmarvel/tree/master/swagger/index.html

## Running

1. Install dependencies
```
$ docker-compose up -d
```

2. Prepare environment variables (**adjust the credentials accordingly**):
```
$ cp env.sample .env
```

3. Run via console:
```
$ make run
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
