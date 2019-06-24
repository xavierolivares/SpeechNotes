Goal: Learning how to work with golang, google cloud api, and speech recognition

RESOURCE -
https://medium.com/@ODSC/quick-start-for-golang-google-cloud-api-and-speech-recognition-f157cda1ce5c

CREATING A SERVICE ACCOUNT -
Setting up a Google Cloud API account: https://cloud.google.com/docs/authentication/getting-started#auth-cloud-implicit-go 

SECURITY -
Managing service account keys: https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en_US&_ga=2.101743105.-1015558750.1559092652#managing_service_account_keys

Securing keys with GODOTENV - 
https://github.com/joho/godotenv

Still unclear:
Is my secret configuration working?
How to get an audio file?
How to import that file into application?

Get the Client Library Tutorial: https://github.com/Drizzy3D/go-speech-recognition-lib

TESTING RESULT 1:
-2019/06/21 12:54:47 Failed to recognize: rpc error: code = PermissionDenied desc = Cloud Speech-to-Text API has not been used in project "soandso" before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/speech.googleapis.com/overview?project="soandso" then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
-Enabled API, needed to set up billing account. Free 12months or $300 credit!

TESTING RESULT 2: 
-Passed! "how old is the Brooklyn Bridge (confidence=0.983505)"

NEW GOAL:
Next step is to figure out to make these changes in realtime..I'm thinking that I need to add to the audio file based on speech. It may get delayed as the audiofile changes. Perhaps I can parse a new audiofile every 10 seconds? Not sure if concurrency will help.

Resources: https://cloud.google.com/speech-to-text/docs/streaming-recognize

gcloud tool: https://cloud.google.com/speech-to-text/docs/quickstart


TESTING RESULT 1: 
2019/06/21 13:39:53 Warning: Speech recognition request exceeded limit of 60 seconds.
2019/06/21 13:39:53 Could not recognize: code:11 message:"Audio Timeout Error: Long duration elapsed without audio. Audio should be sent close to real time." 
-I don't think the program is recognizing the audio from my mic.


Google Cloud SDK Documentation: 
https://cloud.google.com/sdk/docs/
https://cloud.google.com/speech-to-text/docs/quickstart-gcloud

Installation for GO: https://cloud.google.com/appengine/docs/standard/go/download

Steps: Needed to gcloud init and authorize Google Cloud SDK
Confirmation page: https://cloud.google.com/sdk/auth_success

GOAL: test gcloud and code locally

TEST 1: ERROR
PASSED IN - /Users/xavierolivares/go/src/projects/speechtest/audio.raw
RESULT - 2019/06/21 15:40:22 Please pass path to y
our local audio file as a command line ar
gument


TEST 2: SUCCESS
PASSED IN - go build && ./speechtest /Users/xav
ierolivares/go/src/projects/speechtest/au
dio.raw
RESULT - Result: alternatives:<transcript:"how old
 is the Brooklyn Bridge" confidence:0.983
5046 > is_final:true result_end_time:<sec
onds:1 nanos:770000000 > 

FINDING MIC CODES: gst-device-monitor-1.0

REQUIRED: Need to insert mic input into command line with gst-launch1.0.

FOUND this: gst-launch-1.0 osxaudiosrc device=40

Questions:
thirdlevel branch: how can i get the microphone to link up with this code?

stretch:
Is it possible to get results to render on html page?

TAKEAWAYS: how to run audio file locally with cat and livecaption
cat ./audio.raw | go run liv
ecaption.go

//Configure gcloud credentials for communicating with SpeechToText API

export GOOGLE_APPLICATION_CREDENTIALS=[PATH TO CREDENTIALS JSON]

gst-launch-1.0 -v osxaudiosrc device=40 ! audioconvert ! audioresample ! audio/x-raw,channels=1,rate=16000 ! filesink location=/dev/stdout | go run livecaption.go

gst-launch-1.0 -v osxaudiosrc device=40 ! audioconvert ! audioresample ! audio/x-raw,channels=1,rate=16000 ! filesink location=/dev/stdout | go run *.go

Running local file:
cat ./audio.raw | go run *.go

Doesn't work with bluetooth, at least yet.