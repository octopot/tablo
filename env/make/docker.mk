ifndef VERSION
$(error Please define VERSION variable)
endif

.PHONY: docker-build
docker-build:
	docker build -f env/docker/service/Dockerfile \
	             -t octopot/tablo:$(VERSION) \
	             -t octopot/tablo:latest \
	             -t quay.io/octopot/tablo:$(VERSION) \
	             -t quay.io/octopot/tablo:latest \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-publish
docker-publish: docker-build docker-push

.PHONY: docker-push
docker-push:
	docker push octopot/tablo:$(VERSION)
	docker push octopot/tablo:latest
	docker push quay.io/octopot/tablo:$(VERSION)
	docker push quay.io/octopot/tablo:latest

.PHONY: docker-run
docker-run:
	docker run --rm -it \
	           --name tablo-dev \
	           -p 8080:8080 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           octopot/tablo:$(VERSION) run --with-profiling --with-monitoring
