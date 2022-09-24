SET PROTO_PATH=.\auth\api
SET GO_OUT_PATH=.\auth\api\gen\v1

protoc -I=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% auth.proto
protoc -I=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%\auth.yaml:%GO_OUT_PATH% auth.proto

SET PBTS_BIN_DIR=..\wx\miniprogram\node_modules\.bin
SET PBTS_OUT_DIR=..\wx\miniprogram\service\proto_gen\auth

call genPbjs.bat
echo import * as $protobuf from "protobufjs"; > %PBTS_OUT_DIR%\auth_pb.js
type %PBTS_OUT_DIR%\auth_pb_tmp.js >> %PBTS_OUT_DIR%\auth_pb.js
del %PBTS_OUT_DIR%\auth_pb_tmp.js
%PBTS_BIN_DIR%\pbts -o %PBTS_OUT_DIR%\auth_pb.d.ts %PBTS_OUT_DIR%\auth_pb.js