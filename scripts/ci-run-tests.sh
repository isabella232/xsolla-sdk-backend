#!/bin/bash
echo Start CI test

echo "-------> Run API"
cmd/web-api/web-api > /dev/null 2>&1 &
until curl -s 127.0.0.1:8080/healthcheck; do sleep 1; echo "Waiting for api..."; done
echo ""

set | grep X_
echo "-------> Run tests"
./tests/node_modules/mocha/bin/mocha "tests/tests/**/*.js" --recursive --timeout 15000 --colours --exit
