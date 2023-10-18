major=$(strip $(shell awk -F ':' '/major/ {print $$2;}' cfg.yml))
minor=$(strip $(shell awk -F ':' '/minor/ {print $$2;}' cfg.yml))
patch=$(strip $(shell awk -F ':' '/patch/ {print $$2;}' cfg.yml))
version=v$(major).$(minor).$(patch)
vinfo=$(strip $(shell awk -F '$(version)' '/$(version)/ {print $$2;}' vuf.md))


version:
	@echo $(version) $(vinfo)

addtag:
ifeq ($(vinfo),)
	@echo please add version info in vuf.md
else
	@git tag -m '$(vinfo)' $(version)
	@git push origin $(version)
endif

dropTag:
	@git tag -d $(version)
	@git push origin :refs/tags/$(version)

run:
	go run *.go -d


syncDB:
	@scp -P 19529 oa.db root@alco.host:/root/

buildweb:
	@cd oaweb && yarn build
	@rm -rf ./oab/dist/*
	@mv ./oaweb/dist/spa/* ./oab/dist/

nats:
	@nats-server -c ./script/nats.cfg

.PHONY:oab
oab:
	@cd oab && cargo run -- -c ./cfg-demo.yml
