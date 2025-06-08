package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the website URL (e.g. https://example.com): ")
    url, _ := reader.ReadString('\n')
    url = strings.TrimSpace(url)

    fmt.Println("Starting to send requests to:", url)

    for i := 1; ; i++ {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Request error:", err)
        } else {
            fmt.Printf("Request #%d sent, status: %s\n", i, resp.Status)
            resp.Body.Close()
        }
        time.Sleep(5 * time.Second)
    }
}
