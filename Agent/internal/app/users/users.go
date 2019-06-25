package users

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"cloud-opsme.ga/monitoring/agent/internal/app/processes"
	"golang.org/x/crypto/ssh"
)

type ConfigSudoer struct {
	Group string
}

type Config struct {
	User   string
	Host   string
	Id_rsa string
	Script string
}

type User struct {
	Username    string
	HomeDir     string
	UserID      string
	GroupID     string
	DisplayName string
	Sudoer      bool
}

type ReturnList struct {
	StatusCode int
	Body       []User
}

type Response struct {
	StatusCode int
}

func RemoveSudoer(uid string) []byte {
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
		jsn, _ := json.Marshal(Response{
			StatusCode: 1,
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

	confS := parseConf("./configs/sudoers.json")
	usr, _ := user.LookupId(uid)

	err = session.Run("sudo /usr/bin/gpasswd -d " + usr.Username + " " + confS.Group)
	if err != nil {
		log.Println(err)
		jsn, _ := json.Marshal(Response{
			StatusCode: 2,
		})
		return jsn
	}

	jsn, _ := json.Marshal(Response{
		StatusCode: 0,
	})

	return jsn
}

func AddSudoer(uid string) []byte {
	conf := processes.ParseConf("./configs/ssh.json")

	rsa, err := processes.GetKeyFile(conf.Id_rsa)
	if err != nil {
		log.Println("rsa", err)
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
		fmt.Println("client", err)
		jsn, _ := json.Marshal(Response{
			StatusCode: 1,
		})
		return jsn
	}

	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		log.Println("session", err)
	}
	defer session.Close()

	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output

	confS := parseConf("./configs/sudoers.json")
	usr, _ := user.LookupId(uid)

	err = session.Run("sudo /usr/bin/gpasswd -a " + usr.Username + " " + confS.Group)
	if err != nil {
		log.Println("run", err)
		jsn, _ := json.Marshal(Response{
			StatusCode: 2,
		})
		return jsn
	}

	jsn, _ := json.Marshal(Response{
		StatusCode: 0,
	})

	return jsn
}

// GetList will return the users list from /etc/passwd
func GetList() []byte {
	var Users []string
	var users []User

	file, err := os.Open("/etc/passwd")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		// skip all line starting with #
		if equal := strings.Index(line, "#"); equal < 0 {
			// get the username and description
			lineSlice := strings.FieldsFunc(line, func(divide rune) bool {
				return divide == ':' // we divide at colon
			})

			if len(lineSlice) > 0 {
				Users = append(Users, lineSlice[0])
			}

		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	conf := parseConf("./configs/sudoers.json")

	// now we have a list of users
	// iterate(cycle) each of them to
	// print out HomeDir, GroupID, description, etc

	for _, name := range Users {

		usr, err := user.Lookup(name)
		if err != nil {
			panic(err)
		}

		users = append(users, User{
			Username:    usr.Username,
			HomeDir:     usr.HomeDir,
			UserID:      usr.Uid,
			GroupID:     usr.Gid,
			DisplayName: usr.Name,
			Sudoer:      isSudoer(usr.Username, conf.Group),
		})
	}

	jsn, _ := json.Marshal(ReturnList{
		StatusCode: 0,
		Body:       users,
	})

	return jsn
}

func isSudoer(username string, adminGrp string) bool {
	usr, err := user.Lookup(username)
	if err != nil {
		log.Println(err)
	}

	grp, err := user.LookupGroup(adminGrp)
	if err != nil {
		log.Println(err)
	}

	usrgrps, _ := usr.GroupIds()

	for _, group := range usrgrps {
		if group == grp.Gid {
			return true
		}
	}

	return false
}

func parseConf(path string) (conf ConfigSudoer) {
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
