package lanyard

// I'm not entirely sure if these are correct. You can review and open a PR before using it.
// Go port of https://github.com/barbarbar338/react-use-lanyard/blob/main/src/types.ts
type LanyardResponse struct {
	Success bool          `json:"success"`
	Data    *LanyardData  `json:"data"`
	Error   *LanyardError `json:"error"`
}

type LanyardWSResponse struct {
	Op  int          `json:"op"`
	Seq int          `json:"seq"`
	T   string       `json:"t"`
	D   *LanyardData `json:"d"`
}

type LanyardData struct {
	Spotify                Spotify     `json:"spotify"`
	ListeningToSpotify     bool        `json:"listening_to_spotify"`
	DiscordUser            DiscordUser `json:"discord_user"`
	DiscordStatus          string      `json:"discord_status"`
	Activities             []Activity  `json:"activities"`
	ActiveOnDiscordMobile  bool        `json:"active_on_discord_mobile"`
	ActiveOnDiscordDesktop bool        `json:"active_on_discord_desktop"`
}

type LanyardError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type Spotify struct {
	TrackId     string     `json:"track_id"`
	Timestamps  Timestamps `json:"timestamps"`
	Song        string     `json:"song"`
	Artist      string     `json:"artist"`
	Album       string     `json:"album"`
	AlbumArtUrl string     `json:"album_art_url"`
}

type Timestamps struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Activity struct {
	Type          int         `json:"type"`
	State         string      `json:"state"`
	Name          string      `json:"name"`
	Id            string      `json:"id"`
	Emoji         *Emoji      `json:"emoji"`
	CreatedAt     int         `json:"created_at"`
	ApplicationId string      `json:"application_id"`
	Timestamps    *Timestamps `json:"timestamps"`
	SessionId     string      `json:"session_id"`
	Details       *string     `json:"details"`
	Buttons       *[]string   `json:"buttons"`
	Assets        *Assets     `json:"assets"`
}

type Assets struct {
	SmallText  string `json:"small_text"`
	SmallImage string `json:"small_image"`
	LargeText  string `json:"large_text"`
	LargeImage string `json:"large_image"`
}

type Emoji struct {
	Name string `json:"name"`
}

type DiscordUser struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}
