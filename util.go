package main

import (
	//"errors"
	//"flag"
	"fmt"
        "time"
	"./webdavclnt-master"
        "bytes"
	"io/ioutil"
	"os"
	"strings"
        "os/exec"
        "path/filepath"
)



func main() {
        //reading user,password 
	standartroot :="https://cloud.lab215.com/remote.php/dav/files/"
        logpass, err1 := ioutil.ReadFile("/tmp/login:password.txt")//home/sergei/
        if err1 != nil {
          panic(err1)
        }
        str := string(logpass)
        usr := strings.Split(str,":")[0]
        pas := strings.Split(str,":")[1]
        slashedusr := strings.Join([]string{usr,"/"},"")
        root := strings.Join([]string{standartroot,slashedusr},"")//
        //fmt.Println(usr)
        //fmt.Println(pas)
        //fmt.Println(root) 
        client := webdavclnt.
		NewClient(root).
		SetLogin(usr).
		SetPassword(pas)
        //request all files with data modified option
        data_response_map, err2 := client.PropFind("./logs","getlastmodified")
        if err2 != nil {
	  panic(err2)
        }
        fmt.Println(data_response_map)
        //searching last modified file
        lastmodif,err3 := latest_arch(data_response_map)
        if err3 != nil {
	  panic(err3)
        }
        fmt.Println(lastmodif)

        //Getting a file and save
        filePath := filepath.Join(strings.Split(lastmodif,"/")[5],strings.Split(lastmodif,"/")[6])
        fmt.Println(filePath)
        archdata,err4 := client.Get(filePath)//archdata
        if err4 != nil {
	  panic(err4)
        }
        

        //for file
        // filename - name.tar.gz
        
        filename := filepath.Join("",strings.Split(lastmodif,"/")[6])//"./tmp/"
        dirname := filepath.Join("TESTS/",strings.Split(filename,".")[0])
        filename = filepath.Join(dirname,filename)
        os.MkdirAll(dirname,os.ModePerm)
        file, err5 := os.Create(filename)//(filename)
        if err5!=nil {
          return
        }

        defer file.Close()

        fmt.Fprintf(file,"%+s", archdata)
        //return
        //unzip
        fmt.Println(filename, dirname)
        cmd := exec.Command("tar","-xvzf", filename,"-C",dirname) // "-r" - поиск архивов и во вложенных папках
        
        //fmt.Println(filename, dirname)
        cmd.Stdin = strings.NewReader("")
         var out bytes.Buffer
        cmd.Stdout = &out
        err6 := cmd.Run()
        if err6 != nil {
            panic(err6)
        }
        cmd2 := exec.Command("rm",filename)         
        cmd2.Stdin = strings.NewReader("")
        //var out bytes.Buffer
        cmd2.Stdout = &out
        err7 := cmd2.Run()
        if err7 != nil {
            panic(err7)
        }
        return 
        
}

//search latest arch
func latest_arch(datamap map[string]webdavclnt.Properties)(string, error) {
  is_started := true
  latest_root :="fileroot"
  var latest_data time.Time 
  latest_data,err := time.Parse("Mon, 02 Jan 2006 15:04:05","Mon, 02 Jan 2006 15:04:05")
  if err != nil {
	//panic(err)
        return string(""),err
    }
  //all datamap
  for Key_root,Value_opt := range datamap { 
    dataTime,_ := Value_opt["getlastmodified"]
    FmtdataTime := dataTime[:len(dataTime)-4]
    time_parsed,err := time.Parse("Mon, 02 Jan 2006 15:04:05",FmtdataTime)
    if err != nil {
	//panic(err)
        return string(""),err
    }
    //fmt.Println(t,FmtdataTime)
    if( is_started == true || latest_data.Before(time_parsed)) {
      is_started = false
      latest_data = time_parsed
      if(strings.Compare(Key_root,"/remote.php/dav/files/sergei.smirnov/logs/") != 0){
        latest_root = Key_root}
    }
  }
  fmt.Println(latest_root,latest_data)
  return latest_root,nil
}
