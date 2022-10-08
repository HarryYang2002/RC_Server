SET DOMAIN=%1
SET PROTO_PATH=.\%DOMAIN%\api
SET GO_OUT_PATH=.\%DOMAIN%\api\gen\v1
mkdir %GO_OUT_PATH%
protoc -I=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% %DOMAIN%.proto
protoc -I=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%\%DOMAIN%.yaml:%GO_OUT_PATH% %DOMAIN%.proto

SET PBTS_BIN_DIR=..\wx\miniprogram\node_modules\.bin
SET PBTS_OUT_DIR=..\wx\miniprogram\service\proto_gen\%DOMAIN%
mkdir %PBTS_OUT_DIR%
call genPbjs.bat
echo import * as $protobuf from "protobufjs"; > %PBTS_OUT_DIR%\%DOMAIN%_pb.js
type %PBTS_OUT_DIR%\%DOMAIN%_pb_tmp.js >> %PBTS_OUT_DIR%\%DOMAIN%_pb.js
del %PBTS_OUT_DIR%\%DOMAIN%_pb_tmp.js
%PBTS_BIN_DIR%\pbts -o %PBTS_OUT_DIR%\%DOMAIN%_pb.d.ts %PBTS_OUT_DIR%\%DOMAIN%_pb.js
