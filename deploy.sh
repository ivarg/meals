#! /bin/bash

cd web
gulp build
cd ..
gox -osarch linux/amd64
rm meals-img.tar
docker rmi meals-img
docker build --force-rm -t meals-img .
docker save meals-img > meals-img.tar
scp meals-img.tar root@178.62.208.192:/root
ssh root@178.62.208.192 docker stop meals
ssh root@178.62.208.192 docker rm meals
ssh root@178.62.208.192 docker rmi meals-img
ssh root@178.62.208.192 docker load < meals-img.tar
ssh root@178.62.208.192 docker run --name meals -d -p 8080:8080 meals-img
