// package name must be "bot_test". not "bot"
package twitter_bot_test

import (
	"testing"

	"github.com/daimatz/gomock_example/mock_gomock_example"
	"github.com/daimatz/gomock_example/twitter_bot"

	"code.google.com/p/gomock/gomock"
)

func TestAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_twitter_bot.NewMockTwitterService(ctrl)
	service.EXPECT().GetTimeline(10).Return(make([]twitter_bot.Status, 0), nil)
	bot := &twitter_bot.TwitterBotImpl{service}
	tl, _ := bot.Action(10)
	if len(tl) != 0 {
		t.Error("fail")
	}
}
