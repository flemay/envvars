VERSION ?= 0.0.7
GIT_TAG = v$(VERSION)
DOCKER_TAG = $(VERSION)
IMAGE_NAME = flemay/envvars:$(VERSION)
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
ENVFILE ?= env.template

all:
	ENVFILE=env.example $(MAKE) envfile deps test build run buildDockerImage clean

ciTest: envfile deps test sendCoverage build run buildDockerImage clean

_ciRelease:
	TAG=$(GIT_TAG) ./scripts/github_release.sh

envfile:
	cp -f $(ENVFILE) .env

deps:
	docker-compose pull
	$(COMPOSE_RUN_GOLANG) ./scripts/deps.sh

mock:
	$(COMPOSE_RUN_GOLANG) ./scripts/mock.sh

test:
	$(COMPOSE_RUN_GOLANG) ./scripts/test.sh

sendCoverage: $(PROFILE_NAME)
	$(COMPOSE_RUN_GOLANG) ./scripts/coverage.sh

build:
	$(COMPOSE_RUN_GOLANG) bash -c 'VERSION=$(VERSION) ./scripts/build.sh'

run:
	$(COMPOSE_RUN_GOLANG) ./scripts/run.sh

buildDockerImage:
	docker build --no-cache -t $(IMAGE_NAME) .
	docker run --rm $(IMAGE_NAME) --help
	docker run --rm $(IMAGE_NAME) version

pushDockerImage:
	IMAGE_NAME=$(IMAGE_NAME) ./scripts/push.sh

removeDockerImage:
	docker rmi -f $(IMAGE_NAME)

tag:
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)

# this is to be used with caution
overwriteTag:
	-git tag -d $(GIT_TAG)
	-git push origin :refs/tags/$(GIT_TAG)
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)

clean:
	$(COMPOSE_RUN_GOLANG) ./scripts/clean.sh
	docker-compose down --remove-orphans
	-$(MAKE) removeDockerImage
	rm -f .env

shell:
	$(COMPOSE_RUN_GOLANG) bash