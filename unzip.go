package main

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }

    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            panic(err)
        }

        path := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            f, err := os.OpenFile(
                path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                panic(err)
            }

            _, err = io.Copy(f, rc)
            if err != nil {
                panic(err)
            }
            f.Close()
        }
        rc.Close()
    }
    r.Close()
    return nil
}
