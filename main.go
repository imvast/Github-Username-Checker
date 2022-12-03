package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "github.com/RomainMichau/cloudscraper_go/cloudscraper"
    "github.com/imvast/vast.goutil"
)


func init() {
    rand.Seed(time.Now().UnixNano())
    goutil.ClearConsole()
}

func sendReq(name string, debug bool) {
    client, _ := cloudscraper.Init(false, false)
    link := fmt.Sprintf("https://github.com/%v", name)
	  resp, _ := client.Get(link, make(map[string]string), "")
    if resp.Status == 404 {
        goutil.StatLog("green", resp.Status, name)
    }
    if resp.Status == 200 && debug == true {
	      goutil.StatLog("red", resp.Status, name)
    }
    if resp.Status == 429 && debug == true {
        goutil.StatLog("yellow", resp.Status, "log timeout soon ig")//resp.Body)
        time.Sleep(10)
  }
}


func main() {
    length := goutil.Question("Length > ")
    intLength, _ := strconv.Atoi(length)
    for 0==0 {
        var name = goutil.RandStr(intLength)
        sendReq(name, false) // debug: true/false (current: false)
    }
}
