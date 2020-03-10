{{/* Slot Machine Game */}}

{{/* USER VARIABLES */}}
{{$dbName := "CREDITOS"}} {{/* Name of the Key of your DB that stores users currency ammount */}}
{{$gameName := "Máquina da Sorte"}} {{/* Whatever you want the game to be named */}}
{{$user := "Usuário"}} {{/* How should the user be called. For example: "player" or "user" */}}
{{$spinName := "RODANDO"}} {{/* Word to show user that slot machine is currently spinning */}}
{{$lose := "Você perdeu :("}} {{/* Text to tell user he lost */}}
{{$win := "VOCÊ GANHOU!"}} {{/* Text to tell user he won */}}
{{$profit := "Lucro"}} {{/* How should the profit be called */}}
{{$currency := "Créditos"}} {{/* Name of the currency in your server */}}
{{$payOut := "Pagamentos"}} {{/* Name of the currency in your server */}}
{{$youHave := "Você tem "}} {{/* "You have" in your language */}}
{{$helper := "Modo de uso"}} {{/* Helper text title */}}
{{$helpText := "-apostar <quantidade>\nPor exemplo: **-apostar 10**\nAssim você estaria apostando 10 créditos."}} {{/* Your helper text */}}
{{$notEnough := "Você não tem créditos suficientes"}} {{/* Error msg when user doesnt have enough credits to place bet */}}
{{$betBelow1 := "Você só pode fazer apostas de, no mínimo, 1 crédito"}} {{/* Error msg when user try to bet 0 */}}
{{$bettingChannel := 655082852295376922}} {{/* Channel users can player */}}
{{$channels := cslice
	655082852295376922
	682204005723799553
	683859835304804394
	655393004953141251
	655650583818010645
}} {{/* IDs of different channels in your server to prevent the game from lagging */}}
{{/* END USER VARIABLES */}}

{{/* ACTUAL CODE! DON'T TOUCH! */}}
{{$header := (joinStr "" $gameName " | " $user ": " .User.Username)}}
{{$slotEmoji := "<a:slotmoney:686445052284895237>"}}
{{$g := 65280}}{{$y := 16776960}}{{$r := 16711680}}{{$b := 65534}}
{{$emojis := cslice "🥇" "🥇" "🥇" "🥇" "🥇" "🥇" "🥇" "💎" "💎" "💎" "💎" "💎" "💎" "💯" "💯" "💯" "💯" "💵" "💵" "💵" "💰" "💰"}}
{{$choosen := index (shuffle $emojis) 0}}
{{$choosen2 := index (shuffle $emojis) 0}}
{{$choosen3 := index (shuffle $emojis) 0}}
{{$bal := (toInt (dbGet .User.ID $dbName).Value)}}
{{$embed := sdict "color" $g "fields" (cslice (sdict "name" $header "value" (joinStr "" "**-------------------\n| " $slotEmoji " | " $slotEmoji " | " $slotEmoji " |\n-------------------\n- " $spinName " -**") "inline" false))}}

{{if (and (not .ExecData) (eq .Channel.ID $bettingChannel) (not (dbGet .User.ID "block_slot_123456")))}}
	{{if .CmdArgs}}
		{{$bet := toInt (index .CmdArgs 0)}}
		{{if ge $bet 1}}
			{{if ge $bal $bet}}
				{{dbSet .User.ID "block_slot_123456" true}}
				{{$silent := dbIncr .User.ID $dbName (mult -1 $bet)}}
				{{$id := sendMessageRetID nil (cembed $embed)}}
				{{execCC .CCID (index (shuffle $channels) 0) 2 (sdict "depth" 1 "id" $id "bet" $bet)}}
			{{else}}
				{{joinStr "" $notEnough ", " .User.Mention "!"}}
			{{end}}
		{{else}}
			{{joinStr "" $betBelow1 ", " .User.Mention "!"}}
		{{end}}
	{{else}}
		{{$embedHelp := (cembed
			"title" $gameName
			"fields" (cslice (sdict "name" $payOut "value" (joinStr "" "**🥇🥇❓ - 1x\n💎💎❓ - 2x\n💯💯❓ - 3x\n🥇🥇🥇 - 3x\n💎💎💎 - 4x\n💵💵❓ - 4x\n💯💯💯 - 5x\n💰💰❓ - 5x\n💵💵💵 - 10x\n💰💰💰 - 15x**") "inline" false) (sdict "name" $helper "value" $helpText "inline" false))
			"color" $y
		)}}
		{{sendMessage nil $embedHelp}}
	{{end}}
{{end}}

{{with .ExecData}}
	{{if eq .depth 1}}
		{{$embed.Set "fields" (cslice (sdict "name" $header "value" (joinStr "" "**-------------------\n| " $choosen " | " $slotEmoji " | " $slotEmoji " |\n-------------------\n- " $spinName " -**") "inline" false))}}
		{{editMessage $bettingChannel .id (cembed $embed)}}
		{{execCC $.CCID (index (shuffle $channels) 0) 1 (sdict "depth" 2 "id" .id "choosen" $choosen "bet" .bet)}}
	{{else if eq .depth 2}}
		{{$embed.Set "fields" (cslice (sdict "name" $header "value" (joinStr "" "**-------------------\n| " .choosen " | " $choosen2 " | " $slotEmoji " |\n-------------------\n- " $spinName " -**") "inline" false))}}
		{{editMessage $bettingChannel .id (cembed $embed)}}
		{{execCC $.CCID (index (shuffle $channels) 0) 1 (sdict "depth" 3 "id" .id "choosen" .choosen "choosen2" $choosen2 "bet" .bet)}}
	{{else if eq .depth 3}}
		{{$announce := $lose}}
		{{$multiplier := 1}}
		{{if (and (eq .choosen "💎") (eq .choosen2 "💎") (ne $choosen3 "💎"))}}
			{{$multiplier = 2}}
		{{else if (or (and (eq .choosen "🥇") (eq .choosen2 "🥇") (eq $choosen3 "🥇")) (and (eq .choosen "💯") (eq .choosen2 "💯") (ne $choosen3 "💯")))}}
			{{$multiplier = 3}}
		{{else if (or (and (eq .choosen "💎") (eq .choosen2 "💎") (eq $choosen3 "💎")) (and (eq .choosen "💵") (eq .choosen2 "💵") (ne $choosen3 "💵")))}}
			{{$multiplier = 4}}
		{{else if (or (and (eq .choosen "💯") (eq .choosen2 "💯") (eq $choosen3 "💯")) (and (eq .choosen "💰") (eq .choosen2 "💰") (ne $choosen3 "💰")))}}
			{{$multiplier = 5}}
		{{else if (and (eq .choosen "💵") (eq .choosen2 "💵") (eq $choosen3 "💵"))}}
			{{$multiplier = 10}}
		{{else if (and (eq .choosen "💰") (eq .choosen2 "💰") (eq $choosen3 "💰"))}}
			{{$multiplier = 15}}
		{{end}}
		{{$pag1 := (sdict "name" $profit "value" (joinStr "" "**-" .bet " " (lower $currency) "**") "inline" true)}}
		{{$c := $r}}
		{{if eq .choosen .choosen2}}
			{{$c = $b}}
			{{$announce = $win}}
			{{$pag1 = (sdict "name" $profit "value" (joinStr "" "**" (mult .bet $multiplier) " " (lower $currency) "**") "inline" true)}}
			{{$silent2 := dbIncr $.User.ID $dbName (mult .bet $multiplier)}}
		{{end}}
		{{$embed.Set "fields" (cslice (sdict "name" $header "value" (joinStr "" "**-------------------\n| " .choosen " | " .choosen2 " | " $choosen3 " |\n-------------------\n" $announce "**") "inline" false))}}
		{{$embed.Set "color" $c}}
		{{$embed.Set "fields" ($embed.fields.Append $pag1)}}
		{{$saldo := (toInt (dbGet $.User.ID $dbName).Value)}}
		{{$pag2 := (sdict "name" $currency "value" (joinStr "" $youHave " **" $saldo " " (lower $currency) "**") "inline" true)}}
		{{$embed.Set "fields" ($embed.fields.Append $pag2)}}
		{{editMessage $bettingChannel .id (cembed $embed)}}
		{{dbDel $.User.ID "block_slot_123456"}}
	{{end}}
{{end}}