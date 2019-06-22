package file

import (
    "github.com/gin-gonic/gin"
    "github.com/mango-svip/wechat-spider/util"
    "net/http"
    "path"
    "strings"
)

const (
    IMAGE_PRE_URL = "http://localhost:8080/upload"
    IMAGE_SAVE_PATH = "upload"
    IMAGE_RUNTIME_PATH = "upload"

)


func GetImageUrl(name string) string {
    return IMAGE_PRE_URL + "/" + GetImagePath() + "/"+ name
}


func GetImageName(name string) string {
    ext := path.Ext(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = util.EncodeMD5(fileName)

    return fileName + ext
}

func GetImagePath() string {
    return IMAGE_SAVE_PATH
}

func GetImageFullPath() string {
    return IMAGE_RUNTIME_PATH + "/"+ GetImagePath()
}

func CheckImageExt(fileName string) bool {

    return true
}


func UploadImage(c *gin.Context) {
    code := 200
    data := make(map[string]string)

    _, header, e := c.Request.FormFile("image")

    if e != nil {
        code = 400
        c.JSON(http.StatusBadRequest, gin.H{
            "code" : code,
            "msg": e.Error(),
            "data": data,
        })
        return
    }

    if header == nil {
        code = 400
    } else {
        imaeName := GetImageName(header.Filename)
        fullPath := GetImageFullPath()
        savePath := GetImagePath()

        e := IsNotExistMkDir(fullPath)

        if e != nil {
            panic(e)
        }

        src := fullPath + "/" +imaeName

        e = c.SaveUploadedFile(header, src)

        if e != nil {
            panic(e)
        }

        data["image_url"] = GetImageUrl(imaeName)
        data["image_save_url"] = savePath + imaeName
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg": "success",
        "data": data,
    })
}