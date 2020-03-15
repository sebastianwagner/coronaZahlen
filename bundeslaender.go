package main

func regions() caseRegions {
	return caseRegions{
		"Baden-Württemberg": {
			URL:      "https://sozialministerium.baden-wuerttemberg.de/de/gesundheit-pflege/gesundheitsschutz/infektionsschutz-hygiene/informationen-zu-coronavirus/",
			Selector: "figcaption",
			Match:    `([.\d]+) bestätigte Corona-Fälle`,
		},
		"Bayern": {
			URL:      "https://www.lgl.bayern.de/gesundheit/infektionsschutz/infektionskrankheiten_a_z/coronavirus/karte_coronavirus/index.htm",
			Selector: "div.row > div:nth-child(2) > table > tbody > tr:nth-child(9) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Berlin": {
			URL:      "https://www.berlin.de/sen/gpg/service/presse/",
			Selector: ".rss > .html5-section > ul > li > a",
			Match:    `Coronavirus: Derzeit ([.\d]+) bestätigte Fälle`,
		},
		"Brandenburg": {
			URL:      "https://msgiv.brandenburg.de/msgiv/de/start/",
			Selector: "#form_id32516 > div > div > h2",
			Match:    `Insgesamt ([.\d]+) Erkrankungen`,
		},
		"Bremen": {
			URL:      "https://www.gesundheit.bremen.de/sixcms/detail.php?gsid=bremen229.c.32657.de",
			Selector: "#main > div:nth-child(3) > div",
			Match:    ` ([.\d]+) als positiv gemeldete`,
		},
		"Hamburg": {
			URL:      "https://www.hamburg.de/bgv/pressemeldungen/13719254/2020-03-13-bgv-coronavirus-aktuell/",
			Selector: "#pushedContainer > main > div.article > div:nth-child(7) > div > div > div",
			Match:    `insgesamt ([.\d]+) positiv Getestete`,
		},
		"Hessen": {
			URL:      "https://soziales.hessen.de/gesundheit/infektionsschutz/coronavirus-sars-cov-2",
			Selector: "section:nth-child(2) > .block-inner > .blockContent",
			Match:    `insgesamt ([.\d]+) SARS-CoV-2-Fälle`,
		},
		"Mecklenburg-Vorpommern": {
			URL:      "https://www.regierung-mv.de/Landesregierung/wm/Aktuell/?sa.query=neue+Corona-Infektionen&sa.pressemitteilungen.area=11",
			Selector: ".dvz-contenttype-presseserviceassistent",
			Match:    `Insgesamt wurden bislang ([.\d]+) Menschen`,
		},
		"Niedersachsen": {
			URL:      "https://www.niedersachsen.de/Coronavirus",
			Selector: ".maincontent > .content > .complementary > div",
			Match:    ` ([.\d]+)\s+laborbestätigte Covid-19-Fälle`,
		},
		"Nordrhein-Westfalen": {
			URL:      "https://www.mags.nrw/coronavirus-fallzahlen-nrw",
			Selector: "table > tbody > tr:nth-child(54) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Rheinland-Pfalz": {
			URL:      "https://msagd.rlp.de/de/service/presse/detail/news/News/detail/information-der-landesregierung-zum-aktuellen-stand-hinsichtlich-des-coronavirus-bundesratsinitiati/",
			Selector: "h6",
			Match:    `insgesamt ([.\d]+) bestätigte SARS-CoV-2 Fälle`,
		},
		"Saarland": {
			URL:      "https://www.saarland.de/SID-C29AF463-5CFEAE1B/253741.htm",
			Selector: ".textchapter_frame",
			Match:    `landesweit auf ([.\d]+) – `,
		},
		"Sachsen": {
			URL:      "https://www.sms.sachsen.de/coronavirus.html",
			Selector: "table > tbody > tr:nth-child(14) > td:nth-child(2)",
			Match:    `([.\d]+)`,
		},
		"Sachsen-Anhalt": {
			URL:      "https://ms.sachsen-anhalt.de/presse/pressemitteilungen/?no_cache=1",
			Selector: ".tx-rssdisplay > div > div > h2",
			Match:    `([.\d]+)`,
		},
		"Schleswig-Holstein": {
			URL:      "https://schleswig-holstein.de/SiteGlobals/Forms/Suche/DE/Expertensuche_Formular.html?gts=%2526aa42becc-5285-4475-af4d-0413dde5e634_list%253DdateOfIssue_dt%252Bdesc&documentType_=pressrelease&templateQueryString=Gesundheitsministerium+informiert+zum+Corona-Virus",
			Selector: "#content > .singleview > .bodyText",
			Match:    `positiv getesteten Covid19-Fälle in Schleswig-Holstein: Summe ([.\d]+)`,
		},
		"Thüringen": {
			URL:      "https://www.tmasgff.de/covid-19",
			Selector: "h3",
			Match:    `([.\d]+) bestätigte Infektionen`,
		},
	}
}