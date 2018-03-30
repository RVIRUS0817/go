package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    apiUrl := "https://api.chatwork.com/"
    resource := "/v2/rooms/xxxxx/messages"

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := fmt.Sprintf("%v", u) 

    data := url.Values{}
    data.Set("body", "go!go!vuls-chatwork")
    fmt.Println(data.Encode())

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
    r.Header.Add("X-ChatWorkToken", "xxxxxxxx")
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := client.Do(r)

    defer resp.Body.Close()
    contents, _ := ioutil.ReadAll(resp.Body)

    fmt.Println(resp.Status)
    fmt.Printf("result: %s\n", contents)
    fmt.Println(urlStr)
}
