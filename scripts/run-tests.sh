#!/bin/bash
echo Start test with next variables:
cat config/local.env
export `cat config/local.env`
echo ""
	
rm -rf autotetst/report.txt
./tests/node_modules/mocha/bin/mocha "tests/tests/**/*.js" --recursive --timeout 15000 --colours --exit | tee tests/tests_report.txt

