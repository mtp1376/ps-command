package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPpid(pid string) string {
	statFileAddress := fmt.Sprintf("/proc/%s/stat", pid)
	content, _ := os.ReadFile(statFileAddress)
	ppid := strings.Fields(string(content))[3]
	if ppid != "0" {
		return ppid
	}
	return ""
}

func main() {
	entries, _ := os.ReadDir("/proc/")
	fmt.Printf("PID\tCMD\tPPID\n") // Header row
	for _, entry := range entries {
		if _, err := strconv.Atoi(entry.Name()); err != nil || !entry.IsDir() {
			continue
		}
		pid := entry.Name()
		// /proc/{pid}/cmdline
		cmdlineAddress := fmt.Sprintf("/proc/%s/cmdline", pid)
		content, _ := os.ReadFile(cmdlineAddress)
		cleanCmdLine := strings.Replace(string(content), "\000", " ", -1)
		ppid := getPpid(pid)
		fmt.Printf("%s\t%s\t%s\n", pid, cleanCmdLine, ppid)
	}
}
