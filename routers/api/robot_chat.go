package api

import (
	"fmt"
	"net/http"
	"robot_chat/logger"
	"robot_chat/models"
	"robot_chat/myErr"
	"robot_chat/network"
	"robot_chat/store"
	"strings"

	"github.com/gin-gonic/gin"
)

func RobotChat(c *gin.Context) {
	req := models.ChatReq{}
	rsp := make(map[string]interface{})

	err := c.ShouldBindJSON(&req)
	logger.Logger.Info("[GetUserInfo] req:%+v ", req)
	if err != nil {
		InvalidParamErr(c)
		return
	}
	//TODO:信息判断处理
	//查询是否使用管家
	userInfo := store.UserInfoData{}
	userInfo, err = store.SelectRobotStatusMysql(req.UserID)
	if err != nil {
		UniversalProcessingErr(c, myErr.RespErrServer, nil)
		return
	}
	if userInfo.RobotStatus == 0 {
		UniversalProcessingErr(c, myErr.ROBOTCLOSE, nil)
		return
	}
	//TODO:判断消息类型

	var backString string
	//查询特殊回复的关键字
	keys := make([]*store.KeyWordData, 0)
	keys, _ = store.SelectAllKeyWorldsMysql()
	for _, key := range keys {
		loc := strings.Index(req.Content, key.KeyID)
		if loc >= 0 {
			backString = key.Backworld
			break
		}
	}

	backData := models.ChatReq{}
	backData.UserID = req.UserID
	backData.ChatType = req.ChatType
	if len(backString) > 0 {
		backString = ExchangeKey(backString)
		backData.Content = backString
		rsp["err_code"] = myErr.RespOK
		rsp["msg"] = myErr.GetMsg(myErr.RespOK)
		rsp["data"] = backData
		logger.Logger.Info("GetUserInfo rsp = %+v\n", rsp)
		c.JSON(http.StatusOK, &rsp)
		return
	}
	//请求网络数据
	backString, _ = network.XiaoYingNetReq(req.Content)
	fmt.Printf("小影回答 ：%v\n", backString)
	// backString, _ = network.LiaotianNvPuNetReq(req.Content)
	//查询回复错误关键字
	badworldArr := make([]*string, 0)
	badworldArr, _ = store.SelectAllBadWorldMysql()
	isAble := 1
	if len(backString) > 0 {
		for _, badkey := range badworldArr {
			loc := strings.Index(backString, *badkey)
			if loc >= 0 {
				isAble = 0
				break
			}
		}
	} else {
		isAble = 0
	}

	if isAble == 0 {
		fmt.Printf("请求聊天机器人接口 \n\n")
		//请求聊天机器人接口

		backString, _ = network.LiaotianNvPuNetReq(req.Content)
		fmt.Printf("聊天女仆回答 ：%v\n", backString)
		// backString, _ = network.XiaoYingNetReq(req.Content)
		if len(backString) > 0 {
			for _, badkey := range badworldArr {
				loc := strings.Index(backString, *badkey)
				if loc >= 0 {
					backString = "我也不懂怎么回答你了，我要报告老大 你难住我了"
					break
				}
			}
		} else {
			backString = "我也不懂怎么回答你了，我要报告老大 你难住我了"
		}
	}

	backString = ExchangeKey(backString)
	backData.Content = backString
	rsp["err_code"] = myErr.RespOK
	rsp["msg"] = myErr.GetMsg(myErr.RespOK)
	rsp["data"] = backData
	logger.Logger.Info("GetUserInfo rsp = %+v\n", rsp)
	c.JSON(http.StatusOK, &rsp)
	return

}

func ChangeRobotStatus(c *gin.Context) {
	req := models.ChangeStatusReq{}
	rsp := make(map[string]interface{})

	err := c.ShouldBindJSON(&req)
	logger.Logger.Info("[ChangeRobotStatus] req:%v", req)
	if err != nil {
		InvalidParamErr(c)
		return
	}

	err = store.UpdateRobotStatusMysql(req.UserID, req.Status)
	if err != nil {
		UniversalProcessingErr(c, myErr.RespErrServer, nil)
		return
	}

	rsp["err_code"] = myErr.RespOK
	rsp["msg"] = myErr.GetMsg(myErr.RespOK)

	backData := models.ChatReq{}
	backData.UserID = req.UserID
	backData.Content = "那就再见啦！想我的时候直接说‘小管家’"
	rsp["data"] = backData

	logger.Logger.Info("GetUserInfo rsp = %+v\n", rsp)
	c.JSON(http.StatusOK, &rsp)
	return
}

//查询是否需要替换字符，并返回
func ExchangeKey(original string) string {
	exworldArr := make([]*store.ExchangeWordData, 0)
	exworldArr, _ = store.SelectAllExchangeWorldMysql()
	fmt.Printf("exworldArr = %+v \n", exworldArr)
	if len(exworldArr) > 0 {
		for _, exkey := range exworldArr {
			fmt.Printf("original = %v ,exworld = %+v，%v \n", original, exkey.KeyID, exkey.Exworld)
			original = strings.Replace(original, exkey.KeyID, exkey.Exworld, -1)
			fmt.Printf("original = %v ", original)
		}
	}
	return original
}

//通用错误返回
func UniversalProcessingErr(c *gin.Context, errCode int, data interface{}) {
	rsp := make(map[string]interface{})
	rsp["err_code"] = errCode
	rsp["msg"] = myErr.GetMsg(errCode)
	if data != nil {
		rsp["data"] = data
	}
	logger.Logger.Error("[UniversalProcessingErr] rsp Error:%+v", rsp)
	c.JSON(http.StatusOK, &rsp)
}

//参数解析错误
func InvalidParamErr(c *gin.Context) {
	rsp := make(map[string]interface{})
	rsp["err_code"] = myErr.RespErrInvalidParam
	rsp["msg"] = myErr.GetMsg(myErr.RespErrInvalidParam)
	logger.Logger.Error("[InvalidParamErr] rsp Error:%+v", rsp)
	c.JSON(http.StatusOK, &rsp)

}
