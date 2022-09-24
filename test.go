package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"google.golang.org/protobuf/proto"
//	trippb "server/protoc/gen/go"
//)
//
//func main() {
//	fmt.Println("Hello World")
//	trip := trippb.Trip{
//		Start:       "abc",
//		End:         "def",
//		DurationSec: 3600,
//		FeeCent:     10000,
//		StartPos: &trippb.Location{
//			Latitude: 30,
//			Longitude: 120,
//		},
//		EndPos: &trippb.Location{
//			Latitude: 35,
//			Longitude: 115,
//		},
//		PathLocations: []*trippb.Location{
//			{
//				Latitude: 31,
//				Longitude: 119,
//			},
//			{
//				Latitude: 32,
//				Longitude: 118,
//			},
//		},
//	}
//	fmt.Println(&trip)
//
//	b, err := proto.Marshal(&trip)
//	if err != nil {
//		panic(any(err))
//	}
//	fmt.Printf("%X\n", b)
//
//	var trip2 trippb.Trip
//	err = proto.Unmarshal(b, &trip2)
//	if err!=nil{
//		panic(any(err))
//	}
//	fmt.Println(&trip2)
//
//	b, err = json.Marshal(&trip2)
//	if err != nil {
//		panic(any(err))
//	}
//	fmt.Printf("%s\n",b)
//}
