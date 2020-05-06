package zhenai

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>([\d]+)岁</div>`)
var height = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>([\d]+)cm</div>`)
var weight = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>([\d]+)kg</div>`)
var salary = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>月收入:([\d]+)</div>`)
var constellation = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>(天秤座)</div>`)
var marry = regexp.MustCompile(`<div class="m-btn purle" date-v-bff6f798>(已婚)</div>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, height))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weight))
	if err == nil {
		profile.Weight = weight
	}

	profile.Salary = extractString(contents, salary)

	profile.Constellation = extractString(contents, constellation)
	if extractString(contents, marry) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}
