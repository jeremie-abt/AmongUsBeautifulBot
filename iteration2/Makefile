unittestname = $(TESTNAME)
ifneq ($(unittestname),)
	testextrafiles = test/mock.go
else
	testextrafiles = 
endif

run: transformInMainPackage
	@go run *.go

unittest: transformInPackage
	go test test/*$(unittestname)*.go $(testextrafiles) -v

build: transformInMainPackage
	@sed -i "s/package\ main/package\ bot/g" main.go

transformInMainPackage:
	@sed -i "s/package\ bot/package\ main/g" *.go

transformInPackage:
	@sed -i "s/package\ main/package\ bot/g" *.go
