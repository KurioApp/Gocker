package mongo

import (
	"testing"
	"time"

	"github.com/KurioApp/gocker"
)

const (
	mongoImage = "mongo"
	mongoPort  = "27027"
)

func SetupMongoContainer(t *testing.T) (c gocker.ContainerID, ip string) {
	return gocker.SetupContainer(t, mongoImage, 27027, 10*time.Second, func() (string, error) {
		return gocker.Run("-d", "-p", mongoPort+":27017", "--name", "dockertestmongodb", mongoImage)
	})
}
