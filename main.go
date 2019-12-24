package main

import (
    "net/http"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    
)

func GetAir(url string) {
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    //bodyString, err := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(bodyString))
    if resp.StatusCode != 200 {
        fmt.Println("err")
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        panic(err)
    }

    doc.Find("#ctl05_gv tbody tr td").Each(func(i int, s *goquery.Selection) {
        // name
        fmt.Println( s.ChildrenFiltered(`a`).Text(), s.ChildrenFiltered(`span`).Text() )

    })
	return
   
}

func main() {
    url := "https://taqm.epa.gov.tw/taqm/tw/Aqi/Central.aspx?fm=AqiMap"
	GetAir(url)
}