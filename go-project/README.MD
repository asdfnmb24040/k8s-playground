SET CGO_ENABLED=0&SET GOOS=linux&SET GOARCH=amd64&go build main.go

docker build --no-cache -t gin .
docker run -d --name gin -p 8080:8080 gin
docker tag gin asdfnmb24040/gin-server
docker login
docker push asdfnmb24040/gin-server  