package firebaseapp

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var app *FirebaseApp

// FirebaseApp - 내부적으로 쓰는 Firebase 객체 담는 구조체
type FirebaseApp struct {
	App       *firebase.App
	Auth      *auth.Client
	Firestore *firestore.Client
}

// InitFirebaseApp - Firebase 앱 초기화
func InitFirebaseApp(ctx context.Context) {
	fapp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	authClient, err := fapp.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth: %v\n", err)
	}
	firestoreClient, err := fapp.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}

	app = &FirebaseApp{
		App:       fapp,
		Auth:      authClient,
		Firestore: firestoreClient,
	}
}

// App - 초기화 된 Firebase 앱 싱글톤 반환
func App() *FirebaseApp {
	return app
}
