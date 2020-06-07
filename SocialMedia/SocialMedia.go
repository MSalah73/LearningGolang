// SocialMedia Package implements common fucntionality needed for the social media web applications.
package SocialMedia

import "time"

//go:generate stringer -type=MoodState
type MoodState int

// Here we define the possiable mood states using an
// iota enumerator
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
	MoodStateHighAF
)

// This is a type we embed into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// This is the type that represents a Scocial Media Post
type Post struct {
	AuditableContent // Embedded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood state
var Moods map[string]MoodState

// The init() function is responsiable for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine, "highAF": MoodStateHighAF}
}

// This is the function responsiable for creating a new Social Media Post.
// Note: Likers are add later so initailly it should be zero do we dont have to add it
func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContect := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI,
		ThumbnailURI: thumbnailURI, Keywords: keywords, AuthorMood: mood, AuditableContent: auditableContect}
}
