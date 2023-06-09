# docker run -d --name golang-mongo \
# 	-e MONGO_INITDB_ROOT_USERNAME=mongo \
# 	-e MONGO_INITDB_ROOT_PASSWORD=mongo \
# 	-p 27017:27017 mongo

docker run -d --name golang-mongo -p 27017:27017 mongo