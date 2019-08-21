NAMESPACE  := woz5999
REPOSITORY := cuedescriptionstoascii

all:
	docker build --rm -t $(NAMESPACE)/$(REPOSITORY):latest .

linux: | local
darwin: | local
local:
ifneq ($(MAKECMDGOALS), darwin)
ifneq ($(MAKECMDGOALS), linux)
	$(error Valid local build targets are "linux" and "darwin")
endif
endif
	GOOS=$(MAKECMDGOALS) GOARCH=amd64 CGO_ENABLED=0 go build -o ./$(REPOSITORY)

deploy:
	gcloud config set project cue-converter
	gcloud app deploy --quiet app.yaml

run: all
	docker rm -f converter || true
	docker run -d -p 8080:8080 --name converter woz5999/cuedescriptionstoascii
	@echo Access app on http://localhost:8080
	docker logs converter -f
