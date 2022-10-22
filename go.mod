module server

go 1.19

require (
	github.com/docker/docker v20.10.18+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/medivhzhan/weapp/v2 v2.5.0
	go.mongodb.org/mongo-driver v1.10.2
	go.uber.org/zap v1.23.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mozillazg/go-httpheader v0.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aa v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/acp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/advisor v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/af v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/afc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ams v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apcas v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ape v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/api v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asw v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ba v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/batch v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bda v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bi v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bizlive v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bma v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmeip v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmlb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmvpc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bpaas v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bri v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bsca v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/btoe v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/casb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cat v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ccc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cds v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfg v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfw v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ciam v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cii v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cim v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cis v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudhsm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudstudio v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cme v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cmq v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cpdp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cwp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cws v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dasb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dataintegration v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dayu v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbdc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dcdb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dlc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/drm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ds v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dtf v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eiam v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eis v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/fmu v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gme v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gpm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gs v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gse v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/habo v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ic v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/icr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ie v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iecp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iir v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iot v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotcloud v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotexplorer v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iottid v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideoindustry v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/irp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ivld v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lcic v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lowcode v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/market v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/memcached v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mgobe v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mmps v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mna v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mrs v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/msp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mvj v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/npp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/oceanus v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/partners v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pds v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rce v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/region v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rkp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smh v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smpn v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/solar v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssa v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sslpod v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taf v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tan v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tav v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbaas v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcb v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcbr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcex v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcss v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdcpg v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tds v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tia v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ticm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tics v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiems v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tione v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiw v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tkgdq v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trp v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsf v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsw v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ump v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vms v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wav v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wedata v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wss v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yinsuda v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/youmall v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunjing v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunsou v1.0.519 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/zj v1.0.519 // indirect
	github.com/tencentyun/cos-go-sdk-v5 v0.7.39 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gotest.tools/v3 v3.3.0 // indirect
)
