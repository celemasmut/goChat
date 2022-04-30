package main

import (
	"flag"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	host2 = flag.String("host", "scanme.nmap.org", "host to be scanned")
)

func checkEnviron() (string, error) {

	out, err := exec.Command("ping", "-c", "1", *host2).Output()
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`ttl=(.?).[\S]`)
	ttl := fmt.Sprintf("%s", re.FindString(string(out)))
	ttl = strings.Split(ttl, "=")[1]
	ttlNum, err := strconv.Atoi(ttl)
	if err != nil {
		return "", err
	}
	if ttlNum <= 64 {
		return "\n\t[+] Linux system\n", nil
	} else if ttlNum >= 127 {
		return "\n\t[+] Windows system\n", nil
	} else {
		return "\n\t[-] the time to the life of the target system doesn't exists\n", nil
	}

}

func main() {

	flag.Parse()

	environ, _ := checkEnviron()
	fmt.Println(environ)

	var wg sync.WaitGroup
	//wg.Add(1000)
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host2, j))

			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("[+] port %d open!!\n", j)
		}(i)
	}
	wg.Wait()
}
