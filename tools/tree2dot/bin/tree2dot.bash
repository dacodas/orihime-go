#!/bin/bash

HASH="$1"

TREE=$(build/orihime get text-tree ${HASH} --user 'dacoda.strack@gmail.com') 

echo "digraph G {"
echo "$TREE" | jq -r '.[0] | "\"\(.definitionHash)\"[label=\"\(.definition[0:10])...\"];"'
echo "$TREE" | jq -r '.[1:] | .[] | "\"\(.definitionHash)\"[label=\"\(.word)\"];"'
echo "$TREE" | jq -r '.[1:] | .[] | "\"\(.parentTextHash)\" -> \"\(.definitionHash)\";"'
echo "}"
