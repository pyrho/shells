package main

import (
	// UNIX "golang.org/x/sys/unix"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func writeToStdin(s string) {
	var eno syscall.Errno
	for _, c := range s {
		_, _, eno = syscall.RawSyscall(
			syscall.SYS_IOCTL,
			os.Stdin.Fd(),
			syscall.TIOCSTI,
			uintptr(unsafe.Pointer(&c)))
		if eno != 0 {
			log.Fatalln(eno)
		}
	}

}

///////////////////////////////////////////////////////////////////////////////*
/*
        ┌──────────────────────────────────────────────────────────────┐
        │░█░█░▀█▀░▀█▀░█░█░░░▀█▀░█░█░█▀▀░░░█░█░█▀▀░█░░░█▀█░░░█▀█░█▀▀░░░░│
        │░█▄█░░█░░░█░░█▀█░░░░█░░█▀█░█▀▀░░░█▀█░█▀▀░█░░░█▀▀░░░█░█░█▀▀░░▀░│
        │░▀░▀░▀▀▀░░▀░░▀░▀░░░░▀░░▀░▀░▀▀▀░░░▀░▀░▀▀▀░▀▀▀░▀░░░░░▀▀▀░▀░░░░▀░│
        └──────────────────────────────────────────────────────────────┘
- https://www.reddit.com/r/linuxquestions/comments/td29k5/is_there_a_way_to_have_a_program_exit_and_place/
- https://stackoverflow.com/questions/54388088/how-to-ioctl-properly-from-golang
- https://www.reddit.com/r/golang/comments/d3cu3l/writing_to_terminal_input_buffer/
*/
