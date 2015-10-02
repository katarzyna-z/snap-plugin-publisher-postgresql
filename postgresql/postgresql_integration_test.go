//
// +build integration

package postgresql

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"
	"time"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPostgresPublish(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)

	Convey("Pulse Plugin PostgreSQL integration testing with PostgreSQL", t, func() {
		var buf bytes.Buffer

		config["hostname"] = ctypes.ConfigValueStr{Value: os.Getenv("PULSE_POSTGRESQL_HOST")}
		config["port"] = ctypes.ConfigValueInt{Value: 5432}
		config["username"] = ctypes.ConfigValueStr{Value: "postgres"}
		config["password"] = ctypes.ConfigValueStr{Value: ""}
		config["database"] = ctypes.ConfigValueStr{Value: "pulse_test"}
		config["table_name"] = ctypes.ConfigValueStr{Value: "info"}

		ip := NewPostgreSQLPublisher()
		cp := ip.GetConfigPolicy()
		cfg, _ := cp.Get([]string{""}).Process(config)

		Convey("Publish integer metric", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"foo"}, time.Now(), "", 99),
			}
			buf.Reset()
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			err := ip.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish float metric", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"bar"}, time.Now(), "", 3.141),
			}
			buf.Reset()
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			err := ip.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish string metric", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"qux"}, time.Now(), "", "bar"),
			}
			buf.Reset()
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			err := ip.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish boolean metric", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"baz"}, time.Now(), "", true),
			}
			buf.Reset()
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			err := ip.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish multiple metrics", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"foo"}, time.Now(), "", 101),
				*plugin.NewPluginMetricType([]string{"bar"}, time.Now(), "", 5.789),
			}
			buf.Reset()
			enc := gob.NewEncoder(&buf)
			enc.Encode(metrics)
			err := ip.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

	})
}