package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("key").(string)
	log.Println(token)
}
