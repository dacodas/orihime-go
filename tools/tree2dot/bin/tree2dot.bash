#!/bin/bash

HASH=582ae28bf547cae7ea6b0cb8c43cbb6d73f5c5f38ebe9c9b500c59f5cf83abe8 

TREE=$(build/orihime get text-tree ${HASH} --user 'dacoda.strack@gmail.com') 

echo "digraph G {"
echo "$TREE" | jq -r '.[0] | "\"\(.definitionHash)\"[label=\"\(.definition[0:10])...\"];"'
echo "$TREE" | jq -r '.[1:] | .[] | "\"\(.definitionHash)\"[label=\"\(.word)\"];"'
echo "$TREE" | jq -r '.[1:] | .[] | "\"\(.parentTextHash)\" -> \"\(.definitionHash)\";"'
echo "}"
