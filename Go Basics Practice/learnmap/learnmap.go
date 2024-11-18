package learnmap

import "fmt"


type userScore map[string]float64

func (user userScore) outputScore() {
	for key, value := range user {
		fmt.Println("Key is ", key, " Value is ", value)
	}
}

func Learnmap() {

	// Basic ways to create a slice
	var usernames [2]string
	usernames[0] = "abc123"
	usernames[1] = "def123"

	// Basic ways to create a map
	websites := map[string]string {}
	websites["google"] = "google.com"
	websites["aws"] = "aws.com"

	// Using make function
	users := make([]string, 2, 5)
	users[0] = "Suraj"
	users = append(users, "Allen")
	users = append(users, "Job")

	fmt.Println(users)

	userDetails := make(map[string]string, 3)
	userDetails["Suraj"] = "Developer"
	userDetails["Allen"] = "Tester"
	userDetails["Job"] = "Analyst"

	// Alias in Make function
	userScoreDetails := make(userScore, 3)
	userScoreDetails["Suraj"] = 20.22
	userScoreDetails["Allen"] = 24.22
	userScoreDetails["Job"] = 90.22

	userScoreDetails.outputScore()


}