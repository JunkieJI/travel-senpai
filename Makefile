BINARY=travelsenpaid
TESTDIRS=`go list ./... | grep -v /vendor/ `
BASEDIR=$(shell pwd)
GO_SWAGGER=./ops/tools/go-swagger/swagger_0.12.0_darwin_amd64

default:
	rm -rf build/
	mkdir build
	CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags -static" -o build/$(BINARY)

test:
	TRAVEL_SENPAI_CONFIGPATH=$(BASEDIR) TRAVEL_SENPAI_CONFIGFILE=config_test go test --cover -v $(TESTDIRS)


clean:
	rm -rf build/

swagger:
	rm -rf ./api/swagger/models
	rm -rf ./api/swagger/restapi/operations
	$(GO_SWAGGER) generate server --spec=./api/swagger/swagger.yml -t ./api/swagger --name=travel_senpai --exclude-main
