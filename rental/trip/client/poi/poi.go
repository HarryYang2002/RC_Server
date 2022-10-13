package poi

import (
	"context"
	"github.com/golang/protobuf/proto"
	"hash/fnv"
	rentalpb "server/rental/api/gen/v1"
)

var poi = []string{
	"天安门",
	"长城",
	"故宫",
	"天坛",
	"颐和园",
	"圆明园",
	"全聚德",
}

type Manager struct {
}

func (*Manager) Resolve(c context.Context, req *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}
	h := fnv.New32()
	h.Write(b)
	return poi[int(h.Sum32())%len(poi)], nil
}
