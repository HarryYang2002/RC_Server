package mongotesting

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

var mongoURI string

const defaultMongoURI = "mongodb://localhost:27017"

func RunWithMongoInDocker(m *testing.M) int {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(any(err))
	}

	ctx := context.Background()

	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", //填写0会自动开启一共随机的端口
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(any(err))
	}
	containerID := resp.ID
	defer func() {
		err := c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(any(err))
		}

	}()

	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(any(err))
	}
	fmt.Println("container started")
	time.Sleep(5 * time.Second)

	inspRes, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(any(err))
	}
	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)

	return m.Run()
}

func NewClient(c context.Context) (*mongo.Client, error) {
	if mongoURI == "" {
		return nil, fmt.Errorf("mong uri not set. Please run RunWithMongoInDocker in TestMain")
	}
	return mongo.Connect(c, options.Client().ApplyURI(mongoURI))
}

func NewDefaultClient(c context.Context) (*mongo.Client, error) {
	return mongo.Connect(c, options.Client().ApplyURI(defaultMongoURI))
}

func SetupIndexes(c context.Context, d *mongo.Database) error {
	_, err := d.Collection("account").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "open_id",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	_, err = d.Collection("trip").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{
				Key: "trip.accountid", Value: 1,
			},
			{
				Key: "trip.status", Value: 1,
			},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
			"trip.status": 1,
		}),
	})
	if err != nil {
		return err
	}

	_, err = d.Collection("profile").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{Key: "accountid", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	return err
}
