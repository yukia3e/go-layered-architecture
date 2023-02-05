package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func InitFirebaseApp(ctx context.Context) (
	*firebase.App,
	*auth.Client,
	error,
) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, nil, err
	}

	return app, auth, nil
}
