package main

import "fmt"

func main() {
	var message = "This is really a slice of runes."
	fmt.Println(message[3], message[8], message[2], message[8], message[10], message[20], message[1], message[10], message[4], message[6], message[15], message[27], message[20], message[29])
	fmt.Println("\""+message[3:4]+message[8:9]+message[2:3]+message[8:9]+message[10:11]+message[20:21]+message[1:2]+message[10:11]+ message[4:5]+message[6:7]+message[15:16]+message[27:28]+message[20:21]+message[29:30]+"\"")
}