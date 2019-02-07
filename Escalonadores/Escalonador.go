//Feito por Leonardo Tamanhão
//Escalonadores em GO - ATIVIDADE SO I 
//Diogo Branquinho



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


func fcfs() {

    var wtime int
    var tturnaroud int 

	var n_processos int

    p := make([]process, 10)
    fmt.Printf("Número de processos: ")
    fmt.Scanf("%d\n", &n_processos)

    for i := 0; i < n_processos; i++ {
    	fmt.Printf("Burst do processo P%d: ", i)
    	fmt.Scanf("%d\n", &p[i].burst)
    	fmt.Printf("Tempo de chegada do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].t_chegada)
		fmt.Printf("Prioridade do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].prioridade)
		fmt.Printf("Quantum do processo P%d: ", i)
		fmt.Scanf("%d\n", &p[i].quantum)
		fmt.Printf("\n")
    }

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
		fmt.Printf("Burst P%d: %d\t", i, p[i].burst) 
 		fmt.Printf("Tempo de chegada P%d: %d\t", i, p[i].t_chegada)
 		fmt.Printf("Prioridade P%d: %d\t", i, p[i].t_chegada)
 		fmt.Printf("Quantum P%d: %d\n", i, p[i].t_chegada)

 		fmt.Printf("Tempo de espera P%d: %d\t", i, p[i].w_time)
 		fmt.Printf("Tempo de Turnaround P%d: %d\n", i, p[i].t_time)
 		fmt.Printf("\n")

	}  	

	 	fmt.Printf("Tempo médio de espera: %d\n", (wtime/n_processos))
 		fmt.Printf("Tempo médio de Turnaround: %d", (tturnaroud/n_processos))
}

func sjf() {
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

func srtf () {

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
    		fmt.Printf("\nP%d\tTurnaround time: %d\t\tWaiting time: %d", menor, (tempoFinal - t_chegada[menor]), (tempoFinal-burst[menor]-t_chegada[menor]))
    		wait += (tempoFinal - burst[menor] - t_chegada[menor])
    		turnaround += (tempoFinal - t_chegada[menor])	
    	}
    }

	fmt.Printf("\n\nTempo médio de espera: %d\n", (wait/n_processos))
 	fmt.Printf("Tempo médio de Turnaround: %d", (turnaround/n_processos))
}

func rr(){

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

func multinivel () {
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
    fmt.Printf("--------------------------------------------------------------------------\n\n")


    for i := 0; i < n_processos; i++ {
 		fmt.Printf("Tempo de espera P%d: %d\t", i, p[i].w_time)
 		fmt.Printf("Tempo de Turnaround P%d: %d\n", i, p[i].t_time)
 		fmt.Printf("\n")

	}  	

	 	fmt.Printf("Tempo médio de espera: %d\n", (wtime/n_processos))
 		fmt.Printf("Tempo médio de Turnaround: %d", (tturnaroud/n_processos))
}


func menu () {

	var opcao int


	fmt.Printf("Escolha o escalonador: \n")
	fmt.Printf("1 - FCFS\n")
	fmt.Printf("2 - SJF\n")
	fmt.Printf("3 - SRTF\n")
	fmt.Printf("4 - Round Robin\n")
	fmt.Printf("5 - Multinivel\n")
	fmt.Printf("Escolha:")
    fmt.Scanf("%d\n", &opcao)

    for {

	   	switch opcao {
		    case 1 :
		    	fcfs()
		    	break
		    case 2:
		    	sjf()
		    	break
		    case 3:
		    	srtf()
		    	break
		    case 4:
		    	rr()
		    	break
		    case 5:
		    	multinivel()
		    	break
		    }
		  break  
	}

}


func main () {

	fmt.Printf("****Feito por Leonardo*****\n")

	for {
		fmt.Printf("\n")
		menu()
	} 

}