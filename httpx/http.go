package httpx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/lcsin/gopocket/cryptor"
)

func POSTJson(url string, body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json;charset=utf-8", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}

	return respBody, nil
}

func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}

	return respBody, nil
}

func RUN(handler http.Handler, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		log.Printf("listening and serving HTTP on %s\n ", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server forced shutdown: %v\n", err)
	}

	log.Println("server exited ...")
}

// Sign 接口签名
func Sign(params map[string]string, key string) string {
	var fields []string
	for k, _ := range params {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	var sign string
	for _, v := range fields {
		sign += v
		val, _ := params[v]
		sign += val
	}
	sign += key

	return cryptor.MD5(sign)
}
