FROM debian:stable-slim

# Install bash (optional, replace with your preferred shell or tool)
RUN apt-get update && apt-get install -y bash

COPY TideTracker /bin/TideTracker
COPY .env /bin/.env
COPY entrypoint.sh /bin/entrypoint.sh

# Make sure the entrypoint script is executable
RUN chmod +x /bin/entrypoint.sh

CMD ["/bin/entrypoint.sh"]
