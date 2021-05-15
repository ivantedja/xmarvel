# xmarvel

[![codecov](https://codecov.io/gh/ivantedja/xmarvel/branch/master/graph/badge.svg)](https://codecov.io/gh/ivantedja/xmarvel)

## Documentation

Xmarvel provides service to proxy request to Marvel's API (see: https://developer.marvel.com/)

### Endpoints

- `/characters` - Return list of Marvel's "characterId"
- `/characters/{characterId}` - Return Marvel's character by `{characterId}`

### Flows

#### List Character IDs `/characters`

![List Character IDs](https://github.com/ivantedja/xmarvel/blob/master/docs/flows/xmarvels-List.png)

#### Show Character `/characters/{characterId}`

![Show Character](https://github.com/ivantedja/xmarvel/blob/master/docs/flows/xmarvels-Show.png)

### API Doc

See on [SwaggerHub](https://app.swaggerhub.com/apis/ivantedja/Xmarvel/1.0.0)

## Prerequisites

1. [Install Docker](https://docs.docker.com/engine/install/) to setup dependencies
2. [Install Golang](https://golang.org/dl/) > 1.15
3. Register to https://developer.marvel.com/ to get API credentials

## Running

1. Install dependencies
```
$ docker-compose up -d
```

2. Prepare environment variables (**adjust the credentials accordingly**):
```
$ cp env.sample .env
```

3. Generate binary:
```
$ make build
```

4. Run via console:
```
$ ./bin/main
```

OR alternatively you can directly run the application without generating the binary. Step no 3 can be directly changed into

3. Run via console:
```
$ make run
```



Test the endpoint:
```
$ curl localhost:8080/characters
```

## Cache Warming

Xmarvels contains characters' cache which rarely changed.
To keep the cache warm in the background, we can setup a cron to warm it up regularly.

Assuming you already have the binary from `make build` command
```
0 * * * * ./bin/warmer
```

## To do

- [ ] Add datadog metrics
  - [ ] Counter Number of API call to Marvel
  - [ ] Middleware this service's endpoint
