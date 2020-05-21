BUILD = GOOS=js GOARCH=wasm GOPATH=${PWD}/../..:$(GOPATH) go build

all:
	${BUILD} -o contract_audit.wasm audit/audit.go
	${BUILD} -o contract_face.wasm face/service.go
	${BUILD} -o contract_identify.wasm identify/service.go
	${BUILD} -o contract_location.wasm location/service.go
	${BUILD} -o contract_time.wasm time/service.go
