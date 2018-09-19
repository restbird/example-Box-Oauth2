package api
import "callapi"
import "net/http"
import "fmt"
import "io/ioutil"
import	"encoding/json"
type MyDATA struct {
    Id string `json:"id, omitempty"`
}

func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {
	if resp.StatusCode == 201 {
        var body []byte
    	var err error
    	var data MyDATA
    
    	if body, err = ioutil.ReadAll(resp.Body); err != nil {
    		fmt.Println("read body failed.")
    		return false
    	}
    
    	if err = json.Unmarshal(body, &data); err != nil {
    		fmt.Println("conver body to json failed:", err.Error())
    		return false
    	}	    
    	
    	ctx.SetVars("folder_id", data.Id)
    	fmt.Println("folder_id", ctx.GetVars("folder_id"))
    	return true
	}
	return false
}
