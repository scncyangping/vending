package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"vending/app/infrastructure/pkg/log"
)

func Post(url string, json string) (resBody string, err error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(json))
	if err != nil {
		log.Logger().Error("Post url error : %v", err)
		return
	}

	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return
}

func PostForm(url string, json string, action, sessionId string) (resBody string, err error) {
	res, err := http.PostForm(url, neturl.Values{"action": {action}, "data": {json}, "sessionId": {sessionId}})
	if err != nil {
		log.Logger().Error("Post url error : %v", err)
		return
	}
	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return
}

func PostFile(url string, filename string, botId string) (resBody string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("botid", botId)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)

	if err != nil {
		log.Logger().Error("Get file writer error : %v", err)
		return
	}

	fh, err := os.Open(filename)
	if err != nil {
		log.Logger().Error("error opening file : %v", err)
		return
	}

	io.Copy(fileWriter, fh)
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	resp, err := http.Post(url, contentType, bodyBuf)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	resBuffer, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Logger().Error("Read file post url res body error : %v", err)
		return
	}

	resBody = string(resBuffer)
	return resBody
}

func Get(url string) (resBody string) {
	res, err := http.Get(url)

	if err != nil {
		log.Logger().Error("Get url error : %v", err)
		return
	}

	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return resBody
}

func Put(url string, json string) (resBody string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(json))
	if err != nil {
		log.Logger().Error("Put url error : %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	resBuff, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Logger().Error("Read put url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return resBody
}

func Delete(url string) (resBody string) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Logger().Error("Delete url error : %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	resBuff, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Logger().Error("Read delete url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return resBody
}

func PostFormV(url string, v neturl.Values) (resBody string, err error) {
	res, err := http.PostForm(url, v)
	if err != nil {
		log.Logger().Error("Post url error : %v", err)
		return
	}
	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return
	}

	resBody = string(resBuff)
	return
}

func PostByType(url, json, cType string) (resBody string, err error) {
	res, err := http.Post(url, cType, strings.NewReader(json))
	if err != nil {
		log.Logger().Error("Post url error : %v", err)
		return
	}
	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return
	}
	resBody = string(resBuff)
	return
}

func PostR(url, cType string, sendData map[string]any) (map[string]any, error) {
	// 转换参数
	fData, error := json.Marshal(sendData)
	if error != nil {
		log.Logger().Info("json.Marshal Error: ", error, "data: ", sendData)
		return nil, error
	}
	dataString := string(fData)

	// 发送请求
	res, err := http.Post(url, cType, strings.NewReader(dataString))
	if err != nil {
		log.Logger().Error("Post url error : %v", err)
		return nil, err
	}
	defer res.Body.Close()
	resBuff, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Logger().Error("Read post url res body error : %v", err)
		return nil, err
	}
	resBody := string(resBuff)

	dat := make(map[string]any)

	error = json.Unmarshal([]byte(resBody), &dat)
	if error != nil {
		log.Logger().Info("json.Unmarshal Error: %v", resBody)
		return nil, error
	}

	log.Logger().Info("Post url :", url, "result json : ", dat)

	return dat, nil
}
