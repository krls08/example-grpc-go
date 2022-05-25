# example-grpc-go

## generate proto files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/infra/grpc/proto/chat.proto

### using relative paths
navigate to proto folder and run:
protoc --go_out=./chatpb_contract/ --go_opt=paths=source_relative --go-grpc_out=./chatpb_contract/ --go-grpc_opt=paths=source_relative chat.proto

### using absolute paths
from the root of project:
protoc --go_out=./internal/infra/grpc/proto/ --go_opt=paths=import --go-grpc_out=./internal/infra/grpc/proto/ --go-grpc_opt=paths=import internal/infra/grpc/proto/chat.proto
