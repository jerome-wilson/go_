package main

import (
    "fmt"
    "web-servers2/models"
)

func main() {
    u := models.User{
        ID:        2,
        FirstName: "Jerome",
        LastName:  "Wilson",
    }
    fmt.Println(u)
}
