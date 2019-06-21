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
