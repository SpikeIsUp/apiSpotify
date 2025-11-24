package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiData struct {
	Results[] struct {
		Poke string `json:"id"`
	}
}

func main() {
    // URL de l’API
    urlApi := "https://accounts.spotify.com/api/token"

    // Initialisation du client HTTP qui va émettre/demander les requêtes
    httpClient := http.Client{
        Timeout: time.Second * 2, // Timeout après 2sec
    }

    // Création de la requête HTTP vers l’API avec initialisation de la méthode HTTP, la route et le corps de la requête
    req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
    if errReq != nil {
        fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
    }

    // Ajout d’une métadonnée dans le header, User-Agent permet d’identifier l’application, système….
    req.Header.Add("User-Agent", "Ynov Campus Cours")

    // Exécution de la requête HTTP vers l’API
    res, errResp := httpClient.Do(req)
    if errResp != nil {
        fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
        return
    }

    if res.Body != nil {
        defer res.Body.Close()
    }

    // Lecture et récupération du corps de la requête HTTP
    body, errBody := io.ReadAll(res.Body)
    if errBody != nil {
        fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
    }

    // Déclaration de la variable qui va contenir les données
    var decodeData ApiData

    // Décodage des données en format JSON et ajout des donnée à la variable decodeData
    json.Unmarshal(body, &decodeData)

    // Affichage des données
    fmt.Println(decodeData.Results[0])
}
