SHELL=/bin/bash

.PHONY: install
install:
	go install go.osspkg.com/goppy/v2/cmd/goppy@latest
	goppy setup-lib
	cd ./ui && yarn install --force --ignore-scripts

.PHONY: lint
lint:
	goppy lint

.PHONY: license
license:
	goppy license

.PHONY: build_back
build_back:
	goppy build --arch=amd64

.PHONY: build_front
build_front:
	cd ./ui && yarn run build

.PHONY: tests
tests:
	goppy test

.PHONY: pre-commit
pre-commit: install license build_front lint tests build_back

.PHONY: ci
ci: pre-commit

run_back:
	go run cmd/urione/main.go --config=config/config.dev.yaml
run_front:
	cd ./ui && yarn run start

deb:
	deb-builder build
