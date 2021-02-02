# GoLang API

A GoLang library to help to help facilitate RESTful API creation

[Installation](#installation)\
[Usage](#usage)

# Installation

RUN:

```
go get github.com/zeuce/golang-api
```

# Usage

- [Setup Server](#setup-server)
- [Setup Logging](#setup-logging)

### Setup Server

HTTP:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
)

func main() {
    //Start's a HTTP server on port 3000
    goapi.SetupServer(3000)
}
```

HTTPS:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
)

func main() {
    //Start's a HTTPS server on port 3000
    goapi.SetupServerSSL(3000,"/path/to/certfile", "/path/to/keyfile")
}
```

### Setup Logging

```
package main

import (
    goapi "github.com/zeuce/golang-api"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    //Setup logging support
    goapi.SetupLogging(r, "path/to/logDir", "path/to/logFile", "PREFIX")
}
```
