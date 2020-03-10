#!/bin/bash

set -euo pipefail

: ${GENERATE_NEW_KEY:=false}

PRIVATE_KEY=token-signer.private.pem
PUBLIC_KEY=token-signer.public.pem

FILE_TO_SIGN="$1"
SIGNATURE_OUTPUT="$2"

if [ "x${GENERATE_NEW_KEY}" = "xtrue" ]
then
	openssl genrsa -out ${PRIVATE_KEY} 4096
	openssl rsa -outform PEM -in ${PRIVATE_KEY} -pubout -out ${PUBLIC_KEY}
fi

openssl dgst -sha256 -sign ${PRIVATE_KEY} -out ${SIGNATURE_OUTPUT} ${FILE_TO_SIGN}
openssl dgst -sha256 -verify ${PUBLIC_KEY} -signature ${SIGNATURE_OUTPUT} ${FILE_TO_SIGN}
