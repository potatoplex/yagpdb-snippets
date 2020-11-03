{{/*
	Counting Easter Eggs
	Executes set of commands for certain numbers

	author: patatas
	date: 2020/10/31

	Trigger Type: None
*/}}

{{ $input := .ExecData.Input }}
{{ $absInput := (reReplace "-" $input "") }}

{{/* Check for magic numbers */}}

{{/* 21 Guns */}}
{{ if eq $absInput "21" }}
		{{ $url := "https://cdn.discordapp.com/attachments/765047137473265714/773234394122027058/21guns.gif" }}
		{{ sendMessage nil (cembed
				"title" "One... Twety One Guns..."
				"color" (0x228B22)
				"image" (sdict "url" $url)
		) }}
{{end}}

{{/* Kobe */}}
{{ if eq $absInput "24" }}
		{{ addReactions  "🏀" "🐍" }}
		{{ $url := "https://media1.tenor.com/images/48ee91087f61fac64ca41b98f3bfa8ef/tenor.gif?itemid=8764570" }}
		{{ sendMessage nil (cembed
				"title" "KOBE!!!"
				"color" (0x8c15b9)
				"image" (sdict "url" $url)
		) }}
{{end}}

{{/* MJ */}}
{{ if eq $absInput "23" }}
		{{ addReactions  "🏀" "🐐" }}
		{{ $url := "https://media1.tenor.com/images/b054ef6ded01867028e5615cf9a9858d/tenor.gif?itemid=17298668" }}
		{{ sendMessage nil (cembed
				"title" "G. O. A. T."
				"description" "*I believe I can fly...*"
				"color" (0xC12F2F)
				"image" (sdict "url" $url)
		) }}
{{end}}

{{/* almost hehe */}}
{{ if eq $absInput "68" }}
		{{ addReactions  ":arthurFist:749784379529756684" "a:a_sumabog:760366591048417290" }}
{{end}}

{{/* hehe */}}
{{ if eq $absInput "69" }}
		{{ addReactions  ":heheBoye:754459019300700281" ":ab_nice:754459018948247554" }}
		{{ $url := "https://media1.tenor.com/images/f6c24f0de18d6540f657e6b30690c2e0/tenor.gif?itemid=3566649" }}
		{{ sendMessage nil (cembed
				"title" "NICE"
				"color" (0x60A917)
				"image" (sdict "url" $url)
		) }}
{{end}}

{{/* i love you */}}
{{ if eq $absInput "143" }}
		{{ addReactions  ":pepeKilig:764897548288655420" "❤️" }}
		{{ $url := "https://media1.tenor.com/images/4fd7ac67d3e28edcb6fb9aa89df06a66/tenor.gif?itemid=12032436" }}
		{{ sendMessage nil (cembed
				"title" "AYIIEEE"
				"color" (0xC12F2F)
				"image" (sdict "url" $url)
		) }}
{{end}}

{{/* yawa */}}
{{ if or (eq $absInput "66") (eq $absInput "666") }}
		{{ $url := "https://media1.tenor.com/images/b07609be95cee58687557221c87fff52/tenor.gif?itemid=4788073" }}
		{{ sendMessage nil (cembed
				"title" "OMG"
				"color" (0xC12F2F)
				"description" (joinStr "" .User.Mention " has summoned the **DEVIL**!")
				"image" (sdict "url" $url)
		) }}
{{ end }}

{{/* am i real? */}}
{{ if eq $absInput "214" }}
		{{ $url := "https://cdn.discordapp.com/attachments/765047137473265714/771709043244007424/roxnebres-deviantart-com-e1561078810674.png" }}
		{{ addReactions "2️⃣" "1️⃣" "4️⃣" }}
		{{ sendMessage nil (cembed
				"title" "Am I Real?"
				"color" (0xC12F2F)
				"description" "*The world would die and everything may lie...*"
				"image" (sdict "url" $url)
		) }}
{{ end }}

{{/* my favorite song */}}
{{ if eq $absInput "241" }}
		{{ $url := "https://cdn.discordapp.com/attachments/765047137473265714/771709043244007424/roxnebres-deviantart-com-e1561078810674.png" }}
		{{ addReactions "2️⃣" "4️⃣" "1️⃣" }}
		{{ sendMessage nil (cembed
				"title" "I wanted to turn you on"
				"color" (0xC12F2F)
				"description" "*My favorite song...*"
				"image" (sdict "url" $url)
		) }}
{{ end }}


{{/* 420 */}}
{{ if eq $absInput "420" }}
		{{ $url := "https://media1.tenor.com/images/76537617f27e6dee9efdc2071a88c982/tenor.gif?itemid=15415824" }}
		{{ addReactions ":pepeBaked:765389425562615889" }}
		{{ sendMessage nil (cembed
				"title" "Huff and Puff"
				"color" (0x228B22)
				"image" (sdict "url" $url)
		) }}
{{ end }}

{{/* Page not found */}}
{{ if eq $absInput "404" }}
		{{ $url := "https://cdn.discordapp.com/attachments/765047137473265714/772427313227890708/saedx-blog-featured-70.png" }}
		{{ addReactions "🚫" }}
		{{ sendMessage nil (cembed
				"title" "Page Not Found"
				"color" (0x228B22)
				"image" (sdict "url" $url)
		) }}
{{ end }}

{{/* Rey Mysterio */}}
{{ if eq $absInput "619" }}
		{{ $url := "https://media1.tenor.com/images/1484ac2b5738e2708a65d8f3b1ecf9a2/tenor.gif?itemid=13070884" }}
		{{ sendMessage nil (cembed
				"title" "Rey Mysterio 619"
				"color" (0x40e0d0)
				"image" (sdict "url" $url)
		) }}
{{ end }}