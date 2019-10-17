package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/selinplus/data-encrption/models"
	"github.com/selinplus/data-encrption/pkg/app"
	"github.com/selinplus/data-encrption/pkg/e"
	"net/http"
)

//获取部门列表
func GetDepartmentByParentID(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Query("id")
	parentDt, errd := models.GetDepartByID(id)
	if errd != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEPARTMENT_FAIL, nil)
		return
	}
	var dts []interface{}
	data := map[string]interface{}{
		"id":       parentDt.ID,
		"parentid": parentDt.ParentID,
		"name":     parentDt.DepartName,
		"children": dts,
	}
	departments, err := models.GetDepartByParentID(id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEPARTMENT_FAIL, nil)
		return
	}
	if len(departments) > 0 {
		for _, department := range departments {
			leaf := models.IsLeafDepart(department.ID)
			dt := map[string]interface{}{
				"id":       department.ID,
				"parentid": department.ParentID,
				"name":     parentDt.DepartName,
				"isLeaf":   leaf,
			}
			dts = append(dts, dt)
		}
		data["children"] = dts
		appG.Response(http.StatusOK, e.SUCCESS, data)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, data)
	}
}
