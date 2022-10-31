"""Test script for the test time app."""
from json import loads
from multiprocessing import Process
import unittest
from requests import Session
from requests.adapters import HTTPAdapter
from requests.packages.urllib3.util.retry import Retry
import time_app


class TimeAppTests(unittest.TestCase):
    """Tests the time app."""

    @classmethod
    def setUpClass(cls):
        """Start time app server."""
        cls.server = Process(target=time_app.time_app, daemon=True)
        cls.server.start()

        retry_strategy = Retry(
            total=3,
            backoff_factor=0.5,
        )

        cls.http = Session()
        cls.http.mount("http://", HTTPAdapter(max_retries=retry_strategy))

    @classmethod
    def tearDownClass(cls):
        """Stop time app server."""
        cls.server.terminate()
        cls.http.close()

    def test_time_app(self):
        """Test all connections output."""
        response = self.http.get('http://localhost:8080/', timeout=5)

        self.assertTrue(response.status_code == 200)
        self.assertIn("New York", response.text)

    def test_time_app_health(self):
        """Test all connections output."""
        response = loads(self.http.get('http://localhost:8080/health', timeout=5).text)

        self.assertIn("status_code", response)
        self.assertEqual(response["status_code"], 200)

    def test_time_app_wrong_path(self):
        """Test all connections output."""
        response = self.http.get('http://localhost:8080/foo', timeout=5)

        self.assertTrue(response.status_code == 404)


if __name__ == '__main__':
    unittest.main()
