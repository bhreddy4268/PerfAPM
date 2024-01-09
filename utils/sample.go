package main

import (
        //"encoding/json"
        "fmt"
        "net/http"
        "io/ioutil"
        //"reflect"
)
type Asset struct {
    Id   string
    Name string
}

func main() {
//TokenURL := "http://irl75cmgperf06:7085/ldmadmin/web.isp/loginJs"
resp, err := http.Get("http://irl75cmgperf06:7085/ldmadmin/web.isp/loginJs")

fmt.Println(resp)
if err != nil {
    fmt.Println(err)
}
body, error := ioutil.ReadAll(resp.Body)
   if error != nil {
      fmt.Println(error)
   }
   // close response body
   resp.Body.Close()

   // print response body
   fmt.Println(string(body))

}
