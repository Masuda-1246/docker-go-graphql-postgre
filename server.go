package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Masuda-1246/docker-go-graphql-postgre/graph"
	"github.com/Masuda-1246/docker-go-graphql-postgre/db"
	"github.com/Masuda-1246/docker-go-graphql-postgre/loader"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//データベースへの接続処理
	db := db.ConnectGORM() //追加

	ldrs := loader.NewLoaders(db)
	// resolver内でデータベースを扱えるように設定
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB:         db, // 追加
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loader.Middleware(ldrs, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ここで.envファイル全体を読み込みます。
// この読み込み処理がないと、個々の環境変数が取得出来ません。
func loadEnv() {
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}