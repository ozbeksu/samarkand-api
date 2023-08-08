package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/ozbeksu/samarkand-api/api"
	"github.com/ozbeksu/samarkand-api/conf"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/store"
	log "github.com/sirupsen/logrus"
)

var (
	appConfig *conf.AppConfig
	dbConfig  *conf.DBConfig
	db        *ent.Client
)

func init() {
	var err error

	conf.LoadEnvConfig()
	appConfig = conf.NewAppConfig()
	dbConfig = conf.NewDBConfig()

	db, err = ent.Open("postgres", dbConfig.ConnectionUri())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
}

func main() {
	stores := &store.Store{
		User:      store.NewEntUserStore(db),
		Tag:       store.NewEntTagStore(db),
		Topic:     store.NewEntTopicStore(db),
		Community: store.NewEntCommunityStore(db),
		Comment:   store.NewEntCommentStore(db),
	}

	app := fiber.New(fiber.Config{ErrorHandler: api.HandleErrors})
	app.Use(cors.New(appConfig.Cors()))

	authHandler := api.NewAuthHandler(stores)
	userHandler := api.NewUserHandler(stores)
	tagHandler := api.NewTagHandler(stores)
	topicHandler := api.NewTopicHandler(stores)
	communityHandler := api.NewCommunityHandler(stores)
	commentHandler := api.NewCommentHandler(stores)

	// auth endpoint
	auth := app.Group("/api/auth")

	// auth handlers
	auth.Post("/", authHandler.HandleAuthenticate)

	// api/v1 endpoint
	apiV1 := app.Group("/api/v1")

	// users handlers
	apiV1.Post("/users", userHandler.HandlePostUsers)
	apiV1.Get("/users/:username/profile", userHandler.HandleGetUserWithProfile)
	apiV1.Get("/users/:username/posts", userHandler.HandleGetUserWithPosts)
	apiV1.Get("/users/:username/treads", userHandler.HandleGetUserWithTreads)
	apiV1.Get("/users/:username/comments", userHandler.HandleGetUserWithComments)
	apiV1.Get("/users/:username/messages", userHandler.HandleGetUserWithMessages)
	apiV1.Get("/users/:username/media", userHandler.HandleGetUserWithMedia)
	apiV1.Get("/users/:username/bookmarked", userHandler.HandleGetUserWithBookmark)
	apiV1.Get("/users/:username/up-voted", userHandler.HandleGetUserWithUpVoted)
	apiV1.Get("/users/:username/down-voted", userHandler.HandleGetUserWithDownVoted)
	apiV1.Get("/users/:username", userHandler.HandleGetUser)
	apiV1.Get("/users", userHandler.HandleGetUsers)

	// tags handlers
	apiV1.Get("/tags/:slug/comments", tagHandler.HandleGetTagWithComments)
	apiV1.Get("/tags/:slug", tagHandler.HandleGetTag)
	apiV1.Get("/tags", tagHandler.HandleGetTags)

	// topics handlers
	apiV1.Get("/topics/:slug/communities", topicHandler.HandleGetTopicWithCommunities)
	apiV1.Get("/topics/:slug", topicHandler.HandleGetTopic)
	apiV1.Get("/topics", topicHandler.HandleGetTopics)

	// communities handlers
	apiV1.Get("/communities/:slug/comments", communityHandler.HandleGetCommunityWithComments)
	apiV1.Get("/communities/:slug", communityHandler.HandleGetCommunities)
	apiV1.Get("/communities", communityHandler.HandleGetCommunities)

	// comments handlers
	apiV1.Post("/comments/:slug/users/:username/up-vote", commentHandler.HandlePostCommentUpVote)
	apiV1.Post("/comments/:slug/users/:username/down-vote", commentHandler.HandlePostCommentDownVote)
	apiV1.Post("/comments/:slug/users/:username/bookmark", commentHandler.HandlePostCommentBookmark)
	apiV1.Post("/comments", commentHandler.HandlePostComments)
	apiV1.Get("/comments/:slug/comments", commentHandler.HandleGetCommentWithComments)
	apiV1.Get("/comments/:slug", commentHandler.HandleGetComment)
	apiV1.Get("/comments", commentHandler.HandleGetComments)

	// lift-off
	err := app.Listen(appConfig.Port())
	if err != nil {
		log.Fatal(err)
	}
}
