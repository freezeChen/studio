/*
   @Time : 2019-05-16 16:11
   @Author : frozenchen
   @File : proto
   @Software: studio
*/
package template

var Proto_template =
	`syntax = "proto3";

package proto;


service Hello {
	rpc Hello (Req) returns (Reply) {
	}
}


message Req {
	string s = 1;
}

message Reply {
	string message = 1;
}


`
