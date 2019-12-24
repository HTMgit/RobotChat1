package network

import (
	"encoding/json"
	"fmt"
	"robot_chat/global"
	"robot_chat/logger"
	"time"

	"github.com/parnurzeal/gorequest"
)

func LiaotianNvPuNetReq(msg string) (string, error) {
	sendStr := fmt.Sprintf("msg=%v&channelid=2001&openid=865166860721417", msg)
	fmt.Printf("Liaotiannvpuurl = %+v \n", global.Config.ReqUrl.Liaotiannvpuurl)
	resp, body, errs := gorequest.New().
		Timeout(3*time.Second).
		Post(global.Config.ReqUrl.Liaotiannvpuurl).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Send(sendStr).
		EndBytes()
	for _, err := range errs {
		if err != nil {
			logger.Logger.Error("[LiaotianNvPuNetReq] gorequest sendStr = %v,err: %v", sendStr, err)
			return "", err
		}
	}

	if resp.StatusCode != 200 {
		logger.Logger.Error("[LiaotianNvPuNetReq] gorequest sendStr = %v,resp.StatusCode : %v", sendStr, resp.StatusCode)
		return "", nil
	}
	var resBody map[string]interface{}
	err := json.Unmarshal(body, &resBody)
	if err != nil {
		logger.Logger.Error("[LiaotianNvPuNetReq] sendStr = %v,json.Unmarshal err : %v", sendStr, err)
		return "", err
	}

	var backStr string
	backStr = resBody["ret_message"].(string)
	// fmt.Printf("ref = %+v \n\n", resBody["ref"]) 额外数据 MP3，jpg等
	return backStr, nil

}

func XiaoYingNetReq(msg string) (string, error) {
	sendStr := fmt.Sprintf("%v?text=%v&location=&appid=73d2d21af6d8146692069f88b4406b88&cmd=chat&userid=0367C294CE11351385C836DF495C572A2", global.Config.ReqUrl.Xiaoyingurl, msg)
	fmt.Printf("Xiaoyingurl = %+v \n", global.Config.ReqUrl.Xiaoyingurl)
	fmt.Printf("sendStr = %+v \n", sendStr)

	resp, body, errs := gorequest.New().
		Timeout(3*time.Second).
		Post(sendStr).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Send("").
		EndBytes()
	for _, err := range errs {
		if err != nil {
			logger.Logger.Error("[XiaoYingNetReq] gorequest sendStr = %v,err: %v", sendStr, err)
			return "", err
		}
	}
	if resp.StatusCode != 200 {
		logger.Logger.Error("[XiaoYingNetReq] gorequest sendStr = %v,resp.StatusCode : %v", sendStr, resp.StatusCode)
		return "", nil
	}
	var resBody map[string]interface{}
	err := json.Unmarshal(body, &resBody)
	if err != nil {
		logger.Logger.Error("[XiaoYingNetReq] sendStr = %v,json.Unmarshal err : %v", sendStr, err)
		return "", err
	}
	dataMap := resBody["data"].([]interface{})
	data2Map := dataMap[0].(map[string]interface{})
	backStr := data2Map["value"].(string)
	fmt.Printf("resBody = %+v \n", backStr)

	return backStr, nil

}
