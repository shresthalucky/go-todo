up:
		@echo "Starting the project..."
		docker-compose up -d
		@echo "Project started."

build-up:
		@echo "Building the project..."
		docker-compose up --build -d
		@echo "Project built and started."

down:
		@echo "Stopping the project..."
		docker-compose down
		@echo "Project stopped."