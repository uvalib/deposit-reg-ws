#
#
#

# ensure we have and endpoint
if [ -z "$DEPOSITREG_URL" ]; then
   echo "ERROR: DEPOSITREG_URL is not defined"
   exit 1
fi

# issue the command
echo "$DEPOSITREG_URL"
curl $DEPOSITREG_URL/version

exit 0

#
# end of file
#
