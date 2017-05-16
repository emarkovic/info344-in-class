# build dat ish
docker build -t em42/nodetest .

# check if it worked by running 
docker images

# fire it up
docker run -d -p 80:80 --name nodetest em42/nodetest