package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Handler interface {
	Handler(...string)
	Validate(...string) error
	Description() string
	Example() string
}

type Cli struct {
	commands    map[string]Handler
	commandKeys []string
	done        chan struct{}
	cmds        chan []string
	rdr         io.Reader
}

func NewCli() *Cli {
	return NewCliWithReader(nil)
}

func NewCliWithReader(r io.Reader) *Cli {
	if r == nil {
		r = os.Stdin
	}

	c := &Cli{
		commands:    map[string]Handler{},
		commandKeys: make([]string, 0),
		done:        make(chan struct{}, 1),
		cmds:        make(chan []string),
		rdr:         r,
	}

	go func() {
		csignal := make(chan os.Signal, 1)

		signal.Notify(csignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-csignal

		c.Exit()
	}()

	return c
}

func (c *Cli) Reader() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	txt := scanner.Text()
	// 	if txt == "" {
	// 		continue
	// 	}

	// 	cmds, err := BuildCommand(txt)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		continue
	// 	}

	// 	c.cmds <- cmds
	// }

	reader := bufio.NewReader(c.rdr)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			continue
		}

		cmds, _ := BuildCommand(text)

		c.cmds <- cmds
	}
}

func (c *Cli) Start() {
	go c.Reader()

	for {
		select {
		case cmd := <-c.cmds:
			if len(cmd) > 0 {
				c.Handle(cmd[0], cmd[1:]...)
			}
		case <-c.done:
			return
		}
	}
}

func (c *Cli) Handle(command string, params ...string) {
	cmd := strings.ToLower(command)

	if cmd == "exit" {
		c.Exit()
		return
	}

	if cmd == "help" {
		c.Help()
		return
	}

	exe, ok := c.commands[cmd]
	if !ok {
		fmt.Println(fmt.Errorf("err: Command %s is not found", command))
		return
	}

	err := exe.Validate(params...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	exe.Handler(params...)
}

func (c *Cli) Register(cmd string, handler Handler) {
	cmdLower := strings.ToLower(cmd)
	if cmdLower == "help" || cmdLower == "exit" {
		return
	}

	c.commandKeys = append(c.commandKeys, cmd)
	c.commands[strings.ToLower(cmd)] = handler
}

func (c *Cli) Help() {
	bufStr := new(strings.Builder)
	bufStr.WriteString("Command | Description")
	bufStr.WriteString("\n")

	for _, k := range c.commandKeys {
		cmd := c.commands[k]

		bufStr.WriteString(cmd.Example())
		bufStr.WriteString(" | ")
		bufStr.WriteString(cmd.Description())
		bufStr.WriteString("\n")
	}
	bufStr.WriteString("help | show command list")
	bufStr.WriteString("\n")
	bufStr.WriteString("exit | exit quiz master")
	bufStr.WriteString("\n")

	fmt.Println(bufStr.String())
}

func (c *Cli) Exit() {
	c.done <- struct{}{}
	close(c.cmds)
}
