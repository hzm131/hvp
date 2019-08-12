package video

import (
	"com/models/servser_model"
	"com/models/servser_model/users"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Name         string  `gorm:"column:name"json:"name"` //片名
	Pid          int     `gorm:"column:pid"json:"pid"`
	Origin       string  `gorm:"column:origin"json:"origin"`             //产地
	Duration     string  `gorm:"column:duration"json:"duration"`         //时长
	Language     string  `gorm:"column:language"json:"language"`         //语种
	Years        string  `gorm:"column:years"json:"years"`               //年份
	Score        float32 `gorm:"column:score"json:"score"`               //评分
	Introduction string  `gorm:"column:introduction"json:"introduction"` //简介
	Category     string  `gorm:"column:category"json:"category"`         //类别

	VideoSrcId int      `gorm:"column:video_src_id;"json:"video_src_id"validate:"required||integer"` //视频路径
	VideoSrc   VideoSrc `gorm:"ForeignKey:VideoSrcId:AssociationForeignKey:ID"`
	ImageSrcId int      `gorm:"column:image_src_id;"json:"image_src_id"validate:"required||integer"` //封面路径
	ImageSrc   ImageSrc `gorm:"ForeignKey:ImageSrcId:AssociationForeignKey:ID"`
}

//上传视频封面
type ImageSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path"json:"src_path"`
}

// 上传视频路径
type VideoSrc struct {
	gorm.Model
	SrcPath string `gorm:"column:src_path"json:"src_path"`
}

//评论表
type Comment struct {
	gorm.Model
	VideoId int         `gorm:"column:video_id"json:"video_id"` //视频id
	Video   Video       `gorm:"ForeignKey:VideoId:AssociationForeignKey:ID"`
	Content string      `gorm:"column:content"json:"content"` //评论内容
	UserId  int         `gorm:"column:user_id"json:"user_id"` //评论人的id
	User    users.Users `gorm:"ForeignKey:UserId:AssociationForeignKey:ID"`
}

//回复表
type Reply struct {
	gorm.Model
	CommentId int     `gorm:"column:comment_id"json:"comment_id"` //通过评论id可以知道自己属于哪条评论
	Comment   Comment `gorm:"ForeignKey:CommentId:AssociationForeignKey:ID"`
	ReplyId   int     `gorm:"column:reply_id"json:"reply_id"` //回复目标id
	Content   string  `gorm:"column:content"json:"content"`   //回复目标内容
	UserId    int     `gorm:"column:user_id"json:"user_id"`   //回复用户id
}

func (this *VideoSrc) CreatedVideoSrc() (int, error) {
	fmt.Printf("this", this)
	path := VideoSrc{}
	find := servser_model.Db.Create(this).Scan(&path)
	if err := find.Error; err != nil {
		fmt.Println("创建失败", err)
		return 0, err
	}
	//创建成功后返回id
	id := int(path.ID) //拿到id
	return id, nil
}

func (this *ImageSrc) CreatedImageSrc() (int, error) {
	fmt.Printf("this", this)
	path := ImageSrc{}
	find := servser_model.Db.Create(this).Scan(&path)
	if err := find.Error; err != nil {
		fmt.Println("创建失败", err)
		return 0, err
	}
	//创建成功后返回id
	id := int(path.ID) //拿到id
	return id, nil
}

func (this *Video) CreatedVideo() (int, error) {
	video := Video{}
	find := servser_model.Db.Create(this).Scan(&video)
	if err := find.Error; err != nil {
		fmt.Println("创建失败", err)
		return 0, err
	} //创建成功后返回id
	id := int(video.ID) //拿到id
	return id, nil
}

type TotalVideo struct {
	Videos []Video `json:"videos"`
	Total  int     `json:"total"`
}

func (this *Video) QueryVideos(condition string, limit string, offset string) (totalVideo TotalVideo, err error) {
	cond := "%" + condition + "%"
	fmt.Println("cond", cond)
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	count := servser_model.Db.Raw("select * from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where concat(name,origin) like ?", &cond).Scan(&totalVideo.Videos).RowsAffected
	totalVideo.Total = int(count)

	query := servser_model.Db.Raw("select * from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where concat(name,origin) like ? limit ? offset ?", &cond, &limit, &offset).Scan(&totalVideo.Videos)
	if err = query.Error; err != nil {
		return
	}
	return
}

func (this *Video) FindVideo(Id string) (video Video, err error) {
	query := servser_model.Db.Raw("select * from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where video.id = ?", &Id).Scan(&video)
	if err = query.Error; err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}

func (this *Video) UpdateVideo(Id string) (err error) {
	fmt.Println("id:", Id)
	update := servser_model.Db.Exec("update video set name = ?, origin = ?, duration = ?, language = ?, years = ?, score = ?, introduction = ?, category = ?, video_src_id = ?, image_src_id = ? where id = ?", this.Name, this.Origin,
		this.Duration, this.Language, this.Years, this.Score, this.Introduction, this.Category, this.VideoSrcId, this.ImageSrcId, &Id)
	if err = update.Error; err != nil {
		return
	}
	return
}

func (this *Video) DeleteVideo(Id string) (err error) {
	delete := servser_model.Db.Exec("delete video,video_src,image_src from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where video.id = ?", &Id)
	if err = delete.Error; err != nil {
		return
	}
	return
}
