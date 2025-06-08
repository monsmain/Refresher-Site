package main

import (
    "context"
    "fmt"
    "github.com/chromedp/chromedp"
    "time"
)

func main() {
    var url string
    fmt.Print("Enter website URL (e.g. https://example.com): ")
    fmt.Scanln(&url)

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    for i := 1; i <= 5; i++ {
        err := chromedp.Run(ctx,
            chromedp.Navigate(url),
            chromedp.Sleep(10*time.Second),
        )

        if err != nil {
            fmt.Println("Error loading page:", err)
        } else {
            fmt.Printf("Visit #%d completed.\n", i)
        }

        time.Sleep(5 * time.Second)
    }
}
