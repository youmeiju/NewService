package controllers

import (
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"bytes"
)

func pushMsg(orderNO string,toUser interface{})  {
	redis, err := cache.NewCache("redis", `{"conn":"118.31.72.110:6379", "key":"beecacheRedis"}`)
	if err==nil{
		var AccessToken interface{}
		accessToken := redis.Get("access_token")
		if accessToken==nil {
			url:= fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxb4e8fb564d2b9cca&secret=29b49b5a99dc02b7d95703896696df8c")
			res, err := http.Get(url)
			if err == nil {
				defer res.Body.Close()
				m := make(map[string]interface{})
				body, _ := ioutil.ReadAll(res.Body)
				json.Unmarshal(body, &m)
				AccessToken = m["access_token"]
				redis.Put("access_token", AccessToken, 7200*time.Second)
			}
		}else{
			AccessToken=string(accessToken.([]byte))
		}
		url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" +AccessToken.(string)
		PostData := make(map[string]interface{})
		template_id := "mXxbn_Rpwukie05JMhQ5nlqbxo-fu9MdMFJ6Bo4Bkio"
		touser := toUser
		data := make(map[string]interface{})
		keyword1 := make(map[string]string)
		keyword1["value"] = orderNO
		keyword2 := make(map[string]string)
		keyword2["value"] = time.Now().Format("2006-01-02 15:04:05")
		first := make(map[string]string)
		first["value"] = "你有新的订单，请及时处理"
		remark := make(map[string]string)
		remark["value"] = ""
		data["keyword1"] = keyword1
		data["keyword2"] = keyword2
		data["first"] = first
		data["remark"] = remark
		PostData["touser"] = touser
		PostData["template_id"] = template_id
		PostData["data"] = data
		a,_ := json.Marshal(&PostData)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(a))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
