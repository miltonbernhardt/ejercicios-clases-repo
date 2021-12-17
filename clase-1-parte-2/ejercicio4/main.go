package main

import (
	"fmt"
	"time"
)

func main() {
	var mes int8 = 1
	 meses := map[int8]string{}
	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"


fmt.Println(meses[mes])

	//otra forma

	fmt.Println(time.Month(mes).String())
}
