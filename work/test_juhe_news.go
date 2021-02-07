package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	NewType = "top"
	Key     = "2fb875ae24ef13d152a073a906825ff4"
)

type HttpResult struct {
	ErrorCode int32       `json:"error_code"`
	Reason    string      `json:"reason"`
	Result    *JuheResult `json:"result"`
}

type JuheResult struct {
	Stat string      `json:"stat"`
	Data []*JuheItem `json:"data"`
}

type JuheItem struct {
	UniqueKey       string `json:"uniquekey"`
	Title           string `json:"title"`
	Data            string `json:"date"`
	Category        string `json:"category"`
	AuthorName      string `json:"author_name"`
	Url             string `json:"url"`
	ThumbnailPicS   string `json:"thumbnail_pic_s"`
	ThumbnailPicS02 string `json:"thumbnail_pic_s02"`
	ThumbnailPicS03 string `json:"thumbnail_pic_s03"`
}

func main() {
	juheUrl := fmt.Sprintf("http://v.juhe.cn/toutiao/index?type=%s&key=%s", NewType, Key)
	resp, err := http.Get(juheUrl)
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:", err)
		return
	}

	httpResult := new(HttpResult)
	err = json.Unmarshal(b, &httpResult) // 参数无法解析直接丢弃
	if err != nil {
		fmt.Println("json unmarshal failed,err:", err)
		return
	}
	if httpResult.ErrorCode == 0 {
		for k, v := range httpResult.Result.Data {
			fmt.Printf("%d、%s %s\n", k+1, v.Title, v.Url)
		}
	}
}
