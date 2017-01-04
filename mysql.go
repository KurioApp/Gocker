package gocker

import (
	"testing"
	"time"
	"strconv"
)

const (
	containerImage = "mysql"
	mysqlPort = 33066
)

func SetupMysqlContainer(t *testing.T) (c ContainerID, ip string) {
	return SetupContainer(t, containerImage, mysqlPort, 10*time.Second, func() (string, error) {
		return Run("-d", "-p", strconv.Itoa(mysqlPort) + ":3306", "--name", "mysql-test", "-e", "MYSQL_ALLOW_EMPTY_PASSWORD=1",containerImage)
	})
}