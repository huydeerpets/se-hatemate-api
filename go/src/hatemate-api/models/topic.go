package models

import (
	"github.com/astaxie/beego/orm"
)

type Topic struct {
	ID          int    `json:"topic_id"`
	TopicTitle  string `json:"topic_title"`
	Description string `json:"description"`
	CoverImage  string `json:"cover_image"`
}

type SubTopic struct {
	ID            int    `json:"subtopic_id"`
	SubTopicTitle string `json:"subtopic_title"`
	Description   string `json:"subtopic_description"`
	Image         string `json:"subtopic_image"`
	Like          int    `json:"like"`
	Comment       int    `json:"comment"`
	Topic         *Topic `json:"topic_id" orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Topic))
	orm.RegisterModel(new(SubTopic))
}

func AddNewTopic(t Topic) error {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Insert(&t)
	return err
}

func GetAllTopic() []Topic {
	o := orm.NewOrm()
	var topics []Topic
	_, _ = o.QueryTable("topic").All(&topics)
	return topics
}
