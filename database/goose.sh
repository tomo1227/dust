#!/bin/bash

CONFIG_FILE="./dust/config.toml"

DB_USER=$(dasel -f $CONFIG_FILE ".database.user" | sed -e "s/^'//" -e  "s/'\$//")
DB_PASSWORD=$(dasel -f $CONFIG_FILE ".database.password"| sed -e "s/^'//" -e  "s/'\$//")
DB_HOST=$(dasel -f $CONFIG_FILE ".database.host"| sed -e "s/^'//" -e  "s/'\$//")
DB_NAME=$(dasel -f $CONFIG_FILE ".database.dbname"| sed -e "s/^'//" -e  "s/'\$//")
PARSE_TIME=$(dasel -f $CONFIG_FILE ".database.parseTime"| sed -e "s/^'//" -e  "s/'\$//")

goose -dir="./dust/migrations" mysql "$DB_USER:$DB_PASSWORD@tcp($DB_HOST)/$DB_NAME?parseTime=$PARSE_TIME" $1
