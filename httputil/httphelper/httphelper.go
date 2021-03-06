package httphelper

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const ContentType = "Content-Type"
const UserAgent = "User-Agent"

type HttpHelper struct { //线程安全
	DefaultTimeout time.Duration
	clt            *http.Client
}

var DefaultHttpHelper = NewHttpHelper()

func NewHttpHelper() *HttpHelper {
	return &HttpHelper{clt: http.DefaultClient, DefaultTimeout: 30 * time.Second}
}

func (h *HttpHelper) NewProxyFunc(ip, port string) func(*http.Request) (*url.URL, error) {
	proxyAddr := func(_ *http.Request) (*url.URL, error) {
		uristr := fmt.Sprintf("http://%s:%s", ip, port)
		return url.Parse(uristr)
	}
	return proxyAddr
}

func (h *HttpHelper) NewProxyTransport(ip, port string) *http.Transport {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		Proxy:                 h.NewProxyFunc(ip, port)}
	return transport
}

func (h *HttpHelper) SetNewClient(ip string, port uint) *HttpHelper {
	h.clt = &http.Client{Timeout: 30 * time.Second, Transport: h.NewProxyTransport(ip, strconv.Itoa(int(port)))}
	return h
}

func (h *HttpHelper) GetClient() *http.Client {
	return h.clt
}

func (h *HttpHelper) UseDefaultClient() *HttpHelper {
	h.clt = http.DefaultClient
	return h
}

func (h *HttpHelper) SetNewClientWithClient(clt *http.Client) *HttpHelper {
	h.clt = clt
	return h
}

func (h *HttpHelper) SetNewClientWithTransport(trans *http.Transport) *HttpHelper {
	h.clt = &http.Client{Transport: trans}
	return h
}

func (h *HttpHelper) Get(urlText string) (resp string, httpCode int, respHeader http.Header, err error) {
	return h.HttpRequestBase("GET", urlText, "", nil, true)
}

func (h *HttpHelper) GetWithHeader(urlText string, header map[string]string) (resp string, httpCode int, respHeader http.Header, err error) {
	return h.HttpRequestBase("GET", urlText, "", header, true)
}

func (h *HttpHelper) GetNoRedirect(urlText string) (resp string, httpCode int, respHeader http.Header, err error) {
	return h.HttpRequestBase("GET", urlText, "", nil, false)
}

func (h *HttpHelper) Post(urlText string, body string, header map[string]string) (resp string, httpCode int, respHeader http.Header, err error) {
	return h.HttpRequestBase("POST", urlText, body, header, true)
}

func (h *HttpHelper) AutoTryGuessHeader(body string) (header map[string]string) {
	if body == "" {
		return
	}
	header = make(map[string]string)
	header[UserAgent] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36"
	header[ContentType] = "application/x-www-form-urlencoded"
	if string(body[0]) == "{" || string(body[0]) == "[" {
		header[ContentType] = "application/json"
	}
	return
}

func (h *HttpHelper) NewRequest(method, urlText, body string, header map[string]string) (*http.Request, error) {
	var request *http.Request
	var er error
	if len(body) > 0 {
		request, er = http.NewRequest(method, urlText, strings.NewReader(body))
	} else {
		request, er = http.NewRequest(method, urlText, nil)
	}
	if er != nil {
		return nil, er
	}
	for k, v := range h.AutoTryGuessHeader(body) { //尝试猜测需要加的header
		request.Header.Add(k, v)
	}
	for k, v := range header { //如果客户端申明,可以覆盖猜测的
		request.Header.Set(k, v)
	}
	return request, nil
}

func (h *HttpHelper) Timeout(request *http.Request, duration time.Duration) *http.Request {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return request.WithContext(ctx)
}

func (h *HttpHelper) HttpRequestBase(method, urlText, body string, header map[string]string, autoRedirect bool) (respText string, httpCode int, respHeader http.Header, err error) {
	r, err := h.NewRequest(method, urlText, body, header)
	if err != nil {
		return
	}
	//default timeout
	return h.HttpRequest(h.Timeout(r, h.DefaultTimeout), autoRedirect)
}

func (h *HttpHelper) HttpRequestBytes(request *http.Request, autoRedirect bool) (respBytes []byte, httpCode int, respHeader http.Header, err error) {
	if request == nil {
		err = errors.New("request nil")
		return
	}
	clt := h.clt
	var resp *http.Response
	if !autoRedirect {
		trans := clt.Transport
		if trans == nil {
			trans = http.DefaultTransport
		}
		resp, err = trans.RoundTrip(request)
	} else {
		resp, err = clt.Do(request)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if resp == nil || err != nil {
		return
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	return bodyByte, resp.StatusCode, resp.Header, err
}

func (h *HttpHelper) HttpRequest(request *http.Request, autoRedirect bool) (respText string, httpCode int, respHeader http.Header, err error) {
	r, httpCode, respHeader, err := h.HttpRequestBytes(request, autoRedirect)
	respText = string(r)
	return
}
