# Steps to run
#### Backend - Local 
```diff
+ $ go get github.com/gorilla/websocket
+ $ go get github.com/satori/go.uuid
+ $ go controllers.go
 ```
#### Backend - Docker Container
```diff
+ $ docker build --rm -f "Backend\Dockerfile" -t goangularchatapp:latest "Backend"
+ $ docker run --rm -it goangularchatapp:latest
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
+ $ docker build --rm -f "Frontend\Dockerfile-dev" -t goangularchatapp:latest "Frontend"
+ $ npm i
 ```
