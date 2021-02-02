# GoLang API

A GoLang library to help to help facilitate RESTful API creation

[Installation](#installation)\
[Usage](#usage)

## Installation

RUN:

```
go get github.com/zeuce/golang-api
```

## Usage

- [Setup Server](#setup-server)
- [Setup Logging](#setup-logging)
- [Route Handling](#route-handling)
- [Adding Global Headers](#adding-global-headers)

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
)

func main() {
    //Setup logging support
    goapi.SetupLogging("path/to/logDir", "<LOGFILENAME>", "<PREFIX>")
}
```

### Route Handling

GET:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := goapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup GET for /
    goapi.Get("/", handlerFunc)
}
```

POST:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := goapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup POST for /
    goapi.Post("/", handlerFunc)
}
```

DELETE:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := goapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup DELETE for /
    goapi.Delete("/", handlerFunc)
}
```

PATCH:

```
package main

import (
    goapi "github.com/zeuce/golang-api"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := goapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup PATCH for /
    goapi.Patch("/", handlerFunc)
}
```

### Adding Global Headers

```
package main

import (
    goapi "github.com/zeuce/golang-api"
)

func main() {
    //Add global header
    goapi.AddDefaultHeader("key", "val")
    //Adding multiple global headers
    goapi.AddDefaultHeaders([]goapi.Header{
        goapi.Header {
            Key: "foo",
            Value: "bar",
        },
        goapi.Heder {
            Key: "foo1",
            Value: "bar1",
        },
    })
}
```
