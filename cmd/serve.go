package main

import (
	"fmt"
	"net/http"
	"tag-detect/pkg/bow"
	"tag-detect/router"
	"tag-detect/storage"
)

func main() {
	db := storage.Connect()
	defer db.Close()

	bow.ComputeIDF()
	r := router.NewRouter()
	fmt.Println("Listening on port 3030")
	http.ListenAndServe(":3030", r)

	// bow.ExtractWords("Báo cáo thẩm tra công tác!!! phòng chống??? tham nhũng nhấn$$$ mạnh việc thực hiện các biện pháp phòng ngừa tham nhũng, liên quan đến kê khai tài sản, thu nhập.")
}
