package main

import (
	"net/http"
	"testing"
)

const (
	checkMark = "/u02713"
	ballotX   = "/u2717"
)

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the GET call", ballotX, err)
			}
			t.Log("\t\tShould be able to make the GET call", checkMark, nil)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould recieve a \"%d\" status, %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould recieve a \"%d\" status, %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
