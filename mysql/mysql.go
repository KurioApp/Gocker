package mysql

import (
	"testing"
	"github.com/KurioApp/gocker"
	"time"
	"strconv"
)

const (
	containerImage = "mysql"
	mysqlPort = 33066
)

func SetupMysqlContainer(t *testing.T) (c gocker.ContainerID, ip string) {
	return gocker.SetupContainer(t, containerImage, mysqlPort, 10*time.Second, func() (string, error) {
		return gocker.Run("-d", "-p", strconv.Itoa(mysqlPort) + ":3306", "--name", "mysql-test", "-e", "MYSQL_ALLOW_EMPTY_PASSWORD=1",containerImage)
	})
}