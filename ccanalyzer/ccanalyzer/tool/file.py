# encoding=utf-8
# 创建时间：2024-05-05 15:41

# 创建唯一文件名
import os
import uuid

# 创建唯一文件名
def create_unique_file(directory, file_type):
    # 定义文件类型和扩展名的映射
    extension_map = {
        'Python': '.py',
        'C': '.c',
        'CPP': '.cpp',
        'Java': '.java',
        'Go': '.go',
        'HTML': '.html',
        'CSS': '.css',
        'JavaScript': '.js'
    }
    # 获取扩展名，如果没有找到，默认为 '.txt'
    file_extension = extension_map.get(file_type.lower(), '.txt')

    # 生成唯一文件名
    file_name = f"{uuid.uuid4().hex}{file_extension}"

    return os.path.join(directory, file_name)

# 将数据写入文件
def write_to_file(file_path, content):
    with open(file_path, 'w', encoding='utf-8') as file:
        file.write(content)

# 删除文件
def delete_file(file_path):
    if os.path.exists(file_path):
        os.remove(file_path)
        print(f"{file_path} has been deleted")

# 根据文件扩展名判断编程语言
def detect_language(file_name):
    file_extension = file_name.split('.')[-1].lower()
    language_map = {
        'py': 'Python',
        'java': 'Java',
        'c': 'C',
        'h': 'CPP',
        'hpp': 'CPP',
        'cpp': 'CPP',
        'js': 'JavaScript',
        'html': 'HTML',
        'css': 'CSS',
        'go': 'Go'
        # 添加其他文件扩展名和对应的编程语言
    }
    return language_map.get(file_extension, 'Unknown')


def get_directory_tree(path, depth=0, max_depth=10):
    # 转换为绝对路径
    path = os.path.abspath(path)

    print(path)

    # 初始化树结构的根节点
    tree = {'label': os.path.basename(path), 'children': []}

    try:
        # 列出给定目录路径下的所有文件和目录
        for item in os.listdir(path):
            # 获取当前项目的完整路径
            full_path = os.path.join(path, item)

            # 忽略符号链接
            if os.path.islink(full_path):
                continue

            # 如果项目是一个目录
            if os.path.isdir(full_path):
                # 控制最大递归深度以避免无限递归
                if depth < max_depth:
                    tree['children'].append(get_directory_tree(full_path, depth + 1, max_depth))

            # 如果项目是一个文件
            else:
                tree['children'].append({'label': item, 'value': full_path})

    except PermissionError:
        pass

    return tree


""" 打印字典  """
def print_dict_tree(d, indent=0):
    if isinstance(d, dict):  # 如果是字典，遍历键值对
        for key, value in d.items():
            print('    ' * indent + str(key) + ':', end=' ')
            if isinstance(value, (dict, list)):  # 如果值是字典或列表，递归调用
                print()
                print_dict_tree(value, indent + 1)
            else:
                print(value)  # 打印简单类型的值
    elif isinstance(d, list):  # 如果是列表，遍历每个元素
        for i, item in enumerate(d):
            print('    ' * indent + f'[{i}]:', end=' ')
            if isinstance(item, (dict, list)):  # 如果元素是字典或列表，递归调用
                print()
                print_dict_tree(item, indent + 1)
            else:
                print(item)  # 打印简单类型的值


# 检查路径前缀：
def is_path_under_base(path, base_path):
    abs_path = os.path.abspath(path)
    abs_base_path = os.path.abspath(base_path)

    # 比较路径前缀
    return abs_path.startswith(abs_base_path)
