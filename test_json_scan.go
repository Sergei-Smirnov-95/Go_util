package main

import ("fmt"
		"encoding/json"
		"strings"
		"time"
)

type Discs_base struct {
	Smarts map[string]*Smart_info 
}

type Smart_info struct {
	Pgs 	[]int			`json:"pgs"`
	We 		Write_errors 	`json:"we"` 
	Re 		Read_errors     `json:"re"`
	Ve 		Verify_errors 	`json:"ve"`
	Nme		Non_med_err		`json:"nme"`              
	TMP 	Temperature 	`json:"tmp"`
	Ssc 	Start_Stop_Cycles `json:"ssc"`
	Slf 	Self_test		`json:"slf"`                          
	Bkg 	Background_scan `json:"bkg"`
	Prt 	Protocol_spec 	`json:"prt"`
	Ie	    Info_except     `json:"ie"`
	TimesTemp      time.Time          
}

type Non_med_err struct{}

type Self_test struct{}

type Write_errors struct {
	Ewd 	int		`json:"ewd"`
	Et 		int 	`json:"et"`
	Cor 	uint64 	`json:"cor"`
	Bt 		uint64 	`json:"bt"`
}

type Read_errors struct {
	Ewd 	int		`json:"ewd"`
	Et 		int 	`json:"et"`
	Cor 	uint64 	`json:"cor"`
	Bt 		uint64 	`json:"bt"`
}

type Verify_errors struct {
	Ewd 	int 	`json:"ewd"`
	Et 		int 	`json:"et"`                               
	Cor 	uint64 	`json:"cor"`
	Bt 		uint64 	`json:"bt"`
}

type Temperature struct {
	Tmp 	int 	`json:"tmp"`
	Tmpm 	int 	`json:"tmpm"`	
}

type Start_Stop_Cycles struct {
	My 		int 	`json:"my"`
	Mw 		int 	`json:"mw"`
	Ay 		int 	`json:"ay"`
	Aw 		int 	`json:"aw"`
	Sscm 	int 	`json:"sscm"`
	Ssc 	int 	`json:"ssc"`
	Lucm 	int 	`json:"lucm"`
	Luc 	int 	`json:"luc"`
}

type Background_scan struct {
	Pwm 	uint64 	`json:"pwm"`
	Bss 	int 	`json:"bss"`
	Nbss 	int 	`json:"nbss"`
	Nbms 	int 	`json:"nbms"`
}

type Protocol_spec struct {
	Ports 	[]Port_info `json:"Ports"`
}

type Port_info struct {
	Pid 	int 			`json:"pid"`
	Cd 		int 			`json:"cd"`
	Phs 	[]Spec_mod_param 	`json:"phs"`
}

type Spec_mod_param struct {
	Phid 	int 	`json:"phid"`
	Adt 	int 	`json:"adt"`
	Adr 	int 	`json:"adr"`
	Lnk 	int 	`json:"lnk"`
	Ismp 	bool 	`json:"ismp"`
	Tsmp 	bool 	`json:"tsmp"`
	Sadr 	uint64 	`json:"sadr"`
	Asadr 	uint64 	`json:"asadr"`
	Aphid 	int 	`json:"aphid"`
	Nph 	int 	`json:"nph"`
}

type Info_except struct {
	Gie []Temperature `json:"gie"`//Generic_spec `json:"gie"`
	Vie []Vendor_spec `json:"vie"`
}
/*
type Generic_spec struct {
	Tmp 	int 	`json:"tmp"`
	Tmpm 	int 	`json:"tmpm"`
}
*/
type Vendor_spec struct{
	Cd 		int 	`json:"cd"`
	Asc 	int 	`json:"asc"`
	Ascq	int 	`json:"ascq"`
    Vndr    []int `json:"vndr"`
}

func Get_Smart_Info(pars_string string) (*Smart_info, error ){
	si := &Smart_info{}
	splits := strings.Split(pars_string,"smart info: ")
	time_id := strings.Split(splits[0]," ")
	timestemp,_ := time.Parse(time.RFC3339,time_id[0])
	si.TimesTemp = timestemp
	fmt.Println(time_id[9]) // как id в mapе Smart_Info !!!  
	
	pars_json := []byte(splits[1])

	err := json.Unmarshal(pars_json, si) 
	//fmt.Println()
	if err != nil {
		return nil,err
	}
    return si,nil
}

func main() { 
	
	str:=`2018-05-03T14:53:36.248715+03:00 sp-124-168 hwmgr[238748]: I0503 14:53:36.248664  238748 scsi.go:144] Disk scsi-35000cca2530c83f4 smart info: {"pgs":[0,2,3,5,6,8,13,14,15,16,21,24,25,26,47,48,55],"we":{"cor":5639,"bt":1888071680},"re":{"cor":3041,"bt":1131528192},"ve":{"cor":2692},"nme":{},"tmp":{"tmp":25,"tmpm":85},"ssc":{"my":2017,"mw":37,"ay":2017,"aw":37,"sscm":50000,"ssc":12,"lucm":600000,"luc":95},"slf":{},"prt":{"Ports":[{"pid":1,"cd":2,"phs":[{"adt":2,"adr":1,"lnk":11,"ismp":true,"tsmp":true,"sadr":5764832520584332277,"asadr":5820324243979633023,"aphid":19,"nph":4}]},{"pid":2,"cd":2,"phs":[{"phid":1,"adt":2,"adr":1,"lnk":11,"ismp":true,"tsmp":true,"sadr":5764832520584332278,"asadr":5820324243979633535,"aphid":19,"nph":4}]}]},"ie":{"gie":[{"tmp":25,"tmpm":85}],"vie":[{"cd":1,"asc":100,"ascq":100,"vndr":[160]}]}}
	`
	obj,err:= Get_Smart_Info(str)
	if err!= nil {
		panic(err)
	}
    fmt.Println(obj)
	
}