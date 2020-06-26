#!/bin/bash

set -e
# download Kong deb

sudo apt-get update
sudo apt-get install openssl libpcre3 procps perl wget

function setup_kong(){
  SWITCH="1.3.100"

  URL="https://kong.bintray.com/kong-deb/kong-${KONG_VERSION}.xenial.all.deb"

  if [[ "$KONG_VERSION" > "$SWITCH" ]];
  then
  URL="https://kong.bintray.com/kong-deb/kong-${KONG_VERSION}.xenial.amd64.deb"
  fi

  /usr/bin/curl -sL $URL -o kong.deb
}

function setup_kong_enterprise(){
  KONG_VERSION=${enterprise-1.5.0.2#"enterprise-"}
  URL="https://kong.bintray.com/kong-enterprise-edition-deb/dists/kong-enterprise-edition-${KONG_VERSION}.xenial.all.deb"
  RESPONSE_CODE=$(/usr/bin/curl -sL \
    -w "%{http_code}" \
    -u $KONG_ENTERPRISE_REPO_USERNAME:$KONG_ENTERPRISE_REPO_PASSSWORD \
    $URL -o kong.deb)
  if [[ $RESPONSE_CODE != "200" ]]; then
    echo "error retrieving kong enterprise package. response code ${RESPONSE_CODE}"
  fi
}

if [[ $KONG_VERSION == *"enterprise"* ]]; then
  setup_kong_enterprise
else
  setup_kong
fi

sudo dpkg -i kong.deb
echo $KONG_LICENSE_DATA | sudo tee /etc/kong/license.json
export KONG_LICENSE_PATH=/tmp/license.json

sudo kong migrations bootstrap
sudo kong version
sudo kong start
