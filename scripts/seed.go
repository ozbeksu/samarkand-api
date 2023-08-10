package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/migrate"
	"github.com/ozbeksu/samarkand-api/utils"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	err        error
	db         *ent.Client
	faker      *gofakeit.Faker
	ctx        context.Context
	dimensions = []string{"1x1", "3x2", "16x9"}
)

func init() {
	ctx = context.Background()
	setupFaker()
	setupClient()
}

func setupFaker() {
	faker = gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)
}

func setupClient() {
	db, err = ent.Open("postgres", "host=localhost port=5432 user=admin dbname=samarkand password=secret sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	err = db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func main() {
	seed()
}

func seed() {
	users := makeUsers(10)
	addFollowers(users, 10)

	topics := makeTopics(6)
	tags := makeTags(10)
	communities := makeCommunities(5, len(topics), len(users))
	treads := makeComments(10, len(users), len(communities), len(tags), 0, 10)
	comments := makeSubComments(10, len(users), len(communities), len(tags), 1, len(treads))
	subComments := makeSubComments(20, len(users), len(communities), len(tags), len(treads), len(treads)+len(comments))

	makeMessages(10, len(users), len(communities))

	fmt.Printf("%d done", len(comments)+len(treads)+len(subComments))
}

func getScores(createdAt time.Time) (float64, float64) {
	upVotes := faker.RandomInt(makeRange(3000, 9000))
	downVotes := faker.RandomInt(makeRange(1000, 4000))

	hotScore := utils.HotScore(upVotes, downVotes, createdAt)
	bestScore := utils.BestScore(upVotes, upVotes+downVotes)

	return hotScore, bestScore
}

func getRandIntInRange(min, max int) int {
	nums := makeRange(min, max)

	return nums[rand.Intn(len(nums))]
}

func getRandBool() bool {
	return rand.Intn(2) == 0
}

func getRandDate() time.Time {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 6, 31, 0, 0, 0, 0, time.UTC)

	return faker.DateRange(startDate, endDate)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func getImageDimensions(url string) (int, int) {
	var w, h int
	if strings.Contains(url, "1x1") {
		w, h = 1000, 1050
	}
	if strings.Contains(url, "3x2") {
		w, h = 800, 392
	}
	if strings.Contains(url, "16x9") {
		w, h = 600, 338
	}
	return w, h
}

func makeRandImgSrc(n int) (string, int, int) {
	var f string
	d := dimensions[rand.Intn(len(dimensions))]
	i := getRandIntInRange(1, n)

	if i < 10 {
		f = fmt.Sprintf("/assets/posts/%s/0%d.jpg", d, i)
	} else {
		f = fmt.Sprintf("/assets/posts/%s/%d.jpg", d, i)
	}

	w, h := getImageDimensions(f)
	return f, w, h
}

func sliceContains[T comparable](str T, list []T) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
