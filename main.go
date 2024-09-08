package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

var TeleBot *tele.Bot

// MapToken 存储的 Sender 类型
var MapToken = map[int64]*tele.User{}

func createGin() error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/notice", func(context *gin.Context) {
		// 获取参数
		userToken := context.PostForm("token")

		// 判断是否存在
		if userToken == "" {
			context.JSON(200, gin.H{
				"code": 501,
				"msg":  "请输入token",
			})
			return
		}
		i, _ := strconv.ParseInt(userToken, 10, 64)
		userId, exists := MapToken[i]
		if !exists {
			context.JSON(200, gin.H{
				"code": 502,
				"msg":  "token不存在",
			})
			return
		}
		// 获取消息

		title := context.PostForm("title")
		content := context.PostForm("content")

		_, err := TeleBot.Send(userId, title+"\n"+content)
		if err != nil {
			context.JSON(200, gin.H{
				"code": 503,
				"msg":  "发送失败",
			})
			return
		}

		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "发送成功",
		})

	})

	s := &http.Server{
		Addr:           ":6001",
		Handler:        r,
		ReadTimeout:    4 * time.Second,
		WriteTimeout:   4 * time.Second,
		MaxHeaderBytes: 0 << 20,
	}
	err := s.ListenAndServe()

	return err
}

func createBot() (*tele.Bot, error) {
	// 获取命令行参数
	args := os.Args
	// 打印命令本身和所有参数
	teleToken := ""
	for i, arg := range args {
		fmt.Println("Arg", i, "is", arg)
		if i == 1 {
			// 分割
			teleToken = arg
		}
	}

	pref := tele.Settings{
		Token:  teleToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	b.Handle("/start", func(c tele.Context) error {
		// 查找token
		user := MapToken[c.Sender().ID]
		if user != nil {
			return c.Send("你已经注册过了")
		}
		// 存储用户的 token
		MapToken[c.Sender().ID] = c.Sender()
		return c.Send("你好，注册成功")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {

		return c.Send("你好，我是一个机器人")
	})

	println("createBot ok...")

	return b, nil
}

func main() {

	bot, err := createBot()
	if err != nil {
		log.Fatal(err)
		return
	}
	TeleBot = bot

	// 同协
	go func() {
		TeleBot.Start()
	}()

	err = createGin()
	if err != nil {
		log.Fatal(err)
		return
	}
	println("exit...")
}
