# GoLang API

A GoLang library to help to help facilitate RESTful API creation

[Installation](#installation)\
[Usage](#usage)

## Installation

RUN:

```
go get github.com/zeuce/grapi
```

## Usage

- [Setup Server](#setup-server)
- [Setup Logging](#setup-logging)
- [Route Handling](#route-handling)
- [Adding Global Headers](#adding-global-headers)

### Setup Server

HTTP:

```go
package main

import (
    "github.com/zeuce/grapi"
)

func main() {
    //Start's a HTTP server on port 3000
    grapi.SetupServer(3000)
}
```

HTTPS:

```go
package main

import (
    "github.com/zeuce/grapi"
)

func main() {
    //Start's a HTTPS server on port 3000
    grapi.SetupServerSSL(3000,"/path/to/certfile", "/path/to/keyfile")
}
```

### Setup Logging

```go
package main

import (
    "github.com/zeuce/grapi"
)

func main() {
    //Setup logging support
    grapi.SetupLogging("path/to/logDir", "<LOGFILENAME>", "<PREFIX>")
}
```

### Route Handling

GET:

```go
package main

import (
    "github.com/zeuce/grapi"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := grapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup GET for /
    grapi.Get("/", handlerFunc)
}
```

POST:

```go
package main

import (
    "github.com/zeuce/grapi"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := grapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup POST for /
    grapi.Post("/", handlerFunc)
}
```

DELETE:

```go
package main

import (
    "github.com/zeuce/grapi"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := grapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup DELETE for /
    grapi.Delete("/", handlerFunc)
}
```

PATCH:

```go
package main

import (
    "github.com/zeuce/grapi"
    "encoding/json"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    d := grapi.ResponseStruct {
        StatusCode: 200,
        Message: "Yay its working"
    }
    json.NewEncoder(w).Encode(d)
}

func main() {
    //Setup PATCH for /
    grapi.Patch("/", handlerFunc)
}
```

### Adding Global Headers

```go
package main

import (
    "github.com/zeuce/grapi"
)

func main() {
    //Add global header
    grapi.AddDefaultHeader("key", "val")
    //Adding multiple global headers
    grapi.AddDefaultHeaders([]grapi.Header{
        grapi.Header {
            Key: "foo",
            Value: "bar",
        },
        grapi.Heder {
            Key: "foo1",
            Value: "bar1",
        },
    })
}
```
