package file

import (
    "io/ioutil"
    "mime/multipart"
    "os"
    "path"
)

func GetSize(f multipart.File) (int, error) {
    content, e := ioutil.ReadAll(f)

    return len(content),e
}


func GetExt(fileName string) string {
    return path.Ext(fileName)
}


func CheckExist(src string) bool {
    _, e := os.Stat(src)
    return os.IsNotExist(e)
}

func CheckPromission(src string) bool {
    _, err :=  os.Stat(src)

    return os.IsPermission(err)
}


func IsNotExistMkDir(src string) error {

    if  notExist := CheckExist(src); notExist {
        if err := Mkdir(src); err != nil {
            return err
        }
    }
    return nil
}

func Mkdir(src string) error {
    err := os.MkdirAll(src, os.ModePerm)

    if err != nil {
        return err
    }
    return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
    file, e := os.OpenFile(name, flag, perm)

    if e != nil {
        return nil, e
    }

    return file, nil
}