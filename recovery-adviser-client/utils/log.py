import re
import os
from datetime import datetime
import subprocess
from typing import Optional

RECOVERY_TOOL_EXE = 'bomRecoveryCheckTool.exe'

def extract_parameters_from_log(log_file: str, output_folder: str) -> str:
    """ログファイルからパラメータを抽出し、指定されたフォルダに保存する"""
    with open(log_file, 'r', encoding='utf-8') as file:
        log_content = file.read()

    save_parameter(extract_parameter(log_content, 'BOM'), output_folder, "BOM_Parameter.json")
    save_parameter(extract_parameter(log_content, 'CAABatch'), output_folder, "CAABatch_Parameter.json")

    return output_folder

def extract_parameter(log_content: str, parameter_type: str) -> Optional[str]:
    """ログファイルの内容から指定されたパラメータを抽出する"""
    pattern = re.compile(rf'---------- {parameter_type} Parameter ----------\n(.*?)\n---------- {parameter_type} Parameter ----------', re.DOTALL)
    match = pattern.search(log_content)
    return match.group(1) if match else None

def save_parameter(parameter: Optional[str], output_folder: str, filename: str) -> None:
    """指定されたパラメータをファイルに保存する"""
    if parameter:
        filepath = os.path.join(output_folder, filename)
        with open(filepath, 'w', encoding='utf-8') as file:
            file.write(parameter)
        print(f"{filename}を保存しました。")

def run_recovery_tool(kbuban: str, revision: str, krevision: str, output_folder: str) -> str:
    """リカバリーツールを実行し、ログファイルを出力する"""
    current_time = datetime.now().strftime("%Y%m%d%H%M%S")
    log_file = os.path.join(output_folder, f"{current_time}_bomRecoveryCheckTool_output.log")
    exe_path = os.path.join(os.path.dirname(sys.executable), RECOVERY_TOOL_EXE)
    exe_command = [exe_path, kbuban, revision, krevision]

    with open(log_file, "w") as log_output:
        subprocess.run(exe_command, stdout=log_output, stderr=subprocess.STDOUT)
    
    return log_file
