package haproxy

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/renameio"

	"github.com/haproxytech/kubernetes-ingress/pkg/annotations"
	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/api"
	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/config"
	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/process"
	"github.com/haproxytech/kubernetes-ingress/pkg/haproxy/rules"
	"github.com/haproxytech/kubernetes-ingress/pkg/utils"
)

var logger = utils.GetLogger()

// Instance describes and controls a HAProxy Instance
type HAProxy struct {
	api.HAProxyClient
	process.Process
	config.Env
	*config.Config
}

func New(osArgs utils.OSArgs, env config.Env, cfgFile []byte, p process.Process, client api.HAProxyClient, rulesManager rules.Rules) (h HAProxy, err error) {
	err = (&env).Init(osArgs)
	if err != nil {
		err = fmt.Errorf("failed to initialize haproxy environment: %w", err)
		return
	}
	h.Env = env

	err = renameio.WriteFile(h.MainCFGFile, cfgFile, 0755)
	if err != nil {
		err = fmt.Errorf("failed to write haproxy config file: %w", err)
		return
	}

	h.Config, err = config.New(h.Env, rulesManager)
	if err != nil {
		err = fmt.Errorf("failed to initialize haproxy config state: %w", err)
		return
	}
	if client == nil {
		h.HAProxyClient, err = api.New(h.CfgDir, h.MainCFGFile, h.Binary, h.RuntimeSocket)
		if err != nil {
			err = fmt.Errorf("failed to initialize haproxy API client: %w", err)
			return
		}
	}
	if p == nil {
		h.Process = process.New(h.Env, osArgs, h.AuxCFGFile, h.HAProxyClient)
	}
	if !osArgs.Test {
		logVersion(h.Binary)
	}
	return
}

func (h *HAProxy) Refresh(cleanCrts bool) (reload bool, err error) {
	// Certs
	if cleanCrts {
		reload = h.RefreshCerts()
	}
	// Rules
	reload = h.RefreshRules(h.HAProxyClient) || reload
	// Maps
	reload = h.RefreshMaps(h.HAProxyClient) || reload
	// Backends
	deleted, err := h.RefreshBackends()
	logger.Error(err)
	for _, backend := range deleted {
		logger.Debugf("Backend '%s' deleted", backend)
		annotations.RemoveBackendCfgSnippet(backend)
	}
	return
}

func logVersion(program string) {
	//nolint:gosec //checks of HAProxyBinary should be done in Env.Init() .
	cmd := exec.Command(program, "-v")
	res, errExec := cmd.Output()
	if errExec != nil {
		logger.Errorf("unable to get haproxy version: %s", errExec)
		return
	}
	haproxyInfo := strings.Split(string(res), "\n")
	logger.Printf("Running with %s", haproxyInfo[0])
}
