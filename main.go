package main

func main() {
	events := []Event{
		Event{
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
				"\t\tname = me_tylos_1.a\n" +
				"\t\tcurrent_ruler = {\n" +
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
				"\t\t}\n" +
				"\t}",
			Option2: "",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
		Event{
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
				"\t\tname = me_tylos_2.a\n" +
				"\t\tadd_political_influence = -65\n" +
				"\t\tchange_government = babylonian_bureaucracy\n" +
				"\t}",
			Option2: "option = {\n" +
				"\t\tname = me_tylos_2.b\n" +
				"\t\tadd_political_influence = 35\n" +
				"\t}",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
		Event{
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
				"\t\tcustom_tooltip = me_tylos_3_i_tt\n" +
				"\t\thidden_effect = {\n" +
				"\t\t\tarea:babylonia_area = { every_area_province = { add_claim = root } }\n" +
				"\t\t}\n" +
				"\t\tp:918.owner = {\n" +
				"\t\t\tadd_opinion = { modifier = declare_babylon_intent_realtion_tylos target = root }\n" +
				"\t\t\tif = {\n" +
				"\t\t\t\tlimit = {\n" +
				"\t\t\t\t\troot = {\n" +
				"\t\t\t\t\t\tNOT = { is_subject_of = prev }\n" +
				"\t\t\t\t\t}\n" +
				"\t\t\t\t}\n" +
				"\t\t\t\trelease_subject = root\n" +
				"\t\t\t}\n" +
				"\t\t}\n" +
				"\t\trandom_country_culture = {\n" +
				"\t\t\tlimit = {\n" +
				"\t\t\t\tis_culture = babylonian\n" +
				"\t\t\t}\n" +
				"\t\t\tadd_country_culture_modifier = {\n" +
				"\t\t\t\tname = memory_of_babylon\n" +
				"\t\t\t\tduration = 9125\n" +
				"\t\t\t}\n" +
				"\t\t}\n" +
				"\t}",
			Option: "option = {\n" +
				"\t\tname = me_tylos_3.a\n" +
				"\t\tcustom_tooltip = me_tylos_3_a_tt\n" +
				"\t\thidden_effect = {\n" +
				"\t\t\tarea:beth_daraye_area = { every_area_province = { add_claim = root } }\n" +
				"\t\t\tarea:asoristan_area = { every_area_province = { add_claim = root } }\n" +
				"\t\t\tarea:tigris_area = { every_area_province = { add_claim = root } }\n" +
				"\t\t}\n" +
				"\t\tadd_aggressive_expansion = 15\n" +
				"\t\tadd_tyranny = 10\n" +
				"\t\tadd_political_influence = -30\n" +
				"\t}",
			Option2: "option = {\n" +
				"\t\tname = me_tylos_3.b\n" +
				"\t\tadd_political_influence = 30\n" +
				"\t}",
			Option3: "",
			Option4: "",
			Option5: "",
			After:   "",
		},
	}
	eventFile := EventFile{"", "me_tylos_test"}
	eventFile.makeFile(events)
}
