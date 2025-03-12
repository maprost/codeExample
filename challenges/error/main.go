package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

}

func CopyFileNormal(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}


func CopyFile(src, dst string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("copy %s to %s: %v", src, dst, err)
		}
	}()

	r := guard os.Open(src)
	defer must r.Close()

	w := guard os.Create(dst)
	defer must w.Close()

	err = io.Copy(w, r)

	// here we need to do extra stuff when an Copy error happens: now we must use the 'normal' error handling method, and cannot use guard or must
	if err != nil {
		_ := os.Remove(dst) // fail silently if errors happen during error handling
	}
	return
}