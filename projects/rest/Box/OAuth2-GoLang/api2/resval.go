package api
import "callapi"
import "net/http"
import "io/ioutil"
import	"encoding/json"
import "fmt"

type MyDATA struct {
	Access_token    string `json:"access_token,omitempty"`
	Refresh_token    string `json:"refresh_token,omitempty"`
}


func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {

    var body []byte
    var err error
    var data MyDATA = MyDATA{}
    
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
	    fmt.Println("read body failed.")
		return true
	} 
	
	//fmt.Println(body)
	
	if err = json.Unmarshal(body, &data); err != nil {
	    fmt.Println("conver body to json failed")
		return true
	}	
	
	callapi.SetGlobalString("box_access_token", data.Access_token)
	callapi.SetGlobalString("box_refresh_token", data.Refresh_token)
	
	_, box_access_token := callapi.GetGlobalString("box_access_token")
	_, box_refresh_token := callapi.GetGlobalString("box_refresh_token")
	
	fmt.Println("box_access_token: " + box_access_token)
	fmt.Println("box_refresh_token: " + box_refresh_token)

	return true
}
