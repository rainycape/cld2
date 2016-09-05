#!/bin/bash


set -x;
set -o errexit

GIT=https://github.com/CLD2Owners/cld2.git
DIR=cld2

# Checkout a temporary copy of the repository
git clone ${GIT} ${DIR}

cp ${DIR}/internal/*.h .
cp ${DIR}/public/*.h .


# Use the script for compiling the lib to gather the files
# for the Go package.
cd ${DIR}/internal

calls=0

function g++ {
    # First time g++ is called it builds the "non-full" lib
    if [ $calls -lt 1 ]; then
        ((calls+=1))
        return;
    fi

    while (( "$#" )); do
        if [ "$1" = '-o' ]; then
            break;
        fi
        if [[ $1 == -* ]]; then
            shift;
            continue;
        fi
        cp $1 ../../;
        shift;
    done
}

source ./compile_libs.sh

cd -

sed -i '' -e 's/..\/public\///' *.h *.cc
sed -i '' -e 's/..\/internal\///' *.h

rm -fr ${DIR}
