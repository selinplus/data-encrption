package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/selinplus/data-encrption/models"
	"github.com/selinplus/data-encrption/pkg/app"
	"github.com/selinplus/data-encrption/pkg/e"
	"log"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//用户登录
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form LoginForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		log.Printf("BIND: %v", errCode)
		appG.Response(httpCode, errCode, nil)
		return
	}
	/*
		resp   := map[string]interface{}{}
		result := map[string]interface{}{}
		// 引入"crypto/tls":解决golang https请求提示x509: certificate signed by unknown authority
		ts := &tls.Config{InsecureSkipVerify: true}
		pMap := map[string]string{
			"username": form.Username,
			"password": form.Password,
		}
		_, body, errs := gorequest.New().TLSClientConfig(ts).
			Post(setting.DingtalkSetting.OapiHost + "/jeecg-boot/sys/login").
			Type(gorequest.TypeJSON).SendMap(pMap).End()
		if len(errs) > 0 {
			data := fmt.Sprintf("login err:%v", errs[0])
			appG.Response(http.StatusOK, e.ERROR, data)
			return
		} else {
			err := json.Unmarshal([]byte(body), &resp)
			if err != nil {
				data := fmt.Sprintf("unmarshall body error:%v", err)
				appG.Response(http.StatusOK, e.ERROR, data)
				return
			}
			if resp["result"] != nil {
				result = resp["result"].(map[string]interface{})
			}
		}
		data := map[string]interface{}{
			"success": resp["success"],
			"message": resp["message"],
			"token":   result["token"],
		}
	*/
	data := map[string]interface{}{
		"success": true,
		"message": "登录成功",
		"token":   "1571302017.o3in33a6eek.a118abf311e76c23d42357c53a58c906",
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//获取部门用户列表
func GetUserByDepartmentID(c *gin.Context) {
	var appG = app.Gin{C: c}
	DeptID := c.Query("deptId")
	users, err := models.GetUserListByDepartmentID(DeptID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if len(users) > 0 {
		appG.Response(http.StatusOK, e.SUCCESS, users)
	} else {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
	}
}
