package services

import (
	"encoding/json"
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

var Tags = new(TagsService)

type TagsService struct {
}

// Batch insertion
func (this *TagsService) BatchAdd(uid int32, tags []string, tag_colors_str string) bool {
	if len(tags) == 0 {
		return true
	}

	tag_Colors := dto.AbTag_colors{}

	_ = json.Unmarshal([]byte(tag_colors_str), &tag_Colors)

	tagList := []models.Tags{}
	for _, t := range tags {
		tag_color := ""
		if c, found := tag_Colors[t]; found {
			tag_color = strconv.FormatInt(c, 10)
		}

		tagList = append(tagList, models.Tags{
			Uid:   uid,
			Tag:   t,
			Color: tag_color,
		})
	}

	_, err := orm.NewOrm().InsertMulti(3, tagList)
	if err != nil {
		return false
	}
	return true
}

func (this *TagsService) DeleteAll(uid int32) bool {
	_, err := orm.NewOrm().Raw("delete from rustdesk_tags where uid = ?", uid).Exec()
	if err != nil {
		return false
	}
	return true
}

// Query the tag of the user name
func (this *TagsService) FindTags(uid int32) []models.Tags {
	ret := []models.Tags{}
	_, err := orm.NewOrm().QueryTable(new(models.Tags)).Filter("uid", uid).All(&ret, "tag", "color")
	if err != nil {
		return nil
	}
	return ret
}
