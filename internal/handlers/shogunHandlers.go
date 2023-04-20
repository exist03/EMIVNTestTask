package handlers

func initShogunHandlers(command []string) string {
	if command[0] == "show" {

	} else if command[0] == "describe" {
		return "decribe"
	}
	return "Wrong message"
}
