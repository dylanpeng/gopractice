package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	caCert     string = "./certs/ca.crt"
	serverCert string = "./certs/server.crt"
	serverKey  string = "./certs/server.key"
	clientCert string = "./certs/client.crt"
	clientKey  string = "./certs/client.key"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/test", func(c *gin.Context) {
		reqHost := c.Request.Host
		host, _, err := net.SplitHostPort(reqHost)
		if err != nil {
			host = reqHost
		}
		fmt.Printf("Client domain name: %s\n ｜ ip: %s\n", host, c.ClientIP())

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	// 客户端CAPool
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caCert)
	if err != nil {
		fmt.Printf("load ca err: %s", err)
		return
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		fmt.Printf("certpool append ca fail.")
		return
	}

	go func() {
		// 可以直接用注释的代码代替最后两行
		//router.RunTLS("0.0.0.0:10679", "./cert/server.cer", "./cert/server.key")
		server := &http.Server{
			Addr:    "server.com:10679",
			Handler: router,
			TLSConfig: &tls.Config{
				//VerifyPeerCertificate: verifyCertificate,
				ClientAuth: tls.RequireAndVerifyClientCert,
				MinVersion: tls.VersionTLS12,
				//这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
				//RootCAs:    certPool,
				ClientCAs: certPool,
			},
		}

		_ = server.ListenAndServeTLS(serverCert, serverKey)
	}()

	// 调用服务
	httpsClientCall()

	// waitting for exit signal
	exit := make(chan os.Signal, 1)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'", sign)

	fmt.Printf("http server stoped")
}

func httpsClientCall() {
	// 服务端CAPool
	pool := x509.NewCertPool()
	caCrt, err := os.ReadFile(caCert)

	if err != nil {
		log.Fatal("read ca.crt file error:", err.Error())
	}

	pool.AppendCertsFromPEM(caCrt)

	// 客户端自己的证书和私钥用于认证
	cliCrt, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatalln("LoadX509KeyPair error:", err.Error())
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			//这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
			//ClientCAs:    pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://server.com:10679/test")
	if err != nil {
		fmt.Printf("get failed. | err: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func verifyCertificate(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	for _, rawCert := range rawCerts {
		cert, err := x509.ParseCertificate(rawCert)
		if err != nil {
			return err
		}
		// Check if the Common Name in the certificate matches the expected value
		if cert.Subject.CommonName != "server.com" {
			return fmt.Errorf("Unexpected Common Name: %s", cert.Subject.CommonName)
		}
	}
	return nil
}
