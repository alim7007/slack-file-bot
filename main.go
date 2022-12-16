package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main(){
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"WhatsApp Image 2022-12-02 at 16.18.56 (1).jpeg"}

	for i:=0; i<len(fileArr); i++{
		params:=slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err:= api.UploadFile(params)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("Name:%s, URL:%s\n", file.Name, file.URL)
	}
}


// OAUTH token for your workspace => install to workspace, scopes
// EVENT SUBSCRIPTION => enable
// SOCKET MODE => Enable Socket Mode
