package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/context"
	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"github.com/joho/godotenv" 
	"log" 
	"os"
)


func main() {
	fmt.Println("Hello ODSC and ME!")

	//ERROR HANDLING FOR GOOGLE_APPLICATION_CREDENTIALS
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	googleKey := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	//ERROR HANDLING IF API HAS NOT COMPLETED A REQUEST
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if(err != nil) {
	fmt.Println(err)
	}

	//ERROR HANDLING FOR CONNECTING TO AUDIO FILE
	filDir := "Recording.wav";

	audioData, err := ioutil.ReadFile(fileDir);
	if(err != nil) {
	fmt.Println(err)
	}

	response, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
	Config: &speechpb.RecognitionConfig{
	Encoding: speechpb.RecognitionConfig_LINEAR16,
	SampleRateHertz: 22050,
	LanguageCode: "en-US",
	},
	Audio: &speechpb.RecognitionAudio{
	AudioSource: &speechpb.RecognitionAudio_Content{Content: audioData},
	},
	})

	if (err != nil) {
		fmt.Println(err)
	}

	for _, result := range response.Results{
	for _, alt := range result.Alternatives{
	fmt.Println(alt.Transcript)
	
	}
	}
}
