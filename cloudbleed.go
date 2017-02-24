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

func main() {
  if _, err := os.Stat("sorted_unique_cf.txt"); os.IsNotExist(err) {
    downloadLatestCloudbleed()
  }
  
  var domain string
  fmt.Print("What would you like to search in Cloudflare? -> ")
  fmt.Scanln(&domain)

  
  b, err := ioutil.ReadFile("sorted_unique_cf.txt")
  if(err != nil) {
    panic(err)
  }

  s := string(b)

  if(strings.Contains(s, "\n" + domain + "\n")) {
    fmt.Println(domain + " is in the Cloudflare directory")
  } else {
    fmt.Println(domain + " is not in the Cloudflare directory")
  } 
}
