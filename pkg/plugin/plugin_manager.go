package plugin

var plg Plugin = nil

func RegisterPlugin(p Plugin) {
	plg = p
}

func GetPlugin() Plugin {
	return plg
}
