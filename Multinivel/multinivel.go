package main

import (
    "fmt"
)

type process struct {
    burst     int
    t_chegada int
    w_time int
    t_time int
}


func main () {
    var wtime int
    var tturnaroud int 
    var quantum int


	var n_processos int

    p := make([]process, 10)
    fmt.Printf("Número de processos: ")
    fmt.Scanf("%d\n", &n_processos)

    for i := 0; i < n_processos; i++ {
    	fmt.Printf("Burst do processo P%d: ", i)
    	fmt.Scanf("%d\n", &p[i].burst)
    	fmt.Printf("Tempo de chegada do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].t_chegada)
		fmt.Printf("\n")
    }
sou gay
    fmt.Printf("Quantum: ")
    fmt.Scanf("%d\n", &quantum)

    //ROUND ROBIN 1 VEZ
    for i := 0; i < n_processos; i++ {
    	if p[i].burst >= quantum {
    		p[i].burst = p[i].burst - quantum
    	} else if p[i].burst > 0 && p[i].burst < quantum {
    		p[i].burst = p[i].burst - p[i].burst
    	}
    p[i].w_time += quantum - p[i].t_chegada 
    p[i].t_time += p[i].burst
    }


    for i := 0; i < n_processos; i++ {
	if p[i].burst > 0 {
		for i := 1 ; i < n_processos; i++ {
			wtime = p[i-1].burst + p[i-1].w_time
			p[i].w_time = wtime
			}
		for i := 0 ; i < n_processos; i++ {
			tturnaroud += p[i].burst
			p[i].t_time = tturnaroud
			}
		}	
	}

	//Calculando Turnaround

    fmt.Printf("-----------------------------------------------------------------------------\n\n")


    for i := 0; i < n_processos; i++ {
 		fmt.Printf("Tempo de espera P%d: %d\t", i, p[i].w_time)
 		fmt.Printf("Tempo de Turnaround P%d: %d\n", i, p[i].t_time)
 		fmt.Printf("\n")

	}  	

	 	fmt.Printf("Tempo médio de espera: %d\n", (wtime/n_processos))
 		fmt.Printf("Tempo médio de Turnaround: %d", (tturnaroud/n_processos))
}