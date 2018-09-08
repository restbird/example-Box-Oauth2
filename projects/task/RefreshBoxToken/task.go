package main
import "callapi"
import "fmt"
import "time"

import api2 "restbird/Box/OAuth2-GoLang/api2"


func main() {
		for {
			select {
			case <-time.Tick(time.Millisecond * 60000 * 60 ):
				fmt.Println("hello, curren time is :", time.Now())
				
				mytime := time.Now()
				callapi.SetGlobalVars("starttime", mytime)
				
				
				callapi.DoHttpRequestWithEnv("Box/OAuth2-GoLang", "api2", api2.CallBack{},  "Box")
				
				callapi.GetGlobalVars("starttime", &mytime)
				fmt.Println("starttime: ", mytime, "current Time:", time.Now())
				
				_, box_access_token:= callapi.GetGlobalString("box_access_token")
				fmt.Println("++box_access_token: " + box_access_token, "\n")

				_, box_refresh_token:= callapi.GetGlobalString("box_refresh_token")
				fmt.Println("--box_refresh_token: " + box_refresh_token, "\n\n")

				
				
			}
		}

}