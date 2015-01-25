package main

/*
Assigment:

Given a personal information such as an email address, IP address, or a public key,
the program you will write needs to generate a unique avatar. Imagine that you are 
building a new application and you want all of your users to have a default and 
unique avatar. The package you will write will allow the generation of such avatars. 
GitHub recently used such an approach and generates an identicon for all new users 
who don't have a gravatar account attached.
*/

import (
	"fmt"
	"os"
	"net/http"
	"crypto/md5"
	"io/ioutil"
	"io"
	"bytes"
)

func main() {
	var pass string
	fmt.Print("Enter your email: ")
	_,err := fmt.Scanf("%s", &pass)
	
	if err != nil {
		return
	}
	hash := generate_hash(pass)
	generate_img(hash)
}

func generate_img(hash string) {
	resp, err := http.Get("http://identicon.org/?t="+hash+"&s=50")
	contents, err := ioutil.ReadAll(resp.Body)
	
        if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
        }

	out, err := os.Create("new.png")
	defer out.Close()
	
	if err != nil {
		os.Exit(1)
	}

	io.Copy(out, bytes.NewBufferString(string(contents)))
}

func generate_hash(pass string) string {
	data := []byte(pass)
	return fmt.Sprintf("%x", md5.Sum(data))
}
