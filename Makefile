# Bureaucr.at Coding Challenge
# Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/
# London, Jun 2021
# github.com/narcismpap/bureaucrat

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	docker build . -t docker.pkg.github.com/narcismpap/bureaucrat/bureaucrat:latest

test:
	cd bureaucrat && go test

example:
	docker run --rm --network none \
		docker.pkg.github.com/narcismpap/bureaucrat/bureaucrat:latest \
		-s "/GoT.json"\
		-r "GoT-005"\
		-l "GoT-006"
