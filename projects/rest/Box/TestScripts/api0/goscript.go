package api
import "callapi"
import "time"
import "fmt"
import api0 "restbird/Box/File/api0" 
import api1 "restbird/Box/File/api1"

type CallBack struct {}

func (c  CallBack) GoScripts(ctx *callapi.RestBirdCtx) bool{
    var folder_id = ""
    ctx.SetVars("folder_name", "Demo7")
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
     
     //Add sleep to mimick a long test case
      fmt.Println("Sleeping for 30 seconds...")
      time.Sleep(30 * time.Second)
      fmt.Println("wakeup")
     return true
}
