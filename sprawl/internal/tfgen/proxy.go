package tfgen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/hashicorp/consul-topology/topology"
	"github.com/hashicorp/consul-topology/util"
)

const proxyInternalPort = 80

func (g *Generator) writeNginxConfig(net *topology.Network) (bool, string, error) {
	rootdir := filepath.Join(g.workdir, "terraform", "nginx-config-"+net.Name)
	if err := os.MkdirAll(rootdir, 0755); err != nil {
		return false, "", err
	}

	configFile := filepath.Join(rootdir, "nginx.conf")

	body := fmt.Sprintf(`
server {
    listen       %d;

    location / {
        resolver 8.8.8.8;
        proxy_pass http://$http_host$uri$is_args$args;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
`, proxyInternalPort)

	_, err := UpdateFileIfDifferent(
		g.logger,
		[]byte(body),
		configFile,
		0644,
	)
	if err != nil {
		return false, "", fmt.Errorf("error writing %q: %w", configFile, err)
	}

	hash, err := util.HashFile(configFile)
	if err != nil {
		return false, "", fmt.Errorf("error hashing %q: %w", configFile, err)
	}

	return true, hash, err
}

func (g *Generator) getForwardProxyContainer(
	net *topology.Network,
	ipAddress string,
	hash string,
) Resource {
	env := []string{"HASH_FILE_VALUE=" + hash}
	proxy := struct {
		Name              string
		DockerNetworkName string
		InternalPort      int
		IPAddress         string
		Env               []string
	}{
		Name:              net.Name,
		DockerNetworkName: net.DockerName,
		InternalPort:      proxyInternalPort,
		IPAddress:         ipAddress,
		Env:               env,
	}

	return Eval(tfForwardProxyT, &proxy)
}

var tfForwardProxyT = template.Must(template.ParseFS(content, "templates/container-proxy.tf.tmpl"))
