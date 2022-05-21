package Context

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) RunServer(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := newContext(w, req)
	engine.router.handle(context)
}

func Run() {
	r := New()
	r.GET("/", func(context *Context) {
		context.HTML(http.StatusOK, "<h1>hello Gan</h1>")
	})
	r.GET("/hello", func(context *Context) {
		context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
	})
	r.POST("/login", func(context *Context) {
		context.JSON(http.StatusOK, H{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	r.RunServer(":9999")

}
