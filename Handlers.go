package main
import (
	"net/http"
	"fmt"
	"runtime"
	"bytes"
	"encoding/json"
	"crypto/aes"
	"io"
	"crypto/cipher"
	"errors"
	"crypto/rand"
	"encoding/base64"
)

type Website struct {
	ID		int
	Url     	string
	UserName     	string
	Password 	string
}

func (w Website) ToString() string {
	return w.Url+" "+w.UserName+" "+w.Password
}

func getRequestContentFromRequest(req *http.Request) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	data := buf.Bytes()
	var requestContent map[string]interface{}
	err := json.Unmarshal(data, &requestContent)
	if err != nil {
		fmt.Println(err)
	}
	return requestContent
}

func encrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func index_func(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	url := requestContent["url"].(string)
	username := requestContent["username"].(string)
	password := requestContent["password"].(string)
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		_, err := db.Exec("INSERT INTO website("+
			"url, "+
			"username, "+
			"password) "+
			"VALUES(?, ?, ?);",
			url,
			username,
			password,
			)
		if err != nil {
			fmt.Println("Can't insert data into Database: "+err.Error())
			ciphertext, err := encrypt([]byte("Can not insert data into Database"))
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
			rw.Write(ciphertext)
		}else {
			ciphertext, err := encrypt([]byte("Everything worked fine!"))
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
			rw.Write(ciphertext)
		}

	}else{
		ciphertext, err := encrypt([]byte("not enough parameters given"))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	}
}
func get_all_function(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	url := requestContent["url"].(string)

	//TODO unencrypt the the encrypted url
	unencryptedUrl := requestContent["urlCrypted"].(string)
	encryptedurl, err := decrypt([]byte(unencryptedUrl))
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Println(encryptedurl)


	if len(url) != 0 {
		var returnString string = ""
		rows, err := db.Query("SELECT * FROM website WHERE url like ?", url)
		if (err != nil) {
			fmt.Println("can't execute select query: "+err.Error())
		}
		for rows.Next() {
			w := new(Website)
			err := rows.Scan(&w.ID, &w.Url, &w.UserName, &w.Password)
			if (err != nil) {
				fmt.Println("can't read into struct: "+err.Error())
			}
			returnString = returnString + w.ToString()
		}
		err = rows.Err()
		if(err != nil){
			fmt.Println("Error with Row: "+err.Error())
		}

		ciphertext, err := encrypt([]byte(returnString))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	}else{
		ciphertext, err := encrypt([]byte("no get parameter given"))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	}
}