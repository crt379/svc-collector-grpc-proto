.PHONY: tenant service svcapi svcapieg application appsvc register
all: tenant service svcapi svcapieg application appsvc register

.PHONY: tenant
tenant:
	protoc --go_out=. --go_opt=paths=source_relative .\tenant\tenant.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\tenant\tenant.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\tenant\tenant.proto
	protoc --openapiv2_out=. .\tenant\tenant.proto

.PHONY: service
service:
	protoc --go_out=. --go_opt=paths=source_relative .\service\service.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\service\service.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\service\service.proto
	protoc --openapiv2_out=. .\service\service.proto

.PHONY: svcapi
svcapi:
	protoc --go_out=. --go_opt=paths=source_relative .\svcapi\svcapi.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\svcapi\svcapi.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\svcapi\svcapi.proto
	protoc --openapiv2_out=. .\svcapi\svcapi.proto

.PHONY: svcapieg
svcapieg:
	protoc --go_out=. --go_opt=paths=source_relative .\svcapieg\svcapieg.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\svcapieg\svcapieg.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\svcapieg\svcapieg.proto
	protoc --openapiv2_out=. .\svcapieg\svcapieg.proto

.PHONY: application
application:
	protoc --go_out=. --go_opt=paths=source_relative .\application\application.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\application\application.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\application\application.proto
	protoc --openapiv2_out=. .\application\application.proto

.PHONY: appsvc
appsvc:
	protoc --go_out=. --go_opt=paths=source_relative .\appsvc\appsvc.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\appsvc\appsvc.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\appsvc\appsvc.proto
	protoc --openapiv2_out=. .\appsvc\appsvc.proto

.PHONY: register
register:
	protoc --go_out=. --go_opt=paths=source_relative .\register\register.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\register\register.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\register\register.proto
	protoc --openapiv2_out=. .\register\register.proto

.PHONY: processor
processor:
	protoc --go_out=. --go_opt=paths=source_relative .\processor\processor.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\processor\processor.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\processor\processor.proto
	protoc --openapiv2_out=. .\processor\processor.proto

.PHONY: appapi
appapi:
	protoc --go_out=. --go_opt=paths=source_relative .\appapi\appapi.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\appapi\appapi.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\appapi\appapi.proto
	protoc --openapiv2_out=. .\appapi\appapi.proto

.PHONY: appproc
appproc:
	protoc --go_out=. --go_opt=paths=source_relative .\appproc\appproc.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative .\appproc\appproc.proto
	protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative .\appproc\appproc.proto
	protoc --openapiv2_out=. .\appproc\appproc.proto