export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  GOOS=darwin go build -o bin/deposit-reg-ws.darwin depositregws
  res=$?
fi

if [ $res -eq 0 ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/deposit-reg-ws.linux depositregws
  res=$?
fi

exit $res
