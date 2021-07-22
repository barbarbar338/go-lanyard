[![stars](https://img.shields.io/github/stars/barbarbar338/go-lanyard?color=yellow&logo=github&style=for-the-badge)](https://github.com/barbarbar338/go-lanyard)
[![license](https://img.shields.io/github/license/barbarbar338/go-lanyard?logo=github&style=for-the-badge)](https://github.com/barbarbar338/go-lanyard)
[![supportServer](https://img.shields.io/discord/711995199945179187?color=7289DA&label=Support&logo=discord&style=for-the-badge)](https://discord.gg/BjEJFwh)
[![forks](https://img.shields.io/github/forks/barbarbar338/go-lanyard?color=green&logo=github&style=for-the-badge)](https://github.com/barbarbar338/go-lanyard)
[![issues](https://img.shields.io/github/issues/barbarbar338/go-lanyard?color=red&logo=github&style=for-the-badge)](https://github.com/barbarbar338/go-lanyard)

# 🚀 Go Lanyard

Use Lanyard API easily in your Go app!

# 📦 Installation

-   Initialize your project (`go mod init example.com/example`)
-   Add package (`go get github.com/barbarbar338/go-lanyard`)

# 🤓 Usage

Using without websocket:

```golang
package main

import (
	"fmt"

	"github.com/barbarbar338/go-lanyard"
)

func main() {
	//                  User ID here 👇
	res := lanyard.FetchUser("331846231514939392")

	// handle presence data here
	fmt.Println(res.Data.DiscordStatus)
}
```

Using with websocket:

```golang
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/barbarbar338/go-lanyard"
)

func main() {
	//                       User ID here 👇
	ws := lanyard.CreateWS("331846231514939392", func(data *lanyard.LanyardData) {

		// handle presence data here
		fmt.Println(data.DiscordStatus)
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	<-sc

	fmt.Println("Closing client.")

	// destroy ws before exit
	ws.Destroy()
}
```

# 📄 License

Copyright © 2021 [Barış DEMİRCİ](https://github.com/barbarbar338).

Distributed under the [GPL-3.0](https://www.gnu.org/licenses/gpl-3.0.html) License. See `LICENSE` for more information.

# 🧦 Contributing

Feel free to use GitHub's features.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/my-feature`)
3. Commit your Changes (`git commit -m 'my awesome feature my-feature'`)
4. Push to the Branch (`git push origin feature/my-feature`)
5. Open a Pull Request

# 🔥 Show your support

Give a ⭐️ if this project helped you!

# 📞 Contact

-   Mail: demirci.baris38@gmail.com
-   Discord: https://discord.gg/BjEJFwh
-   Instagram: https://www.instagram.com/ben_baris.d/

# ✨ Special Thanks

-   [Phineas](https://github.com/Phineas) - Creator of [Lanyard API](https://github.com/Phineas/lanyard)
