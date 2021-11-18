// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"log"
	"net/smtp"
)

func ExampleSendMailQQMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "769460962@qq.com", "mbaxtfukgubzbbfc", "smtp.qq.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"769460962@qq.com"}
	msg := []byte("To: 769460962@qq.com\r\n" + // 收件人信息
		"From: 769460962@qq.com\r\n" + // 发件人信息
		"Subject: discount Gophers!\r\n" + // 主题
		"\r\n" +
		"This is the email body.\r\n") // 内容
	err := smtp.SendMail("smtp.qq.com:25", auth, "769460962@qq.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	// Output: Hello
}

func ExampleSendMail163Mail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "sb769460962@163.com", "AVBSWKHTMIORVVWZ", "smtp.163.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"sb769460962@163.com"}
	msg := []byte("To: sb769460962@163.com\r\n" + // 收件人信息
		"From: sb769460962@163.com\r\n" + // 发件人信息
		"Subject: discount Gophers!\r\n" + // 主题
		"\r\n" +
		"This is the email body.\r\n") // 内容
	err := smtp.SendMail("smtp.163.com:25", auth, "sb769460962@163.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	// Output: Hello
}
