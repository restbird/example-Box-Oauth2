This is the example repository for demo how to automatiocally retrieve OAuth 2.0 token for BOX API Server with [Restbird](https://restbird.org), which described in this blog, [A practical way to automate testing Oauth 2.0 Service](https://restbird.org/blog/2018/09/01/automate-testing-oauth2.html)

To use this repo
* Download Restbird docker
~~~
docker pull restbird/rest
~~~

* Run Restbird docker with this repo 
~~~
docker run -ti -p {host-port}:8080 -v {path-to-project}:/data/restbird restbird/rest
~~~

The default user credential for Restbird is admin/admin

The Box account used in this exampe is demo@restbird.com/DemoRestbird123!

You can definitely use your own Box account, modify the client_id and client_secret in "Box" environment.

Related articles:

* [OAuth application setup in BOX](https://developer.box.com/docs/setting-up-an-oauth-app)
* [Environment in Restbird](https://restbird.org/docs/rest.html#environment-vaaribles)
