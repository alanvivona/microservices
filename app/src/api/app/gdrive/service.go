package gdrive

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v3"
)

const (
	clientSecretFilePath string = "/go/src/api/app/gdrive/client_secret.json"
	gdriveAPIbaseURL     string = "https://www.googleapis.com"
)

// GdriveService ...
type GdriveService struct {
	CLIENT *drive.Service
}

// HasClient ...
func (s *GdriveService) HasClient() bool {
	return s.CLIENT != nil
}

// SearchInDoc ...
func (s *GdriveService) SearchInDoc(id string, word string) (bool, error) {
	file, err := s.CLIENT.Files.Get(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	regex := regexp.MustCompile(word)
	return regex.FindString(file.Description) == "", err
}

// CreateFile ...
func (s *GdriveService) CreateFile(file *models.File) (*drive.File, error) {
	driveFile, err := s.CLIENT.Files.Create(&drive.File{Name: file.Name, Description: file.Description}).Do()
	return driveFile, err
}

// CreateClient ...
func (s *GdriveService) CreateClient(c *gin.Context, tokenCode string) error {

	s.CLIENT = nil

	b, err := ioutil.ReadFile(clientSecretFilePath)
	if err != nil {
		log.Fatalf("Unable to initialize auth. Error code: 002", err)
		return err
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to initialize auth. Error code: 003", err)
		return err
	}

	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return err
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok, err := config.Exchange(c, tokenCode)
		if err != nil {
			log.Fatalf("Unable to retrieve token from web %v", err)
		}
		saveToken(cacheFile, tok)
	}

	client := config.Client(c, tok)
	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to initialize client. Error code: 005", err)
		return err
	}
	s.CLIENT = srv
	return nil
}

func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("drive-go.json")), err
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// GetAuthURL ...
func (s *GdriveService) GetAuthURL() (string, error) {
	b, err := ioutil.ReadFile(clientSecretFilePath)
	if err != nil {
		log.Fatalf("Unable to initialize auth. Error code: 002", err)
		return "", err
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to initialize auth. Error code: 003", err)
		return "", err
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL, nil
}
