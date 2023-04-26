package main

func main() {
	events := []Event{
		{
			EventType:        "country_event",
			Picture:          "interesting_histories_southern_india",
			Hidden:           false,
			FireOnlyOnce:     false,
			GotoLocation:     "",
			LeftPortrait:     "",
			RightPortrait:    "",
			WeightMultiplier: "",
			Trigger:          "",
			Immediate:        "",
			Options: []Option{
				{
					Effects: "\t\tcurrent_ruler = {\n" +
						"\t\t\tadd_charisma = 2\n" +
						"\t\t\tadd_charisma = 2\n" +
						"\t\t\tadd_zeal = 2\n" +
						"\t\t\tadd_finesse = 2\n" +
						"\t\t\tadd_martial = 2\n" +
						"\t\t\tadd_popularity = 25\n" +
						"\t\t\tadd_nickname = NICKNAME_THE_GREAT\n" +
						"\t\t}\n\n" +
						"\t\tadd_country_modifier = {\n" +
						"\t\t\tname = restored_babylon\n" +
						"\t\t\tduration = 3650\n" +
						"\t\t}\n\n" +
						"\t\tadd_innovation = 4\n" +
						"\t\tadd_4_free_province_investments = yes\n" +
						"\t\tadd_political_influence = 50\n" +
						"\t\tadd_treasury = 350\n\n" +
						"\t\thidden_effect = {\n" +
						"\t\t\tif = {\n" +
						"\t\t\t\tlimit = {\n" +
						"\t\t\t\t\thas_variable = formed_babylon_inv\n" +
						"\t\t\t\t}\n" +
						"\t\t\t\tremove_variable = formed_babylon_inv\n" +
						"\t\t\t}\n" +
						"\t\t}\n",
				},
			},
			After: "",
		},
		{
			EventType:        "country_event",
			Picture:          "interesting_histories_southern_india",
			Hidden:           false,
			FireOnlyOnce:     false,
			GotoLocation:     "",
			LeftPortrait:     "current_ruler",
			RightPortrait:    "",
			WeightMultiplier: "",
			Trigger:          "",
			Immediate:        "",
			Options: []Option{
				{
					Effects: "\t\tadd_political_influence = -65\n" +
						"\t\tchange_government = babylonian_bureaucracy\n",
				},
				{
					Effects: "\t\tadd_political_influence = 35\n",
				},
			},
			After: "",
		},
		{
			EventType:        "country_event",
			Picture:          "interesting_histories_southern_india",
			Hidden:           false,
			FireOnlyOnce:     false,
			GotoLocation:     "",
			LeftPortrait:     "current_ruler",
			RightPortrait:    "",
			WeightMultiplier: "",
			Trigger:          "",
			Immediate: "immediate = {\n" +
				"		custom_tooltip = me_tylos_3_i_tt\n" +
				"		hidden_effect = {\n" +
				"			area:babylonia_area = { every_area_province = { add_claim = root } }\n" +
				"		}\n" +
				"		p:918.owner = {\n" +
				"			add_opinion = { modifier = declare_babylon_intent_realtion_tylos target = root }\n" +
				"			if = {\n" +
				"				limit = {\n" +
				"					root = {\n" +
				"						NOT = { is_subject_of = prev }\n" +
				"					}\n" +
				"				}\n" +
				"				release_subject = root\n" +
				"			}\n" +
				"		}\n" +
				"		random_country_culture = {\n" +
				"			limit = {\n" +
				"				is_culture = babylonian\n" +
				"			}\n" +
				"			add_country_culture_modifier = {\n" +
				"				name = memory_of_babylon\n" +
				"				duration = 9125\n" +
				"			}\n" +
				"		}\n" +
				"	}",
			Options: []Option{
				{
					Effects: "\t\tcustom_tooltip = me_tylos_3_a_tt\n" +
						"\t\thidden_effect = {\n" +
						"\t\t\tarea:beth_daraye_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t\tarea:asoristan_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t\tarea:tigris_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t}\n" +
						"\t\tadd_aggressive_expansion = 15\n" +
						"\t\tadd_tyranny = 10\n" +
						"\t\tadd_political_influence = -30\n",
				},
				{
					Effects: "\t\tcustom_tooltip = me_tylos_3_a_tt\n" +
						"\t\thidden_effect = {\n" +
						"\t\t\tarea:beth_daraye_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t\tarea:asoristan_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t\tarea:tigris_area = { every_area_province = { add_claim = root } }\n" +
						"\t\t}\n" +
						"\t\tadd_aggressive_expansion = 15\n" +
						"\t\tadd_tyranny = 10\n" +
						"\t\tadd_political_influence = -30\n",
				},
			},
			After: "",
		},
	}
	eventFile := EventFile{"", "", "me_tylos_test"}
	eventFile.makeFile(events)
	eventFile.makeLocFile(events)
}
