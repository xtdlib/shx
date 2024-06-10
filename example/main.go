package main

import (
	"log"

	"github.com/xtdlib/shx"
)

func main() {
	// {
	// 	files := shx.FD("")
	// 	for _, file := range files {
	// 		log.Println(file)
	// 	}
	// }

	{
		match := shx.RG("log")
		for _, file := range match {
			log.Printf("%s\n", file)
		}
	}

	log.Println("-----")

	{
		match := shx.RGFile("log")
		for _, file := range match {
			log.Printf("%s\n", file)
		}
	}
}
