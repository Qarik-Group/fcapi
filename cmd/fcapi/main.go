package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/http"
	"github.com/starkandwayne/fcapi/memory"
)

func meat() int {
	cfURL := os.Getenv("CF_ENDPOINT")
	if len(cfURL) == 0 {
		fmt.Println("CF_ENDPOINT not set")
		return 1
	}

	uaaURL := os.Getenv("UAA_ENDPOINT")
	if len(uaaURL) == 0 {
		fmt.Println("UAA_ENDPOINT not set")
		return 1
	}

	urls := map[string]string{
		"cfURL":  cfURL,
		"uaaURL": uaaURL,
	}

	services := memory.NewServices()
	services.Logger = func(c echo.Context) {
		req := c.Request()

		headers := make([]string, 0)

		for key, value := range req.Header {
			headers = append(headers, fmt.Sprintf("%s: %s", key, value))
		}

		params := make([]string, 0)

		for key, value := range params {
			params = append(params, fmt.Sprintf("%s: %s", key, value))
		}

		output := make([]string, 0)
		output = append(output, fmt.Sprintf("%s %s", req.Method, req.URL.String()))
		output = append(output, fmt.Sprintf("Headers {%s}", strings.Join(headers, ", ")))

		if len(params) > 0 {
			output = append(output, fmt.Sprintf("Params {%s}", strings.Join(params, ", ")))
		}

		if req.Method == "PATCH" || req.Method == "POST" || req.Method == "PUT" {
			body, _ := ioutil.ReadAll(req.Body)
			req.Body.Close()

			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			output = append(output, fmt.Sprintf("Body {%s}", string(body)))
		}

		fmt.Println(strings.Join(output, " | "))
	}

	server := http.Server("0.0.0.0:8002", services, urls)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("stopping the server")
		}
	}()

	fmt.Printf("listening for connections on %s\n", server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Close()
		fmt.Println(err.Error())
		return 2
	}

	return 0
}

func main() {
	os.Exit(meat())
}
