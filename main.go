package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    url := os.Getenv("WEBHOOK_URL")
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"text":"` + os.Getenv("INVITE_MESSAGE") + `"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Message successfully sent")
}
