package http

import (
	"io"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/gogf/gf/os/gfile"

	"github.com/gogf/gf/frame/g"
)

func Get(cityCode string) []byte {
	r, err := g.Client().Timeout(60 * time.Second).Get("https://m.weathercn.com/current-weather.do?id=" + cityCode)
	if err != nil {
		return getFromCache()
	}
	all := r.ReadAll()
	cacheToUserHomeLocal(all)
	return all
}

func getFromCache() []byte {
	u, err := user.Current()
	if err != nil {
		return []byte{}
	}
	path := filepath.Join(u.HomeDir, ".local/share/weather-cli", "cache")
	if !gfile.Exists(path) {
		return []byte{}
	}

	fi, err := os.Stat(path)
	if err != nil {
		return []byte{}
	}

	if fi.ModTime().Add(time.Hour).Before(time.Now()) {
		return []byte{}
	}

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return []byte{}
	}
	rb, err := io.ReadAll(f)
	if err != nil {
		return []byte{}
	}

	return rb
}

func cacheToUserHomeLocal(all []byte) {
	u, err := user.Current()
	if err != nil {
		return
	}
	path := filepath.Join(u.HomeDir, ".local/share/weather-cli")
	if !gfile.Exists(path) {
		if err := gfile.Mkdir(path); err != nil {
			return
		}
	}
	fp := filepath.Join(path, "cache")
	if gfile.Exists(fp) {
		if err := gfile.Remove(fp); err != nil {
			return
		}
	}

	f, err := os.Create(fp)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return
	}
	_, _ = f.Write(all)
}
