package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)




func streamHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	vid := p.ByName("vid-id")
	vl := VIDEO_DIR +vid
	video,err := os.Open(vl) //打开文件二进制
	if err != nil{
		sendErrorResponse(w,http.StatusInternalServerError,"internal erro")
		return
	}
	w.Header().Set("Content-Type","video/mp4") //定义播放协议头浏览器自动解析mp4

	http.ServeContent(w,r,"",time.Now(),video) //浏览器播放
	defer video.Close()


}


func uploadHandler(w *http.ResponseWriter,r *http.Request, p httprouter.Params)  {



}


