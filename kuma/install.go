package kuma

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mgfeller/common-adapter-library/adapter"
)

// MeshInstance holds the information of the instance of the mesh
type MeshInstance struct {
	InstallMode     string `json:"installmode,omitempty"`
	InstallPlatform string `json:"installplatform,omitempty"`
	InstallZone     string `json:"installzone,omitempty"`
	InstallVersion  string `json:"installversion,omitempty"`
	MgmtAddr        string `json:"mgmtaddr,omitempty"`
	Kumaaddr        string `json:"kumaaddr,omitempty"`
}

// CreateInstance installs and creates a mesh environment up and running
func (h *KumaAdapter) installKuma(del bool, version string) (string, error) {
	status := "installing"

	if del {
		status = "removing"
	}

	meshinstance := &MeshInstance{
		InstallVersion: version,
	}
	err := h.Config.MeshInstance(meshinstance)
	if err != nil {
		return status, adapter.ErrMeshConfig(err)
	}

	h.Log.Info("Installing Kuma")
	err = meshinstance.installUsingKumactl(del)
	if err != nil {
		h.Log.Err("Kuma installation failed", adapter.ErrInstallMesh(err).Error())
		return status, adapter.ErrInstallMesh(err)
	}

	h.Log.Info("Port forwarding")
	err = meshinstance.portForward()
	if err != nil {
		h.Log.Err("Kuma portforwarding failed", adapter.ErrPortForward(err).Error())
		return status, adapter.ErrPortForward(err)
	}

	return "deployed", nil
}

// installSampleApp installs and creates a sample bookinfo application up and running
func (h *KumaAdapter) installSampleApp(name string) (string, error) {
	// Needs implementation
	return "deployed", nil
}

// installMesh installs the mesh in the cluster or the target location
func (m *MeshInstance) installUsingKumactl(del bool) error {

	Executable, err := exec.LookPath("./scripts/kuma/installer.sh")
	if err != nil {
		return err
	}

	if del {
		Executable, err = exec.LookPath("./scripts/kuma/delete.sh")
		if err != nil {
			return err
		}
		return nil
	}

	cmd := &exec.Cmd{
		Path:   Executable,
		Args:   []string{Executable},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("KUMA_VERSION=%s", m.InstallVersion),
		fmt.Sprintf("KUMA_MODE=%s", m.InstallMode),
		fmt.Sprintf("KUMA_PLATFORM=%s", m.InstallPlatform),
		fmt.Sprintf("KUMA_ZONE=%s", m.InstallZone),
	)

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (m *MeshInstance) portForward() error {
	// Needs implementation
	return nil
}
