package cursedfoxbot

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

// ** SECRETS **
var deepAI_key string = os.Getenv("DEEPAI_KEY")
var api_key string = os.Getenv("TWITTER_APIKEY")
var api_key_secret string = os.Getenv("TWITTER_APISECRET")
var user_token string = os.Getenv("TWUSER_APITOKEN")
var user_secret string = os.Getenv("TWUSER_APISECRET")

var oConfig oauth1.Config = oauth1.Config{
	ConsumerKey:    api_key,
	ConsumerSecret: api_key_secret,
	CallbackURL:    "https://twitter.com/",
	Endpoint:       twitter.AuthorizeEndpoint,
}

var client *http.Client = new(http.Client)

type DeepAIResp struct {
	Id         string
	Output_url string
}

func getDeepAIImage(query string) string {
	u, _ := url.Parse("https://api.deepai.org/api/text2img")

	f := strings.NewReader("text=" + query)
	req, err := http.NewRequest("POST", u.String(), f)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("api-key", deepAI_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	j, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()

	d := new(DeepAIResp)
	json.Unmarshal(j, &d)

	return (d.Output_url)
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func getHTTPBody(u string) []byte {
	res, err := client.Get(u)
	if err != nil {
		fmt.Println(err)
	}

	b, _ := ioutil.ReadAll(res.Body)
	return b
}

// This will get the token and secret that allows you to use the fox bot.
// Be careful!

func GetTwitterAuth() {
	reqToken, reqSecret, err := oConfig.RequestToken()
	if err != nil {
		fmt.Println(err)
	}

	authURL, err := oConfig.AuthorizationURL(reqToken)

	fmt.Println("OPEN THIS WEBSITE IN YOUR BROWSER, AND ACCEPT PERMISSIONS!")
	fmt.Println(authURL)
	fmt.Println("Afterwards, get oauth_token and oauth_verifier and paste them here.")
	fmt.Printf("OAuth token: ")
	reqToken = readInput()
	fmt.Printf("OAuth verifier: ")
	oauthVerifier := readInput()

	userToken, userSecret, err := oConfig.AccessToken(reqToken, reqSecret, oauthVerifier)
	fmt.Println("Here is your token/secret. Use this with the bot by setting environment variables TWUSER_TOKEN and TWUSER_SECRET in order to post photos. Enjoy!")
	fmt.Printf("TOKEN: %s\n", userToken)
	fmt.Printf("SECRET: %s\n", userSecret)
}

func authTwitter(token, secret string) *http.Client {
	oToken := oauth1.NewToken(token, secret)
	return oConfig.Client(oauth1.NoContext, oToken)
}

type TwitterMedia struct {
	Media_ID        int64
	Media_ID_String string
}

func uploadToTwitter(client *http.Client, media []byte) *TwitterMedia {
	b := bytes.NewBuffer(nil)
	m := multipart.NewWriter(b)
	mf, _ := m.CreateFormField("media")
	mf.Write(media)
	m.Close()

	req, _ := http.NewRequest("POST", "https://upload.twitter.com/1.1/media/upload.json", b)
	req.Header.Add("Content-Type", m.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	j, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(j))

	mID := new(TwitterMedia)
	json.Unmarshal(j, &mID)

	return mID
}

func tweetImage(client *http.Client, media *TwitterMedia) {

	// MULTIPART MIME
	/*
		b := bytes.NewBuffer(nil)
		m := multipart.NewWriter(b)

		m.WriteField("media_ids", media.Media_ID_String)
		m.WriteField("status", "test, again again again again again")

		m.Close()
	*/

	// FORM-STYLE

	f := url.Values{}
	f.Add("media_ids", media.Media_ID_String)

	req, _ := http.NewRequest("POST", "https://api.twitter.com/1.1/statuses/update.json", strings.NewReader(f.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	j, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(j))
}

func MakeCursedFox() {
	image := getHTTPBody(getDeepAIImage("red fox"))
	client = authTwitter(user_token, user_secret)
	media := uploadToTwitter(client, image)

	tweetImage(client, media)
}
