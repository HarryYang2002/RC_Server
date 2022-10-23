SET DOMAIN=%1
SET PROTO_PATH=.\%DOMAIN%\api
SET GO_OUT_PATH=.\%DOMAIN%\api\gen\v1
mkdir %GO_OUT_PATH%
protoc -I=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% %DOMAIN%.proto