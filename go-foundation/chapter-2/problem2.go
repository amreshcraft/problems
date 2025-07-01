package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	 file,err := os.Open("./file.amresh")
	 if(err != nil){
		fmt.Println("error in opening a file")
	 }
	 defer file.Close()
	 haser := md5.New()

	 _ , err = io.Copy(haser,file)
	 if err != nil {
		fmt.Println("error in hashing")
	 }

	 hashed := haser.Sum(nil)
	 
	fmt.Printf("%x",hashed)
}