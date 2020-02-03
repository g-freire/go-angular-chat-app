# Steps to run
#### Backend - Local 
```diff
+ $ go get github.com/gorilla/websocket
+ $ go get github.com/satori/go.uuid
+ $ go controllers.go
 ```
#### Backend - Docker Container
```diff
+ $ docker build --rm -f "_Backend\Dockerfile" -t goangularchatapp_back:latest "_Backend"
+ $ docker run --rm -it -p 5000:5000/tcp goangularchatapp_back:latest
 ```
 #
#### Frontend - Local
```diff
+ $ nvm use 11
+ $ npm i
+ $ npm start
 ```
 #### Frontend - Docker Container
```diff
+ $ docker build --rm -f "_Frontend\Dockerfile-dev" -t goangularchatapp_front:latest "_Frontend"
+ $ docker run --rm -it -p 80:80/tcp goangularchatapp_front:latest
 ```
