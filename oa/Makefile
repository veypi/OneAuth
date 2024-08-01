#
# Makefile
# Copyright (C) 2024 veypi <i@veypi.com>
# 2024-07-31 18:38
# Distributed under terms of the MIT license.
#

db=mysql://root:123456@localhost:3306/oa

all:
	@echo "Makefile needs your attention"


.PHONY:build
build:
	@go build -o ./build/oa

fmt:
	@goctl api format --dir ./api

gen_sql:
	@goctl model mysql ddl -c --database oa -d ./models -s ./protoc/sql/base.sql

# add sql script
# sqlx migrate --source ./protoc/sql add  base
# create database
# sqlx database create -D $(db)
gen_db:
	@sqlx database drop -D $(db)
	@sqlx database create -D $(db)
	@sqlx migrate --source ./protoc/sql run -D $(db)
	
gen_api:
	@goctl api go -api ./protoc/api/all.api -dir ./
