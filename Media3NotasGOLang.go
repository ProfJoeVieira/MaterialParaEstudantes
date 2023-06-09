package main

import (
    "fmt"
    "strconv"
)

func main() {
    for {
        var num1, num2, num3, media float64
        var input string

        for {
            fmt.Print("Digite a primeira nota: ")
            fmt.Scanln(&input)
            if n, err := strconv.ParseFloat(input, 64); err == nil {
                num1 = n
                break
            } else {
                fmt.Println("Digite apenas números. Tente novamente.")
            }
        }

        for {
            fmt.Print("Digite a segunda nota: ")
            fmt.Scanln(&input)
            if n, err := strconv.ParseFloat(input, 64); err == nil {
                num2 = n
                break
            } else {
                fmt.Println("Digite apenas números. Tente novamente.")
            }
        }

        for {
            fmt.Print("Digite a terceira nota: ")
            fmt.Scanln(&input)
            if n, err := strconv.ParseFloat(input, 64); err == nil {
                num3 = n
                break
            } else {
                fmt.Println("Digite apenas números. Tente novamente.")
            }
        }

        media = (num1 + num2 + num3) / 3
        fmt.Printf("A média das notas digitadas é: %.2f\n", media)

        if media < 3 {
            fmt.Println("O aluno está reprovado")
        } else if media >= 3 && media < 6 {
            fmt.Println("O aluno está em recuperação")
        } else {
            fmt.Println("O aluno está aprovado")
        }

        var repeat string
        fmt.Print("Deseja fazer outra vez? (s/n)")
        fmt.Scanln(&repeat)
        if repeat != "s" && repeat != "S" {
            break
        }
    }
}


