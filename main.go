package main

import (
    "net/http"
    "log"
)


func main() {
    var err error

    err = http.ListenAndServe(":9002", nil)
    if nil != err {
        log.Fatalf("[FATAL] %s", err)
    }
}
