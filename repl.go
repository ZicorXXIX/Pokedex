package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)


func startRepl(){
    reader := bufio.NewScanner(os.Stdin)
    for{
        fmt.Print("Pokedex>")
        if !reader.Scan() {
            continue
        }
        input := reader.Text()

        cmd, found := commands[input]

        if found {
            if err:= cmd.callback(); err != nil{
                fmt.Printf("Error: %v \n",err)
            }
        }else {
            fmt.Println("Unknown Command")
        }
    }
  // How i would have implemented without instructions from bood.dev
  //   reader:= bufio.NewReader(os.Stdin)
  // //comma ok || error ok syntax
  // input, _ := reader.ReadString('\n')

}

func cleanInput(text string) []string {
    fields := strings.Fields(text)

    for i, s := range fields {
        fields[i] = strings.ToLower(s)
    }

    return fields
}

