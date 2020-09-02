package fetcher

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"gospider/crawler/config"
	"gospider/crawler/proxypool"
)

var (
	rateLimiter = time.Tick(
		time.Second / config.Qps)
	verboseLogging = false
)

func SteVerboseLogging() {
	verboseLogging = true
}

func FetchNoProxy(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko)"+"Chrome/65.0.3325.181 Safari/537.36")
	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("return code is [%d]", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func Fetch(url string) ([]byte, error) {
	// Add Proxy Pool

	// targetUrl := "http://test.abuyun.com"

	// 初始化 proxy http client
	client := proxypool.AbuyunProxy{AppID: proxypool.ProxyUser, AppSecret: proxypool.ProxyPass}.ProxyClient()

	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte(``)))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
	// can remove
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("return code is [%d]", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body, err

}

func FetchOrignal(url string) ([]byte, error) {
	<-rateLimiter
	if verboseLogging {
		log.Printf("Fetching url %s ", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
