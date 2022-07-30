package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

// TODO
// SET TIMEOUTS FOR CONNECTIONS IN NAMED PIPES
// LOOK INTO CHANGING NAMED PIPES TO UNIX DOMAIN SOCKETS
// MAKE PROCESS TO A SERVICE
// MAKE IT MULTIPLATFORM
// MAYBE ADD INDEXING

// https://cli.urfave.org/v2/#subcommands
// https://unix.stackexchange.com/questions/426862/proper-way-to-run-shell-script-as-a-daemon
// https://opensource.com/article/19/4/interprocess-communication-linux-networking
// https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/
// https://www.mtholyoke.edu/courses/dstrahma/cs322/ipc.htm
// https://www.educative.io/answers/what-are-unix-domain-sockets
// https://github.com/devlights/go-unix-domain-socket-example/tree/master/cmd/basic
// https://gist.githubusercontent.com/hakobe/6f70d69b8c5243117787fd488ae7fbf2/raw/75f02abc7742228b24842cecee51837da858055d/client.go
// https://stackoverflow.com/questions/65799886/unable-to-read-from-unix-socket-using-net-conn-read
// https://apple.stackexchange.com/questions/364094/how-to-view-status-of-service-e-g-whether-its-running-in-a-format-similar-to

// Struct for process communication
type response struct {
	Rtn string `json:"rtn"`
	Ok  bool   `json:"ok"`
}

const INPUT_PATH string = "./pipes/input"
const OUTPUT_PATH string = "./pipes/output"

func sendMessage(s ...string) (err error) {
	f, err := os.OpenFile(INPUT_PATH, os.O_WRONLY, 0)
	if err != nil {
		return
	}
	encoded, err := json.Marshal(s)
	if err != nil {
		return
	}
	_, err = f.Write(encoded)
	if err != nil {
		return
	}
	err = f.Close()
	return
}

// Waits for successful response on clipboard server
func waitForResponse() response {
	encoded, _ := os.ReadFile(OUTPUT_PATH)
	var decoded response
	json.Unmarshal(encoded, &decoded)
	return decoded
}

// Copies string to system's clipboard
func copyToClipboard(s string) (err error) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(s)
	err = cmd.Run()
	return
}

// Adds input to clipboard stack
func add(c *cli.Context) (err error) {
	err = sendMessage("add", c.Args().First())
	if err != nil {
		return
	}
	response := waitForResponse()
	if !response.Ok {
		err = errors.New(response.Rtn)
	}
	return
}

// Lists all elements in clipboard stack
func list(c *cli.Context) (err error) {
	err = sendMessage("list")
	if err != nil {
		return
	}
	response := waitForResponse()
	if !response.Ok {
		err = errors.New(response.Rtn)
	} else {
		fmt.Print(response.Rtn)
	}
	return
}

// Pops the top element of the clipboard, copying it to the system's clipboard
func pop(c *cli.Context) (err error) {
	err = sendMessage("pop")
	if err != nil {
		return
	}
	response := waitForResponse()
	if !response.Ok {
		err = errors.New(response.Rtn)
		return
	}
	err = copyToClipboard(response.Rtn)
	if err != nil {
		return
	}
	fmt.Println(response.Rtn, "copied to clipboard!")
	return
}

// Copies the top element on the clipboard stack to the system's clipboard *without popping it*
func front(c *cli.Context) (err error) {
	err = sendMessage("front")
	if err != nil {
		return
	}
	response := waitForResponse()
	if !response.Ok {
		err = errors.New(response.Rtn)
		return
	}
	err = copyToClipboard(response.Rtn)
	if err != nil {
		return
	}
	fmt.Println(response.Rtn, "copied to clipboard!")
	return nil
}

// Clears the elements from the clipboard.
func clear(c *cli.Context) (err error) {
	err = sendMessage("clear")
	if err != nil {
		return
	}
	waitForResponse()
	return
}

//Updates size limit for the clipboard stack (Default=5). Set to 0 for no limit (Not recommended).
func setLimit(c *cli.Context) (err error) {
	err = sendMessage("setLimit", c.Args().First())
	if err != nil {
		return
	}
	waitForResponse()
	return
}

func main() {
	app := &cli.App{
		Name:      "Tapiclipboard",
		Usage:     "Superset of system clipboard! 🐑❤️",
		UsageText: "tc COMMAND [arguments ...]",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Adds input to clipboard stack",
				Action:  add,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "|Lists all elements in clipboard stack",
				Action:  list,
			},
			{
				Name:    "pop",
				Aliases: []string{"p"},
				Usage:   "pop last text to the clipboard",
				Action:  pop,
			},
			{
				Name:    "front",
				Aliases: []string{"f"},
				Usage:   "get last text to the clipboard, without popping it",
				Action:  front,
			},
			{
				Name:    "clear",
				Aliases: []string{"c"},
				Usage:   "Clears the elements from the clipboard",
				Action:  clear,
			},
			{
				Name:    "limit",
				Aliases: []string{"li"},
				Usage:   "Updates size limit for the clipboard stack (Default=5)",
				Action:  setLimit,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
