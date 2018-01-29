package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ryomak/nissy-video-tweet/twitter"
)

var api *twitter.User = &twitter.User{API: twitter.InitTwitterApi()}

var musicName []string = []string{
	"まだ君は知らない MY PRETTIEST GIRL 良すぎる...　#Nissy",
	"Don't let me go かっこいい...　#Nissy",
	"first love　クール... #Nissy",
	"The Eternal Live　やばい... #Nissy",
	"愛tears　ナイス！！　#Nissy",
	"KISS&DIVE　ライブ行きたい　#Nissy",
}

func main() {
	go VideoTweet()
	for {
		regularTweet()
	}
}

func VideoTweet() {
	for num := 0; num < len(musicName); num++ {
		path := "./media/video/" + strconv.Itoa(num) + ".mp4"
		if err := api.UploadMedia(musicName[num], path, "video/mp4", "tweet_video"); err != nil {
			fmt.Printf("err video: %v \n", err)
		}
		time.Sleep(86000 * time.Second)
	}
}

func regularTweet() {
	rand.Seed(time.Now().UnixNano())
	rNum := 0
	aaa, err := twitter.ScrapeAAA()
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	nissy, err := twitter.ScrapeNissy()
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	atae, err := twitter.ScrapeAtae()
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	//AAA
	text := aaa[0].Date + "\n" + aaa[0].Text + aaa[0].URL + " " + aaa[0].Tags
	err = api.ImageTweet(text, aaa[0].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)
	//Nissy
	text = nissy[0].Date + "\n" + nissy[0].Text + nissy[0].URL + " " + nissy[0].Tags
	err = api.ImageTweet(text, nissy[0].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)
	//Atae
	text = atae[0].Date + "\n" + atae[0].Text + atae[0].URL + " " + atae[0].Tags
	err = api.ImageTweet(text, atae[0].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)

	//AAA(random)
	rNum = rand.Intn(len(aaa))
	text = aaa[rNum].Date + "\n" + aaa[rNum].Text + aaa[rNum].URL + " " + aaa[rNum].Tags
	err = api.ImageTweet(text, aaa[rNum].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)
	//Nissy(random)
	rNum = rand.Intn(len(nissy))
	text = nissy[rNum].Date + "\n" + nissy[rNum].Text + nissy[rNum].URL + " " + nissy[rNum].Tags
	err = api.ImageTweet(text, nissy[rNum].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)
	//Atae(random)
	rNum = rand.Intn(len(atae))
	text = atae[rNum].Date + "\n" + atae[rNum].Text + atae[rNum].URL + " " + atae[rNum].Tags
	err = api.ImageTweet(text, atae[rNum].Image)
	if err != nil {
		fmt.Printf("err :%v\n", err)
	}
	fmt.Printf("%s\n", text)
	time.Sleep(7200 * time.Second)
}
