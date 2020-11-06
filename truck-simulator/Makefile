all: build install

build:
	-docker run -dp 5000:5000 --restart=always --name registry registry
	docker build -t localhost:5000/truck-simulator:latest .
	docker push localhost:5000/truck-simulator:latest

install:
	helm upgrade \
		--install \
		--set image.repository="localhost:5000/truck-simulator" \
		--set image.tag="latest" \
		--set slack.webhook_url="${SLACK_WEBHOOK_URL}" \
		--set slack.custom_message="<%s|Eher> is delivering %s." \
		--set world_of_trucks.profile_url="https://www.worldoftrucks.com/en/profile/4992950" \
		track-eher \
		.helm
	helm upgrade \
		--install \
		--set image.repository="localhost:5000/truck-simulator" \
		--set image.tag="latest" \
		--set slack.webhook_url="${SLACK_WEBHOOK_URL}" \
		--set slack.custom_message="<%s|Malukenho> is delivering %s." \
		--set world_of_trucks.profile_url="https://www.worldoftrucks.com/en/profile/5004217" \
		track-malukenho \
		.helm

uninstall:
	helm uninstall track-eher
	helm uninstall track-malukenho
