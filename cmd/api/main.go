package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/marcelowf/First-Go-API/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	//Logger configurado para incluir o arquivo e a linha
	log.SetReportCaller(true)
	//Roteador para gerenciar requisições e respostas
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)
  	//Iniciando servidor HTTP
  	erro := http.ListenAndServe("localhost:8000", router)
  	if erro != nil {
		log.Error(erro)
	}
}