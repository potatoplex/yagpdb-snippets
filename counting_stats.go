{{/*
	Counting Stats
	Display Counting Game Stats

	author: patatas
	date: 2020/11/04

	Trigger Type: Command
	Trigger: countstats
*/}}

{{ $highestStreak := (dbGet 0 "count_highestStreak").Value }}
{{ $highestInput := (dbGet 0 "count_highestInput").Value }}
{{ $lowestInput := (dbGet 0 "count_lowestInput").Value }}
{{ $a :=  (joinStr "" $highestStreak "") }}
{{ $b :=  (joinStr "" $highestInput "") }}
{{ $c :=  (joinStr "" $lowestInput "") }}

{{ sendMessage nil (cembed
	"title" "Counting Stats"
	"color" (0x60A917)
	"fields" (cslice 
		(sdict "name" "Highest Streak" "value" $a "inline" false) 
		(sdict "name" "Highest Correct Answer" "value" $b "inline" false) 
		(sdict "name" "Lowest Correct Answer" "value" $c "inline" false) 
	)
) }}