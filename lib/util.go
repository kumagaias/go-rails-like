package lib

import (
	"os"
	"path"
	"runtime"
)

func GetDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// func PluckId(data []interface{}) []uint {
//   var ids []uint
//   for _, v := range data {
//     ids = append(ids, v.ID)
//   }
//   return ids
// }
