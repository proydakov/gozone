package main

import (
    "net/http"
)

var PAGE = `
<!DOCTYPE html>
<html>
<head>
<title>Welcome to Golang!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to Golang!</h1>
<p>If you see this page, the poco web server is successfully compiled and
working.</p>

<p>For online documentation please refer to
<a href="https://github.com/proydakov/gozone">gozone</a>.<br/>

<p><em>Thank you for using Golang.</em></p>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(PAGE))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":11111", nil)
}
