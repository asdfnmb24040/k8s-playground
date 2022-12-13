docker run -d --name redis -p 6379:6379 redis
docker tag redis asdfnmb24040/redis-server
docker push asdfnmb24040/redis-server  