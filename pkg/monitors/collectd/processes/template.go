//go:build linux
// +build linux

package processes

// AUTOGENERATED BY scripts/collectd-template-to-go.  DO NOT EDIT!!

import (
	"text/template"

	"github.com/signalfx/signalfx-agent/pkg/monitors/collectd"
)

// CollectdTemplate is a template for a processes collectd config file
var CollectdTemplate = template.Must(collectd.InjectTemplateFuncs(template.New("processes")).Parse(`
<LoadPlugin "processes">
  Interval {{.IntervalSeconds}}
</LoadPlugin>
<Plugin processes>
  {{range .Processes -}}
  Process "{{.}}"
  {{end}}
  {{range $name, $regex := .ProcessMatch -}}
  ProcessMatch "{{$name}}" "{{$regex}}"
  {{end}}
  CollectContextSwitch {{.CollectContextSwitch}}
  {{with .ProcFSPath -}}
  ProcFSPath "{{.}}"
  {{- end}}
</Plugin>

<Chain "PostCache"> 
  <Rule "set_processes_monitor_id"> 
    <Match "regex"> 
      Plugin "^processes$" 
    </Match> 
    <Target "set"> 
      MetaData "monitorID" "{{.MonitorID}}" 
    </Target> 
  </Rule> 
</Chain>
`)).Option("missingkey=error")
