package main

var (
	// Known Payload headers
	connectHeader   = PayloadHeader{Type: "CONNECT"}
	closeHeader     = PayloadHeader{Type: "CLOSE"}
	getStatusHeader = PayloadHeader{Type: "GET_STATUS"}
	pongHeader      = PayloadHeader{Type: "PONG"}       // Response to PING payload
	launchHeader    = PayloadHeader{Type: "LAUNCH"}     // Launches a new chromecast app
	playHeader      = PayloadHeader{Type: "PLAY"}       // Plays / unpauses the running app
	pauseHeader     = PayloadHeader{Type: "PAUSE"}      // Pauses the running app
	seekHeader      = PayloadHeader{Type: "SEEK"}       // Pauses the running app
	volumeHeader    = PayloadHeader{Type: "SET_VOLUME"} // Sets the volume
	loadHeader      = PayloadHeader{Type: "LOAD"}       // Loads an application onto the chromecast
)

type Payload interface {
	SetRequestId(id int)
}

type PayloadHeader struct {
	Type      string `json:"type"`
	RequestId int    `json:"requestId,omitempty"`
}

func (p *PayloadHeader) SetRequestId(id int) {
	p.RequestId = id
}

type MediaHeader struct {
	PayloadHeader
	MediaSessionId int     `json:"mediaSessionId"`
	CurrentTime    float32 `json:"currentTime"`
	ResumeState    string  `json:"resumeState"`
}

type Volume struct {
	Level float32 `json:"level"`
	Muted bool    `json:"muted"`
}

type ReceiverStatusResponse struct {
	PayloadHeader
	Status struct {
		Applications []struct {
			AppId        string `json:"appId"`
			DisplayName  string `json:"displayName"`
			IsIdleScreen bool   `json:"isIdleScreen"`
			SessionId    string `json:"sessionId"`
			StatusText   string `json:"statusText"`
			TransportId  string `json:"transportId"`
		} `json:"applications"`

		Volume Volume `json:"volume"`
	} `json:"status"`
}

type Application struct {
	AppId        string `json:"appId"`
	DisplayName  string `json:"displayName"`
	IsIdleScreen bool   `json:"isIdleScreen"`
	SessionId    string `json:"sessionId"`
	StatusText   string `json:"statusText"`
	TransportId  string `json:"transportId"`
}

type ReceiverStatusRequest struct {
	PayloadHeader
	Applications []Application `json:"applications"`

	Volume Volume `json:"volume"`
}

type LaunchRequest struct {
	PayloadHeader
	AppId string `json:"appId"`
}

type LoadMediaCommand struct {
	PayloadHeader
	Media       MediaItem   `json:"media"`
	CurrentTime int         `json:"currentTime"`
	Autoplay    bool        `json:"autoplay"`
	CustomData  interface{} `json:"customData"`
}

type MediaItem struct {
	ContentId   string  `json:"contentId"`
	ContentType string  `json:"contentType"`
	StreamType  string  `json:"streamType"`
	Duration    float32 `json:"duration"`
	Metadata    struct {
		MetadataType int    `json:"metadataType`
		Title        string `json:"title"`
		SongName     string `json:"songName"`
		Artist       string `json:"artist"`
	} `json:"metadata"`
}

type Media struct {
	MediaSessionId int     `json:"mediaSessionId"`
	PlayerState    string  `json:"playerState"`
	CurrentTime    float32 `json:"currentTime"`
	IdleReason     string  `json:"idleReason"`
	Volume         Volume  `json:"volume"`

	Media MediaItem `json:"media"`
}

type MediaStatusResponse struct {
	PayloadHeader
	Status []Media `json:"status"`
}
