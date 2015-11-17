package apibase

import (
	"flag"
	"gopkg.in/macaron.v1"
	"io/ioutil"
)

var mode *string = flag.String("m", "dev", "Use -m <dev/product>")

type ApiServer struct {
	macaron.Macaron
	Config *ApiBaseConfig
}

func Api() *ApiServer {
	flag.Parse()
	var configPath string
	if *mode == "dev" {
		configPath = "./conf/config.json"
		macaron.Env=macaron.DEV
	} else if *mode == "product" {
		macaron.Env=macaron.PROD
		configPath = "./conf/config.product.json"
	}
	apiServer:= CustomApi(configPath)
	return apiServer
}
func CustomApi(configFilepath string) *ApiServer {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Renderer())
	apiServer := &ApiServer{*m, nil}
	apiServer.loadConfig(configFilepath)
	return apiServer
}

//加载配置文件
func (this *ApiServer) loadConfig(filename string) (err error) {

	data, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		return readErr
	}

	if this.Config, err = TransferConfig(data); err != nil {
		return err
	}

	return nil
}
