## 1.0.0 (2026-04-10)

### :sparkles: Features

* add `--skip-path-prompt` possibility to use computed paths only ([4681312](https://github.com/FournyP/deepsearch-mockgen-cli/commit/46813120770b33b54c4c1f5150d98684879f7ea9))
* add `-all` option to accept all generation and do not prompt for every interfaces ([bc70ef2](https://github.com/FournyP/deepsearch-mockgen-cli/commit/bc70ef23dd16758be3f73ca1b31e97aa28fb4450))
* add CTRL+C handling to get out of the CLI ([780c1c1](https://github.com/FournyP/deepsearch-mockgen-cli/commit/780c1c18e1a6e201f295a45ddf0cef82cf55f4b1))
* add progress bar and display errors only at interface generation ([49baffb](https://github.com/FournyP/deepsearch-mockgen-cli/commit/49baffb181fb0a4da056fdaa9c57c4834b7975fc))
* add setup instructions to `AGENTS.md` and `copilot-instructions.md` ([699d312](https://github.com/FournyP/deepsearch-mockgen-cli/commit/699d312c4ba118855ad691040ae3a9b581014200))
* enhance CLI usage information and update flag descriptions ([a4c321c](https://github.com/FournyP/deepsearch-mockgen-cli/commit/a4c321ce966f30885b049d2103da74f900bda6ef))
* make list display variable according to terminal size ([dc22bbc](https://github.com/FournyP/deepsearch-mockgen-cli/commit/dc22bbc38741982717ffe7090d658764b42ee707))
* publish prebuilt binaries for linux, macos and windows ([9ad5928](https://github.com/FournyP/deepsearch-mockgen-cli/commit/9ad59286302c2f8318667eb3e9a1f70ae32bc195))
* usage of bubbletea in every interactions ([483b870](https://github.com/FournyP/deepsearch-mockgen-cli/commit/483b87099c9a895cdc126390a5be1038a33209ca))

### :bug: Bug Fixes

* handle acronyms in snake_case filename conversion ([1d281a2](https://github.com/FournyP/deepsearch-mockgen-cli/commit/1d281a2ebd4106d77f275dcff2350091d2a00065))
* remove '_i' replace ([a1fe4d4](https://github.com/FournyP/deepsearch-mockgen-cli/commit/a1fe4d4d73db00c147dbfa5b52cb34de85a36e62))

### :wrench: Miscellaneous Chores

* move copilot-instructions to the right place ([f8ed57b](https://github.com/FournyP/deepsearch-mockgen-cli/commit/f8ed57b7916f0a0b69dbc2890cbfcd4c0ff5e77a))
* upgrade dependencies ([3c48e4b](https://github.com/FournyP/deepsearch-mockgen-cli/commit/3c48e4ba6417e1b84520e485fb09141e38737b18))

### :recycle: Code Refactors

* rename project and make it deepsearch by design ([068307c](https://github.com/FournyP/deepsearch-mockgen-cli/commit/068307ca427c75125ea8615b45c5ac7d79bcb458))
* suffix name of the project with `-cli` ([6c0b9f9](https://github.com/FournyP/deepsearch-mockgen-cli/commit/6c0b9f982e9f4b74696eb17799103172464740a5))

### :green_heart: CI/CD

* add build gate to prevent releasing broken code ([0b379ee](https://github.com/FournyP/deepsearch-mockgen-cli/commit/0b379eecc6e17a3dedf04d2a89b32ae65df5b986))
