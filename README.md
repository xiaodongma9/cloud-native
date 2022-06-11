# cloud-native


# build docker image
docker build --tag look4xiaodong/cloudnative-module2:0.1 .

# push image
docker push look4xiaodong/cloudnative-module2:0.1

# run image
docker run -d -p 80:80 look4xiaodong/cloudnative-module2:0.1

# check log
docker logs -f <containerID>
