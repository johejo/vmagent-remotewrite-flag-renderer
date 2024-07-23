// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"

	promconfig "github.com/prometheus/prometheus/config"
)

type PrometheusRemoteWriteConfig = promconfig.RemoteWriteConfig

type PrometheusConfig struct {
	RemoteWriteConfigs []PrometheusRemoteWriteConfig `yaml:"remote_write"`
}

func (c PrometheusConfig) String() string {
	b := newBuilder()
	for _, rwc := range c.RemoteWriteConfigs {
		if rwc.URL != nil {
			b.W("-remoteWrite.url='%s'", rwc.URL.String())
		}
		if rwc.HTTPClientConfig.BasicAuth != nil {
			b.W("-remoteWrite.basicAuth.username='%s'", rwc.HTTPClientConfig.BasicAuth.Username)
			b.W("-remoteWrite.basicAuth.password='%s'", rwc.HTTPClientConfig.BasicAuth.Password)
		}
		if rwc.HTTPClientConfig.Authorization != nil {
			if rwc.HTTPClientConfig.Authorization.Credentials != "" {
				b.W("-remoteWrite.bearerToken='%s'", rwc.HTTPClientConfig.Authorization.Credentials)
			}
			if rwc.HTTPClientConfig.Authorization.CredentialsFile != "" {
				b.W("-remoteWrite.bearerTokenFile='%s'", rwc.HTTPClientConfig.Authorization.CredentialsFile)
			}
		}
		if rwc.Headers != nil {
			for k, v := range rwc.Headers {
				b.W("-remoteWrite.headers='%s=%s'", k, v)
			}
		}
		if rwc.SigV4Config != nil {
			b.W("-remoteWrite.aws.useSigv4=true")
			if rwc.SigV4Config.AccessKey != "" {
				b.W("-remoteWrite.aws.accessKey='%s'", rwc.SigV4Config.AccessKey)
			}
			if rwc.SigV4Config.SecretKey != "" {
				b.W("-remoteWrite.aws.secretKey='%s'", rwc.SigV4Config.SecretKey)
			}
			if rwc.SigV4Config.Region != "" {
				b.W("-remoteWrite.aws.region='%s'", rwc.SigV4Config.Region)
			}
			if rwc.SigV4Config.RoleARN != "" {
				b.W("-remoteWrite.aws.roleARN='%s'", rwc.SigV4Config.RoleARN)
			}
		}
		if rwc.HTTPClientConfig.OAuth2 != nil {
			if rwc.HTTPClientConfig.OAuth2.ClientID != "" {
				b.W("-remoteWrite.oauth2.clientID='%s'", rwc.HTTPClientConfig.OAuth2.ClientID)
			}
			if rwc.HTTPClientConfig.OAuth2.ClientSecret != "" {
				b.W("-remoteWrite.oauth2.clientSecret='%s'", rwc.HTTPClientConfig.OAuth2.ClientSecret)
			}
			if rwc.HTTPClientConfig.OAuth2.ClientSecretFile != "" {
				b.W("-remoteWrite.oauth2.clientSecretFile='%s'", rwc.HTTPClientConfig.OAuth2.ClientSecretFile)
			}
			if len(rwc.HTTPClientConfig.OAuth2.Scopes) != 0 {
				b.W("-remoteWrite.oauth2.scopes='%s'", strings.Join(rwc.HTTPClientConfig.OAuth2.Scopes, ";"))
			}
			if rwc.HTTPClientConfig.OAuth2.TokenURL != "" {
				b.W("-remoteWrite.oauth2.tokenUrl='%s'", rwc.HTTPClientConfig.OAuth2.TokenURL)
			}
		}
		if rwc.HTTPClientConfig.ProxyURL.URL != nil {
			b.W("-remoteWrite.provyURL='%s'", rwc.HTTPClientConfig.ProxyURL.String())
		}
		if rwc.HTTPClientConfig.TLSConfig.CAFile != "" {
			b.W("-remoteWrite.tlsCAFile='%s'", rwc.HTTPClientConfig.TLSConfig.CAFile)
		}
		if rwc.HTTPClientConfig.TLSConfig.CertFile != "" {
			b.W("-remoteWrite.tlsCertFile='%s'", rwc.HTTPClientConfig.TLSConfig.CertFile)
		}
		if rwc.HTTPClientConfig.TLSConfig.KeyFile != "" {
			b.W("-remoteWrite.tlsKeyFile='%s'", rwc.HTTPClientConfig.TLSConfig.KeyFile)
		}
		if rwc.HTTPClientConfig.TLSConfig.ServerName != "" {
			b.W("-remoteWrite.tlsServerName='%s'", rwc.HTTPClientConfig.TLSConfig.ServerName)
		}
		if rwc.RemoteTimeout.String() != "" {
			b.W("-remoteWrite.sendTimmeout='%s'", rwc.RemoteTimeout.String())
		}
	}
	return b.String()
}
