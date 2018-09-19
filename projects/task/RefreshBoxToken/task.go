package main
import "callapi"
import "fmt"
import "time"

import api2 "restbird/Box/OAuth2-GoLang/api2"


func main() {
        fmt.Println("Start Task ")
        task_RefereshBoxToken_count := 0
		for {
			select {
    			case <-time.Tick(time.Millisecond * 60000 * 60 ): //Set time interval of 1 hour
    				fmt.Println("hello, curren time is :", time.Now())
    				
    				mytime := time.Now()
    				callapi.SetGlobalVars("task_RefereshBoxToken_starttime", mytime)
    				
    				task_RefereshBoxToken_count++
    				callapi.SetGlobalVars("task_RefereshBoxToken_count", task_RefereshBoxToken_count)
    				
    				callapi.DoHttpRequestWithEnv("Box/OAuth2-GoLang", "api2", api2.CallBack{},  "Box")
    				
    				callapi.GetGlobalVars("task_RefereshBoxToken_starttime", &mytime)
    				fmt.Println("task_RefereshBoxToken_starttime: ", mytime, "current Time:", time.Now())
    				
    				callapi.GetGlobalVars("task_RefereshBoxToken_count", &task_RefereshBoxToken_count)
    				fmt.Println("task_RefereshBoxToken_count: ", task_RefereshBoxToken_count)
    				
    				_, box_access_token:= callapi.GetGlobalString("box_access_token")
    				fmt.Println("++box_access_token: " , box_access_token)
    
    				_, box_refresh_token:= callapi.GetGlobalString("box_refresh_token")
    				fmt.Println("--box_refresh_token: " + box_refresh_token, "\n\n")

				
				
			}
		}

}