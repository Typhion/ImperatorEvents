package main

func main() {
	events := []Event{
		{
			Namespace:        "me_tylos_test",
			EventId:          0,
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
			Option: "option = {\n" +
				"		name = me_tylos_1.a\n" +
				"		current_ruler = {\n" +
				"			add_charisma = 2\n" +
				"			add_zeal = 2\n" +
				"			add_finesse = 2\n" +
				"			add_martial = 2\n" +
				"			add_popularity = 25\n" +
				"			add_nickname = NICKNAME_THE_GREAT\n" +
				"		}\n\n" +
				"		add_country_modifier = {\n" +
				"			name = restored_babylon\n" +
				"			duration = 3650\n" +
				"		}\n\n" +
				"		add_innovation = 4\n" +
				"		add_4_free_province_investments = yes\n" +
				"		add_political_influence = 50\n" +
				"		add_treasury = 350\n\n" +
				"		hidden_effect = {\n" +
				"			if = {\n" +
				"				limit = {\n" +
				"					has_variable = formed_babylon_inv\n" +
				"				}\n" +
				"				remove_variable = formed_babylon_inv\n" +
				"			}\n" +
				"		}\n" +
				"	}",
			Option2: "",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
		{
			Namespace:        "me_tylos_test",
			EventId:          0,
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
			Option: "option = {\n" +
				"		name = me_tylos_2.a\n" +
				"		add_political_influence = -65\n" +
				"		change_government = babylonian_bureaucracy\n" +
				"	}",
			Option2: "option = {\n" +
				"		name = me_tylos_2.b\n" +
				"		add_political_influence = 35\n" +
				"	}",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
		{
			Namespace:        "me_tylos_test",
			EventId:          0,
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
			Option: "option = {\n" +
				"		name = me_tylos_3.a\n" +
				"		custom_tooltip = me_tylos_3_a_tt\n" +
				"		hidden_effect = {\n" +
				"			area:beth_daraye_area = { every_area_province = { add_claim = root } }\n" +
				"			area:asoristan_area = { every_area_province = { add_claim = root } }\n" +
				"			area:tigris_area = { every_area_province = { add_claim = root } }\n" +
				"		}\n" +
				"		add_aggressive_expansion = 15\n" +
				"		add_tyranny = 10\n" +
				"		add_political_influence = -30\n" +
				"	}",
			Option2: "option = {\n" +
				"		name = me_tylos_3.b\n" +
				"		add_political_influence = 30\n" +
				"	}",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
	}
	eventFile := EventFile{"", "me_tylos_test"}
	eventFile.makeFile(events)
}
