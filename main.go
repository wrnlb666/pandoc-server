package main

import (
	"fmt"
	"os"
    "time"
    "strconv"
    
    "github.com/labstack/echo/v4"
)

func main() {
    filename, port := handleFileName()

    modTime := getFileModify(filename)
    show := getHtml(filename)

    go func() {
        for range time.Tick(1 * time.Second) {
            m := getFileModify(filename)
            if m.After(modTime) {
                modTime = m
                show = getHtml(filename)
            }
        }
    }()

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        component := html(show)
        return component.Render(c.Request().Context(), c.Response())
    })
    e.GET("/body", func(c echo.Context) error {
        component := body(show)
        return component.Render(c.Request().Context(), c.Response())
    })
    e.GET("/main.go", func(c echo.Context) error {return nil})

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}


func handleFileName() (string, int) {
    args := os.Args[1:]
	if len(args) == 0 {
		fmt.Print(printHelp())
        os.Exit(255)
	}

    filename := ""
    port := 8080

    for i := 0; i < len(args); i++ {
        switch args[i] {
        case "-h": {
            printHelp()
            os.Exit(0)
        }
        
        case "-p": {
            i++
            if i == len(args) {
                fmt.Printf("optiona requires an argument '-- p'\n%s", printHelp())
                os.Exit(1)
            }
            p, err := strconv.Atoi(args[i])
            port = p
            if err != nil {
                fmt.Printf("Bad port %s\n", args[i])
                os.Exit(1)
            }
        }

        default: {
            if filename != "" {
                fmt.Printf("Bad argument '%s'\n%s", args[i], printHelp())
                os.Exit(1)
            }
            filename = args[i]
        }
        }
    }

    if filename == "" {
        fmt.Print(printHelp())
        os.Exit(255)
    }

    if _, err := os.Stat(filename); err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return filename, port
}


func getFileModify(filename string) time.Time {
    stat, err := os.Stat(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return stat.ModTime()
}


func printHelp() string {
	return fmt.Sprintf("[usage]: %s <filename> [optional]<flags>\n"+
        "   [-h] display this help menu\n"+
        "   [-p] specify port to start the server\n", os.Args[0])
}

