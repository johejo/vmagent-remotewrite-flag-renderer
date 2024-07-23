package main

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	for _, test := range []struct {
		config string
		format string
		want   string
	}{
		{
			config: "./testdata/prometheus.yaml",
			format: "prometheus",
			want:   "-remoteWrite.url='http://localhost:9000' -remoteWrite.basicAuth.username='foo' -remoteWrite.basicAuth.password='bar' -remoteWrite.sendTimmeout='30s' -remoteWrite.url='http://localhost:9001' -remoteWrite.basicAuth.username='hello' -remoteWrite.basicAuth.password='world' -remoteWrite.sendTimmeout='30s'",
		},
		{
			config: "./testdata/vmagent.yaml",
			format: "vmagent",
			want:   "-remoteWrite.url='http://localhost:9000' -remoteWrite.basicAuth.username='foo' -remoteWrite.basicAuth.password='hello' -remoteWrite.url='http://localhost:9001' -remoteWrite.basicAuth.username='bar' -remoteWrite.basicAuth.password='world'",
		},
		{
			config: "./testdata/vmagent.json",
			format: "vmagent",
			want:   "-remoteWrite.url='http://localhost:9000' -remoteWrite.basicAuth.username='foo' -remoteWrite.basicAuth.password='hello' -remoteWrite.url='http://localhost:9001' -remoteWrite.basicAuth.username='bar' -remoteWrite.basicAuth.password='world'",
		},
	} {
		t.Run(test.config, func(t *testing.T) {
			got, err := run(test.config, test.format)
			if err != nil {
				t.Fatal(err)
			}
			s1 := strings.Split(got, " ")
			slices.Sort(s1)
			s2 := strings.Split(test.want, " ")
			slices.Sort(s2)
			if !reflect.DeepEqual(s1, s2) {
				t.Fatalf("\nwant=%s\n got=%s", strings.Join(s2, " "), strings.Join(s1, " "))
			}
		})
	}
}
