// Package tupian 图片获取集合
package tupian

import (
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	shouer   = "https://iw233.cn/api.php?sort=cat&type"
	baisi    = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=ergbskjhebrgkjlhkerjsbkbregsbg"
	heisi    = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=rsetbgsekbjlghelkrabvfgheiv"
	siwa     = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=dsrgvkbaergfvyagvbkjavfwe"
	bizhi    = "https://iw233.cn/api.php?sort=iw233&type"
	baimao   = "https://iw233.cn/api.php?sort=yin&type"
	xing     = "https://iw233.cn/api.php?sort=xing&type"
	yuan     = "https://sakura.iw233.cn/Tag/API/web/api/op.php"
	sese     = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=qwuydcuqwgbvwgqefvbwgueahvbfkbegh"
	biaoqing = "https://iw233.cn/api.php?sort=img"
	wife     = "http://bh.ayud.top/meyu"
)

func init() {
	engine := control.Register("tupian", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "全部图片指令\n" +
			"- 兽耳\n" +
			"- 白毛\n" +
			"- ！原神\n" +
			"- 黑丝\n" +
			"- 白丝\n" +
			"- 丝袜\n" +
			"- 老婆照片\n" +
			"- 随机壁纸\n" +
			"- 星空\n" +
			"- 随机表情包\n" +
			"- 涩涩哒咩/我要涩涩\n",
	})
	engine.OnFullMatch("兽耳").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(shouer)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("白丝").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(baisi)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("黑丝").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(heisi)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("丝袜").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(siwa)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("随机壁纸").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(bizhi)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("白毛").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(baimao)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("星空").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(xing)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("！原神").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(yuan)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatchGroup([]string{"涩涩哒咩", "我要涩涩"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(sese)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("随机表情包").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(biaoqing)
			if err != nil {
				ctx.SendChain(message.Text("获取图片失败惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
	engine.OnFullMatch("老婆照片").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.GetData(wife)
			if err != nil {
				ctx.SendChain(message.Text("你老婆丢掉惹", err))
				return
			}
			ctx.SendChain(message.ImageBytes(data))
		})
}
