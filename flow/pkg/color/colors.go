package color

func Red(msg string) string {
	return "\033[31m" + msg + "\033[0m"
}

func Green(msg string) string {
	return "\033[32m" + msg + "\033[0m"
}

func Yellow(msg string) string {
	return "\033[33m" + msg + "\033[0m"
}

func Blue(msg string) string {
	return "\033[34m" + msg + "\033[0m"
}

func Magenta(msg string) string {
	return "\033[35m" + msg + "\033[0m"
}

func Cyan(msg string) string {
	return "\033[36m" + msg + "\033[0m"
}

func White(msg string) string {
	return "\033[37m" + msg + "\033[0m"
}
