package main

import (
	"strings"
	"text/template"
)

const (
	confPerm  = 0644
	confPath  = "/srv/ipsec/ipsec.conf"
	reloadCmd = "/usr/sbin/ipsec update"
)

var (
	funcMap = template.FuncMap{
		"replace": strings.Replace,
	}

	cfgTemplate = template.Must(template.New("cfg").Funcs(funcMap).Parse(
		`# IPSec configuration generated by https://github.com/appscode/swanc
# DO NOT EDIT!

config setup
        # strictcrlpolicy=yes
        # uniqueids = no

conn %default
        ikelifetime=60m
        keylife=20m
        rekeymargin=3m
        keyingtries=1
        mobike=no
        keyexchange=ikev2

{{with $ip := .HostIP}}{{range $peer_ip := .NodeIPs }}{{ if ne $ip $peer_ip }}
conn {{replace $ip "." "_" -1}}_{{replace $peer_ip "." "_" -1}}
        authby=secret
        left={{$ip}}
        right={{$peer_ip}}
        type=transport
        auto=start
        esp=aes128gcm16!
{{ end }}{{end}}{{end}}
`))
)
