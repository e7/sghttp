package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
)

var (
    flag_help bool
    flag_basedir string
    flag_host string
    flag_port string
)


func main() {
    var err error

    // flags
    flag.BoolVar(&flag_help, "h", false, "show this message")
    flag.Usage = func () {
        fmt.Fprintf(os.Stderr, "usage: xxxxx\n")
    }
    flag.StringVar(&flag_basedir, "b", ".", "base directory")
    flag.Parse()

    if flag_help {
        flag.Usage()
        return
    }

    err = http.ListenAndServe("127.0.0.1:9002", nil)
    if nil != err {
        log.Fatalf("[FATAL] %s", err)
    }
}
