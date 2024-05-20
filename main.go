package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

func PlayAudio() error {
	cmd := exec.Command("afplay", "y1873.wav")
	err := cmd.Run()
	return err
}

type Response struct {
	RegDate     string `json:"regDate"`
	LeaveStatus int    `json:"leaveStatus"`
	IsWaiting   string `json:"isWaiting"`
	Leave       bool   `json:"leave"`
}

type ResponseTotal struct {
	Data []Response `json:"data"`
	Code string     `json:"resultCode"`
	Desc string     `json:"resultDesc"`
}

func do() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://mp.mhealth100.com/gateway/registration/appointment/dateSchedule/isLeave?branchCode=100238001&deptId=1100101", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Host", "mp.mhealth100.com")
	//req.Header.Add("Cookie", "Hm_lpvt_4f5e6d65812072c49089f068396b8513=1715939851; Hm_lvt_4f5e6d65812072c49089f068396b8513=1715876026,1715905473,1715906979,1715939851; gray-active=false; providerId=wechat; token=ZC7y_kyuPdR37mMK0tQdnw; userId=9A787DF81B894FA89774683B3E87C3E7; wechat_access_token=80_w71rUVKnQ8DC9-B3VsiqW95VGBKHBqNKANDJGc5wsj_esjaptsNxJKBFzcLZOVqAsXVsORvW3-XA-t4N2tyTQ0S6L5p777Y89GOlcOCrj3s; wechat_openId=oyGdsuIXQdexTXh_px8ZTFblweso")
	req.Header.Add("Cookie", "Hm_lpvt_4f5e6d65812072c49089f068396b8513=1716170513; Hm_lvt_4f5e6d65812072c49089f068396b8513=1715987592,1716024331,1716166232,1716170513; gray-active=false; providerId=wechat; token=wZC8vxWuWj0G-q2k5foYfA; userId=9A787DF81B894FA89774683B3E87C3E7; wechat_access_token=80_b5Aj3gHDCejXseALQCFlVPnuArPyAUlIFoTtR71xyTgHfbv-_zzhRC7GzaTRowqeW-0SE4s8HSc4HblJjXLPweFkfS3yFMcOQ1GPTlyaFGU; wechat_openId=oyGdsuIXQdexTXh_px8ZTFblweso")
	req.Header.Add("accept", "application/json")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.49(0x1800312c) NetType/4G Language/zh_CN")
	req.Header.Add("referer", "https://mp.mhealth100.com/patient/registration/")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	var respTemp ResponseTotal
	err = json.Unmarshal(body, &respTemp)
	fmt.Println(err)
	fmt.Println(respTemp)
	for _, v := range respTemp.Data {
		if respTemp.Code != "0" {
			fmt.Println("请求失败")
			return
		}
		if v.Leave {
			timeNow := time.Now()
			for {
				timeAfter := time.Now()
				if timeNow.Add(time.Minute * 5).After(timeAfter) {
					PlayAudio()
					fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "播放音频")
				} else {
					break
				}
			}
		}
	}
}

func main() {
	fmt.Println("系统开始运行")
	for {
		fmt.Println("时间：", time.Now().Format("2006/01/02 15:04:05"), "运行")
		do()
		time.Sleep(time.Minute)
	}
}
