package main
import "fmt"
func main(){
	var dia int;
	var mes int;
	fmt.Scan(&dia);
	fmt.Scan(&mes);
	if mes == 12{
		if dia < 22{
			fmt.Println("sagitario")
		}else{
			fmt.Println("sapriconio")
		}		
	} else if mes == 1{
		if dia < 20{
			fmt.Println("capricornio")
		}else{
			fmt.Println("acuario")
		}
	}else if mes == 2{
		if dia < 19{
			fmt.Println("acuario")
		}else{
			fmt.Println("piscis")
		}
	}else if mes == 3{
		if dia < 21{
			fmt.Println("piscis")
		}else{
			fmt.Println("aries")
		}	
	}else if mes == 4{
		if dia < 20{
			fmt.Println("aries")
		}else{
			fmt.Println("tauro")
		}
	}else if mes == 5{
		if dia < 21{
			fmt.Println("tauro")
		}else{
			fmt.Println("geminis")
		}
	}else if mes == 6{
		if dia < 21{
			fmt.Println("geminis")
		}else{
			fmt.Println("cancer")
		}
	}else if mes == 7{
		if dia < 23{
			fmt.Println("cancer")
		}else{
			fmt.Println("leo")
		}
	}else if mes == 8{
		if dia < 23{
			fmt.Println("leo")
		}else{
			fmt.Println("virgo")
		}
	}else if mes == 9{
		if dia < 23{
			fmt.Println("virgo")
		}else{
			fmt.Println("libra")
		}
	}else if mes == 10{
		if dia < 23{
			fmt.Println("libra")
		}else{
			fmt.Println("scorpio")
		}
	}else if mes == 11{
		if dia < 22{
			fmt.Println("scorpio")
		}else{
			fmt.Println("sagitario")
		}
	}
}