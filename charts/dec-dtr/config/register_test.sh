#!/bin/sh

DAPS_URL=http://127.0.0.1:50816
DAPS_TOKEN_URL=$DAPS_URL/token
DAPS_ADMIN_API_URL=$DAPS_URL/api/v1
DAPS_ADMIN=clientId
DAPS_ADMIN_SECRET=clientSecret

getJsonAdminToken() {
curl -s -X POST --location $DAPS_TOKEN_URL \
    --data-urlencode "grant_type=client_credentials" \
    --data-urlencode "client_id=$DAPS_ADMIN" \
    --data-urlencode "client_secret=$DAPS_ADMIN_SECRET" \
    --data-urlencode "scope=omejdn:admin"
}


getAdminToken() {
  getJsonAdminToken | jq -r '.access_token'
}

getJson() {
  jq --null-input \
    --arg CLIENT_ID $CLIENT_ID \
    --arg CLIENT_NAME $CLIENT_NAME \
    --arg CLIENT_SECURITY_PROFILE $CLIENT_SECURITY_PROFILE \
    --arg REFERRING_CONNECTOR $REFERRING_CONNECTOR \
  '
  {
      "client_id": $CLIENT_ID,
      "name": $CLIENT_NAME,
      "token_endpoint_auth_method": "private_key_jwt",
      "scope": [
        "idsc:IDS_CONNECTOR_ATTRIBUTES_ALL"
      ],
      "grant_types": [
        "client_credentials"
      ],
      "attributes": [
        {
          "key": "idsc",
          "value": "IDS_CONNECTOR_ATTRIBUTES_ALL"
        },
        {
          "key": "@type",
          "value": "ids:DatPayload"
        },
        {
          "key": "@context",
          "value": "https://w3id.org/idsa/contexts/context.jsonld"
        },
        {
          "key": "securityProfile",
          "value": $CLIENT_SECURITY_PROFILE
        },
        {
          "key": "referringConnector",
          "value": $REFERRING_CONNECTOR
        }
      ]
  }'
}

getCertJson() {
  jq --null-input --arg cert "$(openssl x509 -in $1)" '{"certificate": $cert}'
}

if [ ! $# -ge 1 ] || [ ! $# -le 4 ]; then
    echo "Usage: $0 CLIENT_NAME REFERRING_CONNECTOR (SECURITY_PROFILE) (CERTFILE) or $0 -check"
    exit 1
fi

ADMIN_TOKEN=$(getAdminToken)

if [ $1 = "-check" ]; then
  curl --location $DAPS_ADMIN_API_URL/config/clients \
      --oauth2-bearer "$ADMIN_TOKEN"
  exit 0
fi

CLIENT_NAME=$1
REFERRING_CONNECTOR=$2
CLIENT_SECURITY_PROFILE=$3
CERTFILE=$4

if [ -z "$CLIENT_SECURITY_PROFILE" ] || [ -f "$CLIENT_SECURITY_PROFILE" ]; then
  CLIENT_SECURITY_PROFILE="idsc:BASE_SECURITY_PROFILE"
  CERTFILE=$3
fi
if [ -n "$CERTFILE" ]; then
    [ ! -f "$CERTFILE" ] && (echo "Cert not found"; exit 1)
else
    openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout "keys/${CLIENT_NAME}.key" -out "keys/${CLIENT_NAME}.crt" -config openssl.cnf
    CERTFILE=keys/${CLIENT_NAME}.crt
fi

SKI="$(openssl x509 -in "$CERTFILE" -text | grep -A1 "Subject Key Identifier" | tail -n 1 | tr -d ' ')"
AKI="$(openssl x509 -in "$CERTFILE" -text | grep -A1 "Authority Key Identifier" | tail -n 1 | tr -d ' ')"
CLIENT_ID="$SKI:$AKI"
[ -z $CLIENT_ID ] && (echo "certificate does not contain SKI:AKI extension"; exit 1)

curl -X POST --location $DAPS_ADMIN_API_URL/config/clients \
  --oauth2-bearer "$ADMIN_TOKEN" \
  --data-raw "$(getJson)"

curl -X POST --location $DAPS_ADMIN_API_URL/config/clients/$CLIENT_ID/keys   \
  --oauth2-bearer "$ADMIN_TOKEN" \
  --data-raw "$(getCertJson "$CERTFILE")"

echo "Client has been registered successfully!!!"

