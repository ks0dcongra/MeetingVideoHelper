package contract

const (
	Success = iota
	ERROR_Upload_Video
	ERROR_Download_Video
	ERROR_Download_Video_NotExist
	ERROR_Video_To_Hex
	ERROR_Upload_Video_NotExist
	ERROR_Video_Format
	ERROR_IP_Request_Limit
	ERROR_Upload_Video_FileTooLarge
)

var Message = map[int]string{
	Success:                         "Upload video successfully.",
	ERROR_Upload_Video:              "Upload video fail.",
	ERROR_Download_Video:            "Download video fail.",
	ERROR_Download_Video_NotExist:   "VideoID not found.",
	ERROR_Video_To_Hex:              "Video bytes to hex fail.",
	ERROR_Upload_Video_NotExist:     "Video not found.",
	ERROR_Video_Format:              "Video format has some problem.",
	ERROR_IP_Request_Limit:          "Request limit reached.",
	ERROR_Upload_Video_FileTooLarge: "upload video too large.",
}
