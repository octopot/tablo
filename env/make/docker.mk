IMAGE   = octopot/tablo
BACKUP  = quay.io/octopot/tablo
VERSION = 1.x

.PHONY: docker-backup
docker-backup:
	@docker tag  $(IMAGE):$(VERSION)  $(BACKUP):$(VERSION)
	@docker tag  $(IMAGE):latest      $(BACKUP):latest
	@docker push $(BACKUP):$(VERSION)
	@docker push $(BACKUP):latest

.PHONY: docker-build
docker-build:
	docker build -f env/docker/service/Dockerfile \
	             -t $(IMAGE):$(VERSION) \
	             -t $(IMAGE):latest \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-publish
docker-publish: docker-build docker-push

.PHONY: docker-push
docker-push:
	@docker push $(IMAGE):$(VERSION)
	@docker push $(IMAGE):latest

.PHONY: docker-run
docker-run:
	@docker run --rm -it \
	            --name $(subst /,-,$(IMAGE))-dev \
	            -p 8080:8080 \
	            -p 8090:8090 \
	            -p 8091:8091 \
	            -p 8092:8092 \
	            $(IMAGE):$(VERSION) run --with-monitoring --with-profiling
