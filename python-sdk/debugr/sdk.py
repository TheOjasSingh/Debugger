import requests

class DebugrClient:
    def __init__(self, base_url: str):
        self.base_url = base_url.rstrip('/')

    def trace(self, pid: int) -> str:
        response = requests.get(f"{self.base_url}/trace", params={"pid": pid})
        response.raise_for_status()
        return response.json().get("stack", "")