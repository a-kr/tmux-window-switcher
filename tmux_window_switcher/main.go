package main

import (
	"strconv"
	"fmt"
	"log"
	"os"
	"syscall"
	"bufio"
	"strings"
	"github.com/benizi/termstate"
)

const (
	Hotkeys = "0123456789abcdefghijklmnopqrstuvwxyz"
	MenuItem = "\033[1;33m%s\033[0m %s\n"
)


func main() {
	numberToHotkey := make(map[int]string)
	for i, hk := range Hotkeys {
		numberToHotkey[i] = string(hk)
	}
	hotkeyToItem := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) > 2 {
			panic("wtf")
		}
		hotkey := ""
		ok := false
		if len(parts) == 2 {
			num, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			hotkey, ok = numberToHotkey[num]
			if !ok {
				log.Fatalf("Cannot assign a hotkey to number %d", num)
			}
			hotkeyToItem[hotkey] = line
		}
		fmt.Fprintf(os.Stderr, MenuItem, hotkey, line)
	}

	stdin, err := os.OpenFile("/dev/tty", syscall.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("Cannot open /dev/tty: %s", err)
	}

	if !termstate.IsSupported() {
		log.Fatalf("cbreak is not supported")
	}
	initial, err := termstate.GetFD(int(stdin.Fd()))
	if err != nil {
		log.Fatalf("Cannot get termstate: %s", err)
	}
	initial.Cbreak().SetFD(int(stdin.Fd()))
	defer termstate.DeferredReset(termstate.State.Cbreak, termstate.State.EchoOff)

	for {
		buf := []byte{0}
		_, err := stdin.Read(buf)
		if err != nil {
			log.Fatal("Reading from stdin failed: %s", err)
		}

		key := string(buf[0])

		if line, ok := hotkeyToItem[key]; ok {
			fmt.Printf("%s\n", line)
			return
		}
		if buf[0] == 'q' {
			return
		}
	}

}
