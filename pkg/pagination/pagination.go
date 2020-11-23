/*
@Time : 2020/11/7 下午11:49
@Author : hoastar
@File : pagination
@Software: GoLand
*/

package pagination

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/jinzhu/gorm"
	"math"
)

type Param struct {
	C *gin.Context
	DB *gorm.DB
	ShowSQL	bool
}

type Paginator struct {
	TotalCount	int			`json:"total_count"`
	TotalPage	int			`json:"total_page"`
	Data		interface{} `json:"data"`
	PerPage		int			`json:"per_json"`
	Page		int			`json:"page"`
}


// ListRequest 分页请求中的参数
// Page 页数
// PerPage 当前页记录行数
// Sort 排序方式对应的int值
type ListRequest struct {
	Page	int	`json:"page" form:"page"`
	PerPage int `json:"per_page" form:"per_page"`
	Sort 	int `json:"sort" form:"sort"`
}

// Paging 分页查询操作
func Paging(p *Param, result interface{}, args ...interface{}) (*Paginator, error) {
	var (
		param		ListRequest
		paginator	Paginator
		count		int
		offset 		int
		tableName 	string
	)

	if err := p.C.Bind(&param); err != nil {
		logger.Errorf("参数绑定失败，错误：%v", err)
		return nil, err
	}

	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}

	if param.Page < 1 {
		param.Page = 1
	}

	if param.PerPage == 0 {
		param.PerPage = 10
	}

	if param.Sort == 0 || param.Sort == -1 {
		// ListRequest sort
		// 0 or -1 为结果倒序
		db = db.Order("id desc")
	}

	// args 为未经过初始化的nil切片
	if len(args) > 1 {
		// 不定参数用来存储 表名以及查询方式(如like, equal)和 预设的字段值
		tableName = fmt.Sprintf("'%s'.", args[1].(string))
	}

	if len(args) > 0 {
		for paramType, paramsValue := range args[0].(map[string]map[string]interface{}) {
			if paramType == "like" {
				for key, value := range paramsValue {
					db = db.Where(fmt.Sprintf("%v%v like ?", tableName, key), fmt.Sprintf("%%%v%%", value))
				}
			} else if paramType == "equal" {
				for key, value := range paramsValue {
					db = db.Where(fmt.Sprintf("%v%v = ?", tableName, key), value)
				}
			}
		}
	}

	done := make(chan bool, 1)
	go countRecords(db, result, done, &count)

	if param.Page == 1 {
		offset = 0
	} else {
		offset = (param.Page -1) * param.PerPage
	}


	err := db.Limit(param.PerPage).Offset(offset).Scan(result).Error
	if err != nil {
		logger.Errorf("数据查询失败，错误：%v", err)
		return nil, err
	}

	<-done

	paginator.TotalCount = count
	paginator.Data = result
	paginator.Page = param.Page
	paginator.PerPage = param.PerPage
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(param.PerPage)))

	return &paginator, nil
}


func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
