package main

import (
	"log"
  "fmt"
	"github.com/xuri/excelize/v2"
)

func main(){
  file,err:=excelize.OpenFile("test.xlsx")
  if err!=nil{
    log.Fatal(err)
  }else{
    fmt.Println("file opend")
  }
  defer file.Close()
  rows,err:= file.GetRows("Sheet1")
  if err!=nil{
    fmt.Println(err)
  }
  for _,currentRow := range rows{
    for _,col:= range currentRow{
      fmt.Print(col,"\t")
    }
    fmt.Println()
  }
  
  
}