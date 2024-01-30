package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type Quiz struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

var quizzes []Quiz

func main() {
	r := gin.Default()

	r.GET("/quiz", getRandomQuiz)
	r.Run(":8080")
}

func getRandomQuiz(c *gin.Context) {
	if len(quizzes) == 0 {
		c.JSON(400, gin.H{"error": "No quizzes available"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quizzes))
	quiz := quizzes[randomIndex]
	c.JSON(200, quiz)
}

func init() {
	quizzes = append(quizzes, Quiz{
		ID:       "1",
		Question: "What is the Capital of France?",
		Options:  []string{"Berlin", "Paris", "London", "Rome"},
		Answer:   "Paris",
	})
	quizzes = append(quizzes, Quiz{
		ID:       "2",
		Question: "Which planet is known as the Red Planet?",
		Options:  []string{"Earth", "Mars", "Jupiter", "Venus"},
		Answer:   "Mars",
	})
	quizzes = append(quizzes, Quiz{
		ID:       "3",
		Question: "What is the largest mammal on Earth?",
		Options:  []string{"Elephant", "Blue Whale", "Giraffe", "Hippopotamus"},
		Answer:   "Blue Whale",
	})
}
