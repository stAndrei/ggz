# ggz

[![GoDoc](https://godoc.org/github.com/go-ggz/ggz?status.svg)](https://godoc.org/github.com/go-ggz/ggz)
[![Build Status](http://drone.wu-boy.com/api/badges/go-ggz/ggz/status.svg)](http://drone.wu-boy.com/go-ggz/ggz)
[![codecov](https://codecov.io/gh/go-ggz/ggz/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ggz/ggz)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ggz/ggz)](https://goreportcard.com/report/github.com/go-ggz/ggz)
[![codebeat badge](https://codebeat.co/badges/0a4eff2d-c9ac-46ed-8fd7-b59942983390)](https://codebeat.co/projects/github-com-appleboy-gorush)
[![Docker Pulls](https://img.shields.io/docker/pulls/go-ggz/ggz.svg)](https://hub.docker.com/r/go-ggz/ggz/)
[![](https://images.microbadger.com/badges/image/go-ggz/ggz.svg)](https://microbadger.com/images/go-ggz/ggz "Get your own image badge on microbadger.com")
[![Release](https://github-release-version.herokuapp.com/github/go-ggz/ggz/release.svg?style=flat)](https://github.com/go-ggz/ggz/releases/latest)

An URL shortener service written in Golang

## Prepare

copy `.env.example` to `.env`

```sh
$ cp .env.example `.env`
```

open `.env` file 

[embedmd]:# (.env.example ini)
```ini
GGZ_DB_DRIVER=mysql
GGZ_DB_USERNAME=root
GGZ_DB_PASSWORD=123456
GGZ_DB_NAME=ggz
GGZ_DB_HOST=127.0.0.1:3307
GGZ_SERVER_ADDR=:8080
GGZ_SHORTEN_SERVER_ADDR=:8081
GGZ_DEBUG=true
GGZ_SERVER_HOST=http://localhost:8080
GGZ_SERVER_SHORTEN_HOST=http://localhost:8081
```
