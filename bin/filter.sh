#!/bin/bash

#
# Filter JSON payload from the array contained in the file passed as a parameter
# to this script and output corresponding files for each item in the array on
# the folder specified as the second parameter of this script.
#

FILE=$1
OUTPUT=$2
SIZE=$(jq '. | length' $FILE)

for N in $(seq 0 $((SIZE-1)))
do
    COMMAND="jq '.[${N}]' data.json"
    SERVER=$(eval ${COMMAND})
    echo ${SERVER} > ${OUTPUT}/data-${N}.json
    echo "Wrote data-${N}.json in ${OUTPUT}"
done