package ping

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/aTTiny73/ThreadPoolmService/pkg/pool"

	"github.com/aTTiny73/ThreadPoolmService/internal/mail"
)

// Pinger function pings the hosts ip addresses and send email if there is no response
func Pinger(pingData []byte) {

	data := pool.Host{}
	err := json.Unmarshal(pingData, &data)
	if err != nil {
		fmt.Println(err)
	}

	//Ping syscall, -c ping count, -i interval, -w timeout
	out, _ := exec.Command("ping", data.IP, "-c 5", "-i 3", "-w 10").Output()
	if (strings.Contains(string(out), "Destination Host Unreachable")) || (strings.Contains(string(out), "100% packet loss")) {

		fmt.Println("Host not reachable")
		fmt.Println(data.IP, " ", data.Recipients)

		pool.CoordinatorInstance.Enqueue(mail.Mailer, pingData)

	} else {
		fmt.Println("Host ping successful")
	}

}
