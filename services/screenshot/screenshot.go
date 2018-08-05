// Example screenshot shows how to take a screenshot of the current desktop
// and show it in a window. In a comment, it also shows how to save it as
// a png.
//
// It works by getting the image of the root window, which automatically
// includes all child windows.
package screenshot

import (
  "context"
	"log"
	"time"

	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	
	"github.com/gammazero/nexus/client"
	"github.com/gammazero/nexus/wamp"
)

func TakeScreenshot(ctx context.Context, args wamp.List, kwargs, details wamp.Dict) *client.InvokeResult {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	// Use the "NewDrawable" constructor to create an xgraphics.Image value
	// from a drawable. (Usually this is done with pixmaps, but drawables
	// can also be windows.)
	ximg, err := xgraphics.NewDrawable(X, xproto.Drawable(X.RootWin()))
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now().Local().Format("2006-01-02 15:04:05")

	// Shows the screenshot in a window.
	ximg.XShowExtra(currentTime, true)

	// If you'd like to save it as a png, use:
	err = ximg.SavePng(currentTime + ".png")
	if err != nil {
    log.Fatal(err)
  }

	xevent.Main(X)

  return &client.InvokeResult{Args: wamp.List{nil}}
}
