import sys
import os
from datetime import datetime
from typing import Optional, Dict

from utils.api import fetch_data_from_api
from utils.log import extract_parameters_from_log, run_recovery_tool
from utils.update import handle_job_status, fetch_recovery_job_status, update_bom, update_caabatch

USERNAME = os.getenv('USERNAME')

def fetch_part_info(seppenbuban: str) -> Optional[Dict[str, str]]:
    """指定された部品番号の部品情報を取得する"""
    return fetch_data_from_api(f"part/{seppenbuban}", params={"usr_id": USERNAME})

def main(seppenbuban: str) -> None:
    part_info = fetch_part_info(seppenbuban)
    if not part_info:
        sys.exit(f"SEPPENBUBAN: {seppenbuban}の部品情報が見つかりません。")

    kbuban, revision, krevision = part_info['kbuban'], part_info['revision'], part_info['krevision']

    output_folder = seppenbuban
    log_file = run_recovery_tool(kbuban, revision, krevision, output_folder)
    output_folder = extract_parameters_from_log(log_file, output_folder)

    if update_bom(output_folder, kbuban, revision, krevision):
        bom_parameter_file = os.path.join(output_folder, "BOM_Parameter.json")
        if os.path.exists(bom_parameter_file):
            new_bom_filename = f"BOM_Parameter_{datetime.now().strftime('%Y%m%d%H%M%S')}.json"
            os.rename(bom_parameter_file, os.path.join(output_folder, new_bom_filename))

    if os.path.exists(os.path.join(output_folder, "CAABatch_Parameter.json")):
        job_status = fetch_recovery_job_status(seppenbuban)
        if handle_job_status(job_status, kbuban, revision, krevision, output_folder):
            caa_batch_parameter_file = os.path.join(output_folder, "CAABatch_Parameter.json")
            if os.path.exists(caa_batch_parameter_file):
                new_caa_filename = f"CAABatch_Parameter_{datetime.now().strftime('%Y%m%d%H%M%S')}.json"
                os.rename(caa_batch_parameter_file, os.path.join(output_folder, new_caa_filename))
            if os.path.exists(bom_parameter_file):
                new_bom_filename = f"BOM_Parameter_{datetime.now().strftime('%Y%m%d%H%M%S')}.json"
                os.rename(bom_parameter_file, os.path.join(output_folder, new_bom_filename))

    print(f"kbuban: {kbuban}, revision: {revision}, krevision: {krevision}")
    print(f"ログファイル: {log_file}")
    print(f"出力フォルダ: {output_folder}")

if __name__ == "__main__":
    print('Hell World!')
    if len(sys.argv) < 2:
        print("引数1（SEPPENBUBAN）が必要です。")
        sys.exit(1)

    main(sys.argv[1])
