package video

import (
	"com/models/wx"
	"fmt"
	"os"
)

type Video struct {
	wx.Model
	Name         *string  `gorm:"column:name"json:"name"` //片名
	Pid          *int     `gorm:"column:pid"json:"pid"`
	Origin       *string  `gorm:"column:origin"json:"origin"`             //产地
	Duration     *string  `gorm:"column:duration"json:"duration"`         //时长
	Language     *string  `gorm:"column:language"json:"language"`         //语种
	Years        *string  `gorm:"column:years"json:"years"`               //年份
	Score        *float32 `gorm:"column:score;type:float;"json:"score"`   //评分
	Introduction *string  `gorm:"column:introduction"json:"introduction"` //简介
	Category     *string  `gorm:"column:category"json:"category"`         //类别

	VideoSrcId int      `gorm:"column:video_src_id;not null"json:"video_src_id"validate:"required||integer"` //视频路径
	VideoSrc   VideoSrc `gorm:"ForeignKey:VideoSrcI;AssociationForeignKey:ID"json:"video_src"`
	ImageSrcId int      `gorm:"column:image_src_id;not null"json:"image_src_id"validate:"required||integer"` //封面路径
	ImageSrc   ImageSrc `gorm:"ForeignKey:ImageSrcId;AssociationForeignKey:ID"json:"image_src"`

	Director string `gorm:"column:director"json:"director"` //导演
	Actor string `gorm:"column:actor"json:"actor"`

	Count *int `gorm:"column:count"json:"count"` //播放量

}

//上传视频封面
type ImageSrc struct {
	wx.Model
	Name    *string `gorm:"column:name"json:"name"`
	SrcPath *string `gorm:"column:src_path"json:"src_path"`
}

func (this *ImageSrc) DeleteImageSrc(id string) (err error) {
	image_src := ImageSrc{}
	query := wx.Db.Raw("select * from image_src where id = ?", &id).Scan(&image_src)
	if err = query.Error; err != nil {
		return err
	}
	err = os.Remove("./public/upload/images/" + *image_src.Name)
	if err != nil {
		return
	}
	delete := wx.Db.Exec("delete from image_src where id = ?", &id)
	if err = delete.Error; err != nil {
		return err
	}
	return
}

// 上传视频路径
type VideoSrc struct {
	wx.Model
	Name    *string `gorm:"column:name"json:"name"`
	SrcPath *string `gorm:"column:src_path"json:"src_path"`
}

func (this *VideoSrc) DeleteVideoSrc(id string) (err error) {
	video_src := VideoSrc{}
	query := wx.Db.Raw("select * from video_src where id = ?", &id).Scan(&video_src)
	if err = query.Error; err != nil {
		return err
	}
	err = os.Remove("./public/upload/videos/" + *video_src.Name)
	if err != nil {
		return
	}
	delete := wx.Db.Exec("delete from video_src where id = ?", &id)
	if err = delete.Error; err != nil {
		return err
	}
	return
}

func (this *VideoSrc) CreatedVideoSrc() (int, error) {
	fmt.Printf("this", this)
	path := VideoSrc{}
	find := wx.Db.Create(this).Scan(&path)
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
	find := wx.Db.Create(this).Scan(&path)
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
	find := wx.Db.Create(this).Scan(&video)
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

func (this *Video) QueryVideos(condition string, orderBy string, limit string, offset string) (totalVideo TotalVideo, err error) {
	cond := "%" + condition + "%"
	fmt.Println("cond", cond)
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	count := wx.Db.Raw("select * from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where concat(video.name,video.origin,video.years,video.score,video.duration,video.category) like ? order by ? Desc", &cond, &orderBy).Scan(&totalVideo.Videos).RowsAffected
	totalVideo.Total = int(count)
	totalVideo.Videos = nil
	rows, err := wx.Db.Raw("select video.id,video.name,pid,origin,duration,language,years,score,introduction,category,director,actor,video_src_id,image_src_id,image_src.id,image_src.name,image_src.src_path,video_src.id,video_src.name,video_src.src_path from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where concat(video.name,video.origin,video.years,video.score,video.duration,video.category) like ? order by ? Desc limit ? offset ?", &cond, &orderBy, &limit, &offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		video := Video{}
		err = rows.Scan(&video.ID, &video.Name, &video.Pid, &video.Origin, &video.Duration, &video.Language, &video.Years, &video.Score, &video.Introduction, &video.Category,&video.Director,&video.Actor, &video.VideoSrcId, &video.ImageSrcId, &video.ImageSrc.ID, &video.ImageSrc.Name, &video.ImageSrc.SrcPath, &video.VideoSrc.ID, &video.VideoSrc.Name, &video.VideoSrc.SrcPath)
		if err != nil {
			return totalVideo, err
		}
		totalVideo.Videos = append(totalVideo.Videos, video)
	}
	return
}

func (this *Video) FindVideo(Id string) (video Video, err error) {
	query := wx.Db.Raw("select * from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where video.id = ?", &Id).Scan(&video)
	if err = query.Error; err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}

func (this *Video) UpdateVideo(Id string) (err error) {
	fmt.Println("id:", Id)
	update := wx.Db.Exec("update video set name = ?, origin = ?, duration = ?, language = ?, years = ?, score = ?, introduction = ?, category = ?,director = ?,actor = ?, video_src_id = ?, image_src_id = ? where id = ?", this.Name, this.Origin,
		this.Duration, this.Language, this.Years, this.Score, this.Introduction, this.Category,this.Director,this.Actor, this.VideoSrcId, this.ImageSrcId, &Id)
	if err = update.Error; err != nil {
		return
	}
	return
}

func (this *Video) DeleteVideo(Id string) (err error) {
	fmt.Println("Id", Id)
	delete := wx.Db.Exec("delete video,video_src,image_src from video left join video_src on video.video_src_id = video_src.id left join image_src on video.image_src_id = image_src.id where video.id = ?", &Id)
	if err = delete.Error; err != nil {
		return
	}
	return
}
