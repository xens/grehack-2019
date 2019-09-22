package iputil

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"os/exec"
	"strings"
	"time"
)

func LookupAddr(ip net.IP) (string, error) {
	names, err := net.LookupAddr(ip.String())
	if err != nil || len(names) == 0 {
		return "", err
	}
	// Always return unrooted name
	return strings.TrimRight(names[0], "."), nil
}

func LookupPort(ip net.IP, port uint64) error {
	address := fmt.Sprintf("[%s]:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func ToDecimal(ip net.IP) *big.Int {
	i := big.NewInt(0)
	if to4 := ip.To4(); to4 != nil {
		i.SetBytes(to4)
	} else {
		i.SetBytes(ip)
	}
	return i
}

func IsItDown(address string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	pingCommand := fmt.Sprintf("ping -c 1 -W 150 %s", address)
	command := exec.CommandContext(ctx, "sh", "-c", pingCommand)
	trace, err := command.Output()
	return string(trace), err
}
