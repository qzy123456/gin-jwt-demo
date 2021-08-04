package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwtDemo/utils"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//单个文件上传
func Upload(ctx *gin.Context) {
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	address := fmt.Sprintf("%s:%d%s", GlobalService.Conf.Server.Address, GlobalService.Conf.Server.HttpPort,"/")
	file, err := ctx.FormFile("file")
	var data = make(map[string]interface{})
	if err != nil {
		fmt.Println("获取数据失败")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "获取数据失败",
			"data":    data,
		})
	} else {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			ctx.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		fmt.Println("接收的数据", file.Filename)
		//获取文件名称
		fmt.Println(file.Filename)
		//文件大小
		fmt.Println(file.Size)
		//获取文件的后缀名
		extstring := path.Ext(file.Filename)
		fmt.Println(extstring)
		//根据当前时间鹾生成一个新的文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		//新的文件名
		fileName := fileNameStr + extstring
		//保存上传文件
		filePath := filepath.Join(utils.Mkdir("upload"), "/", fileName)
		ctx.SaveUploadedFile(file, filePath)
		data["url"] = address + filePath
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    data,
		})
	}
}

//多个文件上传
func Uploads(ctx *gin.Context) {
	var data = make(map[string]interface{})
	address := fmt.Sprintf("%s:%d%s", GlobalService.Conf.Server.Address, GlobalService.Conf.Server.HttpPort,"/")
	if form, err := ctx.MultipartForm(); err == nil {
		//1.获取文件
		files := form.File["file"]
		//2.循环全部的文件
		for k, file := range files {
			fileExt := strings.ToLower(path.Ext(file.Filename))
			if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
				ctx.JSON(200, gin.H{
					"code": 400,
					"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
				})
				return
			}
			// 3.根据时间鹾生成文件名
			fileNameInt := time.Now().Unix()
			fileNameStr := strconv.FormatInt(fileNameInt, 10)
			//4.新的文件名(如果是同时上传多张图片的时候就会同名，因此这里使用时间鹾加文件名方式)
			fileName := fileNameStr + file.Filename
			//5.保存上传文件
			filePath := filepath.Join(utils.Mkdir("upload"), "/", fileName)
			ctx.SaveUploadedFile(file, filePath)
			url := "url_" + strconv.Itoa(k)
			data[url] = address + filePath
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "上传成功",
			"data":    data,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取数据失败",
			"data":    data,
		})
	}
}
