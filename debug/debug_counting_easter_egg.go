{{/*
	Counting Easter Eggs Debugger
	Simulates actual response of easter eggs

	author: patatas
	date: 2020/10/31

	Trigger Type: Command (prefix)
	Trigger: testcount (or whatever)

	Sample command: -testcount 69
*/}}

{{/* CCID of easter egg custom command */}}
{{ $easterEggCC := 2 }}

{{ $input:= toString (index .CmdArgs 0 ) }}
{{ execCC $easterEggCC nil 0 (sdict "Input" $input) }}