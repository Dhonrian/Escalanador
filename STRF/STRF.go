package main

import (
    "fmt"
)

func main() {

	var menor int 
	var restante int
	var wait int
	var turnaround int
	var n_processos int

	var t_chegada [10]int
	var burst [10]int
	var r_time [10]int

	restante = 0

	fmt.Printf("Número de processos: ")
    fmt.Scanf("%d\n", &n_processos)

    //t_chegada := make([]int, n_processos)
    //burst := make([]int, n_processos)
    //r_time := make([]int, n_processos)
    	



    for i := 0; i < n_processos; i++{
    	fmt.Printf("Burst do processo P%d: ", i)
    	fmt.Scanf("%d\n", &burst[i])
    	fmt.Printf("Tempo de chegada do processo P%d: ", i)
		fmt.Scanf("%d\n", &t_chegada[i])
		r_time[i] = burst[i]
 		fmt.Printf("\n")

    }

    r_time[len(r_time)-1] = 9999



    for tempo := 0; restante != n_processos; tempo++ {
    	//fmt.Printf("%d\n", tempo)
    	menor = 9
    	for i := 0; i < n_processos; i++{
    		//fmt.Printf("%d\n", i)

    		if t_chegada[i] <= tempo && r_time[i] < r_time[menor] && r_time[i] > 0 {
    			menor = i
    		}
    	}

    	r_time[menor]--;
    	if r_time[menor] == 0{
    		restante++
    		tempoFinal := tempo+1
    		fmt.Printf("\nP%d\tTurnaroud time: %d\t\tWaiting time: %d", menor, (tempoFinal - t_chegada[menor]), (tempoFinal-burst[menor]-t_chegada[menor]))
    		wait += (tempoFinal - burst[menor] - t_chegada[menor])
    		turnaround += (tempoFinal - t_chegada[menor])	
    	}
    }

	fmt.Printf("\n\nTempo médio de espera: %d\n", (wait/n_processos))
 	fmt.Printf("Tempo médio de Turnaround: %d", (turnaround/n_processos))
}