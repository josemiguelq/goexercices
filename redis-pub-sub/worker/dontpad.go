package worker

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
    "redis-lab/redisb"
)

func PostDontpad(msg string) {
    resp, err := http.PostForm(os.Getenv("DONTPAD_URL")+"/golangoexercisestoday", url.Values{
        "text": {msg},
    })

    if nil != err {
        fmt.Println("errorination happened getting the response", err)
        return
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if nil != err {
        fmt.Println("errorination happened reading the body", err)
        return
    }

    fmt.Println(string(body[:]))
}

var latest string
func Changed(){
    current := redisb.Get(redisb.CreateClient())
    fmt.Println(current)
    fmt.Println(latest)

    if current != latest {
        latest = current
        PostDontpad(latest)
    }
    fmt.Println("n mudou")
}