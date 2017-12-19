# golang_http_request_header_check
GolangによるHTTPリクエストヘッダの確認


## GET

```sh
$ telnet localhost 8080
Trying ::1...
Connected to localhost.
Escape character is '^]'.
// リクエストヘッダ
GET / HTTP/1.1
Host: localhost

// レスポンスヘッダ
HTTP/1.1 200 OK
Date: Tue, 19 Dec 2017 13:28:25 GMT
Content-Length: 21
Content-Type: text/html; charset=utf-8

// レスポンスボディ
<b>Hello, World :</b>
```

```sh
$ telnet localhost 8080
Trying ::1...
Connected to localhost.
Escape character is '^]'.
// リクエストヘッダ
GET /?hoge=fuga HTTP/1.1
Host: localhost

// レスポンスヘッダ
HTTP/1.1 200 OK
Date: Tue, 19 Dec 2017 13:34:41 GMT
Content-Length: 25
Content-Type: text/html; charset=utf-8

// レスポンスボディ
<b>Hello, World :fuga</b>
```

## POST


```sh
$ telnet localhost 8080                                                                                                            [master]
Trying ::1...
Connected to localhost.
Escape character is '^]'.
// リクエストヘッダ
POST / HTTP/1.1
Host: localhost
Content-Type: application/x-www-form-urlencoded
Content-Length: 9

// リクエストボディ
hoge=fuga
// レスポンスヘッダ
HTTP/1.1 200 OK
Date: Tue, 19 Dec 2017 14:21:06 GMT
Content-Length: 25
Content-Type: text/html; charset=utf-8

// レスポンスボディ
<b>Hello, World :fuga</b>
```
