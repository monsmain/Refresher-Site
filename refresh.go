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

    fmt.Print("Enter the website URL: ")
    url, _ := reader.ReadString('\n')
    url = strings.TrimSpace(url)

    client := &http.Client{}

    for i := 1; ; i++ {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Println("Error creating request:", err)
            continue
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
        req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
        req.Header.Set("Accept-Language", "en-US,en;q=0.5")
        req.Header.Set("Connection", "keep-alive")

        resp, err := client.Do(req)
        if err != nil {
            fmt.Println("Request error:", err)
            continue
        }

        fmt.Printf("Request #%d sent. Status: %s\n", i, resp.Status)
        resp.Body.Close()

        time.Sleep(2 * time.Second)
    }
}
