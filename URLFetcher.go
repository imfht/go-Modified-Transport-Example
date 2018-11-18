package awesomeProject1

import (
  "fmt"
  "log"
  "net/http"
)

func IPv4Fetcher(url string) {
  client := http.Client{Transport: MyTransportWrapper{IPv6Only: false}}
  out, err := client.Head(url)
  if err != nil {
    fmt.Println("error")
    log.Fatal(err)
  }
  fmt.Println(out.Header)
}
