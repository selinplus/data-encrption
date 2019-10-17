package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/selinplus/data-encrption/middleware/cors"
	"github.com/selinplus/data-encrption/pkg/export"
	"github.com/selinplus/data-encrption/pkg/qrcode"
	"github.com/selinplus/data-encrption/pkg/upload"
	"github.com/selinplus/data-encrption/routers/api"
	v1 "github.com/selinplus/data-encrption/routers/api/v1"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORSMiddleware())
	//port := strconv.Itoa(setting.ServerSetting.HttpPort + 1)

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//上传文件
	r.POST("/file/upload", api.UploadFile)

	r.POST("/login", v1.Login)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//获取部门列表
		apiv1.GET("/department/list", v1.GetDepartmentByParentID)
		//获取部门用户列表
		apiv1.GET("/user/list", v1.GetUserByDepartmentID)
	}
	return r
}
