package gl

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func Ept() int {
	return 1
}

// https://developers.google.com/sheets/api/quickstart/go
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Read() [2]string {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1PKORfLXcTv9CK1gBbPPO4RDY0_XJA6n83Mgz2nO0oy4"
	//16K8Yxnq1s6a4COX1srG3HDa6eAVyjtGcsOYjdl-zP0M
	readRange := "день!E4:E8"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var str string
	str = ""
	if len(resp.Values) == 0 {
		str = "No data found."
	} else {

		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.

			//fmt.Printf("%T\n", row[0])
			//	fmt.Println(row)
			if row == nil {
				os.Exit(3)
				str = str + fmt.Sprintf(" %v", row[0]) + "  ... " + "\r\n"
			} else {

				str = str + fmt.Sprintf(" %v", row) + "  ... " + "\r\n"
			}

		}
	}

	var dayMonth [2]string
	dayMonth[0] = str

	readRange = "10_дней!E4:E6"
	resp, err = srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	str = ""
	if len(resp.Values) == 0 {
		str = "No data found."
	} else {

		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			if row[0] != nil {
				str = str + fmt.Sprintf(" %v", row) + "  ... " + "\r\n"
			}

		}
	}
	dayMonth[1] = str
	return dayMonth
}
