package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var logger = new(zap.Logger)

type param struct {
	TitleId   int `json:"title_id"`
	AnswerOpt int `json:"answer_opt"`
}

func handleCommitAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm)
	//fmt.Println("header:", r.Header)
	if r.Method == "GET" {
		w.Write([]byte("hello1"))
		return
	}

	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		//fmt.Println("len(b):", len(b), "cap(b):", cap(b))
		//fmt.Println("body:", string(b))

		arg := new(param)
		err := json.Unmarshal(b, arg)
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}
		fmt.Println("arg:", arg)

		w.Write([]byte("hello2"))
		return
	}
}

func handleCommitAnswerList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm)
	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		list := make([]*param, 0)
		err := json.Unmarshal(b, &list)
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}
		logger.Info("handleCommitAnswerList", zap.Any("list", list))

		w.Write([]byte("hello2"))
		return
	}

	w.Write([]byte("error request"))
}

type question struct {
	Id    int64    `json:"id"`
	Title string   `json:"title"`
	Opts  []string `json:"opts"`
}

func handlePullQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path2:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm, "header:", r.Header)
	//fmt.Println("header:", r.Header)
	if r.Method == "GET" {
		ques := question{
			Id:    123,
			Title: "下面哪个是中国的四大发明???",
			Opts:  []string{"火药1", "中药2", "火锅3", "指北针4"},
		}
		bin, _ := json.Marshal(ques)
		w.Write(bin)
		return
	}
}

func handlePullQuestionList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path3:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm, "header:", r.Header)
	//fmt.Println("header:", r.Header)
	if r.Method == "GET" {
		list := make([]*question, 0, 3)
		list = append(list, &question{
			Id:    1,
			Title: "下面哪个是中国的四大发明",
			Opts:  []string{"火药", "中药", "火锅", "指北针"},
		})
		list = append(list, &question{
			Id:    2,
			Title: "徐悲鸿以画什么闻名世界",
			Opts:  []string{"虾", "山水画", "奔马", "大雁"},
		})
		list = append(list, &question{
			Id:    3,
			Title: "\"人生自古谁无死，留取丹心照汗青\"的作者是？",
			Opts:  []string{"陆游", "辛弃疾", "岳飞", "文天祥"},
		})
		bin, _ := json.Marshal(list)
		w.Write(bin)
		return
	}
}

func main() {
	fmt.Println("answer server start...")
	logger, _ = zap.NewProduction()

	handler := http.NewServeMux()
	handler.HandleFunc("/commit_answer", handleCommitAnswer)
	handler.HandleFunc("/commit_answer_list", handleCommitAnswerList)
	handler.HandleFunc("/pull_question", handlePullQuestion)
	handler.HandleFunc("/pull_question_list", handlePullQuestionList)

	server := &http.Server{Addr: ":8080", Handler: handler}
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
