{{/*
    Idea by boss tokwa
    Counting with varying increments, because +1 are for normies.

    author: daimos
    modified: patatas, tokwa
    date: 2020/10/26

    Trigger Type: Regex
    Trigger: ^-?\d+$
*/}}

{{/* Change into correct easter egg CCID */}}
{{ $easterEggCC := 5 }}

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

        {{/* Check for magic numbers */}}
        {{ execCC $easterEggCC nil 0 (sdict "Input" (toString $input)) }}
        
    {{ end }}
{{ end }}

{{ dbSet 0 "prevInput" $input}}
