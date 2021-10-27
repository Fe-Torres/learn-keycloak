package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = "myclient"
	clientSecret = "d560fb6a-6a14-4439-a92b-7742c3ac2c93"
)

func main() {
	//Pacote para controlar o fluxo da solicitação
	//Quando a gente faz a solicitação, de acordo com o que a gente quiser, podemos parar a solicitação pela metade
	//Raramente utilizamos... Alguns pacotes que solicitam.
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/myrealm")
	if err != nil {
		log.Fatal(err)
	}

	//Redirect url = Depois de logado
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	//É recomendável que seja dinâmico
	//Mas onde irei gerar e onde ele ficará armazenado?
	state := "123"

	//Toda vez que ele acessar essa url "/", ele irá redirecionar o usuário para
	// o writer vai redireicionar ele para uma página que exige uma autenticação
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("state") != state {
			http.Error(writer, "State inválido", http.StatusBadRequest)
			return
		}

		//Trocando o code da url para um token
		token, err := config.Exchange(ctx, request.URL.Query().Get("code"))
		if err != nil {
			http.Error(writer, "Falha ao trocar o token", http.StatusInternalServerError)
			return
		}

		//Gerando o id token para usar no processo de autenticação
		IdToken, ok := token.Extra("id_token").(string)
		if !ok {
			http.Error(writer, "Falha ao gerar o IdToken", http.StatusInternalServerError)
			return
		}

		//Pegando as informações do usuário, precisa do token!
		userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(token))
		if !ok {
			http.Error(writer, "Erro ao pegar as informações do usuário", http.StatusInternalServerError)
			return
		}

		//http://localhost:8080/auth/realms/myrealm/protocol/openid-connect/userinfo

		//Resposta
		resp := struct {
			AccessToken *oauth2.Token
			IDToken     string
			UserInfo    *oidc.UserInfo
		}{
			AccessToken: token,
			IDToken:     IdToken,
			UserInfo:    userInfo,
		}

		//Pegando os dados da resposta da rota de callback
		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		//response
		writer.Write(data)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
