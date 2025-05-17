# Changelog

## 1.1.0 (2025-05-17)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/OneBusAway/go-sdk/compare/v1.0.0...v1.1.0)

### Features

* **api:** api update ([99d15da](https://github.com/OneBusAway/go-sdk/commit/99d15da3372e265f6a0647da473d7f19c8279c24))

## 1.0.0 (2025-05-17)

Full Changelog: [v0.1.0-alpha.24...v1.0.0](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.24...v1.0.0)

### Chores

* **internal:** codegen related update ([681c135](https://github.com/OneBusAway/go-sdk/commit/681c135a5d30db41ce1f3abf8968e36314161d4b))

## 0.1.0-alpha.24 (2025-04-30)

Full Changelog: [v0.1.0-alpha.23...v0.1.0-alpha.24](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.23...v0.1.0-alpha.24)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#71](https://github.com/OneBusAway/go-sdk/issues/71)) ([249be00](https://github.com/OneBusAway/go-sdk/commit/249be00c6f37b9927493404b8d15f25cfce2a09c))
* **client:** accept RFC6838 JSON content types ([#72](https://github.com/OneBusAway/go-sdk/issues/72)) ([8496c56](https://github.com/OneBusAway/go-sdk/commit/8496c56f4eeef5023732a2df3c9520245273a693))
* **client:** add support for reading base URL from environment variable ([aa9e078](https://github.com/OneBusAway/go-sdk/commit/aa9e078338065441f57bbd9da3e0ab5c8ca31ee7))
* **client:** allow custom baseurls without trailing slash ([#69](https://github.com/OneBusAway/go-sdk/issues/69)) ([bd6d607](https://github.com/OneBusAway/go-sdk/commit/bd6d6078a909b7997f82794c70f3336dfc320802))
* **client:** improve default client options support ([#74](https://github.com/OneBusAway/go-sdk/issues/74)) ([703d7bc](https://github.com/OneBusAway/go-sdk/commit/703d7bce4e38e23971f26e64a99b335b2417785b))
* **client:** support custom http clients ([#81](https://github.com/OneBusAway/go-sdk/issues/81)) ([19d138e](https://github.com/OneBusAway/go-sdk/commit/19d138e89e3926508836ef3ff7cf6b7b95abef3f))


### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#80](https://github.com/OneBusAway/go-sdk/issues/80)) ([70a742d](https://github.com/OneBusAway/go-sdk/commit/70a742db5646d7a6c7c5ded09029a1fe76c989b2))
* handle empty bodies in WithJSONSet ([f117b7d](https://github.com/OneBusAway/go-sdk/commit/f117b7d53c948c4b39f423a3bea6253db5883a9d))
* **test:** return early after test failure ([#78](https://github.com/OneBusAway/go-sdk/issues/78)) ([a3975d5](https://github.com/OneBusAway/go-sdk/commit/a3975d5cfc9c9323ae81c7be878459ce73eac1a0))


### Chores

* add request options to client tests ([#77](https://github.com/OneBusAway/go-sdk/issues/77)) ([a1409d6](https://github.com/OneBusAway/go-sdk/commit/a1409d69548e61a745a2ccaf308657ee43f1bfcb))
* **ci:** add timeout thresholds for CI jobs ([6c4e058](https://github.com/OneBusAway/go-sdk/commit/6c4e05895395d160766d8186a39b923398927d9d))
* **ci:** only use depot for staging repos ([27ea834](https://github.com/OneBusAway/go-sdk/commit/27ea83482c078d8cff0077f180289833a0b4c7f5))
* **docs:** document pre-request options ([1d1d81c](https://github.com/OneBusAway/go-sdk/commit/1d1d81c5d5b5fc251bcf6c23faaf17356a765e8c))
* **docs:** improve security documentation ([#76](https://github.com/OneBusAway/go-sdk/issues/76)) ([301ca94](https://github.com/OneBusAway/go-sdk/commit/301ca94a7f6f7925d76ce9031fc9733ccda4dfce))
* fix typos ([#79](https://github.com/OneBusAway/go-sdk/issues/79)) ([1b71db2](https://github.com/OneBusAway/go-sdk/commit/1b71db2b5d785c5657c8e2efed3202f562a941c8))
* **internal:** codegen related update ([db7e619](https://github.com/OneBusAway/go-sdk/commit/db7e619fd67b7365885ca9964ffaa61d52e984e1))
* **internal:** expand CI branch coverage ([39a0bd5](https://github.com/OneBusAway/go-sdk/commit/39a0bd532bd1884c5c0354cec004ec8be274732c))
* **internal:** reduce CI branch coverage ([dd831b8](https://github.com/OneBusAway/go-sdk/commit/dd831b8fa18ed73c2a8170cf428f63697cfbf675))
* **internal:** remove extra empty newlines ([#75](https://github.com/OneBusAway/go-sdk/issues/75)) ([b943deb](https://github.com/OneBusAway/go-sdk/commit/b943debfa05d346f776ea79477e83d6d382d81f5))


### Documentation

* update documentation links to be more uniform ([82d2142](https://github.com/OneBusAway/go-sdk/commit/82d21422f00dfe21f1562ea49406bc3671a9686e))


### Refactors

* tidy up dependencies ([#73](https://github.com/OneBusAway/go-sdk/issues/73)) ([d2fd02d](https://github.com/OneBusAway/go-sdk/commit/d2fd02d6c29f38cffaef48b61e50b3e34101daae))

## 0.1.0-alpha.23 (2025-03-01)

Full Changelog: [v0.1.0-alpha.22...v0.1.0-alpha.23](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.22...v0.1.0-alpha.23)

### Documentation

* update URLs from stainlessapi.com to stainless.com ([#66](https://github.com/OneBusAway/go-sdk/issues/66)) ([bcc4333](https://github.com/OneBusAway/go-sdk/commit/bcc4333c0a96ff289d4a8f43a522ae8395c06b62))

## 0.1.0-alpha.22 (2025-02-22)

Full Changelog: [v0.1.0-alpha.21...v0.1.0-alpha.22](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.21...v0.1.0-alpha.22)

### Chores

* **internal:** fix devcontainers setup ([#63](https://github.com/OneBusAway/go-sdk/issues/63)) ([2cb616e](https://github.com/OneBusAway/go-sdk/commit/2cb616ed77433ea83e112385a05b0a4c7d9fd9d2))

## 0.1.0-alpha.21 (2025-02-15)

Full Changelog: [v0.1.0-alpha.20...v0.1.0-alpha.21](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.20...v0.1.0-alpha.21)

### Bug Fixes

* **client:** don't truncate manually specified filenames ([#60](https://github.com/OneBusAway/go-sdk/issues/60)) ([b6d50a6](https://github.com/OneBusAway/go-sdk/commit/b6d50a62f390597b0856d8c2a6eeeb037959f2a7))

## 0.1.0-alpha.20 (2025-02-11)

Full Changelog: [v0.1.0-alpha.19...v0.1.0-alpha.20](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.19...v0.1.0-alpha.20)

### Bug Fixes

* do not call path.Base on ContentType ([#57](https://github.com/OneBusAway/go-sdk/issues/57)) ([82c787d](https://github.com/OneBusAway/go-sdk/commit/82c787de2108dfdb18a50482d927860336e3f285))

## 0.1.0-alpha.19 (2025-02-07)

Full Changelog: [v0.1.0-alpha.18...v0.1.0-alpha.19](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.18...v0.1.0-alpha.19)

### Bug Fixes

* fix early cancel when RequestTimeout is provided for streaming requests ([#54](https://github.com/OneBusAway/go-sdk/issues/54)) ([7e43254](https://github.com/OneBusAway/go-sdk/commit/7e43254657035434db3af5e3b3ace71cca4d833c))

## 0.1.0-alpha.18 (2025-02-06)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Chores

* add UnionUnmarshaler for responses that are interfaces ([#51](https://github.com/OneBusAway/go-sdk/issues/51)) ([15989be](https://github.com/OneBusAway/go-sdk/commit/15989bedd34cf5d1ed5239c9c60fb8ce59fac6fb))

## 0.1.0-alpha.17 (2025-02-04)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **client:** send `X-Stainless-Timeout` header ([#48](https://github.com/OneBusAway/go-sdk/issues/48)) ([d70788b](https://github.com/OneBusAway/go-sdk/commit/d70788bd54bf0ac439771a5fbca338ec0258f621))

## 0.1.0-alpha.16 (2025-02-01)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Documentation

* document raw responses ([#45](https://github.com/OneBusAway/go-sdk/issues/45)) ([7dd5d87](https://github.com/OneBusAway/go-sdk/commit/7dd5d87d0079a94213a35eabc1da46d53723ffdd))

## 0.1.0-alpha.15 (2025-02-01)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Bug Fixes

* fix unicode encoding for json ([#42](https://github.com/OneBusAway/go-sdk/issues/42)) ([800d2aa](https://github.com/OneBusAway/go-sdk/commit/800d2aabff07f74ff117e7efb9002d8f19549ff0))

## 0.1.0-alpha.14 (2025-01-29)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Chores

* refactor client tests ([#39](https://github.com/OneBusAway/go-sdk/issues/39)) ([d061c5a](https://github.com/OneBusAway/go-sdk/commit/d061c5ae8dd13f5cd6d02f24737124603aab6e4d))

## 0.1.0-alpha.13 (2025-01-21)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Bug Fixes

* fix apijson.Port for embedded structs ([#36](https://github.com/OneBusAway/go-sdk/issues/36)) ([27bdcc1](https://github.com/OneBusAway/go-sdk/commit/27bdcc1ba6f1a7a00b73e0107c273738c74086ae))

## 0.1.0-alpha.12 (2025-01-21)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Bug Fixes

* fix apijson.Port for embedded structs ([#33](https://github.com/OneBusAway/go-sdk/issues/33)) ([3747123](https://github.com/OneBusAway/go-sdk/commit/374712355b6b1e5d63b2a47b036559b2b6c0183e))

## 0.1.0-alpha.11 (2025-01-09)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Chores

* **internal:** codegen related update ([#30](https://github.com/OneBusAway/go-sdk/issues/30)) ([023c0fa](https://github.com/OneBusAway/go-sdk/commit/023c0fa99c466d432ac3e8c083d1b9abe0acce41))

## 0.1.0-alpha.10 (2025-01-02)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Chores

* **internal:** codegen related update ([#27](https://github.com/OneBusAway/go-sdk/issues/27)) ([00d5fc8](https://github.com/OneBusAway/go-sdk/commit/00d5fc8974188cfbc3242a56a4e280d8dfb6b7d3))

## 0.1.0-alpha.9 (2024-12-20)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Chores

* **internal:** codegen related update ([#24](https://github.com/OneBusAway/go-sdk/issues/24)) ([cfebdbd](https://github.com/OneBusAway/go-sdk/commit/cfebdbd057cf1cd6ef6fc0e71633fa17e5a4a414))

## 0.1.0-alpha.8 (2024-11-29)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** api update ([#21](https://github.com/OneBusAway/go-sdk/issues/21)) ([436994c](https://github.com/OneBusAway/go-sdk/commit/436994cc63cdbd815b2b61c53b781a210a340434))

## 0.1.0-alpha.7 (2024-11-29)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** api update ([#18](https://github.com/OneBusAway/go-sdk/issues/18)) ([61db3f0](https://github.com/OneBusAway/go-sdk/commit/61db3f0cc1a6c87589adb04f6c9cdb2e734e547a))

## 0.1.0-alpha.6 (2024-11-12)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Chores

* rebuild project due to codegen change ([#15](https://github.com/OneBusAway/go-sdk/issues/15)) ([2eba407](https://github.com/OneBusAway/go-sdk/commit/2eba407b918f5dad3d6199313587c0ce7d9d19e8))

## 0.1.0-alpha.5 (2024-11-08)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Chores

* rebuild project due to codegen change ([#12](https://github.com/OneBusAway/go-sdk/issues/12)) ([a00ba97](https://github.com/OneBusAway/go-sdk/commit/a00ba9724810098cc2fa259b6bcf1a15aabcedfc))

## 0.1.0-alpha.4 (2024-10-02)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Chores

* **internal:** codegen related update ([#9](https://github.com/OneBusAway/go-sdk/issues/9)) ([cc22b69](https://github.com/OneBusAway/go-sdk/commit/cc22b6911bbd64a2dbfdb43bd5af757da38ba651))

## 0.1.0-alpha.3 (2024-09-27)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Chores

* **internal:** codegen related update ([#6](https://github.com/OneBusAway/go-sdk/issues/6)) ([aac16f7](https://github.com/OneBusAway/go-sdk/commit/aac16f7a973c736257e7661b767e5743d9e35655))

## 0.1.0-alpha.2 (2024-09-12)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/OneBusAway/go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* feat: Go sdk examples ([39f90f2](https://github.com/OneBusAway/go-sdk/commit/39f90f281c61fe40cd9e7a511166b66dd23d65f0))

## 0.1.0-alpha.1 (2024-09-12)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/OneBusAway/go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* Add go sdk examples ([f7b1d31](https://github.com/OneBusAway/go-sdk/commit/f7b1d319199a5379f9a6f9096764383950e2b417))
* Add go sdk examples ([231fd0d](https://github.com/OneBusAway/go-sdk/commit/231fd0da4d821ad27f2fb53320be7550e0fc7109))
* agency endpoint example ([aaff320](https://github.com/OneBusAway/go-sdk/commit/aaff320c025046d34f4558c04cea9b384f6ad1f3))
* agency example ([30741c6](https://github.com/OneBusAway/go-sdk/commit/30741c64dc257830b4518e9663becb65483be6d7))
* **api:** OpenAPI spec update via Stainless API ([71cee9f](https://github.com/OneBusAway/go-sdk/commit/71cee9f771749c814ee1869a4e652f73386164ea))
* **api:** OpenAPI spec update via Stainless API ([4f8fe99](https://github.com/OneBusAway/go-sdk/commit/4f8fe997ee45e480e369bfce2c72026ac03e5027))
* **api:** update via SDK Studio ([ad2aba0](https://github.com/OneBusAway/go-sdk/commit/ad2aba00888e42c28dbdf34f4eaa0ddb6745e0b1))
* **api:** update via SDK Studio ([e1f68d7](https://github.com/OneBusAway/go-sdk/commit/e1f68d70d11592658e1c8ed6794bb4ed369aa978))
* **api:** update via SDK Studio ([90c0a8c](https://github.com/OneBusAway/go-sdk/commit/90c0a8c5abafb3820610fb704ccb9f1b064938f1))
* **api:** update via SDK Studio ([2258ba8](https://github.com/OneBusAway/go-sdk/commit/2258ba827a3b3956b06c8b8d5c9f6ac4ffc2cbb4))
* **api:** update via SDK Studio ([0604148](https://github.com/OneBusAway/go-sdk/commit/0604148868d23a7a0163d9b961ef8133aa98b94f))
* **api:** update via SDK Studio ([4468268](https://github.com/OneBusAway/go-sdk/commit/446826825e8597dfbb85593f7a087e74903d38f2))
* **api:** update via SDK Studio ([b0d4789](https://github.com/OneBusAway/go-sdk/commit/b0d4789faa6ca26d78136a5c446d367f8187683b))
* **api:** update via SDK Studio ([f5dd9a2](https://github.com/OneBusAway/go-sdk/commit/f5dd9a295b7c8c789cc901101d1640dccfc15436))
* **api:** update via SDK Studio ([c8ed506](https://github.com/OneBusAway/go-sdk/commit/c8ed5061284e177fe25fa4bcf5091b179cbae9fd))
* **api:** update via SDK Studio ([8ae72ab](https://github.com/OneBusAway/go-sdk/commit/8ae72ab351d29fe6b2e1d4da06ffb97daef7b3a2))
* **api:** update via SDK Studio ([fcdcec7](https://github.com/OneBusAway/go-sdk/commit/fcdcec72ec8bde000ababf82558c4fd45eb69a35))
* **api:** update via SDK Studio ([d52db59](https://github.com/OneBusAway/go-sdk/commit/d52db597c072fc53689c4bf9ec10de2da0ae50af))
* **api:** update via SDK Studio ([3075e60](https://github.com/OneBusAway/go-sdk/commit/3075e60851633c3fa2380d7ab7c766217619b3d6))


### Bug Fixes

* **requestconfig:** copy over more fields when cloning ([c3b242b](https://github.com/OneBusAway/go-sdk/commit/c3b242bada9dc8a2150dfe396fd87c9cd2cc5111))


### Chores

* configure new SDK language ([7806d6d](https://github.com/OneBusAway/go-sdk/commit/7806d6d318dcdb542672de078ff385c06fdd15cd))
* go live ([#1](https://github.com/OneBusAway/go-sdk/issues/1)) ([d68e58b](https://github.com/OneBusAway/go-sdk/commit/d68e58b11500a7ec3a2ce6873b045fea2b827cd3))
* **internal:** codegen related update ([46fc949](https://github.com/OneBusAway/go-sdk/commit/46fc9493350f1a19c086d92fc1642820624be106))
