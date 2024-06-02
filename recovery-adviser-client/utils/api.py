import requests
from typing import Optional, Dict, Any
import os

API_BASE_URL = "http://localhost:8080"
USERNAME = os.getenv('USERNAME')

def fetch_data_from_api(endpoint: str, params: Optional[Dict[str, Any]] = None) -> Optional[Dict[str, Any]]:
    """APIからデータを取得する共通関数"""
    try:
        response = requests.get(f"{API_BASE_URL}/{endpoint}", params=params)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.HTTPError as err:
        if response.status_code == 404:
            print(f"{endpoint}のデータが見つかりません。")
        else:
            print(f"{endpoint}のデータ取得に失敗しました: {response.status_code}")
        return None

def post_data_to_api(endpoint: str, json_data: Dict[str, Any]) -> bool:
    """APIにデータをポストする共通関数"""
    try:
        response = requests.post(f"{API_BASE_URL}/{endpoint}", json=json_data, params={"usr_id": USERNAME})
        response.raise_for_status()
        print(f"API呼び出し成功: {endpoint}")
        return True
    except requests.exceptions.HTTPError as err:
        print(f"API呼び出し失敗: {endpoint} (ステータスコード: {response.status_code})")
        return False
