package examples

import (
	"context"
	"flag"
	"fmt"
	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"testing"
)

func TestAuth(t *testing.T)  {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	data, err := ioutil.ReadFile(*kubeconfig)
	if err!=nil {
		fmt.Println(err.Error())
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {

	}

    client,_:=k8s.NewClient(&config)

	dataCm,_:=ioutil.ReadFile("test-cm.yaml")
	res,_:=yaml.YAMLToJSON(dataCm)
	fmt.Println(string(res))

	client.DoSend(context.TODO(),"PUT","https://10.5.1.111:6443/api/v1/namespaces/test/configmaps/game-demo",res,nil)
	//cm:=`{"apiVersion":"v1","data":{"auth":"test"},"kind":"ConfigMap","metadata":{"name":"game-demo","namespace":"test"}}`

	client.DoSend(context.TODO(),"DELETE","https://10.5.1.111:6443/api/v1/namespaces/test/configmaps",res,nil)

	resource,_:=ioutil.ReadFile("test-deploy.yaml")
	result,_:=yaml.YAMLToJSON(resource)
	fmt.Println(string(result))

	ers:=client.DoSend(context.TODO(),"DELETE","https://10.5.1.111:6443/apis/apps/v1/namespaces/test/deployments/nginx-deployment",nil,nil)


	if ers!=nil {
		fmt.Println(ers.Error())
	}



}