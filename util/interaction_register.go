package util

type InteractionRegister interface {
	Command(Command)
	Component(string, ComponentHandle)
	Autocomplete(Command, AutocompleteHandle)
	Modal(string, ModalHandle)
	// TODO: Add specific as Button()
}

// Add Command handler to Manager
func (m *Imanager) Command(cmd Command) {
	defer func() {
		if err := recover(); err != nil {
			m.Logger.Error("command failed to Init()", "command", cmd, "err", err)
		} else {
			m.interactions = append(m.interactions, cmd)
		}

	}()
	cmd.Init(m)
}

func (m *Imanager) AddCommandsFromPackage(cmds []Command) {
	for _, c := range cmds {
		c.Init(m)
	}
	m.interactions = cmds
}

func (m *Imanager) Component(customId string, handle ComponentHandle) {
	if customId != "" {
		m.components = append(m.components, Component{customId, handle})
	}
}

func (m *Imanager) Autocomplete(c Command, f AutocompleteHandle) {
	m.autocompletes = append(m.autocompletes, Autocomplete{c, f})
}

func (m *Imanager) Modal(_ string, _ ModalHandle) {
	panic("not implemented") // TODO: Implement
}