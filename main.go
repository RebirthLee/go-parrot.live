package main

import (
	_ "embed"
	"fmt"
	termcolor "github.com/rebirthlee/go-term-colorize"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	frameDelay    = 70 * time.Millisecond
	clearTerminal = "\x1b[2J\x1b[3J\x1b[H"
)

var (
	colors = []termcolor.ColorCode{
		termcolor.StdRed,
		termcolor.StdYellow,
		termcolor.StdGreen,
		termcolor.StdBlue,
		termcolor.StdMagenta,
		termcolor.StdCyan,
		termcolor.StdWhite,
	}

	frames = []string{
		"                         .cccc;;cc;';c.           \n                      .,:dkdc:;;:c:,:d:.          \n                     .loc'.,cc::c:::,..;:.        \n                   .cl;....;dkdccc::,...c;        \n                  .c:,';:'..ckc',;::;....;c.      \n                .c:'.,dkkoc:ok:;llllc,,c,';:.     \n               .;c,';okkkkkkkk:;lllll,:kd;.;:,.   \n               co..:kkkkkkkkkk:;llllc':kkc..oNc   \n             .cl;.,oxkkkkkkkkkc,:cll;,okkc'.cO;   \n             ;k:..ckkkkkkkkkkkl..,;,.;xkko:',l'   \n            .,...';dkkkkkkkkkkd;.....ckkkl'.cO;   \n         .,,:,.;oo:ckkkkkkkkkkkdoc;;cdkkkc..cd,   \n      .cclo;,ccdkkl;llccdkkkkkkkkkkkkkkkd,.c;     \n     .lol:;;okkkkkxooc::coodkkkkkkkkkkkko'.oc     \n   .c:'..lkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkd,.oc     \n  .lo;,:cdkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkd,.c;     \n,dx:..;lllllllllllllllllllllllllllllllllc'...     \ncNO;........................................      \n",
		"                .ckx;'........':c.                \n             .,:c:::::oxxocoo::::,',.             \n            .odc'..:lkkoolllllo;..;d,             \n            ;c..:o:..;:..',;'.......;.            \n           ,c..:0Xx::o:.,cllc:,'::,.,c.           \n           ;c;lkXKXXXXl.;lllll;lKXOo;':c.         \n         ,dc.oXXXXXXXXl.,lllll;lXXXXx,c0:         \n         ;Oc.oXXXXXXXXo.':ll:;'oXXXXO;,l'         \n         'l;;kXXXXXXXXd'.'::'..dXXXXO;,l'         \n         'l;:0XXXXXXXX0x:...,:o0XXXXx,:x,         \n         'l;;kXXXXXXXXXKkol;oXXXXXXXO;oNc         \n        ,c'..ckk0XXXXXXXXXX00XXXXXXX0:;o:.        \n      .':;..:do::ooookXXXXXXXXXXXXXXXo..c;        \n    .',',:co0XX0kkkxxOXXXXXXXXXXXXXXXOc..;l.      \n  .:;'..oXXXXXXXXXXXXXXXXXXXXXXXXXXXXXko;';:.     \n.ldc..:oOXKXXXXXXKXXKXXXXXXXXXXXXXXXXXXXo..oc     \n:0o...:dxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxo,.:,     \ncNo........................................;'     \n",
		"            .cc;.  ...  .;c.                      \n         .,,cc:cc:lxxxl:ccc:;,.                   \n        .lo;...lKKklllookl..cO;                   \n      .cl;.,:'.okl;..''.;,..';:.                  \n     .:o;;dkd,.ll..,cc::,..,'.;:,.                \n     co..lKKKkokl.':lloo;''ol..;dl.               \n   .,c;.,xKKKKKKo.':llll;.'oOxl,.cl,.             \n   cNo..lKKKKKKKo'';llll;;okKKKl..oNc             \n   cNo..lKKKKKKKko;':c:,'lKKKKKo'.oNc             \n   cNo..lKKKKKKKKKl.....'dKKKKKxc,l0:             \n   .c:'.lKKKKKKKKKk;....lKKKKKKo'.oNc             \n     ,:.'oxOKKKKKKKOxxxxOKKKKKKxc,;ol:.           \n     ;c..'':oookKKKKKKKKKKKKKKKKKk:.'clc.         \n   ,xl'.,oxo;'';oxOKKKKKKKKKKKKKKKOxxl:::;,.      \n  .dOc..lKKKkoooookKKKKKKKKKKKKKKKKKKKxl,;ol.     \n  cx,';okKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKl..;lc.   \n  co..:dddddddddddddddddddddddddddddddddl::',::.  \n  co...........................................   \n",
		"           .ccccccc.                              \n      .,,,;cooolccoo;;,,.                         \n     .dOx;..;lllll;..;xOd.                        \n   .cdo;',loOXXXXXkll;';odc.                      \n  ,ol:;c,':oko:cccccc,...ckl.                     \n  ;c.;kXo..::..;c::'.......oc                     \n,dc..oXX0kk0o.':lll;..cxxc.,ld,                   \nkNo.'oXXXXXXo',:lll;..oXXOo;cOd.                  \nKOc;oOXXXXXXo.':lol;..dXXXXl';xc                  \nOl,:k0XXXXXX0c.,clc'.:0XXXXx,.oc                  \nKOc;dOXXXXXXXl..';'..lXXXXXo..oc                  \ndNo..oXXXXXXXOx:..'lxOXXXXXk,.:; ..               \ncNo..lXXXXXXXXXOolkXXXXXXXXXkl,..;:';.            \n.,;'.,dkkkkk0XXXXXXXXXXXXXXXXXOxxl;,;,;l:.        \n  ;c.;:''''':doOXXXXXXXXXXXXXXXXXXOdo;';clc.      \n  ;c.lOdood:'''oXXXXXXXXXXXXXXXXXXXXXk,..;ol.     \n  ';.:xxxxxocccoxxxxxxxxxxxxxxxxxxxxxxl::'.';;.   \n  ';........................................;l'   \n",
		"                                                  \n        .;:;;,.,;;::,.                            \n     .;':;........'co:.                           \n   .clc;'':cllllc::,.':c.                         \n  .lo;;o:coxdllllllc;''::,,.                      \n.c:'.,cl,.'l:',,;;'......cO;                      \ndo;';oxoc;:l;;llllc'.';;'.,;.                     \nc..ckkkkkkkd,;llllc'.:kkd;.':c.                   \n'.,okkkkkkkkc;lllll,.:kkkdl,cO;                   \n..;xkkkkkkkkc,ccll:,;okkkkk:,co,                  \n..,dkkkkkkkkc..,;,'ckkkkkkkc;ll.                  \n..'okkkkkkkko,....'okkkkkkkc,:c.                  \nc..ckkkkkkkkkdl;,:okkkkkkkkd,.',';.               \nd..':lxkkkkkkkkxxkkkkkkkkkkkdoc;,;'..'.,.         \no...'';llllldkkkkkkkkkkkkkkkkkkdll;..'cdo.        \no..,l;'''''';dkkkkkkkkkkkkkkkkkkkkdlc,..;lc.      \no..;lc;;;;;;,,;clllllllllllllllllllllc'..,:c.     \no..........................................;'     \n",
		"                                                  \n           .,,,,,,,,,.                            \n         .ckKxodooxOOdcc.                         \n      .cclooc'....';;cool.                        \n     .loc;;;;clllllc;;;;;:;,.                     \n   .c:'.,okd;;cdo:::::cl,..oc                     \n  .:o;';okkx;';;,';::;'....,:,.                   \n  co..ckkkkkddkc,cclll;.,c:,:o:.                  \n  co..ckkkkkkkk:,cllll;.:kkd,.':c.                \n.,:;.,okkkkkkkk:,cclll;.ckkkdl;;o:.               \ncNo..ckkkkkkkkko,.;loc,.ckkkkkc..oc               \n,dd;.:kkkkkkkkkx;..;:,.'lkkkkko,.:,               \n  ;:.ckkkkkkkkkkc.....;ldkkkkkk:.,'               \n,dc..'okkkkkkkkkxoc;;cxkkkkkkkkc..,;,.            \nkNo..':lllllldkkkkkkkkkkkkkkkkkdcc,.;l.           \nKOc,c;''''''';lldkkkkkkkkkkkkkkkkkc..;lc.         \nxx:':;;;;,.,,...,;;cllllllllllllllc;'.;od,        \ncNo.....................................oc        \n",
		"                                                  \n                                                  \n                   .ccccccc.                      \n               .ccckNKOOOOkdcc.                   \n            .;;cc:ccccccc:,:c::,,.                \n         .c;:;.,cccllxOOOxlllc,;ol.               \n        .lkc,coxo:;oOOxooooooo;..:,               \n      .cdc.,dOOOc..cOd,.',,;'....':l.             \n      cNx'.lOOOOxlldOc..;lll;.....cO;             \n     ,do;,:dOOOOOOOOOl'':lll;..:d:''c,            \n     co..lOOOOOOOOOOOl'':lll;.'lOd,.cd.           \n     co.'dOOOOOOOOOOOo,.;llc,.,dOOc..dc           \n     co..lOOOOOOOOOOOOc.';:,..cOOOl..oc           \n   .,:;.'::lxOOOOOOOOOo:'...,:oOOOc.'dc           \n   ;Oc..cl'':lldOOOOOOOOdcclxOOOOx,.cd.           \n  .:;';lxl''''':lldOOOOOOOOOOOOOOc..oc            \n,dl,.'cooc:::,....,::coooooooooooc'.c:            \ncNo.................................oc            \n",
		"                                                  \n                                                  \n                                                  \n                        .cccccccc.                \n                  .,,,;;cc:cccccc:;;,.            \n                .cdxo;..,::cccc::,..;l.           \n               ,do:,,:c:coxxdllll:;,';:,.         \n             .cl;.,oxxc'.,cc,.';;;'...oNc         \n             ;Oc..cxxxc'.,c;..;lll;...cO;         \n           .;;',:ldxxxdoldxc..;lll:'...'c,        \n           ;c..cxxxxkxxkxxxc'.;lll:'','.cdc.      \n         .c;.;odxxxxxxxxxxxd;.,cll;.,l:.'dNc      \n        .:,''ccoxkxxkxxxxxxx:..,:;'.:xc..oNc      \n      .lc,.'lc':dxxxkxxxxxxxol,...',lx:..dNc      \n     .:,',coxoc;;ccccoxxxxxxxxo:::oxxo,.cdc.      \n  .;':;.'oxxxxxc''''';cccoxxxxxxxxxxxc..oc        \n,do:'..,:llllll:;;;;;;,..,;:lllllllll;..oc        \ncNo.....................................oc        \n",
		"                                                  \n                                                  \n                              .ccccc.             \n                         .cc;'coooxkl;.           \n                     .:c:::c:,,,,,;c;;,.'.        \n                   .clc,',:,..:xxocc;'..c;        \n                  .c:,';:ox:..:c,,,,,,...cd,      \n                .c:'.,oxxxxl::l:.,loll;..;ol.     \n                ;Oc..:xxxxxxxxx:.,llll,....oc     \n             .,;,',:loxxxxxxxxx:.,llll;.,,.'ld,   \n            .lo;..:xxxxxxxxxxxx:.'cllc,.:l:'cO;   \n           .:;...'cxxxxxxxxxxxxoc;,::,..cdl;;l'   \n         .cl;':,'';oxxxxxxdxxxxxx:....,cooc,cO;   \n     .,,,::;,lxoc:,,:lxxxxxxxxxxxo:,,;lxxl;'oNc   \n   .cdxo;':lxxxxxxc'';cccccoxxxxxxxxxxxxo,.;lc.   \n  .loc'.'lxxxxxxxxocc;''''';ccoxxxxxxxxx:..oc     \nolc,..',:cccccccccccc:;;;;;;;;:ccccccccc,.'c,     \nOl;......................................;l'      \n",
		"                                                  \n                              ,ddoodd,            \n                         .cc' ,ooccoo,'cc.        \n                      .ccldo;...',,...;oxdc.      \n                   .,,:cc;.,'..;lol;;,'..lkl.     \n                  .dOc';:ccl;..;dl,.''.....oc     \n                .,lc',cdddddlccld;.,;c::'..,cc:.  \n                cNo..:ddddddddddd;':clll;,c,';xc  \n               .lo;,clddddddddddd;':clll;:kc..;'  \n             .,c;..:ddddddddddddd:';clll,;ll,..   \n             ;Oc..';:ldddddddddddl,.,c:;';dd;..   \n           .''',:c:,'cdddddddddddo:,''..'cdd;..   \n         .cdc';lddd:';lddddddddddddd;.';lddl,..   \n      .,;::;,cdddddol;;lllllodddddddlcldddd:.'l;  \n     .dOc..,lddddddddlcc:;'';cclddddddddddd;;ll.  \n   .coc,;::ldddddddddddddlcccc:ldddddddddl:,cO;   \n,xl::,..,cccccccccccccccccccccccccccccccc:;':xx,  \ncNd.........................................;lOc  \n",
	}
)

func newColorPallets() *ColorPallets {
	return &ColorPallets{
		shuffler:     rand.New(rand.NewSource(time.Now().UnixNano())),
		currentIndex: 0,
		pallets:      append([]termcolor.ColorCode{}, colors...),
	}
}

// ColorPallets not thread-safety
type ColorPallets struct {
	shuffler     *rand.Rand
	currentIndex int
	pallets      []termcolor.ColorCode
}

func (p *ColorPallets) shuffle() {
	p.shuffler.Shuffle(len(p.pallets), func(i, j int) {
		p.pallets[i], p.pallets[j] = p.pallets[j], p.pallets[i]
	})
}

func (p *ColorPallets) get() (res termcolor.ColorCode) {
	if len(p.pallets) <= p.currentIndex {
		p.currentIndex = 0
		p.shuffle()
	}

	res = p.pallets[p.currentIndex]
	p.currentIndex++
	return
}

var _ http.Handler = (*handler)(nil)

type handler struct{}

type try func() error

func tryCatches(tries ...try) error {
	for _, t := range tries {
		err := t()
		if err != nil {
			return err
		}
	}

	return nil
}

func tryWriteString(writer io.Writer, raw string) try {
	return func() (err error) {
		_, err = writer.Write([]byte(raw))
		return
	}
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ua := request.Header.Get("User-Agent")
	if !strings.Contains(ua, "curl") {
		writer.Header().Set("Location", "https://github.com/rebirthlee/go-parrot.live")
		writer.WriteHeader(http.StatusFound)
		return
	}

	flusher, ok := writer.(http.Flusher)
	if !ok {
		writer.Write([]byte("hasn't stream"))
		writer.WriteHeader(http.StatusBadGateway)
	}

	var colorPallets = newColorPallets()
	var tint = colorPallets.get()
	var index uint64 = 0
	for {
		err := tryCatches(
			tryWriteString(writer, clearTerminal),
			tryWriteString(writer, tint.Foreground(frames[int(index%uint64(len(frames)))])),
			tryWriteString(writer, termcolor.Reset),
		)
		if err != nil {
			//TODO: error check, logging error
			break
		}

		flusher.Flush()

		prev := time.Now()
		index++
		tint = colorPallets.get()
		diff := time.Now().Sub(prev)
		if diff < frameDelay {
			time.Sleep(frameDelay - diff)
		}
	}
}

func main() {
	port := os.Getenv("PARROT_PORT")
	if len(port) == 0 {
		port = "3000"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), &handler{})
	if err != nil {
		fmt.Println("error http.ListenAndServe")
		fmt.Println(err)
	}
}
