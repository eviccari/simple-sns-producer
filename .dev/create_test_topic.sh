TOPIC=$1

aws --endpoint-url=http://localhost:4566  sns create-topic --name $TOPIC