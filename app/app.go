package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/qeesung/image2ascii/convert"
	"github.com/urfave/cli"
)

type Follower struct {
	Login        string `json:"login"`
	ID           int    `json:"id"`
	AvatarURL    string `json:"avatar_url"`
	URL          string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	Type         string `json:"type"`
	SiteAdmin    bool   `json:"site_admin"`
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "My github on terminal made with GoLang"
	app.Usage = "Basicaly you can get yours public projects and followers by terminal flags"

	app.Commands = []cli.Command{
		{
			Name:   "mygitlist",
			Usage:  "Get your followers",
			Action: getMyFollowers,
		},
	}

	return app
}

func getMyFollowers(c *cli.Context) {

	token := os.Getenv("GITHUB_TOKEN")
	url := os.Getenv("GITHUB_URL")

	if token == "" {
		fmt.Println("GitHub access token not found on .env file")
		return
	}

	if url == "" {
		fmt.Println("URL not found on .env file")
		return
	}

	// Make a new HTTP GET request
	req, err := http.NewRequest("GET", url+"followers", nil)
	if err != nil {
		fmt.Println("Erro to create new request:", err)
		return
	}

	// Change "YOUR_AUTH_TOKEN" to your authtoken from GitHub
	req.Header.Set("Authorization", "token "+token)

	// Make a HTTP client
	client := &http.Client{}

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro to send request:", err)
		return
	}

	// Close request after
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro to read response:", err)
		return
	}

	// Create a slice variable to store result
	var followers []Follower

	// Do JSON parse to feed followers slice
	err = json.Unmarshal(body, &followers)
	if err != nil {
		fmt.Println("Erro on JSON parsing:", err)
		return
	}

	SaveAvatar(followers)
}

func SaveAvatar(followers []Follower) {
	folderPath := "./followers"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755) // Permissões 0755 para a pasta
		if err != nil {
			fmt.Println("Erro to create folder:", err)
			return
		}
	}

	for _, follower := range followers {
		imageURL := "https://github.com/" + follower.Login + ".png"
		filename := filepath.Join(folderPath, filepath.Base(imageURL))

		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println("Erro to request HTTP:", err)
			return
		}
		defer response.Body.Close()

		// Check if the request was successful (code 200)
		if response.StatusCode != http.StatusOK {
			fmt.Println("Erro: status code isn not 200 OK")
			return
		}

		outputFile, err := os.Create(filename)
		if err != nil {
			fmt.Println("Erro to create file:", err)
			return
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, response.Body)
		if err != nil {
			fmt.Println("Erro copy response content to file:", err)
			return
		}
		fmt.Println("Save imagem with successful!")

		canal := PrintAvatar(filename)
		fmt.Print(<-canal)

	}
}

func PrintAvatar(filename string) <-chan string {
	canal := make(chan string)

	go func() {
		time.Sleep(time.Second)
		convertOptions := convert.DefaultOptions
		convertOptions.FixedWidth = 100
		convertOptions.FixedHeight = 40
		converter := convert.NewImageConverter()
		canal <- converter.ImageFile2ASCIIString(filename, &convertOptions)
	}()

	return canal
}
