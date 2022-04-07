/*
 * @Author: photowey
 * @Date: 2022-02-01 15:05:05
 * @LastEditTime: 2022-02-01 17:13:45
 * @LastEditors: photowey
 * @Description: okhttp.go
 * @FilePath: /wechat-pay/internal/okhttp/okhttp.go
 * Copyright (c) 2022 by photowey<photowey@gmail.com>, All Rights Reserved.
 */

package okhttp

import (
    "net/url"
    "strings"

    gokhttp "github.com/BRUHItsABunny/gOkHttp"
)

// Post fire the gOkhttp post request
func Post(addr string, jsonBody []byte, headers map[string]string) (string, error) {
    headers = populateHeaders(headers)
    client := gokhttp.GetHTTPClient(nil)
    request, err := client.MakeMultiPartPOSTRequest(addr, "application/json", strings.NewReader(string(jsonBody)), headers)
    if err != nil {
        return "", err
    }
    response, _ := client.Do(request)
    bodyBytes, err := response.Bytes()
    bodyString := string(bodyBytes)

    return bodyString, err
}

// Get fire the gOkhttp post request
func Get(addr string, parameters, headers map[string]string) (string, error) {
    headers = populateHeaders(headers)
    client := gokhttp.GetHTTPClient(nil)

    values := toUrlValues(parameters)

    request, err := client.MakeGETRequest(addr, values, headers)
    if err != nil {
        return "", err
    }
    response, _ := client.Do(request)
    bodyBytes, err := response.Bytes()
    bodyString := string(bodyBytes)

    return bodyString, err
}

// populateHeaders populate default header, if necessary
func populateHeaders(headers map[string]string) map[string]string {
    if headers == nil {
        headers = map[string]string{}
        headers["Accept"] = "*/*"
    }

    return headers
}

// toUrlValues convert the map to url.values, if necessary
func toUrlValues(parameters map[string]string) url.Values {
    values := url.Values{}

    if parameters == nil {
        return values
    }

    for k, v := range parameters {
        values[k] = []string{v}
    }

    return values
}
