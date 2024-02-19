FROM debian:stable-slim

# Install bash (optional, replace with your preferred shell or tool)
RUN apt-get update && apt-get install -y bash && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY TideTracker /bin/TideTracker
COPY .env /bin/.env
COPY entrypoint.sh /bin/entrypoint.sh

# Make sure the entrypoint script is executable
RUN chmod +x /bin/entrypoint.sh

CMD ["/bin/entrypoint.sh"]
