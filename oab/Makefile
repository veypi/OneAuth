#
# Makefile
# Copyright (C) 2022 veypi <i@veypi.com>
# 2022-07-16 23:43
# Distributed under terms of the Apache license.
#

run:
	@cargo run -- -c ./cfg-demo.yml
init:
	@cargo run -- init

sqlx:
	@sqlx migrate run

entity:
	@sea-orm-cli generate entity --database-url mysql://root:123456@localhost:3306/oneauth -o ./src/models/entity/ --with-serde both --model-extra-derives Default,sqlx::FromRow --serde-skip-hidden-column --date-time-crate chrono

# install_target:
# 	@rustup target add x86_64-unknown-linux-gnu

build_linux:
	@cargo build --release --target-dir ./target/linux/
