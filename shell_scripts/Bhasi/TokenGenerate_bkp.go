package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"reflect"
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

var obj Asset
if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
        fmt.Println(resp.Body)
	fmt.Println(err)
}
//stringData := json.parse(body)
mySimpleMap := make(map[string]interface{})
err1 := json.Unmarshal([]byte(string(body)), &mySimpleMap) 
if err1 != nil {
    fmt.Println("error while unmarshal", err1)
} else {
    fmt.Println(reflect.TypeOf(mySimpleMap["Id"]))
    fmt.Printf("%f\n", mySimpleMap["Id"])
    //fmt.Println(int(mySimpleMap["Id"].(float64)))
}
myAsset := Asset{} 
err = json.Unmarshal([]byte(body), &myAsset) 
if err != nil {
    fmt.Println("error while unmarshal", err)
} else {
    fmt.Printf("%d", myAsset.Id)	
}   

//fmt.Println(obj.Id, obj.Name)
//err_3 := json.Unmarshal([]byte(body), &jsonData)
//fmt.Println(err3)
//fmt.Println(jsonData)
//ldmHeader.InfaToken = jsonData.InfaToken
//fmt.Println(jsonData.InfaToken)
//fmt.Println(ldmHeader.InfaToken)

}
