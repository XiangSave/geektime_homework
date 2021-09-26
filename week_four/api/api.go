package api

import (
	"fmt"
	"net/http"
	"week_four/internal/biz"
	"week_four/internal/pkg/mysql"
)

// func UnorderHandler(info *mysql.DBModel) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		restr := biz.UnOrderedString(info, 11111)
// 		fmt.Fprintln(w, restr)
// 	})
// }

func UnorderHandler(info *mysql.DBModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		restr := biz.UnOrderedString(info, 11111)
		fmt.Fprintln(w, restr)
	})
}
