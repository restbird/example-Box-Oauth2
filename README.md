This is the example repository for demo how to automatiocally retrieve OAuth 2.0 token for BOX API Server with [Restbird](https://restbird.org), which described in this blog, [A practical way to automate testing Oauth 2.0 Service](https://restbird.org/blog/2018/09/01/automate-testing-oauth2.html) and [A practical way to automate testing Oauth 2.0 Service - Part 2](https://restbird.org/blog/2018/09/07/automate-testing-oauth2-2.html)

To use this repo


* Download Restbird docker3

~~~
docker pull restbird/rest
~~~

* Run Restbird docker with this repo 

~~~
docker run -ti -p {host-port}:8080 -v {path-to-project}:/data/restbird restbird/rest
~~~

Check [Restbird tutorial] (https://restbird.org/tutorial/tutorial.html) for detail 

The default user credential for Restbird is admin/admin

The Box account used in this exampe is demo@restbird.com/DemoRestbird123!

You can definitely use your own Box account, modify the client_id and client_secret in "Box" environment.

Related articles:

* [OAuth application setup in BOX](https://developer.box.com/docs/setting-up-an-oauth-app)

Jenkins pipeline example used to test this demo 

~~~
node {
    def payload 
   
    stage('Build') {
    //...
    }
    stage('Deploy') {
    //..
    }
    stage('Run test') {
        def host = 'http://192.168.1.178:10000'
        def basicAuth = 'Basic YWRtaW46YWRtaW4='
      
        println('Call Restbird API to run test:')
        def reqBody = '{"casepath":"Box/TestScripts","apis":["api0"]}'
        def response = httpRequest httpMode:'POST', customHeaders: [[name: 'Authorization', value: basicAuth]], requestBody: reqBody, url:host+"/v1/rest/run"
        println('Status: '+response.status)
        println('Response: '+response.content)
        payload = readJSON text: response.content
        def historyId = payload.his.id
        println('History_id: '+ historyId)
        
        def historyReqBody = '{"project":"Box/TestScripts","id":"' + payload.his.id + '", "immediatereturn": true}'
   
        for(int i=0;i<10;i++){
            println('Call Restbird API to get result: '  + i)

            def historyResponse = httpRequest httpMode:'POST', customHeaders: [[name: 'Authorization', value: basicAuth]], requestBody: historyReqBody, url:host+"/v1/rest/runresult"
            println('Status: '+historyResponse.status)
            println('Response: '+historyResponse.content)
          
              if(historyResponse.status == 200){
                  
                printConsoleLog(host, basicAuth, historyId)
                
                payload = readJSON text: historyResponse.content
                if(payload.code == 0){
                    if(payload.his.responseval.result == true){
                         currentBuild.result = 'SUCCESS'
                    }else{
                         currentBuild.result = 'FAILURE'
                    }
                    return
                }else if(payload.code == -1){
                    println("Test unfinshed, check back in 10 seconds")
                    sleep 10
                }else{
                    println("Test error" + payload.code + ", " + payload.info)
                    currentBuild.result = 'FAILURE'
                    return
                }
            }else{
                currentBuild.result = 'FAILURE'
                return
            }

        }
        
        //return timeout
        currentBuild.result = 'FAILURE'
        return
    }
}
def printConsoleLog(host, basicAuth, historyId){
   println('--Call Restbird API to get console log:')
   def consoleResponse = httpRequest httpMode:'GET', customHeaders: [[name: 'Authorization', value: basicAuth]], url:host+"/v1/rest/his/console?project=Box/TestScripts&id=" + historyId
   println('--Console Status: '+consoleResponse.status)
   println('--Console Response: '+consoleResponse.content)  
}
~~~
