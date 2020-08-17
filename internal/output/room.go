package output

import "github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"

type RoomResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Sender   string `json:"sender"`
	Text     string `json:"text"`
	ImageURL string `json:"image_url"`
}

func ToGetRoomResponse(room *model.Room) RoomResponse {
	var res RoomResponse
	res.ID = room.ID
	res.Name = room.Name
	for _, v := range room.Comments {
		comment := Comment{
			Sender:   v.Sender,
			Text:     v.Text,
			ImageURL: v.ImageURL,
		}
		res.Comments = append(res.Comments, comment)
	}
	return res
}
