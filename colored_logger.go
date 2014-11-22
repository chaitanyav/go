package main

import (
	"log"
	"os"
)

func fatal(msg string) {
	log.Fatal("\033[1;91;40mFATAL:\033[0m ", msg)
}

func error(msg string) {
	log.Println("\033[1;35;40mERROR:\033[0m ", msg)
}

func debug(msg string) {
	log.Println("\033[1;37;40mDEBUG:\033[0m ", msg)
}

func warn(msg string) {
	log.Println("\033[1;33;40mWARN:\033[0m ", msg)
}

func info(msg string) {
	log.Println("\033[1;36;40mINFO:\033[0m ", msg)
}

func main() {
	log.SetOutput(os.Stdout)
	info("This is a info message")
	warn("This is a warn message")
	debug("This is a debug message")
	error("This is a error message")
	fatal("This is a error message")
}
