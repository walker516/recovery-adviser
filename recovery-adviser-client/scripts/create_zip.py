import zipfile
import os

def create_zip(zip_name, source_dir):
    with zipfile.ZipFile(zip_name, 'w', zipfile.ZIP_DEFLATED) as zipf:
        for root, _, files in os.walk(source_dir):
            for file in files:
                file_path = os.path.join(root, file)
                arcname = os.path.relpath(file_path, start=source_dir)
                zipf.write(file_path, arcname)

if __name__ == "__main__":
    source_directory = "dist"
    output_zip_file = os.path.join("dist", "RecoveryAdviser.zip")
    create_zip(output_zip_file, source_directory)
    print(f"{output_zip_file} が作成されました。")
