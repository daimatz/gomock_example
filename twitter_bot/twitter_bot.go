package twitter_bot

import (
	"fmt"
	"time"
)

type Status struct {
	UserId string
	Id     int
}

type TwitterService interface {
	Tweet(text string, inReplyTo int) (int, error)
	GetTimeline(count int) ([]Status, error)
}

type TwitterServiceImpl struct{}

func (t *TwitterServiceImpl) Tweet(text string, inReplyTo int) (int, error) {
	fmt.Printf("%d %s\n", inReplyTo, text)
	return inReplyTo + 1, nil
}

func (t *TwitterServiceImpl) GetTimeline(count int) ([]Status, error) {
	return make([]Status, count), nil
}

type TwitterBot interface {
	EventLoop() error
	Action(count int) ([]int, error)
}

type TwitterBotImpl struct {
	TwitterService TwitterService
}

func (t *TwitterBotImpl) EventLoop() error {
	for {
		_, err := t.Action(10)
		if err != nil {
			return err
		}

		time.Sleep(60 * time.Second)
	}
}

func (t *TwitterBotImpl) Action(count int) ([]int, error) {
	tl, _ := t.TwitterService.GetTimeline(count)
	ret := make([]int, 0)
	for i := range tl {
		status := tl[i]
		statusId, _ := t.TwitterService.Tweet("@"+status.UserId+" Good morning!", status.Id)
		ret = append(ret, statusId)
	}
	return ret, nil
}
