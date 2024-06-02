def confirm_action(action_type: str) -> bool:
    """指定されたアクションを実行するかどうかをユーザーに確認する"""
    while True:
        user_input = input(f"{action_type}を続行しますか？ (yes/no): ").lower()
        if user_input in ["yes", "no"]:
            return user_input == "yes"
        print("無効な入力です。'yes'または'no'を入力してください。")
