package main

import ( 
     "fmt"
     "os/exec"
   )
   
func main(){
      com := "kubectl"
      para1 := "get"
      para2 := "pods"
      
     cmd := exec.Command(com, para1, para2)
     stdout, err := cmd.Output()
     if err != nil {
     fmt.Println(err.Error())
     return
     }
fmt.Println("List of Pods")
fmt.Println(string(stdout))
}
