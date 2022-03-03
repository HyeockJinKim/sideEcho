package config

import (
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadFlags(t *testing.T) {
	// 1. env variable 설정
	t.Setenv(authUsername, "test-username")
	t.Setenv(authPassword, "test-password")
	flags := ReadFlags()

	assert.Equal(t, []byte("test-username"), flags.Auth.Username)
	assert.Equal(t, []byte("test-password"), flags.Auth.Password)
}

func Test_ReadFile(t *testing.T) {
	t.Parallel()

	t.Run("Success Case - config 파일에서 config를 읽으면 입력한 flags와 함께 반환합니다.", func(t *testing.T) {
		t.Parallel()
		_, file, _, _ := runtime.Caller(0)
		directory := path.Join(path.Dir(file))
		configFile := path.Join(path.Dir(directory), defaultConfigPath)

		flags := &Flags{
			ConfigPath: configFile,
		}
		config, err := ReadConfigFile(flags)
		assert.NoError(t, err)
		assert.True(t, assert.ObjectsAreEqual(*flags, config.Flags))
		assert.True(t, assert.ObjectsAreEqual(ServerConfig{
			Api: Api{
				Port:               8080,
				InternalPort:       8090,
				TimeoutSec:         2,
				ShutDownTimeoutSec: 5,
			},
		}, config.Config))
	})

	t.Run("Failure Case - config 파일이 없다면 에러를 반환합니다.", func(t *testing.T) {
		t.Parallel()
		_, file, _, _ := runtime.Caller(0)
		directory := path.Join(path.Dir(file))
		configFile := path.Join(path.Dir(directory), "no-config-path.toml")

		flags := &Flags{
			ConfigPath: configFile,
		}
		_, err := ReadConfigFile(flags)
		assert.Error(t, err)
	})
}
