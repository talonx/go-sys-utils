package main

import (
	"os/exec"
    "syscall"
    "fmt"
    "bytes"
    "reflect"
)

func execute(binaryName string, args []string) ([]byte, []byte, int) {
    cmd := exec.Command(binaryName, args...) 

    stdout := &bytes.Buffer {}
    stderr := &bytes.Buffer {}
    cmd.Stdout = stdout
    cmd.Stderr = stderr
    exitcode := 1

    //TODO control over environment

    /*
    * The command might not have been found (exec.go), the command might have
    * exited with an error code or there might be a problem with IO.
    */
    err := cmd.Run()

    //TODO handle non *nix
    if err != nil {
        switch err.(type) {
            case *exec.ExitError:
                e := err.(*exec.ExitError)
                if status, ok := e.Sys().(syscall.WaitStatus); ok {
                    exitcode = status.ExitStatus()
                }
            case *exec.Error:
                e := err.(*exec.Error)
                stderr.WriteString(e.Err.Error())
            default:
                panic("Unknown err type: " + reflect.TypeOf(err).String())
        }
    } else {
        if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
            exitcode = status.ExitStatus()
        }
    }

    return stdout.Bytes(), stderr.Bytes(), exitcode
}

func main() {
    args := []string {"-l", "-t", "-r", "/asd"}
	stdout, stderr, exitcode := execute("lsa", args)
    fmt.Println(string(stdout))
    fmt.Println(string(stderr))
    fmt.Println(exitcode)

}
