package lxdpkg

import (
	"fmt"

	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func Connect() (container lxd.ContainerServer) {
	container, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		fmt.Println(err)
	}
	return container
}

func Status(name string, container lxd.ContainerServer) *api.ContainerState {
	var stat *api.ContainerState

	stat, str, err := container.GetContainerState(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return stat
}

func Create(name string, container lxd.ContainerServer) (string, error) {
	resp := "[Complete] create " + name
	req := api.ContainersPost{
		Name: name,
		Source: api.ContainerSource{
			Type: "image",
			Alias: "debian",
		},
	}

	op, err := container.CreateContainer(req)
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	err = op.Wait()
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	return resp, err
}

func Delete(name string, container lxd.ContainerServer) (string, error) {
	resp := "[Complete] delete" + name
	op, err := container.DeleteContainer(name)
	if err != nil{
		resp = "[Failed]"
		return resp, err
	}
	err = op.Wait()
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	return resp, err
}

func Start(name string, container lxd.ContainerServer) (string, error) {
	resp := "[Complete] start" + name
	reqState := api.ContainerStatePut {
		Action: "start",
		Timeout: -1,
	}

	op, err := container.UpdateContainerState(name, reqState, "")
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	err = op.Wait()
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	return resp, err
}

func Stop(name string, container lxd.ContainerServer) (string, error) {
	resp := "[Complete] stop" + name
	reqState := api.ContainerStatePut {
		Action: "stop",
		Timeout: -1,
	}

	op, err := container.UpdateContainerState(name, reqState, "")
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	err = op.Wait()
	if err != nil {
		resp = "[Failed]"
		return resp, err
	}
	return resp, err
}
