package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "/u02713"
	ballotX   = "/u2717"
)

var feed = `<?xml version="1.0" encoding="UTF-8"?>
	<rss>
	<channel>
		<title>Going Go Programming</title>
		<description>Golang : https://github.com/goinggo</description>
		<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
			<title>Object Oriented Programming Mechanics</title>
			<description>Go is an object oriented language.</description>
			<link>http://www.goinggo.net/2015/03/object-oriented</link>
		</item>
	</channel>
	</rss>`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {

	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading content.")
	{
		for _, url := range urls {

			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url.url, url.statusCode)
			{
				resp, err := http.Get(url.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the GET call", ballotX, err)
				}
				t.Log("\t\tShould be able to make the GET call", checkMark, nil)

				defer resp.Body.Close()

				if resp.StatusCode == url.statusCode {
					t.Logf("\t\tShould recieve a \"%d\" status, %v", url.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould recieve a \"%d\" status, %v %v", url.statusCode, ballotX, resp.StatusCode)
				}

			}
		}
	}
}
