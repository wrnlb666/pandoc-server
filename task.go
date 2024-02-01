package main

import (
    "os"
    "fmt"
    // "runtime"
    "strings"
    "os/exec"
)

func getHtml(filename string) string {
    // shell := ""
    // shellFlag := ""
    // if runtime.GOOS == "linux" {
    //     shell = "sh"
    //     shellFlag = "-c"
    // } else if runtime.GOOS == "windows" {
    //     shell = "cmd"
    //     shellFlag = "/c"
    // }

    fileSplited := strings.Split(filename, ".")
    if len(fileSplited) < 2 {
        fmt.Printf("[ERRO]: cannot determine type of the file '%s'\n", filename)
        os.Exit(1)
    }
    fileExtension := fileSplited[len(fileSplited) - 1]
    fileType := ""
    switch fileExtension {
    case "md", "MARKDOWN": {
        fileType = "markdown"
    }
    case "tex": {
        fileType = "latex"
    }
    default: {
        fmt.Printf("[ERRO]: cannot determine type of the file '%s'\n", filename)
        os.Exit(1)
    }
    }

    // cmd := exec.Command(shell, shellFlag, "pandoc", "-f", fileType, filename, "-t", "html")
    // cmd := exec.Command(shell, shellFlag, fmt.Sprintf("pandoc -f %s %s -t html", fileType, filename))
    cmd := exec.Command("pandoc", "-f", fileType, filename, "-t", "html")

    response, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return string(response)
}

