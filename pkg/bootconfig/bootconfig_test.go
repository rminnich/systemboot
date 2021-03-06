package bootconfig

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBootConfig(t *testing.T) {
	data := []byte(`{
	"name": "some_conf",
	"kernel": "/path/to/kernel",
	"initramfs": "/path/to/initramfs",
	"kernel_args": "init=/bin/bash",
	"devicetree": "some data here"
}`)
	c, err := NewBootConfig(data)
	require.NoError(t, err)
	require.Equal(t, "some_conf", c.Name)
	require.Equal(t, "/path/to/kernel", c.Kernel)
	require.Equal(t, "/path/to/initramfs", c.Initramfs)
	require.Equal(t, "init=/bin/bash", c.KernelArgs)
	require.Equal(t, "some data here", c.DeviceTree)
	require.Equal(t, true, c.Validate())
}

func TestNewBootConfigInvalidJSON(t *testing.T) {
	data := []byte(`{
	"name": "broken
}`)
	_, err := NewBootConfig(data)
	require.Error(t, err)
}

func TestNewBootConfigMissingKernel(t *testing.T) {
	data := []byte(`{
	"name": "some_conf",
	"kernel_is_missing": "/path/to/kernel",
	"initramfs": "/path/to/initramfs",
	"kernel_args": "init=/bin/bash",
	"devicetree": "some data here"
}`)
	c, err := NewBootConfig(data)
	require.NoError(t, err)
	require.Equal(t, false, c.Validate())
}
