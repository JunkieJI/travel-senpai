BINARY=travelsenpaid
TESTDIRS=`go list ./... | grep -v /vendor/ `
BASEDIR=$(shell pwd)

default:
	rm -rf build/
	mkdir build
	CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags -static" -o build/$(BINARY)

test:
	TRAVEL_SENPAI_CONFIGPATH=$(BASEDIR) TRAVEL_SENPAI_CONFIGFILE=config_test go test --cover -v $(TESTDIRS)


clean:
	rm -rf build/
