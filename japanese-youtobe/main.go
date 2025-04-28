package main

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Video struct {
	Title   string `bson:"title"`
	Link    string `bson:"link"`
	VideoId string `bson:"video_id"`
}

func SaveVideosFromYoutobe(ctx context.Context, collection *mongo.Collection, url string) error {
	cmd := exec.Command("yt-dlp", "--flat-playlist", url, "--get-title", "--get-id")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)

	var allVideos []Video
	for scanner.Scan() {
		title := scanner.Text()
		if !scanner.Scan() {
			break
		}
		id := scanner.Text()
		allVideos = append(allVideos, Video{
			Title:   title,
			Link:    fmt.Sprintf("https://www.youtobe.com/watch?v=%s", id),
			VideoId: id,
		})
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	// Prepare a list of all VideoIds to check existence in MongoDB
	videoIds := make([]string, len(allVideos))
	for i, v := range allVideos {
		videoIds[i] = v.VideoId
	}

	// Query existing VideoIds in one request
	cursor, err := collection.Find(ctx, bson.M{"video_id": bson.M{"$in": videoIds}})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	existing := make(map[string]struct{})
	for cursor.Next(ctx) {
		var v Video
		if err := cursor.Decode(&v); err == nil {
			existing[v.VideoId] = struct{}{}
		}
	}

	// Filter out videos that already exist
	var results []interface{}
	for _, v := range allVideos {
		if _, found := existing[v.VideoId]; !found {
			results = append(results, v)
		}
	}

	if len(results) > 0 {
		_, err := collection.InsertMany(ctx, results)
		if err != nil {
			return err
		}
		log.Printf("Inserted %d videos into MongoDB.\n", len(results))
	} else {
		log.Println("No new videos found.")
	}
	return nil
}

func homeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome Home!")
}

var collection *mongo.Collection

func main() {

	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	//log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Hello World!")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/")
	client, _ := mongo.Connect(ctx, clientOptions)
	collection = client.Database("youtube_japanese").Collection("rikki")

	err := SaveVideosFromYoutobe(ctx, collection, "https://www.youtube.com/@congongazura5123/videos")
	if err != nil {
		log.Fatal(err)
	}

}
