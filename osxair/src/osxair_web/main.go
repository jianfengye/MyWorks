package main

import (
    "log"
    "net/http"
    "os"
    "os/exec"
    "path"
)

func main() {
    file, err := exec.LookPath(os.Args[0])
    if err != nil {
        panic(err)
    }

    dir, _ := path.Split(file)
    os.Chdir(dir + "/../src/osxair_web/public")
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    log.Println("please listen in port 3200")
    err = http.ListenAndServe(":3200", http.FileServer(http.Dir(wd)))
    if err != nil {
        panic(err)
    }
}
