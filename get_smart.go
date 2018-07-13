package smart

import ("fmt"
)
/*smart info: {
	"pgs":[0,2,3,5,6,8,13,14,15,16,21,24,25,26,47,48,55],
	"we":{
		"cor":5639,
		"bt":1888071680},
	"re":{
		"cor":3041,
		"bt":1131528192},
	"ve":{"cor":2692},
	"nme":{},
	"tmp":{
		"tmp":25,
		"tmpm":85},
	"ssc":{
		"my":2017,
		"mw":37,
		"ay":2017,
		"aw":37,
		"sscm":50000,
		"ssc":12,
		"lucm":600000,
		"luc":95},
	"slf":{},
	"prt":{
		"Ports":
		[{
			"pid":1,
			"cd":2,
			"phs":[
				{
				"adt":2,
				"adr":1,
				"lnk":11,
				"ismp":true,
				"tsmp":true,
				"sadr":5764832520584332277,
				"asadr":5820324243979633023,
				"aphid":19,"nph":4}]},
		{
			"pid":2,
			"cd":2,
			"phs":[{
				"phid":1,
				"adt":2,
				"adr":1,
				"lnk":11,
				"ismp":true,
				"tsmp":true,
				"sadr":5764832520584332278,
				"asadr":5820324243979633535,
				"aphid":19,
				"nph":4}]}
		]},
	"ie":{
		"gie":[{"tmp":25,"tmpm":85}],
	    "vie":[{"cd":1,"asc":100,"ascq":100,"vndr":[160]}]}
}

*/
type Discs_base map[string]*Smart_info
type Smart_info struct {
	Pgs 	[]int			`json"pgs"`
	We 		Write_errors 	`json"we"` 
	Re 		Read_errors 	`json"re"`
	Ve 		Verify_errors 	`json"ve"`
	Nme		string 			`json"nme"`              //?
	TMP 	Temperature 	`json"tmp"`
	Ssc 	Start_Stop_Cycles `json"ssc"`
	Slf 	string 			`json""slf`                          //?
	Bkg 	Background_scan `json"bkg"`
	Prt 	Protocol_spec 	`json"prt"`
	Ie	    Info_except     `json"ie"`
}

type Write_errors struct {
	Ewd 	int		`json"ewd"`
	Et 		int 	`json"et"`
	Cor 	int64 	`json"cor"`
	Bt 		int64 	`json"bt"`
}

type Read_errors struct {
	Ewd 	int 	`json"ewd"`
	Et 		int 	`json"et"`                               //??
	Cor 	int64 	`json"cor"`
	Bt 		int64 	`json"bt"`
}

type Verify_errors struct {
	Ewd 	int 	`json"ewd"`
	Et 		int 	`json"et"`                               //??
	Cor 	int64 	`json"cor"`
	Bt 		int64 	`json"bt"`
}

type Temperature struct {
	Tmp 	int 	`json"tmp"`
	Tmpm 	int 	`json"tmpm"`	
}

type Start_Stop_Cycles struct {
	My 		int 	`json"my"`
	Mw 		int 	`json"mw"`
	Ay 		int 	`json"ay"`
	Aw 		int 	`json"aw"`
	Sscm 	int 	`json"sscm"`
	Ssc 	int 	`json"ssc"`
	Lucm 	int 	`json"lucm"`
	Luc 	int 	`json"luc"`
}

type Background_scan struct {
	Pwm 	int64 	`json"pwm"`
	Bss 	int 	`json"bss"`
	Nbss 	int 	`json"nbss"`
	Nbms 	int 	`json"nbms"`
}

type Protocol_spec struct {
	Ports 	map[string]*Port_info `json"Ports"`
}

type Port_info struct {
	Pid 	int 			`json"pid"`
	Cd 		int 			`json"cd"`
	Phs 	Spec_mod_param 	`json"phs"`
}

type Spec_mod_param struct {
	Phid 	int 	`json"phid"`
	Adt 	int 	`json"adt"`
	Adr 	int 	`json"adr"`
	Lnk 	int 	`json"lnk"`
	Ismp 	bool 	`json"ismp"`
	Tsmp 	bool 	`json"tsmp"`
	Sadr 	int64 	`json"sadr"`
	Asadr 	int64 	`json"asadr"`
	Aphid 	int 	`json"aphid"`
	Nph 	int 	`json"nph"`
}

type Info_except struct {
	Gie map[string]*generic_spec `json"gie"`
	Vie map[string]*vendor_spec `json"vie"`
}

type generic_spec struct {
	Tmp 	int 	`json"tmp"`
	Tmpm 	int 	`json"tmpm"`
}

type vendor_spec struct{
	Cd 		int 	`json"cd"`
	Asc 	int 	`json"asc"`
	Ascq	int 	`json"ascq"`
    Vndr    map[string]*int `json"vndr"`
}

func get_smart_info(){

}

func main() { 
	get_smart_info()
	//fmt.Printf("hello, world2\n")
}