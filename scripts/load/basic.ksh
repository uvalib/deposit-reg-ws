#
# basic load test
#

if [ -z "$DEPOSITREG_URL" ]; then
   echo "ERROR: DEPOSITREG_URL is not defined"
   exit 1
fi

if [ -z "$API_TOKEN" ]; then
   echo "ERROR: API_TOKEN is not defined"
   exit 1
fi

LT=../../bin/bombardier
if [ ! -f "$LT" ]; then
   echo "ERROR: Bombardier is not available"
   exit 1
fi

# set the test parameters
endpoint=$DEPOSITREG_URL
concurrent=10
count=5000
url=options?auth=$API_TOKEN

CMD="$LT -c $concurrent -n $count -l $endpoint/$url"
echo "Host = $endpoint, count = $count, concurrency = $concurrent"
echo $CMD
$CMD
exit $?

#
# end of file
#
