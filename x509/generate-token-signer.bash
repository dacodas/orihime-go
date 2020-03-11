#!/bin/bash

set -euo pipefail

: ${GENERATE_NEW_KEY:=false}

PRIVATE_KEY=token-signer.private.pem
PUBLIC_KEY=token-signer.public.pem

FILE_TO_SIGN="$1"
DIGEST_OUTPUT="$2"
SIGNATURE_OUTPUT="$3"

PKEYUTL_OPTIONS="-pkeyopt digest:sha256 -pkeyopt rsa_padding_mode:pss -inkey ${PRIVATE_KEY} "
PKEYUTL_SIGN_OPTIONS="-pkeyopt rsa_pss_saltlen:256 -pkeyopt rsa_mgf1_md:sha256"

if [ "x${GENERATE_NEW_KEY}" = "xtrue" ]
then
	openssl genrsa -out ${PRIVATE_KEY} 4096
	openssl rsa -outform PEM -in ${PRIVATE_KEY} -pubout -out ${PUBLIC_KEY}
fi

openssl dgst -sha256 -binary -out ${DIGEST_OUTPUT} ${FILE_TO_SIGN}
openssl pkeyutl ${PKEYUTL_OPTIONS} ${PKEYUTL_SIGN_OPTIONS} -in ${DIGEST_OUTPUT} -sign -out ${SIGNATURE_OUTPUT}
openssl pkeyutl ${PKEYUTL_OPTIONS} -in ${DIGEST_OUTPUT} -verify -sigfile ${SIGNATURE_OUTPUT}
# openssl dgst -sha256 -verify ${PUBLIC_KEY} -signature ${SIGNATURE_OUTPUT} ${FILE_TO_SIGN}
