package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// struct to hold the VTT file data, name, url
type VTTFile struct {
	Name string
	URL  string
}

func main() {
	// array of VTT files EP1 to EP39
	vttFiles := []VTTFile{
		{
			Name: "EP1",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01C8F946BE1C6D7B382BC5F0FF4253D8D83737B7458731A8B96B7EB4D0BBE40826379948DC74C833FDECC53D689996AABAEF1FE33F87B49AAEB0BA85997F80B8103FB342999200269E47241B2633F7E823E53094BDC11EEB10BBBFFDB36564FA913F134B246AD029732FF5D6B5128B0C7267CA37A78052D0898C039CC73F9DDC46/gzc_1000117_0b534uafgaaaquamb4pxcjs43zodkpuqavsa.f451507.vtt?ver=4",
		},
		{
			Name: "EP2",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01518A6E98CEAD1E9CA404989A33DF943665728C64857C960C1378440F1AAFA3EB93359B1108B309C39748BE59E4B7B74AEB1B7517CE9752EFF856BD7964B4C649A7A764C1681C798E1ED58ACC9B866226594C5C54FC406CCB393E1C0C1F5B00EABD4D17386B8375B8A6C751461B0EC29FE1D2A7561581C808D4EBB9D50B21B557/gzc_1000117_0b53fqafgaaajqamaopxcjs42lgdkmqaavsa.f865507.vtt?ver=4",
		},
		{
			Name: "EP3",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/015DCA3263F0C5B804F129411C04BD371945592383CF170B36AAFC644BF24872B7A8B127132CC61811FFE9C9B62428F83E89356486904172633F5224FF1478C10D67363F9D448CC7EA411DE7731663398E63C5FEE4FA16EF642E8247A4FCC429B961E2423286A6F590713650DF88332E03AFDD82BE97DE0FBD9AC87ECEFC501E42/gzc_1000117_0b53raafiaaa5aampg7xcvs43cgdkscaauka.f300507.vtt?ver=4",
		},
		{
			Name: "EP4",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0111408F97A97A996FB05C0783FA48441B9CEE1D9C4379E56169A2D97AC39B1DD34D8E1D71BF0A5D5F5E367230E6B61CC95E232C789DD3327503187162991B3F679F56F48A6A852615DBD6B0A2ECEDDC83E5C29B9C3728A44623759702C97DE8423A0A344FE20F9F571C9F6F4B2AE923473FF8DDC0ECDB3756F4313FB6573DF342/gzc_1000117_0b53diaakaaarqanvzhxzbs42gwdaulaaaca.f246507.vtt?ver=4",
		},
		{
			Name: "EP5",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/018D000A693043BAC63CBCDBE2A0FE5C441FCF1B3017D8D1C540BCDFFDA68608C67274BB3F6DA841DD1CD331A88623F96AF77275D3B72A33F353AEF42D81332A518AE667CB0AC1DD6FF6B3E3E8A8C3223839EF56F802DAF4016E8FE3693B01A0A8003104C161219D3121BC4596FEA474C6A7C924C4A1DE935D9C25B5F04D438889/gzc_1000117_0b53yaafgaaaa4alknpyczs43qgdkpgaavsa.f235507.vtt?ver=4",
		},
		{
			Name: "EP6",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01AD9C6385D678B5117E23C2135C117D354AE9BC68BA9094063DA85483C4EB5DA47BC6E3887BD72C0F8788997BD68235A9B5194EFA0677F7DC788F0F91E93730823530D0888673D2F851781A6BED216DBDA547587678B6302019C431033DDD3935FC5E60EB7380E6810C2234A49A92E3BEE8314F571EAF90781ED1711874A37B8F/gzc_1000117_0b53yeafgaaaayalknhyczs43qodkpgqavsa.f31507.vtt?ver=4",
		},
		{
			Name: "EP7",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/019DF232A059B63ABC591EE73BFA2D8D234DAC6FE4A2DE75FE66DB66F4667861B476940F2A9D4246D7734D51B9EC1817BD25D20CD694E9B2644AABAFFA56F26D7F48BC1427D29ECAB548476B880568E81F6832A3F5017DFCE7954D818504915A149A7EC38ACED80358C4FF216D1E17FD4C4B96D2C5FD762F985D4E433A3121F2EE/gzc_1000117_0b53feakqaaagaajwoh2vvs42kodvasqblka.f884507.vtt?ver=4",
		},
		{
			Name: "EP8",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0175907D290133880ABD2CAFEA114BF9937DC0DB2F0A6C28CB3BADE6A77D61B347330DC1F8AFBE00441E142DD7FAF00C9AE4FA17250805A9E44C7EFBE7E0B69E8130A689BB560A5AED27AF4633852DD9D298E8FBACE98BCC5034DF7F28CB233138A383A20D785140990C89CD40D2011F2A2452D103EAEF8D4B2BC6542796207A34/gzc_1000117_0b53emab2aaahiac4qp2sns42i6dduxqagca.f734507.vtt?ver=4",
		},
		{
			Name: "EP9",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0174D6E0E8BAB0D66EC15B1F23BC10BE74365F4E2CE7BDABC967DEF962B518FEBE9B1FFF57AC5BBABBF82D3A2FB28BAB9F2392D48A5B3A0804D8F5331DADDFB4C9528607783076A51EE74808CF277F950C8B7005589A04C681B24D76A053B77D480528CA0877F6F3A1A7B82B815419496D223CD85C25E19D5809B5C299701F30A6/gzc_1000117_0b532aanqaaa4eaflyx3pvs43ugd3doabxka.f961507.vtt?ver=4",
		},
		{
			Name: "EP10",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01778E88B85DE01BBE39C99D27AEB2EE6588F11778102493F1BDEED56D45F42E3481419898D5A7C4C800FCCA338C86912E93D9B54C1835AD3BB8013B0AA2C17D617C9BAD9EE050609C3941AD61328CBB7ECE81D767CF474D849E3014CDC38885793C3CC10D47A9E378719E91D82F7DE4BDA41A53293C8301EF964C7127CA44AF55/gzc_1000117_0b53xeadcaaaraalznp3xvs43oodgg2qanca.f319507.vtt?ver=4",
		},
		{
			Name: "EP11",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01190E1B4BFD223449FC276BF96724420F4D3106A770D38AF3516A6DEBD66675358D70147E58371DEB4BFB7AEF34272DDE6E607FD8C1DFFE480DAFDF4D137BE688E2B115E01464B1ACD6B7452B5102E6B59C7EC373E9C1360DFC16D4EA09AD9223D8927862D50BDB946C0AC4528378E9ACFD3629B26A650BD2E881F0788EBA7CAA/gzc_1000117_0b53iiaqwaabnaanpa75frs4yqwdbnhacdsa.f446507.vtt?ver=4",
		},
		{
			Name: "EP12",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0183BFCF3C25B427D48D55817CB3726E23C533090F666A5A64D2AE9E8B73F830D28B3D8A8CA8513EF19C34FB09DAE40BCFFB11F85C2F25A67B3A093A555D984E75F14DB1E450FFD22BD1CEEFBB65BD73E780F082BE7127D2D7A05F593C729C2A5415205192D10BF67708C2D976AF8F4148496F6932D77C2CFDF592BF27B4343AD8/gzc_1000117_0b53raaleaaaoyam7gp6w5s43cgdwkcabn2a.f248507.vtt?ver=4",
		},
		{
			Name: "EP13",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B7A31A908A8BAC0164510CFECB5BAAFAE91F67C0FF91414814FE4CF0186655BB4DD835C1EF6C6EC543F2D8E2FD3A41D21E058347E7476F6C48A356E39A70E442D74E9ED8971EE2E6C1B5040B11EC59ED008E7C6CFB9726A55CDCCA3098889F622307DB022AF453DBB5675C331336897303467DE93EDD690283AED40F2B66BC27/gzc_1000117_0b53iiacoaaabuafsh77vfs42qwde5haaisa.f830507.vtt?ver=4",
		},
		{
			Name: "EP14",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B1348C2EFE4C8791EC2D0A42816FC9394219B777908556211B99EACAFDDCB4F2884C2B4558B6E6160964DE14D8EDF538F41E0C02DA3DA0F8A02F368A7E267B4E8D774F42549AFCD642916584AD3B3641B628997ED6C393E4733807D24782FDD600429303CC031FB0D3981777184B3B73073FAC59C3ED7E2EBCA2BBB9929C2D0D/gzc_1000117_0b53qmag2aaa5aakfeyafrtm3a6dnwhqa2ca.f481507.vtt?ver=4",
		},
		{
			Name: "EP15",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0136AE1A32B67B6A239E0AD587AA1A26944952717002EA8E0C7C1702DE94AA6CCCB526ED5E2863991B727BF66B0500A1585621FF53BC330B547A636644EEBD75AA2B40DD227B09EBC8F5E88452F53BB6B88D1C4BDD65108343E614BAA2378352BA761BA1C2BDC4947E11BAE263C8331020EBD95F34CDB75BF5B08D6B9A73EBE7ED/gzc_1000117_0b53saaqqaab7mamotibiftmzegdbcoacdka.f297507.vtt?ver=4",
		},
		{
			Name: "EP16",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0160B375C5685D7CA43A31B98775C3CFBE3B01F03D4CE78C6694780050507C12EBA25BC4A68EE3C2C769DCBE89C6E33174F4278E9A71D0D7132ACA4D91C0B22B552D9C655736B8D8FFC1242A9599F21A52C80BCB74FF4072BCCDF075106CE11C85546A8EF387EBD4A04734737B9DF3FCFFE93AB4EA5BCA40B74E99DDC7855BA338/gzc_1000117_0b53hmaasaaai4almcqhiztm2o6dbe3qadca.f626507.vtt?ver=4",
		},
		{
			Name: "EP17",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/018E1E6463260A36E38F6C221E5655ED81910EF15820361AA2A1A2100B462EEA2AE7F8D6E3DC7C369970343A2509D17659D772F967A5D17449AC6DD07D239D8B4FBEA2D8491EFA74CFB24E38E845F8DA08D18AB4475C7ABD58B4BBBCFB5F1517F19EECC048FB64BE0CAA1C89631B8635D3E010D32B4D15D9EF86A2332D94BC3350/gzc_1000117_0b53cyasqaabdaaasiydhftmyfwdfanaclka.f194507.vtt?ver=4",
		},
		{
			Name: "EP18",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01F0E7EE7E141960FA6BD498842778C0B0B3E5F4D8F1B5D8A09FA0D796C5A8778300DDBAFC9D4F93CC90D30EB41C4CFDF74C277DCD03A4F5AF3BC0799EC30AD8EE48B5F109792F1FD09644509748C42E91C27202BE27A592504CFCF5FA98A5F3515794A6C25EDA99F876F33C73C9597353C5A06139EA9FE61ED1A53AB6DED3A34C/gzc_1000117_0b53t4afsaaanyacxyydfvtm3h6dlgjqaxca.f579507.vtt?ver=4",
		},
		{
			Name: "EP19",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0179244475E9FD01D45EA3BA1DA7BDBC749EE1E68AF278AACB8764671C6A40B1434EFCFD36430408D5511830D5AE40CE962BC361DF83AE93F2639AD82584A14EA2E5F7B6B7381A1F53F3FC53EFAA2A9A906DE74574604D5589CD8ABC06E818DD4F6599A752CFA0B8D483149DA0D4251EDD274E9D496260899775B7B45636F426DB/gzc_1000117_0b53fmafiaaa3iacmuid25tm2k6dkqtqauka.f468507.vtt?ver=4",
		},
		{
			Name: "EP20",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01571E09732BF8312551D3F480BE17595898A0BFBA314497566254476A00AC362A360E51216F0DF14FD1A1B5C3F2A624ABA3C23C5D6712A7BF7CEE8ACA8FF480347CCFDC26956FE82D7FDFFCEE7862C07B4882A074ED0ED20DDFC0872EB4C08BFF9B956E33B5DB5CDC7208112911949EDA72A423DCF6153D7227C1E907D9497843/gzc_1000117_0b53xyamqaaaweaa23aezjtm3pwdzczabtka.f259507.vtt?ver=4",
		},
		{
			Name: "EP21",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B25B110FD7B36398FE896E1F9C8B08A8BE89154C5C35ABC416C4D3099B26A806748702A64A42ECE5281737607EA7F5E8148FA999B297D4F82D37A48218CC8CB2AC9826D0027DE04AA5BDEBFC42D20FA3780E8161F133B21B41C45DD732AB74AA945F912701D42D539DB6FC3EAA4197BE4130F4135B8061AB84612A41955F2C2B/gzc_1000117_0b53ouaagaaapiamnliejntm25odan4qabsa.f511507.vtt?ver=4",
		},
		{
			Name: "EP22",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/011FE845A74BB405FB1AB58CB7342245B20BFED3A3A143FC1C04BFF403A5F62FBB0E4E2E4E6C4D768FDF9EC691AA08B1ABC210A4F58CEDAE3DE04662013D73F750C1A60E9043B960940DF708089BB7FF451E1A4E31609FFA003B1EED638D9BE5DE92791B7040AA3570B98C57004220161A7935E9FA95C941C57411FD10B598D4E9/gzc_1000117_0b53fyacgaaawiadttag2ntm2lwdemraajsa.f516507.vtt?ver=4",
		},
		{
			Name: "EP23",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B8D23ADAFCF400897A44F4060C16DC63ECA399EBF9996DA74F51EBFAB7E8A63B0784057522A1EBC7A7E62A4DAA38472852E9171B974EBE319809693377A65A9FF784CE3D28F2D9D30DAE66BBED1AB5F63636FBC7171B0D39C1D331EB5CCF108350F27ADEE2AF4D4BB1EBA16A0A92EE63C71421921D8EAB3D5B30A5DBDD6ECB0A/gzc_1000117_0b53iqaakaaa5uag4qihqvtm2rgdaveaaaca.f884507.vtt?ver=4",
		},
		{
			Name: "EP24",
			URL:  "https://subtitle.wetvinfo.com/svp_50125/hHfTiwjl8EXOHsSTBpCSqc4GXScJttjo_nkYJpGqolCjY--2YaofxtCdOPd42mYE_aEyqUhUQBGew6YNpVaGHa_eNi0geUG0oYf865SYNLL5htLpxF-nEfzCeEZuEFQPCvnKehUQPeE0hhvfe-XafRV0zsLz9ChH/gzc_1000117_0b53uaaaqaaagaalm5yipftm3igdbcwaadka.f815507.vtt?ver=4",
		},
		{
			Name: "EP25",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B10FBAB1E1239438934246FDEA4FFC4F8AF7B28D7FA73E463D2F1A3DCD343811ACDEF56D1E9EF1E6380E925E677614D399DD6A830FAAA1DA2FFD8094B11BED67926C3CF8BD2D6F5F6AB77D8B30A3C3579EC5C1A9B8456164B8A0B796A904C19CAFB4ADF7D06DDAB08DAC56E17CFE34452F02EAFB1FFD6A3691A8D892C305DC1B/gzc_1000117_0b53zmaauaaapaaansykrjtm3s6dbldqad2a.f143507.vtt?ver=4",
		},
		{
			Name: "EP26",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/019F846C3A9568D0C2828B76FA6CA3278E3EB64EB10259160E879664E7FFFAAE936991DBDB55A4FD278E1B7263829D9825D047B7A01CFAAB4DCE55038D717691F0A77340953AE3130747E5FA26EE38BB0E2BC75AFA7C1125197BDB9495A7523430189BDEFEEA76A05AE0CA32B9EE8C76804BC9365A268AAF5E1F627EBB0238587C/gzc_1000117_0b532aaauaaammaanvakrjtm3ugdbloaad2a.f688507.vtt?ver=4",
		},
		{
			Name: "EP27",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01F49AE78B64C2B028C171D426F539CDC136476D50A8482B6B9776D917B56C547F41208CF7B82690CF06F4CB7E90437B62DC6AA9ABAC59ADD149B05928D3E7B860A00269A58B6B042CEF0E07CA6A159F0745C8F9F51B0721ECD3D45800BFEDE3D71CC4CA3E09CCB3767661914BE0C9B5EECA3087DF3B0014056F78BCB4D54AC04B/gzc_1000117_0b53yqabeaaavmae3pqlmbtm3rgdcleaaf2a.f109507.vtt?ver=4",
		},
		{
			Name: "EP28",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01B3D9E66E1E151664930ED9277E37ABF0B8CC16061C143D7AF648EBDA15CB2248463311700C9672960D6F679794406FB3F7BB9CC5436DDB819B2BB784B8B5794377706AA0065E8830D9E0AE24C5A0421A2D40D965E54AA70DB766741093E94BA6C1CD9BF084401D3B8638476F715768AFB704662BBF357299C5608FE55C564D5B/gzc_1000117_0b532mabcaaaxqae5nalmntm3u6dchpqafca.f941507.vtt?ver=4",
		},
		{
			Name: "EP29",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/0194148D13D93EF4641875FCFDFE42DD860FB085E2868442FC75B7D0F28E19416BE8084CA9A86050DCA3D9F38CFF5AD5347AA1B03DE2851A9DE5C378E75D3F8F63BAF7C4997D3BF41A40F4201B230ECFBABD435C736D1A3154B17A02C2E408CD5976C56F021452543C45DD485E565843988B4F944C924C02C1828851EAC021150A/gzc_1000117_0b53uyabkaaaeyal5bqmljtm3jwdcwvaaeca.f154507.vtt?ver=4",
		},
		{
			Name: "EP30",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01FA1F037DFDE9BB8353D5AB92823217CE0F33B512E58D50397105690A113C9E8D2425710FF66091330079765F0D80B0432F97B45C88EA0E4EDDD85ED284FF4C873FBDD6AB88F9245E20FCB32F0F48AB1BE1F33EE94D3660D01A0DC715B1A098E4E19901097A27071331166FBBBCF8C9866925DF0851664E2D8EBF3474BA1B37E9/gzc_1000117_0b53h4abiaaax4al7mymnrtm2p6dcqzqaeka.f401507.vtt?ver=4",
		},
		{
			Name: "EP31",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/010345D8BCD7DE09A2DCC8E0D5A4D939893BE40D3ACA87A55600E2FECC2D5211043175BC54DE5E60FDA49C5BDFA0572CC89963382744EA05FBF0696D2D9218AADB18D1F383274AB8B72A6D0FC9BED67E65A70D40B2E83F15294DD9C9641C3EEA883CFA3474E2404D733AD4B7D597B10EA003771DC9F71969E61FD59A456EDFCA9B/gzc_1000117_0b53tqabuaaaxyabviaokvtm3hgddkiaah2a.f407507.vtt?ver=4",
		},
		{
			Name: "EP32",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01533183F942C831ADFA9D88E2091D36A83ED00F5FFB7C9CFC6AB92C379F3B45F9E1E194B077A56503E252F6F8F1360BBAC5C2408FCF9745A570BE575C244D2E6BB3B274A1F9C3B890E0D972772A6E8FDFCB616DD95A61466CEC6D920496BFB45898A40700DCE7C4C144FB92111ED8A9AF3EE3A2532234D4D082A54B1C1C920A36/gzc_1000117_0b53yaabwaaavmabejionrtm3qgddpgaahsa.f444507.vtt?ver=4",
		},
		{
			Name: "EP33",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01A80A8850D00957B58CDF79D3552F7B1F1D6F710E029DAA0F97049F90AA59B46D28108594605D9DE5BDBC34602EC7E19B50EC6F35D37D24C0D1EB8C806AF2DE505344A0928DE042D9D5E3258BA0C608BB4506A97CCDF8AA0A645230455FBAA167AFA13088D65CE34DEBE4DCD5F5453B80A714DB3F727F22D8AB64B6DC3A0EB3C3/gzc_1000117_0b53taabwaaakmaereypnjtm3ggddokaahsa.f35507.vtt?ver=4",
		},
		{
			Name: "EP34",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/011D9C5148DD28EA50C5428D7A95012AC1C0DAA72224B1561C03686F96208F295258BEEA13AA13A178D527E60993706D570D8FD717599AB239F3B8F3E8D695BFCE45144261660A0F6EF78B48918CD5611F852C7A3D1F25DCBF528DDCEDB34606190BC08476B357E9560C4B3056136E9F9AF5F4CCCA5D7BDEC9DCB46192D075DC43/gzc_1000117_0b53beabyaaayiae6aypnvtm2coddqcqagka.f539507.vtt?ver=4",
		},
		{
			Name: "EP35",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/015F13B7BBB304CE3FEF6302BFF0C1DDFEF39B4AB777D670F4FA5D49B854A0D9CBBCDC42034C58B1DAEB2D77B94ED4139FDA1B66D582985A5DE76BEB09A08099D6E461A14F94CB39E196EC6401DCA90EF196D5D548E8D497B8CA16E4593D15EC6B017B0C05675D79119090CB8F612242DCBCE3DD21FB8E6E9C648A838D5B8E1457/gzc_1000117_0b536qabuaaah4aetvipmjtm35gddl4aah2a.f526507.vtt?ver=4",
		},
		{
			Name: "EP36",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01F75FC5024A4810BD1B41D9BA82D5D0E30967980C55393F6D7B501ED15A74A103FD54C04E98EBD7669E5C6E0B3D7987713A42AAC017A19AFDE60DDD53AD5EBE9CF93C449D9A5D0D12B111E4D3F3E36B159A942024EF0FBBE3BC906A60D963AACC8FAE5318C17D6ACF2328F42B247F9CD0B5DE48B7537C1869A17FB2CE4CE5BE06/gzc_1000117_0b53teabwaaakiaereqpnjtm3goddokqahsa.f757507.vtt?ver=4",
		},
		{
			Name: "EP37",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/014D8AE62908A80CE94839277F49E7F67B7D153348E10CD2320729E40A1B79B4E7A3EE535C168804F4640AD1C6CB7669DC89859B02326B2D0591E2BA25F89195B85E3E15D0EEA54BE5D6525E154ED54E3B1FF5A49077F54F133431D80EE6E61C48F79A859EC1AF819F220702AA951313831D19DDA29A2F21CE785606D076E93C01/gzc_1000117_0b533eabmaaaciaekdipm5tm3wodc3kqae2a.f385507.vtt?ver=4",
		},
		{
			Name: "EP38",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01120721B29E4241B7EE3D08D628B98B9474B1384E4433555955D451337ED1B7268D6A3A1A64C141AA1FE229898D2834D4CC30B13896C89D19BB9EA57078E935F329BC729DF73CA1BB3133AEE4E4319448EAFAA77664E2B57002A23D3090219DD66C8D51BC6CD7E7AC864B1A85D338E33C3EE4460A9DC7DC3DDC8131F355EA94B4/gzc_1000117_0b53rmabmaaaiaaekxapm5tm3c6dc2dqae2a.f103507.vtt?ver=4",
		},
		{
			Name: "EP39",
			URL:  "https://cffaws.wetvinfo.com/svp_50125/01C2A9C4A14DFE08E8D2C47B5C9F48E0BD660861B2CFD433BE763E318BF60D52A560B6AD9D6028434FC22350EC66289B8C7D1868C56F35050676708EF967253E9A68808FC8D38F5C170E556353627AEAD0DA6C4D37A4B50010E46DE8DE532222EC6E54CE13F4DCCD15695545AF1C97DFAAB718FCB7C64327F2104FDBC5A3DF628E/gzc_1000117_0b53t4abuaaakqaetvypmrtm3h6ddkjqah2a.f378507.vtt?ver=4",
		},
	}

	// loop through the VTT files and download them
	for _, vttFile := range vttFiles {
		log.Printf("Downloading VTT file: %s\n", vttFile.Name)
		err := downloadFile(vttFile)
		if err != nil {
			log.Fatal("Error downloading VTT file:", err)
		}
	}

}

func downloadFile(file VTTFile) error {
	resp, err := http.Get(file.URL)
	if err != nil {
		log.Fatal("Error fetching VTT file:", err)
	}
	defer resp.Body.Close()

	// Create the output file
	outputFile, err := os.Create(file.Name + ".vtt")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()

	// Copy the content from response body to the output file
	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		log.Fatal("Error copying content to output file:", err)
	}

	log.Println("VTT file downloaded successfully!")

	return nil
}
