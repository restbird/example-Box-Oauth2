package api
import "callapi"
import "net/http"
import "io/ioutil"
import	"encoding/json"
import "fmt"
type MyEntry struct {
	Id  string        `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	Type string			`json:"type,omitempty"`
}	

type MyCollection struct{
	Entries []MyEntry	`json:"entries,omitempty"`
}
type MyDATA struct {
	Item_collection  MyCollection        `json:"item_collection,omitempty"`
}


func (c  CallBack) ResponseValidate(resp *http.Response, ctx *callapi.RestBirdCtx) bool {

    var body []byte
    var err error
    var data MyDATA
    
    if resp.StatusCode == 200 {
    	if body, err = ioutil.ReadAll(resp.Body); err != nil {
    	    fmt.Println("read body failed.")
    		return false
    	} 
    	
    	//fmt.Println(body)
    	
    	if err = json.Unmarshal(body, &data); err != nil {
    	    fmt.Println("conver body to json failed")
    		return false
    	}	
    	
    //	fmt.Println("data: ", data)
    	for i, v := range data.Item_collection.Entries {
    
    // 		fmt.Println(i)
    // 		fmt.Println(v.Name)
    		if v.Name == ctx.GetVars("folder_name") {
                 ctx.SetVars("folder_id", v.Id)
                 fmt.Println("Found folder: ")
    	         fmt.Println("i, folder_id, folder_name: " , i, ",", ctx.GetVars("folder_id"), ",", ctx.GetVars("folder_name"))
    	         return true
    	    }
    		
    	}
    	
    	fmt.Println("Can't find folder: ", ctx.GetVars("folder_name"))
    	return true
    }

	return false
}
