# BookMarks

[![Docker Automated buil](https://img.shields.io/docker/automated/jrottenberg/ffmpeg.svg)](https://hub.docker.com/r/batazor/go-bookmarks/)

Simple service for saving link (tag, share and etc)

- [Docs](https://documenter.getpostman.com/view/95030/go-bookmarks/6n7Srf9)

### RUN

```
docker-compose build
docker-compose up
```

### ENV

| Name ENV         | Default value             |
|------------------|---------------------------|
| PORT             | 4070                      |
| MONGO_URL        | localhost/bookmarks       |

### technology stack

#### Back-End

* Go
* MongoDB

#### Library

+ [logrus](github.com/Sirupsen/logrus) - for logging
+ [chi](github.com/pressly/chi) - for routing
+ [glide](github.com/Masterminds/glide) - for vendoring
+ [prometheus](https://github.com/prometheus/client_golang) - for monitoring