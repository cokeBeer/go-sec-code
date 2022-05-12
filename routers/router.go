package routers

import (
	"fmt"
	"go-sec-code/controllers"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/favicon.ico", &controllers.FaviconController{})
	beego.Router("/commandInject/vuln", &controllers.CommandInjectVuln1Controller{})
	beego.Router("/commandInject/vuln/host", &controllers.CommandInjectVuln2Controller{})
	beego.Router("/commandInject/safe", &controllers.CommandInjectSafe1Controller{})
	beego.Router("/cors/vuln/reflect", &controllers.CorsVuln1Controller{})
	beego.Router("/cors/vuln/any-origin-with-credential", &controllers.CorsVuln2Controller{})
	beego.Router("/cors/safe", &controllers.CorsSafe1Controller{})
	beego.Router("/crlfInjection/safe", &controllers.CRLFSafe1Controller{})
	beego.Router("/fileUpload/vuln", &controllers.FileUploadVuln1Controller{})
	beego.Router("/fileUpload/safe", &controllers.FileUploadSafe1Controller{})
	beego.Router("/jsonp/vuln/noCheck", &controllers.JsonpVuln1Controller{})
	beego.Router("/jsonp/vuln/emptyReferer", &controllers.JsonpVuln1Controller{})
	beego.Router("/jsonp/safe", &controllers.JsonpSafe1Controller{})
	beego.Router("/pathTraversal/vuln", &controllers.PathTraversalVuln1Controller{})
	beego.Router("/pathTraversal/safe", &controllers.PathTraversalSafe1Controller{})
	beego.Router("/sqlInjection/native/vuln/integer", &controllers.SqlInjectionVuln1Controller{})
	beego.Router("/sqlInjection/native/vuln/string", &controllers.SqlInjectionVuln2Controller{})
	beego.Router("/sqlInjection/native/safe/integer", &controllers.SqlInjectionSafe1Controller{})
	beego.Router("/sqlInjection/native/safe/string", &controllers.SqlInjectionSafe2Controller{})
	beego.Router("/ssrf/vuln", &controllers.SSRFVuln1Controller{})
	beego.Router("/ssrf/vuln/confuse", &controllers.SSRFVuln2Controller{})
	beego.Router("/ssrf/vuln/302", &controllers.SSRFVuln3Controller{})
	beego.Router("/ssrf/safe/whitelists", &controllers.SSRFSafe1Controller{})
	beego.Router("/ssti/vuln", &controllers.SSTIVuln1Controller{})
	beego.Router("/ssti/safe", &controllers.SSTISafe1Controller{})
	beego.Router("/xss/vuln", &controllers.XSSVuln1Controller{})
	beego.Router("/xss/vuln/store", &controllers.XSSVuln2Controller{})
	beego.Router("/xss/safe", &controllers.XSSSafe1Controller{})
	beego.Router("/xxe/vuln", &controllers.XXEVuln1Controller{})
	// beego.Router("/xxe/vuln/gokogiri", &controllers.XXEVuln2Controller{})
	beego.Router("/xxe/safe", &controllers.XXESafe1Controller{})
	beego.Router("/abc/aaa", &controllers.BypassController{})
	beego.InsertFilter("/abc/aaa", beego.BeforeStatic, func(ctx *context.Context) {
		fmt.Println(fmt.Sprintf("filtered url:%s", ctx.Input.URL()))
	})
}