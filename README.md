# SpeechNotes

SpeechNotes is a note taking application and assistant for those that love to attend meetings. The application converts the spoken word to text and stores each session for future review. The base of this application contains Golang and the Google Cloud Speech-to-Text API. The Golang server uses websockets to connect to a server on the frontend, which uses Node/Express and React. The spoken word is delivered in real-time to the webpage and the application allows a user to save the notes for future review.

## System Design

![SpeechNotes](https://i.ibb.co/2hDLhjD/Speech-To-Text-Design.png)
