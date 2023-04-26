package details

import (
	"log"
	"net"
	"os"
)

func Hostname() (name string, err error) {
	hostname, err := os.Hostname()
	return hostname, err
}
func GetIpAddress() (net.IP) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
