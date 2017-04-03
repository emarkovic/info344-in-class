GOOS=linux go build
docker build -t em42/zipsvr .
docker run -d -p 80:80 -e ADDR=:80 em42/zipsvr