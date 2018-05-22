package controllers

import (
	"encoding/json"
	"hatemate-api/models"
)

// Add new Topic

func (uc *UserController) AddNewTopic() {
	var topic models.Topic
	json.Unmarshal(uc.Ctx.Input.RequestBody, &topic)
	err := models.AddNewTopic(topic)

	if err == nil {
		uc.Ctx.Output.SetStatus(200)
		uc.Ctx.Output.Body([]byte("Success added new topic"))
	} else {
		uc.Ctx.Output.SetStatus(400)
		uc.Ctx.Output.Body([]byte("Error adding new topic"))
	}
}

// Get All Topic
func (uc *UserController) GetAllTopicController() {
	topics := models.GetAllTopic()
	uc.Data["json"] = topics
	uc.ServeJSON()
}
