package apis

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Address struct{}

func NewAddress() *Address {
	return &Address{}
}

func (a Address) GetAddressByIP(ip string) string {
	if ip == "127.0.0.1" || ip == "::1" {
		return "本地"
	}
	if a.check(ip) {
		// ip转地址
		if res := a.getAddress(ip); res != "" {
			return res
		}
		return "地址查询失败"
	}
	return "非法ip地址"
}

func (a Address) check(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func (a Address) getAddress(ip string) string {
	type response struct {
		Addr string `json:"addr"`
	}
	var resp response
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	_ = a.request(url, &resp, true)
	return resp.Addr
}

func (a Address) request(url string, dst any, toUtf8 bool) error {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if toUtf8 {
		body, _ = a.gbkToUtf8(body)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

func (a Address) gbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
