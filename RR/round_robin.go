package main


import (
    "fmt"
)


func main() {

	var n_processos int
	var quantum int
	var flag int
	var restante int
	var w_time int
	var turnaround int
	flag = 0
	w_time = 0
	turnaround = 0

	var at[10] int
	var burst[10] int
	var r_time[10] int
	var tempo int
	var i int

    fmt.Printf("Número de processos: ")
    fmt.Scanf("%d\n", &n_processos)
    restante  = n_processos

    for i := 0; i < n_processos; i++ {
    	fmt.Printf("Burst do processo P%d: ", i)
    	fmt.Scanf("%d\n", &burst[i])
    	fmt.Printf("Tempo de chegada P%d: ", i)
    	fmt.Scanf("%d\n", &at[i])
    	r_time[i] = burst[i]
    }

    fmt.Printf("Quantum: ")
    fmt.Scanf("%d\n", &quantum)

    fmt.Printf("Processo\tTurnaround\tWaiting Time\n")

    tempo = 0
    for i = 0; restante != 0;{

       	
    	if r_time[i] <= quantum && r_time[i] > 0 {
    		tempo += r_time[i]
    		r_time[i] = 0
    		flag = 1
    	} else if r_time[i] > 0 {
    		r_time[i] -= quantum
    		tempo += quantum
    	} 

    	if r_time[i] == 0 && flag == 1 {
    		restante--

    		//fmt.Printf("\nP%d\t%d\t%d\t%d\n", i,tempo, at[i],burst[i])
    		fmt.Printf("P%d\t\t%d\t\t%d\n", i, tempo-at[i], tempo-at[i]-burst[i])
    		w_time += tempo - at[i] - burst[i]
    		turnaround += tempo - at[i]
    		flag = 0
    	}

    	if i == n_processos-1 {
    		i = 0
    	} else if at[i+1] <= tempo {
    		i++
    	} else {
    		i = 0
    	}
    }


    fmt.Printf("Tempo médio de espera: %d\n", (w_time/n_processos))
 	fmt.Printf("Tempo médio de Turnaround: %d", (turnaround/n_processos))


}



