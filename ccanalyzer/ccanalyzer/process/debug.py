# encoding=utf-8
# 创建时间：2024-05-01 13:12

# ---------------- | 功能 | --------------------
# 工具类
# ---------------------------------------------

class Degbug:

    """ 打印字典  """
    @staticmethod
    def print_dict_tree(d, indent=0):
        if isinstance(d, dict):  # 如果是字典，遍历键值对
            for key, value in d.items():
                print('    ' * indent + str(key) + ':', end=' ')
                if isinstance(value, (dict, list)):  # 如果值是字典或列表，递归调用
                    print()
                    Degbug.print_dict_tree(value, indent + 1)
                else:
                    print(value)  # 打印简单类型的值
        elif isinstance(d, list):  # 如果是列表，遍历每个元素
            for i, item in enumerate(d):
                print('    ' * indent + f'[{i}]:', end=' ')
                if isinstance(item, (dict, list)):  # 如果元素是字典或列表，递归调用
                    print()
                    Degbug.print_dict_tree(item, indent + 1)
                else:
                    print(item)  # 打印简单类型的值