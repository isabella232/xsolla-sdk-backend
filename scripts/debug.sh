#!/bin/bash
clear
echo Server Debug-mode will be started with next variables:
cat config/local.env
echo ""
export `cat config/local.env`
dlv exec --headless --continue --listen :2345 --accept-multiclient cmd/web-api/web-api