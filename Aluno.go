//go 1.10.4

package main
import 
    (
        "fmt"
        "strconv"
    )
//Estrutura aluno.
type Aluno struct{
    Disciplina
    ra int
    nome string
}
//Estrutura disciplina.
type Disciplina struct{
    disc string

    a1 [3]float64
    a2 [3]float64

    qtd_aulas  int
    qtd_faltas int

    media_final float64
    pct_presenca float64
    status string
}

//Função que retorna a media.
func (aluno *Aluno) calculaMedia() float64{

    var a1 float64
    var a2 float64

    a1 = aluno.a1[0] * 0.7 + aluno.a1[1] * 0.2 + aluno.a1[2] * 0.1
    a2 = aluno.a2[0] * 0.7 + aluno.a2[1] * 0.2 + aluno.a2[2] * 0.1

    aluno.media_final = (a1 + a2 * 2) / 3

    return aluno.media_final
}

//Função que retorna a porcentagem de aulas frequentadas.
func (aluno *Aluno) calculaPresenca() float64{
    
    aluno.pct_presenca = float64 ( 100 - (aluno.qtd_faltas * 100 / aluno.qtd_aulas))

    return aluno.pct_presenca
}

//Função que retornar o status do aluno com base na media e nas aulas frequentadas.
func (aluno *Aluno) calculaStatus() string{

    if (aluno.media_final >= 5.0 && aluno.pct_presenca >= 75.0){
        return "APROVADO"
    }else if ((aluno.media_final >= 3.0 && aluno.media_final < 5) && (aluno.pct_presenca >= 75)){
        return "RECUPERACAO"
    }else if (aluno.media_final < 3.0 && aluno.pct_presenca >= 75){
        return "REPROVADO POR NOTA"
    }else if (aluno.pct_presenca < 75){
        return "REPROVADO POR FALTA"
    }
    return "CASO PERDIDO"

}

/*
    Função para validar numeros.
    Retorna 1 se estiver ok
    Retorna 0 se estiver fora de intervalo
*/

func (aluno *Aluno) validarNumeros() bool{
    var err_validacao string
    var i int
    var iCont int
    var flag int

    err_validacao = " : Número menor que zero ou maior que dez.\n"
    iCont = 1
    flag = 0

    //Valida vetor a1.
    for i = 0 ; i < 3; i++{

        if (aluno.a1[i] < 0 || aluno.a1[i] > 10){
                
            flag = 1

            if (i == 0){
                fmt.Printf("p" + strconv.Itoa(iCont) + err_validacao)
            }
            if(i == 1){
                fmt.Printf("ma" + strconv.Itoa(iCont) + err_validacao)
            }
            if (i == 2){
                fmt.Printf("mb" + strconv.Itoa(iCont) + err_validacao)
            }
        }
    }

    iCont += 1

    //Valida vetor a2.
    for i = 0 ; i < 3; i++{

        if (aluno.a2[i] < 0 || aluno.a2[i] > 10){
                
            flag = 1

            if (i == 0){
                fmt.Printf("p" + strconv.Itoa(iCont) + err_validacao)
            }
            if(i == 1){
                fmt.Printf("ma" + strconv.Itoa(iCont) + err_validacao)
            }
            if (i == 2){
                fmt.Printf("mb" + strconv.Itoa(iCont) + err_validacao)
            }
        }
    }

    if (flag == 1){
        return false
    }

    return true
}

func main(){

    aluno1 := &Aluno{}
    aluno1.disc = "pkoop"    
    
    fmt.Printf("Nome da disciplina: ")
    fmt.Scan(&aluno1.disc)

    fmt.Printf("P1...MA1...MB1: ")
    fmt.Scan(&aluno1.a1[0])
    fmt.Scan(&aluno1.a1[1])
    fmt.Scan(&aluno1.a1[2])

    fmt.Printf("\n")
    
    fmt.Printf("P2...MA2...MB2: ")
    fmt.Scan(&aluno1.a2[0])
    fmt.Scan(&aluno1.a2[1])
    fmt.Scan(&aluno1.a2[2])

    fmt.Printf("qtd aulas...qtd faltas: ")
    fmt.Scan(&aluno1.qtd_aulas)
    fmt.Scan(&aluno1.qtd_faltas)

    fmt.Printf("\n")

    aluno1.media_final = 0
    aluno1.pct_presenca = 0
    aluno1.status = ""

    if (!aluno1.validarNumeros()){
        fmt.Printf("\nOppsss...erros ocorreram na validação dos números...\n")
        return
    }

    /*
    //Criação do aluno.
    aluno1 := &Aluno{
        Disciplina{
            "NomeDaDisciplina",
            10,8,7,5,5,6,20,4,0,0,"",
        },
        103758,
        "José",
    }
    */
    fmt.Printf("Media...: %f \n", aluno1.calculaMedia())
    fmt.Printf("Presenca: %f \n", aluno1.calculaPresenca())
    fmt.Printf("Status..: %s \n", aluno1.calculaStatus())
}
