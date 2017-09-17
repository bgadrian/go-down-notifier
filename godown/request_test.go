package godown

import (
	"testing"
)

func runOne(s int, method, url string, errorShould bool, m string, t *testing.T) {
	alertError := HTTPRequest(s, "", method, url)

	if errorShould {
		if alertError == nil {
			t.Error(m)
		}
		return
	}

	//errorShould == false
	if alertError != nil {
		t.Error(m + alertError.Error())
	}
}

func TestWebReqBasic(t *testing.T) {
	runOne(1, "GET", "http://httpbin.org/status/200", false, "httpbin GET should not trigger erro", t)
	runOne(1, "POST", "http://httpbin.org/status/200", false, "httpbin POST should not trigger erro", t)
	runOne(1, "PUT", "http://httpbin.org/status/200", false, "httpbin PUT should not trigger erro", t)
	runOne(1, "DELETE", "http://httpbin.org/status/200", false, "httpbin DELETE should not trigger erro", t)

	runOne(1, "GET", "http://httpbin.org:80/status/200", false, "httpbin :80 200 should not trigger erro", t)
	runOne(1, "GET", "https://httpbin.org/status/200", false, "httpbin SSL 200 should not trigger erro", t)
	runOne(1, "GET", "http://httpbin.org/status/500", true, "httpbin 500 should trigger error", t)
	runOne(1, "GET", "http://httpbin.org/delay/2", true, "timeout requests should return error", t)
	runOne(2, "GET", "http://httpbin.org/delay/1", false, "delay should not return error", t)
	runOne(2, "'[]'", "http://httpbin.org/status/200", true, "invalid method should return error", t)
}

func TestWebReqFormats(t *testing.T) {
	runOne(2, "GET", "https://httpbin.org/image/png", false, "png misc type", t)
	runOne(2, "GET", "https://httpbin.org/image/jpeg", false, "jpeg misc type", t)
	runOne(2, "GET", "https://httpbin.org/image/webp", false, "webp misc type", t)
	runOne(2, "GET", "https://httpbin.org/image/svg", false, "svg misc type", t)
	runOne(2, "GET", "https://httpbin.org/forms/post", false, "post misc type", t)
	runOne(2, "GET", "https://httpbin.org/xml", false, "xml misc type", t)
	runOne(2, "GET", "https://httpbin.org/html", false, "html misc type", t)
	runOne(2, "GET", "https://httpbin.org/robots.txt", false, "robots misc type", t)
}

func TestWebReqRedirects(t *testing.T) {
	/*
		/redirect/:n 302 Redirects n times.
		/redirect-to?url=foo 302 Redirects to the foo URL.
		/redirect-to?url=foo&status_code=307 307 Redirects to the foo URL.
		/relative-redirect/:n 302 Relative redirects n times.
		/absolute-redirect/:n 302 Absolute redirects n times. \
	*/
	runOne(2, "GET", "https://httpbin.org/redirect/1", false, "redirect 1", t)
	runOne(1, "GET", "https://httpbin.org/redirect-to?url=https://httpbin.org", false, "redirect to", t)
	runOne(1, "GET", "https://httpbin.org/redirect-to?url=https://httpbin.org&status_code=308", false, "redirect to 308", t)

}
