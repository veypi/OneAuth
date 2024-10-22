#
# Makefile
# Copyright (C) 2024 veypi <i@veypi.com>
# 2024-10-22 16:36
# Distributed under terms of the GPL license.
#
cfg=~/.config/oa.yaml

run:
	@go run main.go -f ${cfg}

dbrest:
	@go run main.go -f ${cfg} db drop
	@go run main.go -f ${cfg} db migrate
	@go run main.go -f ${cfg} db init

