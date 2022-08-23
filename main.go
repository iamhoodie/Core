package core

import (
	"fmt"
)

// Agent must be implemented by the CORE code running.
type Agent interface {
	Start()
	Stop()
	Restart()
}

type Core struct {
	Manager string
	CmdBus  chan AgentCmd
	Agents  []Agent
}

type AgentCmd struct {
	ID  int64
	Cmd string
}

func New(agents []Agent, busSize int) *Core {
	return &Core{
		Manager: "",
		Agents:  agents,
		CmdBus:  make(chan AgentCmd, busSize),
	}
}

func (c *Core) InitializeAndListen() {
	//var d []Agent
	//
	//// Get a list of all the Agents and append them to be spun up. How do we want to keep track of them...by name or index?
	//d = append(d, &dm.DM{})
	//
	//coreManager := NewCore(agents)

	// Start all the Agents initially
	for _, agent := range c.Agents {
		agent.Start()
	}

	// Sit here and wait for commands
	for {
		cmd := <-c.CmdBus

		switch cmd.Cmd {
		case "START":
			c.Agents[cmd.ID].Start()
			break
		case "STOP":
			c.Agents[cmd.ID].Stop()
			break
		case "RESTART":
			c.Agents[cmd.ID].Restart()
			break
		default:
			fmt.Println("Unknown Command")
		}

	}
}
