package main

import (
    "bytes"
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "io/ioutil"
    "net/http"
    "os"
)

func postJson(url string, json []byte) (error) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    return err
}

func HandleRequest(ctx context.Context) (string, error) {
    var url = os.Getenv("WEBHOOK_URL")
    var json = []byte(`{"text":"` + os.Getenv("INVITE_MESSAGE") + `"}`)
    err := postJson(url, json)
    if err != nil {
        return "Error", err
    }
    return "Message successfully sent", nil
}

func main() {
    lambda.Start(HandleRequest)
}
