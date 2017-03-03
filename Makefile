PACKAGE = github.com/hackebrot/go-librariesio

.DEFAULT_GOAL := help

.PHONY: cmd
cmd:  ##  Install application binaries
	@echo "+ $@"
	@go install $(PACKAGE)/...


.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
