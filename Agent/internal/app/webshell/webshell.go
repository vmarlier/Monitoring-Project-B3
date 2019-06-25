package webshell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"cloud-opsme.ga/monitoring/agent/internal/app/processes"
	"golang.org/x/crypto/ssh"
)

// Output is the basic output for a well executed Command
type Output struct {
	StatusCode int
	Body       string
}

/*
Command allowed execution of:
cat, cp, date, df, dmesg, echo, grep, hostname, ip, journalctl, less, ls, mkdir,
mv, netstat, networkctl, ps, pwd, rm, rmdir, systemctl, touch, uname, which
*/
func Command(cmd string) []byte {
	conf := processes.ParseConf("./configs/ssh.json")

	rsa, err := processes.GetKeyFile(conf.Id_rsa)
	if err != nil {
		log.Println(err)
	}

	// Authentication
	config := &ssh.ClientConfig{
		User:            conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(rsa),
		},
	}

	// Connect
	client, err := ssh.Dial("tcp", conf.Host+":22", config)
	if err != nil {
		fmt.Println(err)
		jsn, _ := json.Marshal(Output{
			StatusCode: 1,
			Body:       "ERROR: SSH connection initialization failed",
		})
		return jsn
	}

	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output

	err = session.Run(cmd)
	if err != nil {
		log.Println(err)
		jsn, _ := json.Marshal(Output{
			StatusCode: 2,
			Body:       fmt.Sprintln(err),
		})
		return jsn
	}

	jsn, _ := json.Marshal(Output{
		StatusCode: 0,
		Body:       b.String(),
	})

	return jsn
}
