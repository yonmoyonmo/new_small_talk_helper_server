docker stop new_smthp_server
docker rm new_smthp_server
docker rmi new_smalltalk_server:latest

docker build -t new_smalltalk_server:latest .
docker run --name new_smthp_server -d -p 5000:5555 new_smalltalk_server:latest