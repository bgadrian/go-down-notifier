# Go Down Notifier
A small service checker built in Go. Every X seconds checks a list of URL's for a valid response (2xx or 3xx) status code.

If the response fails an alert email is sent.

## Usage

```bash
go install github.com/bgadrian/go-down-notifier
go build github.com/bgadrian/go-down-notifier
./go-down-notifier -interval=3 -debug=true -web=https://url1,https://url2
```

## Parameters
```bash 
  -debug
        more verbose, even when success
  -gpass string
        gmail password
  -guser string
        gmail username
  -interval uint
        seconds interval to make the check (default 300)
  -mail string
        email to receive alerts
  -method string
        web request method (default "GET")
  -uagent string
        user agent for web requests, default Chrome 60 (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
  -web string
        a CSV of URL's for web request checks (default "http://httpbin.org/get?a=test")

```

### GMAIL
If provided, it uses a gmail account to send the failure. Email example:

```txt
Subject: Alert! [https://provided.url]
Body: Alert!

 A service didn't respond.
 target: https://provided.url
 error: Get https://provided.url: dial tcp: lookup sss: no such host
 ```

 #### TODO
 * recheck N times before alerting
 * implement a PORT checker too, for other services like databases
 * add more alert services like pager duty, web hooks, etc.

 #### License MIT
 Copyright (c) 2017 B.G.Adrian.