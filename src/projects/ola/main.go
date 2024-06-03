package main

import (
	"fmt"
	"gopractice/projects/ola/im"
)

func main() {
	query := map[string]string{
		"out_random_id":    "95627954",
		"user_uuid":        "95627954",
		"out_message_id":   "R171716493594442850",
		"audit_op":         "4",
		"message_template": "1",
		"store_id":         "1489753817728680455",
		"app_id":           "7064034436848812032",
		"app_key":          "618c0310110b4b1a9505244fe28a43bc",
	}
	originData := im.JoinStringsInASCII(query, "&")
	signN := im.Md5Encrypt(originData)
	fmt.Println(signN)
}
