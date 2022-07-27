package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	r := strings.NewReader(getString())
	doc, err := html.Parse(r)
	if err != nil {
		panic(err.Error())
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			if n.FirstChild != nil {
				fmt.Printf("{name: \"%s\", id:\"%s\"},\n", n.FirstChild.Data, n.FirstChild.Data)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func getString() string {
	return `<div data-test="all-carriers">
    <div class="patient-web-app__sc-53koon-2 lmCwqN"></div>
    <div data-test="all-header" class="patient-web-app__sc-53koon-4 coAGHm">All carriers</div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">#</div>
    <div data-row-key="440-1199S" data-test="insurance-picker-row" data-uem-id="440-1199S">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>1199SEIU</span>
        </div>
    </div>
    <div data-row-key="1520-1st A" data-test="insurance-picker-row" data-uem-id="1520-1st A">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>1st Agency</span>
        </div>
    </div>
    <div data-row-key="876-20/20" data-test="insurance-picker-row" data-uem-id="876-20/20">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>20/20 Eyecare Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">A</div>
    <div data-row-key="290-AARP" data-test="insurance-picker-row" data-uem-id="290-AARP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AARP</span>
        </div>
    </div>
    <div data-row-key="480-ACE" data-test="insurance-picker-row" data-uem-id="480-ACE">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ACE</span>
        </div>
    </div>
    <div data-row-key="291-AIG" data-test="insurance-picker-row" data-uem-id="291-AIG">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AIG</span>
        </div>
    </div>
    <div data-row-key="391-APWU" data-test="insurance-picker-row" data-uem-id="391-APWU">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>APWU</span>
        </div>
    </div>
    <div data-row-key="926-ATRIO" data-test="insurance-picker-row" data-uem-id="926-ATRIO">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ATRIO Health Plans</span>
        </div>
    </div>
    <div data-row-key="1522-AVMA " data-test="insurance-picker-row" data-uem-id="1522-AVMA ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AVMA Life</span>
        </div>
    </div>
    <div data-row-key="1239-Absol" data-test="insurance-picker-row" data-uem-id="1239-Absol">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Absolute Total Care</span>
        </div>
    </div>
    <div data-row-key="1215-Acces" data-test="insurance-picker-row" data-uem-id="1215-Acces">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Access Medicare (NY)</span>
        </div>
    </div>
    <div data-row-key="526-Accou" data-test="insurance-picker-row" data-uem-id="526-Accou">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Accountable Health Plan of Ohio</span>
        </div>
    </div>
    <div data-row-key="1592-Advan" data-test="insurance-picker-row" data-uem-id="1592-Advan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Advanced Health</span>
        </div>
    </div>
    <div data-row-key="1354-Advan" data-test="insurance-picker-row" data-uem-id="1354-Advan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AdvantageMD</span>
        </div>
    </div>
    <div data-row-key="771-Advan" data-test="insurance-picker-row" data-uem-id="771-Advan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Advantica</span>
        </div>
    </div>
    <div data-row-key="882-Adven" data-test="insurance-picker-row" data-uem-id="882-Adven">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Advent Health</span>
        </div>
    </div>
    <div data-row-key="533-Adven" data-test="insurance-picker-row" data-uem-id="533-Adven">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Adventist Health</span>
        </div>
    </div>
    <div data-row-key="1060-Advoc" data-test="insurance-picker-row" data-uem-id="1060-Advoc">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Advocate Health Care</span>
        </div>
    </div>
    <div data-row-key="300-Aetna" data-test="insurance-picker-row" data-uem-id="300-Aetna">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Aetna</span>
        </div>
    </div>
    <div data-row-key="1298-Aetna" data-test="insurance-picker-row" data-uem-id="1298-Aetna">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Aetna Better Health</span>
        </div>
    </div>
    <div data-row-key="340-Affin" data-test="insurance-picker-row" data-uem-id="340-Affin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Affinity Health Plan</span>
        </div>
    </div>
    <div data-row-key="1654-AgeRi" data-test="insurance-picker-row" data-uem-id="1654-AgeRi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AgeRight Advantage</span>
        </div>
    </div>
    <div data-row-key="1291-AgeWe" data-test="insurance-picker-row" data-uem-id="1291-AgeWe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AgeWell New York</span>
        </div>
    </div>
    <div data-row-key="1324-Agile" data-test="insurance-picker-row" data-uem-id="1324-Agile">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Agile Health Insurance</span>
        </div>
    </div>
    <div data-row-key="877-Alame" data-test="insurance-picker-row" data-uem-id="877-Alame">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Alameda Alliance for Health</span>
        </div>
    </div>
    <div data-row-key="1597-Alier" data-test="insurance-picker-row" data-uem-id="1597-Alier">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Aliera Health Care</span>
        </div>
    </div>
    <div data-row-key="1433-Align" data-test="insurance-picker-row" data-uem-id="1433-Align">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Alignment Health Plan</span>
        </div>
    </div>
    <div data-row-key="1314-All S" data-test="insurance-picker-row" data-uem-id="1314-All S">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>All Savers Insurance</span>
        </div>
    </div>
    <div data-row-key="1474-AllCa" data-test="insurance-picker-row" data-uem-id="1474-AllCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AllCare Health</span>
        </div>
    </div>
    <div data-row-key="865-AllSt" data-test="insurance-picker-row" data-uem-id="865-AllSt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AllState</span>
        </div>
    </div>
    <div data-row-key="1544-AllWa" data-test="insurance-picker-row" data-uem-id="1544-AllWa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AllWays Health Partners</span>
        </div>
    </div>
    <div data-row-key="1085-Alleg" data-test="insurance-picker-row" data-uem-id="1085-Alleg">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Allegiance Life and Health</span>
        </div>
    </div>
    <div data-row-key="1084-Allia" data-test="insurance-picker-row" data-uem-id="1084-Allia">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Alliant Health Plans</span>
        </div>
    </div>
    <div data-row-key="1311-Allia" data-test="insurance-picker-row" data-uem-id="1311-Allia">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Allianz Worldwide Care</span>
        </div>
    </div>
    <div data-row-key="1471-Allwe" data-test="insurance-picker-row" data-uem-id="1471-Allwe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Allwell</span>
        </div>
    </div>
    <div data-row-key="1173-Aloha" data-test="insurance-picker-row" data-uem-id="1173-Aloha">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AlohaCare</span>
        </div>
    </div>
    <div data-row-key="1290-Alpha" data-test="insurance-picker-row" data-uem-id="1290-Alpha">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AlphaCare</span>
        </div>
    </div>
    <div data-row-key="957-AltaM" data-test="insurance-picker-row" data-uem-id="957-AltaM">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AltaMed Senior BuenaCare (PACE)</span>
        </div>
    </div>
    <div data-row-key="325-Altiu" data-test="insurance-picker-row" data-uem-id="325-Altiu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Altius (Coventry Health Care)</span>
        </div>
    </div>
    <div data-row-key="708-Alway" data-test="insurance-picker-row" data-uem-id="708-Alway">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AlwaysCare</span>
        </div>
    </div>
    <div data-row-key="1323-Ambet" data-test="insurance-picker-row" data-uem-id="1323-Ambet">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ambetter</span>
        </div>
    </div>
    <div data-row-key="388-Ameri" data-test="insurance-picker-row" data-uem-id="388-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AmeriGroup</span>
        </div>
    </div>
    <div data-row-key="326-Ameri" data-test="insurance-picker-row" data-uem-id="326-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AmeriHealth</span>
        </div>
    </div>
    <div data-row-key="622-Ameri" data-test="insurance-picker-row" data-uem-id="622-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AmeriHealth Caritas</span>
        </div>
    </div>
    <div data-row-key="909-Ameri" data-test="insurance-picker-row" data-uem-id="909-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>America's 1st Choice</span>
        </div>
    </div>
    <div data-row-key="1228-Ameri" data-test="insurance-picker-row" data-uem-id="1228-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>American Behavioral</span>
        </div>
    </div>
    <div data-row-key="1054-Ameri" data-test="insurance-picker-row" data-uem-id="1054-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>American Eldercare</span>
        </div>
    </div>
    <div data-row-key="1431-Ameri" data-test="insurance-picker-row" data-uem-id="1431-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>American Healthcare Alliance</span>
        </div>
    </div>
    <div data-row-key="1521-Ameri" data-test="insurance-picker-row" data-uem-id="1521-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>American Maritime Officers Plans</span>
        </div>
    </div>
    <div data-row-key="711-Ameri" data-test="insurance-picker-row" data-uem-id="711-Ameri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>American Republic Insurance Company</span>
        </div>
    </div>
    <div data-row-key="976-Amida" data-test="insurance-picker-row" data-uem-id="976-Amida">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Amida Care</span>
        </div>
    </div>
    <div data-row-key="991-Ampli" data-test="insurance-picker-row" data-uem-id="991-Ampli">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Amplifon Hearing Health Care</span>
        </div>
    </div>
    <div data-row-key="550-Anthe" data-test="insurance-picker-row" data-uem-id="550-Anthe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Anthem Blue Cross</span>
        </div>
    </div>
    <div data-row-key="324-Anthe" data-test="insurance-picker-row" data-uem-id="324-Anthe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Anthem Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="1618-Apost" data-test="insurance-picker-row" data-uem-id="1618-Apost">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Apostrophe</span>
        </div>
    </div>
    <div data-row-key="1426-ArchC" data-test="insurance-picker-row" data-uem-id="1426-ArchC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ArchCare</span>
        </div>
    </div>
    <div data-row-key="1160-Arise" data-test="insurance-picker-row" data-uem-id="1160-Arise">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Arise Health Plan</span>
        </div>
    </div>
    <div data-row-key="1545-Arizo" data-test="insurance-picker-row" data-uem-id="1545-Arizo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Arizona Complete Health</span>
        </div>
    </div>
    <div data-row-key="678-Arizo" data-test="insurance-picker-row" data-uem-id="678-Arizo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Arizona Foundation for Medical Care</span>
        </div>
    </div>
    <div data-row-key="1010-Arkan" data-test="insurance-picker-row" data-uem-id="1010-Arkan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Arkansas Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="1611-Arkan" data-test="insurance-picker-row" data-uem-id="1611-Arkan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Arkansas Total Care</span>
        </div>
    </div>
    <div data-row-key="1634-Ascen" data-test="insurance-picker-row" data-uem-id="1634-Ascen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ascension Complete</span>
        </div>
    </div>
    <div data-row-key="816-Ascen" data-test="insurance-picker-row" data-uem-id="816-Ascen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ascension Health</span>
        </div>
    </div>
    <div data-row-key="1652-Aspir" data-test="insurance-picker-row" data-uem-id="1652-Aspir">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Aspire Health Plan</span>
        </div>
    </div>
    <div data-row-key="1053-Assur" data-test="insurance-picker-row" data-uem-id="1053-Assur">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Assurant Employee Benefits</span>
        </div>
    </div>
    <div data-row-key="335-Assur" data-test="insurance-picker-row" data-uem-id="335-Assur">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Assurant Health</span>
        </div>
    </div>
    <div data-row-key="716-Asuri" data-test="insurance-picker-row" data-uem-id="716-Asuri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Asuris Northwest Health</span>
        </div>
    </div>
    <div data-row-key="1130-Aultc" data-test="insurance-picker-row" data-uem-id="1130-Aultc">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Aultcare</span>
        </div>
    </div>
    <div data-row-key="560-AvMed" data-test="insurance-picker-row" data-uem-id="560-AvMed">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>AvMed</span>
        </div>
    </div>
    <div data-row-key="1195-Avera" data-test="insurance-picker-row" data-uem-id="1195-Avera">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Avera Health Plans</span>
        </div>
    </div>
    <div data-row-key="549-Avesi" data-test="insurance-picker-row" data-uem-id="549-Avesi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Avesis</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">B</div>
    <div data-row-key="680-BMC H" data-test="insurance-picker-row" data-uem-id="680-BMC H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>BMC HealthNet Plan</span>
        </div>
    </div>
    <div data-row-key="1352-Banke" data-test="insurance-picker-row" data-uem-id="1352-Banke">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Banker's Life</span>
        </div>
    </div>
    <div data-row-key="717-Banne" data-test="insurance-picker-row" data-uem-id="717-Banne">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Banner Health</span>
        </div>
    </div>
    <div data-row-key="1318-Bapti" data-test="insurance-picker-row" data-uem-id="1318-Bapti">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Baptist Health Plan</span>
        </div>
    </div>
    <div data-row-key="1546-BayCa" data-test="insurance-picker-row" data-uem-id="1546-BayCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>BayCarePlus Medicare Advantage</span>
        </div>
    </div>
    <div data-row-key="519-Beaco" data-test="insurance-picker-row" data-uem-id="519-Beaco">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Beacon Health Options</span>
        </div>
    </div>
    <div data-row-key="720-Beaum" data-test="insurance-picker-row" data-uem-id="720-Beaum">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Beaumont Employee Health Plan</span>
        </div>
    </div>
    <div data-row-key="303-Beech" data-test="insurance-picker-row" data-uem-id="303-Beech">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Beech Street</span>
        </div>
    </div>
    <div data-row-key="721-Best " data-test="insurance-picker-row" data-uem-id="721-Best ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Best Choice Plus</span>
        </div>
    </div>
    <div data-row-key="1483-Best " data-test="insurance-picker-row" data-uem-id="1483-Best ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Best Doctors Insurance</span>
        </div>
    </div>
    <div data-row-key="993-Best " data-test="insurance-picker-row" data-uem-id="993-Best ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Best Life And Health</span>
        </div>
    </div>
    <div data-row-key="879-Bette" data-test="insurance-picker-row" data-uem-id="879-Bette">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Better Health (Florida Medicaid)</span>
        </div>
    </div>
    <div data-row-key="1632-Bind " data-test="insurance-picker-row" data-uem-id="1632-Bind ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Bind HealthCare</span>
        </div>
    </div>
    <div data-row-key="1109-Blue " data-test="insurance-picker-row" data-uem-id="1109-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Choice Health Plan</span>
        </div>
    </div>
    <div data-row-key="304-Blue " data-test="insurance-picker-row" data-uem-id="304-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield (BCBS)</span>
        </div>
    </div>
    <div data-row-key="545-Blue " data-test="insurance-picker-row" data-uem-id="545-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield Federal Employee Program</span>
        </div>
    </div>
    <div data-row-key="793-Blue " data-test="insurance-picker-row" data-uem-id="793-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Alabama</span>
        </div>
    </div>
    <div data-row-key="572-Blue " data-test="insurance-picker-row" data-uem-id="572-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Arizona</span>
        </div>
    </div>
    <div data-row-key="568-Blue " data-test="insurance-picker-row" data-uem-id="568-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Georgia</span>
        </div>
    </div>
    <div data-row-key="451-Blue " data-test="insurance-picker-row" data-uem-id="451-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Illinois</span>
        </div>
    </div>
    <div data-row-key="810-Blue " data-test="insurance-picker-row" data-uem-id="810-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Kansas</span>
        </div>
    </div>
    <div data-row-key="726-Blue " data-test="insurance-picker-row" data-uem-id="726-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Kansas City</span>
        </div>
    </div>
    <div data-row-key="853-Blue " data-test="insurance-picker-row" data-uem-id="853-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Louisiana</span>
        </div>
    </div>
    <div data-row-key="573-Blue " data-test="insurance-picker-row" data-uem-id="573-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Massachusetts</span>
        </div>
    </div>
    <div data-row-key="676-Blue " data-test="insurance-picker-row" data-uem-id="676-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Michigan</span>
        </div>
    </div>
    <div data-row-key="698-Blue " data-test="insurance-picker-row" data-uem-id="698-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Minnesota</span>
        </div>
    </div>
    <div data-row-key="758-Blue " data-test="insurance-picker-row" data-uem-id="758-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Mississippi</span>
        </div>
    </div>
    <div data-row-key="1207-Blue " data-test="insurance-picker-row" data-uem-id="1207-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Montana</span>
        </div>
    </div>
    <div data-row-key="756-Blue " data-test="insurance-picker-row" data-uem-id="756-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Nebraska</span>
        </div>
    </div>
    <div data-row-key="1096-Blue " data-test="insurance-picker-row" data-uem-id="1096-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of New Mexico</span>
        </div>
    </div>
    <div data-row-key="791-Blue " data-test="insurance-picker-row" data-uem-id="791-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of North Carolina</span>
        </div>
    </div>
    <div data-row-key="1201-Blue " data-test="insurance-picker-row" data-uem-id="1201-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of North Dakota</span>
        </div>
    </div>
    <div data-row-key="854-Blue " data-test="insurance-picker-row" data-uem-id="854-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Oklahoma</span>
        </div>
    </div>
    <div data-row-key="724-Blue " data-test="insurance-picker-row" data-uem-id="724-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Rhode Island</span>
        </div>
    </div>
    <div data-row-key="797-Blue " data-test="insurance-picker-row" data-uem-id="797-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of South Carolina</span>
        </div>
    </div>
    <div data-row-key="828-Blue " data-test="insurance-picker-row" data-uem-id="828-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Tennessee</span>
        </div>
    </div>
    <div data-row-key="509-Blue " data-test="insurance-picker-row" data-uem-id="509-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Texas</span>
        </div>
    </div>
    <div data-row-key="996-Blue " data-test="insurance-picker-row" data-uem-id="996-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Vermont</span>
        </div>
    </div>
    <div data-row-key="997-Blue " data-test="insurance-picker-row" data-uem-id="997-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Western New York</span>
        </div>
    </div>
    <div data-row-key="1073-Blue " data-test="insurance-picker-row" data-uem-id="1073-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross Blue Shield of Wyoming</span>
        </div>
    </div>
    <div data-row-key="1094-Blue " data-test="insurance-picker-row" data-uem-id="1094-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross of Idaho</span>
        </div>
    </div>
    <div data-row-key="808-Blue " data-test="insurance-picker-row" data-uem-id="808-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Cross of Northeastern Pennsylvania</span>
        </div>
    </div>
    <div data-row-key="438-Blue " data-test="insurance-picker-row" data-uem-id="438-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Shield of California</span>
        </div>
    </div>
    <div data-row-key="999-Blue " data-test="insurance-picker-row" data-uem-id="999-Blue ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Blue Shield of Northeastern New York</span>
        </div>
    </div>
    <div data-row-key="944-Brand" data-test="insurance-picker-row" data-uem-id="944-Brand">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Brand New Day</span>
        </div>
    </div>
    <div data-row-key="1637-Brave" data-test="insurance-picker-row" data-uem-id="1637-Brave">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Braven Health (Horizon Blue Cross Blue Shield of New Jersey)</span>
        </div>
    </div>
    <div data-row-key="1115-Bridg" data-test="insurance-picker-row" data-uem-id="1115-Bridg">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>BridgeSpan</span>
        </div>
    </div>
    <div data-row-key="579-Bridg" data-test="insurance-picker-row" data-uem-id="579-Bridg">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Bridgeway Health Solutions</span>
        </div>
    </div>
    <div data-row-key="1423-Brigh" data-test="insurance-picker-row" data-uem-id="1423-Brigh">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Bright Health</span>
        </div>
    </div>
    <div data-row-key="1131-Bucke" data-test="insurance-picker-row" data-uem-id="1131-Bucke">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Buckeye Health Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">C</div>
    <div data-row-key="1017-CBA B" data-test="insurance-picker-row" data-uem-id="1017-CBA B">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CBA Blue</span>
        </div>
    </div>
    <div data-row-key="352-CDPHP" data-test="insurance-picker-row" data-uem-id="352-CDPHP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CDPHP</span>
        </div>
    </div>
    <div data-row-key="730-CHAMP" data-test="insurance-picker-row" data-uem-id="730-CHAMP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CHAMPVA</span>
        </div>
    </div>
    <div data-row-key="1087-CHP G" data-test="insurance-picker-row" data-uem-id="1087-CHP G">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CHP Group</span>
        </div>
    </div>
    <div data-row-key="1303-CHRIS" data-test="insurance-picker-row" data-uem-id="1303-CHRIS">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CHRISTUS Health Plan</span>
        </div>
    </div>
    <div data-row-key="929-CalOp" data-test="insurance-picker-row" data-uem-id="929-CalOp">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CalOptima</span>
        </div>
    </div>
    <div data-row-key="989-CalPE" data-test="insurance-picker-row" data-uem-id="989-CalPE">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CalPERS</span>
        </div>
    </div>
    <div data-row-key="1289-CalVi" data-test="insurance-picker-row" data-uem-id="1289-CalVi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CalViva Health</span>
        </div>
    </div>
    <div data-row-key="1513-Calif" data-test="insurance-picker-row" data-uem-id="1513-Calif">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>California Foundation for Medical Care</span>
        </div>
    </div>
    <div data-row-key="1241-Calif" data-test="insurance-picker-row" data-uem-id="1241-Calif">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>California Health and Wellness</span>
        </div>
    </div>
    <div data-row-key="1535-Calvo" data-test="insurance-picker-row" data-uem-id="1535-Calvo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Calvos</span>
        </div>
    </div>
    <div data-row-key="1540-Cambr" data-test="insurance-picker-row" data-uem-id="1540-Cambr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cambridge Health Alliance (CHA)</span>
        </div>
    </div>
    <div data-row-key="566-Capit" data-test="insurance-picker-row" data-uem-id="566-Capit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Capital Blue Cross</span>
        </div>
    </div>
    <div data-row-key="908-Capit" data-test="insurance-picker-row" data-uem-id="908-Capit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Capital Health Plan</span>
        </div>
    </div>
    <div data-row-key="728-Care " data-test="insurance-picker-row" data-uem-id="728-Care ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Care Access Health Plan</span>
        </div>
    </div>
    <div data-row-key="527-Care " data-test="insurance-picker-row" data-uem-id="527-Care ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Care Improvement Plus</span>
        </div>
    </div>
    <div data-row-key="1030-Care " data-test="insurance-picker-row" data-uem-id="1030-Care ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Care N' Care</span>
        </div>
    </div>
    <div data-row-key="581-Care1" data-test="insurance-picker-row" data-uem-id="581-Care1">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Care1st</span>
        </div>
    </div>
    <div data-row-key="1358-CareC" data-test="insurance-picker-row" data-uem-id="1358-CareC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CareConnect</span>
        </div>
    </div>
    <div data-row-key="305-CareF" data-test="insurance-picker-row" data-uem-id="305-CareF">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CareFirst Blue Cross Blue Shield  (Health)</span>
        </div>
    </div>
    <div data-row-key="605-CareM" data-test="insurance-picker-row" data-uem-id="605-CareM">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CareMore</span>
        </div>
    </div>
    <div data-row-key="807-CareO" data-test="insurance-picker-row" data-uem-id="807-CareO">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CareOregon</span>
        </div>
    </div>
    <div data-row-key="1554-CareP" data-test="insurance-picker-row" data-uem-id="1554-CareP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CarePartners of Connecticut</span>
        </div>
    </div>
    <div data-row-key="341-CareP" data-test="insurance-picker-row" data-uem-id="341-CareP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CarePlus Health Plans (Florida Medicare)</span>
        </div>
    </div>
    <div data-row-key="1078-CareS" data-test="insurance-picker-row" data-uem-id="1078-CareS">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CareSource</span>
        </div>
    </div>
    <div data-row-key="925-Casca" data-test="insurance-picker-row" data-uem-id="925-Casca">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cascade Health Alliance</span>
        </div>
    </div>
    <div data-row-key="457-Cater" data-test="insurance-picker-row" data-uem-id="457-Cater">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Caterpillar</span>
        </div>
    </div>
    <div data-row-key="633-Celti" data-test="insurance-picker-row" data-uem-id="633-Celti">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CeltiCare Health Plan</span>
        </div>
    </div>
    <div data-row-key="930-CenCa" data-test="insurance-picker-row" data-uem-id="930-CenCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CenCal Health</span>
        </div>
    </div>
    <div data-row-key="1364-Cenpa" data-test="insurance-picker-row" data-uem-id="1364-Cenpa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cenpatico</span>
        </div>
    </div>
    <div data-row-key="1454-Cente" data-test="insurance-picker-row" data-uem-id="1454-Cente">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Centennial Care</span>
        </div>
    </div>
    <div data-row-key="931-Cente" data-test="insurance-picker-row" data-uem-id="931-Cente">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Center for Elders' Independence (PACE)</span>
        </div>
    </div>
    <div data-row-key="729-Cente" data-test="insurance-picker-row" data-uem-id="729-Cente">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CenterLight Healthcare</span>
        </div>
    </div>
    <div data-row-key="1480-Cente" data-test="insurance-picker-row" data-uem-id="1480-Cente">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Centers Plan for Healthy Living</span>
        </div>
    </div>
    <div data-row-key="1449-Cente" data-test="insurance-picker-row" data-uem-id="1449-Cente">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Centers for Medicare &amp; Medicaid Services</span>
        </div>
    </div>
    <div data-row-key="1625-Centi" data-test="insurance-picker-row" data-uem-id="1625-Centi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Centivo</span>
        </div>
    </div>
    <div data-row-key="1573-Centr" data-test="insurance-picker-row" data-uem-id="1573-Centr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Centra Health</span>
        </div>
    </div>
    <div data-row-key="932-Centr" data-test="insurance-picker-row" data-uem-id="932-Centr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Central California Alliance for Health</span>
        </div>
    </div>
    <div data-row-key="933-Centr" data-test="insurance-picker-row" data-uem-id="933-Centr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Central Health Plan of California</span>
        </div>
    </div>
    <div data-row-key="1356-Centu" data-test="insurance-picker-row" data-uem-id="1356-Centu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Century Healthcare - CHC</span>
        </div>
    </div>
    <div data-row-key="1163-Child" data-test="insurance-picker-row" data-uem-id="1163-Child">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Children's Community Health Plan</span>
        </div>
    </div>
    <div data-row-key="1580-Child" data-test="insurance-picker-row" data-uem-id="1580-Child">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Children's Medical Center Health Plan</span>
        </div>
    </div>
    <div data-row-key="880-Child" data-test="insurance-picker-row" data-uem-id="880-Child">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Children's Medical Services (CMS)</span>
        </div>
    </div>
    <div data-row-key="420-Chine" data-test="insurance-picker-row" data-uem-id="420-Chine">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Chinese Community Health Plan</span>
        </div>
    </div>
    <div data-row-key="394-Choic" data-test="insurance-picker-row" data-uem-id="394-Choic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Choice Care Network</span>
        </div>
    </div>
    <div data-row-key="1366-Chris" data-test="insurance-picker-row" data-uem-id="1366-Chris">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Christian Healthcare Ministries</span>
        </div>
    </div>
    <div data-row-key="307-Cigna" data-test="insurance-picker-row" data-uem-id="307-Cigna">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cigna</span>
        </div>
    </div>
    <div data-row-key="510-Cigna" data-test="insurance-picker-row" data-uem-id="510-Cigna">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cigna-HealthSpring</span>
        </div>
    </div>
    <div data-row-key="1621-Clari" data-test="insurance-picker-row" data-uem-id="1621-Clari">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Clarion Health</span>
        </div>
    </div>
    <div data-row-key="1304-Clark" data-test="insurance-picker-row" data-uem-id="1304-Clark">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Clark County Self-Funded Health</span>
        </div>
    </div>
    <div data-row-key="1572-Clear" data-test="insurance-picker-row" data-uem-id="1572-Clear">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Clear Spring Health</span>
        </div>
    </div>
    <div data-row-key="546-Cleme" data-test="insurance-picker-row" data-uem-id="546-Cleme">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Clements Worldwide</span>
        </div>
    </div>
    <div data-row-key="1300-Cleve" data-test="insurance-picker-row" data-uem-id="1300-Cleve">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cleveland Clinic Employee Health Plan</span>
        </div>
    </div>
    <div data-row-key="1306-Clove" data-test="insurance-picker-row" data-uem-id="1306-Clove">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Clover Health</span>
        </div>
    </div>
    <div data-row-key="679-Cofin" data-test="insurance-picker-row" data-uem-id="679-Cofin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cofinity</span>
        </div>
    </div>
    <div data-row-key="1083-Color" data-test="insurance-picker-row" data-uem-id="1083-Color">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Colorado Access</span>
        </div>
    </div>
    <div data-row-key="1588-Colum" data-test="insurance-picker-row" data-uem-id="1588-Colum">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Columbia Pacific CCO</span>
        </div>
    </div>
    <div data-row-key="1052-Colum" data-test="insurance-picker-row" data-uem-id="1052-Colum">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Columbia United Providers</span>
        </div>
    </div>
    <div data-row-key="532-ComPs" data-test="insurance-picker-row" data-uem-id="532-ComPs">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ComPsych</span>
        </div>
    </div>
    <div data-row-key="1292-Commo" data-test="insurance-picker-row" data-uem-id="1292-Commo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Common Ground Healthcare Cooperative</span>
        </div>
    </div>
    <div data-row-key="1297-Commo" data-test="insurance-picker-row" data-uem-id="1297-Commo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Commonwealth Care Alliance</span>
        </div>
    </div>
    <div data-row-key="1346-Commu" data-test="insurance-picker-row" data-uem-id="1346-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Behavioral Health</span>
        </div>
    </div>
    <div data-row-key="1424-Commu" data-test="insurance-picker-row" data-uem-id="1424-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Care Alliance of Illinois</span>
        </div>
    </div>
    <div data-row-key="1537-Commu" data-test="insurance-picker-row" data-uem-id="1537-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Care Associates</span>
        </div>
    </div>
    <div data-row-key="1223-Commu" data-test="insurance-picker-row" data-uem-id="1223-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Care Behavioral Health Organization</span>
        </div>
    </div>
    <div data-row-key="1566-Commu" data-test="insurance-picker-row" data-uem-id="1566-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Care Plan</span>
        </div>
    </div>
    <div data-row-key="831-Commu" data-test="insurance-picker-row" data-uem-id="831-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Care of North Carolina</span>
        </div>
    </div>
    <div data-row-key="1058-Commu" data-test="insurance-picker-row" data-uem-id="1058-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Eye Care</span>
        </div>
    </div>
    <div data-row-key="512-Commu" data-test="insurance-picker-row" data-uem-id="512-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community First Health Plans</span>
        </div>
    </div>
    <div data-row-key="513-Commu" data-test="insurance-picker-row" data-uem-id="513-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Health Choice</span>
        </div>
    </div>
    <div data-row-key="935-Commu" data-test="insurance-picker-row" data-uem-id="935-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Health Group</span>
        </div>
    </div>
    <div data-row-key="1266-Commu" data-test="insurance-picker-row" data-uem-id="1266-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Health Options</span>
        </div>
    </div>
    <div data-row-key="647-Commu" data-test="insurance-picker-row" data-uem-id="647-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Health Partners</span>
        </div>
    </div>
    <div data-row-key="677-Commu" data-test="insurance-picker-row" data-uem-id="677-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Community Health Plan of Washington</span>
        </div>
    </div>
    <div data-row-key="1105-Commu" data-test="insurance-picker-row" data-uem-id="1105-Commu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CommunityCare of Oklahoma</span>
        </div>
    </div>
    <div data-row-key="661-CompB" data-test="insurance-picker-row" data-uem-id="661-CompB">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CompBenefits</span>
        </div>
    </div>
    <div data-row-key="1444-Compa" data-test="insurance-picker-row" data-uem-id="1444-Compa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Companion Life</span>
        </div>
    </div>
    <div data-row-key="466-Compr" data-test="insurance-picker-row" data-uem-id="466-Compr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Comprehensive Health Insurance Plan (CHIP) of Illinois</span>
        </div>
    </div>
    <div data-row-key="556-Compr" data-test="insurance-picker-row" data-uem-id="556-Compr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Comprehensive Medical and Dental Program (CMDP)</span>
        </div>
    </div>
    <div data-row-key="1180-Conne" data-test="insurance-picker-row" data-uem-id="1180-Conne">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Connect Care</span>
        </div>
    </div>
    <div data-row-key="329-Conne" data-test="insurance-picker-row" data-uem-id="329-Conne">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ConnectiCare</span>
        </div>
    </div>
    <div data-row-key="536-Conso" data-test="insurance-picker-row" data-uem-id="536-Conso">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Consolidated Health Plans</span>
        </div>
    </div>
    <div data-row-key="1519-Const" data-test="insurance-picker-row" data-uem-id="1519-Const">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Constitution Life</span>
        </div>
    </div>
    <div data-row-key="395-Consu" data-test="insurance-picker-row" data-uem-id="395-Consu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Consumer Health Network</span>
        </div>
    </div>
    <div data-row-key="937-Contr" data-test="insurance-picker-row" data-uem-id="937-Contr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Contra Costa Health Plan</span>
        </div>
    </div>
    <div data-row-key="973-Cook " data-test="insurance-picker-row" data-uem-id="973-Cook ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cook Children's Health Plan</span>
        </div>
    </div>
    <div data-row-key="1067-Coord" data-test="insurance-picker-row" data-uem-id="1067-Coord">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Coordinated Care Health</span>
        </div>
    </div>
    <div data-row-key="396-Corve" data-test="insurance-picker-row" data-uem-id="396-Corve">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Corvel</span>
        </div>
    </div>
    <div data-row-key="1200-Count" data-test="insurance-picker-row" data-uem-id="1200-Count">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CountyCare (Cook County)</span>
        </div>
    </div>
    <div data-row-key="369-Coven" data-test="insurance-picker-row" data-uem-id="369-Coven">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Coventry Health Care</span>
        </div>
    </div>
    <div data-row-key="736-Cox H" data-test="insurance-picker-row" data-uem-id="736-Cox H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Cox HealthPlans</span>
        </div>
    </div>
    <div data-row-key="1464-Creat" data-test="insurance-picker-row" data-uem-id="1464-Creat">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Create</span>
        </div>
    </div>
    <div data-row-key="1310-Cryst" data-test="insurance-picker-row" data-uem-id="1310-Cryst">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Crystal Run Health Plans</span>
        </div>
    </div>
    <div data-row-key="753-Culin" data-test="insurance-picker-row" data-uem-id="753-Culin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Culinary Health Fund</span>
        </div>
    </div>
    <div data-row-key="1644-CuraL" data-test="insurance-picker-row" data-uem-id="1644-CuraL">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>CuraLinc Healthcare</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">D</div>
    <div data-row-key="901-DAKOT" data-test="insurance-picker-row" data-uem-id="901-DAKOT">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>DAKOTACARE</span>
        </div>
    </div>
    <div data-row-key="1072-DC Me" data-test="insurance-picker-row" data-uem-id="1072-DC Me">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>DC Medicaid</span>
        </div>
    </div>
    <div data-row-key="738-DMC C" data-test="insurance-picker-row" data-uem-id="738-DMC C">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>DMC Care</span>
        </div>
    </div>
    <div data-row-key="924-DOCS " data-test="insurance-picker-row" data-uem-id="924-DOCS ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>DOCS (Doctors of the Oregon South Coast)</span>
        </div>
    </div>
    <div data-row-key="1598-David" data-test="insurance-picker-row" data-uem-id="1598-David">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>DavidShield</span>
        </div>
    </div>
    <div data-row-key="538-Davis" data-test="insurance-picker-row" data-uem-id="538-Davis">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Davis Vision</span>
        </div>
    </div>
    <div data-row-key="1578-Deaco" data-test="insurance-picker-row" data-uem-id="1578-Deaco">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Deaconess Health Plans</span>
        </div>
    </div>
    <div data-row-key="1159-Dean " data-test="insurance-picker-row" data-uem-id="1159-Dean ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Dean Health Plan</span>
        </div>
    </div>
    <div data-row-key="1653-Dell " data-test="insurance-picker-row" data-uem-id="1653-Dell ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Dell Children's Health Plan</span>
        </div>
    </div>
    <div data-row-key="1316-Denve" data-test="insurance-picker-row" data-uem-id="1316-Denve">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Denver Health Medical Plan</span>
        </div>
    </div>
    <div data-row-key="1415-Depar" data-test="insurance-picker-row" data-uem-id="1415-Depar">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Department of Medical Assistance Services</span>
        </div>
    </div>
    <div data-row-key="899-Deser" data-test="insurance-picker-row" data-uem-id="899-Deser">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Deseret Mutual</span>
        </div>
    </div>
    <div data-row-key="397-Devon" data-test="insurance-picker-row" data-uem-id="397-Devon">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Devon Health Services</span>
        </div>
    </div>
    <div data-row-key="1547-Devot" data-test="insurance-picker-row" data-uem-id="1547-Devot">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Devoted Health</span>
        </div>
    </div>
    <div data-row-key="1523-Dimen" data-test="insurance-picker-row" data-uem-id="1523-Dimen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Dimension Health</span>
        </div>
    </div>
    <div data-row-key="1550-Docto" data-test="insurance-picker-row" data-uem-id="1550-Docto">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Doctors HealthCare Plans</span>
        </div>
    </div>
    <div data-row-key="1185-Drisc" data-test="insurance-picker-row" data-uem-id="1185-Drisc">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Driscoll Health Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">E</div>
    <div data-row-key="1468-EHP S" data-test="insurance-picker-row" data-uem-id="1468-EHP S">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EHP Significa</span>
        </div>
    </div>
    <div data-row-key="1143-EMI H" data-test="insurance-picker-row" data-uem-id="1143-EMI H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EMI Health</span>
        </div>
    </div>
    <div data-row-key="504-ESSEN" data-test="insurance-picker-row" data-uem-id="504-ESSEN">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ESSENCE Healthcare</span>
        </div>
    </div>
    <div data-row-key="1587-Easte" data-test="insurance-picker-row" data-uem-id="1587-Easte">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Eastern Oregon Coordinated Care Organization</span>
        </div>
    </div>
    <div data-row-key="938-Easy " data-test="insurance-picker-row" data-uem-id="938-Easy ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Easy Choice Health Plan (California)</span>
        </div>
    </div>
    <div data-row-key="693-Easy " data-test="insurance-picker-row" data-uem-id="693-Easy ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Easy Choice Health Plan of New York</span>
        </div>
    </div>
    <div data-row-key="1184-El Pa" data-test="insurance-picker-row" data-uem-id="1184-El Pa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>El Paso First Health Plans</span>
        </div>
    </div>
    <div data-row-key="435-Elder" data-test="insurance-picker-row" data-uem-id="435-Elder">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Elderplan</span>
        </div>
    </div>
    <div data-row-key="349-Emble" data-test="insurance-picker-row" data-uem-id="349-Emble">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EmblemHealth</span>
        </div>
    </div>
    <div data-row-key="338-Emble" data-test="insurance-picker-row" data-uem-id="338-Emble">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EmblemHealth (formerly known as GHI)</span>
        </div>
    </div>
    <div data-row-key="337-Emble" data-test="insurance-picker-row" data-uem-id="337-Emble">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EmblemHealth (formerly known as HIP)</span>
        </div>
    </div>
    <div data-row-key="1581-Emory" data-test="insurance-picker-row" data-uem-id="1581-Emory">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Emory Health Care Plan</span>
        </div>
    </div>
    <div data-row-key="336-Empir" data-test="insurance-picker-row" data-uem-id="336-Empir">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Empire Blue Cross Blue Shield (Health)</span>
        </div>
    </div>
    <div data-row-key="1355-Empir" data-test="insurance-picker-row" data-uem-id="1355-Empir">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Empire BlueCross BlueShield HealthPlus</span>
        </div>
    </div>
    <div data-row-key="354-Empir" data-test="insurance-picker-row" data-uem-id="354-Empir">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Empire Plan</span>
        </div>
    </div>
    <div data-row-key="1612-Empow" data-test="insurance-picker-row" data-uem-id="1612-Empow">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Empower Healthcare Solutions</span>
        </div>
    </div>
    <div data-row-key="459-Encor" data-test="insurance-picker-row" data-uem-id="459-Encor">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Encore Health Network</span>
        </div>
    </div>
    <div data-row-key="1056-Envol" data-test="insurance-picker-row" data-uem-id="1056-Envol">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Envolve Benefit Options</span>
        </div>
    </div>
    <div data-row-key="1456-Eon H" data-test="insurance-picker-row" data-uem-id="1456-Eon H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Eon Health</span>
        </div>
    </div>
    <div data-row-key="1086-Epic " data-test="insurance-picker-row" data-uem-id="1086-Epic ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Epic Hearing Health Care</span>
        </div>
    </div>
    <div data-row-key="1429-Erick" data-test="insurance-picker-row" data-uem-id="1429-Erick">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Erickson Advantage</span>
        </div>
    </div>
    <div data-row-key="1466-EverC" data-test="insurance-picker-row" data-uem-id="1466-EverC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EverCare</span>
        </div>
    </div>
    <div data-row-key="1134-Everg" data-test="insurance-picker-row" data-uem-id="1134-Everg">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Evergreen Health Cooperative</span>
        </div>
    </div>
    <div data-row-key="596-Evolu" data-test="insurance-picker-row" data-uem-id="596-Evolu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Evolutions Healthcare Systems</span>
        </div>
    </div>
    <div data-row-key="557-Excel" data-test="insurance-picker-row" data-uem-id="557-Excel">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Excellus Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="1607-Exper" data-test="insurance-picker-row" data-uem-id="1607-Exper">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Experience HealthND</span>
        </div>
    </div>
    <div data-row-key="1538-Exten" data-test="insurance-picker-row" data-uem-id="1538-Exten">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Extended Managed Long Term Care</span>
        </div>
    </div>
    <div data-row-key="539-EyeMe" data-test="insurance-picker-row" data-uem-id="539-EyeMe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>EyeMed</span>
        </div>
    </div>
    <div data-row-key="1059-Eyeto" data-test="insurance-picker-row" data-uem-id="1059-Eyeto">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Eyetopia Vision Care</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">F</div>
    <div data-row-key="618-Fallo" data-test="insurance-picker-row" data-uem-id="618-Fallo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Fallon Community Health Plan (FCHP)</span>
        </div>
    </div>
    <div data-row-key="740-Famil" data-test="insurance-picker-row" data-uem-id="740-Famil">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Family Health Network</span>
        </div>
    </div>
    <div data-row-key="806-Famil" data-test="insurance-picker-row" data-uem-id="806-Famil">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>FamilyCare Health Plans</span>
        </div>
    </div>
    <div data-row-key="343-Fidel" data-test="insurance-picker-row" data-uem-id="343-Fidel">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Fidelis Care (NY)</span>
        </div>
    </div>
    <div data-row-key="662-First" data-test="insurance-picker-row" data-uem-id="662-First">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>First Choice Health</span>
        </div>
    </div>
    <div data-row-key="1286-First" data-test="insurance-picker-row" data-uem-id="1286-First">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>First Choice Health Plan of Mississippi</span>
        </div>
    </div>
    <div data-row-key="309-First" data-test="insurance-picker-row" data-uem-id="309-First">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>First Health (Coventry Health Care)</span>
        </div>
    </div>
    <div data-row-key="745-First" data-test="insurance-picker-row" data-uem-id="745-First">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>FirstCare Health Plans</span>
        </div>
    </div>
    <div data-row-key="788-First" data-test="insurance-picker-row" data-uem-id="788-First">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>FirstCarolinaCare</span>
        </div>
    </div>
    <div data-row-key="638-Flori" data-test="insurance-picker-row" data-uem-id="638-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida Blue: Blue Cross Blue Shield of Florida</span>
        </div>
    </div>
    <div data-row-key="1561-Flori" data-test="insurance-picker-row" data-uem-id="1561-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida Community Care</span>
        </div>
    </div>
    <div data-row-key="911-Flori" data-test="insurance-picker-row" data-uem-id="911-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida Health Care Plans</span>
        </div>
    </div>
    <div data-row-key="1014-Flori" data-test="insurance-picker-row" data-uem-id="1014-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida Health Partners</span>
        </div>
    </div>
    <div data-row-key="1034-Flori" data-test="insurance-picker-row" data-uem-id="1034-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida Hospital Healthcare System (FHHS)</span>
        </div>
    </div>
    <div data-row-key="646-Flori" data-test="insurance-picker-row" data-uem-id="646-Flori">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Florida KidCare</span>
        </div>
    </div>
    <div data-row-key="1063-Fort " data-test="insurance-picker-row" data-uem-id="1063-Fort ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Fort Bend County Indigent Health Care</span>
        </div>
    </div>
    <div data-row-key="555-Forti" data-test="insurance-picker-row" data-uem-id="555-Forti">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Fortified Provider Network</span>
        </div>
    </div>
    <div data-row-key="692-Freed" data-test="insurance-picker-row" data-uem-id="692-Freed">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Freedom Health</span>
        </div>
    </div>
    <div data-row-key="1512-Frese" data-test="insurance-picker-row" data-uem-id="1512-Frese">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Fresenius Health Plans</span>
        </div>
    </div>
    <div data-row-key="1543-Frida" data-test="insurance-picker-row" data-uem-id="1543-Frida">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Friday Health Plans</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">G</div>
    <div data-row-key="310-GEHA" data-test="insurance-picker-row" data-uem-id="310-GEHA">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>GEHA</span>
        </div>
    </div>
    <div data-row-key="941-GEMCa" data-test="insurance-picker-row" data-uem-id="941-GEMCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>GEMCare Health Plan</span>
        </div>
    </div>
    <div data-row-key="313-GWH-C" data-test="insurance-picker-row" data-uem-id="313-GWH-C">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>GWH-Cigna (formerly Great West Healthcare)</span>
        </div>
    </div>
    <div data-row-key="400-Galax" data-test="insurance-picker-row" data-uem-id="400-Galax">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Galaxy Health</span>
        </div>
    </div>
    <div data-row-key="631-Gatew" data-test="insurance-picker-row" data-uem-id="631-Gatew">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Gateway Health</span>
        </div>
    </div>
    <div data-row-key="904-Geisi" data-test="insurance-picker-row" data-uem-id="904-Geisi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Geisinger Health Plan</span>
        </div>
    </div>
    <div data-row-key="1032-Gener" data-test="insurance-picker-row" data-uem-id="1032-Gener">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>General Vision Services (GVS)</span>
        </div>
    </div>
    <div data-row-key="1025-GeoBl" data-test="insurance-picker-row" data-uem-id="1025-GeoBl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>GeoBlue</span>
        </div>
    </div>
    <div data-row-key="768-Gilsb" data-test="insurance-picker-row" data-uem-id="768-Gilsb">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Gilsbar 360 Alliance</span>
        </div>
    </div>
    <div data-row-key="963-Globa" data-test="insurance-picker-row" data-uem-id="963-Globa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Global Health</span>
        </div>
    </div>
    <div data-row-key="1301-Gold " data-test="insurance-picker-row" data-uem-id="1301-Gold ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Gold Coast Health Plan</span>
        </div>
    </div>
    <div data-row-key="311-Golde" data-test="insurance-picker-row" data-uem-id="311-Golde">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Golden Rule</span>
        </div>
    </div>
    <div data-row-key="849-Golde" data-test="insurance-picker-row" data-uem-id="849-Golde">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Golden State Medicare Health Plan</span>
        </div>
    </div>
    <div data-row-key="1609-Green" data-test="insurance-picker-row" data-uem-id="1609-Green">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Green Mountain Care (Vermont)</span>
        </div>
    </div>
    <div data-row-key="671-Group" data-test="insurance-picker-row" data-uem-id="671-Group">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Group Health Cooperative</span>
        </div>
    </div>
    <div data-row-key="1172-Group" data-test="insurance-picker-row" data-uem-id="1172-Group">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Group Health Cooperative of Eau Claire</span>
        </div>
    </div>
    <div data-row-key="1162-Group" data-test="insurance-picker-row" data-uem-id="1162-Group">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Group Health Cooperative of South Central Wisconsin</span>
        </div>
    </div>
    <div data-row-key="312-Guard" data-test="insurance-picker-row" data-uem-id="312-Guard">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Guardian</span>
        </div>
    </div>
    <div data-row-key="1171-Gunde" data-test="insurance-picker-row" data-uem-id="1171-Gunde">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Gundersen Health Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">H</div>
    <div data-row-key="301-HAP (" data-test="insurance-picker-row" data-uem-id="301-HAP (">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HAP (Health Alliance Plan)</span>
        </div>
    </div>
    <div data-row-key="1055-HAP M" data-test="insurance-picker-row" data-uem-id="1055-HAP M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HAP Midwest Health Plan</span>
        </div>
    </div>
    <div data-row-key="454-HFN" data-test="insurance-picker-row" data-uem-id="454-HFN">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HFN</span>
        </div>
    </div>
    <div data-row-key="503-HFS M" data-test="insurance-picker-row" data-uem-id="503-HFS M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HFS Medical Benefits</span>
        </div>
    </div>
    <div data-row-key="972-HUSKY" data-test="insurance-picker-row" data-uem-id="972-HUSKY">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HUSKY Health</span>
        </div>
    </div>
    <div data-row-key="1576-Hamas" data-test="insurance-picker-row" data-uem-id="1576-Hamas">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hamaspik Choice</span>
        </div>
    </div>
    <div data-row-key="1326-Harke" data-test="insurance-picker-row" data-uem-id="1326-Harke">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Harken Health</span>
        </div>
    </div>
    <div data-row-key="470-Harmo" data-test="insurance-picker-row" data-uem-id="470-Harmo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Harmony Health Plan</span>
        </div>
    </div>
    <div data-row-key="615-Harva" data-test="insurance-picker-row" data-uem-id="615-Harva">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Harvard Pilgrim Health Care</span>
        </div>
    </div>
    <div data-row-key="1206-Hawai" data-test="insurance-picker-row" data-uem-id="1206-Hawai">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hawaii Medical Assurance Association (HMAA)</span>
        </div>
    </div>
    <div data-row-key="1198-Hawai" data-test="insurance-picker-row" data-uem-id="1198-Hawai">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hawaii Medical Service Association (HMSA)</span>
        </div>
    </div>
    <div data-row-key="1122-Healt" data-test="insurance-picker-row" data-uem-id="1122-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Alliance</span>
        </div>
    </div>
    <div data-row-key="586-Healt" data-test="insurance-picker-row" data-uem-id="586-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Choice Arizona</span>
        </div>
    </div>
    <div data-row-key="1322-Healt" data-test="insurance-picker-row" data-uem-id="1322-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health First (FL)</span>
        </div>
    </div>
    <div data-row-key="1453-Healt" data-test="insurance-picker-row" data-uem-id="1453-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health First Colorado</span>
        </div>
    </div>
    <div data-row-key="910-Healt" data-test="insurance-picker-row" data-uem-id="910-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health First Health Plans (Florida)</span>
        </div>
    </div>
    <div data-row-key="295-Healt" data-test="insurance-picker-row" data-uem-id="295-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Net</span>
        </div>
    </div>
    <div data-row-key="628-Healt" data-test="insurance-picker-row" data-uem-id="628-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health New England</span>
        </div>
    </div>
    <div data-row-key="1362-Healt" data-test="insurance-picker-row" data-uem-id="1362-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Partners Plans (Pennsylvania)</span>
        </div>
    </div>
    <div data-row-key="761-Healt" data-test="insurance-picker-row" data-uem-id="761-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Plan of Nevada</span>
        </div>
    </div>
    <div data-row-key="942-Healt" data-test="insurance-picker-row" data-uem-id="942-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Plan of San Joaquin</span>
        </div>
    </div>
    <div data-row-key="943-Healt" data-test="insurance-picker-row" data-uem-id="943-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Plan of San Mateo</span>
        </div>
    </div>
    <div data-row-key="429-Healt" data-test="insurance-picker-row" data-uem-id="429-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Plus</span>
        </div>
    </div>
    <div data-row-key="1590-Healt" data-test="insurance-picker-row" data-uem-id="1590-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Share of Oregon</span>
        </div>
    </div>
    <div data-row-key="858-Healt" data-test="insurance-picker-row" data-uem-id="858-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Health Sun</span>
        </div>
    </div>
    <div data-row-key="1106-Healt" data-test="insurance-picker-row" data-uem-id="1106-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthChoice Oklahoma</span>
        </div>
    </div>
    <div data-row-key="1619-Healt" data-test="insurance-picker-row" data-uem-id="1619-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthChoice of Michigan</span>
        </div>
    </div>
    <div data-row-key="366-Healt" data-test="insurance-picker-row" data-uem-id="366-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthFirst (NY)</span>
        </div>
    </div>
    <div data-row-key="1111-Healt" data-test="insurance-picker-row" data-uem-id="1111-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthNow</span>
        </div>
    </div>
    <div data-row-key="570-Healt" data-test="insurance-picker-row" data-uem-id="570-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthPartners</span>
        </div>
    </div>
    <div data-row-key="1435-Healt" data-test="insurance-picker-row" data-uem-id="1435-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthScope Benefits</span>
        </div>
    </div>
    <div data-row-key="452-Healt" data-test="insurance-picker-row" data-uem-id="452-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthSmart</span>
        </div>
    </div>
    <div data-row-key="1636-Healt" data-test="insurance-picker-row" data-uem-id="1636-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthTeam Advantage</span>
        </div>
    </div>
    <div data-row-key="1469-Healt" data-test="insurance-picker-row" data-uem-id="1469-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Healthcare Highways Health Plan</span>
        </div>
    </div>
    <div data-row-key="461-Healt" data-test="insurance-picker-row" data-uem-id="461-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Healthlink</span>
        </div>
    </div>
    <div data-row-key="851-Healt" data-test="insurance-picker-row" data-uem-id="851-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Healthy Texas Women</span>
        </div>
    </div>
    <div data-row-key="1139-Healt" data-test="insurance-picker-row" data-uem-id="1139-Healt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>HealthyCT</span>
        </div>
    </div>
    <div data-row-key="990-Hear " data-test="insurance-picker-row" data-uem-id="990-Hear ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hear In America</span>
        </div>
    </div>
    <div data-row-key="1603-Henne" data-test="insurance-picker-row" data-uem-id="1603-Henne">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hennepin Health</span>
        </div>
    </div>
    <div data-row-key="961-Herit" data-test="insurance-picker-row" data-uem-id="961-Herit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Heritage Vision Plans</span>
        </div>
    </div>
    <div data-row-key="564-Highm" data-test="insurance-picker-row" data-uem-id="564-Highm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Highmark Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="640-Highm" data-test="insurance-picker-row" data-uem-id="640-Highm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Highmark Blue Cross Blue Shield of Delaware</span>
        </div>
    </div>
    <div data-row-key="565-Highm" data-test="insurance-picker-row" data-uem-id="565-Highm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Highmark Blue Shield</span>
        </div>
    </div>
    <div data-row-key="1153-Highm" data-test="insurance-picker-row" data-uem-id="1153-Highm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Highmark BlueCross BlueShield of West Virginia</span>
        </div>
    </div>
    <div data-row-key="1212-Hills" data-test="insurance-picker-row" data-uem-id="1212-Hills">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hillsborough Health Care Plan</span>
        </div>
    </div>
    <div data-row-key="1226-Home " data-test="insurance-picker-row" data-uem-id="1226-Home ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Home State Health Plan</span>
        </div>
    </div>
    <div data-row-key="713-Homet" data-test="insurance-picker-row" data-uem-id="713-Homet">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hometown Health</span>
        </div>
    </div>
    <div data-row-key="314-Horiz" data-test="insurance-picker-row" data-uem-id="314-Horiz">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Horizon Blue Cross Blue Shield of New Jersey</span>
        </div>
    </div>
    <div data-row-key="1413-Horiz" data-test="insurance-picker-row" data-uem-id="1413-Horiz">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Horizon Blue Cross Blue Shield of New Jersey For Barnabas Health</span>
        </div>
    </div>
    <div data-row-key="1412-Horiz" data-test="insurance-picker-row" data-uem-id="1412-Horiz">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Horizon Blue Cross Blue Shield of New Jersey For Novartis</span>
        </div>
    </div>
    <div data-row-key="673-Horiz" data-test="insurance-picker-row" data-uem-id="673-Horiz">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Horizon NJ Health</span>
        </div>
    </div>
    <div data-row-key="351-Hudso" data-test="insurance-picker-row" data-uem-id="351-Hudso">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Hudson Health Plan</span>
        </div>
    </div>
    <div data-row-key="315-Human" data-test="insurance-picker-row" data-uem-id="315-Human">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Humana</span>
        </div>
    </div>
    <div data-row-key="1002-Human" data-test="insurance-picker-row" data-uem-id="1002-Human">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Humana Behavioral Health (LifeSynch)</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">I</div>
    <div data-row-key="1452-IHC H" data-test="insurance-picker-row" data-uem-id="1452-IHC H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>IHC Health Solutions</span>
        </div>
    </div>
    <div data-row-key="529-IMS (" data-test="insurance-picker-row" data-uem-id="529-IMS (">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>IMS (Independent Medical Systems)</span>
        </div>
    </div>
    <div data-row-key="1328-IU He" data-test="insurance-picker-row" data-uem-id="1328-IU He">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>IU Health Plans (Indiana University Health)</span>
        </div>
    </div>
    <div data-row-key="632-Illin" data-test="insurance-picker-row" data-uem-id="632-Illin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Illinicare Health</span>
        </div>
    </div>
    <div data-row-key="1517-Illin" data-test="insurance-picker-row" data-uem-id="1517-Illin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Illinois' Primary Care Case Management (PCCM)</span>
        </div>
    </div>
    <div data-row-key="1375-Imagi" data-test="insurance-picker-row" data-uem-id="1375-Imagi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Imagine Health</span>
        </div>
    </div>
    <div data-row-key="1532-Imper" data-test="insurance-picker-row" data-uem-id="1532-Imper">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Imperial Health Plan of California</span>
        </div>
    </div>
    <div data-row-key="1557-Imper" data-test="insurance-picker-row" data-uem-id="1557-Imper">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Imperial Insurance Company of Texas</span>
        </div>
    </div>
    <div data-row-key="563-Indep" data-test="insurance-picker-row" data-uem-id="563-Indep">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Independence Blue Cross</span>
        </div>
    </div>
    <div data-row-key="1099-Indep" data-test="insurance-picker-row" data-uem-id="1099-Indep">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Independence Care System</span>
        </div>
    </div>
    <div data-row-key="403-Indep" data-test="insurance-picker-row" data-uem-id="403-Indep">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Independent Health</span>
        </div>
    </div>
    <div data-row-key="1419-India" data-test="insurance-picker-row" data-uem-id="1419-India">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Indiana Medicaid</span>
        </div>
    </div>
    <div data-row-key="1359-Ingha" data-test="insurance-picker-row" data-uem-id="1359-Ingha">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ingham Health Plan</span>
        </div>
    </div>
    <div data-row-key="946-Inlan" data-test="insurance-picker-row" data-uem-id="946-Inlan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Inland Empire Health Plan</span>
        </div>
    </div>
    <div data-row-key="1142-Innov" data-test="insurance-picker-row" data-uem-id="1142-Innov">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Innovation Health</span>
        </div>
    </div>
    <div data-row-key="1340-Integ" data-test="insurance-picker-row" data-uem-id="1340-Integ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Integra</span>
        </div>
    </div>
    <div data-row-key="947-Inter" data-test="insurance-picker-row" data-uem-id="947-Inter">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Inter Valley Health Plan</span>
        </div>
    </div>
    <div data-row-key="1596-Inter" data-test="insurance-picker-row" data-uem-id="1596-Inter">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>InterCommunity Health Network CCO</span>
        </div>
    </div>
    <div data-row-key="1467-Inter" data-test="insurance-picker-row" data-uem-id="1467-Inter">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Intergroup Services</span>
        </div>
    </div>
    <div data-row-key="1189-Iowa " data-test="insurance-picker-row" data-uem-id="1189-Iowa ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Iowa MediPASS</span>
        </div>
    </div>
    <div data-row-key="1601-Iowa " data-test="insurance-picker-row" data-uem-id="1601-Iowa ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Iowa Total Care</span>
        </div>
    </div>
    <div data-row-key="1257-Itasc" data-test="insurance-picker-row" data-uem-id="1257-Itasc">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Itasca Medical Care</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">J</div>
    <div data-row-key="1591-Jacks" data-test="insurance-picker-row" data-uem-id="1591-Jacks">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Jackson Care Connect</span>
        </div>
    </div>
    <div data-row-key="1531-Jacks" data-test="insurance-picker-row" data-uem-id="1531-Jacks">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Jackson Health Plan</span>
        </div>
    </div>
    <div data-row-key="1562-Jai M" data-test="insurance-picker-row" data-uem-id="1562-Jai M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Jai Medical Systems</span>
        </div>
    </div>
    <div data-row-key="624-Johns" data-test="insurance-picker-row" data-uem-id="624-Johns">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Johns Hopkins Employer Health Programs</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">K</div>
    <div data-row-key="773-KPS H" data-test="insurance-picker-row" data-uem-id="773-KPS H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>KPS Health Plans</span>
        </div>
    </div>
    <div data-row-key="385-Kaise" data-test="insurance-picker-row" data-uem-id="385-Kaise">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Kaiser Permanente</span>
        </div>
    </div>
    <div data-row-key="741-Kansa" data-test="insurance-picker-row" data-uem-id="741-Kansa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Kansas HealthWave</span>
        </div>
    </div>
    <div data-row-key="1472-Kansa" data-test="insurance-picker-row" data-uem-id="1472-Kansa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Kansas Superior Select</span>
        </div>
    </div>
    <div data-row-key="1186-Kelse" data-test="insurance-picker-row" data-uem-id="1186-Kelse">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>KelseyCare Advantage</span>
        </div>
    </div>
    <div data-row-key="569-Keyst" data-test="insurance-picker-row" data-uem-id="569-Keyst">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Keystone First</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">L</div>
    <div data-row-key="949-L.A. " data-test="insurance-picker-row" data-uem-id="949-L.A. ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>L.A. Care Health Plan</span>
        </div>
    </div>
    <div data-row-key="1530-L.A. " data-test="insurance-picker-row" data-uem-id="1530-L.A. ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>L.A. Care Health Plan</span>
        </div>
    </div>
    <div data-row-key="1128-Land " data-test="insurance-picker-row" data-uem-id="1128-Land ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Land of Lincoln Health</span>
        </div>
    </div>
    <div data-row-key="1227-Landm" data-test="insurance-picker-row" data-uem-id="1227-Landm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Landmark Healthplan</span>
        </div>
    </div>
    <div data-row-key="1574-Lasso" data-test="insurance-picker-row" data-uem-id="1574-Lasso">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lasso Healthcare</span>
        </div>
    </div>
    <div data-row-key="1397-Legac" data-test="insurance-picker-row" data-uem-id="1397-Legac">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Legacy Health</span>
        </div>
    </div>
    <div data-row-key="1539-Lehig" data-test="insurance-picker-row" data-uem-id="1539-Lehig">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lehigh Valley Health Network Health Plan</span>
        </div>
    </div>
    <div data-row-key="1524-Leon " data-test="insurance-picker-row" data-uem-id="1524-Leon ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Leon Medical Centers Health Plans</span>
        </div>
    </div>
    <div data-row-key="427-Liber" data-test="insurance-picker-row" data-uem-id="427-Liber">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Liberty Health Advantage</span>
        </div>
    </div>
    <div data-row-key="1583-Liber" data-test="insurance-picker-row" data-uem-id="1583-Liber">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Liberty HealthShare</span>
        </div>
    </div>
    <div data-row-key="484-Liber" data-test="insurance-picker-row" data-uem-id="484-Liber">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Liberty Mutual</span>
        </div>
    </div>
    <div data-row-key="659-LifeW" data-test="insurance-picker-row" data-uem-id="659-LifeW">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>LifeWise</span>
        </div>
    </div>
    <div data-row-key="1629-Lifes" data-test="insurance-picker-row" data-uem-id="1629-Lifes">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lifestyle Health</span>
        </div>
    </div>
    <div data-row-key="1339-Light" data-test="insurance-picker-row" data-uem-id="1339-Light">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lighthouse Guild</span>
        </div>
    </div>
    <div data-row-key="1563-Light" data-test="insurance-picker-row" data-uem-id="1563-Light">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lighthouse Health Plan</span>
        </div>
    </div>
    <div data-row-key="1066-Linco" data-test="insurance-picker-row" data-uem-id="1066-Linco">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lincoln Financial Group</span>
        </div>
    </div>
    <div data-row-key="1643-Live3" data-test="insurance-picker-row" data-uem-id="1643-Live3">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Live360 Health Plan</span>
        </div>
    </div>
    <div data-row-key="1553-Longe" data-test="insurance-picker-row" data-uem-id="1553-Longe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Longevity Health Plan</span>
        </div>
    </div>
    <div data-row-key="1244-Louis" data-test="insurance-picker-row" data-uem-id="1244-Louis">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Louisiana Healthcare Connections</span>
        </div>
    </div>
    <div data-row-key="465-Luthe" data-test="insurance-picker-row" data-uem-id="465-Luthe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Lutheran Preferred</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">M</div>
    <div data-row-key="1003-MCM M" data-test="insurance-picker-row" data-uem-id="1003-MCM M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MCM Maxcare</span>
        </div>
    </div>
    <div data-row-key="1114-MDwis" data-test="insurance-picker-row" data-uem-id="1114-MDwis">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MDwise</span>
        </div>
    </div>
    <div data-row-key="1061-MHNet" data-test="insurance-picker-row" data-uem-id="1061-MHNet">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MHNet Behavioral Health</span>
        </div>
    </div>
    <div data-row-key="1600-MINES" data-test="insurance-picker-row" data-uem-id="1600-MINES">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MINES &amp; Associates</span>
        </div>
    </div>
    <div data-row-key="1548-MMM o" data-test="insurance-picker-row" data-uem-id="1548-MMM o">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MMM of Florida (Medicare and Much More)</span>
        </div>
    </div>
    <div data-row-key="1031-MO He" data-test="insurance-picker-row" data-uem-id="1031-MO He">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MO HealthNet</span>
        </div>
    </div>
    <div data-row-key="1533-MOAA " data-test="insurance-picker-row" data-uem-id="1533-MOAA ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MOAA (Miltary Officers Association of America)</span>
        </div>
    </div>
    <div data-row-key="353-MVP H" data-test="insurance-picker-row" data-uem-id="353-MVP H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MVP Health Care</span>
        </div>
    </div>
    <div data-row-key="592-Magel" data-test="insurance-picker-row" data-uem-id="592-Magel">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Magellan Health</span>
        </div>
    </div>
    <div data-row-key="339-Magna" data-test="insurance-picker-row" data-uem-id="339-Magna">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MagnaCare</span>
        </div>
    </div>
    <div data-row-key="1245-Magno" data-test="insurance-picker-row" data-uem-id="1245-Magno">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Magnolia Health Plan</span>
        </div>
    </div>
    <div data-row-key="445-Mail " data-test="insurance-picker-row" data-uem-id="445-Mail ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mail Handlers Benefit Plan</span>
        </div>
    </div>
    <div data-row-key="1192-Maine" data-test="insurance-picker-row" data-uem-id="1192-Maine">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MaineCare</span>
        </div>
    </div>
    <div data-row-key="558-Manag" data-test="insurance-picker-row" data-uem-id="558-Manag">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Managed Health Network (MHN)</span>
        </div>
    </div>
    <div data-row-key="1246-Manag" data-test="insurance-picker-row" data-uem-id="1246-Manag">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Managed Health Services (Indiana)</span>
        </div>
    </div>
    <div data-row-key="1247-Manag" data-test="insurance-picker-row" data-uem-id="1247-Manag">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Managed Health Services (Wisconsin)</span>
        </div>
    </div>
    <div data-row-key="1437-Manag" data-test="insurance-picker-row" data-uem-id="1437-Manag">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Managed HealthCare Northwest</span>
        </div>
    </div>
    <div data-row-key="774-March" data-test="insurance-picker-row" data-uem-id="774-March">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>March Vision Care</span>
        </div>
    </div>
    <div data-row-key="584-Maric" data-test="insurance-picker-row" data-uem-id="584-Maric">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Maricopa Health Plan</span>
        </div>
    </div>
    <div data-row-key="1268-Marti" data-test="insurance-picker-row" data-uem-id="1268-Marti">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Martin's Point HealthCare</span>
        </div>
    </div>
    <div data-row-key="642-Maryl" data-test="insurance-picker-row" data-uem-id="642-Maryl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Maryland Medical Assistance (Medicaid)</span>
        </div>
    </div>
    <div data-row-key="559-Maryl" data-test="insurance-picker-row" data-uem-id="559-Maryl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Maryland Physicians Care</span>
        </div>
    </div>
    <div data-row-key="669-MassH" data-test="insurance-picker-row" data-uem-id="669-MassH">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MassHealth</span>
        </div>
    </div>
    <div data-row-key="1541-Massa" data-test="insurance-picker-row" data-uem-id="1541-Massa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Massachusetts Laborers' Health &amp; Welfare Fund</span>
        </div>
    </div>
    <div data-row-key="1102-Mayo " data-test="insurance-picker-row" data-uem-id="1102-Mayo ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mayo Medical Plan</span>
        </div>
    </div>
    <div data-row-key="1135-McLar" data-test="insurance-picker-row" data-uem-id="1135-McLar">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>McLaren Health Plan</span>
        </div>
    </div>
    <div data-row-key="775-MedSt" data-test="insurance-picker-row" data-uem-id="775-MedSt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MedStar Family Choice</span>
        </div>
    </div>
    <div data-row-key="1230-MedSt" data-test="insurance-picker-row" data-uem-id="1230-MedSt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MedStar Select</span>
        </div>
    </div>
    <div data-row-key="450-Medi-" data-test="insurance-picker-row" data-uem-id="450-Medi-">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medi-Cal</span>
        </div>
    </div>
    <div data-row-key="1068-MediG" data-test="insurance-picker-row" data-uem-id="1068-MediG">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MediGold</span>
        </div>
    </div>
    <div data-row-key="966-MediP" data-test="insurance-picker-row" data-uem-id="966-MediP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MediPass</span>
        </div>
    </div>
    <div data-row-key="704-Medic" data-test="insurance-picker-row" data-uem-id="704-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medica</span>
        </div>
    </div>
    <div data-row-key="664-Medic" data-test="insurance-picker-row" data-uem-id="664-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medica HealthCare Plans (Florida)</span>
        </div>
    </div>
    <div data-row-key="358-Medic" data-test="insurance-picker-row" data-uem-id="358-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medicaid</span>
        </div>
    </div>
    <div data-row-key="683-Medic" data-test="insurance-picker-row" data-uem-id="683-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medical Eye Services (MES Vision)</span>
        </div>
    </div>
    <div data-row-key="600-Medic" data-test="insurance-picker-row" data-uem-id="600-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medical Mutual</span>
        </div>
    </div>
    <div data-row-key="356-Medic" data-test="insurance-picker-row" data-uem-id="356-Medic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Medicare</span>
        </div>
    </div>
    <div data-row-key="1622-Memor" data-test="insurance-picker-row" data-uem-id="1622-Memor">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Memorial Healthcare System</span>
        </div>
    </div>
    <div data-row-key="779-Memor" data-test="insurance-picker-row" data-uem-id="779-Memor">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Memorial Hermann</span>
        </div>
    </div>
    <div data-row-key="578-Mercy" data-test="insurance-picker-row" data-uem-id="578-Mercy">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mercy Care</span>
        </div>
    </div>
    <div data-row-key="965-Merid" data-test="insurance-picker-row" data-uem-id="965-Merid">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Meridian Health Plan</span>
        </div>
    </div>
    <div data-row-key="1488-Merit" data-test="insurance-picker-row" data-uem-id="1488-Merit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Meritain Health</span>
        </div>
    </div>
    <div data-row-key="883-MetLi" data-test="insurance-picker-row" data-uem-id="883-MetLi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MetLife</span>
        </div>
    </div>
    <div data-row-key="1447-Metro" data-test="insurance-picker-row" data-uem-id="1447-Metro">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MetroHealth</span>
        </div>
    </div>
    <div data-row-key="409-Metro" data-test="insurance-picker-row" data-uem-id="409-Metro">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MetroPlus Health Plan</span>
        </div>
    </div>
    <div data-row-key="850-Metro" data-test="insurance-picker-row" data-uem-id="850-Metro">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Metropolitan Health Plan</span>
        </div>
    </div>
    <div data-row-key="1549-Miami" data-test="insurance-picker-row" data-uem-id="1549-Miami">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Miami Children's Health Plan</span>
        </div>
    </div>
    <div data-row-key="1516-Michi" data-test="insurance-picker-row" data-uem-id="1516-Michi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Michigan Complete Health</span>
        </div>
    </div>
    <div data-row-key="1640-Michi" data-test="insurance-picker-row" data-uem-id="1640-Michi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Michigan No-Fault</span>
        </div>
    </div>
    <div data-row-key="1262-Minut" data-test="insurance-picker-row" data-uem-id="1262-Minut">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Minuteman Health</span>
        </div>
    </div>
    <div data-row-key="1399-Missi" data-test="insurance-picker-row" data-uem-id="1399-Missi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mississippi Division of Medicaid</span>
        </div>
    </div>
    <div data-row-key="780-Misso" data-test="insurance-picker-row" data-uem-id="780-Misso">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Missouri Care</span>
        </div>
    </div>
    <div data-row-key="706-Moda " data-test="insurance-picker-row" data-uem-id="706-Moda ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Moda Health</span>
        </div>
    </div>
    <div data-row-key="675-Molin" data-test="insurance-picker-row" data-uem-id="675-Molin">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Molina Healthcare</span>
        </div>
    </div>
    <div data-row-key="1327-Monta" data-test="insurance-picker-row" data-uem-id="1327-Monta">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Montana Health Cooperative</span>
        </div>
    </div>
    <div data-row-key="1556-Monte" data-test="insurance-picker-row" data-uem-id="1556-Monte">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Montefiore HMO</span>
        </div>
    </div>
    <div data-row-key="1650-MoreC" data-test="insurance-picker-row" data-uem-id="1650-MoreC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>MoreCare</span>
        </div>
    </div>
    <div data-row-key="1319-Mount" data-test="insurance-picker-row" data-uem-id="1319-Mount">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mountain Health Co-Op</span>
        </div>
    </div>
    <div data-row-key="360-Multi" data-test="insurance-picker-row" data-uem-id="360-Multi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Multiplan PHCS</span>
        </div>
    </div>
    <div data-row-key="318-Mutua" data-test="insurance-picker-row" data-uem-id="318-Mutua">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Mutual of Omaha</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">N</div>
    <div data-row-key="1475-NALC " data-test="insurance-picker-row" data-uem-id="1475-NALC ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NALC Health Benefit Plan</span>
        </div>
    </div>
    <div data-row-key="1579-NECA/" data-test="insurance-picker-row" data-uem-id="1579-NECA/">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NECA/IBEW Family Medical Care Plan</span>
        </div>
    </div>
    <div data-row-key="502-NY St" data-test="insurance-picker-row" data-uem-id="502-NY St">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NY State No-Fault</span>
        </div>
    </div>
    <div data-row-key="1309-NY: Y" data-test="insurance-picker-row" data-uem-id="1309-NY: Y">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NY: YourCare Health Plan</span>
        </div>
    </div>
    <div data-row-key="1428-Natio" data-test="insurance-picker-row" data-uem-id="1428-Natio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>National Congress of Employers (NCE)</span>
        </div>
    </div>
    <div data-row-key="547-Natio" data-test="insurance-picker-row" data-uem-id="547-Natio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>National Vision Administrators</span>
        </div>
    </div>
    <div data-row-key="296-Natio" data-test="insurance-picker-row" data-uem-id="296-Natio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Nationwide</span>
        </div>
    </div>
    <div data-row-key="1451-Navaj" data-test="insurance-picker-row" data-uem-id="1451-Navaj">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Navajo Nation</span>
        </div>
    </div>
    <div data-row-key="1616-Nebra" data-test="insurance-picker-row" data-uem-id="1616-Nebra">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Nebraska Total Care (Heritage Health)</span>
        </div>
    </div>
    <div data-row-key="617-Neigh" data-test="insurance-picker-row" data-uem-id="617-Neigh">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Neighborhood Health Plan (Massachusetts)</span>
        </div>
    </div>
    <div data-row-key="1023-Neigh" data-test="insurance-picker-row" data-uem-id="1023-Neigh">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Neighborhood Health Plan of Rhode Island</span>
        </div>
    </div>
    <div data-row-key="1321-Netwo" data-test="insurance-picker-row" data-uem-id="1321-Netwo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Network Health Plan</span>
        </div>
    </div>
    <div data-row-key="787-Nevad" data-test="insurance-picker-row" data-uem-id="787-Nevad">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Nevada Preferred</span>
        </div>
    </div>
    <div data-row-key="1430-New D" data-test="insurance-picker-row" data-uem-id="1430-New D">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>New Directions Behavioral Health</span>
        </div>
    </div>
    <div data-row-key="1267-New H" data-test="insurance-picker-row" data-uem-id="1267-New H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>New Hampshire Healthy Families</span>
        </div>
    </div>
    <div data-row-key="1119-New M" data-test="insurance-picker-row" data-uem-id="1119-New M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>New Mexico Health Connections</span>
        </div>
    </div>
    <div data-row-key="1372-New Y" data-test="insurance-picker-row" data-uem-id="1372-New Y">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>New York Hotel Trades Council</span>
        </div>
    </div>
    <div data-row-key="1473-NextL" data-test="insurance-picker-row" data-uem-id="1473-NextL">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NextLevelHealth</span>
        </div>
    </div>
    <div data-row-key="782-Nippo" data-test="insurance-picker-row" data-uem-id="782-Nippo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Nippon Life Benefits</span>
        </div>
    </div>
    <div data-row-key="834-North" data-test="insurance-picker-row" data-uem-id="834-North">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>North Carolina Health Choice (NCHC) for Children</span>
        </div>
    </div>
    <div data-row-key="1138-North" data-test="insurance-picker-row" data-uem-id="1138-North">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>North Shore LIJ CareConnect</span>
        </div>
    </div>
    <div data-row-key="1608-North" data-test="insurance-picker-row" data-uem-id="1608-North">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Northland PACE</span>
        </div>
    </div>
    <div data-row-key="597-NovaN" data-test="insurance-picker-row" data-uem-id="597-NovaN">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>NovaNet</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">O</div>
    <div data-row-key="927-OHMS " data-test="insurance-picker-row" data-uem-id="927-OHMS ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>OHMS (Oregon Health Management Services)</span>
        </div>
    </div>
    <div data-row-key="1088-OSU H" data-test="insurance-picker-row" data-uem-id="1088-OSU H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>OSU Health Plan</span>
        </div>
    </div>
    <div data-row-key="1432-Ohara" data-test="insurance-picker-row" data-uem-id="1432-Ohara">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ohara, LLC</span>
        </div>
    </div>
    <div data-row-key="1365-Ohio " data-test="insurance-picker-row" data-uem-id="1365-Ohio ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ohio Health Choice</span>
        </div>
    </div>
    <div data-row-key="1614-OhioH" data-test="insurance-picker-row" data-uem-id="1614-OhioH">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>OhioHealthy</span>
        </div>
    </div>
    <div data-row-key="1476-On Lo" data-test="insurance-picker-row" data-uem-id="1476-On Lo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>On Lok Lifeways</span>
        </div>
    </div>
    <div data-row-key="950-On Lo" data-test="insurance-picker-row" data-uem-id="950-On Lo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>On Lok Lifeways (PACE)</span>
        </div>
    </div>
    <div data-row-key="367-OneNe" data-test="insurance-picker-row" data-uem-id="367-OneNe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>OneNet PPO</span>
        </div>
    </div>
    <div data-row-key="783-Optic" data-test="insurance-picker-row" data-uem-id="783-Optic">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Opticare of Utah</span>
        </div>
    </div>
    <div data-row-key="320-Optim" data-test="insurance-picker-row" data-uem-id="320-Optim">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Optima Health</span>
        </div>
    </div>
    <div data-row-key="884-Optim" data-test="insurance-picker-row" data-uem-id="884-Optim">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Optimum HealthCare</span>
        </div>
    </div>
    <div data-row-key="651-Optum" data-test="insurance-picker-row" data-uem-id="651-Optum">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Optum Health</span>
        </div>
    </div>
    <div data-row-key="1127-Oscar" data-test="insurance-picker-row" data-uem-id="1127-Oscar">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Oscar Health Insurance Co.</span>
        </div>
    </div>
    <div data-row-key="1075-Oxfor" data-test="insurance-picker-row" data-uem-id="1075-Oxfor">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Oxford (UnitedHealthcare)</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">P</div>
    <div data-row-key="1615-PA He" data-test="insurance-picker-row" data-uem-id="1615-PA He">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PA Health and Wellness</span>
        </div>
    </div>
    <div data-row-key="1481-PA Me" data-test="insurance-picker-row" data-uem-id="1481-PA Me">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PA Medical Assistance</span>
        </div>
    </div>
    <div data-row-key="534-PBA (" data-test="insurance-picker-row" data-uem-id="534-PBA (">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PBA (Patrolmen's Benefit Association)</span>
        </div>
    </div>
    <div data-row-key="344-POMCO" data-test="insurance-picker-row" data-uem-id="344-POMCO">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>POMCO</span>
        </div>
    </div>
    <div data-row-key="1229-Pacif" data-test="insurance-picker-row" data-uem-id="1229-Pacif">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Pacific Health Alliance</span>
        </div>
    </div>
    <div data-row-key="697-Pacif" data-test="insurance-picker-row" data-uem-id="697-Pacif">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PacificSource Health Plans</span>
        </div>
    </div>
    <div data-row-key="430-Palme" data-test="insurance-picker-row" data-uem-id="430-Palme">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Palmetto GBA</span>
        </div>
    </div>
    <div data-row-key="1360-Pan-A" data-test="insurance-picker-row" data-uem-id="1360-Pan-A">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Pan-American Life Insurance Group</span>
        </div>
    </div>
    <div data-row-key="1599-Parad" data-test="insurance-picker-row" data-uem-id="1599-Parad">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Paradigm Senior Care Advantage</span>
        </div>
    </div>
    <div data-row-key="1092-Param" data-test="insurance-picker-row" data-uem-id="1092-Param">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Paramount Healthcare</span>
        </div>
    </div>
    <div data-row-key="514-Parkl" data-test="insurance-picker-row" data-uem-id="514-Parkl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Parkland Community Health Plan</span>
        </div>
    </div>
    <div data-row-key="1417-Parkv" data-test="insurance-picker-row" data-uem-id="1417-Parkv">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Parkview Total Health</span>
        </div>
    </div>
    <div data-row-key="1427-Partn" data-test="insurance-picker-row" data-uem-id="1427-Partn">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Partners Health Plan</span>
        </div>
    </div>
    <div data-row-key="951-Partn" data-test="insurance-picker-row" data-uem-id="951-Partn">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Partnership HealthPlan of California</span>
        </div>
    </div>
    <div data-row-key="1174-Passp" data-test="insurance-picker-row" data-uem-id="1174-Passp">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Passport Health Plan (Kentucky)</span>
        </div>
    </div>
    <div data-row-key="1205-Passp" data-test="insurance-picker-row" data-uem-id="1205-Passp">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Passport To Health (Montana Medicaid)</span>
        </div>
    </div>
    <div data-row-key="1112-Patie" data-test="insurance-picker-row" data-uem-id="1112-Patie">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Patient 1st (Alabama Medicaid)</span>
        </div>
    </div>
    <div data-row-key="599-Peach" data-test="insurance-picker-row" data-uem-id="599-Peach">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Peach State Health Plan</span>
        </div>
    </div>
    <div data-row-key="571-Peach" data-test="insurance-picker-row" data-uem-id="571-Peach">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PeachCare for Kids</span>
        </div>
    </div>
    <div data-row-key="1305-PennC" data-test="insurance-picker-row" data-uem-id="1305-PennC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PennCare</span>
        </div>
    </div>
    <div data-row-key="1165-Peopl" data-test="insurance-picker-row" data-uem-id="1165-Peopl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Peoples Health</span>
        </div>
    </div>
    <div data-row-key="582-Phoen" data-test="insurance-picker-row" data-uem-id="582-Phoen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Phoenix Health Plan</span>
        </div>
    </div>
    <div data-row-key="1191-Physi" data-test="insurance-picker-row" data-uem-id="1191-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Physician Assured Access System</span>
        </div>
    </div>
    <div data-row-key="1374-Physi" data-test="insurance-picker-row" data-uem-id="1374-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Physician Benefits Trust</span>
        </div>
    </div>
    <div data-row-key="1584-Physi" data-test="insurance-picker-row" data-uem-id="1584-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Physicians Eyecare Plan</span>
        </div>
    </div>
    <div data-row-key="1222-Physi" data-test="insurance-picker-row" data-uem-id="1222-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Physicians Health Plan</span>
        </div>
    </div>
    <div data-row-key="1148-Physi" data-test="insurance-picker-row" data-uem-id="1148-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Physicians Health Plan of Northern Indiana, Inc.</span>
        </div>
    </div>
    <div data-row-key="1443-Physi" data-test="insurance-picker-row" data-uem-id="1443-Physi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PhysiciansCare</span>
        </div>
    </div>
    <div data-row-key="1149-Piedm" data-test="insurance-picker-row" data-uem-id="1149-Piedm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Piedmont Community Health Plan</span>
        </div>
    </div>
    <div data-row-key="1211-Piedm" data-test="insurance-picker-row" data-uem-id="1211-Piedm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Piedmont WellStar Health Plans</span>
        </div>
    </div>
    <div data-row-key="784-Posit" data-test="insurance-picker-row" data-uem-id="784-Posit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Positive Health Care</span>
        </div>
    </div>
    <div data-row-key="1347-Prefe" data-test="insurance-picker-row" data-uem-id="1347-Prefe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Preferential Care Network</span>
        </div>
    </div>
    <div data-row-key="681-Prefe" data-test="insurance-picker-row" data-uem-id="681-Prefe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Preferred Care Partners</span>
        </div>
    </div>
    <div data-row-key="699-Prefe" data-test="insurance-picker-row" data-uem-id="699-Prefe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PreferredOne</span>
        </div>
    </div>
    <div data-row-key="601-Preme" data-test="insurance-picker-row" data-uem-id="601-Preme">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Premera Blue Cross</span>
        </div>
    </div>
    <div data-row-key="1329-Premi" data-test="insurance-picker-row" data-uem-id="1329-Premi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Premier Health Plan</span>
        </div>
    </div>
    <div data-row-key="604-Presb" data-test="insurance-picker-row" data-uem-id="604-Presb">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Presbyterian Health Plan/Presbyterian Insurance Company</span>
        </div>
    </div>
    <div data-row-key="887-Prest" data-test="insurance-picker-row" data-uem-id="887-Prest">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Prestige Health Choice</span>
        </div>
    </div>
    <div data-row-key="1203-Prima" data-test="insurance-picker-row" data-uem-id="1203-Prima">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Primary Care Case Management (North Dakota Medicaid)</span>
        </div>
    </div>
    <div data-row-key="748-Prime" data-test="insurance-picker-row" data-uem-id="748-Prime">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Prime Health Services, Inc</span>
        </div>
    </div>
    <div data-row-key="1623-Prime" data-test="insurance-picker-row" data-uem-id="1623-Prime">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Prime Healthcare</span>
        </div>
    </div>
    <div data-row-key="1253-Prime" data-test="insurance-picker-row" data-uem-id="1253-Prime">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>PrimeWest Health</span>
        </div>
    </div>
    <div data-row-key="321-Princ" data-test="insurance-picker-row" data-uem-id="321-Princ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Principal Financial Group</span>
        </div>
    </div>
    <div data-row-key="786-Prior" data-test="insurance-picker-row" data-uem-id="786-Prior">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Priority Health</span>
        </div>
    </div>
    <div data-row-key="447-Prior" data-test="insurance-picker-row" data-uem-id="447-Prior">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Priority Partners</span>
        </div>
    </div>
    <div data-row-key="1645-ProCa" data-test="insurance-picker-row" data-uem-id="1645-ProCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ProCare Advantage</span>
        </div>
    </div>
    <div data-row-key="867-Progr" data-test="insurance-picker-row" data-uem-id="867-Progr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Progressive</span>
        </div>
    </div>
    <div data-row-key="714-Promi" data-test="insurance-picker-row" data-uem-id="714-Promi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Prominence Health Plan</span>
        </div>
    </div>
    <div data-row-key="979-Provi" data-test="insurance-picker-row" data-uem-id="979-Provi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ProviDRs Care (WPPA)</span>
        </div>
    </div>
    <div data-row-key="331-Provi" data-test="insurance-picker-row" data-uem-id="331-Provi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Providence Health Plans</span>
        </div>
    </div>
    <div data-row-key="1571-Provi" data-test="insurance-picker-row" data-uem-id="1571-Provi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Provider Partners Health Plan</span>
        </div>
    </div>
    <div data-row-key="500-Publi" data-test="insurance-picker-row" data-uem-id="500-Publi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Public Aid (Illinois Medicaid)</span>
        </div>
    </div>
    <div data-row-key="1157-Publi" data-test="insurance-picker-row" data-uem-id="1157-Publi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Public Employees Health Program (PEHP)</span>
        </div>
    </div>
    <div data-row-key="1633-Puget" data-test="insurance-picker-row" data-uem-id="1633-Puget">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Puget Sound Electrical Workers Trusts</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">Q</div>
    <div data-row-key="347-QualC" data-test="insurance-picker-row" data-uem-id="347-QualC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>QualCare</span>
        </div>
    </div>
    <div data-row-key="1188-QualC" data-test="insurance-picker-row" data-uem-id="1188-QualC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>QualChoice Arkansas</span>
        </div>
    </div>
    <div data-row-key="912-Quali" data-test="insurance-picker-row" data-uem-id="912-Quali">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Quality Health Plans of New York</span>
        </div>
    </div>
    <div data-row-key="1558-Quart" data-test="insurance-picker-row" data-uem-id="1558-Quart">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Quartz</span>
        </div>
    </div>
    <div data-row-key="1624-Quest" data-test="insurance-picker-row" data-uem-id="1624-Quest">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Quest Behavioral Health</span>
        </div>
    </div>
    <div data-row-key="1559-Quikt" data-test="insurance-picker-row" data-uem-id="1559-Quikt">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Quiktrip</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">R</div>
    <div data-row-key="1509-RLI I" data-test="insurance-picker-row" data-uem-id="1509-RLI I">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>RLI Indemnity Company</span>
        </div>
    </div>
    <div data-row-key="602-Regen" data-test="insurance-picker-row" data-uem-id="602-Regen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Regence Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="827-Regen" data-test="insurance-picker-row" data-uem-id="827-Regen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Regence Blue Shield of Washington</span>
        </div>
    </div>
    <div data-row-key="1511-Regen" data-test="insurance-picker-row" data-uem-id="1511-Regen">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Regent Insurance</span>
        </div>
    </div>
    <div data-row-key="1628-Relia" data-test="insurance-picker-row" data-uem-id="1628-Relia">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Reliance Medicare Advantage</span>
        </div>
    </div>
    <div data-row-key="1627-Renai" data-test="insurance-picker-row" data-uem-id="1627-Renai">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Renaissance</span>
        </div>
    </div>
    <div data-row-key="1510-Repub" data-test="insurance-picker-row" data-uem-id="1510-Repub">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Republic-Franklin Insurance</span>
        </div>
    </div>
    <div data-row-key="1357-River" data-test="insurance-picker-row" data-uem-id="1357-River">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>RiverLink Health</span>
        </div>
    </div>
    <div data-row-key="1425-River" data-test="insurance-picker-row" data-uem-id="1425-River">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>RiverSpring Health Plans</span>
        </div>
    </div>
    <div data-row-key="1551-River" data-test="insurance-picker-row" data-uem-id="1551-River">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>RiverSpring at Home</span>
        </div>
    </div>
    <div data-row-key="1264-River" data-test="insurance-picker-row" data-uem-id="1264-River">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Riverside Health</span>
        </div>
    </div>
    <div data-row-key="1567-Rocke" data-test="insurance-picker-row" data-uem-id="1567-Rocke">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Rockefeller Health Plan</span>
        </div>
    </div>
    <div data-row-key="695-Rocky" data-test="insurance-picker-row" data-uem-id="695-Rocky">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Rocky Mountain Health Plans</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">S</div>
    <div data-row-key="1577-SAG A" data-test="insurance-picker-row" data-uem-id="1577-SAG A">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SAG AFTRA Health Plan</span>
        </div>
    </div>
    <div data-row-key="1525-SAMBA" data-test="insurance-picker-row" data-uem-id="1525-SAMBA">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SAMBA</span>
        </div>
    </div>
    <div data-row-key="593-SCAN " data-test="insurance-picker-row" data-uem-id="593-SCAN ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SCAN Health Plan</span>
        </div>
    </div>
    <div data-row-key="1147-SIHO " data-test="insurance-picker-row" data-uem-id="1147-SIHO ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SIHO Insurance Services</span>
        </div>
    </div>
    <div data-row-key="1065-SSM H" data-test="insurance-picker-row" data-uem-id="1065-SSM H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SSM Health Care</span>
        </div>
    </div>
    <div data-row-key="468-Sagam" data-test="insurance-picker-row" data-uem-id="468-Sagam">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sagamore Health Network</span>
        </div>
    </div>
    <div data-row-key="918-Samar" data-test="insurance-picker-row" data-uem-id="918-Samar">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Samaritan Health Plan Operations</span>
        </div>
    </div>
    <div data-row-key="421-San F" data-test="insurance-picker-row" data-uem-id="421-San F">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>San Francisco Health Plan</span>
        </div>
    </div>
    <div data-row-key="1648-Sana" data-test="insurance-picker-row" data-uem-id="1648-Sana">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sana</span>
        </div>
    </div>
    <div data-row-key="1256-Sanfo" data-test="insurance-picker-row" data-uem-id="1256-Sanfo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sanford Health Plan</span>
        </div>
    </div>
    <div data-row-key="952-Santa" data-test="insurance-picker-row" data-uem-id="952-Santa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Santa Clara Family Health Plan</span>
        </div>
    </div>
    <div data-row-key="516-Scott" data-test="insurance-picker-row" data-uem-id="516-Scott">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Scott &amp; White Health Plan</span>
        </div>
    </div>
    <div data-row-key="1161-Secur" data-test="insurance-picker-row" data-uem-id="1161-Secur">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Security Health Plan of Wisconsin, Inc.</span>
        </div>
    </div>
    <div data-row-key="700-Selec" data-test="insurance-picker-row" data-uem-id="700-Selec">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Select Care</span>
        </div>
    </div>
    <div data-row-key="469-Selec" data-test="insurance-picker-row" data-uem-id="469-Selec">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Select Health Network</span>
        </div>
    </div>
    <div data-row-key="1610-Selec" data-test="insurance-picker-row" data-uem-id="1610-Selec">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Select Health of South Carolina</span>
        </div>
    </div>
    <div data-row-key="332-Selec" data-test="insurance-picker-row" data-uem-id="332-Selec">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SelectHealth</span>
        </div>
    </div>
    <div data-row-key="800-Sende" data-test="insurance-picker-row" data-uem-id="800-Sende">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sendero Health Plans</span>
        </div>
    </div>
    <div data-row-key="1508-Senec" data-test="insurance-picker-row" data-uem-id="1508-Senec">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Seneca Insurance</span>
        </div>
    </div>
    <div data-row-key="1470-Senio" data-test="insurance-picker-row" data-uem-id="1470-Senio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Senior Dimensions</span>
        </div>
    </div>
    <div data-row-key="629-Senio" data-test="insurance-picker-row" data-uem-id="629-Senio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Senior Whole Health</span>
        </div>
    </div>
    <div data-row-key="1507-Sentr" data-test="insurance-picker-row" data-uem-id="1507-Sentr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sentry Insurance</span>
        </div>
    </div>
    <div data-row-key="765-Seton" data-test="insurance-picker-row" data-uem-id="765-Seton">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Seton Health Plan</span>
        </div>
    </div>
    <div data-row-key="953-Sharp" data-test="insurance-picker-row" data-uem-id="953-Sharp">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sharp Health Plan</span>
        </div>
    </div>
    <div data-row-key="755-Sierr" data-test="insurance-picker-row" data-uem-id="755-Sierr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sierra Health and Life</span>
        </div>
    </div>
    <div data-row-key="801-Sight" data-test="insurance-picker-row" data-uem-id="801-Sight">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SightCare</span>
        </div>
    </div>
    <div data-row-key="1617-Silve" data-test="insurance-picker-row" data-uem-id="1617-Silve">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SilverSummit Healthplan</span>
        </div>
    </div>
    <div data-row-key="803-Simpl" data-test="insurance-picker-row" data-uem-id="803-Simpl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Simply Healthcare</span>
        </div>
    </div>
    <div data-row-key="1527-Simpr" data-test="insurance-picker-row" data-uem-id="1527-Simpr">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Simpra Advantage</span>
        </div>
    </div>
    <div data-row-key="1602-Solis" data-test="insurance-picker-row" data-uem-id="1602-Solis">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Solis Health Plans</span>
        </div>
    </div>
    <div data-row-key="971-Solst" data-test="insurance-picker-row" data-uem-id="971-Solst">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Solstice</span>
        </div>
    </div>
    <div data-row-key="1103-Soone" data-test="insurance-picker-row" data-uem-id="1103-Soone">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SoonerCare (Oklahoma Medicaid)</span>
        </div>
    </div>
    <div data-row-key="1024-Sound" data-test="insurance-picker-row" data-uem-id="1024-Sound">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Soundpath Health</span>
        </div>
    </div>
    <div data-row-key="1252-South" data-test="insurance-picker-row" data-uem-id="1252-South">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>South Country Health Alliance</span>
        </div>
    </div>
    <div data-row-key="767-South" data-test="insurance-picker-row" data-uem-id="767-South">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>South Florida Community Care Network</span>
        </div>
    </div>
    <div data-row-key="1575-South" data-test="insurance-picker-row" data-uem-id="1575-South">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Southwestern Health Resources (SWHR)</span>
        </div>
    </div>
    <div data-row-key="650-Spect" data-test="insurance-picker-row" data-uem-id="650-Spect">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Spectera</span>
        </div>
    </div>
    <div data-row-key="1022-Stand" data-test="insurance-picker-row" data-uem-id="1022-Stand">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Standard Life and Accident Insurance Company</span>
        </div>
    </div>
    <div data-row-key="1515-Stanf" data-test="insurance-picker-row" data-uem-id="1515-Stanf">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Stanford Health Care Advantage</span>
        </div>
    </div>
    <div data-row-key="1506-StarN" data-test="insurance-picker-row" data-uem-id="1506-StarN">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>StarNet Insurance</span>
        </div>
    </div>
    <div data-row-key="297-State" data-test="insurance-picker-row" data-uem-id="297-State">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>State Farm®</span>
        </div>
    </div>
    <div data-row-key="1505-State" data-test="insurance-picker-row" data-uem-id="1505-State">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>State National Insurance</span>
        </div>
    </div>
    <div data-row-key="1536-Stayw" data-test="insurance-picker-row" data-uem-id="1536-Stayw">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Staywell Insurance</span>
        </div>
    </div>
    <div data-row-key="1136-Stewa" data-test="insurance-picker-row" data-uem-id="1136-Stewa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Steward Health Care Network - Health Choice Arizona</span>
        </div>
    </div>
    <div data-row-key="1210-Stewa" data-test="insurance-picker-row" data-uem-id="1210-Stewa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Steward Health Choice</span>
        </div>
    </div>
    <div data-row-key="846-Strat" data-test="insurance-picker-row" data-uem-id="846-Strat">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Stratose</span>
        </div>
    </div>
    <div data-row-key="334-Summa" data-test="insurance-picker-row" data-uem-id="334-Summa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SummaCare</span>
        </div>
    </div>
    <div data-row-key="1613-Summi" data-test="insurance-picker-row" data-uem-id="1613-Summi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Summit Community Care</span>
        </div>
    </div>
    <div data-row-key="1249-Sunfl" data-test="insurance-picker-row" data-uem-id="1249-Sunfl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sunflower Health Plan</span>
        </div>
    </div>
    <div data-row-key="1552-Sunri" data-test="insurance-picker-row" data-uem-id="1552-Sunri">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sunrise Advantage Plan</span>
        </div>
    </div>
    <div data-row-key="769-Sunsh" data-test="insurance-picker-row" data-uem-id="769-Sunsh">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sunshine Health</span>
        </div>
    </div>
    <div data-row-key="517-Super" data-test="insurance-picker-row" data-uem-id="517-Super">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Superior HealthPlan</span>
        </div>
    </div>
    <div data-row-key="553-Super" data-test="insurance-picker-row" data-uem-id="553-Super">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Superior Vision</span>
        </div>
    </div>
    <div data-row-key="1504-Susse" data-test="insurance-picker-row" data-uem-id="1504-Susse">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sussex Insurance</span>
        </div>
    </div>
    <div data-row-key="1542-Sutte" data-test="insurance-picker-row" data-uem-id="1542-Sutte">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Sutter Health Plus</span>
        </div>
    </div>
    <div data-row-key="554-Sutte" data-test="insurance-picker-row" data-uem-id="554-Sutte">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>SutterSelect</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">T</div>
    <div data-row-key="1503-TNUS " data-test="insurance-picker-row" data-uem-id="1503-TNUS ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>TNUS Insurance</span>
        </div>
    </div>
    <div data-row-key="1440-TakeC" data-test="insurance-picker-row" data-uem-id="1440-TakeC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>TakeCare</span>
        </div>
    </div>
    <div data-row-key="1100-Teach" data-test="insurance-picker-row" data-uem-id="1100-Teach">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Teachers Health Trust</span>
        </div>
    </div>
    <div data-row-key="567-Texan" data-test="insurance-picker-row" data-uem-id="567-Texan">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>TexanPlus</span>
        </div>
    </div>
    <div data-row-key="518-Texas" data-test="insurance-picker-row" data-uem-id="518-Texas">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Texas Children's Health Plan</span>
        </div>
    </div>
    <div data-row-key="1462-Texas" data-test="insurance-picker-row" data-uem-id="1462-Texas">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Texas Health Aetna</span>
        </div>
    </div>
    <div data-row-key="1646-Texas" data-test="insurance-picker-row" data-uem-id="1646-Texas">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Texas Independence Health Plan (TIHP)</span>
        </div>
    </div>
    <div data-row-key="751-Texas" data-test="insurance-picker-row" data-uem-id="751-Texas">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Texas Kids First</span>
        </div>
    </div>
    <div data-row-key="1270-The H" data-test="insurance-picker-row" data-uem-id="1270-The H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>The HSC Health Care System</span>
        </div>
    </div>
    <div data-row-key="483-The H" data-test="insurance-picker-row" data-uem-id="483-The H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>The Hartford</span>
        </div>
    </div>
    <div data-row-key="1183-The H" data-test="insurance-picker-row" data-uem-id="1183-The H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>The Health Plan of the Upper Ohio Valley, Inc.</span>
        </div>
    </div>
    <div data-row-key="476-Three" data-test="insurance-picker-row" data-uem-id="476-Three">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Three Rivers Providers Network (TRPN)</span>
        </div>
    </div>
    <div data-row-key="842-Total" data-test="insurance-picker-row" data-uem-id="842-Total">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Total Health Care</span>
        </div>
    </div>
    <div data-row-key="535-Touch" data-test="insurance-picker-row" data-uem-id="535-Touch">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Touchstone</span>
        </div>
    </div>
    <div data-row-key="1585-Trans" data-test="insurance-picker-row" data-uem-id="1585-Trans">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Transamerica</span>
        </div>
    </div>
    <div data-row-key="487-Trave" data-test="insurance-picker-row" data-uem-id="487-Trave">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Travelers</span>
        </div>
    </div>
    <div data-row-key="377-Trica" data-test="insurance-picker-row" data-uem-id="377-Trica">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Tricare</span>
        </div>
    </div>
    <div data-row-key="919-Trill" data-test="insurance-picker-row" data-uem-id="919-Trill">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Trillium Community Health Plan</span>
        </div>
    </div>
    <div data-row-key="1293-Trilo" data-test="insurance-picker-row" data-uem-id="1293-Trilo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Trilogy Health Insurance</span>
        </div>
    </div>
    <div data-row-key="1080-Tripl" data-test="insurance-picker-row" data-uem-id="1080-Tripl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Triple-S Salud: Blue Cross Blue Shield of Puerto Rico</span>
        </div>
    </div>
    <div data-row-key="1570-Triwe" data-test="insurance-picker-row" data-uem-id="1570-Triwe">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Triwest Healthcare Alliance</span>
        </div>
    </div>
    <div data-row-key="1526-TrueH" data-test="insurance-picker-row" data-uem-id="1526-TrueH">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>TrueHealth New Mexico</span>
        </div>
    </div>
    <div data-row-key="1177-Trust" data-test="insurance-picker-row" data-uem-id="1177-Trust">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Trusted Health Plan</span>
        </div>
    </div>
    <div data-row-key="922-Tuali" data-test="insurance-picker-row" data-uem-id="922-Tuali">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Tuality Health Alliance</span>
        </div>
    </div>
    <div data-row-key="1344-Tufts" data-test="insurance-picker-row" data-uem-id="1344-Tufts">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Tufts Health Freedom Plan</span>
        </div>
    </div>
    <div data-row-key="616-Tufts" data-test="insurance-picker-row" data-uem-id="616-Tufts">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Tufts Health Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">U</div>
    <div data-row-key="467-UCHP " data-test="insurance-picker-row" data-uem-id="467-UCHP ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UCHP (University of Chicago Health Plan)</span>
        </div>
    </div>
    <div data-row-key="701-UCare" data-test="insurance-picker-row" data-uem-id="701-UCare">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UCare</span>
        </div>
    </div>
    <div data-row-key="1178-UHA H" data-test="insurance-picker-row" data-uem-id="1178-UHA H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UHA Health Insurance</span>
        </div>
    </div>
    <div data-row-key="1502-ULLIC" data-test="insurance-picker-row" data-uem-id="1502-ULLIC">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>ULLICO Casualty Company</span>
        </div>
    </div>
    <div data-row-key="621-UPMC " data-test="insurance-picker-row" data-uem-id="621-UPMC ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UPMC Health Plan</span>
        </div>
    </div>
    <div data-row-key="414-US Fa" data-test="insurance-picker-row" data-uem-id="414-US Fa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>US Family Health Plan</span>
        </div>
    </div>
    <div data-row-key="974-US He" data-test="insurance-picker-row" data-uem-id="974-US He">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>US Health Group</span>
        </div>
    </div>
    <div data-row-key="384-USA M" data-test="insurance-picker-row" data-uem-id="384-USA M">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>USA Managed Care Organization</span>
        </div>
    </div>
    <div data-row-key="1331-USAbl" data-test="insurance-picker-row" data-uem-id="1331-USAbl">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>USAble Mutual Insurance Company</span>
        </div>
    </div>
    <div data-row-key="1457-Ultim" data-test="insurance-picker-row" data-uem-id="1457-Ultim">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ultimate Health Plans</span>
        </div>
    </div>
    <div data-row-key="1593-Umpqu" data-test="insurance-picker-row" data-uem-id="1593-Umpqu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Umpqua Health Alliance</span>
        </div>
    </div>
    <div data-row-key="322-UniCa" data-test="insurance-picker-row" data-uem-id="322-UniCa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UniCare</span>
        </div>
    </div>
    <div data-row-key="623-Unifo" data-test="insurance-picker-row" data-uem-id="623-Unifo">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Uniform Medical Plan</span>
        </div>
    </div>
    <div data-row-key="1465-Union" data-test="insurance-picker-row" data-uem-id="1465-Union">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Union Eye Care</span>
        </div>
    </div>
    <div data-row-key="475-Union" data-test="insurance-picker-row" data-uem-id="475-Union">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Union Health Services, Inc</span>
        </div>
    </div>
    <div data-row-key="346-Union" data-test="insurance-picker-row" data-uem-id="346-Union">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Union Plans</span>
        </div>
    </div>
    <div data-row-key="987-Unite" data-test="insurance-picker-row" data-uem-id="987-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>United American</span>
        </div>
    </div>
    <div data-row-key="449-Unite" data-test="insurance-picker-row" data-uem-id="449-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>United Behavioral Health</span>
        </div>
    </div>
    <div data-row-key="995-Unite" data-test="insurance-picker-row" data-uem-id="995-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UnitedHealthOne</span>
        </div>
    </div>
    <div data-row-key="323-Unite" data-test="insurance-picker-row" data-uem-id="323-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UnitedHealthcare</span>
        </div>
    </div>
    <div data-row-key="644-Unite" data-test="insurance-picker-row" data-uem-id="644-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UnitedHealthcare Community Plan</span>
        </div>
    </div>
    <div data-row-key="330-Unite" data-test="insurance-picker-row" data-uem-id="330-Unite">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>UnitedHealthcare Oxford</span>
        </div>
    </div>
    <div data-row-key="1170-Unity" data-test="insurance-picker-row" data-uem-id="1170-Unity">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Unity Health Insurance</span>
        </div>
    </div>
    <div data-row-key="542-Unive" data-test="insurance-picker-row" data-uem-id="542-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Univera Healthcare</span>
        </div>
    </div>
    <div data-row-key="531-Unive" data-test="insurance-picker-row" data-uem-id="531-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Universal American</span>
        </div>
    </div>
    <div data-row-key="1501-Unive" data-test="insurance-picker-row" data-uem-id="1501-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Universal Underwriters Insurance</span>
        </div>
    </div>
    <div data-row-key="1287-Unive" data-test="insurance-picker-row" data-uem-id="1287-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University Hospitals (Health Design Plus)</span>
        </div>
    </div>
    <div data-row-key="1371-Unive" data-test="insurance-picker-row" data-uem-id="1371-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University Physician Network (UPN)</span>
        </div>
    </div>
    <div data-row-key="1082-Unive" data-test="insurance-picker-row" data-uem-id="1082-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University of Arizona Health Plans</span>
        </div>
    </div>
    <div data-row-key="1448-Unive" data-test="insurance-picker-row" data-uem-id="1448-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University of Maryland Health Advantage</span>
        </div>
    </div>
    <div data-row-key="1564-Unive" data-test="insurance-picker-row" data-uem-id="1564-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University of Maryland Health Partners</span>
        </div>
    </div>
    <div data-row-key="1647-Unive" data-test="insurance-picker-row" data-uem-id="1647-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University of St. Mary of the Lake - Mundelein Seminary</span>
        </div>
    </div>
    <div data-row-key="1332-Unive" data-test="insurance-picker-row" data-uem-id="1332-Unive">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>University of Utah Health Plans</span>
        </div>
    </div>
    <div data-row-key="1631-Upper" data-test="insurance-picker-row" data-uem-id="1631-Upper">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Upper Peninsula Health Plan</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">V</div>
    <div data-row-key="416-VNS C" data-test="insurance-picker-row" data-uem-id="416-VNS C">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>VNS Choice Health Plans</span>
        </div>
    </div>
    <div data-row-key="441-VSP" data-test="insurance-picker-row" data-uem-id="441-VSP">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>VSP</span>
        </div>
    </div>
    <div data-row-key="1500-Valle" data-test="insurance-picker-row" data-uem-id="1500-Valle">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Valley Forge Insurance</span>
        </div>
    </div>
    <div data-row-key="861-Valle" data-test="insurance-picker-row" data-uem-id="861-Valle">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Valley Health Plan</span>
        </div>
    </div>
    <div data-row-key="1626-Valor" data-test="insurance-picker-row" data-uem-id="1626-Valor">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Valor Health Plan</span>
        </div>
    </div>
    <div data-row-key="1166-Vanta" data-test="insurance-picker-row" data-uem-id="1166-Vanta">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vantage Health Plan, Inc.</span>
        </div>
    </div>
    <div data-row-key="958-Ventu" data-test="insurance-picker-row" data-uem-id="958-Ventu">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Ventura County Health Care Plan</span>
        </div>
    </div>
    <div data-row-key="1442-Vibra" data-test="insurance-picker-row" data-uem-id="1442-Vibra">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vibra Health Plan</span>
        </div>
    </div>
    <div data-row-key="1098-Villa" data-test="insurance-picker-row" data-uem-id="1098-Villa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>VillageCareMax</span>
        </div>
    </div>
    <div data-row-key="1416-Virgi" data-test="insurance-picker-row" data-uem-id="1416-Virgi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Virginia Coordinated Care (VCC)</span>
        </div>
    </div>
    <div data-row-key="479-Virgi" data-test="insurance-picker-row" data-uem-id="479-Virgi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Virginia Health Network</span>
        </div>
    </div>
    <div data-row-key="426-Virgi" data-test="insurance-picker-row" data-uem-id="426-Virgi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Virginia Premier Health Plan</span>
        </div>
    </div>
    <div data-row-key="548-Visio" data-test="insurance-picker-row" data-uem-id="548-Visio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vision Benefits of America</span>
        </div>
    </div>
    <div data-row-key="778-Visio" data-test="insurance-picker-row" data-uem-id="778-Visio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vision Care Direct</span>
        </div>
    </div>
    <div data-row-key="839-Visio" data-test="insurance-picker-row" data-uem-id="839-Visio">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vision Plan of America</span>
        </div>
    </div>
    <div data-row-key="1113-Viva " data-test="insurance-picker-row" data-uem-id="1113-Viva ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Viva Health Plan</span>
        </div>
    </div>
    <div data-row-key="1565-Vivid" data-test="insurance-picker-row" data-uem-id="1565-Vivid">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vivida Health</span>
        </div>
    </div>
    <div data-row-key="857-Volus" data-test="insurance-picker-row" data-uem-id="857-Volus">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Volusia Health Network</span>
        </div>
    </div>
    <div data-row-key="417-Vytra" data-test="insurance-picker-row" data-uem-id="417-Vytra">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Vytra</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">W</div>
    <div data-row-key="1294-WEA T" data-test="insurance-picker-row" data-uem-id="1294-WEA T">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>WEA Trust</span>
        </div>
    </div>
    <div data-row-key="1333-WPS H" data-test="insurance-picker-row" data-uem-id="1333-WPS H">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>WPS Health Plan</span>
        </div>
    </div>
    <div data-row-key="1499-WRM A" data-test="insurance-picker-row" data-uem-id="1499-WRM A">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>WRM America Indemnity Company</span>
        </div>
    </div>
    <div data-row-key="1288-Well " data-test="insurance-picker-row" data-uem-id="1288-Well ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Well Sense Health Plan</span>
        </div>
    </div>
    <div data-row-key="1638-WellF" data-test="insurance-picker-row" data-uem-id="1638-WellF">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>WellFirst Health</span>
        </div>
    </div>
    <div data-row-key="418-Wellc" data-test="insurance-picker-row" data-uem-id="418-Wellc">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Wellcare</span>
        </div>
    </div>
    <div data-row-key="856-Wellm" data-test="insurance-picker-row" data-uem-id="856-Wellm">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Wellmark Blue Cross Blue Shield</span>
        </div>
    </div>
    <div data-row-key="1498-West " data-test="insurance-picker-row" data-uem-id="1498-West ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>West American Insurance</span>
        </div>
    </div>
    <div data-row-key="1458-West " data-test="insurance-picker-row" data-uem-id="1458-West ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>West Virginia Senior Advantage</span>
        </div>
    </div>
    <div data-row-key="422-Weste" data-test="insurance-picker-row" data-uem-id="422-Weste">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Western Health Advantage</span>
        </div>
    </div>
    <div data-row-key="1630-Weste" data-test="insurance-picker-row" data-uem-id="1630-Weste">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Western Sky Community Care</span>
        </div>
    </div>
    <div data-row-key="1595-Willa" data-test="insurance-picker-row" data-uem-id="1595-Willa">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Willamette Valley Community Health</span>
        </div>
    </div>
    <div data-row-key="501-Worke" data-test="insurance-picker-row" data-uem-id="501-Worke">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Workers' Compensation</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">Y</div>
    <div data-row-key="1594-Yamhi" data-test="insurance-picker-row" data-uem-id="1594-Yamhi">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Yamhill Community Care Organization</span>
        </div>
    </div>
    <div class="patient-web-app__sc-e9ndvy-0 iLLdUB">Z</div>
    <div data-row-key="486-Zenit" data-test="insurance-picker-row" data-uem-id="486-Zenit">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Zenith</span>
        </div>
    </div>
    <div data-row-key="1651-Zing " data-test="insurance-picker-row" data-uem-id="1651-Zing ">
        <div class="patient-web-app__sc-53koon-5 dxKQOT">
            <span>Zing Health</span>
        </div>
    </div>
</div>
`
}
