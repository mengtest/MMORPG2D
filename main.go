package main

import(
	// "github.com/urfave/cli"
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	// "os"
)

func main(){
	g,err:= gocui.NewGui(gocui.OutputNormal)

	// g := gocui.NewGui(gocui.OutputNormal)

	if err !=nil{
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	if err := g.SetKeybinding("",gocui.KeyCtrlC,gocui.ModNone,quit);err != nil{
		log.Panicln(err)
	}
	if err:=g.MainLoop();err!=nil&&err!=gocui.ErrQuit{
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


type character struct{
	name string
	AT int
	AD int
	HP int
}

var yezhu=character{
	name :"野猪",
	AT: 3,
	AD : 2,
	HP : 10,
}
var avatar = character{
	name : "玩家",
	AT:2,
	AD:3,
	HP:10,
}

// func Avatar_Attack(target character){
// 	damage(avatar,target);
// }

// func Ai_Attack(){
// 	damage(yezhu,avatar)
// }

// func damage(attacker character,attacked character,callback CallBack){
// 	damageNum := attacker.AT - attacked.AD
// 	attacked.HP = attacked.HP - damageNum
// 	if callback{
// 		callback()
// 	}
// 	fmt.print("%s attack %s,damage is %d",attacker.name,attacked.name,damageNum )
// 	fmt.print("%s Hp is %d",attacked.name,attacked.HP)
// }

// func onDamage(){
// 	fmt.print("is damage!!!")
// }

