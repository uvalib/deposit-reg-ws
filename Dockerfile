FROM alpine:3.8

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates && rm -fr /var/cache/apk/*

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# set the timezone appropriatly
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /depositreg-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_HOME/assets
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# Add the RDS certificates
COPY data/rds-combined-ca-bundle.pem /etc/ssl/certs

# Specify the user
USER webservice

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY data/container_bash_profile /home/webservice/.profile
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY bin/deposit-reg-ws.linux $APP_HOME/bin/deposit-reg-ws
COPY assets/* $APP_HOME/assets/

# Add the build tag
COPY buildtag.* $APP_HOME/
