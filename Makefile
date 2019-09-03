.PHONY: build-docker

build-docker:
	docker build -t "sebiwi/connect-four" .

run-docker:
	docker run -ti --rm sebiwi/connect-four
