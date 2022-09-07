package commands

import "filec2/Handlers/sessions"

var (
	Commands = make(map[string]*Command)
)

//Command can be executed by a opperator via the terminal
type Command struct {
	Name string

	Admin bool

	Execute func(session *sessions.Session, cmd []string) error
}

//Register adds the command to the handler
func Register(cmd *Command) {
	if _, ok := Commands[cmd.Name]; ok == true {
		panic("more than one command with the same name")
	}

	Commands[cmd.Name] = cmd
}

//Get returns a command via it's name
func Get(name string) *Command {
	cmd := Commands[name]
	return cmd
}

//HasPermission returns true if the current user is allowed to run this command
func (c *Command) HasPermission(session *sessions.Session) bool {
	if c.Admin == true && session.User.Admin == false {
		return false
	}

	return true
}
