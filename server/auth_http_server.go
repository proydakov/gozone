package main

import (
    "log"
    "net/http"
)

func checkAuth(username string, password string, ok bool) bool {
    return ok
}

func handler(w http.ResponseWriter, r *http.Request) {
    username, password, ok := r.BasicAuth()
    log.Printf("Username: '%s', Password: '%s', Ok: %t", username, password, ok)
    if !checkAuth(username, password, ok) {
        w.Header().Set("WWW-Authenticate", "Basic realm=\"Zzz\"")
        http.Error(w, "Authorization failed", 401)
        return
    }

    names := r.Header["Name-File"]
    name := ""
    if len(names) == 1 {
        name = names[0]
    }
    log.Printf("NAME-FILE: %s", name)
    if "" == name {
        http.Error(w, "Not found header 'NAME-FILE'", 400)
        return
    }
    w.Write([]byte("ok"))
}

func main() {
    http.HandleFunc("/upload", handler)
    http.ListenAndServe(":11111", nil)
}
