package main

import(
    "fmt"
    "encoding/json"
)

type NekoStruct struct {
    Name string
    Color string
}

func main() {
    stcData := NekoStruct{Name: "Kuro", Color: "black"}

    // Marchal関数でjsonエンコード
    // ->返り値jsonDataにはエンコード結果が[]byte形式で格納される
    jsonData, err := json.Marshal(stcData)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%s\n", jsonData)
}

