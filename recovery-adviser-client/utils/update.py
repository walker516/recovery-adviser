from typing import Dict, Any, Optional
import os
import sys
import json
import time

import requests

from utils.api import fetch_data_from_api, post_data_to_api
from utils.prompts import confirm_action
from utils.log import extract_parameters_from_log, run_recovery_tool

API_BASE_URL = "http://localhost:8080"
BOM_UPDATE_URL = "http://example.com/bom_update"
CAABATCH_UPDATE_URL = "http://example.com/caabatch_update"
USERNAME = os.getenv('USERNAME')

def fetch_recovery_job_status(seppenbuban: str) -> Dict[str, Any]:
    """リカバリジョブステータスを取得する"""
    return fetch_data_from_api(f"recovery-job-status/{seppenbuban}", params={"usr_id": USERNAME})

def update_bom(output_folder: str, kbuban: str, revision: str, krevision: str) -> bool:
    """BOMの更新を行う"""
    bom_file = os.path.join(output_folder, "BOM_Parameter.json")
    if os.path.exists(bom_file):
        with open(bom_file, 'r', encoding='utf-8') as file:
            bom_data = file.read()
        if confirm_action("BOM更新") and post_data_to_api(BOM_UPDATE_URL, bom_data):
            log_file = run_recovery_tool(kbuban, revision, krevision, output_folder)
            extract_parameters_from_log(log_file, output_folder)
            return True
    return False

def update_caabatch(output_folder: str, kbuban: str, revision: str, krevision: str) -> bool:
    """CAABatchの更新を行う"""
    caabatch_file = os.path.join(output_folder, "CAABatch_Parameter.json")
    if os.path.exists(caabatch_file):
        with open(caabatch_file, 'r', encoding='utf-8') as file:
            caabatch_data = json.load(file)
        if confirm_action("CAABatch更新") and post_data_to_api(CAABATCH_UPDATE_URL, caabatch_data):
            if wait_for_job_completion(caabatch_data["process_order"], kbuban):
                log_file = run_recovery_tool(kbuban, revision, krevision, output_folder)
                extract_parameters_from_log(log_file, output_folder)
                return True
    return False

def wait_for_job_completion(process_order: str, kbuban: str) -> bool:
    """ジョブが完了するのを待つ"""
    start_time = time.time()
    while time.time() - start_time < 180:  # 3分待つ
        job_status = fetch_data_from_api(f"job-queue/{process_order}", params={"seppenbuban": kbuban})
        if job_status['status'] == '3':
            print(f"ジョブ {process_order} が正常に完了しました。")
            return True
        elif job_status['status'] == '4':
            print(f"ジョブ {process_order} でエラーが発生しました。")
            return False
        time.sleep(10)  # 10秒ごとにステータスをチェック
    print(f"ジョブ {process_order} が3分以内に完了しませんでした。")
    return False

def handle_job_status(result: Dict[str, Any], kbuban: str, revision: str, krevision: str, output_folder: str) -> bool:
    """ジョブステータスに基づいて適切なアクションを実行する"""
    if result['needs_investigation'] == 1 or result['needs_detailed_review'] == 1:
        sys.exit("プロセスには調査が必要です。")
    elif result['job_not_completed_correctly'] == 1:
        return update_caabatch(output_folder, kbuban, revision, krevision)
    elif result['error_occurred_during_job'] == 1 and confirm_action("CAD再処理"):
        reprocess_cad(result['latest_process_order'], kbuban)
        return True
    return False

def check_and_delete_lock(process_order: str) -> None:
    """ジョブロックを確認し、存在する場合は削除する"""
    response = fetch_data_from_api(f"job-lock/{process_order}", params={"usr_id": USERNAME})
    if response:
        print(f"ジョブロックの変更前: {response}")
        delete_response = requests.delete(f"{API_BASE_URL}/job-lock/{process_order}", params={"usr_id": USERNAME})
        if delete_response.status_code == 200:
            print("ジョブロックの削除に成功しました。")
        else:
            delete_response.raise_for_status()
        print(f"ジョブロックの変更後: None")

def fetch_and_update_job_queue(process_order: str, status: str, host: Optional[str]) -> None:
    """ジョブキューを取得し、更新する"""
    job_queue = fetch_data_from_api(f"job-queue/{process_order}", params={"usr_id": USERNAME})
    if job_queue:
        log_job_queue_change(job_queue, status, host)
        update_response = requests.put(f"{API_BASE_URL}/job-queue/{process_order}", json={"status": status, "host": host}, params={"usr_id": USERNAME})
        if update_response.status_code == 200:
            print("ジョブキューの更新に成功しました。")
        else:
            update_response.raise_for_status()

def log_job_queue_change(job_queue: Dict[str, Any], new_status: str, new_host: Optional[str]) -> None:
    """ジョブキューの変更前後をログ出力する"""
    print(f"ジョブキューの変更前: {job_queue}")
    print(f"ジョブキューの変更後: {{'process_order': job_queue['process_order'], 'status': new_status, 'host': new_host}}")

def reprocess_cad(process_order: str, kbuban: str) -> None:
    """CADの再処理を行う"""
    check_and_delete_lock(process_order)
    fetch_and_update_job_queue(process_order, '1', None)
    print("CADの再処理が正常に完了しました。")
