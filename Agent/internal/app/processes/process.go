package processes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-ps"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	User   string
	Host   string
	Id_rsa string
	Script string
}

type Resp struct {
	StatusCode int
	Body       string
}

type List struct {
	Procs []Procss
}

type Procss struct {
	Pid        int
	PPid       int
	Executable string
}

func GetList() []byte {
	var prcs []Procss

	processList, err := ps.Processes()
	if err != nil {
		log.Println(err)
	}

	for _, p := range processList {
		prcs = append(prcs, Procss{
			Pid:        p.Pid(),
			PPid:       p.PPid(),
			Executable: p.Executable(),
		})
	}

	jsn, _ := json.Marshal(List{
		Procs: prcs,
	})

	return jsn
}

func KillProcess(pid string) []byte {
	conf := ParseConf("./configs/ssh.json")

	rsa, err := GetKeyFile(conf.Id_rsa)
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
		jsn, _ := json.Marshal(Resp{
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

	// Finally, run the command
	err = session.Run("sudo " + conf.Script + " " + pid)
	if err != nil {
		log.Println(err)
		jsn, _ := json.Marshal(Resp{
			StatusCode: 2,
			Body:       fmt.Sprintln(err),
		})
		return jsn
	}

	jsn, _ := json.Marshal(Resp{
		StatusCode: 0,
		Body:       "Process killed",
	})

	return jsn
}

// ParseConf return an Config type after parsing the specified configuration file
func ParseConf(path string) (conf Config) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("ERROR: CONFIGURATIONS FILES MISSING")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(bytes, &conf)

	return conf
}

// GetKeyFile will return a ssh.signer from the parsed file
func GetKeyFile(path string) (key ssh.Signer, err error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}
