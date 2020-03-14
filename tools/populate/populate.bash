#!/bin/bash

set -euo pipefail

NUMBER="$1"

TEXT="Here is a text from some source ${NUMBER}"
TEXT_SOURCE="Some source"
WORD="source ${NUMBER}"
DEFINITION="Place to get stuff from ${NUMBER}"
DEFINITION_SOURCE="Some dictionary"

PARENT_TEXT_HASH="$(echo -n "${TEXT}" | sha256sum | sed 's/  -//')"

# orihime add source "${TEXT_SOURCE}"
# orihime add source "${DEFINITION_SOURCE}"

orihime add text "${TEXT}" --source "${TEXT_SOURCE}"
orihime add child-word "${WORD}" --user 'dacoda.strack@gmail.com' --source "${DEFINITION_SOURCE}" --definition "${DEFINITION}" --parent-text "${PARENT_TEXT_HASH}"

echo ${PARENT_TEXT_HASH}
