package ModuleExample

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func makeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	jsonStr := `{"Code":0,"Message":"获取数据成功","Data":"热榜数据","CreateTime":1569673821}]}`
	w.Write([]byte(jsonStr))
}

func Cmain() {
	http.HandleFunc("/a", makeGzipHandler(handler)) // 设置访问的路由
	http.ListenAndServe(":1113", nil)
}
