package main

// go run Luna.go TOKEN SERVERID

import (
    "bufio"
    "os"
    "log"
    "fmt"

    "net/http"
)

func main() {
	log.Print("[~] Welcome To Luna!")
	
	token := os.Args[1]
	guild := os.Args[2]

	BanAll(token, guild)

	log.Print("[~] Finished!")
}

func BanAll(token string, guild string) {

	log.Print("[~] Loading IDs")

    file, err := os.Open("Members.txt")
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        go Send_Request(token, guild, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return
    }
}

func Send_Request(token string, guild string, user string) {
	url := fmt.Sprintf("https://discord.com/api/v8/guilds/%s/bans/%s", guild, user)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode == 204 {
		log.Print("Successfully Banned ", user)
	} 

}