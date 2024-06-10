package shx

import (
	"bytes"
	"os/exec"
	"strings"
)

var defaultShell = "/bin/sh"

func FD_(args string) ([]string, error) {
	lst, err := stringarray("fd --print0 "+args, 0x00)
	if err != nil {
		return nil, err
	}
	return removelastemptystring(lst), nil
}

func FD(args string) []string {
	return must1(FD_(args))
}

func RG_(args string) ([]string, error) {
	lst, err := stringarray("rg --no-config --no-filename "+args, '\n')
	return removelastemptystring(lst), err
}

func RG(args string) []string {
	return must1(RG_(args))
}

func RGFile_(args string) ([]string, error) {
	lst, err := stringarray("rg --no-config --files-with-matches "+args, '\n')
	return removelastemptystring(lst), err
}

func RGFile(args string) []string {
	return must1(RGFile_(args))
}

func stringarray(command string, seperator byte) ([]string, error) {
	cmd := exec.Command(defaultShell, "-c", command)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lst := strings.Split(buf.String(), string(seperator))
	return lst, nil
}

func removelastemptystring(lst []string) []string {
	if len(lst) == 0 {
		return lst
	}

	if lst[len(lst)-1] == "" {
		return lst[:len(lst)-1]
	}

	return lst
}

func must1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func must2[T1 any, T2 any](v T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v, v2
}

func must3[T1 any, T2 any, T3 any](v T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v, v2, v3
}
