# go-web-bootstrap
***自用Go Web代码模版***
# 如何运行?
1. clone project
   ```shell
   cd $GOPATH/src && git clone git@github.com:FromMeToMyself/go-web-bootstrap.git
   ```

2. build http server
   ```shell
   cd cmd/http && go build server.go
   ```

3. run http server
   ```shell
   ./server -config $GOPATH/src/go-web-bootstrap/config
   ```
