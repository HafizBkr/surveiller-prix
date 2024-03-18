package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/surveillance-prix/scraping"
)

type Computer struct {
	Name  string
	Price string
}

func main() {
	// Récupérer les données des ordinateurs depuis Amazon
	amazonComputers, err := ScrapeAmazonComputers("ordinateur")
	if err != nil {
		log.Fatal("Erreur lors du scraping depuis Amazon :", err)
	}

	// Récupérer les données des ordinateurs depuis Coin Afrique
	coinAfriqueComputers, err := ScrapeCoinAfriqueComputers("ordinateur")
	if err != nil {
		log.Fatal("Erreur lors du scraping depuis Coin Afrique :", err)
	}

	// Créer une nouvelle instance de Gin
	router := gin.Default()

	// Définir la route pour le dashboard
	router.GET("/", func(c *gin.Context) {
		// Passez les données des ordinateurs au modèle HTML
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"AmazonComputers":    amazonComputers,
			"CoinAfriqueComputers": coinAfriqueComputers,
		})
	})

	// Démarrer le serveur
	router.Run(":8080")
}
