// https://github.com/alura-cursos/web_com_golang/tree/aula_2
package main

import (
	"go-web/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
