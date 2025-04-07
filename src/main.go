package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)


func main(){

exibeMenu();
comando:=lecomando();

for{

switch comando{
case 1:
	AnalisarSites()


case 2:

	fmt.Println("Exibindo Logs", "\n")
	imprimeLogs()

case 3:
	fmt.Println("saindo do Sistema")
	os.Exit(0)

default:
	fmt.Println("Comando inválido")
	os.Exit(-1)
}

}
}

func exibeMenu(){

fmt.Println(" 1-Analisar sites"+ "\n" +
			" 2-Exibir Logs"+ "\n" +
			" 3-Sair do sistema" + "\n")			

}

func lecomando()int{
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando Digitado foi",comando)

		return comando
}

func AnalisarSites(){


	fmt.Println("iniciando a analise...")

	for i:=0;i<=5;i++{
		fmt.Println("Monitorando")
		sites := leArquivo()

		for i,site:=range sites{
			fmt.Println("posição",i,"meu site", site)
			testaSite(site)
		}

		time.Sleep(10* time.Minute)
	}




}

func testaSite(site string){
	
	resp,err:= http.Get(site)

	if err!=nil{
		fmt.Println("Ocoreu um erro",err)
	}

	if resp.StatusCode==200{
		fmt.Println("site", site, "foi carregado com sucesso")
		registraLog(site,true)
	}else{
		fmt.Println("site", site, "Está com problemas",resp.StatusCode)
		registraLog(site,false)
	}

}


func leArquivo()[]string{

var sites[]string

 arquivo,err :=os.Open("sites.txt")

 if err!=nil{
	fmt.Println("Não foi possivel ler o arquivo", err)
 }

 leitor:=bufio.NewReader(arquivo)



for{
	linha,err:= leitor.ReadString('\n')
	linha=strings.TrimSpace(linha)
	sites=append(sites,linha)

	if err== io.EOF{
		
		break
	}
	


}

fmt.Println(sites)
arquivo.Close()

return sites
}


func registraLog(site string,status bool){

arquivo,err:= os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)

if err !=nil{

	fmt.Println(err)

}

arquivo.WriteString(time.Now().Format("02/01/2006  15:04:05")+ "-" + site + " "+"online"+ " "+  strconv.FormatBool(status)+ "\n" )

arquivo.Close()
}

func imprimeLogs(){
	arquivo,err:=ioutil.ReadFile("log.txt")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}

