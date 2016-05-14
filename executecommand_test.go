package main

import (
    "testing"
    "github.com/stretchr/testify/assert" 
)

func TestAllOk(t *testing.T) {
    args := []string {"-l", "-t", "-r", "/tmp"}
    _, stderr, exitcode := execute("ls", args)
    assert := assert.New(t)

    assert.Empty(stderr)
    assert.Equal(0, exitcode)
}


func TestCmdNotFound(t *testing.T) {
    args := []string {"-l", "-t", "-r", "/tmp"}
    stdout, stderr, exitcode := execute("lswontbehere", args)
    assert := assert.New(t)

    assert.Empty(stdout)
    assert.Equal("executable file not found in $PATH", string(stderr)) //Linux only
    assert.Equal(1, exitcode)
}

func TestInvalidArgs(t *testing.T) {
    args := []string {"-l", "-t", "-r", "/tmpnothere"}
    stdout, stderr, exitcode := execute("ls", args)
    assert := assert.New(t)

    assert.Empty(stdout)
    assert.Equal("ls: cannot access '/tmpnothere': No such file or directory\n", string(stderr)) //Linux only
    assert.Equal(2, exitcode)
}
