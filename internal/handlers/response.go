package handlers

type Response struct {
	FileCode        string `json:"file_code"`
	DownloadMessage string `json:"download_message"`
	LoginMessage    string `json:"login_message"`
}
