// SPDX-License-Identifier: Apache-2.0

package main

import "reflect"

type VMAgentConfig struct {
	RemoteWriteConfigs []vmagentRemoteWriteConfig `yaml:"remoteWrite" json:"remoteWrite"`
}

func (c *VMAgentConfig) String() string {
	b := newBuilder()
	for _, rwc := range c.RemoteWriteConfigs {
		c.render(b, rwc)
	}
	return b.String()
}

func (c *VMAgentConfig) render(b *builder, e any) {
	v := reflect.ValueOf(e)
	t := reflect.TypeOf(e)
	n := t.NumField()
	for i := 0; i < n; i++ {
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Struct:
			c.render(b, v.Field(i).Interface())
		default:
			fi := v.Field(i)
			if flg := ft.Tag.Get("vmagent-flag"); flg != "" && !fi.IsZero() {
				b.W("-%s='%s'", flg, fi)
			}
		}
	}
}
