"""Python example time API."""
from datetime import datetime
from http.server import BaseHTTPRequestHandler, HTTPServer
import json
import logging
import os
from signal import signal, SIGTERM, SIGINT
import sys
import pytz

# Create logger
_logger = logging.getLogger("time_app")
# The log level can be overridden by setting the LOG_LEVEL environment variable
_log_level = int(os.environ.get("LOG_LEVEL", "20"))
_logger.setLevel(_log_level)
_ch = logging.StreamHandler(sys.stdout)
_ch.setLevel(_log_level)
_logger.addHandler(_ch)


class TimeServer(BaseHTTPRequestHandler):
    """Extend the HTTP server BaseHTTPRequestHandler."""

    # The configuration file path can be overridden by setting the CONFIG_FILE environment variable
    config = os.environ.get("CONFIG_FILE", "config.json")

    with open(file=config,  mode='r', encoding='utf-8') as config_file:
        time_zones = json.load(config_file)

    if not time_zones:
        _logger.critical("No time zones configured, exiting")
        sys.exit(1)

    def do_GET(self):
        """Serve GET requests for the time server.

        The server will return the times on the root path or a JSON status code on /health.
        """
        if self.path == "/":
            self.time_request()

        elif self.path == "/health":
            self.health_request()

        else:
            self.not_found()

    def time_request(self):
        """Serve time requests."""
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>Local times</title></head>", "utf-8"))
        self.wfile.write(bytes("<body>", "utf-8"))

        fmt = "%Y-%m-%d %H:%M:%S %Z%z"

        self.wfile.write(bytes("<table><tr><th>City</th><th>Time</th></tr>", "utf-8"))

        for time_zone in self.time_zones:
            tz_time = datetime.now(tz=pytz.timezone(time_zone["time_zone"]))
            self.wfile.write(bytes(f"<tr><td>{time_zone['city']}</td><td>{tz_time.strftime(fmt)}</td></tr>", "utf-8"))

        self.wfile.write(bytes("</body></html>", "utf-8"))

    def health_request(self):
        """Serve health requests."""
        self.send_response(200)
        self.send_header("Content-type", "application/json")
        self.end_headers()

        response = {
            "status_code": 200
        }

        self.wfile.write(bytes(json.dumps(response), "utf-8"))

    def not_found(self):
        """Return 404 responses."""
        self.send_response(404)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>Not found</title></head></html>", "utf-8"))
        self.wfile.write(bytes("<body>Not found</body>", "utf-8"))


def sig_handler(_signo, _stack_frame):
    """Handle SIGTERM and SIGINT signals and exit."""
    _logger.info("Shutting down")
    sys.exit(0)


def time_app():
    """Start an HTTP server to get world times."""
    # The hostname and port can be overridden by setting the HOST and PORT environment variables
    host = os.environ.get("HOST", "localhost")
    port = int(os.environ.get("PORT", 8080))

    web_server = HTTPServer((host, port), TimeServer)
    _logger.info(f"Server started http://{host}:{port}")

    signal(SIGINT, sig_handler)
    signal(SIGTERM, sig_handler)

    try:
        web_server.serve_forever()
    finally:
        web_server.server_close()
        _logger.info("Server stopped")


if __name__ == "__main__":

    time_app()
