package main

// go run Luna.go TOKEN SERVERID [ur code is horrible]

import (
	"time"
    "bufio"
    "os"
    "fmt" // why use log? 
    "net/http"
)

func main() {
	fmt.Println("[~] Welcome To Luna!")
	
	token := os.Args[1]
	guild := os.Args[2]

	BanAll(token, guild)

	fmt.Println("[~] Finished!")
}

func BanAll(token, guild string) {

    fmt.Println("[~] Loading IDs")

    file, err := os.Open("Members.txt")
    if err != nil {
		fmt.Println(err) // Print error
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        go Send_Request(token, guild, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
	fmt.Println(err) // Print error
        return
    }
}

func Send_Request(token, guild, user string) { // no need to add string besides at the end (duh)
	url := fmt.Sprintf("https://discord.com/api/v8/guilds/%s/bans/%s", guild, user)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err) // Print the error
		return
	}

	if resp.StatusCode == 204 {
		fmt.Printf("Successfully Banned %s\n", user)
		return // kill it
	} 
	if resp.StatusCode == 429 {
		fmt.Println("[!] Ratelimited")
		time.Sleep(1 * time.Second) // one second is more then enough
		Send_Request(token, guild, user)
	}

}
