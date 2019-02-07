package main
import (
    "fmt"
    "sort"
)

type process struct {
    id int
    burst     int
    t_chegada int
    prioridade int
    quantum int
    w_time int
    t_time int
}

func main() {

    var wtime int
    var tturnaroud int 

	var n_processos int

    


    fmt.Printf("Número de processos: ")
    fmt.Scanf("%d\n", &n_processos)


    p := make([]process, n_processos)

    for i := 0; i < n_processos; i++ {
    	fmt.Printf("Burst do processo P%d: ", i)
    	fmt.Scanf("%d\n", &p[i].burst)
    	fmt.Printf("Tempo de chegada do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].t_chegada)
		fmt.Printf("Prioridade do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].prioridade)
		fmt.Printf("Quantum do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].quantum)
        p[i].id = i
		fmt.Printf("\n")
    }


sort.Slice(p[:], func(i, j int) bool {
  return p[i].burst < p[j].burst
})


    //Calculando waiting time
	for i := 1 ; i < n_processos; i++ {
		wtime = p[i-1].burst + p[i-1].w_time
		p[i].w_time = wtime
	}

	//Calculando Turnaround
	for i := 0 ; i < n_processos; i++ {
		tturnaroud += p[i].burst
		p[i].t_time = tturnaroud
	}

    fmt.Printf("-----------------------------------------------------------------------------\n\n")


    for i := 0; i < n_processos; i++ {
		fmt.Printf("Burst P%d: %d\t", p[i].id, p[i].burst) 
 		fmt.Printf("Tempo de chegada P%d: %d\t", p[i].id, p[i].t_chegada)
 		fmt.Printf("Prioridade P%d: %d\t", p[i].id, p[i].prioridade)
 		fmt.Printf("Quantum P%d: %d\n", p[i].id, p[i].quantum)

 		fmt.Printf("Tempo de espera P%d: %d\t", p[i].id, p[i].w_time)
 		fmt.Printf("Tempo de Turnaround P%d: %d\n", p[i].id, p[i].t_time)
 		fmt.Printf("\n")

	}  	

	 	fmt.Printf("Tempo médio de espera: %d\n", (wtime/n_processos))
 		fmt.Printf("Tempo médio de Turnaround: %d", (tturnaroud/n_processos))
}