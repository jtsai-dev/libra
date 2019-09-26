package conf

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	bootstrapConf = "conf/bootstrap.yaml"
	localConf     = "conf/app.yaml"
)

type NacosYaml struct {
	Nacos Nacos `yaml:"Nacos"`
}

var Configs Config
var nacos Nacos

// Setup initialize the configuration instance
func Setup() {
	yaml := new(NacosYaml)
	err := loadConfig(bootstrapConf, yaml)
	if err != nil {
		panic(err)
	}
	err = loadFromServer(yaml.Nacos)
	if err != nil {
		log.Fatalln(err)
	}

	if &Configs == nil {
		err = loadConfig(localConf, &Configs)
		if err != nil {
			panic(err)
		}
	}
}

func loadConfig(path string, data interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, data)
	if err != nil {
		return err
	}

	return nil
}

func loadFromServer(nacos Nacos) (err error) {
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: nacos.ServerAddr,
				Port:   nacos.ServerPort,
			},
		},
		"clientConfig": constant.ClientConfig{
			NamespaceId:         nacos.Tenant,
			TimeoutMs:           nacos.TimeoutMs,
			ListenInterval:      nacos.ListenInterval,
			LogDir:              nacos.LogDir,
			NotLoadCacheAtStart: true,
		},
	})
	if err != nil {
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
	})
	if err != nil {
		return
	}
	yaml.Unmarshal([]byte(content), &Configs)

	configClient.ListenConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			yaml.Unmarshal([]byte(data), &Configs)
		},
	})

	return nil
}
