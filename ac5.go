package main

import "fmt"
 
func main() {
	fmt.Println("Hello, World!")

  var a, b, c float64

  fmt.Println("Digite os três lados do triângulo:")
  fmt.Scan(&a, &b, &c)

  tipoTriangulo(a, b, c)

var T int

fmt.Println("Digite o número de casos de teste:")
fmt.Scan(&T)

for i := 0; i < T; i++ {
    var PA, PB int
    var G1, G2 float64

    fmt.Println("Digite a população de A, a população de B e as taxas de crescimento de A e B (em %):")
    fmt.Scan(&PA, &PB, &G1, &G2)

    fmt.Println(crescimentoPop(PA, PB, G1, G2))
 }
}

func tipoTriangulo(a, b, c float64) {
   
    if a < b {
        a, b = b, a
    }
    if a < c {
        a, c = c, a
    }
    if b < c {
        b, c = c, b
    }

    if a >= b+c {
        fmt.Println("NAO FORMA TRIANGULO")
    } else {
        if a*a == b*b+c*c {
            fmt.Println("TRIANGULO RETANGULO")
        }
        if a*a > b*b+c*c {
            fmt.Println("TRIANGULO OBTUSANGULO")
        }
        if a*a < b*b+c*c {
            fmt.Println("TRIANGULO ACUTANGULO")
        }
        if a == b && b == c {
            fmt.Println("TRIANGULO EQUILATERO")
        } else if a == b || b == c || a == c {
            fmt.Println("TRIANGULO ISOSCELES")
        }
    }
}

func crescimentoPop(PA, PB int, G1, G2 float64) string {
    anos := 0
    for PA <= PB {
        PA += int(float64(PA) * G1 / 100)
        PB += int(float64(PB) * G2 / 100)
        anos++
        if anos > 100 {
            return "Mais de 1 seculo."
        }
    }
    return fmt.Sprintf("%d anos.", anos)
}
