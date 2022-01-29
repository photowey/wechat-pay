/*
 * @Author: photowey
 * @Date: 2022-01-29 21:37:29
 * @LastEditTime: 2022-01-29 21:47:51
 * @LastEditors: photowey
 * @Description: server.go
 * @FilePath: \wechat-pay\server\server.go
 * Copyright (c) 2022 by photowey<photowey@gmail.com>, All Rights Reserved.
 */

package server

import (
	"net/http"
)

type PaymentServer struct {
	srv *http.Server
}
