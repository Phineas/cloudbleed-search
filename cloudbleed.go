package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

func downloadLatestCloudbleed() {
  url := "https://raw.githubusercontent.com/pirate/sites-using-cloudflare/master/sorted_unique_cf.txt"
  tokens := strings.Split(url, "/")
  fileName := tokens[len(tokens)-1]
  fmt.Println("Downloading latest Cloudbleed file...")

  output, err := os.Create(fileName)
  if err != nil {
    fmt.Println("Error while creating", fileName, "-", err)
    return
  }

  defer output.Close()

  response, err := http.Get(url)
  if err != nil {
    fmt.Println("Error while downloading", url, "-", err)
    return
  }

  defer response.Body.Close()

  n, err := io.Copy(output, response.Body)
  if err != nil {
    fmt.Println("Error while downloading", url, "-", err)
    return
  }

  fmt.Println(n, "bytes downloaded.")
}

func checkDomain(domain string) string {
  if _, err := os.Stat("sorted_unique_cf.txt"); os.IsNotExist(err) {
    downloadLatestCloudbleed()
  }

  b, err := ioutil.ReadFile("sorted_unique_cf.txt")
  if(err != nil) {
    panic(err)
  }

  s := string(b)

  if(strings.Contains(s, "\n" + strings.ToLower(domain) + "\n")) {
    return domain + " is in the Cloudflare directory"
  } else if(strings.Contains(s, strings.ToLower(domain))) {
    return domain + " is not specifically in the Cloudflare directory, but there are domains that contain " + domain + " as a substring"
  } else {
    return domain + " is not in the Cloudflare directory"
  }
}

func main() {
  if _, err := os.Stat("sorted_unique_cf.txt"); os.IsNotExist(err) {
    downloadLatestCloudbleed()
  }
  
  var domain string
  fmt.Print("What would you like to search in Cloudflare? -> ")
  fmt.Scanln(&domain)

  fmt.Println(checkDomain(domain))
}
