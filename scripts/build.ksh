export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  env GOOS=darwin go build -o bin/deposit-reg-ws.darwin depositregws
  res=$?
fi

if [ $res -eq 0 ]; then
  env GOOS=linux go build -o bin/deposit-reg-ws.linux depositregws
  res=$?
fi

exit $res
