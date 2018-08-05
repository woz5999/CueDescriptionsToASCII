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
