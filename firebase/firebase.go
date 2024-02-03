package firebase

import (
	"context"
	"fmt"
	"time"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/samuelmunyoki/ProxyDNS/utils"

	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
var ProxyDNSFirestore *firestore.Client

// Initializes connection to firebase
func InitFirestore()(error){

	config := firebase.Config{
		ProjectID: "newflixbot",
	}

	opt := option.WithCredentialsJSON([]byte(`{
	  "type": "service_account",
	  "project_id": "newflixbot",
	  "private_key_id": "0c82a1a9dc7f93286bfaba78c726b6a042165037",
	  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQD0AeINKtURp88r\nv7VlLil5zXGKUQOxmNFrfKxOzxVRx6YqoGU1sEnKAtanOQAw9TrRt1sP7AeCue5A\n/IlAV2lNR3sKSctblbPEuhYjh2T4kP2SntnMGZ2bCAlqUw+QALHf2h30H7BQjyoz\n4huDjXvIBu3I1jBQkxPPRIjwAWcxvNAzgLSR60Vg/7TXuOBXHj1Usuv9xvws/X+z\nqtU9gXHWnxuUfv6FqKx7os78NgupeL8cUdtzDbM0phtYPuP6k5boWg1dtDuyebRb\neVm1dsFXV6O5x1CVbFxaIYHyDLibXm+UyasAw9rmpNW+M8cPccJk4JmJ7YXLicy1\nxBp0QSI9AgMBAAECggEADNwki8kcHagYdRfPeZurN+4p8749UZjaQK37btPfLcY1\n7b0yWFgIK4tmwL1yUyI5jV/6fqZT5wHhmq80lJ2GwTnpNCubeiIzrUSZchnqqmcJ\n1jZlCCq5cbhEtsPV6CMBPOkD9x/MbRJ+iOl7xb0pLuuekJ0pQrXdr2jPRsbsJ2TP\nw9zSoJKy/abB1wN2dfHKuhJIBuiE4WE1qKklSWPny3ZXUfnXF/swv3/6C7IbRd3n\nJChJHQeC6CXCMr6lEb9U0/6YYfMn2T2AAZxiCjuosymk72vFvJYaaUvORysNbP+l\nSRLoHmxStQ8J+wS1bKKoGGuMAIWnRLY3UgCzZUJBAQKBgQD7rOJKDva+EPMey0B+\nzdpz+t67oEDSCM7i450oSzdYc8ZtHkHz31E33aVJDNRVPshgDDDM5H4wNCBPRlY0\nvbd/pw0un2hOfd7BdxpH4rZqtH6l0zU3PajpurabjectyGxSb0fAL388Dn//cuk0\nUXIcif122aYD/bTebA3X75U1IQKBgQD4M0SMzGGDSM3rvrL47U3REbZi+SEVsUtK\n/0fJhBAzOcaJ+QsB6s39FuU/GoGox/0647OQ58BJEV87D1OTMUSNmIErMRBZ+teh\n379MoF/E3HwwqnvHMdPT6KEJc1DuSQJAE46TtKBLI1E68NFI+GWUIqizNij8g8DE\nxCbQYZDtnQKBgQDlgnp2WQEQwTpE6cuuF6HQxIWcCv8xytCIPlPSLA2TvzjDx6UT\ndaKGVL1nSajU+EUYueVC5FhjMxYH1TfGLwCJC9lMnBguBEFAopG33nrGAkXAiURt\nVPCV/SyL3LNmn/CQmGSRUX8xUHPPi4Y5rNBoUDpxyPfJifMIJvfU5OpnoQKBgHcd\nQ5y+yV2UJx8oWgQU/u+DLKC7JyGeAKBxeyY+9vdeluXIW3wED++SeVqbgfZaZDFK\n7fZxAlsOt0FEzbsqPdgmhHcSCOLl+254TvqbTNrRZdiFOPoT87ETR5Wdfg2dhDL0\nL8y7NuJYhLdgs0/txfId4BCBwZHOOUk1Sdtft4oRAoGBAM33b1SVhJJUaei5aZ8P\nvIifO3w1jZjwGqi5uWOuaIGHtixJP3DizZb33NP3loKuTTWD6dm1eU8hlFS9ZD2e\n4KpgUo8erFy4evZOSQpGhqtYMUg2XeoNFvKoqYdfLAxmI4KEZ1roANYhGul9k2lk\naLRLMpL8SfMsQ9qpeZi962ak\n-----END PRIVATE KEY-----\n",
	  "client_email": "firebase-adminsdk-eb6se@newflixbot.iam.gserviceaccount.com",
	  "client_id": "106794844845020968546",
	  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	  "token_uri": "https://oauth2.googleapis.com/token",
	  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-eb6se%40newflixbot.iam.gserviceaccount.com",
	  "universe_domain": "googleapis.com"
	}
	`))

	app, err := firebase.NewApp(context.Background(), &config, opt)
	if err != nil {
		utils.Flog.Error(err.Error())
		return fmt.Errorf("error initializing app: %v", err)
	}
	
	firestore, err:= app.Firestore(context.Background())
	if err != nil {
		utils.Flog.Error(err.Error())
		return fmt.Errorf("error creating firestore client: %v", err)
	}

	ProxyDNSFirestore = firestore
	utils.Flog.Info("Connected to firestore")
	fmt.Println("Connected to firestore")
	return nil 
}


func AddLog(ip string, log string) {
	go func() {
		// Get the current timestamp
		currentTime := time.Now()
		timestamp := currentTime.Format("2006-01-02 15:04:05")

		// Append timestamp to the log
		logWithTimestamp := fmt.Sprintf("%s: %s", timestamp, log)

		// Run a transaction
		err := ProxyDNSFirestore.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
			// Check if the document exists
			docRef := ProxyDNSFirestore.Collection("proxydnslogs").Doc(ip)
			docSnapshot, err := tx.Get(docRef)
			if err != nil && status.Code(err) != codes.NotFound {
				utils.Flog.Error(err.Error())
				return fmt.Errorf("error checking document existence: %v", err)
			}

			if !docSnapshot.Exists() {
				// Document doesn't exist, create it
				return tx.Set(docRef, map[string]interface{}{"logs": []string{logWithTimestamp}})
			}

			// Document exists, update the array field
			return tx.Update(docRef, []firestore.Update{
				{Path: "logs", Value: firestore.ArrayUnion(logWithTimestamp)},
			})
		})

		if err != nil {
			utils.Flog.Error(err.Error())
			// fmt.Printf("transaction failed: %v", err)
			return
		}

		// fmt.Println("transaction succeeded: document updated successfully")
	}()
}

