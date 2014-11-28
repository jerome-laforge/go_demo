package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var nbCon, unicPort int32 = 0, 14000

func main() {
	nbGoroutine := 500
	wg.Add(nbGoroutine)
	for ; nbGoroutine != 0; nbGoroutine-- {
		go openSocketReadOrTimeout()
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\npid :", os.Getpid())

	execThis(fmt.Sprintf("netstat -anpe | grep %d | wc -l", os.Getpid()), true)
	execThis(fmt.Sprintf("cat /proc/%d/status | grep Threads", os.Getpid()), true)
	cmd := `echo "Hello World" | nc -uc 127.0.0.1 ` + strconv.Itoa(int(unicPort))
	fmt.Println(cmd)
	time.Sleep(10 * time.Second)
	execThis(cmd, false)

	wg.Wait()
}

func openSocketReadOrTimeout() {
	port := atomic.AddInt32(&unicPort, 1)
	ip := &net.UDPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(port),
	}
	con, err := net.ListenUDP("udp", ip)

	if err == nil {
		defer func() {
			if con.Close() == nil {
				fmt.Printf("Connexion est fermée (port = %d). Nombre de connexions restantes ouvertes : %d\n", port, atomic.AddInt32(&nbCon, -1))
			}
		}()

		fmt.Printf("Connexion est ouverte (port = %d). Nombre de connexions actuellement ouvertes : %d\n", port, atomic.AddInt32(&nbCon, 1))

		buf := make([]byte, 1500)
		con.SetDeadline(time.Now().Add(30 * time.Second))
		n, _, err := con.ReadFromUDP(buf)
		if err == nil {
			fmt.Printf("buf[0:%d] = %s", n, buf[0:n])
		}
	} else {
		fmt.Println("Erreur sur l'ouverture de la connexion", err)
	}
	wg.Done()
}

func execThis(cmd string, print bool) {
	out, _ := exec.Command("sh", "-c", cmd).Output()
	if print {
		fmt.Printf("%s => %s", cmd, out)
	}
}
