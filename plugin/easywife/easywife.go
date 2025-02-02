// Package easywife 简单本地老婆,可以去https://github.com/FloatTech/zbpwife下载素材
package easywife

import (
	"os"
	"regexp"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	fcext "github.com/FloatTech/floatbox/ctxext"
	"github.com/FloatTech/floatbox/file"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
)

func init() {
	engine := control.Register("easywife", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "本地老婆",
		Help:              " - 抽老婆\nTips: 可以去https://github.com/FloatTech/zbpwife下载素材",
		PrivateDataFolder: "easywife",
	})
	cachePath := engine.DataFolder() + "wife/"
	err := os.MkdirAll(cachePath, 0755)
	if err != nil {
		panic(err)
	}
	engine.OnFullMatch("抽老婆").SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			wifes, err := os.ReadDir(cachePath)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
			}
			uid := ctx.Event.UserID
			name := ctx.CardOrNickName(uid)
			n := fcext.RandSenderPerDayN(uid, len(wifes))
			wn := wifes[n].Name()
			reg := regexp.MustCompile(`[^\.]+`)
			list := reg.FindAllString(wn, -1)
			ctx.SendChain(
				message.Text(name, "さんが二次元で結婚するであろうヒロインは、", "\n"),
				message.Image("file:///"+file.BOTPATH+"/"+cachePath+wn),
				message.Text("\n【", list[0], "】です！"))
		})
}
