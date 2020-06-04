# Weasley Clock 🕰️
> Stolen to the [Weasley family][weasley-clock], monitor each of your family members whereabouts.

Each family member will update their own status using a [`telegram`][telegram-bot] bot you'll have to set up.
The data will be stored in a [`firebase` database][firebase] you'll maintain.
The clock will refresh itself every now and then, no need to refresh any page!

And the whole thing can be hosted at `netlify` at no cost! 💸

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/thisiserico/weasley-clock)

[weasley-clock]: https://harrypotter.fandom.com/wiki/Weasley_Clock
[telegram-bot]: https://core.telegram.org/bots
[firebase]: https://firebase.google.com

## 🛠️ Set up
It's recommended to set up the dependencies before `netlify` itself.
It should take no more than 10 minutes and you can start using your clock right away after deploying it 🚀

### 🗄️ Firebase database
Create a new project in the [`firebase` console][firebase-console] (there's no need to add analytics to it).
Once created, head over to the database section and provision a new [_realtime database_][firebase-realtime] in _test mode_.

⚠️ _That makes your data visible to anyone with the link to it, hiding this data is an upcoming feature._

Modify and run the following `curl` command specifying your database name in the address, the statuses you want in your clock and the people that will be displayed.

```sh
curl -X PUT 'https://{your-database-name}.firebaseio.com/.json' \
-H 'Content-Type: application/json' \
--data-raw '{
	"statuses": [
		"dentist",
		"prison",
		"lost",
		"quidditch",
		"dying",
		"tailor",
		"bed",
		"holidays",
		"forest",
		"work",
		"garden",
		"school",
		"home"
	],
	"people": {
		"ron": {
			"name": "Ron",
			"status": "bed"
		},
		"ginny": {
			"name": "Ginny",
			"status": "quidditch"
		},
		"fred": {
			"name": "Fred",
			"status": "dying"
		},
		"george": {
			"name": "George",
			"status": "work"
		}
	}
}'
```

[firebase-console]: https://console.firebase.google.com
[firebase-realtime]: https://firebase.google.com/products/realtime-database

### 🤖 Telegram bot
Reach out to [`BotFather`][bot-father] within telegram and use the `/newbot` command.
Simply follow their instructions (give your bot a name and username) so you can get your bot access token back.

After that, you're free to customize your bot a bit setting up a profile picture for it (`/setuserpic`) and other niceties, but I'll let you explore that yourself.

Now you're gonna need the `telegram chat ID` for each person that will be using the clock.
To get that, each person will have to reach out to [`IDBot`][id-bot] and trigger the `/getid` command to get their chat ID.
You can also complete this step yourself for now and come back to it once the clock is running.

[bot-father]: https://telegram.me/BotFather
[id-bot]: https://telegram.me/myidbot

### 🌍 Netlify environment
Click the `deploy to netlify` button above.
At some point, you'll be asked to configure environment variables.
It's here where you'll need to enter data that you got in previous steps:

- The `ACCEPTED_MEMBERS` variable controls who can interact with the bot.
It follows the `\w+:\d+( \w+:\d+)+` format, where `\w+` is the key for that person's entry in firebase and `\d+` the number you got after spending a quality minute with `IDBot`.
Add a single `ACCEPTED_MEMBERS` entry, splitting each `name_id:chat_id` pair with a space, something like `ron:123 ginny:456`.

- The `FIREBASE_ADDRESS` variable let the functions running within `netlify` know where the data is stored.
Use `https://{your-database-name}.firebaseio.com` as value, replacing the subdomain in that address.

- Lastly, the `TELEGRAM_TOKEN` variable will be used to send replies from a `netlify` function to you.
Use the value you previously got when talking to `BotFather`, something like `number:string`.

Continue the set up process until the end and let `netlify` do its magic.
After these steps, you'll be able to see your clock in the provided address (something like `https://{subdomain}.netlify.app`).

### 🔌 Telegram webhook
The last step is to let the `telegram bot` you created know where to send the messages it receives.
For that, a `webhook` will be released as part of your `netlify` deploy.
Run the following `curl` command, replacing the `{telegram_token}` placeholder in the address and the `{subdomain}` from the body.

```sh
curl -X GET 'https://api.telegram.org/bot{telegram_token}/setWebhook' \
-H 'Content-Type: application/json' \
--data-raw '{
	"url": "https://{subdomain}.netlify.app/.netlify/functions/webhook"
}'
```

That's it, you're ready to go! ⚡⚡⚡
Start talking to your bot 🤖 and see how the clock updates (it will take 20 seconds at most!) 🕰️

## 🏗️ How the clock works
The clock UI is being powered by a [`vue.js`][vue] application.
The clock itself is a bunch of `svg` elements put together.
All the credit goes to [Marta][marta], who helped compose the `svg` in a way that could be understood and maintained.

The data itself comes from `firebase`, although it reaches the UI through a [`netlify lambda function`][netlify-functions] written in `go`.
This prevents the `firebase` address to leak into the UI.

The webhook that `telegram` uses to update the statuses is another `go` `lambda function`.
Some validations –chat ID, status– are being performed as part of the execution.

```
+----------+       +--------------------------+       +-------------------+
| telegram | ----> | netlify webhook function | ----> | firebase database |
+----------+       +--------------------------+       +-------------------+
                                                             ^
                                                             |
+----------+       +---------------------------+             |
| clock UI | ----> | netlify statuses function | ------------+
+----------+       +---------------------------+
```

[vue]: https://vuejs.org
[marta]: https://github.com/mbondyra
[netlify-functions]: https://www.netlify.com/products/functions

## 🕵️‍♀️ Contributing
Security around the `firebase` data is coming at some point.
At the same time, support to display people faces, as opposed to their initial, is also coming!

Besides those, feel free to tweak and extend the clock to your needs, we'd love to see what you come up with! 🙌

