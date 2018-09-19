package api
import "callapi"
//import "time"
import "fmt"
import api0 "restbird/Box/File/api0" 
import api1 "restbird/Box/File/api1"

type CallBack struct {}

func (c  CallBack) GoScripts(ctx *callapi.RestBirdCtx) bool{
    var folder_id = ""
    ctx.SetVars("folder_name", "Demo6")
    ctx.SetVars("folder_id", "")
    
    reterr, retbool,retMsg := callapi.DoHttpRequest("Box/File", "api0", api0.CallBack{},  ctx)
    fmt.Println(reterr, retbool,retMsg)
    if !retbool {
        fmt.Println("callapi failed", retMsg)
        return false
    }
    
    folder_id = ctx.GetVars("folder_id")
    
    if folder_id == ""{
        fmt.Println("Folder doesn't exist, create a new one: ", ctx.GetVars("folder_name"))
        reterr, retbool,retMsg := callapi.DoHttpRequest("Box/File", "api1", api1.CallBack{},  ctx)
        fmt.Println(reterr, retbool,retMsg)
        
        if !retbool {
          fmt.Println("callapi failed", retMsg)
          return false
        }
    
        folder_id = ctx.GetVars("folder_id")
    }
    
     fmt.Println(folder_id)
     
    //  fmt.Println("Sleeping for 60 seconds...")
    //  time.Sleep(60 * time.Second)
     return true
}
