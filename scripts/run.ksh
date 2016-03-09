if [ -z "$DBPASSWD" ]; then
   echo "ERROR: DBPASSWD must be defined"
   exit 1
fi

bin/deposit-reg-ws.darwin --dbpassword $DBPASSWD
