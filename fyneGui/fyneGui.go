package fyneGui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-toast/toast"
	"newTarot/tarot"
	"newTarot/tarots"
	"newTarot/todo"
	"reflect"
	"time"
)

var (
	StartTime  time.Time
	GoHomeTime time.Time
	Once       = 0
)

func GuiStart() {

	myapp := app.New()
	myapp.Settings().SetTheme(theme.LightTheme())
	w := myapp.NewWindow("tarot")

	w.Resize(fyne.Size{
		Width:  200,
		Height: 200,
	})
	//计时器组件
	f := binding.NewFloat()
	f.Set(1.0)
	progress := widget.NewProgressBarWithData(f)

	//按钮1
	tarotButton := widget.NewButton("tarot", func() {
		w2 := myapp.NewWindow("tarot")
		card := tarot.GetTarotCard()

		cardfile := "tarotCards/" + card.Id

		file, err := tarots.TarotCards.Open(cardfile)
		if err != nil {
			fmt.Println(err)
		}

		img := canvas.NewImageFromReader(file, "tarot1.jpg")
		img.FillMode = canvas.ImageFillOriginal
		w2.SetContent(img)
		w2.Show()
	})

	// todolist button
	todoBtn := widget.NewButton("todo", func() {

		w := myapp.NewWindow("TODO App")

		todos := binding.NewUntypedList()

		newTodoDescTxt := widget.NewEntry()
		newTodoDescTxt.PlaceHolder = "New Todo Description..."

		addBtn := widget.NewButton("Add", func() {
			newtodo := todo.NewTodo(newTodoDescTxt.Text)
			todos.Prepend(&newtodo)

			list, _ := todos.Get()
			todo.AddTodo(list)
			newTodoDescTxt.SetText("")
		})
		addBtn.Disable()

		newTodoDescTxt.OnChanged = func(s string) {
			if len(s) >= 3 {
				addBtn.Enable()
				return
			}
			addBtn.Disable()
		}
		data := todo.LoadTodo()

		for _, t := range data.TodoList {
			todos.Append(t)
		}

		saveBtn := widget.NewButton("save", func() {
			list, _ := todos.Get()
			todo.AddTodo(list)
		})

		w.SetContent(container.NewBorder(nil,
			container.NewBorder(nil,
				saveBtn,
				nil,
				addBtn,
				newTodoDescTxt,
			),
			nil,
			nil,
			widget.NewListWithData(
				todos,
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil,
						widget.NewLabel(""),
						widget.NewButton("remove", func() {
						}),
						widget.NewCheck("", func(b bool) {

						}),
					)
				},
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)

					l := new(widget.Label)
					c := new(widget.Check)
					b := new(widget.Button)
					for _, v := range ctr.Objects {
						switch v.(type) {
						case *widget.Label:
							l = v.(*widget.Label)
						case *widget.Check:
							c = v.(*widget.Check)
						case *widget.Button:
							b = v.(*widget.Button)
						}
					}

					diu, _ := di.(binding.Untyped).Get()
					newtodo := diu.(*todo.Todo)
					b.OnTapped = func() {
						idx := reflect.ValueOf(di).Elem().FieldByName("index")
						list, _ := todos.Get()
						todos.Set(make([]interface{}, 0, todos.Length()-1))

						for i, v := range list {
							if int64(i) == idx.Int() {
								continue
							}
							todos.Append(v)
						}

						newtodos, _ := todos.Get()
						todo.RemoveTodo(newtodos)
					}

					c.OnChanged = func(b bool) {
						newtodo.Done = b
					}
					l.SetText(newtodo.Description)
					c.Bind(binding.BindBool(&newtodo.Done))

				},
			)))

		w.Resize(fyne.Size{
			Width:  400,
			Height: 800,
		})
		w.Show()

	})

	startButton := widget.NewButton("start", func() {
		//开始倒计时
		if Once == 0 {
			Once = 1
			start(f)
		}

	})

	w.SetContent(container.New(layout.NewGridLayout(2), tarotButton, startButton, todoBtn, progress))

	//设置主窗口
	w.SetMaster()

	w.ShowAndRun()
}

func start(f binding.Float) {

	StartTime = time.Now()
	GoHomeTime = StartTime.Add(9 * time.Hour)

	go func() {
		for range time.Tick(1 * time.Second) {
			now := time.Now()
			//距离下班的时间差的int64
			since := GoHomeTime.Unix() - now.Unix()
			//距离下班的分钟

			//sincemin := GoHomeTime.Sub(now).Minutes()
			//
			//s.Set("\nRemaining Time:\n" + strconv.Itoa(int(sincemin)) + " min")

			num := float64(since) / float64(GoHomeTime.Unix()-StartTime.Unix())

			f.Set(num)

			if num <= 0.0 && Once == 1 {
				Notify()
				Once = 0
			}
		}
	}()
}

func Notify() {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "下班",
		Message: "到点了该打卡下班了",
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "打卡!", "https://www.dingtalk.com/"},
			//{"protocol", "按钮2", "https://github.com/"},
		},
	}
	notification.Push()

}
