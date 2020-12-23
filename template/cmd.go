/*
   @Time : 2019-05-08 17:09
   @Author : frozenchen
   @File : cmd
   @Software: studio
*/
package template

var Conf_template = `
name: "{{.Appname}}"
env: "dev"
version: "v0.0.1"
mysql:
	source: "test:test@tcp(172.16.0.148:3306)/studio?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"
	active: 100
	idle: 20
log:
	name: "{{.Appname}}"
	debug: true
	LogFileDir: "./log/"
	KafkaAddr: "127.0.0.1:9092"
	WriteKafka: false
redis:
	addr: "127.0.0.1:6379"
	auth: ""
	active: 100
	idle: 20
	dialTimeout: "1ms"
	readTimeout: "1ms"
	writeTimeout: "1ms"
	idleTimeout: "1ms"
kafka: "127.0.0.1:9092"
`
