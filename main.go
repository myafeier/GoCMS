package main

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/renders"
	"github.com/go-macaron/session"
	"github.com/myafeier/GoCMS/routers/index"
	"gopkg.in/macaron.v1"
	"runtime"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	m := macaron.Classic()
	m.Use(cache.Cacher())
	// m.Use(session.Sessioner())
	m.Use(session.Sessioner(session.Options{
		Provider:       "memory",
		ProviderConfig: "",
		CookieName:     "Kfx",
		CookiePath:     "/",
		Gclifetime:     3600,
		Maxlifetime:    3600,
		Secure:         false,
		CookieLifeTime: 0,
		Domain:         "/",
		IDLength:       16,
		Section:        "session",
	}))
	m.Use(csrf.Csrfer())
	m.Use(captcha.Captchaer(captcha.Options{
		// 获取验证码图片的 URL 前缀，默认为 "/captcha/"
		URLPrefix: "/captcha/",
		// 表单隐藏元素的 ID 名称，默认为 "captcha_id"
		FieldIdName: "captcha_id",
		// 用户输入验证码值的元素 ID，默认为 "captcha"
		FieldCaptchaName: "captcha",
		// 验证字符的个数，默认为 6
		ChallengeNums: 6,
		// 验证码图片的宽度，默认为 240 像素
		Width: 240,
		// 验证码图片的高度，默认为 80 像素
		Height: 80,
		// 验证码过期时间，默认为 600 秒
		Expiration: 600,
		// 用于存储验证码正确值的 Cache 键名，默认为 "captcha_"
		CachePrefix: "captcha_",
	}))
	m.Use(renders.Renderer(renders.Options{
		Directory:       "templates",
		Extensions:      []string{".html"},
		Charset:         "UTF-8",
		IndentJSON:      true,
		IndentXML:       true,
		HTMLContentType: "text/html",
	}))

	m.Get("/", index.Index)

	m.NotFound(func(r renders.Render) {
		r.HTML(200, "404.html", map[string]interface{}{"Title": "Home"})
	})
	m.Run()
}
