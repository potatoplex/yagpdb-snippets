{{/*
    Idea by boss tokwa
    Counting with varying increments, because +1 are for normies.

    author: daimos
    modified: patatas, tokwa
    date: 2020/10/26
*/}}

{{/* Get value from DB */}}
{{ $tofuNum := (dbGet 0 "tofuNum").Value}}
{{ $streak := (dbGet 0 "streak").Value}}
{{ $prevInput := toInt (dbGet 0 "prevInput").Value}}
{{ $failed := false }}
{{ $newRound := false }}
{{ $absTofuNum := (reReplace "-" ($tofuNum | toString) "") }}


{{/* Get user input */}}
{{ $input := toInt .Message.Content }}

{{/* ignore duplicate inputs */}}
{{ if ne $prevInput $input }}

    {{/* Get expected running value */}}
    {{ $expected := dbIncr 0 "expected" $tofuNum | toInt }}

    {{/* check if NOT EQUAL to expected value */}}
    {{ if ne $expected $input }}
        {{ $failed = true }}
        {{ $newRound = true }}
    {{ else }}
        {{/* Round Randomizer */}}
        {{$rate := 10}}
        {{$roll := randInt 100}}
        {{if lt $roll $rate}}
            {{ $failed = false }}
            {{ $newRound = true }}
        {{ end }}
    {{ end }}
    {{/* Create new round rules */}}
        {{if $newRound }}
        {{/* Initialization */}}
        {{ $tofuPrev := $tofuNum | toInt  }}
        {{ $tofuNum = randInt 10 | add 1 }}
        {{ $operator := "" }}
        {{ $absTofuNum := (reReplace "-" ($tofuNum | toString) "") }}

        {{/* Addition Block */}}
        {{if eq ($fun := randInt -1 1) 0}}
            {{ $operator = "+" }}
        {{/* Subtraction Block, No need to specify operator on signed number */}}
        {{else if lt $fun 0}}
            {{ $operator = "-" }}
            {{ $tofuNum = mult $tofuNum -1}}
            {{ $absTofuNum := (reReplace "-" ($tofuNum | toString) "") }}
        {{end}}

        {{/* Save new round rules in DB */}}
        {{ dbSet 0 "tofuNum" $tofuNum}}

        {{if $failed }}
            {{/* Restart on fail */}}
            {{ $restartNum := randInt 10 | add 1}}
            {{ dbSet 0 "expected" $restartNum}}
            

            {{/* Notify the imbecile */}}

            {{ $embed := cembed
            "title" "Ennggkk..."
            "color" (0x3498DB)
            "description" (joinStr ""  .User.Mention "  failed! correct sequence is " $expected "\nNew Rule `n " $operator " " $absTofuNum "` Start with **" $restartNum "**")
            "footer" (sdict "text" (joinStr "" "Streak: " $streak) "icon_url" "")
            }}
            {{ sendMessage nil $embed }}
        {{else if (ne $tofuPrev $tofuNum) }}
            
            {{ sendMessage nil (cembed
                "title" "New rule!!!"
                "color" (0x60A917)
                "description" (joinStr "" "`n " $operator " " $absTofuNum  "`")
            ) }}
        {{end}}

    {{end}} 

    {{/* Reset streak to 0 on fail */}}
    {{ if $failed }}
        {{ dbSet 0 "streak" 0}}
    {{ else }}
        {{/* Increment streak by 1 for each correct answer */}}
        {{ $streak = dbIncr 0 "streak" 1 }}
        {{ $absInput := (reReplace "-" ($input | toString) "") }}

        {{/* Check for magic numbers */}}

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
        
    {{ end }}
{{ end }}

{{ dbSet 0 "prevInput" $input}}
