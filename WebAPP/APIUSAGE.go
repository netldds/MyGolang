package main

import (
	"net/http"

	"fmt"
	"os"
	"path/filepath"

	"log"
	"time"

	"strings"

	"io"

	"html/template"

	"context"
	"os/signal"

	"bytes"

	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"
	"mime"

	"database/sql"
	"math/rand"

	"strconv"

	"crypto/md5"

	"path"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	"golang.org/x/crypto/acme/autocert"
)

type Login struct {
	//绑定用的TAG声明
	// form对应form形式的提交(user=manu&password=123)
	//json就对应json形式的提交 ({"user": "manu", "password": "123"})
	//bind对应是否是必须字段
	User     string `form:"user" json:"user" bind:"required"`
	Password string `form:"password" json:"password bind:required" `
}

func APITestforPJT() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprint(c.Keys))
	})
	r.Run(":80")
}

//XML, JSON and YAML rendering
func renderData() {
	r := gin.Default()
	origin, _ := os.Getwd()
	//path.Clean("/src/netldds/WebAPP/"+"APIUSAGE.go")
	origin = filepath.Join(origin, "/src/github.com/netldds/WebAPP")

	//view specified file.
	//r.Static("/static", origin)

	//list directory
	r.StaticFS("/static", gin.Dir("./", true))
	origin = filepath.Join(origin, "/APIUSAGE.go")
	fmt.Println(origin)
	r.StaticFile("/staticfile", origin)

	fmt.Println(os.Getwd())

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/secureJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	r.Run(":80")
}

func Middleware() {

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in."})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		}
	})
	// Example for binding a HTML form (user=manu&password=123)
	f := func(c *gin.Context) {
		//绑定操作不影响正常表单操作。
		if c.Request.Method == "POST" {
			m, err := c.MultipartForm()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			}
			fl := m.File["file"]
			c.String(200, fl[0].Filename)
		}

		var form Login
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in."})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized."})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		}
	}
	r.POST("/LoginForm", f)
	r.GET("/loginForm", f)
	r.Run(":80")

}

func DefaultAPI() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", c.Request.Host)
	})

	r.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "page found.%s\n%s", c.Query("firstname"), c.Query("lastname"))
	})
	r.POST("/form_post", func(c *gin.Context) {
		name := c.DefaultPostForm("name", "lalala")
		c.JSON(200, gin.H{
			"status": "posted",
			"name":   name,
		})
	})
	r.POST("/query_form_post", func(c *gin.Context) {
		field1 := c.Query("id")
		field2 := c.PostForm("data")
		c.JSON(200, gin.H{
			"ID":   field1,
			"DATA": field2,
		})
	})

	r.MaxMultipartMemory = 8 << 20
	r.Static("/static", ".")
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for _, v := range files {
			c.SaveUploadedFile(v, v.Filename)
		}
		c.String(http.StatusOK, "%d files uploaded.", len(files))
	})

	v1 := r.Group("v1")
	v1.GET("/login", func(c *gin.Context) {
		c.String(200, "v1 login")
	})
	v2 := r.Group("v2")
	v2.GET("/login", func(c *gin.Context) {
		c.String(200, "v2 login")
	})

	//r.Run(":8080")
	r.RunTLS(":8080", "server.crt", "server.key")

}

func ClockStream() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		clientGone := c.Writer.(http.CloseNotifier).CloseNotify()
		c.Writer.Header().Set("Content-Type", "text/plain")
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		io.WriteString(c.Writer, strings.Repeat("# xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n", 13))
		var counter uint
		for {
			fmt.Fprintf(c.Writer, "%v\n", time.Now())
			c.Writer.(http.Flusher).Flush()
			select {
			case <-ticker.C:
				if counter < 40 {
					counter++
				} else {
					return
				}
			case <-clientGone:
				log.Printf("Client %v disconnected from the clock", c.Request.RemoteAddr)
				return
			}
		}
	})
	r.Run(":8080")
}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func PusherExample() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务
	r.Run(":8080")
}

//Serving data from reader
func ServingData() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		defer response.Body.Close()
		//contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		//extraHeader := map[string]string{
		//	"Content-Disposition": `attachment; filename="gopher.png"`,
		//}

		//c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeader)
		//直接转发内容
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		c.Data(http.StatusOK, contentType, buf.Bytes())
	})
	//r.Run(":80")
	r.RunTLS(":80", "server.crt", "server.key")
}
func DownloadFile() {
	fname := "F:/GOProjects/src/github.com/netldds/MyGolang/Misc/TempleExample.go"
	baseFilename := path.Base(fname)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.Header("Content-Description", "File Transfer")
		//c.Header("Content-Transfer-Encoding", "binary")
		//c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+baseFilename)
		c.File(fname)

	})
	r.Run(":80")
}

//HTML rendering
func htmlRendering() {
	r := gin.Default()

	//Custom Template Funcs
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	//这里不能同时匹配不同目录下的模板文件
	//可以用 https://github.com/gin-contrib/multitemplate
	r.LoadHTMLGlob("./templates/*.*")
	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.html", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "Users",
		})
	})

	r.Run(":80")
}
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

//Custom Middleware
//注册的中间件，内部是把它放在路由函数的前面执行。
func Logger() gin.HandlerFunc {
	//返回HandlerFunc函数之前是注册阶段会执行的部分
	return func(c *gin.Context) {
		//先执行中间件
		t := time.Now()
		//Set example variable
		c.Set("example", "12345")
		//before request
		//执行下一个中间件，可能是路由函数，就把content传给它执行 P302
		c.Next()
		//执行完路由函数后，接着执行中间件。
		//如果么有NEXT这个函数会先把这个中间件执行完毕，并在c.handlers[c.index](c)执行下一个中间件。
		latency := time.Since(t)
		log.Println(latency)

		//access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

	}
}
func middleRun() {
	r := gin.New()
	//r.Use(Logger())

	r.GET("/", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})
	r.Run(":80")
}

//Using BasicAuth() middleware
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func UsingBasicAuth() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.GET("/secrets", func(c *gin.Context) {

		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"secret": secret,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"secret": "NO SECRET :(",
			})
		}
	})
	r.Run(":80")
}

//Goroutines inside a middleware

func Goroutinesmiddleware() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})
	r.Run(":80")
}

//Support Let's Encrypt
func Encrypt() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	log.Fatal(autotls.Run(r, "example.com", "example2.com"))
}

//example for custom autocert manager
func AutoCert() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("./.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}

//Run multiple service using GIN
var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "welcome server 01",
		})
	})
	return e
}
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome server 02",
		})
	})
	return e
}
func RunMultiServices() {
	server01 := http.Server{
		Addr:         ":80",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := http.Server{
		Addr:         ":8080",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

//graceful shutdwon
func GracefulShutdown() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin NetWork.")
	})
	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	//^C退出操作会触发向信道发送信号
	signal.Notify(quit, os.Interrupt, os.Kill)
	//把服务跑在GOROUTINE里面，主线程完成退出前的操作。
	<-quit
	log.Println("Shutdown NetWork...")

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("NetWork Shutdown:", err)
	}
	log.Println("NetWork exiting")

}

//Bind form-data request with custom struct
type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_X"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}
func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
func BindFormData() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	//r.Run(":80")
	r.RunTLS(":443", "server.crt", "server.key")
}

func Practice() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.StaticFS("/pdf", gin.Dir("./templates/", true))

	g1 := r.Group("/login", gin.BasicAuth(gin.Accounts{
		"user1": "pwd1",
	}))

	g1.GET("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(http.StatusOK, "%s", user)
	})
	r.GET("/pdf1", func(c *gin.Context) {
		fl, _ := os.Open("templates/supp.pdf")
		data, _ := ioutil.ReadAll(fl)
		//"Content-Disposition": `attachment; filename="gopher.png"`,
		//c.Header("Content-Disposition",`attachment; filename="gopher.pdf"`)
		c.Data(http.StatusOK, mime.TypeByExtension(filepath.Ext("a.pdf")), data)
	})
	r.Run(":80")

}

//DB Operation
var DBSourceName1 string = "root:0000@tcp(192.168.1.146)/person"
var DBSourceName2 string = "ml:0000@tcp(127.0.0.1)/person"

func queryDB() {
	r := gin.Default()

	db, err := sql.Open("mysql", DBSourceName1)
	if err != nil {
		log.Fatal(err)
	}

	//insert
	r.GET("/databatch", func(c *gin.Context) {
		//id,_:=strconv.Atoi(c.Query("id"))
		var res sql.Result
		//var err error
		var stmt *sql.Stmt

		stmt, err = db.Prepare("insert labor_bak set name=?,age=?,occupation=?,password=?")

		for i := 0; i < 1000; i++ {

			name, age := randSqe(6)

			md := md5.New()
			io.WriteString(md, name)
			if rand.Intn(2) == 1 {
				res, err = stmt.Exec(name, age, "Friar", fmt.Sprintf("%x", md.Sum(nil)))

			} else {
				res, err = stmt.Exec(name, age, "Prist", fmt.Sprintf("%x", md.Sum(nil)))

			}

			ids, err := res.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(c.Writer, strconv.Itoa(int(ids))+"\n")
			c.Writer.(http.Flusher).Flush()
		}

	})

	//delete
	r.GET("/delete", func(c *gin.Context) {
		var stmt *sql.Stmt
		stmt, err := db.Prepare("delete from labor_bak where ID=?")
		if err != nil {
			log.Fatal(err)
		}
		str_id := c.Query("id")
		id, err := strconv.Atoi(str_id)
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(id)
		rowAffect, err := res.RowsAffected()
		c.String(http.StatusOK, "RowsAffected:%d\n", rowAffect)
	})

	/*
			int64
		float64
		bool
		[]byte
		string   [*]除了Rows.Next返回的不能是string.
		time.Time
	*/
	//select
	r.GET("/select", func(c *gin.Context) {
		str_id := c.Query("id")
		id, err := strconv.Atoi(str_id)
		rows, err := db.Query("select * from labor_bak where ID=?", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var age int
			var occupation string
			var password string
			if err := rows.Scan(&id, &name, &age, &occupation, &password); err != nil {
				log.Fatal(err)
			}
			io.WriteString(c.Writer, fmt.Sprintf("id:%d name=%s age=%d occupation=%s password=%s", id, name, age, occupation, password))
		}

	})

	//update
	r.GET("/update", func(c *gin.Context) {
		str_id := c.Query("id")
		str_age := c.Query("age")
		id, err := strconv.Atoi(str_id)
		age, err := strconv.Atoi(str_age)
		stmt, err := db.Prepare(`update labor_bak set age=? where ID=?`)

		res, err := stmt.Exec(age, id)
		if err != nil {
			log.Fatal(err)
		}
		i, err := res.RowsAffected()
		c.String(http.StatusOK, "RowsAffected:%d", i)
	})
	r.Run(":80")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSqe(n int) (string, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i, _ := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b), r.Intn(100)
}

func SetCookie() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.SetCookie()
	})
	r.Run(":80")
}

//http2 server push
//https 默认端口是443
func ServerPush() {

}
