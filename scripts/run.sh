#!/bin/bash
clear
echo Server will be started with next variables:
cat config/local.env
echo ""
export `cat config/local.env`
cmd/web-api/web-api